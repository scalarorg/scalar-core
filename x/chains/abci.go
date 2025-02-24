package chains

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/chains/types"
	mexported "github.com/scalarorg/scalar-core/x/multisig/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	pexported "github.com/scalarorg/scalar-core/x/protocol/exported"
	scalarnet "github.com/scalarorg/scalar-core/x/scalarnet/exported"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(sdk.Context, abci.RequestBeginBlock, types.BaseKeeper) {}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper, p types.ProtocolKeeper) ([]abci.ValidatorUpdate, error) {
	clog.Yellow("Chains ABCI ENDBLOCKER")
	handleConfirmedEvents(ctx, bk, n, m, p)
	handleMessages(ctx, bk, n, m)

	return nil, nil
}

func handleConfirmedEvents(ctx sdk.Context, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper, p types.ProtocolKeeper) {

	// This will handle all chains except Scalarnet.
	for _, chain := range slices.Filter(n.GetChains(ctx), types.IsSupportedChain) {
		handleConfirmedEventsForChain(ctx, chain, bk, n, m, p)
	}
}

func handleConfirmedEventsForChain(ctx sdk.Context, chain nexus.Chain, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper, p types.ProtocolKeeper) {
	ck := funcs.Must(bk.ForChain(ctx, chain.Name))
	queue := ck.GetConfirmedEventQueue(ctx)
	endBlockerLimit := ck.GetParams(ctx).EndBlockerLimit

	var events []types.Event
	var event types.Event
	for int64(len(events)) < endBlockerLimit && queue.Dequeue(&event) {
		events = append(events, event)
	}

	clog.Yellow("handleConfirmedEventsForChain", "chain", chain.Name.String(), "events", events)

	for _, event := range events {
		success := utils.RunCached(ctx, bk, func(ctx sdk.Context) (bool, error) {
			if err := handleConfirmedEvent(ctx, event, bk, n, m, p); err != nil {
				ck.Logger(ctx).Debug(fmt.Sprintf("failed handling event: %s", err.Error()),
					"chain", chain.Name.String(),
					"eventID", event.GetID(),
				)
				clog.Magentaf("[x/chains] [ABCI]-handleConfirmedEventsForChain %++v of type %T failed with error: %+v", event, event.GetEvent(), err)
				return false, err
			}

			ck.Logger(ctx).Debug("completed handling event",
				"chain", chain.Name.String(),
				"eventID", event.GetID(),
			)

			return true, nil
		})

		if !success {
			funcs.MustNoErr(ck.SetEventFailed(ctx, event.GetID()))
			continue
		}

		funcs.MustNoErr(ck.SetEventCompleted(ctx, event.GetID()))
	}
}

func handleConfirmedEvent(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper, p types.ProtocolKeeper) error {
	if err := validateEvent(ctx, event, bk, n); err != nil {
		return err
	}
	switch event.GetEvent().(type) {
	case *types.Event_SourceTxConfirmationEvent:
		return handleSourceConfirmationEvent(ctx, event, bk, n, m)
	case *types.Event_ContractCallWithToken:
		return handleContractCallWithToken(ctx, event, bk, n, m, p)
	case *types.Event_TokenSent:
		return handleTokenSent(ctx, event, bk, n)
	case *types.Event_Transfer:
		return handleConfirmDeposit(ctx, event, bk, n)
	case *types.Event_TokenDeployed:
		return handleTokenDeployed(ctx, event, bk, n)
	case *types.Event_MultisigOperatorshipTransferred:
		return handleMultisigTransferKey(ctx, event, bk, n, m)
	// TODO: add other event types here

	default:
		panic(fmt.Errorf("unsupported event type %T", event))
	}
}

