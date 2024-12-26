package evm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/vald/xchain"
	"github.com/scalarorg/scalar-core/x/chains/types"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *EthereumClient) ProcessSourceTxsConfirmation(event *types.EventConfirmSourceTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {

	clog.Red("ProcessSourceTxsConfirmation", "event", event)

	txIDs := slices.Map(event.PollMappings, func(m types.PollMapping) xchain.Hash { return m.TxID })
	txReceipts, _ := client.GetTxReceiptsIfFinalized(txIDs, event.ConfirmationHeight)

	var votes []sdk.Msg
	for i, txReceipt := range txReceipts {
		pollID := event.PollMappings[i].PollID
		if txReceipt.Err() != nil {
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain)))
			clog.Redf("broadcasting empty vote for poll %s: %s", pollID.String(), txReceipt.Err().Error())
		} else {
			events := client.processTxReceipt(event, txReceipt.Ok().(ETHTxReceipt))
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain, events...)))
			clog.Redf("broadcasting vote %v for poll %s", events, pollID.String())
		}
	}

	return votes, nil
}

func (c *EthereumClient) processTxReceipt(event *types.EventConfirmSourceTxsStarted, receipt ETHTxReceipt) []types.Event {
	var events []types.Event

	for _, txlog := range receipt.Logs {
		if len(txlog.Topics) == 0 {
			continue
		}

		clog.Red("processTxReceipt", "txlog", txlog)

		switch txlog.Topics[0] {
		case ContractCallSig:
			contractCallEvent, err := c.decodeSourceTxConfirmationEvent(txlog)
			if err != nil {
				c.logger().Error(sdkerrors.Wrap(err, "decode event ContractCall failed").Error())
				continue
			}
			clog.Red("processTxReceipt", "contractCallEvent", contractCallEvent)

			if err := contractCallEvent.ValidateBasic(); err != nil {
				c.logger().Error(sdkerrors.Wrap(err, "invalid event ContractCall").Error())
				continue
			}

			events = append(events, types.Event{
				Chain: event.Chain,
				TxID:  xchain.Hash(receipt.TxHash),
				Event: &types.Event_SourceTxConfirmationEvent{
					SourceTxConfirmationEvent: contractCallEvent,
				},
				Index: uint64(txlog.Index),
			})
		default:
			c.logger().Errorf("unknown event type: %s", txlog.Topics[0])
		}
	}

	return events
}
