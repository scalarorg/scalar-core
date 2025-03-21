package btc

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	xcommon "github.com/scalarorg/scalar-core/vald/xchain/common"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	covTypes "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *BtcClient) ProcessRedeemTxsConfirmation(event *covTypes.ConfirmRedeemTxStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	txIDs := slices.Map(event.PollMappings, func(m chainsTypes.PollMapping) xcommon.Hash { return xcommon.Hash(m.TxID) })
	txReceipts, _ := client.GetTxReceiptsIfFinalized(txIDs, event.ConfirmationHeight)

	var votes []sdk.Msg
	for i, txReceipt := range txReceipts {
		pollID := event.PollMappings[i].PollID
		if txReceipt.Err() != nil {
			// This ensures all of the txs are valid
			return nil, fmt.Errorf("exist error in txReceipt: %s, txID: %s", txReceipt.Err().Error(), txIDs[i].String())
		} else {
			event := client.processRedeemTxReceipt(event, txReceipt.Ok().(BTCTxReceipt))
			if event == nil {
				continue
			}

			events := []chainsTypes.Event{*event}
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, chainsTypes.NewVoteEvents(event.Chain, events...)))
			clog.Redf("broadcasting vote %v for poll %s", events, pollID.String())
		}
	}

	return votes, nil
}

func (client *BtcClient) processRedeemTxReceipt(event *covTypes.ConfirmRedeemTxStarted, receipt BTCTxReceipt) *chainsTypes.Event {
	// TODO: 🛑 validate the btc protocol address from the event
	redeemTxConfirmed, err := client.createEventRedeemTxConfirmed(event, &receipt)
	if err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "decode event EventConfirmRedeemTxStarted failed").Error())
		return nil
	}
	clog.Greenf("[BTC] btcEvent: %+v\n", redeemTxConfirmed)

	if err := redeemTxConfirmed.ValidateBasic(); err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "invalid event RedeemTxConfirmed").Error())
		return nil
	}

	// Note: TxID is the reversed-order hash of the txid aka RPC TxID, aka Mempool TxID
	txID, err := exported.HashFromHex(receipt.Raw.Txid)
	if err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "invalid tx id").Error())
		return nil
	}

	return &chainsTypes.Event{
		Chain: event.Chain,
		TxID:  txID,
		// Event: &covTypes.Event_RedeemTxConfirmed{
		// 	RedeemTxConfirmed: redeemTxConfirmed,
		// },
		Index: uint64(receipt.TransactionIndex),
	}
}

func (client *BtcClient) createEventRedeemTxConfirmed(event *covTypes.ConfirmRedeemTxStarted, tx *BTCTxReceipt) (*covTypes.RedeemTxConfirmed, error) {
	// TODO: Validate the session sequence in OP_RETURN

	// Note: TxID is the reversed-order hash of the txid aka RPC TxID, aka Mempool TxID
	txID, err := exported.HashFromHex(tx.Raw.Txid)
	if err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "invalid tx id").Error())
		return nil, fmt.Errorf("invalid tx id %s", tx.Raw.Txid)
	}

	eventId := chainsTypes.NewEventID(txID, uint64(tx.TransactionIndex))

	return &covTypes.RedeemTxConfirmed{
		EventID: eventId,
		Chain:   nexus.ChainName(event.Chain),
	}, nil
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
