package btc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/vald/xchain"
	"github.com/scalarorg/scalar-core/x/chains/types"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *BtcClient) ProcessSourceTxsConfirmation(event *types.EventConfirmSourceTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	txIDs := slices.Map(event.PollMappings, func(m types.PollMapping) xchain.Hash { return m.TxID })
	txReceipts, _ := client.GetTxReceiptsIfFinalized(txIDs, event.ConfirmationHeight)

	clog.Redf("[BTC] txReceipts: %+v", txReceipts)

	var votes []sdk.Msg
	// TODO: handle multiple tx receipts
	for i, txReceipt := range txReceipts {
		pollID := event.PollMappings[i].PollID
		if txReceipt.Err() != nil {
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain)))
			clog.Redf("broadcasting empty vote for poll %s: %s", pollID.String(), txReceipt.Err().Error())
		} else {
			events := client.processSrcTxReceipt(event, txReceipt.Ok().(BTCTxReceipt))
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain, events...)))
			clog.Redf("broadcasting vote %v for poll %s", events, pollID.String())
		}
	}

	return votes, nil
}

func (client *BtcClient) processSrcTxReceipt(event *types.EventConfirmSourceTxsStarted, receipt BTCTxReceipt) []types.Event {

	var events []types.Event

	btcEvent, err := client.decodeSourceTxConfirmationEvent(&receipt)
	if err != nil {
		client.logger().Debug(sdkerrors.Wrap(err, "decode event ContractCall failed").Error())
	}

	if err := btcEvent.ValidateBasic(); err != nil {
		client.logger().Debug(sdkerrors.Wrap(err, "invalid event ContractCall").Error())
	}

	txID, err := types.HashFromHex(receipt.Raw.Txid)
	if err != nil {
		client.logger().Debug(sdkerrors.Wrap(err, "invalid tx id").Error())
	}

	events = append(events, types.Event{
		Chain: event.Chain,
		TxID:  txID,
		Event: &types.Event_SourceTxConfirmationEvent{
			SourceTxConfirmationEvent: btcEvent,
		},
		Index: 0, // TODO: fix this hardcoded index, this is used to identify the staking tx in the event
	})

	return events
}
