package btc

import (
	"encoding/json"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/log"
)

func (c *BtcClient) logger(keyvals ...any) log.Logger {
	keyvals = append([]any{"rpc", "btc"}, keyvals...)
	return log.WithKeyVals(keyvals...)
}

func (client *BtcClient) isFinalized(txReceipt *btcjson.GetTransactionResult, confHeight uint64) (bool, error) {
	blockHeightCache := client.blockHeightCache.Get(txReceipt.BlockHash)
	if blockHeightCache == nil {
		blockHeight, err := client.GetBlockHeight(txReceipt.BlockHash)
		if err != nil {
			return false, err
		}

		clog.Cyanf("block_height_not_found, block_hash: %s, conf_block_height: %d, block_height: %d", txReceipt.BlockHash, confHeight, blockHeight)
		client.blockHeightCache.Set(txReceipt.BlockHash, blockHeight)
		blockHeightCache = &blockHeight
	} else {
		clog.Cyanf("[BTC] blockHeightCache already exists, block_hash: %s, conf_block_height: %d, block_height: %d", txReceipt.BlockHash, confHeight, *blockHeightCache)
	}

	latestFinalizedBlockCache := client.latestFinalizedBlockCache.Get()
	clog.Cyanf("[BTC] blockHeightCache: %d", *blockHeightCache)
	clog.Cyanf("[BTC] latestFinalizedBlockCache: %d", latestFinalizedBlockCache)

	if latestFinalizedBlockCache != 0 && latestFinalizedBlockCache >= *blockHeightCache {
		return true, nil
	}

	latestFinalizedBlockHeight, err := client.LatestFinalizedBlockHeight(confHeight)
	if err != nil {
		return false, err
	}

	client.blockHeightCache.Set(txReceipt.BlockHash, latestFinalizedBlockHeight)
	client.latestFinalizedBlockCache.Set(latestFinalizedBlockHeight)

	// This is a rare case, but it can happen if the block height is not updated in the cache
	if latestFinalizedBlockHeight < *blockHeightCache {
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
