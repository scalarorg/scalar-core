package btc

import (
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/log"
)

func (c *BtcClient) logger(keyvals ...any) log.Logger {
	keyvals = append([]any{"rpc", "btc"}, keyvals...)
	return log.WithKeyVals(keyvals...)
}

func (client *BtcClient) isFinalized(txReceipt *btcjson.TxRawResult, confHeight uint64) (bool, error) {
	block := client.blockCache.Get(txReceipt.BlockHash)
	if block == nil {
		return false, fmt.Errorf("block not found")
	}

	blockHeight := block.Height
	clog.Cyanf("[BTC] block already exists, block_hash: %s, conf_block_height: %d, block_height: %d", txReceipt.BlockHash, confHeight, blockHeight)

	latestFinalizedBlockCache := client.latestFinalizedBlockCache.Get()
	clog.Cyanf("[BTC] latestFinalizedBlockCache: %d", latestFinalizedBlockCache)

	if latestFinalizedBlockCache != 0 && latestFinalizedBlockCache >= uint64(blockHeight) {
		return true, nil
	}

	latestFinalizedBlockHeight, err := client.LatestFinalizedBlockHeight(confHeight)
	if err != nil {
		return false, err
	}

	client.latestFinalizedBlockCache.Set(latestFinalizedBlockHeight)

	// This is a rare case, but it can happen if the block height is not updated in the cache
	if latestFinalizedBlockHeight < uint64(blockHeight) {
		return false, nil
	}

	return true, nil
}

// The reason for this function is the GetBlockChainInfo() or btcd parsing error
func (c *BtcClient) getBlockChainInfo() (*btcjson.GetBlockChainInfoResult, error) {
	rawResponse, err := c.client.RawRequest("getblockchaininfo", nil)
	if err != nil {
		return nil, err
	}

	var chainInfo btcjson.GetBlockChainInfoResult
	if err := json.Unmarshal(rawResponse, &chainInfo); err != nil {
		return nil, err
	}

	return &chainInfo, nil
}
