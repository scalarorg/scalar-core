package evm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	xcommon "github.com/scalarorg/scalar-core/vald/xchain/common"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/chains/types"
	covTypes "github.com/scalarorg/scalar-core/x/covenant/types"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *EthereumClient) ProcessSwitchedPhaseConfirmation(event *covTypes.ConfirmSwitchedPhaseStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	txIDs := slices.Map([]chains.Hash{event.TxID}, func(txid chains.Hash) xcommon.Hash { return xcommon.Hash(txid.Bytes()) })
	txReceipts, _ := client.GetTxReceiptsIfFinalized(txIDs, event.ConfirmationHeight)

	var votes []sdk.Msg
	for _, txReceipt := range txReceipts {
		pollID := event.PollID
		if txReceipt.Err() != nil {
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain)))
			clog.Redf("[ProcessSwitchedPhaseConfirmation] broadcasting empty vote for poll %s: %s", pollID.String(), txReceipt.Err().Error())
		} else {
			events := client.processConfirmSwitched(event, txReceipt.Ok().(ETHTxReceipt))
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, covTypes.NewVoteEvents(event.Chain, events...)))
			clog.Bluef("[ProcessSwitchedPhaseConfirmation] broadcasting vote for poll %s, %++v", pollID.String(), events)
		}
	}
	return votes, nil
}

func (c *EthereumClient) processConfirmSwitched(event *covTypes.ConfirmSwitchedPhaseStarted, receipt ETHTxReceipt) []covTypes.Event {
	events := []covTypes.Event{}

	for _, txlog := range receipt.Logs {

		// TODO: 🛑 validate the pool address from the event

		// if !bytes.Equal(gatewayAddress.Bytes(), txlog.Address.Bytes()) {
		// 	continue
		// }

		if len(txlog.Topics) == 0 {
			continue
		}

		clog.Red("processConfirmSwitched", "txlog", txlog)

		switch txlog.Topics[0] {
		case SwitchPhaseSig:
			switchPhaseEvent, err := DecodeEventSwitchPhase(txlog)
			if err != nil {
				c.logger().Infof(sdkerrors.Wrap(err, "decode event SwitchPhase failed").Error())
				continue
			}

			if err := switchPhaseEvent.ValidateBasic(); err != nil {
				c.logger().Debug(sdkerrors.Wrap(err, "invalid event SwitchPhase").Error())
				continue
			}

			hash := chains.Hash(txlog.TxHash)

			events = append(events, covTypes.Event{
				Chain: event.Chain,
				Hash:  &hash,
				Index: uint64(txlog.Index),
				Event: &covTypes.Event_SwitchedPhaseConfirmed{
					SwitchedPhaseConfirmed: switchPhaseEvent,
				},
			})
		default:
			c.logger().Debugf("unknown event type: %s", txlog.Topics[0])
		}
	}

	return events
}
