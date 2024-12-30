package common

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/scalarorg/scalar-core/utils/monads/results"
	btcTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

type TxReceipt interface{}

type TxResult = results.Result[TxReceipt]

type Hash = common.Hash

func HashToString(h Hash) string {
	return hex.EncodeToString(h[:])
}

func HashToChainHash(h Hash) chainhash.Hash {
	var reversedTxID [32]byte
	for i := 0; i < 32; i++ {
		reversedTxID[i] = h[31-i]
	}
	return chainhash.Hash(reversedTxID)
}

// TxReceiptResult is a custom type that allows moq to correctly generate the mock for
// results.TxReceiptResult with *types.Receipt.
type TxReceiptResult results.Result[types.Receipt]

type Client interface {
	ProcessSourceTxsConfirmation(event *btcTypes.EventConfirmSourceTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error)
	GetTxReceiptsIfFinalized(txIDs []Hash, confHeight uint64) ([]TxResult, error)
	GetTransactions(txIDs []Hash) ([]TxResult, error)
	GetTransaction(txID Hash) (TxResult, error)
	LatestFinalizedBlockHeight(confHeight uint64) (uint64, error)
	Close()
}
