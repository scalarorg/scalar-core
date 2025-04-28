package covenant

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	goutils "github.com/scalarorg/bitcoin-vault/go-utils/types"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	protocol "github.com/scalarorg/scalar-core/x/protocol/exported"
	abci "github.com/tendermint/tendermint/abci/types"
)

type neededKeeper struct {
	keeper      types.Keeper
	chains      types.BaseKeeper
	protocol    types.ProtocolKeeper
	nexus       types.Nexus
	multisig    types.MultisigKeeper
	rewarder    types.Rewarder
	scalar      types.ScalarnetKeeper
	vote        types.Voter
	snapshotter types.Snapshotter
	slashing    types.SlashingKeeper
}

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, bk types.Keeper) {}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock,
	k types.Keeper,
	b types.BaseKeeper,
	pk types.ProtocolKeeper,
	n types.Nexus,
	multisig types.MultisigKeeper,
	rewarder types.Rewarder,
	s types.ScalarnetKeeper,
	vote types.Voter,
	snapshotter types.Snapshotter,
	slashing types.SlashingKeeper,
) ([]abci.ValidatorUpdate, error) {
	clog.Greenf("[x/covenant] [ABCI] EndBlocker, ctx.BlockHeight: %+v", ctx.BlockHeight())

	supportedBtcChains := []nexus.ChainName{}
	for _, chain := range n.GetChains(ctx) {
		if chainsTypes.IsBitcoinChain(chain.Name) {
			supportedBtcChains = append(supportedBtcChains, chain.Name)
		}
	}
	neededKeepers := &neededKeeper{
		keeper:      k,
		chains:      b,
		protocol:    pk,
		nexus:       n,
		multisig:    multisig,
		rewarder:    rewarder,
		scalar:      s,
		vote:        vote,
		snapshotter: snapshotter,
		slashing:    slashing,
	}

	for _, btcChain := range supportedBtcChains {
		clog.Greenf("[x/covenant] [ABCI] Start EndBlocker, chain: %+v", btcChain)
		handleEnqueuedEvents(ctx, neededKeepers, btcChain)
		handleSwitchPhase(ctx, neededKeepers, btcChain)
		clog.Greenf("[x/covenant] [ABCI] Finish EndBlocker, chain: %+v", btcChain)
	}

	handleSignings(ctx, k, rewarder)
	return nil, nil
}

func handleSignings(ctx sdk.Context, k types.Keeper, rewarder types.Rewarder) {
	// we handle sessions that'll expire on the next block,
	// to avoid waiting for an additional block
	for _, signing := range k.GetSigningSessionsByExpiry(ctx, ctx.BlockHeight()+1) {
		clog.Bluef("handleSignings, signing.GetID(): %+v", signing.GetID())
		_ = utils.RunCached(ctx, k, func(cachedCtx sdk.Context) ([]abci.ValidatorUpdate, error) {
			k.DeleteSigningSession(cachedCtx, signing.GetID())
			module := signing.GetModule()

			pool := rewarder.GetPool(cachedCtx, types.ModuleName)
			slices.ForEach(signing.GetMissingParticipants(), pool.ClearRewards)

			if signing.State != exported.Completed {
				events.Emit(cachedCtx, types.NewSigningPsbtExpired(signing.GetID()))
				k.Logger(cachedCtx).Info("signing session expired",
					"sig_id", signing.GetID(),
				)

				funcs.MustNoErr(k.GetCovenantRouter().GetHandler(module).HandleFailed(cachedCtx, signing.GetMetadata()))
				return nil, nil
			}

			// finalize the psbt
			err := FinalizeMultiPsbt(&signing.PsbtMultiSig)
			//serr := signing.PsbtMultiSig.Finalize()
			if err != nil {
				return nil, sdkerrors.Wrap(err, "failed to finalize psbt")
			}

			sig := funcs.Must(signing.Result())

			// TODO: must validate the signature in the submit signature request then release the rewards
			slices.ForEach(sig.GetParticipants(), func(p sdk.ValAddress) { funcs.MustNoErr(pool.ReleaseRewards(p)) })

			if err := k.GetCovenantRouter().GetHandler(module).HandleCompleted(cachedCtx, &sig, signing.GetMetadata()); err != nil {
				return nil, sdkerrors.Wrap(err, "failed to handle completed signature")
			}

			for _, p := range sig.GetMultiPsbt() {
				clog.Greenf("CovenantHandler: HandleCompleted, Psbts: %x", p.Bytes())
			}

			for _, tx := range sig.GetFinalizedTxs() {
				clog.Greenf("CovenantHandler: HandleCompleted, FinalizedTx: %x", tx)
			}

			events.Emit(cachedCtx, types.NewSigningPsbtCompleted(signing.GetID()))
			k.Logger(cachedCtx).Info("signing session completed",
				"sig_id", signing.GetID(),
				"key_id", sig.GetKeyID(),
				"module", module,
			)

			return nil, nil
		})
	}
}