func validateEvent(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus) error {
	var destinationChainName nexus.ChainName
	var contractAddress string
	switch event := event.GetEvent().(type) {
	case *types.Event_SourceTxConfirmationEvent:
		destinationChainName = event.SourceTxConfirmationEvent.DestinationChain
		contractAddress = event.SourceTxConfirmationEvent.DestinationContractAddress
	case *types.Event_ContractCallWithToken:
		destinationChainName = event.ContractCallWithToken.DestinationChain
		contractAddress = event.ContractCallWithToken.ContractAddress
	case *types.Event_TokenSent:
		destinationChainName = event.TokenSent.DestinationChain
	case *types.Event_Transfer, *types.Event_TokenDeployed,
		*types.Event_MultisigOperatorshipTransferred:
		// skip checks for non-gateway tx event
		return nil
	default:
		panic(fmt.Errorf("unsupported event type %T", event))
	}

	// skip if destination chain is not registered
	destinationChain, ok := n.GetChain(ctx, destinationChainName)
	if !ok {
		return fmt.Errorf("destination chain not found")
	}

	// skip if destination chain is not activated
	if !n.IsChainActivated(ctx, destinationChain) {
		return fmt.Errorf("destination chain de-activated")
	}

	// TODO: Here is validate the contract address for EVM, need to handle more general cases
	if len(contractAddress) != 0 && !common.IsHexAddress(contractAddress) {
		return fmt.Errorf("invalid contract address")
	}

	// skip further destination chain keeper checks if it is not an evm chain
	if !destinationChain.IsFrom(types.ModuleName) {
		return nil
	}
	// Todo: Verify destination's gatewayAddress if destination is an evm chain
	// destinationCk, err := bk.ForChain(ctx, destinationChainName)
	// if err != nil {
	// 	return fmt.Errorf("destination chain not EVM-compatible")
	// }

	// // skip if destination chain has not got gateway set yet
	// if _, ok := destinationCk.GetGatewayAddress(ctx); !ok {
	// 	return fmt.Errorf("destination chain gateway not deployed yet")
	// }

	return nil
}

func handleSourceConfirmationEvent(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper) error {
	e := event.GetSourceTxConfirmationEvent()
	if e == nil {
		panic(fmt.Errorf("event is nil"))
	}

	return setMessageToNexus(ctx, n, event, nil)
}

func handleTokenSent(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus) error {
	e := event.GetEvent().(*types.Event_TokenSent).TokenSent
	if e == nil {
		panic(fmt.Errorf("event is nil"))
	}

	sourceChain := funcs.MustOk(n.GetChain(ctx, event.Chain))
	destinationChain := funcs.MustOk(n.GetChain(ctx, e.DestinationChain))
	sourceCk := funcs.Must(bk.ForChain(ctx, sourceChain.Name))
	// Find token by symbol, which user request to transfer on the EVM
	token := sourceCk.GetERC20TokenBySymbol(ctx, e.Asset.Denom)
	if !token.Is(types.Confirmed) {
		return fmt.Errorf("token with symbol %s not confirmed on source chain", e.Asset.Denom)
	}
	//Get asset string from found token
	//asset.Name is unique in the whole scalar network
	asset := token.GetAsset()

	// check erc20 token status if destination is an evm chain
	if destinationCk, err := bk.ForChain(ctx, destinationChain.Name); err == nil {
		if token := destinationCk.GetERC20TokenByAsset(ctx, asset); !token.Is(types.Confirmed) {
			return fmt.Errorf("token with asset %s not confirmed on destination chain", e.Asset.Denom)
		}
	}

	recipient := nexus.CrossChainAddress{Chain: destinationChain, Address: e.DestinationAddress}
	amount := sdk.NewCoin(asset, sdk.Int(e.Asset.Amount))
	transferID, err := n.EnqueueCrossChainTransfer(ctx, sourceChain, common.Hash(event.TxID), recipient, amount)
	if err != nil {
		return sdkerrors.Wrap(err, "failed enqueuing transfer for event")
	}
	bk.Logger(ctx).Debug(fmt.Sprintf("enqueued transfer for event from chain %s", sourceChain.Name),
		"chain", destinationChain.Name,
		"eventID", event.GetID(),
		"transferID", transferID.String(),
	)
	clog.Magentaf("[x/chains] [ABCI]-Emited EventTokenSent")
	events.Emit(ctx, &types.EventTokenSent{
		Chain:              event.Chain,
		EventID:            event.GetID(),
		TransferID:         transferID,
		CommandID:          event.TxID.Hex(),
		Sender:             e.Sender,
		DestinationChain:   e.DestinationChain,
		DestinationAddress: e.DestinationAddress,
		Asset:              amount,
	})

	return nil
}

