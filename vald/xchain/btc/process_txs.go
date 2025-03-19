package btc

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/btcsuite/btcd/txscript"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	vault "github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	go_utils "github.com/scalarorg/bitcoin-vault/go-utils/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	xcommon "github.com/scalarorg/scalar-core/vald/xchain/common"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/chains/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
	"golang.org/x/crypto/sha3"
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
	txID, err := exported.HashFromHex(receipt.Raw.Txid)
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

func (client *BtcClient) ProcessRedeemTxConfirmation(event *types.ConfirmRedeemTxStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
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
			events := client.processRedeemTxReceipt(event, txReceipt.Ok().(BTCTxReceipt))
			if len(events) == 0 {
				continue
			}
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain, events...)))
			clog.Redf("broadcasting vote %v for poll %s", events, pollID.String())
		}
	}

	return votes, nil
}

func (client *BtcClient) processRedeemTxReceipt(event *types.ConfirmRedeemTxStarted, receipt BTCTxReceipt) []types.Event {
	// TODO: 🛑 validate the btc protocol address from the event
	var events []types.Event
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
	events = append(events, types.Event{
		Chain: event.Chain,
		TxID:  txID,
		Event: &types.Event_RedeemTxConfirmed{
			RedeemTxConfirmed: redeemTxConfirmed,
		},
		Index: uint64(receipt.TransactionIndex),
	})

	clog.Bluef("[BTC] RedeemTxConfirmationEvent: %+v\n", events)
	return events
}

func (client *BtcClient) createEventRedeemTxConfirmed(event *types.ConfirmRedeemTxStarted, tx *BTCTxReceipt) (*types.RedeemTxConfirmed, error) {
	if len(tx.MsgTx.TxOut) < MinNumberOfOutputs {
		return nil, ErrInvalidTxOutCount
	}

	// Note: TxID is the reversed-order hash of the txid aka RPC TxID, aka Mempool TxID
	txID, err := exported.HashFromHex(tx.Raw.Txid)
	if err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "invalid tx id").Error())
		return nil, fmt.Errorf("invalid tx id %s", tx.Raw.Txid)
	}

	eventId := chainsTypes.NewEventID(txID, uint64(tx.TransactionIndex))

	embeddedDataTxOut := tx.MsgTx.TxOut[EmbeddedDataOutputIndex]
	if embeddedDataTxOut == nil || embeddedDataTxOut.PkScript == nil || embeddedDataTxOut.PkScript[0] != txscript.OP_RETURN {
		return nil, ErrInvalidOpReturn
	}

	output, err := vault.ParseVaultEmbeddedData(embeddedDataTxOut.PkScript)
	if err != nil || output == nil {
		return nil, ErrInvalidOpReturnData
	}

	if output.TransactionType != go_utils.TransactionTypeStaking {
		return nil, ErrInvalidTransactionType
	}

	var redeemAmount int64 = tx.MsgTx.TxOut[StakingOutputIndex].Value

	var destinationRecipientAddress chainsTypes.Address
	err = destinationRecipientAddress.Unmarshal(output.DestinationRecipientAddress)
	if err != nil {
		return nil, err
	}

	// queryClient := grpc_client.QueryManager().GetProtocolClient()

	// response, err := queryClient.Protocol(context.Background(), &protocolTypes.ProtocolRequest{
	// 	OriginChain: event.Chain,
	// 	MinorChain:  nexus.ChainName(destinationChain.ToBytes().String()),
	// 	Address:     hex.EncodeToString(output.DestinationTokenAddress),
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// clog.Greenf("EventTokenSent/Asset Response: %+v", response.Protocol)

	return &types.RedeemTxConfirmed{
		EventID:            eventId,
		Chain:              nexus.ChainName(event.Chain),
		DestinationAddress: chainsTypes.Address(destinationRecipientAddress).Hex(),
		RedeemAmount:       redeemAmount,
	}, nil
}

func (client *BtcClient) GetUtxoLists(event *types.UpdateUtxoListsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	utxos, err := client.getUtxoList(event.TaprootAddress)
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	for _, utxo := range utxos {
		buffer.Write(utxo.TxID.Bytes())
		binary.Write(&buffer, binary.BigEndian, utxo.Vout)
		binary.Write(&buffer, binary.BigEndian, utxo.AmountInSats)
	}
	//Use hash of utxos as txID
	txID := sha3.Sum256(buffer.Bytes())

	voteEvent := types.NewVoteEvents(event.Chain, types.Event{
		Chain: event.Chain,
		TxID:  txID,
		Event: &types.Event_UtxoListConfirmed{
			UtxoListConfirmed: &types.UtxoListConfirmed{
				Utxos: utxos,
			},
		},
		Index: uint64(len(utxos)),
	})
	var votes []sdk.Msg
	votes = append(votes, voteTypes.NewVoteRequest(proxy, event.PollID, voteEvent))

	return votes, nil
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
// 	txID, err := exported.HashFromHex(receipt.Raw.TxID)
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
