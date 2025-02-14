package btc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	xcommon "github.com/scalarorg/scalar-core/vald/xchain/common"
	"github.com/scalarorg/scalar-core/x/chains/types"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *BtcClient) ProcessSourceTxsConfirmation(event *types.EventConfirmSourceTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	txIDs := slices.Map(event.PollMappings, func(m types.PollMapping) xcommon.Hash { return xcommon.Hash(m.TxID) })
	txReceipts, _ := client.GetTxReceiptsIfFinalized(txIDs, event.ConfirmationHeight)

	var votes []sdk.Msg
	// TODO: handle multiple tx receipts
	for i, txReceipt := range txReceipts {
		pollID := event.PollMappings[i].PollID
		if txReceipt.Err() != nil {
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain)))
			clog.Redf("broadcasting empty vote for poll %s: %s", pollID.String(), txReceipt.Err().Error())
		} else {
			events := client.processSrcTxReceipt(event, txReceipt.Ok().(BTCTxReceipt))
			if len(events) == 0 {
				continue
			}
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain, events...)))
			clog.Redf("broadcasting vote %v for poll %s", events, pollID.String())
		}
	}

	return votes, nil
}
func (client *BtcClient) processSrcTxReceipt(event *types.EventConfirmSourceTxsStarted, receipt BTCTxReceipt) []types.Event {
	// TODO: 🛑 validate the btc protocol address from the event
	clog.Redf("[BTC] txReceipt.Raw.Txid: %+v", receipt.Raw)
	clog.Redf("[BTC] txReceipt.TransactionIndex: %+v", receipt.TransactionIndex)
	clog.Redf("[BTC] txReceipt.Raw.Hash: %+v", receipt.Raw.Hash)
	var events []types.Event
	tokenSent, err := client.createEventTokenSent(event, &receipt)
	if err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "decode event EventConfirmSourceTxsStarted failed").Error())
		return nil
	}
	clog.Greenf("[BTC] btcEvent: %+v\n", tokenSent)

	if err := tokenSent.ValidateBasic(); err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "invalid event TokenSent").Error())
		return nil
	}

	// Note: TxID is the reversed-order hash of the txid aka RPC TxID, aka Mempool TxID
	txID, err := types.HashFromHex(receipt.Raw.Txid)
	if err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "invalid tx id").Error())
		return nil
	}
	events = append(events, types.Event{
		Chain: event.Chain,
		TxID:  txID,
		Event: &types.Event_TokenSent{
			TokenSent: tokenSent,
		},
		Index: uint64(receipt.TransactionIndex),
	})

	clog.Bluef("[BTC] SourceTxConfirmationEvent: %+v\n", events)
	return events
}

// 2025 Jan 06, Use EventTokenSent insteadof SourceTxConfirmation
// func (client *BtcClient) processSrcTxReceipt2(event *types.EventConfirmSourceTxsStarted, receipt BTCTxReceipt) []types.Event {

// 	var events []types.Event

// 	btcEvent, err := client.decodeSourceTxConfirmationEvent(&receipt)
// 	if err != nil {
// 		client.logger().Error(sdkerrors.Wrap(err, "decode event ContractCall failed").Error())
// 		return nil
// 	}

// 	clog.Greenf("[BTC] btcEvent: %+v\n", btcEvent)

// 	if err := btcEvent.ValidateBasic(); err != nil {
// 		client.logger().Error(sdkerrors.Wrap(err, "invalid event ContractCall").Error())
// 		return nil
// 	}

// 	// Note: TxID is the reversed-order hash of the txid aka RPC TxID, aka Mempool TxID
// 	txID, err := types.HashFromHex(receipt.Raw.TxID)
// 	if err != nil {
// 		client.logger().Error(sdkerrors.Wrap(err, "invalid tx id").Error())
// 		return nil
// 	}
// 	//Support transfer only, not contract call
// 	events = append(events, types.Event{
// 		Chain: event.Chain,
// 		TxID:  txID,
// 		Event: &types.Event_SourceTxConfirmationEvent{
// 			SourceTxConfirmationEvent: btcEvent,
// 		},
// 		Index: uint64(receipt.Raw.BlockIndex),
// 	})

// 	clog.Bluef("[BTC] SourceTxConfirmationEvent: %+v\n", events)

// 	return events
// }
