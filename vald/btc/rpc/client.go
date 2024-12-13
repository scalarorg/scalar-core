package rpc

import (
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/scalarorg/scalar-core/utils/monads/results"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

type TxResult = results.Result[TxReceipt]

type Client interface {
	GetTransaction(txID types.Hash) TxResult
	GetTransactions(txIDs []types.Hash) ([]TxResult, error)
	LatestFinalizedBlockHeight(confHeight uint64) (uint64, error)
	GetBlockHeight(blockHash string) (uint64, error)
	Close()
	GetTxOut(outpoint wire.OutPoint) (*btcjson.GetTxOutResult, error)
}

func NewClient(cfg *types.BTCConfig) (Client, error) {
	client, error := NewBtcClient(mapBTCConfigToRPCConfig(cfg))

	if error != nil {
		return nil, error
	}

	return client, nil
}

func mapBTCConfigToRPCConfig(cfg *types.BTCConfig) *rpcclient.ConnConfig {
	return &rpcclient.ConnConfig{
		Host:   cfg.RpcHost,
		User:   cfg.RpcUser,
		Pass:   cfg.RpcPass,
		Params: cfg.Chain.String(),
	}
}
