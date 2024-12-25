package chains

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/chains/types"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	scalarnet "github.com/scalarorg/scalar-core/x/scalarnet/exported"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(sdk.Context, abci.RequestBeginBlock, types.BaseKeeper) {}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper) ([]abci.ValidatorUpdate, error) {
	clog.Yellow("Chains ABCI ENDBLOCKER")
	handleConfirmedEvents(ctx, bk, n, m)
	handleMessages(ctx, bk, n, m)

	return nil, nil
}

func handleConfirmedEvents(ctx sdk.Context, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper) {

	// This will handle all chains except Scalarnet.
	for _, chain := range slices.Filter(n.GetChains(ctx), types.IsSupportedChain) {
		handleConfirmedEventsForChain(ctx, chain, bk, n, m)
	}
}

func handleConfirmedEventsForChain(ctx sdk.Context, chain nexus.Chain, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper) {
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
			if err := handleConfirmedEvent(ctx, event, bk, n, m); err != nil {
				ck.Logger(ctx).Debug(fmt.Sprintf("failed handling event: %s", err.Error()),
					"chain", chain.Name.String(),
					"eventID", event.GetID(),
				)

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

func handleConfirmedEvent(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper) error {
	if err := validateEvent(ctx, event, bk, n); err != nil {
		return err
	}

	switch event.GetEvent().(type) {
	case *types.Event_SourceTxConfirmationEvent:
		return handleSourceConfirmationEvent(ctx, event, bk, n, m)

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

	return nil
}

func handleSourceConfirmationEvent(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper) error {
	e := event.GetSourceTxConfirmationEvent()
	if e == nil {
		panic(fmt.Errorf("event is nil"))
	}

	return setMessageToNexus(ctx, n, event, nil)
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

		message = nexus.NewGeneralMessage(
			string(event.GetID()),
			sender,
			recipient,
			e.SourceTxConfirmationEvent.PayloadHash.Bytes(),
			event.TxID.Bytes(),
			event.Index,
			nil,
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

func handleMessages(ctx sdk.Context, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper) {
	allChains := n.GetChains(ctx)

	// This will handle all chains except Scalarnet.
	for _, chain := range slices.Filter(allChains, types.IsSupportedChain) {
		destCk := funcs.Must(bk.ForChain(ctx, chain.Name))
		endBlockerLimit := destCk.GetParams(ctx).EndBlockerLimit
		msgs := n.GetProcessingMessages(ctx, chain.Name, endBlockerLimit)

		bk.Logger(ctx).Info(fmt.Sprintf("handling %d general messages", len(msgs)), types.AttributeKeyChain, chain.Name)

		for _, msg := range msgs {
			success := false
			_ = utils.RunCached(ctx, bk, func(ctx sdk.Context) (bool, error) {
				if err := validateMessage(ctx, destCk, n, m, chain, msg); err != nil {
					bk.Logger(ctx).Info(fmt.Sprintf("failed validating message: %s", err.Error()),
						types.AttributeKeyChain, msg.GetDestinationChain(),
						types.AttributeKeyMessageID, msg.ID,
					)
					return false, err
				}

				chainID := funcs.MustOk(destCk.GetChainID(ctx))
				keyID := funcs.MustOk(m.GetCurrentKeyID(ctx, chain.Name))

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

				events.Emit(ctx, &types.DestCallFailed{
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
	_, ok := m.GetCurrentKeyID(ctx, chain.Name)
	if !ok {
		return fmt.Errorf("current key not set for chain %v", chain.Name)
	}

	if !n.IsChainActivated(ctx, chain) {
		return fmt.Errorf("destination chain de-activated")
	}
	//Check gateway address is set for evm chain
	if types.IsEvmChain(chain) {
		if _, ok := ck.GetGatewayAddress(ctx); !ok {
			return fmt.Errorf("destination chain gateway for chain %v not deployed yet", chain.Name)
		}
	}
	if !common.IsHexAddress(msg.GetDestinationAddress()) {
		return fmt.Errorf("invalid contract address")
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

func handleMessageWithToken(ctx sdk.Context, ck types.ChainKeeper, n types.Nexus, chainID sdk.Int, keyID multisig.KeyID, msg nexus.GeneralMessage) error {
	token := ck.GetERC20TokenByAsset(ctx, msg.Asset.GetDenom())

	if err := n.RateLimitTransfer(ctx, msg.GetDestinationChain(), *msg.Asset, nexus.TransferDirectionTo); err != nil {
		return err
	}

	cmd := types.NewApproveContractCallWithMintGeneric(chainID, keyID, common.BytesToHash(msg.SourceTxID), msg.SourceTxIndex, msg, token.GetDetails().Symbol)
	funcs.MustNoErr(ck.EnqueueCommand(ctx, cmd))

	events.Emit(ctx, &types.DestCallWithMintApproved{
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

func handleMessage(ctx sdk.Context, ck types.ChainKeeper, chainID sdk.Int, keyID multisig.KeyID, msg nexus.GeneralMessage) {
	cmd := types.NewApproveBridgeCallCommandGeneric(chainID, keyID, common.HexToAddress(msg.GetDestinationAddress()), common.BytesToHash(msg.PayloadHash), common.BytesToHash(msg.SourceTxID), msg.GetSourceChain(), msg.GetSourceAddress(), msg.SourceTxIndex, msg.ID)
	funcs.MustNoErr(ck.EnqueueCommand(ctx, cmd))
	clog.Redf("[Chains] msg: %+v", msg)
	clog.Redf("[Chains] EnqueueCommand: %+v", cmd)

	destCallApproved := &types.DestCallApproved{
		Chain:            msg.GetSourceChain(),
		EventID:          types.EventID(msg.ID),
		CommandID:        cmd.ID,
		Sender:           msg.GetSourceAddress(),
		DestinationChain: msg.GetDestinationChain(),
		ContractAddress:  msg.GetDestinationAddress(),
		PayloadHash:      types.Hash(common.BytesToHash(msg.PayloadHash)),
	}

	clog.Redf("[Chains] destCallApproved: %+v", destCallApproved)

	events.Emit(ctx, destCallApproved)

	ck.Logger(ctx).Debug("completed handling general message",
		types.AttributeKeyChain, msg.GetDestinationChain(),
		types.AttributeKeyMessageID, msg.ID,
		types.AttributeKeyCommandsID, cmd.ID,
	)
}
