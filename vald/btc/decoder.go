package btc

import (
	"bytes"
	"encoding/hex"
	"errors"

	evmTypes "github.com/axelarnetwork/axelar-core/x/evm/types"
	"github.com/axelarnetwork/utils/log"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/ethereum/go-ethereum/common"
	vault "github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	evmUtils "github.com/scalarorg/bitcoin-vault/go-utils/evm"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/vald/btc/rpc"
	btcTypes "github.com/scalarorg/scalar-core/x/btc/types"
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

func DecodeStakingTransaction(tx *rpc.TxReceipt) (btcTypes.EventStakingTx, error) {
	log.Infof("Decoding BTC transaction %+v\n", tx)

	// Decode the hex string into bytes
	txRaw, err := hex.DecodeString(tx.Data.Hex)
	if err != nil {
		log.Errorf("Failed to decode hex string %v", err)
		return btcTypes.EventStakingTx{}, err
	}

	// Parse the transaction
	msgTx := wire.NewMsgTx(wire.TxVersion)
	err = msgTx.Deserialize(bytes.NewReader(txRaw))
	if err != nil {
		log.Errorf("Failed to parse transaction %v", err)
		return btcTypes.EventStakingTx{}, err
	}

	if len(msgTx.TxOut) < MinNumberOfOutputs {
		return btcTypes.EventStakingTx{}, ErrInvalidTxOutCount
	}

	embeddedDataTxOut := msgTx.TxOut[1]
	if embeddedDataTxOut == nil || embeddedDataTxOut.PkScript == nil || embeddedDataTxOut.PkScript[0] != txscript.OP_RETURN {
		return btcTypes.EventStakingTx{}, ErrInvalidOpReturn
	}

	output, err := vault.ParseVaultEmbeddedData(embeddedDataTxOut.PkScript)
	if err != nil || output == nil {
		return btcTypes.EventStakingTx{}, ErrInvalidOpReturnData
	}

	stakingMetadata, err := mapOutputToEventStakingTx(output)
	if err != nil {
		return btcTypes.EventStakingTx{}, err
	}

	var txIdBytes [32]byte
	txId := msgTx.TxID()
	txBytes, err := hex.DecodeString(txId)
	if err != nil {
		return btcTypes.EventStakingTx{}, ErrInvalidTxId
	}
	copy(txIdBytes[:], txBytes)

	var mintingAmount int64 = msgTx.TxOut[0].Value

	_, payloadHash, err := evmUtils.CalculateStakingPayloadHash(stakingMetadata.DestinationRecipientAddress, mintingAmount, txIdBytes)
	if err != nil {
		return btcTypes.EventStakingTx{}, ErrInvalidPayloadHash
	}

	clog.Redf("Decoded BTC transaction %+v\n", stakingMetadata)
	clog.Redf("Payload hash %+v\n", payloadHash)
	clog.Redf("Minting amount %+v\n", mintingAmount)
	clog.Redf("Tx ID %+v\n", txId)

	return btcTypes.EventStakingTx{
		PrevOutPoint: msgTx.TxIn[0].PreviousOutPoint.String(),
		Amount:       uint64(mintingAmount),
		Asset:        "satoshi", // TODO: Fix hard coded
		Metadata:     *stakingMetadata,
		PayloadHash:  evmTypes.Hash(common.BytesToHash(payloadHash)),
	}, nil

}

func mapOutputToEventStakingTx(output *vault.VaultReturnTxOutput) (*btcTypes.StakingTxMetadata, error) {

	var vaultTag btcTypes.VaultTag
	err := vaultTag.Unmarshal(output.Tag)
	if err != nil {
		return nil, err
	}

	var destinationContractAddress evmTypes.Address
	err = destinationContractAddress.Unmarshal(output.DestinationContractAddress)
	if err != nil {
		return nil, err
	}

	var destinationRecipientAddress evmTypes.Address
	err = destinationRecipientAddress.Unmarshal(output.DestinationRecipientAddress)
	if err != nil {
		return nil, err
	}

	parsedDestinationChain := chain.NewDestinationChainFromBytes(output.DestinationChain)
	if parsedDestinationChain == nil {
		return nil, ErrInvalidDestinationChain
	}

	return &btcTypes.StakingTxMetadata{
		Tag:                         vaultTag,
		Version:                     btcTypes.VersionFromInt(int(output.Version)),
		NetworkId:                   btcTypes.NetworkKind(output.NetworkID),
		Flags:                       output.Flags,
		ServiceTag:                  output.ServiceTag,
		HaveOnlyCovenants:           output.HaveOnlyCovenants,
		CovenantQuorum:              output.CovenantQuorum,
		DestinationChainType:        parsedDestinationChain.ChainType,
		DestinationChainId:          parsedDestinationChain.ChainID,
		DestinationContractAddress:  destinationContractAddress,
		DestinationRecipientAddress: destinationRecipientAddress,
	}, nil

}