func FinalizeMultiPsbt(p *types.PsbtMultiSig) error {
	var tapScriptSigsMapByEachPsbt = make([]map[string]*exported.TapScriptSigsMap, len(p.MultiPsbt))

	// collect the map for each psbt
	// ParticipantListTapScriptSigs = {
	// "Alice": [sigOfPsbt1, sigOfPsbt2, sigOfPsbt3],
	// "Bob": [sigOfPsbt1, sigOfPsbt2, sigOfPsbt3],
	// "Charlie": [sigOfPsbt1, sigOfPsbt2, sigOfPsbt3],
	//}
	// => output: [
	//    map[Alice:[sigOfPsbt1] Bob:[sigOfPsbt1] Charlie:[sigOfPsbt1]],
	//    map[Alice:[sigOfPsbt2] Bob:[sigOfPsbt2] Charlie:[sigOfPsbt2]],
	//    map[Alice:[sigOfPsbt3] Bob:[sigOfPsbt3] Charlie:[sigOfPsbt3]],
	// ]

	for party, listOfEachParty := range p.ParticipantListTapScriptSigs {
		for index, sig := range listOfEachParty.Inner {
			if tapScriptSigsMapByEachPsbt[index] == nil {
				tapScriptSigsMapByEachPsbt[index] = make(map[string]*exported.TapScriptSigsMap)
			}
			tapScriptSigsMapByEachPsbt[index][party] = sig
		}
	}

	return processPsbt(p, tapScriptSigsMapByEachPsbt)
}

func processPsbt(p *types.PsbtMultiSig, tapScriptSigsMapByEachPsbt []map[string]*exported.TapScriptSigsMap) error {
	type result struct {
		index     int
		tx        []byte
		psbtBytes []byte
		err       error
	}

	resultChan := make(chan result, len(p.MultiPsbt))

	// Launch goroutines for each PSBT
	for index, psbt := range p.MultiPsbt {
		go func(idx int, psbtData []byte, tapScriptSigsMap map[string]*exported.TapScriptSigsMap) {
			psbtBytes := psbtData
			var err error

			// Process tap script signatures
			for _, m := range tapScriptSigsMap {
				raw := m.ToRaw()
				psbtBytes, err = vault.AggregateTapScriptSigs(psbtBytes, raw)
				if err != nil {
					resultChan <- result{idx, nil, nil, err}
					return
				}
			}

			clog.Greenf("CovenantHandler: Finalize, Psbt: %x", psbtBytes)

			// Finalize PSBT and extract transaction
			tx, err := vault.FinalizePsbtAndExtractTx(psbtBytes)
			if err != nil {
				clog.Redf("CovenantHandler: Finalize, Error: %s", err)
				resultChan <- result{idx, nil, nil, err}
				return
			}

			resultChan <- result{idx, tx, psbtBytes, nil}
		}(index, psbt, tapScriptSigsMapByEachPsbt[index])
	}

	// Collect results
	for i := 0; i < len(p.MultiPsbt); i++ {
		res := <-resultChan
		if res.err != nil {
			return res.err
		}
		p.FinalizedTxs[res.index] = res.tx
		p.MultiPsbt[res.index] = res.psbtBytes
	}

	return nil
}

// Hande switch phase from Prepaing to Executing
func handleSwitchPhase(ctx sdk.Context, nk *neededKeeper, btcChain nexus.ChainName) {
	expiredEvmSessions, expiredRedeemSessions := findExpiredEvmSessionsAndRenewable(ctx, nk, nk.protocol, nk.chains, btcChain)
	if len(expiredEvmSessions) > 0 || len(expiredRedeemSessions) > 0 {
		log.Info().
			Str("Chain", btcChain.String()).
			Int("expiredEvmSessions", len(expiredEvmSessions)).
			Int("expiredRedeemSessions", len(expiredRedeemSessions)).
			Msg("[x/covenant] [handleSwitchPhase] [Found expired evm sessions]")
	}
	for _, evmSession := range expiredEvmSessions {
		success := utils.RunCached(ctx, nk.keeper, func(ctx sdk.Context) (bool, error) {
			switchPhaseForEvmChain(ctx, nk.chains, nk.multisig, evmSession, exported.Executing)
			return true, nil
		})
		_ = success
	}

	for _, redeemSession := range expiredRedeemSessions {
		log.Info().Msgf("turn on the flag isSwitching for redeem session %s", redeemSession.CustodianGroupUID.Hex())
		nk.keeper.SetSwitchingForRedeemSession(ctx, redeemSession.CustodianGroupUID)
	}
}

func handleEnqueuedEvents(
	ctx sdk.Context,
	nk *neededKeeper,
	chain nexus.ChainName,
) {
	queue := nk.keeper.GetEventsQueue(ctx)
	endBlockerLimit := 100 // TODO: move to the module.params

	var events []types.Event
	var event types.Event
	// Note: this ensures the blockchain is not frozen by processing all events in the queue
	for len(events) < endBlockerLimit && queue.Dequeue(&event) {
		events = append(events, event)
	}

	for _, event := range events {
		success := utils.RunCached(ctx, nk.keeper, func(ctx sdk.Context) (bool, error) {
			err := handleEnqueueEvent(ctx, &event, nk, chain)
			if err != nil {
				nk.keeper.Logger(ctx).Debug(fmt.Sprintf("failed handling event: %s", err.Error()),
					"chain", event.Chain.String(),
				)
				clog.Redf("[x/covenent] [ABCI]-handle event %++v of type %T failed with error: %+v", event, event.GetEvent(), err)
				return false, err
			}

			nk.keeper.Logger(ctx).Debug("completed handling event",
				"chain", event.Chain.String(),
			)

			return true, nil
		})

		_ = success

		// if !success {
		// 	funcs.MustNoErr(ck.SetEventFailed(ctx, event.GetID()))
		// 	continue
		// }

		// funcs.MustNoErr(ck.SetEventCompleted(ctx, event.GetID()))
	}

}

