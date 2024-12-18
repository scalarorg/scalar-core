package btc

import (
	"errors"

	"github.com/btcsuite/btcd/txscript"
	vault "github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	evmUtils "github.com/scalarorg/bitcoin-vault/go-utils/evm"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/log"
	btcTypes "github.com/scalarorg/scalar-core/x/btc/types"
	evmTypes "github.com/scalarorg/scalar-core/x/evm/types"
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

func (client *BtcClient) decodeStakingTransaction(tx *BTCTxReceipt) (btcTypes.EventStakingTx, error) {
	log.Infof("Decoding BTC transaction %+v\n", tx)

	if len(tx.MsgTx.TxOut) < MinNumberOfOutputs {
		return btcTypes.EventStakingTx{}, ErrInvalidTxOutCount
	}

	embeddedDataTxOut := tx.MsgTx.TxOut[EmbeddedDataOutputIndex]
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

	txHash := tx.MsgTx.TxHash()
	txBytes := txHash.CloneBytes()
	txBytes = txBytes[:32]
	var copiedTxHash [32]byte
	copy(copiedTxHash[:], txBytes)

	var stakingAmount int64 = tx.MsgTx.TxOut[StakingOutputIndex].Value

	_, payloadHash, err := evmUtils.CalculateStakingPayloadHash(stakingMetadata.DestinationRecipientAddress, stakingAmount, copiedTxHash)
	if err != nil {
		return btcTypes.EventStakingTx{}, ErrInvalidPayloadHash
	}

	chainHash, err := btcTypes.HashFromBytes(payloadHash)
	if err != nil {
		return btcTypes.EventStakingTx{}, ErrInvalidPayloadHash
	}

	clog.Redf("[VALD] Staking metadata %+v\n", stakingMetadata)
	clog.Redf("[VALD] Payload hash %+v\n", payloadHash)
	clog.Redf("[VALD] Minting amount %+v\n", stakingAmount)
	clog.Redf("[VALD] Chain hash %+v\n", chainHash)
	clog.Redf("[VALD] tx: %+v", tx)

	// tx.MsgTx.TxHash()

	h := btcTypes.FromChainHash(tx.MsgTx.TxHash())
	stakingMetadata.StakingOutpoint = btcTypes.OutPoint{
		Hash:  &h,
		Index: uint32(StakingOutputIndex),
	}

	return btcTypes.EventStakingTx{
		Sender:      tx.PrevTxOuts[0].ScriptPubKey.Address, // TODO: Fix hard coded
		Amount:      uint64(stakingAmount),
		Asset:       "satoshi", // TODO: Fix hard coded
		Metadata:    *stakingMetadata,
		PayloadHash: chainHash,
	}, nil
}

func mapOutputToEventStakingTx(output *vault.VaultReturnTxOutput) (*btcTypes.EventStakingTx_StakingTxMetadata, error) {

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

	parsedDestinationChain := chain.NewChainInfoFromBytes(output.DestinationChain)
	if parsedDestinationChain == nil {
		return nil, ErrInvalidDestinationChain
	}

	return &btcTypes.EventStakingTx_StakingTxMetadata{
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
