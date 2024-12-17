package btc

import (
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/monads/results"
	"github.com/scalarorg/scalar-core/vald/xchain"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

type BtcClient struct {
	client                    *rpcclient.Client
	cfg                       *rpcclient.ConnConfig
	blockHeightCache          *BlockHeightCache
	latestFinalizedBlockCache xchain.LatestFinalizedBlockCache
}

type BTCTxReceipt struct {
	Raw        btcjson.TxRawResult
	PrevTxOuts []*btcjson.GetTxOutResult
	MsgTx      *wire.MsgTx
}

type BTCTxResult = results.Result[xchain.TxReceipt]

var _ xchain.Client = &BtcClient{}

func NewClient(cfg *types.BTCConfig) (xchain.Client, error) {
	rpcConfig := mapBTCConfigToRPCConfig(cfg)
	clog.Red("NewClient", "rpcConfig", rpcConfig)
	rpcClient, error := rpcclient.New(rpcConfig, nil)
	if error != nil {
		return nil, error
	}

	blockHeightCache := NewBlockHeightCache()
	latestFinalizedBlockCache := xchain.NewLatestFinalizedBlockCache()

	client := &BtcClient{
		client:                    rpcClient,
		cfg:                       rpcConfig,
		blockHeightCache:          blockHeightCache,
		latestFinalizedBlockCache: latestFinalizedBlockCache,
	}

	return client, nil
}

func mapBTCConfigToRPCConfig(cfg *types.BTCConfig) *rpcclient.ConnConfig {
	params := cfg.Chain.String()
	if params == types.BtcChainName[types.Testnet4BtcChain] {
		params = chaincfg.TestNet3Params.Name
	}

	return &rpcclient.ConnConfig{
		Host:                 cfg.RpcHost,
		User:                 cfg.RpcUser,
		Pass:                 cfg.RpcPass,
		Params:               params,
		DisableTLS:           cfg.DisableTLS,
		DisableConnectOnNew:  cfg.DisableConnectOnNew,
		DisableAutoReconnect: cfg.DisableAutoReconnect,
		HTTPPostMode:         cfg.HttpPostMode,
	}
}

func (c *BtcClient) Close() {
	c.client.Shutdown()
}