func handleEnqueueEvent(
	ctx sdk.Context,
	event *types.Event,
	nk *neededKeeper,
	chain nexus.ChainName,
) error {
	// if err := validateEvent(ctx, event, bk, n); err != nil {
	// 	return err
	// }
	switch event.GetEvent().(type) {
	case *types.Event_RedeemTxsConfirmed:
		return handleRedeemTxsConfirmed(ctx, event, nk)
	case *types.Event_SwitchedPhaseConfirmed:
		return handleSwitchedPhaseConfirmed(ctx, event, nk, chain)
	case *types.Event_IntializeUtxoSnapshotCompleted:
		return handleInitializeUtxoSnapshotCompleted(ctx, event, nk.keeper)
	default:
		panic(fmt.Errorf("unsupported event type %T", event))
	}
}

func handleInitializeUtxoSnapshotCompleted(ctx sdk.Context, event *types.Event, k types.Keeper) error {
	ctx.Logger().Info("[x/covenant] [ABCI] start handling initialize utxo snapshot completed event")
	confirmedEvent, ok := event.GetEvent().(*types.Event_IntializeUtxoSnapshotCompleted)
	if !ok {
		return fmt.Errorf("invalid event type")
	}

	utxos := confirmedEvent.IntializeUtxoSnapshotCompleted.GetUtxoSnapshot()

	k.SetUtxoSnapshot(ctx, utxos)

	return nil

}

func handleRedeemTxsConfirmed(ctx sdk.Context, event *types.Event, nk *neededKeeper) error {
	// event := event.GetRedeemTxsConfirmed()
	// keyID := event.GetKeyID()
	// chain
	// txs := event.GetTxs()
	confirmedEvent, ok := event.GetEvent().(*types.Event_RedeemTxsConfirmed)
	if !ok {
		return fmt.Errorf("invalid event type")
	}

	utxos := confirmedEvent.RedeemTxsConfirmed.GetUtxoSnapshot()

	group, ok := nk.keeper.GetCustodianGroup(ctx, utxos.CustodianGroupUID)
	if !ok {
		return fmt.Errorf("not found custodian group")
	}

	redeemSession, ok := nk.keeper.GetRedeemSession(ctx, group.UID)
	if !ok {
		return fmt.Errorf("not found redeem session")
	}

	if redeemSession.CurrentPhase != exported.Executing {
		return fmt.Errorf("redeem session is not in executing phase")
	}

	protocols := nk.protocol.FindProtocolInfoByCustodianGroupUID(ctx, [][]byte{group.UID.Bytes()})
	if len(protocols) == 0 {
		return fmt.Errorf("[handleRedeemTxsConfirmed] protocol ")
	}

	evmSessions := make(map[string]*types.ExpiredEvmSession)
	for _, protocol := range protocols {
		for _, chain := range protocol.MinorAddresses {
			s, ok := evmSessions[chain.ChainName.String()]
			if !ok {
				s = &types.ExpiredEvmSession{
					CustodianGroupUID: protocol.CustodianGroupUID,
					Chain:             chain.ChainName,
					Sequence:          redeemSession.Sequence,
					CurrentPhase:      redeemSession.CurrentPhase,
					Tokens:            []string{},
				}
			}
			s.Tokens = append(s.Tokens, protocol.Symbol)
			evmSessions[chain.ChainName.String()] = s
		}
	}

	for _, session := range evmSessions {
		success := utils.RunCached(ctx, nk.keeper, func(ctx sdk.Context) (bool, error) {
			switchPhaseForEvmChain(ctx, nk.chains, nk.multisig, session, exported.Preparing)
			return true, nil
		})
		_ = success
	}
	currentUtxoSnapshot, present := nk.keeper.GetUtxoSnapshot(ctx, utxos.CustodianGroupUID)
	if present {
		//Check if new utxo don't contain any resrved utxo
		reservedUtxos := map[string]bool{}
		for _, utxo := range currentUtxoSnapshot.Utxos {
			if len(utxo.Reservations) > 0 {
				key := fmt.Sprintf("%s:%d", utxo.TxID.Hex(), utxo.Vout)
				reservedUtxos[key] = true
			}
		}
		for _, utxo := range utxos.Utxos {
			key := fmt.Sprintf("%s:%d", utxo.TxID.Hex(), utxo.Vout)
			if reservedUtxos[key] {
				log.Error().Str("ReservedUtxo", key).Err(fmt.Errorf("reserved utxo found"))
			}
		}
	}
	err := nk.keeper.SetUtxoSnapshot(ctx, utxos)
	if err != nil {
		return err
	}
	return nk.keeper.SetSwitchingForRedeemSession(ctx, utxos.CustodianGroupUID /*, keyID*/)
}

