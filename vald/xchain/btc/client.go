package btc

import (
	"fmt"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/scalarorg/scalar-core/utils/monads/results"
	"github.com/scalarorg/scalar-core/vald/config"
	"github.com/scalarorg/scalar-core/vald/xchain"
	btcChain "github.com/scalarorg/bitcoin-vault/go-utils/chain"
)

type BtcClient struct {
	client                    *rpcclient.Client
	cfg                       *rpcclient.ConnConfig
	blockHeightCache          *BlockHeightCache
	latestFinalizedBlockCache xchain.LatestFinalizedBlockCache
}
type BTCTxReceipt struct {
	Raw        btcjson.TxRawResult
	PrevTxOuts []*btcjson.Vout
	MsgTx      *wire.MsgTx
}

type BTCTxResult = results.Result[xchain.TxReceipt]

var _ xchain.Client = &BtcClient{}

func NewClient(cfg *config.BTCConfig) (xchain.Client, error) {
	rpcConfig := mapBTCConfigToRPCConfig(cfg)
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

func validateChain(cfg *config.BTCConfig) error {
	_, ok := btcChain.BtcChainConfigValueInt[cfg.Chain]
	if !ok {
		return fmt.Errorf("invalid chain %s", cfg.Chain)
	}
	return nil
}

func mapBTCConfigToRPCConfig(cfg *config.BTCConfig) *rpcclient.ConnConfig {
	err := validateChain(cfg)
	if err != nil {
		panic("invalid btc chain when setting the params")
	}

	params := cfg.Chain

	if params == btcChain.ChaincfgTestnet4ParamsName {
		params = chaincfg.TestNet3Params.Name
	}

	return &rpcclient.ConnConfig{
		Host:                 cfg.RPCHost,
		User:                 cfg.RPCUser,
		Pass:                 cfg.RPCPass,
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
