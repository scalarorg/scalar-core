package btc

import (
	"github.com/btcsuite/btcd/btcjson"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/log"
)

func (c *BtcClient) logger(keyvals ...any) log.Logger {
	keyvals = append([]any{"rpc", "btc"}, keyvals...)
	return log.WithKeyVals(keyvals...)
}

func (client *BtcClient) isFinalized(txReceipt btcjson.TxRawResult, confHeight uint64) (bool, error) {
	blockHeightCache := client.blockHeightCache.Get(txReceipt.BlockHash)
	if blockHeightCache == nil {
		clog.Redf("block_height_not_found, block_hash: %s, block_height: %d", txReceipt.BlockHash, confHeight)
		blockHeight, err := client.GetBlockHeight(txReceipt.BlockHash)
		if err != nil {
			return false, err
		}
		client.blockHeightCache.Set(txReceipt.BlockHash, blockHeight)
		blockHeightCache = &blockHeight
	}

	if client.latestFinalizedBlockCache.Get() != 0 && client.latestFinalizedBlockCache.Get() >= *blockHeightCache {
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