func handleContractCallWithToken(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper, p types.ProtocolKeeper) error {
	e := event.GetContractCallWithToken()
	if e == nil {
		panic(fmt.Errorf("event is nil"))
	}

	sourceChain := funcs.MustOk(n.GetChain(ctx, event.Chain))
	destinationChain := funcs.MustOk(n.GetChain(ctx, e.DestinationChain))

	sourceCk := funcs.Must(bk.ForChain(ctx, sourceChain.Name))
	token := sourceCk.GetERC20TokenBySymbol(ctx, e.Symbol)
	if !token.Is(types.Confirmed) {
		return fmt.Errorf("token with symbol %s not confirmed on source chain", e.Symbol)
	}
	asset := token.GetAsset()

	if err := n.RateLimitTransfer(ctx, sourceChain.Name, sdk.NewCoin(asset, sdk.Int(e.Amount)), nexus.TransferDirectionFrom); err != nil {
		return err
	}

	switch destinationChain.Module {
	case types.ModuleName:
		if types.IsEvmChain(destinationChain.Name) {
			return handleContractCallWithTokenToEVM(ctx, event, bk, n, m, sourceChain.Name, destinationChain.Name, asset)
		} else if types.IsBitcoinChain(destinationChain.Name) {
			return handleContractCallWithTokenToBTC(ctx, event, bk, n, p, sourceChain.Name, destinationChain.Name, asset)
		}
		return nil
	default:
		coin := sdk.NewCoin(asset, sdk.Int(e.Amount))
		// set as general message in nexus, so the dest module can handle the message
		return setMessageToNexus(ctx, n, event, &coin)
	}
}

func handleContractCallWithTokenToBTC(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus, p types.ProtocolKeeper, sourceChain, destinationChain nexus.ChainName, asset string) error {
	e := event.GetContractCallWithToken()
	if e == nil {
		panic(fmt.Errorf("event is nil"))
	}

	destinationCk := funcs.Must(bk.ForChain(ctx, destinationChain))

	destinationToken := destinationCk.GetERC20TokenByAsset(ctx, asset)
	if !destinationToken.Is(types.Confirmed) {
		return fmt.Errorf("token with asset %s not confirmed on destination chain", e.Symbol)
	}

	if !common.IsHexAddress(e.ContractAddress) {
		return fmt.Errorf("invalid contract address %s", e.ContractAddress)
	}

	coin := sdk.NewCoin(asset, sdk.Int(e.Amount))

	if err := n.RateLimitTransfer(ctx, destinationChain, coin, nexus.TransferDirectionTo); err != nil {
		return err
	}
	// keyId, ok := m.GetCurrentKeyID(ctx, destinationChain)
	// if !ok {
	// 	keyId = multisigexported.KeyID(destinationChain)
	// }
	protocolInfo, err := p.FindProtocolByExternalSymbol(ctx, destinationChain, sourceChain, e.Symbol)
	if err != nil {
		return err
	}

	var keyId mexported.KeyID
	switch protocolInfo.LiquidityModel {
	case pexported.LIQUIDITY_MODEL_POOL:
		clog.Yellowf("[abci/chains] Pooling protocol: %+v", protocolInfo)
		keyId = protocolInfo.KeyID
	case pexported.LIQUIDITY_MODEL_UPC:
		buffer := event.TxID[:]
		buffer = append(buffer, protocolInfo.CustodiansPubkey...)
		keyId = mexported.KeyID(hex.EncodeToString(buffer))
		clog.Yellowf("[abci/chains] Transactional protocol: %+v, keyId: %+v, txID: %s", protocolInfo, keyId, hex.EncodeToString(event.TxID[:]))
	}

	cmd := types.NewApproveContractCallWithMintCommandWithPayload(
		funcs.MustOk(destinationCk.GetChainID(ctx)),
		keyId,
		sourceChain,
		event.TxID,
		event.Index,
		*e,
		e.Amount,
		destinationToken.GetDetails().Symbol,
		event.GetContractCallWithToken().Payload,
	)

	clog.Magentaf("[abci/chains] created %s command for event: %+v", cmd.Type, cmd)

	funcs.MustNoErr(destinationCk.EnqueueCommand(ctx, cmd))

	bk.Logger(ctx).Info(fmt.Sprintf("created %s command for event", cmd.Type),
		"chain", destinationChain,
		"eventID", event.GetID(),
		"commandID", cmd.ID.Hex(),
	)

	approvedEvent := &types.EventContractCallWithMintApproved{
		Chain:            event.Chain,
		EventID:          event.GetID(),
		CommandID:        cmd.ID,
		Sender:           e.Sender.Hex(),
		DestinationChain: e.DestinationChain,
		ContractAddress:  e.ContractAddress,
		PayloadHash:      e.PayloadHash,
		Asset:            coin,
	}

	clog.Yellowf("[abci/chains] emitted EventContractCallWithMintApproved event for event: %v", approvedEvent)

	events.Emit(ctx, approvedEvent)

	return nil
}

