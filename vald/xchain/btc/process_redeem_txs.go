package btc

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/slices"
	xcommon "github.com/scalarorg/scalar-core/vald/xchain/common"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	covTypes "github.com/scalarorg/scalar-core/x/covenant/types"
)

func (client *BtcClient) ProcessRedeemTxsConfirmation(event *covTypes.ConfirmRedeemTxStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	txIDs := slices.Map(event.TxIDs, func(m chains.Hash) xcommon.Hash { return xcommon.Hash(m.Bytes()) })
	txReceipts, _ := client.GetTxReceiptsIfFinalized(txIDs, event.ConfirmationHeight)

	var eventIds []*chainsTypes.EventID
	var votes []sdk.Msg
	for i, txReceipt := range txReceipts {
		if txReceipt.Err() != nil {
			// This ensures all of the txs are valid
			return nil, fmt.Errorf("exist error in txReceipt: %s, txID: %s", txReceipt.Err().Error(), txIDs[i].String())
		} else {
			eventId := client.processRedeemTxReceipt(event, txReceipt.Ok().(BTCTxReceipt))
			if event == nil {
				continue
			}
			eventIds = append(eventIds, eventId)
		}
	}

	// events := []covTypes.Event{*event}
	// 		votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, covTypes.NewVoteEvents(event.Chain, events...)))
	// 		clog.Redf("broadcasting vote %v for poll %s", events, pollID.String())

	return votes, nil
}

func (client *BtcClient) processRedeemTxReceipt(event *covTypes.ConfirmRedeemTxStarted, receipt BTCTxReceipt) *chainsTypes.EventID {
	// TODO: 🛑 validate the btc protocol address from the event

	// TODO: Validate the session sequence, custodian_group_id in OP_RETURN

	// Note: TxID is the reversed-order hash of the txid aka RPC TxID, aka Mempool TxID
	txID, err := exported.HashFromHex(receipt.Raw.Txid)
	if err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "invalid tx id").Error())
		return nil
	}

	eventId := chainsTypes.NewEventID(txID, uint64(receipt.TransactionIndex))

	return &eventId
}

// func (client *BtcClient) GetUtxoLists(event *chainsTypes.UpdateUtxoListsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
// 	utxos, err := client.getUtxoList(event.TaprootAddress)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var buffer bytes.Buffer
// 	for _, utxo := range utxos {
// 		buffer.Write(utxo.TxID.Bytes())
// 		binary.Write(&buffer, binary.BigEndian, utxo.Vout)
// 		binary.Write(&buffer, binary.BigEndian, utxo.AmountInSats)
// 	}
// 	//Use hash of utxos as txID
// 	txID := sha3.Sum256(buffer.Bytes())

// 	voteEvent := chainsTypes.NewVoteEvents(event.Chain, chainsTypes.Event{
// 		Chain: event.Chain,
// 		TxID:  txID,
// 		Event: &chainsTypes.Event_UtxoListConfirmed{
// 			UtxoListConfirmed: &chainsTypes.UtxoListConfirmed{
// 				Utxos: utxos,
// 			},
// 		},
// 		Index: uint64(len(utxos)),
// 	})
// 	var votes []sdk.Msg
// 	votes = append(votes, voteTypes.NewVoteRequest(proxy, event.PollID, voteEvent))

// 	return votes, nil
// }
