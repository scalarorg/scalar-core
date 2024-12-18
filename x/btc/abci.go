package btc

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/btc/types"
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
	clog.Yellow("BTC ABCI ENDBLOCKER")
	handleConfirmedEvents(ctx, bk, n, m)
	handleMessages(ctx, bk, n, m)

	return nil, nil
}

func handleConfirmedEvents(ctx sdk.Context, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper) {
	for _, chain := range slices.Filter(n.GetChains(ctx), types.IsBTCChain) {
		handleConfirmedEventsForChain(ctx, chain, bk, n, m)
	}
}

func handleConfirmedEventsForChain(ctx sdk.Context, chain nexus.Chain, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper) {
	clog.Red("handleConfirmedEventsForChain", "chain", chain.Name.String())
	ck := funcs.Must(bk.ForChain(ctx, chain.Name))
	queue := ck.GetConfirmedEventQueue(ctx)
	endBlockerLimit := ck.GetParams(ctx).EndBlockerLimit

	var events []types.Event
	var event types.Event
	for int64(len(events)) < endBlockerLimit && queue.Dequeue(&event) {
		events = append(events, event)
	}

	for _, event := range events {
		clog.Redf("[BTC] handleConfirmedEvent: %+v", event)
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
	clog.Red("handleConfirmedEvent", "event", event.GetID())
	if err := validateEvent(ctx, event, bk, n); err != nil {
		return err
	}

	switch event.GetEvent().(type) {
	case *types.Event_StakingTx:
		return handleStakingTx(ctx, event, bk, n, m)

	// TODO: add other event types here

	default:
		panic(fmt.Errorf("unsupported event type %T", event))
	}
}

func validateEvent(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus) error {
	clog.Red("validateEvent", "event", event.GetID())
	var destinationChainName nexus.ChainName
	var contractAddress string
	switch event := event.GetEvent().(type) {
	case *types.Event_StakingTx:
		destinationChainName = event.StakingTx.DestinationChain
		contractAddress = event.StakingTx.Metadata.DestinationContractAddress.String()
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

	// TODO: handle for the same btc chain

	return nil
}

func handleStakingTx(ctx sdk.Context, event types.Event, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper) error {
	clog.Red("handleStakingTx", "event", event.GetID())
	e := event.GetStakingTx()
	if e == nil {
		panic(fmt.Errorf("event is nil"))
	}

	destinationChain := funcs.MustOk(n.GetChain(ctx, e.DestinationChain))
	clog.Red("handleStakingTx", "destinationChain", destinationChain.Name.String())
	switch destinationChain.Module {
	case types.ModuleName:
		// handleStakingTxToBTC(ctx, bk, multisig, n, destinationChain.Name, event)
		// TODO: implement
		clog.Red("handleStakingTx for the same btc chain", "destinationChain", destinationChain.Name.String())
		return nil
	default:
		// set as general message in nexus, so the dest module can handle the message
		return setMessageToNexus(ctx, n, event, nil)
	}
}

func setMessageToNexus(ctx sdk.Context, n types.Nexus, event types.Event, asset *sdk.Coin) error {

	clog.Red("setMessageToNexus", "event", event)

	sourceChain := funcs.MustOk(n.GetChain(ctx, event.Chain))

	var message nexus.GeneralMessage
	switch e := event.GetEvent().(type) {
	case *types.Event_StakingTx:
		sender := nexus.CrossChainAddress{
			Chain:   sourceChain,
			Address: e.StakingTx.Sender,
		}

		recipient := nexus.CrossChainAddress{
			Chain:   funcs.MustOk(n.GetChain(ctx, e.StakingTx.DestinationChain)),
			Address: e.StakingTx.Metadata.DestinationContractAddress.String(),
		}

		message = nexus.NewGeneralMessage(
			string(event.GetID()),
			sender,
			recipient,
			e.StakingTx.PayloadHash.Bytes(),
			event.TxID.Bytes(),
			event.Index,
			nil,
		)

	// TODO: add other event types here

	default:
		return fmt.Errorf("unsupported event type %T", event)
	}

	clog.Red("setMessageToNexus", "message", message)

	if message.Recipient.Chain.Name.Equals(scalarnet.Scalarnet.Name) {
		return fmt.Errorf("%s is not a supported recipient", scalarnet.Scalarnet.Name)
	}

	return n.SetNewMessage(ctx, message)
}

func handleMessages(ctx sdk.Context, bk types.BaseKeeper, n types.Nexus, m types.MultisigKeeper) {
	for _, chain := range slices.Filter(n.GetChains(ctx), types.IsBTCChain) {
		destCk := funcs.Must(bk.ForChain(ctx, chain.Name))
		endBlockerLimit := destCk.GetParams(ctx).EndBlockerLimit
		msgs := n.GetProcessingMessages(ctx, chain.Name, endBlockerLimit)

		bk.Logger(ctx).Debug(fmt.Sprintf("handling %d general messages", len(msgs)), types.AttributeKeyChain, chain.Name)

		for _, msg := range msgs {
			clog.Red("handleMessages", "msg", msg)
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
					// if err := handleMessageWithToken(ctx, destCk, n, chainID, keyID, msg); err != nil {
					// 	return false, err
					// }
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

				clog.Red("failed handling general message",
					types.AttributeKeyChain, chain.Name.String(),
					types.AttributeKeyMessageID, msg.ID,
				)
				events.Emit(ctx, &types.BridgeCallFailed{
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
		return fmt.Errorf("current key not set")
	}

	if !n.IsChainActivated(ctx, chain) {
		return fmt.Errorf("destination chain de-activated")
	}

	// TODO: How to check if destination chain is EVM-compatible?
	clog.Yellow("TODO: How to check if destination chain is EVM-compatible?")
	// if _, ok := ck.GetGatewayAddress(ctx); !ok {
	// 	return fmt.Errorf("destination chain gateway not deployed yet")
	// }

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

func handleMessage(ctx sdk.Context, ck types.ChainKeeper, chainID sdk.Int, keyID multisig.KeyID, msg nexus.GeneralMessage) {
	cmd := types.NewApproveBridgeCallCommandGeneric(chainID, keyID, common.HexToAddress(msg.GetDestinationAddress()), common.BytesToHash(msg.PayloadHash), common.BytesToHash(msg.SourceTxID), msg.GetSourceChain(), msg.GetSourceAddress(), msg.SourceTxIndex, msg.ID)
	clog.Redf("[BTC] EnqueueCommand: %+v", cmd)
	funcs.MustNoErr(ck.EnqueueCommand(ctx, cmd))

	events.Emit(ctx, &types.BridgeCallApproved{
		Chain:            msg.GetSourceChain(),
		EventID:          types.EventID(msg.ID),
		CommandID:        cmd.ID,
		Sender:           msg.GetSourceAddress(),
		DestinationChain: msg.GetDestinationChain(),
		ContractAddress:  msg.GetDestinationAddress(),
		PayloadHash:      types.Hash(common.BytesToHash(msg.PayloadHash)),
	})

	ck.Logger(ctx).Debug("completed handling general message",
		types.AttributeKeyChain, msg.GetDestinationChain(),
		types.AttributeKeyMessageID, msg.ID,
		types.AttributeKeyCommandsID, cmd.ID,
	)
}