// Update the redeem session for the chain
// Check and update the redeem session for custodian group to the lowest phase of all evm chains
func handleSwitchedPhaseConfirmed(
	ctx sdk.Context,
	event *types.Event,
	nk *neededKeeper,
	chain nexus.ChainName,
) error {
	confirmedEvent, ok := event.GetEvent().(*types.Event_SwitchedPhaseConfirmed)
	if !ok {
		return fmt.Errorf("invalid event type")
	}
	ctx.Logger().Debug("[handleSwitchedPhaseConfirmed] start handling switch phase confirmed event")
	switchPhaseEvent := confirmedEvent.SwitchedPhaseConfirmed
	chainRedeemSession := chainsTypes.RedeemSession{
		CustodianGroupUID: switchPhaseEvent.CustodianGroupUID,
		Sequence:          switchPhaseEvent.Sequence,
		CurrentPhase:      switchPhaseEvent.ToPhase,
	}
	ck, err := nk.chains.ForChain(ctx, event.Chain)
	if err != nil {
		ctx.Logger().Error("[handleSwitchedPhaseConfirmed] failed to get chain keeper for chain %s", event.Chain, err)
		return err
	}
	existingRedeemSession, ok := nk.keeper.GetRedeemSession(ctx, switchPhaseEvent.CustodianGroupUID)
	if ok && !existingRedeemSession.IsSwitching {
		ctx.Logger().Info("[handleSwitchedPhaseConfirmed] redeem session is not in switching state, scalar received switch phase to Executing due to recovering mode, set it to switching for safety update to Execution")
		nk.keeper.SetSwitchingForRedeemSession(ctx, switchPhaseEvent.CustodianGroupUID)

	}
	err = ck.SetRedeemSession(ctx, &chainRedeemSession)
	if err != nil {
		ctx.Logger().Error("[handleSwitchedPhaseConfirmed] failed to set redeem session for chain %s", event.Chain, err)
		return err
	} else {
		ctx.Logger().Info(fmt.Sprintf("[handleSwitchedPhaseConfirmed] set redeem session %+v for chain %s", chainRedeemSession, event.Chain.String()))
	}
	// Get the highest and lowest redeem session for the custodian group
	// Result: highestRedeemSession, lowestRedeemSession are not nil because we have at least one evm chain
	highestRedeemSession, lowestRedeemSession, missing := getMinMaxRedeemSession(ctx, nk, switchPhaseEvent.CustodianGroupUID)
	ctx.Logger().Info(fmt.Sprintf("[handleSwitchedPhaseConfirmed] highest redeem session: %s, lowest redeem session: %s missing: %d",
		highestRedeemSession.ToString(), lowestRedeemSession.ToString(), missing))
	if missing > 0 {
		ctx.Logger().Info(fmt.Sprintf("[handleSwitchedPhaseConfirmed] missing %d redeem sessions. Waiting for them", missing))
		return nil
	}

	diff := highestRedeemSession.Cmp(lowestRedeemSession)
	if diff == 0 {
		if switchPhaseEvent.ToPhase == exported.Preparing {
			err := nk.keeper.UpdateExecutingToPreparing(ctx, switchPhaseEvent.CustodianGroupUID, switchPhaseEvent.Sequence)
			if err != nil {
				ctx.Logger().Error("[handleSwitchedPhaseConfirmed] failed to update executing to preparing", err)
				return err
			}

			// err = startInitializeUtxoEvent(ctx, k, b, n, v, snapshotter, slashing, switchPhaseEvent.CustodianGroupUID.Bytes())
			// if err != nil {
			// 	return err
			// }

			// TODO:
		} else if switchPhaseEvent.ToPhase == exported.Executing {
			err := nk.keeper.UpdatePreparingToExecuting(ctx, switchPhaseEvent.CustodianGroupUID, switchPhaseEvent.Sequence)
			if err != nil {
				return err
			}

			err = signAllPendingRedeemCommands(ctx, nk, chain, switchPhaseEvent.CustodianGroupUID, switchPhaseEvent.Sequence)
			if err != nil {
				return err
			}
		}
	} else if diff == 1 {
		//Set the redeem session to the lowest phase
		covRedeemSession, ok := nk.keeper.GetRedeemSession(ctx, switchPhaseEvent.CustodianGroupUID)
		if !ok {
			return fmt.Errorf("not found redeem session")
		}
		covRedeemSession.Sequence = lowestRedeemSession.Sequence
		covRedeemSession.CurrentPhase = lowestRedeemSession.CurrentPhase
		covRedeemSession.IsSwitching = true
		if covRedeemSession.PhaseExpiredAt <= uint64(ctx.BlockHeight()) {
			covRedeemSession.PhaseExpiredAt = uint64(ctx.BlockHeight()) + 1
		}
		nk.keeper.SetRedeemSession(ctx, covRedeemSession)
	} else if diff > 1 {
		//this must not happen
		ctx.Logger().Info("[handleSwitchedPhaseConfirmed] difference between highest and lowest redeem sessions is greater than 1")
		//Set the redeem session to the lowest phase
		covRedeemSession, ok := nk.keeper.GetRedeemSession(ctx, switchPhaseEvent.CustodianGroupUID)
		if !ok {
			return fmt.Errorf("not found redeem session")
		}
		covRedeemSession.Sequence = lowestRedeemSession.Sequence
		covRedeemSession.CurrentPhase = lowestRedeemSession.CurrentPhase
		covRedeemSession.IsSwitching = true
		if covRedeemSession.PhaseExpiredAt <= uint64(ctx.BlockHeight()) {
			covRedeemSession.PhaseExpiredAt = uint64(ctx.BlockHeight()) + 1
		}
		nk.keeper.SetRedeemSession(ctx, covRedeemSession)
	}
	return nil
}

