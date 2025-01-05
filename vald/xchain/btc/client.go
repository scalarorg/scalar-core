package btc

import (
	"fmt"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	btcChain "github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/scalar-core/utils/monads/results"
	"github.com/scalarorg/scalar-core/vald/config"
	"github.com/scalarorg/scalar-core/vald/xchain/common"
)

type BtcClient struct {
	client                    *rpcclient.Client
	cfg                       *rpcclient.ConnConfig
	blockHeightCache          *BlockHeightCache
	latestFinalizedBlockCache common.LatestFinalizedBlockCache
}
type BTCTxReceipt struct {
	Raw        *btcjson.GetTransactionResult
	PrevTxOuts []*btcjson.Vout
	MsgTx      *wire.MsgTx
}

type BTCTxResult = results.Result[common.TxReceipt]

var _ common.Client = &BtcClient{}

func NewClient(cfg *config.BTCConfig) (common.Client, error) {
	rpcConfig := MapBTCConfigToRPCConfig(cfg)
	rpcClient, error := rpcclient.New(rpcConfig, nil)
	if error != nil {
		return nil, error
	}

	blockHeightCache := NewBlockHeightCache()
	latestFinalizedBlockCache := common.NewLatestFinalizedBlockCache()

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

func MapBTCConfigToRPCConfig(cfg *config.BTCConfig) *rpcclient.ConnConfig {
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
