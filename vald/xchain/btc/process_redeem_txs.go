package btc

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/go-common/btc"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	xcommon "github.com/scalarorg/scalar-core/vald/xchain/common"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	covTypes "github.com/scalarorg/scalar-core/x/covenant/types"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *BtcClient) ProcessRedeemTxsConfirmation(event *covTypes.ConfirmRedeemTxStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	txIDs := slices.Map(event.TxIDs, func(m chains.Hash) xcommon.Hash { return xcommon.Hash(m.Bytes()) })
	txReceipts, _ := client.GetTxReceiptsIfFinalized(txIDs, event.ConfirmationHeight)
	if len(txReceipts) != len(txIDs) {
		// This ensures all of the txs are valid
		return nil, fmt.Errorf("invalid txReceipts length: %d, txIDs length: %d", len(txReceipts), len(txIDs))
	}

	var eventIds []chainsTypes.EventID
	for i, txReceipt := range txReceipts {
		if txReceipt.Err() != nil {
			// This ensures all of the txs are valid
			return nil, fmt.Errorf("exist error in txReceipt: %s, txID: %s", txReceipt.Err().Error(), txIDs[i].String())
		} else {
			eventId := client.processRedeemTxReceipt(txReceipt.Ok().(BTCTxReceipt))
			if event == nil {
				continue
			}
			eventIds = append(eventIds, *eventId)
		}
	}

	taprootAddress, err := btc.ScriptPubKeyToAddress(event.ScriptPubkey, event.NetworkParams)
	if err != nil {
		return nil, err
	}

	utxos, _, err := client.getUtxoList(taprootAddress.String())
	if err != nil {
		return nil, err
	}
	txIDStrs := slices.Map(event.TxIDs, func(tx chains.Hash) string { return tx.Hex() })
	maxBlockHeight := maxInt64(slices.Map(txReceipts, func(m BTCTxResult) int64 { return *m.Ok().(BTCTxReceipt).BlockHeight }))
	log.Debug().Any("RedeemTx Ids", txIDStrs).Int64("MaxBlockHeight", maxBlockHeight).Msg("Create new UtxoSnapshot")
	utxoSnapshot := covTypes.UTXOSnapshot{
		CustodianGroupUID: event.CustodianGroupUID,
		BlockHeight:       uint64(maxBlockHeight),
		Utxos:             utxos,
	}
	hash := utxoSnapshot.GetHash()
	voteEvent := covTypes.NewVoteEvents(event.Chain, covTypes.Event{
		Chain: event.Chain,
		Hash:  &hash,
		Event: &covTypes.Event_RedeemTxsConfirmed{
			RedeemTxsConfirmed: &covTypes.RedeemTxsConfirmed{
				EventIDs:     eventIds,
				UtxoSnapshot: &utxoSnapshot,
			},
		},
		Index: 0,
	})

	votes := []sdk.Msg{
		voteTypes.NewVoteRequest(proxy, event.PollID, voteEvent),
	}
	clog.Redf("[ProcessRedeemTxsConfirmation] broadcasting vote for poll %s", event.PollID.String())

	return votes, nil
}

func (client *BtcClient) processRedeemTxReceipt(receipt BTCTxReceipt) *chainsTypes.EventID {
	// TODO: 🛑 validate the btc protocol address from the event

	// TODO: Validate the session sequence, custodian_group_id in OP_RETURN

	// Note: TxID is the reversed-order hash of the txid aka RPC TxID, aka Mempool TxID
	txID, err := chains.HashFromHex(receipt.Raw.Txid)
	if err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "invalid tx id").Error())
		return nil
	}

	eventId := chainsTypes.NewEventID(txID, uint64(receipt.TransactionIndex))

	return &eventId
}