// Get the highest and lowest redeem session for the custodian group
// Result: highestRedeemSession, lowestRedeemSession are not nil because we have at least one evm chain
// missing: the number of missing redeem sessions
func getMinMaxRedeemSession(ctx sdk.Context, nk *neededKeeper, custodianGroupUID chains.Hash) (*chainsTypes.RedeemSession, *chainsTypes.RedeemSession, int) {
	allChains := nk.nexus.GetChains(ctx)
	//Store the highest and lowest redeem session for the chain
	//If them are equal, we can switch the phase, otherwise we need to handle the slower chains
	var highestRedeemSession *chainsTypes.RedeemSession
	var lowestRedeemSession *chainsTypes.RedeemSession
	missing := 0
	//Loop through all chains and find the highest and lowest redeem session
	for _, c := range allChains {
		if chainsTypes.IsEvmChain(c.Name) {
			ck, err := nk.chains.ForChain(ctx, c.Name)
			if err != nil {
				log.Debug().
					Str("chain", c.Name.String()).
					Msg("[getMinMaxRedeemSession] failed to get chain keeper")
				missing++
				continue
			}
			redeemSession, ok := ck.GetRedeemSession(ctx, custodianGroupUID)
			if !ok {
				ctx.Logger().Info(fmt.Sprintf("[handleSwitchedPhaseConfirmed] not found redeem session with custodian group uid %s for chain %s", hex.EncodeToString(custodianGroupUID.Bytes()), c.Name.String()))
				//Missing redeem session
				missing++
				continue
			}
			if highestRedeemSession == nil || redeemSession.Cmp(highestRedeemSession) > 0 {
				highestRedeemSession = &redeemSession
			}
			if lowestRedeemSession == nil || redeemSession.Cmp(lowestRedeemSession) < 0 {
				lowestRedeemSession = &redeemSession
			}
		}
	}
	return highestRedeemSession, lowestRedeemSession, missing
}
func signAllPendingRedeemCommands(
	ctx sdk.Context,
	nk *neededKeeper,
	chain nexus.ChainName,
	custodianGroupUID chains.Hash,
	seq uint64,
) error {
	group, ok := nk.keeper.GetCustodianGroup(ctx, custodianGroupUID)
	if !ok {
		return fmt.Errorf("not found custodian group")
	}

	commandBatch, err := getCommandBatchToSign(ctx, nk.chains, chain, group.BitcoinPubkey)
	if err != nil {
		return err
	}

	if len(commandBatch.GetCommandIDs()) == 0 {
		return nil
	}

	psbt, err := aggregatePsbtFromCommandBatch(ctx, nk.keeper, nk.scalar, nk.chains, chain, commandBatch, group, seq)
	if err != nil {
		return err
	}

	if err := nk.keeper.SignPsbt(
		ctx,
		commandBatch.GetKeyID(),
		[]exported.Psbt{psbt},
		chainsTypes.ModuleName,
		chain,
		chainsTypes.NewSigMetadata(chainsTypes.SigCommand, chain, commandBatch.GetID()),
	); err != nil {
		return err
	}

	if !commandBatch.SetStatus(chainsTypes.BatchSigning) {
		return fmt.Errorf("failed setting status of command batch %s to be signing", hex.EncodeToString(commandBatch.GetID()))
	}

	clog.Yellowf("[keeper] [msg_server_sign_btc_commands] commandBatch: %+v", commandBatch)

	batchedCommandsIDHex := hex.EncodeToString(commandBatch.GetID())
	commandList := chainsTypes.CommandIDsToStrings(commandBatch.GetCommandIDs())
	for _, commandID := range commandList {
		nk.keeper.Logger(ctx).Info(
			fmt.Sprintf("signing command %s in batch %s for chain %s using key %s", commandID, batchedCommandsIDHex, chain, string(commandBatch.GetKeyID())),
			chainsTypes.AttributeKeyChain, chain,
			chainsTypes.AttributeKeyKeyID, string(commandBatch.GetKeyID()),
			"commandBatchID", batchedCommandsIDHex,
			"commandID", commandID,
		)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			chainsTypes.EventTypeSign,
			sdk.NewAttribute(sdk.AttributeKeyAction, chainsTypes.AttributeValueStart),
			sdk.NewAttribute(sdk.AttributeKeyModule, chainsTypes.ModuleName),
			sdk.NewAttribute(chainsTypes.AttributeKeyChain, chain.String()),
			sdk.NewAttribute(chainsTypes.AttributeKeyBatchedCommandsID, batchedCommandsIDHex),
			sdk.NewAttribute(chainsTypes.AttributeKeyCommandsIDs, strings.Join(commandList, ",")),
		),
	)

	return nil
}

