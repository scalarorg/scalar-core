package covenant

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
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

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, bk types.Keeper) {}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock,
	k types.Keeper,
	b types.BaseKeeper,
	pk types.ProtocolKeeper,
	nexus types.Nexus,
	multisig types.MultisigKeeper,
	rewarder types.Rewarder,
	s types.ScalarnetKeeper,
) ([]abci.ValidatorUpdate, error) {
	clog.Greenf("Covenant EndBlocker, ctx.BlockHeight: %+v", ctx.BlockHeight())
	handleSignings(ctx, k, rewarder)
	handleEnqueuedEvents(ctx, k, b, pk, nexus, multisig, s)
	handleSwitchPhase(ctx, k, b, pk, multisig)
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
func handleSwitchPhase(ctx sdk.Context, k types.Keeper, b types.BaseKeeper, pk types.ProtocolKeeper, multisig types.MultisigKeeper) {
	expiredEvmSessions, expiredRedeemSessions := findExpiredEvmSessions(ctx, k, pk)

	for _, evmSession := range expiredEvmSessions {
		success := utils.RunCached(ctx, k, func(ctx sdk.Context) (bool, error) {
			switchPhaseForEvmChain(ctx, k, b, multisig, evmSession, exported.Executing)
			return true, nil
		})
		_ = success
	}
	for _, redeemSession := range expiredRedeemSessions {
		k.SetSwitchingForRedeemSession(ctx, redeemSession.CustodianGroupUID[:])
	}
}

func handleEnqueuedEvents(
	ctx sdk.Context,
	k types.Keeper,
	b types.BaseKeeper,
	pk types.ProtocolKeeper,
	nexus types.Nexus,
	m types.MultisigKeeper,
	s types.ScalarnetKeeper,
) {
	queue := k.GetEventsQueue(ctx)
	endBlockerLimit := 100 // TODO: move to the module.params

	var events []types.Event
	var event types.Event
	// Note: this ensures the blockchain is not frozen by processing all events in the queue
	for len(events) < endBlockerLimit && queue.Dequeue(&event) {
		events = append(events, event)
	}

	for _, event := range events {
		success := utils.RunCached(ctx, k, func(ctx sdk.Context) (bool, error) {
			if err := handleEnqueueEvent(ctx, &event, k, b, pk, nexus, m, s); err != nil {
				k.Logger(ctx).Debug(fmt.Sprintf("failed handling event: %s", err.Error()),
					"chain", event.Chain.String(),
				)
				clog.Magentaf("[x/covenent] [ABCI]-handle event %++v of type %T failed with error: %+v", event, event.GetEvent(), err)
				return false, err
			}

			k.Logger(ctx).Debug("completed handling event",
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
	k types.Keeper,
	b types.BaseKeeper,
	pk types.ProtocolKeeper,
	nexus types.Nexus,
	m types.MultisigKeeper,
	s types.ScalarnetKeeper,
) error {
	// if err := validateEvent(ctx, event, bk, n); err != nil {
	// 	return err
	// }
	switch event.GetEvent().(type) {
	case *types.Event_RedeemTxsConfirmed:
		return handleRedeemTxsConfirmed(ctx, event, k, b, m, pk)
	case *types.Event_SwitchedPhaseConfirmed:
		return handleSwitchedPhaseConfirmed(ctx, event, k, b, nexus, s)
	default:
		panic(fmt.Errorf("unsupported event type %T", event))
	}
}

func handleRedeemTxsConfirmed(ctx sdk.Context, event *types.Event, k types.Keeper, b types.BaseKeeper, m types.MultisigKeeper, pk types.ProtocolKeeper) error {
	// event := event.GetRedeemTxsConfirmed()
	// keyID := event.GetKeyID()
	// chain
	// txs := event.GetTxs()
	confirmedEvent, ok := event.GetEvent().(*types.Event_RedeemTxsConfirmed)
	if !ok {
		return fmt.Errorf("invalid event type")
	}

	utxos := confirmedEvent.RedeemTxsConfirmed.GetUtxoSnapshot()

	group, ok := k.GetCustodianGroup(ctx, chains.Hash(utxos.CustodianGroupUID.Bytes()))
	if !ok {
		return fmt.Errorf("not found custodian group")
	}

	redeemSession, ok := k.GetRedeemSession(ctx, group.UID.Bytes())
	if !ok {
		return fmt.Errorf("not found redeem session")
	}

	if redeemSession.CurrentPhase != exported.Preparing {
		return fmt.Errorf("redeem session is not in preparing phase")
	}

	protocols := pk.FindProtocolInfoByCustodianGroupUID(ctx, [][]byte{group.UID.Bytes()})
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

	for _, s := range evmSessions {
		success := utils.RunCached(ctx, k, func(ctx sdk.Context) (bool, error) {
			switchPhaseForEvmChain(ctx, k, b, m, s, exported.Preparing)
			return true, nil
		})
		_ = success
	}

	k.SetUtxoSnapshot(ctx, utxos)
	k.SetSwitchingForRedeemSession(ctx, utxos.CustodianGroupUID[:] /*, keyID*/)
	return nil
}

func handleSwitchedPhaseConfirmed(
	ctx sdk.Context,
	event *types.Event,
	k types.Keeper,
	b types.BaseKeeper,
	n types.Nexus,
	s types.ScalarnetKeeper,
) error {
	confirmedEvent, ok := event.GetEvent().(*types.Event_SwitchedPhaseConfirmed)
	if !ok {
		return fmt.Errorf("invalid event type")
	}

	switchPhaseEvent := confirmedEvent.SwitchedPhaseConfirmed
	chainRedeemSession := chainsTypes.RedeemSession{
		CustodianGroupUID: switchPhaseEvent.CustodianGroupUID,
		Sequence:          switchPhaseEvent.Sequence,
		CurrentPhase:      switchPhaseEvent.ToPhase,
	}
	ck, err := b.ForChain(ctx, event.Chain)
	if err != nil {
		return err
	}
	ck.SetRedeemSession(ctx, &chainRedeemSession)
	allChains := n.GetChains(ctx)
	//Store the slower chains which have old session or phase
	//If this array is empty, all evm chains have the same session and phase, we can switch the phase
	slowerChains := []nexus.Chain{}
	//Check if all evm chains have the same session and phase
	for _, c := range allChains {
		if c.Name != event.Chain && chainsTypes.IsEvmChain(c.Name) {
			ck, err := b.ForChain(ctx, c.Name)
			if err != nil {
				return err
			}
			redeemSession, ok := ck.GetRedeemSession(ctx, switchPhaseEvent.CustodianGroupUID.Bytes())
			if !ok {
				ctx.Logger().Debug("not found redeem session for chain %s", c.Name)
				slowerChains = append(slowerChains, c)
			} else if redeemSession.Sequence < switchPhaseEvent.Sequence ||
				(redeemSession.Sequence == switchPhaseEvent.Sequence && redeemSession.CurrentPhase < switchPhaseEvent.ToPhase) {
				ctx.Logger().Debug("[handleSwitchedPhaseConfirmed] slower chain %s with session %++v", c.Name, redeemSession)
				slowerChains = append(slowerChains, c)
			}
		}
	}
	if len(slowerChains) == 0 {
		ctx.Logger().Debug("[handleSwitchedPhaseConfirmed] all evm chains have the same session and phase, we can switch the phase")
		if switchPhaseEvent.ToPhase == exported.Preparing {
			return k.UpdateExecutingToPreparing(ctx, switchPhaseEvent.CustodianGroupUID.Bytes())
		} else if switchPhaseEvent.ToPhase == exported.Executing {
			err := k.UpdatePreparingToExecuting(ctx, switchPhaseEvent.CustodianGroupUID.Bytes())
			if err != nil {
				return err
			}

			err = handleSignCommands(ctx, k, b, s, switchPhaseEvent.CustodianGroupUID.Bytes())
			if err != nil {
				return err
			}
		}
	} else {
		ctx.Logger().Debug("[handleSwitchedPhaseConfirmed] there are %s slower chains, we need to handle them", len(slowerChains))
		//TODO: handle the slower chains
	}

	return fmt.Errorf("invalid phase")
}

func handleSignCommands(
	ctx sdk.Context,
	k types.Keeper,
	b types.BaseKeeper,
	s types.ScalarnetKeeper,
	custodianGroupUID []byte,
) error {
	// TODO: Fix the chain
	mockChain := nexus.ChainName("bitcoin|4")

	group, ok := k.GetCustodianGroup(ctx, chains.Hash(custodianGroupUID))
	if !ok {
		return fmt.Errorf("not found custodian group")
	}

	commandBatch, err := getCommandBatchToSign(ctx, b, mockChain, group.BitcoinPubkey)
	if err != nil {
		return err
	}

	if len(commandBatch.GetCommandIDs()) == 0 {
		return nil
	}

	psbt, err := aggregatePsbtFromCommandBatch(ctx, s, b, mockChain, commandBatch, group)
	if err != nil {
		return err
	}

	if err := k.SignPsbt(
		ctx,
		commandBatch.GetKeyID(),
		[]exported.Psbt{psbt},
		chainsTypes.ModuleName,
		mockChain,
		types.NewSigMetadata(types.SigCommand, mockChain, commandBatch.GetID()),
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
		k.Logger(ctx).Info(
			fmt.Sprintf("signing command %s in batch %s for chain %s using key %s", commandID, batchedCommandsIDHex, mockChain, string(commandBatch.GetKeyID())),
			chainsTypes.AttributeKeyChain, mockChain,
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
			sdk.NewAttribute(chainsTypes.AttributeKeyChain, mockChain.String()),
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

	bytesType := funcs.Must(abi.NewType("bytes", "bytes", nil))
	uint256Type := funcs.Must(abi.NewType("uint256", "uint256", nil))
	uint256ArrayType := funcs.Must(abi.NewType("uint256[]", "uint256[]", nil))
	stringArrayType := funcs.Must(abi.NewType("string[]", "string[]", nil))

	arg := abi.Arguments{
		{Type: uint256Type},
		{Type: bytesType},
		{Type: stringArrayType},
		{Type: uint256ArrayType},
		{Type: uint256ArrayType},
	}

	visited := map[string]bool{}
	inputs := []goutils.PreviousStakingUTXO{}
	outputs := []goutils.UnstakingOutput{}

	for _, payload := range multiPayload {
		params, err := chainsTypes.StrictDecode(arg, payload)
		if err != nil {
			return nil, err
		}

		reqAmount := params[0].(*uint64)
		lockingScript := params[1].([]byte)

		outputs = append(outputs, goutils.UnstakingOutput{
			Amount:        *reqAmount,
			LockingScript: lockingScript,
		})

		txIds := params[2].([]string)
		vouts := params[3].([]*uint64)
		amountInSats := params[4].([]*uint64)

		for i, txId := range txIds {
			if visited[txId] {
				continue
			}
			visited[txId] = true
			txHash, err := chainhash.NewHashFromStr(txId)
			if err != nil {
				return nil, err
			}
			inputs = append(inputs, goutils.PreviousStakingUTXO{
				OutPoint: goutils.OutPoint{
					Txid: [32]byte(txHash.CloneBytes()),
					Vout: uint32(*vouts[i]),
				},
				Amount: *amountInSats[i],
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
	// TODO: fix me
	serviceTag := []byte("no-tag")
	network := ck.GetParams(ctx).NetworkKind
	custodianPubKeys := slices.Map(group.Custodians, func(c *exported.Custodian) goutils.PublicKey {
		pk := make([]byte, 33)
		copy(pk, c.BitcoinPubkey)
		return goutils.PublicKey(pk)
	})
	custodianQuorum := group.Quorum
	rbf := false
	feeRate := uint64(1)

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
	_ types.Keeper,
	b types.BaseKeeper,
	multisig types.MultisigKeeper,
	evmSession *types.ExpiredEvmSession,
	newPhase exported.Phase,
) error {
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
func findExpiredEvmSessions(ctx sdk.Context, k types.Keeper, pk types.ProtocolKeeper) (map[string]*types.ExpiredEvmSession, map[string]*types.RedeemSession) {
	result := map[string]*types.ExpiredEvmSession{}
	expiredGroups := [][]byte{}
	expiredSessions := map[string]*types.RedeemSession{}
	groups, ok := k.GetAllCustodianGroups(ctx)
	if !ok {
		return result, expiredSessions
	}
	currentHeight := ctx.BlockHeight()
	for _, group := range groups {
		redeemSession, ok := k.GetRedeemSession(ctx, group.UID.Bytes())
		if !ok {
			continue
		}
		if redeemSession.CurrentPhase == exported.Preparing && redeemSession.PhaseExpiredAt < uint64(currentHeight) {
			expiredGroups = append(expiredGroups, group.UID.Bytes())
			expiredSessions[group.UID.Hex()] = redeemSession
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