func handleContractCallWithTokenToEVM(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus, multisig types.MultisigKeeper, sourceChain, destinationChain nexus.ChainName, asset string) error {
	e := event.GetContractCallWithToken()
	if e == nil {
		panic(fmt.Errorf("event is nil"))
	}

	destinationCk := funcs.Must(bk.ForChain(ctx, destinationChain))

	destinationToken := destinationCk.GetERC20TokenByAsset(ctx, asset)
	if !destinationToken.Is(types.Confirmed) {
		return fmt.Errorf("token with asset %s not confirmed on destination chain", e.Symbol)
	}

	if !common.IsHexAddress(e.ContractAddress) {
		return fmt.Errorf("invalid contract address %s", e.ContractAddress)
	}

	coin := sdk.NewCoin(asset, sdk.Int(e.Amount))

	if err := n.RateLimitTransfer(ctx, destinationChain, coin, nexus.TransferDirectionTo); err != nil {
		return err
	}

	cmd := types.NewApproveContractCallWithMintCommand(
		funcs.MustOk(destinationCk.GetChainID(ctx)),
		funcs.MustOk(multisig.GetCurrentKeyID(ctx, destinationChain)),
		sourceChain,
		event.TxID,
		event.Index,
		*e,
		e.Amount,
		destinationToken.GetDetails().Symbol,
	)
	funcs.MustNoErr(destinationCk.EnqueueCommand(ctx, cmd))
	bk.Logger(ctx).Debug(fmt.Sprintf("created %s command for event", cmd.Type),
		"chain", destinationChain,
		"eventID", event.GetID(),
		"commandID", cmd.ID.Hex(),
	)

	events.Emit(ctx, &types.EventContractCallWithMintApproved{
		Chain:            event.Chain,
		EventID:          event.GetID(),
		CommandID:        cmd.ID,
		Sender:           e.Sender.Hex(),
		DestinationChain: e.DestinationChain,
		ContractAddress:  e.ContractAddress,
		PayloadHash:      e.PayloadHash,
		Asset:            coin,
	})

	return nil
}

func setMessageToNexus(ctx sdk.Context, n types.Nexus, event types.Event, asset *sdk.Coin) error {

	sourceChain := funcs.MustOk(n.GetChain(ctx, event.Chain))

	var message nexus.GeneralMessage
	switch e := event.GetEvent().(type) {
	case *types.Event_SourceTxConfirmationEvent:
		sender := nexus.CrossChainAddress{
			Chain:   sourceChain,
			Address: e.SourceTxConfirmationEvent.Sender,
		}

		recipient := nexus.CrossChainAddress{
			Chain:   funcs.MustOk(n.GetChain(ctx, e.SourceTxConfirmationEvent.DestinationChain)),
			Address: e.SourceTxConfirmationEvent.DestinationContractAddress,
		}

		// message = nexus.NewGeneralMessage(
		// 	string(event.GetID()),
		// 	sender,
		// 	recipient,
		// 	e.SourceTxConfirmationEvent.PayloadHash.Bytes(),
		// 	event.TxID.Bytes(),
		// 	event.Index,
		// 	nil,
		// )
		message = nexus.NewGeneralMessageWithPayload(
			string(event.GetID()),
			sender,
			recipient,
			e.SourceTxConfirmationEvent.PayloadHash.Bytes(),
			event.TxID.Bytes(),
			event.Index,
			nil,
			e.SourceTxConfirmationEvent.Payload,
		)
	case *types.Event_ContractCallWithToken:
		if asset == nil {
			return fmt.Errorf("expect asset for ContractCallWithToken")
		}

		sender := nexus.CrossChainAddress{
			Chain:   sourceChain,
			Address: e.ContractCallWithToken.Sender.Hex(),
		}

		recipient := nexus.CrossChainAddress{
			Chain:   funcs.MustOk(n.GetChain(ctx, e.ContractCallWithToken.DestinationChain)),
			Address: e.ContractCallWithToken.ContractAddress,
		}

		message = nexus.NewGeneralMessageWithPayload(
			string(event.GetID()),
			sender,
			recipient,
			e.ContractCallWithToken.PayloadHash.Bytes(),
			event.TxID.Bytes(),
			event.Index,
			asset,
			event.GetContractCallWithToken().Payload,
		)
	// TODO: add other event types here

	default:
		return fmt.Errorf("unsupported event type %T", event)
	}

	if message.Recipient.Chain.Name.Equals(scalarnet.Scalarnet.Name) {
		return fmt.Errorf("%s is not a supported recipient", scalarnet.Scalarnet.Name)
	}

	return n.SetNewMessage(ctx, message)
}

