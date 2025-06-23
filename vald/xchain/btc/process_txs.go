package btc

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	xcommon "github.com/scalarorg/scalar-core/vald/xchain/common"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/chains/types"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *BtcClient) ProcessSourceTxsConfirmation(event *types.EventConfirmSourceTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	txIDs := slices.Map(event.PollMappings, func(m types.PollMapping) xcommon.Hash { return xcommon.Hash(m.Hash) })
	txReceipts, _ := client.GetTxReceiptsIfFinalized(txIDs, event.ConfirmationHeight)

	var votes []sdk.Msg
	// TODO: handle multiple tx receipts
	for i, txReceipt := range txReceipts {
		pollID := event.PollMappings[i].PollID
		if txReceipt.Err() != nil {
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain)))
			clog.Bluef("broadcasting empty vote for poll %s: %s", pollID.String(), txReceipt.Err().Error())
		} else {
			events := client.processSrcTxReceipt(event, txReceipt.Ok().(BTCTxReceipt))
			if len(events) == 0 {
				continue
			}
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain, events...)))
			clog.Bluef("broadcasting vote %v for poll %s", events, pollID.String())
		}
	}

	return votes, nil
}
func (client *BtcClient) processSrcTxReceipt(event *types.EventConfirmSourceTxsStarted, receipt BTCTxReceipt) []types.Event {
	// TODO: 🛑 validate the btc protocol address from the event
	clog.Bluef("[BTC] txReceipt.Raw.Txid: %+v, TxIndex: %+v", receipt.Raw.Txid, receipt.TransactionIndex)
	var events []types.Event
	tokenSent, err := client.CreateEventTokenSent(event, &receipt)
	if err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "decode event EventConfirmSourceTxsStarted failed").Error())
		client.logger().Error(fmt.Sprintf("receipt txid: %s, raw hex: %s", receipt.Raw.Txid, receipt.Raw.Hex))
		return nil
	}
	clog.Greenf("[BTC] btcEvent: %+v\n", tokenSent)

	if err = tokenSent.ValidateBasic(); err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "invalid event TokenSent").Error())
		return nil
	}

	// Note: TxID is the reversed-order hash of the txid aka RPC TxID, aka Mempool TxID
	txID, err := exported.HashFromHex(receipt.Raw.Txid)
	if err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "invalid tx id").Error())
		return nil
	}
	events = append(events, types.Event{
		Chain: event.Chain,
		Hash:  txID,
		Event: &types.Event_TokenSent{
			TokenSent: tokenSent,
		},
		Index: uint64(receipt.TransactionIndex),
	})

	clog.Bluef("[BTC] SourceTxConfirmationEvent: %+v\n", events)
	return events
}