func aggregatePsbtFromCommandBatch(
	ctx sdk.Context,
	cov types.Keeper,
	s types.ScalarnetKeeper,
	b types.BaseKeeper,
	chainName nexus.ChainName,
	commandBatch chainsTypes.CommandBatch,
	group *exported.CustodianGroup,
	seq uint64,
) (exported.Psbt, error) {
	clog.Yellow("[abci/covenant] [aggregatePsbtFromCommandBatch] start")
	multiPayload := commandBatch.GetExtraData()
	clog.Yellowf("[abci/covenant] [aggregatePsbtFromCommandBatch] multiPayload: %+v", multiPayload)
	params := types.RedeemTokenPayloadWithType{}

	inputs := []goutils.PreviousOutpoint{}
	outputs := []goutils.UnlockingOutput{}

	visited := map[string]bool{}

	utxoSnapshot, ok := cov.GetUtxoSnapshot(ctx, group.UID)
	if !ok {
		return nil, fmt.Errorf("[abci/covenant]: utxo snapshot not found")
	}

	for _, payload := range multiPayload {
		//First byte is the payload type
		err := params.AbiUnpack(payload)
		if err != nil {
			return nil, err
		}

		clog.Yellowf("[abci/covenant] [aggregatePsbtFromCommandBatch] payload: %+x", payload)

		outputs = append(outputs, goutils.UnlockingOutput{
			Amount:        params.Amount,
			LockingScript: params.LockingScript,
		})

		for _, utxo := range params.Utxos {
			key := fmt.Sprintf("%x:%d", utxo.TxID, utxo.Vout)
			if visited[key] {
				continue
			}
			visited[key] = true
			foundUtxo := utxoSnapshot.FindUtxo(utxo.TxID, utxo.Vout)
			if foundUtxo == nil {
				log.Error().Str("TxID", utxo.TxID.Hex()).
					Str("UTXOs", utxoSnapshot.ToString()).
					Uint32("Vout", utxo.Vout).Msg("[abci/covenant]: utxo not found in snapshot")
				return nil, fmt.Errorf("[abci/covenant]: utxo not found in snapshot")
			}

			txHash, err := chainhash.NewHashFromStr(strings.TrimPrefix(foundUtxo.TxID.Hex(), "0x"))
			if err != nil {
				return nil, err
			}
			inputs = append(inputs, goutils.PreviousOutpoint{
				OutPoint: goutils.OutPoint{
					Txid: [32]byte(txHash.CloneBytes()),
					Vout: foundUtxo.Vout,
				},
				Amount: foundUtxo.AmountInSats,
				Script: group.BitcoinPubkey,
			})
		}
	}

	var choosenUtxo *types.UTXO

	for _, utxo := range utxoSnapshot.Utxos {
		key := fmt.Sprintf("%x:%d", utxo.TxID, utxo.Vout)
		if visited[key] {
			continue
		}

		choosenUtxo = utxo
		break
	}

	if choosenUtxo != nil {
		txHash, err := chainhash.NewHashFromStr(strings.TrimPrefix(choosenUtxo.TxID.Hex(), "0x"))
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, goutils.PreviousOutpoint{
			OutPoint: goutils.OutPoint{
				Txid: [32]byte(txHash.CloneBytes()),
				Vout: choosenUtxo.Vout,
			},
			Amount: choosenUtxo.AmountInSats,
			Script: group.BitcoinPubkey,
		})
	}

	ck, err := b.ForChain(ctx, chainName)
	if err != nil {
		return nil, err
	}

	scalarnetParams := s.GetParams(ctx)

	tag := scalarnetParams.Tag
	version := scalarnetParams.Version

	// Note: because of we have multiple protocols in one signle redeem transaction so we hardcode the service tag to "pools"
	serviceTag := []byte("pools")
	network := ck.GetParams(ctx).NetworkKind
	custodianPubKeys := slices.Map(group.Custodians, func(c *exported.Custodian) goutils.PublicKey {
		pk := make([]byte, 33)
		copy(pk, c.BitcoinPubkey)
		return goutils.PublicKey(pk)
	})
	custodianQuorum := group.Quorum
	rbf := false
	feeRate := uint64(1)

	sort.Slice(inputs, func(i, j int) bool {
		txidCmp := bytes.Compare(inputs[i].OutPoint.Txid[:], inputs[j].OutPoint.Txid[:])
		if txidCmp != 0 {
			return txidCmp < 0
		}
		return inputs[i].OutPoint.Vout < inputs[j].OutPoint.Vout
	})

	sort.Slice(outputs, func(i, j int) bool {
		lockingScriptCmp := bytes.Compare(outputs[i].LockingScript[:], outputs[j].LockingScript[:])
		if lockingScriptCmp != 0 {
			return lockingScriptCmp < 0
		}
		return outputs[i].Amount < outputs[j].Amount
	})

	for _, input := range inputs {
		clog.Greenf("Input: %+v\n", input)
	}

	for _, output := range outputs {
		clog.Greenf("Output: %+v\n", output)
	}

	for _, pk := range custodianPubKeys {
		clog.Greenf("CustodianPubKey: %+x\n", pk)
	}

	clog.Greenf("CustodianQuorum: %+v\n", custodianQuorum)

	psbt, err := vault.BuildPoolingRedeemTx(
		tag,
		serviceTag,
		uint8(version),
		network,
		inputs,
		outputs,
		custodianPubKeys,
		uint8(custodianQuorum),
		rbf,
		feeRate,
		seq,
		group.UID.Bytes(),
	)
	return psbt, err
}