func handleConfirmDeposit(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus) error {
	e := event.GetEvent().(*types.Event_Transfer).Transfer
	if e == nil {
		panic(fmt.Errorf("event is nil"))
	}

	chain := funcs.MustOk(n.GetChain(ctx, event.Chain))
	ck := funcs.Must(bk.ForChain(ctx, event.Chain))

	// get deposit address
	burnerInfo := ck.GetBurnerInfo(ctx, e.To)
	if burnerInfo == nil {
		return fmt.Errorf("no burner info found for address %s", e.To.Hex())
	}

	depositAddr := nexus.CrossChainAddress{Chain: chain, Address: e.To.Hex()}
	recipient, ok := n.GetRecipient(ctx, depositAddr)
	if !ok {
		return fmt.Errorf("cross-chain sender has no recipient %s", e.To.Hex())
	}

	// this check is only needed for historical reason.
	// if _, _, ok := ck.GetLegacyDeposit(ctx, event.TxID, burnerInfo.BurnerAddress); ok {
	// 	return fmt.Errorf("%s deposit %s-%s already exists", chain.Name.String(), event.TxID.Hex(), burnerInfo.BurnerAddress.Hex())
	// }

	amount := sdk.NewCoin(burnerInfo.Asset, sdk.NewIntFromBigInt(e.Amount.BigInt()))
	transferID, err := n.EnqueueForCrossChainTransfer(ctx, depositAddr, common.Hash(event.TxID), amount)
	if err != nil {
		return err
	}

	// set confirmed deposit
	erc20Deposit := types.ERC20Deposit{
		TxID:             event.TxID,
		LogIndex:         event.Index,
		Amount:           e.Amount,
		Asset:            burnerInfo.Asset,
		DestinationChain: burnerInfo.DestinationChain,
		BurnerAddress:    burnerInfo.BurnerAddress,
	}
	if _, _, ok := ck.GetDeposit(ctx, erc20Deposit.TxID, erc20Deposit.LogIndex); ok {
		panic(fmt.Errorf("%s deposit %s-%d already exists", chain.Name.String(), erc20Deposit.TxID.Hex(), erc20Deposit.LogIndex))
	}
	ck.SetDeposit(ctx, erc20Deposit, types.DepositStatus_Confirmed)
	ck.Logger(ctx).Info(fmt.Sprintf("confirmed deposit to %s with amount %s", e.To.Hex(), e.Amount),
		"chain", chain.Name,
		"depositAddress", depositAddr.Address,
		"eventID", event.GetID(),
		"txID", event.TxID.Hex(),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeDepositConfirmation,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyChain, chain.Name.String()),
			sdk.NewAttribute(types.AttributeKeySourceChain, chain.Name.String()),
			sdk.NewAttribute(types.AttributeKeyDestinationChain, recipient.Chain.Name.String()),
			sdk.NewAttribute(types.AttributeKeyDestinationAddress, recipient.Address),
			sdk.NewAttribute(types.AttributeKeyAmount, e.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyAsset, burnerInfo.Asset),
			sdk.NewAttribute(types.AttributeKeyDepositAddress, depositAddr.Address),
			sdk.NewAttribute(types.AttributeKeyTokenAddress, burnerInfo.TokenAddress.Hex()),
			sdk.NewAttribute(types.AttributeKeyTxID, event.TxID.Hex()),
			sdk.NewAttribute(types.AttributeKeyTransferID, transferID.String()),
			sdk.NewAttribute(types.AttributeKeyEventID, string(event.GetID())),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueConfirm),
		))

	return nil
}

