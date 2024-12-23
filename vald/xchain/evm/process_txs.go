package evm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/vald/xchain"
	"github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

type newEventWrapper[T types.ConfirmationEvent] func(*types.TxConfirmationEvent) T

func (client *EthereumClient) ProcessDestinationTxsConfirmation(event *types.EventConfirmDestTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	clog.Blue("[EVM] ProcessDestinationTxsConfirmation", "txIDs", event.PollMappings)
	return processTxsConfirmation(
		client,
		event.Chain,
		event.PollMappings,
		event.ConfirmationHeight,
		proxy,
		func(e *types.TxConfirmationEvent) types.ConfirmationEvent {
			return &types.Event_DestTxConfirmationEvent{
				DestTxConfirmationEvent: e,
			}
		},
	)
}

func (client *EthereumClient) ProcessSourceTxsConfirmation(event *types.EventConfirmSourceTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	return processTxsConfirmation(
		client,
		event.Chain,
		event.PollMappings,
		event.ConfirmationHeight,
		proxy,
		func(e *types.TxConfirmationEvent) types.ConfirmationEvent {
			return &types.Event_SourceTxConfirmationEvent{
				SourceTxConfirmationEvent: e,
			}
		},
	)
}

func processTxsConfirmation[T types.ConfirmationEvent](
	client *EthereumClient,
	chain nexus.ChainName,
	pollMappings []types.PollMapping,
	confirmationHeight uint64,
	proxy sdk.AccAddress,
	newEventWrapper newEventWrapper[T],
) ([]sdk.Msg, error) {
	txIDs := slices.Map(pollMappings, func(m types.PollMapping) xchain.Hash { return m.TxID })
	clog.Redf("[ETH] txIDs: %+v", txIDs)
	txReceipts, _ := client.GetTxReceiptsIfFinalized(txIDs, confirmationHeight)

	clog.Redf("[ETH] txReceipts: %+v", txReceipts)

	var votes []sdk.Msg
	for i, txReceipt := range txReceipts {
		pollID := pollMappings[i].PollID
		if txReceipt.Err() != nil {
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(chain)))
			clog.Redf("broadcasting empty vote for poll %s: %s", pollID.String(), txReceipt.Err().Error())
		} else {
			events := processTxReceipt(client, chain, txReceipt.Ok().(ETHTxReceipt), newEventWrapper)
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(chain, events...)))
			clog.Redf("broadcasting vote %v for poll %s", events, pollID.String())
		}
	}

	return votes, nil
}

func processTxReceipt[T types.ConfirmationEvent](
	client *EthereumClient,
	chain nexus.ChainName,
	receipt ETHTxReceipt,
	newEventWrapper newEventWrapper[T],
) []types.Event {
	var events []types.Event

	for _, txlog := range receipt.Logs {
		if len(txlog.Topics) == 0 {
			continue
		}

		switch txlog.Topics[0] {
		case ContractCallSig:
			contractCallEvent, err := client.decodeEventContractCall(txlog)
			if err != nil {
				client.logger().Debug(sdkerrors.Wrap(err, "decode event ContractCall failed").Error())
				continue
			}

			if err := contractCallEvent.ValidateBasic(); err != nil {
				client.logger().Debug(sdkerrors.Wrap(err, "invalid event ContractCall").Error())
				continue
			}

			events = append(events, types.Event{
				Chain: chain,
				TxID:  xchain.Hash(receipt.TxHash),
				Event: newEventWrapper(contractCallEvent),
				Index: uint64(txlog.Index),
			})
		default:
			client.logger().Debugf("unknown event type: %s", txlog.Topics[0])
		}
	}

	return events
}
