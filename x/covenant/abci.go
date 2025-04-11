package covenant

import (
	"encoding/hex"
	"fmt"
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

	for _, chain := range supportedBtcChains {
		clog.Greenf("[x/covenant] [ABCI] EndBlocker, chain: %+v", chain)
		handleEnqueuedEvents(ctx, neededKeepers, chain)
		handleSwitchPhase(ctx, neededKeepers, chain)
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
func handleSwitchPhase(ctx sdk.Context, nk *neededKeeper, chain nexus.ChainName) {
	expiredEvmSessions, expiredRedeemSessions := findExpiredEvmSessionsAndRenewable(ctx, nk.keeper, nk.protocol, nk.chains, chain)

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

	if redeemSession.CurrentPhase != exported.Preparing {
		return fmt.Errorf("redeem session is not in preparing phase")
	}

	protocols := nk.protocol.FindProtocolInfoByCustodianGroupUID(ctx, [][]byte{group.UID.Bytes()})
	if len(protocols) != 1 {
		return fmt.Errorf("not found protocol")
	}

	protocol := protocols[0]

	evmSessions := make(map[string]*types.ExpiredEvmSession)

	for _, chain := range protocol.MinorAddresses {
		s, ok := evmSessions[chain.ChainName.String()]
		if !ok {
			s = &types.ExpiredEvmSession{
				CustodianGroupUID: protocol.CustodianGroupUID,
				Chain:             chain.ChainName,
				Sequence:          redeemSession.Sequence,
				CurrentPhase:      exported.Preparing,
				Tokens:            []string{},
			}
		}
		s.Tokens = append(s.Tokens, protocol.Symbol)
		evmSessions[chain.ChainName.String()] = s
	}

	for _, session := range evmSessions {
		success := utils.RunCached(ctx, nk.keeper, func(ctx sdk.Context) (bool, error) {
			switchPhaseForEvmChain(ctx, nk.chains, nk.multisig, session, exported.Preparing)
			return true, nil
		})
		_ = success
	}

	nk.keeper.SetUtxoSnapshot(ctx, utxos)
	nk.keeper.SetSwitchingForRedeemSession(ctx, utxos.CustodianGroupUID /*, keyID*/)
	return nil
}

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
	err = ck.SetRedeemSession(ctx, &chainRedeemSession)
	if err != nil {
		ctx.Logger().Error("[handleSwitchedPhaseConfirmed] failed to set redeem session for chain %s", event.Chain, err)
		return err
	} else {
		ctx.Logger().Info(fmt.Sprintf("[handleSwitchedPhaseConfirmed] set redeem session %+v for chain %s", chainRedeemSession, event.Chain.String()))
	}
	allChains := nk.nexus.GetChains(ctx)
	//Store the slower chains which have old session or phase
	//If this array is empty, all evm chains have the same session and phase, we can switch the phase
	slowerChains := []nexus.Chain{}
	//Check if all evm chains have the same session and phase
	for _, c := range allChains {
		if c.Name.String() != event.Chain.String() && chainsTypes.IsEvmChain(c.Name) {
			ck, err := nk.chains.ForChain(ctx, c.Name)
			if err != nil {
				ctx.Logger().Error(fmt.Sprintf("[handleSwitchedPhaseConfirmed] failed to get chain keeper for chain %s", c.Name.String()), err)
				return err
			}
			redeemSession, ok := ck.GetRedeemSession(ctx, switchPhaseEvent.CustodianGroupUID)
			if !ok {
				ctx.Logger().Info(fmt.Sprintf("[handleSwitchedPhaseConfirmed] not found redeem session with custodian group uid %s for chain %s", hex.EncodeToString(switchPhaseEvent.CustodianGroupUID.Bytes()), c.Name.String()))
				slowerChains = append(slowerChains, c)
			} else if redeemSession.Sequence < switchPhaseEvent.Sequence ||
				(redeemSession.Sequence == switchPhaseEvent.Sequence && redeemSession.CurrentPhase < switchPhaseEvent.ToPhase) {
				ctx.Logger().Info(fmt.Sprintf("[handleSwitchedPhaseConfirmed] slower chain %s with session %++v", c.Name.String(), redeemSession))
				slowerChains = append(slowerChains, c)
			} else {
				ctx.Logger().Info(fmt.Sprintf("[handleSwitchedPhaseConfirmed] chain %s is already switch to phase %+v", c.Name.String(), chainRedeemSession))
			}
		}
	}
	if len(slowerChains) == 0 {
		ctx.Logger().Info("[handleSwitchedPhaseConfirmed] all evm chains have the same session and phase, we can switch the phase")
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
			err := nk.keeper.UpdatePreparingToExecuting(ctx, switchPhaseEvent.CustodianGroupUID)
			if err != nil {
				return err
			}

			err = signAllPendingRedeemCommands(ctx, nk, chain, switchPhaseEvent.CustodianGroupUID)
			if err != nil {
				return err
			}
		}
	} else {
		ctx.Logger().Info(fmt.Sprintf("[handleSwitchedPhaseConfirmed] there are %d slower chains, we need to handle them", len(slowerChains)))
		//TODO: handle the slower chains
	}

	return nil
}