func getCommandBatchToSign(ctx sdk.Context, bk types.BaseKeeper, chain nexus.ChainName, scriptPubkey []byte) (chainsTypes.CommandBatch, error) {
	latest := bk.GetLatestCommandBatchForChain(ctx, chain)

	switch latest.GetStatus() {
	case chainsTypes.BatchSigning:
		return chainsTypes.CommandBatch{}, sdkerrors.Wrapf(chainsTypes.ErrSignCommandsInProgress, "command batch '%s'", hex.EncodeToString(latest.GetID()))
	case chainsTypes.BatchAborted:
		return latest, nil
	default:
		return bk.CreateNewBtcPoolingBatchToSign(ctx, chain, scriptPubkey)
	}
}

func switchPhaseForEvmChain(ctx sdk.Context,
	b types.BaseKeeper,
	multisig types.MultisigKeeper,
	evmSession *types.ExpiredEvmSession,
	newPhase exported.Phase,
) error {
	log.Info().
		Str("chain", evmSession.Chain.String()).
		Str("CustodianGroupUID", hex.EncodeToString(evmSession.CustodianGroupUID.Bytes())).
		Any("newPhase", newPhase).
		Msg("[x/covenant] [switchPhaseForEvmChain]")
	// Start signing session for reserve redeem utxos
	keyID, ok := multisig.GetCurrentKeyID(ctx, evmSession.Chain)
	if !ok {
		return fmt.Errorf("could not find key ID for '%s'", evmSession.Chain)
	}

	chainID := funcs.Must(evmSession.Chain.GetChainID())

	cmd := types.NewSwitchPhaseCommandWithExpiredSessioin(
		evmSession,
		*chainID,
		keyID,
		newPhase,
	)

	ck, err := b.ForChain(ctx, evmSession.Chain)
	if err != nil {
		return err
	}

	ck.EnqueueCommand(ctx, cmd)
	return nil
}

// TODO: review this method

// return map[chainName]ExpiredEvmSession
func findExpiredEvmSessionsAndRenewable(ctx sdk.Context, nk *neededKeeper, pk types.ProtocolKeeper, c types.BaseKeeper, btcChain nexus.ChainName) (map[string]*types.ExpiredEvmSession, map[string]*types.RedeemSession) {
	result := map[string]*types.ExpiredEvmSession{}
	expiredGroups := [][]byte{}
	expiredSessions := map[string]*types.RedeemSession{}
	groups, ok := nk.keeper.GetAllCustodianGroups(ctx)
	if !ok {
		return result, expiredSessions
	}
	currentHeight := ctx.BlockHeight()

	ck, err := c.ForChain(ctx, btcChain)
	if err != nil {
		panic(err)
	}

	for _, group := range groups {
		redeemSession, ok := nk.keeper.GetRedeemSession(ctx, group.UID)
		if !ok {
			log.Debug().
				Str("CustodianGroupUID", group.UID.Hex()).
				Msg("[x/covenant] [findExpiredEvmSessionsAndRenewable] [Not found redeem session]")
			continue
		}
		if redeemSession.CurrentPhase != exported.Preparing {
			//Switching phase on expired apply for preparing phase only
			continue
		}
		//Check if redeem session is switching we need to continue switch process
		//TODO: check if the redeem session is switching but all evm sessions are successfully switched
		if redeemSession.IsSwitching {
			log.Info().
				Str("CustodianGroupUID", group.UID.Hex()).
				Msg("[x/covenant] [Switching redeem session]")
			expiredGroups = append(expiredGroups, group.UID.Bytes())
			expiredSessions[group.UID.Hex()] = redeemSession
		} else if redeemSession.PhaseExpiredAt <= uint64(currentHeight) {
			log.Info().
				Str("CustodianGroupUID", group.UID.Hex()).
				Int64("currentHeight", currentHeight).
				Uint64("phaseExpiredAt", redeemSession.PhaseExpiredAt).
				Msg("[x/covenant] Found expired redeem session")
			// We switch the expired session to the executing phase if there are pending redeem commands
			// Otherwise, we extend current prepering phase
			if ck.HasBtcPoolingCommands(ctx, group.BitcoinPubkey) {
				log.Info().
					Str("CustodianGroupUID", group.UID.Hex()).
					Msg("[x/covenant][Switching redeem session to executing phase]")
				expiredGroups = append(expiredGroups, group.UID.Bytes())
				expiredSessions[group.UID.Hex()] = redeemSession
			} else {
				log.Info().
					Str("CustodianGroupUID", group.UID.Hex()).
					Msg("[x/covenant] [Renewing redeem session]")
				err := nk.keeper.RenewRedeemSession(ctx, group.UID)
				if err != nil {
					log.Error().
						Err(err).
						Str("CustodianGroupUID", group.UID.Hex()).
						Msg("failed to renew redeem session")
				}
			}
		}
	}

	protocols := pk.FindProtocolInfoByCustodianGroupUID(ctx, expiredGroups)
	result, err = findExpiredEvmRedeemSessions(ctx, protocols, expiredSessions, nk)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[x/covenant] [findExpiredEvmSessionsAndRenewable] [Failed to find expired evm sessions]")
	}
	return result, expiredSessions
}

