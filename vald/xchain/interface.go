package xchain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/monads/results"
	btcTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

type TxReceipt interface{}

type TxResult = results.Result[TxReceipt]

type Hash = [32]byte

type Client interface {
	ProcessSourceTxsConfirmation(event *btcTypes.EventConfirmSourceTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error)
	ProcessDestinationTxsConfirmation(event *btcTypes.EventConfirmDestTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error)
	GetTxReceiptsIfFinalized(txIDs []Hash, confHeight uint64) ([]TxResult, error)
	GetTransaction(txID Hash) (TxResult, error)
	GetTransactions(txIDs []Hash) ([]TxResult, error)
	LatestFinalizedBlockHeight(confHeight uint64) (uint64, error)
	GetBlockHeight(blockHash string) (uint64, error)
	Close()
}
