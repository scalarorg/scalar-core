package btc

import (
	"errors"

	"github.com/btcsuite/btcd/txscript"
	vault "github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/bitcoin-vault/go-utils/encode"
	"github.com/scalarorg/scalar-core/utils/clog"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var (
	StakingOutputIndex      = 0
	EmbeddedDataOutputIndex = 1
)

var (
	ErrInvalidTxOutCount       = errors.New("btcLocking tx must have at least 3 outputs")
	ErrInvalidOpReturn         = errors.New("transaction does not have expected payload op return output")
	ErrInvalidOpReturnData     = errors.New("cannot parse payload op return data")
	ErrInvalidTxId             = errors.New("failed to decode tx id")
	ErrInvalidPayloadHash      = errors.New("failed to get payload hash")
	ErrInvalidDestinationChain = errors.New("failed to parse destination chain")
)

const (
	MinNumberOfOutputs = 2
)

func (client *BtcClient) decodeStakingTransaction(tx *BTCTxReceipt) (*chainsTypes.TxConfirmationEvent, error) {
	if len(tx.MsgTx.TxOut) < MinNumberOfOutputs {
		return nil, ErrInvalidTxOutCount
	}

	embeddedDataTxOut := tx.MsgTx.TxOut[EmbeddedDataOutputIndex]
	if embeddedDataTxOut == nil || embeddedDataTxOut.PkScript == nil || embeddedDataTxOut.PkScript[0] != txscript.OP_RETURN {
		return nil, ErrInvalidOpReturn
	}

	output, err := vault.ParseVaultEmbeddedData(embeddedDataTxOut.PkScript)
	if err != nil || output == nil {
		return nil, ErrInvalidOpReturnData
	}

	txHash := tx.MsgTx.TxHash()
	txBytes := txHash.CloneBytes()
	txBytes = txBytes[:32]
	var copiedTxHash [32]byte
	copy(copiedTxHash[:], txBytes)

	var stakingAmount int64 = tx.MsgTx.TxOut[StakingOutputIndex].Value

	destinationChain := chain.NewChainInfoFromBytes(output.DestinationChain)
	if destinationChain == nil {
		return nil, ErrInvalidDestinationChain
	}

	var destinationContractAddress chainsTypes.Address
	err = destinationContractAddress.Unmarshal(output.DestinationContractAddress)
	if err != nil {
		return nil, err
	}

	var destinationRecipientAddress chainsTypes.Address
	err = destinationRecipientAddress.Unmarshal(output.DestinationRecipientAddress)
	if err != nil {
		return nil, err
	}

	payload, payloadHash, err := encode.CalculateStakingPayloadHash(destinationRecipientAddress, uint64(stakingAmount), copiedTxHash)
	if err != nil {
		return nil, ErrInvalidPayloadHash
	}

	// , err := chainsTypes.HashFromBytes(payloadHash)
	// if err != nil {
	// 	return chainsTypes.ConfirmationEvent{}, ErrInvalidPayloadHash
	// }chainHash

	clog.Redf("[VALD] Payload hash %+v\n", payloadHash)
	clog.Redf("[VALD] Minting amount %+v\n", stakingAmount)
	// clog.Redf("[VALD] Chain hash %+v\n", chainHash)
	clog.Redf("[VALD] tx: %+v", tx)

	return &chainsTypes.TxConfirmationEvent{
		Sender:                      tx.PrevTxOuts[0].ScriptPubKey.Address, // TODO: Fix hard coded
		DestinationChain:            nexus.ChainName(destinationChain.ToBytes().String()),
		Amount:                      uint64(stakingAmount),
		Asset:                       "satoshi", // TODO: Fix hard coded
		PayloadHash:                 chainsTypes.Hash(payloadHash),
		Payload:                     payload,
		DestinationContractAddress:  chainsTypes.Address(destinationContractAddress).Hex(),
		DestinationRecipientAddress: chainsTypes.Address(destinationRecipientAddress).Hex(),
	}, nil
}
