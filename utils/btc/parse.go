package btc

import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	vault "github.com/scalarorg/bitcoin-vault/ffi/go"
	"github.com/scalarorg/go-common/chain"
	go_utils "github.com/scalarorg/go-common/types"
)

var (
	EmbeddedDataOutputIndex = 0
	LockingOutputIndex      = 1
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

type VaultInfo struct {
	TxID                        []byte
	StakingAmount               int64
	ScriptPubkey                []byte
	DestinationChain            string
	DestinationRecipientAddress []byte
	DestinationTokenAddress     string //Hex encoded
	Vout                        uint32
	MsgTx                       *wire.MsgTx
}

func ParseTx(rawTx []byte) (*VaultInfo, error) {
	reader := bytes.NewReader(rawTx)
	var msgTx wire.MsgTx
	if err := msgTx.Deserialize(reader); err != nil {
		return nil, err
	}
	txHash := msgTx.TxHash()
	embeddedDataTxOut := msgTx.TxOut[EmbeddedDataOutputIndex]
	if embeddedDataTxOut == nil || embeddedDataTxOut.PkScript == nil || embeddedDataTxOut.PkScript[0] != txscript.OP_RETURN {
		return nil, ErrInvalidOpReturn
	}

	output, err := vault.ParseVaultEmbeddedData(embeddedDataTxOut.PkScript)
	if err != nil || output == nil {
		return nil, ErrInvalidOpReturnData
	}

	if output.TransactionType != go_utils.TransactionTypeLocking {
		return nil, ErrInvalidTransactionType
	}

	var stakingAmount int64 = msgTx.TxOut[LockingOutputIndex].Value
	var scriptPubkey []byte = msgTx.TxOut[LockingOutputIndex].PkScript
	destinationChain := chain.NewChainInfoFromBytes(output.DestinationChain)
	if destinationChain == nil {
		return nil, ErrInvalidDestinationChain
	}

	return &VaultInfo{
		TxID:                        txHash.CloneBytes(),
		StakingAmount:               stakingAmount,
		ScriptPubkey:                scriptPubkey,
		DestinationChain:            destinationChain.ToBytes().String(),
		DestinationRecipientAddress: output.DestinationRecipientAddress,
		DestinationTokenAddress:     hex.EncodeToString(output.DestinationTokenAddress),
		Vout:                        uint32(LockingOutputIndex),
		MsgTx:                       &msgTx,
	}, nil
}