func findExpiredEvmRedeemSessions(ctx sdk.Context, protocols []*protocol.ProtocolInfo, expiredSessions map[string]*types.RedeemSession,
	nk *neededKeeper) (map[string]*types.ExpiredEvmSession, error) {
	result := map[string]*types.ExpiredEvmSession{}
	for _, protocol := range protocols {
		redeemSession, ok := expiredSessions[protocol.CustodianGroupUID.Hex()]
		if !ok {
			continue
		}
		for _, chain := range protocol.MinorAddresses {
			if !chainsTypes.IsEvmChain(chain.ChainName) {
				continue
			}
			ck, err := nk.chains.ForChain(ctx, chain.ChainName)
			if err != nil {
				return nil, err
			}
			evmRedeemSession, ok := ck.GetRedeemSession(ctx, protocol.CustodianGroupUID)
			if !ok {
				log.Error().
					Str("chain", chain.ChainName.String()).
					Str("CustodianGroupUID", protocol.CustodianGroupUID.Hex()).
					Msg("[x/covenant] [findExpiredEvmRedeemSessions] [Not found evm redeem session]")
				continue
			}
			if evmRedeemSession.Sequence > redeemSession.Sequence ||
				(evmRedeemSession.Sequence == redeemSession.Sequence && evmRedeemSession.CurrentPhase > redeemSession.CurrentPhase) {
				log.Info().
					Str("CustodianGroupUID", protocol.CustodianGroupUID.Hex()).
					Str("Chain", chain.ChainName.String()).
					Any("evmRedeemSession", evmRedeemSession).
					Any("redeemSession", redeemSession).
					Msg("[x/covenant] [findExpiredEvmRedeemSessions] evm session is already switch. No need to request again")
				continue
			}
			evmSession, ok := result[chain.ChainName.String()]
			if !ok {
				evmSession = &types.ExpiredEvmSession{
					CustodianGroupUID: protocol.CustodianGroupUID,
					Chain:             chain.ChainName,
					Sequence:          redeemSession.Sequence,
					CurrentPhase:      exported.Preparing,
					Tokens:            []string{},
				}
			}
			log.Info().
				Str("CustodianGroupUID", protocol.CustodianGroupUID.Hex()).
				Str("Chain", chain.ChainName.String()).
				Any("evmRedeemSession", evmRedeemSession).
				Any("redeemSession", redeemSession).
				Msg("[x/covenant] [Found expired evm session]")
			evmSession.Tokens = append(evmSession.Tokens, protocol.Symbol)
			result[chain.ChainName.String()] = evmSession
		}
	}
	return result, nil
}

// func hasPendingRedeemCommands(ctx sdk.Context, k types.Keeper, ck chainsTypes.ChainKeeper) bool {
// 	batch := ck.GetLatestBtcPoolingBatch(ctx)
// 	if batch == nil || len(batch.GetCommandIDs()) == 0 {
// 		log.Debug().
// 			Msg("[x/covenant] [hasPendingRedeemCommands] [No pending redeem commands]")
// 		return false
// 	}
// 	extraData := batch.GetExtraData()
// 	// extra data is the array of payload of each redeem command, check contract for the payload
// 	// the payload of each reserve redeem utxo command doesn't contain the command ID,
// 	// so we cannot detect which reserve redeem utxo command is sent to the evm
// 	// extraDataIncludesAllReserveCommandsInTheRedeemSession := true

// 	// because all of reserve redeem utxo commands must be confirmed in one batch, we can just check if all command can be taken by the command ID
// 	for _, payload := range extraData {
// 		redeemTokenPayload := types.RedeemTokenPayload{}
// 		err := redeemTokenPayload.AbiUnpack(payload)
// 		if err != nil {
// 			panic(err)
// 		}
// 		log.Debug().
// 			Any("payload", payload).
// 			Msg("[x/covenant] [hasPendingRedeemCommands]")
// 		//todo: check if the command is sent to the evm
// 		command := k.GetReserveUTXOCommandByID(ctx, redeemTokenPayload.RequestId[:])
// 		if command.Is(types.StandaloneCommandStatusNonExistent) {
// 			return true
// 			// extraDataIncludesAllReserveCommandsInTheRedeemSession = false
// 			// break
// 		}
// 	}

// 	return false
// 	// if !extraDataIncludesAllReserveCommandsInTheRedeemSession {
// 	// 	// renew redeem session
// 	// 	err := k.RenewRedeemSession(ctx, group.UID.Bytes())
// 	// 	if err != nil {
// 	// 		log.Error().
// 	// 			Err(err).
// 	// 			Str("CustodianGroupUID", group.UID.Hex()).
// 	// 			Msg("failed to renew redeem session")
// 	// 		panic(err)
// 	// 	}
// 	// }
// }
