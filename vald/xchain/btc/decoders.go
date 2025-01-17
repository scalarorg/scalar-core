package btc

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/txscript"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	vault "github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/scalar-core/utils/clog"
	grpc_client "github.com/scalarorg/scalar-core/vald/grpc-client"
	"github.com/scalarorg/scalar-core/x/chains/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	protocolTypes "github.com/scalarorg/scalar-core/x/protocol/types"
)

var (
	StakingOutputIndex      = 0
	EmbeddedDataOutputIndex = 1
)

var (
	ErrInvalidTxOutCount       = errors.New("btcLocking tx must have at least 3 outputs")
	ErrInvalidOpReturn         = errors.New("transaction does not have expected payload op return output")
	ErrInvalidOpReturnData     = errors.New("cannot parse payload op return data")
	ErrInvalidTransactionType  = errors.New("invalid transaction type, expected staking")
	ErrInvalidTxId             = errors.New("failed to decode tx id")
	ErrInvalidPayloadHash      = errors.New("failed to get payload hash")
	ErrInvalidDestinationChain = errors.New("failed to parse destination chain")
)

const (
	MinNumberOfOutputs = 2
	SYMBOL_SCALAR_BTC  = "sBtc" //Todo: get from keeper
)

func (client *BtcClient) createEventTokenSent(event *types.EventConfirmSourceTxsStarted, tx *BTCTxReceipt) (*chainsTypes.EventTokenSent, error) {
	if len(tx.MsgTx.TxOut) < MinNumberOfOutputs {
		return nil, ErrInvalidTxOutCount
	}

	// Note: TxID is the reversed-order hash of the txid aka RPC TxID, aka Mempool TxID
	txID, err := types.HashFromHex(tx.Raw.Txid)
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

	if output.TransactionType != vault.TransactionTypeStaking {
		return nil, ErrInvalidTransactionType
	}

	var stakingAmount int64 = tx.MsgTx.TxOut[StakingOutputIndex].Value

	destinationChain := chain.NewChainInfoFromBytes(output.DestinationChain)
	if destinationChain == nil {
		return nil, ErrInvalidDestinationChain
	}

	var destinationRecipientAddress chainsTypes.Address
	err = destinationRecipientAddress.Unmarshal(output.DestinationRecipientAddress)
	if err != nil {
		return nil, err
	}

	queryClient := grpc_client.QueryManager().GetProtocolClient()

	response, err := queryClient.ProtocolAsset(context.Background(), &protocolTypes.ProtocolAssetRequest{
		SourceChain:      event.Chain,
		TokenAddress:     hex.EncodeToString(output.DestinationTokenAddress),
		DestinationChain: nexus.ChainName(destinationChain.ToBytes().String()),
	})
	if err != nil {
		return nil, err
	}

	clog.Greenf("EventTokenSent/Asset Response: %+v", response.Asset)

	return &chainsTypes.EventTokenSent{
		EventID:            eventId,
		Sender:             tx.PrevTxOuts[0].ScriptPubKey.Address,
		Chain:              nexus.ChainName(event.Chain),
		TransferID:         nexus.TransferID(1),
		DestinationChain:   nexus.ChainName(destinationChain.ToBytes().String()),
		DestinationAddress: chainsTypes.Address(destinationRecipientAddress).Hex(),
		Asset:              sdk.NewCoin(response.Asset.Name, sdk.NewInt(stakingAmount)),
	}, nil
}

// func (client *BtcClient) decodeSourceTxConfirmationEvent(tx *BTCTxReceipt) (*chainsTypes.SourceTxConfirmationEvent, error) {
// 	if len(tx.MsgTx.TxOut) < MinNumberOfOutputs {
// 		return nil, ErrInvalidTxOutCount
// 	}

// 	embeddedDataTxOut := tx.MsgTx.TxOut[EmbeddedDataOutputIndex]
// 	if embeddedDataTxOut == nil || embeddedDataTxOut.PkScript == nil || embeddedDataTxOut.PkScript[0] != txscript.OP_RETURN {
// 		return nil, ErrInvalidOpReturn
// 	}

// 	output, err := vault.ParseVaultEmbeddedData(embeddedDataTxOut.PkScript)
// 	if err != nil || output == nil {
// 		return nil, ErrInvalidOpReturnData
// 	}

// 	var stakingAmount int64 = tx.MsgTx.TxOut[StakingOutputIndex].Value

// 	destinationChain := chain.NewChainInfoFromBytes(output.DestinationChain)
// 	if destinationChain == nil {
// 		return nil, ErrInvalidDestinationChain
// 	}

// 	var destinationContractAddress chainsTypes.Address
// 	err = destinationContractAddress.Unmarshal(output.DestinationTokenAddress)
// 	if err != nil {
// 		return nil, err
// 	}

// 	payload, payloadHash, err := encode.SafeCalculateDestPayload(uint64(stakingAmount), tx.MsgTx.TxID(), output.DestinationRecipientAddress)
// 	if err != nil {
// 		return nil, ErrInvalidPayloadHash
// 	}

// 	var destinationRecipientAddress chainsTypes.Address
// 	err = destinationRecipientAddress.Unmarshal(output.DestinationRecipientAddress)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &chainsTypes.SourceTxConfirmationEvent{
// 		Sender:                      tx.PrevTxOuts[0].ScriptPubKey.Address,
// 		DestinationChain:            nexus.ChainName(destinationChain.ToBytes().String()),
// 		Amount:                      uint64(stakingAmount),
// 		Asset:                       "satoshi",
// 		PayloadHash:                 chainsTypes.Hash(payloadHash),
// 		Payload:                     payload,
// 		DestinationContractAddress:  chainsTypes.Address(destinationContractAddress).Hex(),
// 		DestinationRecipientAddress: chainsTypes.Address(destinationRecipientAddress).Hex(),
// 	}, nil
// }