func handleTokenDeployed(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus) error {
	fmt.Println("HandleTokenDeployed")
	e := event.GetEvent().(*types.Event_TokenDeployed).TokenDeployed
	if e == nil {
		panic(fmt.Errorf("event is nil"))
	}

	chain := funcs.MustOk(n.GetChain(ctx, event.Chain))
	ck := funcs.Must(bk.ForChain(ctx, event.Chain))

	token := ck.GetERC20TokenBySymbol(ctx, e.Symbol)
	if token.Is(types.NonExistent) {
		return fmt.Errorf("token %s does not exist", e.Symbol)
	}

	if token.GetAddress() != e.TokenAddress {
		return fmt.Errorf("token address %s does not match expected %s", e.TokenAddress.Hex(), token.GetAddress().Hex())
	}

	if err := token.ConfirmDeployment(); err != nil {
		return err
	}

	ck.Logger(ctx).Info(fmt.Sprintf("token %s deployment confirmed on chain %s", e.Symbol, chain.Name),
		"chain", chain.Name,
		"asset", token.GetAsset(),
		"eventID", event.GetID(),
		"txID", event.TxID.Hex(),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeTokenConfirmation,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyChain, chain.Name.String()),
			sdk.NewAttribute(types.AttributeKeyAsset, token.GetAsset()),
			sdk.NewAttribute(types.AttributeKeySymbol, token.GetDetails().Symbol),
			sdk.NewAttribute(types.AttributeKeyTokenAddress, token.GetAddress().Hex()),
			sdk.NewAttribute(types.AttributeKeyTxID, event.TxID.Hex()),
			sdk.NewAttribute(types.AttributeKeyEventID, string(event.GetID())),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueConfirm),
		))

	return nil
}

func handleMultisigTransferKey(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus, multisig types.MultisigKeeper) error {
	e := event.GetEvent().(*types.Event_MultisigOperatorshipTransferred).MultisigOperatorshipTransferred
	if e == nil {
		panic(fmt.Errorf("event is nil"))
	}

	chain := funcs.MustOk(n.GetChain(ctx, event.Chain))
	ck := funcs.Must(bk.ForChain(ctx, event.Chain))
	newAddresses := e.NewOperators
	newWeights := e.NewWeights
	newThreshold := e.NewThreshold

	nextKeyID, ok := multisig.GetNextKeyID(ctx, chain.Name)
	if !ok {
		return fmt.Errorf("next key for chain %s not found", chain.Name)
	}

	nextKey, ok := multisig.GetKey(ctx, nextKeyID)
	if !ok {
		return fmt.Errorf("key %s not found", nextKeyID)
	}

	expectedAddressWeights, expectedThreshold := types.ParseMultisigKey(nextKey)

	if len(newAddresses) != len(expectedAddressWeights) {
		return fmt.Errorf("new addresses length does not match, expected %d got %d", len(expectedAddressWeights), len(newAddresses))
	}

	addressSeen := make(map[string]bool)
	for i, newAddress := range newAddresses {
		newAddressHex := newAddress.Hex()
		if addressSeen[newAddressHex] {
			return fmt.Errorf("duplicate address in new addresses")
		}
		addressSeen[newAddressHex] = true

		expectedWeight, ok := expectedAddressWeights[newAddressHex]
		if !ok {
			return fmt.Errorf("new addresses do not match")
		}

		if !expectedWeight.Equal(newWeights[i]) {
			return fmt.Errorf("new weights do not match")
		}
	}

	if !newThreshold.Equal(expectedThreshold) {
		return fmt.Errorf("new threshold does not match, expected %s got %s", expectedThreshold.String(), newThreshold.String())
	}

	if err := multisig.RotateKey(ctx, chain.Name); err != nil {
		return err
	}

	ck.Logger(ctx).Info(fmt.Sprintf("successfully confirmed key transfer for chain %s", chain.Name),
		"chain", chain.Name,
		"txID", event.TxID.Hex(),
		"eventID", event.GetID(),
		"keyID", nextKeyID,
	)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeTransferKeyConfirmation,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(types.AttributeKeyChain, chain.Name.String()),
		sdk.NewAttribute(types.AttributeKeyTxID, event.TxID.Hex()),
		sdk.NewAttribute(types.AttributeKeyEventID, string(event.GetID())),
		sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueConfirm),
	))

	return nil
}