func signAllPendingRedeemCommands(
	ctx sdk.Context,
	nk *neededKeeper,
	chain nexus.ChainName,
	custodianGroupUID chains.Hash,
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

	psbt, err := aggregatePsbtFromCommandBatch(ctx, nk.scalar, nk.chains, chain, commandBatch, group)
	if err != nil {
		return err
	}

	if err := nk.keeper.SignPsbt(
		ctx,
		commandBatch.GetKeyID(),
		[]exported.Psbt{psbt},
		chainsTypes.ModuleName,
		chain,
		types.NewSigMetadata(types.SigCommand, chain, commandBatch.GetID()),
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
	s types.ScalarnetKeeper,
	b types.BaseKeeper,
	chainName nexus.ChainName,
	commandBatch chainsTypes.CommandBatch,
	group *exported.CustodianGroup) (exported.Psbt, error) {
	multiPayload := commandBatch.GetExtraData()

	params := types.RedeemTokenPayload{}

	visited := map[string]bool{}
	inputs := []goutils.PreviousStakingUTXO{}
	outputs := []goutils.UnstakingOutput{}

	for _, payload := range multiPayload {
		err := params.AbiUnpack(payload)
		if err != nil {
			return nil, err
		}

		outputs = append(outputs, goutils.UnstakingOutput{
			Amount:        params.Amount,
			LockingScript: params.LockingScript,
		})

		for _, utxo := range params.Utxos {
			key := fmt.Sprintf("%x:%d", utxo.TxID, utxo.Vout)
			if visited[key] {
				continue
			}
			visited[key] = true
			txHash, err := chainhash.NewHashFromStr(strings.TrimPrefix(utxo.TxID.Hex(), "0x"))
			if err != nil {
				return nil, err
			}
			inputs = append(inputs, goutils.PreviousStakingUTXO{
				OutPoint: goutils.OutPoint{
					Txid: [32]byte(txHash.CloneBytes()),
					Vout: utxo.Vout,
				},
				Amount: utxo.AmountInSats,
				Script: group.BitcoinPubkey,
			})
		}
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

	psbt, err := vault.BuildCustodianOnlyUnstakingTx(
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
		Msg("switchPhaseForEvmChain")
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
func findExpiredEvmSessionsAndRenewable(ctx sdk.Context, k types.Keeper, pk types.ProtocolKeeper, c types.BaseKeeper, chain nexus.ChainName) (map[string]*types.ExpiredEvmSession, map[string]*types.RedeemSession) {
	result := map[string]*types.ExpiredEvmSession{}
	expiredGroups := [][]byte{}
	expiredSessions := map[string]*types.RedeemSession{}
	groups, ok := k.GetAllCustodianGroups(ctx)
	if !ok {
		return result, expiredSessions
	}
	currentHeight := ctx.BlockHeight()

	ck, err := c.ForChain(ctx, chain)
	if err != nil {
		panic(err)
	}

	for _, group := range groups {
		redeemSession, ok := k.GetRedeemSession(ctx, group.UID)
		if !ok {
			continue
		}
		if redeemSession.CurrentPhase == exported.Preparing && redeemSession.PhaseExpiredAt <= uint64(currentHeight) {
			// We switch the expired session to the executing phase if there are pending redeem commands
			// Otherwise, we extend current prepering phase
			if hasPendingRedeemCommands(ctx, k, ck) {
				expiredGroups = append(expiredGroups, group.UID.Bytes())
				expiredSessions[group.UID.Hex()] = redeemSession
			} else {
				err := k.RenewRedeemSession(ctx, group.UID)
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
	for _, protocol := range protocols {
		redeemSession, ok := expiredSessions[protocol.CustodianGroupUID.Hex()]
		if !ok {
			continue
		}
		for _, chain := range protocol.MinorAddresses {
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
			evmSession.Tokens = append(evmSession.Tokens, protocol.Symbol)
			result[chain.ChainName.String()] = evmSession
		}
	}
	return result, expiredSessions
}

func hasPendingRedeemCommands(ctx sdk.Context, k types.Keeper, ck chainsTypes.ChainKeeper) bool {
	batch := ck.GetLatestBtcPoolingBatch(ctx)
	if batch == nil || len(batch.GetCommandIDs()) == 0 {
		return false
	}
	extraData := batch.GetExtraData()
	// extra data is the array of payload of each redeem command, check contract for the payload
	// the payload of each reserve redeem utxo command doesn't contain the command ID,
	// so we cannot detect which reserve redeem utxo command is sent to the evm
	// extraDataIncludesAllReserveCommandsInTheRedeemSession := true

	// because all of reserve redeem utxo commands must be confirmed in one batch, we can just check if all command can be taken by the command ID
	for _, payload := range extraData {
		redeemTokenPayload := types.RedeemTokenPayload{}
		err := redeemTokenPayload.AbiUnpack(payload)
		if err != nil {
			panic(err)
		}
		//todo: check if the command is sent to the evm
		command := k.GetReserveUTXOCommandByID(ctx, redeemTokenPayload.RequestId[:])
		if command.Is(types.StandaloneCommandStatusNonExistent) {
			return true
			// extraDataIncludesAllReserveCommandsInTheRedeemSession = false
			// break
		}
	}

	return false
	// if !extraDataIncludesAllReserveCommandsInTheRedeemSession {
	// 	// renew redeem session
	// 	err := k.RenewRedeemSession(ctx, group.UID.Bytes())
	// 	if err != nil {
	// 		log.Error().
	// 			Err(err).
	// 			Str("CustodianGroupUID", group.UID.Hex()).
	// 			Msg("failed to renew redeem session")
	// 		panic(err)
	// 	}
	// }
}
