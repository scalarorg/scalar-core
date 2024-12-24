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

func (client *EthereumClient) ProcessSourceTxsConfirmation(event *types.EventConfirmSourceTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {

	txIDs := slices.Map(event.PollMappings, func(m types.PollMapping) xchain.Hash { return m.TxID })
	clog.Redf("[ETH] txIDs: %+v", txIDs)
	txReceipts, _ := client.GetTxReceiptsIfFinalized(txIDs, event.ConfirmationHeight)

	clog.Redf("[ETH] txReceipts: %+v", txReceipts)

	var votes []sdk.Msg
	for i, txReceipt := range txReceipts {
		pollID := event.PollMappings[i].PollID
		if txReceipt.Err() != nil {
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain)))
			clog.Redf("broadcasting empty vote for poll %s: %s", pollID.String(), txReceipt.Err().Error())
		} else {
			events := client.processTxReceipt(event.Chain, txReceipt.Ok().(ETHTxReceipt))
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain, events...)))
			clog.Redf("broadcasting vote %v for poll %s", events, pollID.String())
		}
	}

	return votes, nil
}

func (c *EthereumClient) processTxReceipt(chain nexus.ChainName, receipt ETHTxReceipt) []types.Event {
	var events []types.Event

	clog.Red("processTxReceipt", "receipt", receipt)
	clog.Red("processTxReceipt", "logs", receipt.Logs)

	for _, txlog := range receipt.Logs {
		if len(txlog.Topics) == 0 {
			continue
		}

		switch txlog.Topics[0] {
		case ContractCallSig:
			contractCallEvent, err := c.decodeSourceTxConfirmationEvent(txlog)
			if err != nil {
				c.logger().Debug(sdkerrors.Wrap(err, "decode event ContractCall failed").Error())
				continue
			}
			clog.Red("processTxReceipt", "contractCallEvent", contractCallEvent)

			if err := contractCallEvent.ValidateBasic(); err != nil {
				c.logger().Error(sdkerrors.Wrap(err, "invalid event ContractCall").Error())
				continue
			}

			events = append(events, types.Event{
				Chain: chain,
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