func handleMessages(ctx sdk.Context, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper) {
	allChains := n.GetChains(ctx)

	// This will handle all chains except Scalarnet.
	for _, chain := range slices.Filter(allChains, types.IsSupportedChain) {
		destCk := funcs.Must(bk.ForChain(ctx, chain.Name))
		if types.IsBitcoinChain(chain.Name) {
			bk.Logger(ctx).Info(fmt.Sprintf("Bitcoin network %s cannot receive general messages", chain.Name))
			continue
		}
		endBlockerLimit := destCk.GetParams(ctx).EndBlockerLimit
		msgs := n.GetProcessingMessages(ctx, chain.Name, endBlockerLimit)

		bk.Logger(ctx).Info(fmt.Sprintf("handling %d general messages", len(msgs)), types.AttributeKeyChain, chain.Name)

		for _, msg := range msgs {
			success := false
			_ = utils.RunCached(ctx, bk, func(ctx sdk.Context) (bool, error) {
				if err := validateMessage(ctx, destCk, n, m, chain, msg); err != nil {
					bk.Logger(ctx).Error(fmt.Sprintf("failed validating message: %s", err.Error()),
						types.AttributeKeyChain, msg.GetDestinationChain(),
						types.AttributeKeyMessageID, msg.ID,
					)
					return false, err
				}

				chainID := funcs.MustOk(destCk.GetChainID(ctx))

				var keyID mexported.KeyID
				if types.IsEvmChain(chain.Name) {
					keyID = funcs.MustOk(m.GetCurrentKeyID(ctx, chain.Name))
				} else {
					return false, fmt.Errorf("unsupported chain %v for key id", chain.Name)
				}

				switch msg.Type() {
				case nexus.TypeGeneralMessage:
					handleMessage(ctx, destCk, chainID, keyID, msg)
				case nexus.TypeGeneralMessageWithToken:
					if err := handleMessageWithToken(ctx, destCk, n, chainID, keyID, msg); err != nil {
						return false, err
					}
				default:
					panic(fmt.Sprintf("unrecognized message type %d", msg.Type()))
				}

				success = true
				return true, nil
			})

			if !success {
				destCk.Logger(ctx).Error("failed handling general message",
					types.AttributeKeyChain, chain.Name.String(),
					types.AttributeKeyMessageID, msg.ID,
				)

				events.Emit(ctx, &types.ContractCallFailed{
					Chain:     chain.Name,
					MessageID: msg.ID,
				})

				funcs.MustNoErr(n.SetMessageFailed(ctx, msg.ID))

				continue
			}

			funcs.MustNoErr(n.SetMessageExecuted(ctx, msg.ID))
		}
	}
}

func validateMessage(ctx sdk.Context, ck types.ChainKeeper, n types.Nexus, m types.MultisigKeeper, chain nexus.Chain, msg nexus.GeneralMessage) error {
	// TODO refactor to do these checks earlier so we don't fail in the end blocker

	if !utils.ValidateChainID(chain.Name.String()) {
		return fmt.Errorf("invalid chain id")
	}

	if !n.IsChainActivated(ctx, chain) {
		return fmt.Errorf("destination chain de-activated")
	}

	//Check gateway address is set for evm chain
	if types.IsEvmChain(chain.Name) {
		if _, ok := ck.GetGatewayAddress(ctx); !ok {
			return fmt.Errorf("destination chain gateway for chain %v not deployed yet", chain.Name)
		}
	}

	if !common.IsHexAddress(msg.GetDestinationAddress()) {
		return fmt.Errorf("invalid contract address")
	}

	if msg.Sender.Address == "" {
		return fmt.Errorf("sender address is empty")
	}

	if msg.Recipient.Address == "" {
		return fmt.Errorf("recipient address is empty")
	}

	if msg.PayloadHash == nil {
		return fmt.Errorf("payload hash is empty")
	}

	switch msg.Type() {
	case nexus.TypeGeneralMessage:
		return nil
	case nexus.TypeGeneralMessageWithToken:
		// TODO: handle multiple assets on btc: brc20, ordinals, runes, etc.
		return nil
	default:
		return fmt.Errorf("unrecognized message type %d", msg.Type())
	}
}

