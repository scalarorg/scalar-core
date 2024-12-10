package rpc

import (
	"github.com/axelarnetwork/utils/monads/results"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

type TxResult results.Result[TxReceipt]

type Client interface {
	GetTransaction(txID types.Hash) (TxReceipt, error)
	GetTransactions(txIDs []types.Hash) ([]TxResult, error)
	LatestFinalizedBlockHeight(confHeight uint64) (uint64, error)
	GetBlockHeight(blockHash string) (uint64, error)
	Close()
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
		Host: cfg.RPCAddr,
		User: cfg.RPCUser,
		Pass: cfg.RPCPassword,
	}
}
