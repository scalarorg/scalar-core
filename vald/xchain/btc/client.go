package btc

import (
	"fmt"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	btcChain "github.com/scalarorg/bitcoin-vault/go-utils/btc"
	"github.com/scalarorg/scalar-core/utils/monads/results"
	"github.com/scalarorg/scalar-core/vald/config"
	"github.com/scalarorg/scalar-core/vald/xchain/common"
)

type BtcClient struct {
	client                    *rpcclient.Client
	cfg                       *rpcclient.ConnConfig
	blockCache                *BlockCache
	latestFinalizedBlockCache common.LatestFinalizedBlockCache
}
type BTCTxReceipt struct {
	Raw *btcjson.TxRawResult
	// blockIndex field in "gettransaction" and the index of the transaction in the block in "getblock"
	TransactionIndex int
	PrevTxOuts       []*btcjson.Vout
	MsgTx            *wire.MsgTx
}

type BTCTxResult = results.Result[common.TxReceipt]

var _ common.Client = &BtcClient{}

func NewClient(cfg *config.BTCConfig) (common.Client, error) {
	rpcConfig := MapBTCConfigToRPCConfig(cfg)
	rpcClient, error := rpcclient.New(rpcConfig, nil)
	if error != nil {
		return nil, error
	}

	blockCache := NewBlockCache()
	latestFinalizedBlockCache := common.NewLatestFinalizedBlockCache()

	client := &BtcClient{
		client:                    rpcClient,
		cfg:                       rpcConfig,
		blockCache:                blockCache,
		latestFinalizedBlockCache: latestFinalizedBlockCache,
	}

	return client, nil
}

func validateChain(cfg *config.BTCConfig) error {
	if btcChain.BtcChainsRecords().GetChainParamsByName(cfg.Chain) == nil {
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

	if params == "testnet4" {
		params = chaincfg.TestNet3Params.Name // TODO: update this field when btc supports
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