func handleMessageWithToken(ctx sdk.Context, ck types.ChainKeeper, n types.Nexus, chainID sdk.Int, keyID mexported.KeyID, msg nexus.GeneralMessage) error {
	token := ck.GetERC20TokenByAsset(ctx, msg.Asset.GetDenom())

	if err := n.RateLimitTransfer(ctx, msg.GetDestinationChain(), *msg.Asset, nexus.TransferDirectionTo); err != nil {
		return err
	}

	cmd := types.NewApproveContractCallWithMintGeneric(chainID, keyID, common.BytesToHash(msg.SourceTxID), msg.SourceTxIndex, msg, token.GetDetails().Symbol)
	funcs.MustNoErr(ck.EnqueueCommand(ctx, cmd))

	events.Emit(ctx, &types.EventContractCallWithMintApproved{
		Chain:            msg.GetSourceChain(),
		EventID:          types.EventID(msg.ID),
		CommandID:        cmd.ID,
		Sender:           msg.GetSourceAddress(),
		DestinationChain: msg.GetDestinationChain(),
		ContractAddress:  msg.GetDestinationAddress(),
		PayloadHash:      types.Hash(common.BytesToHash(msg.PayloadHash)),
		Asset:            *msg.Asset,
	})

	ck.Logger(ctx).Debug("completed handling general message with token",
		types.AttributeKeyChain, msg.GetDestinationChain(),
		types.AttributeKeyMessageID, msg.ID,
		types.AttributeKeyCommandsID, cmd.ID,
	)

	return nil
}

func handleMessage(ctx sdk.Context, ck types.ChainKeeper, chainID sdk.Int, keyID mexported.KeyID, msg nexus.GeneralMessage) {
	params := &types.ApproveContractCallCommandParams{
		ContractAddress:  common.HexToAddress(msg.GetDestinationAddress()),
		PayloadHash:      common.BytesToHash(msg.PayloadHash),
		SourceTxID:       common.BytesToHash(msg.SourceTxID),
		SourceChain:      msg.GetSourceChain(),
		Sender:           msg.GetSourceAddress(),
		SourceEventIndex: msg.SourceTxIndex,
		Payload:          msg.Payload,
	}

	clog.Yellow("[abci/chains/handleMessage] ==== COMMAND PARAMS ====")
	clog.Yellowf("chainID: %+v", chainID)
	clog.Yellowf("keyID: %+x", keyID)
	clog.Yellowf("msgID: %+x", msg.ID)
	clog.Yellowf("params: %+v", params)

	cmd := types.NewApproveContractCallCommandGeneric(chainID, keyID, msg.ID, params)
	funcs.MustNoErr(ck.EnqueueCommand(ctx, cmd))

	destCallApproved := &types.ContractCallApproved{
		Chain:            msg.GetSourceChain(),
		EventID:          types.EventID(msg.ID),
		CommandID:        cmd.ID,
		Sender:           msg.GetSourceAddress(),
		DestinationChain: msg.GetDestinationChain(),
		ContractAddress:  msg.GetDestinationAddress(),
		PayloadHash:      types.Hash(common.BytesToHash(msg.PayloadHash)),
	}

	clog.Redf("[abci/chains] DestCallApproved: %+v", destCallApproved)
	clog.Redf("[abci/chains] commandID: %+x", cmd.ID)

	events.Emit(ctx, destCallApproved)

	ck.Logger(ctx).Debug("completed handling general message",
		types.AttributeKeyChain, msg.GetDestinationChain(),
		types.AttributeKeyMessageID, msg.ID,
		types.AttributeKeyCommandsID, cmd.ID,
	)
}
