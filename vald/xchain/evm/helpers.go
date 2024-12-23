package evm

import (
	"fmt"

	"github.com/scalarorg/scalar-core/utils/log"
)

func (c *EthereumClient) logger(keyvals ...any) log.Logger {
	keyvals = append([]any{"rpc", "eth"}, keyvals...)
	return log.WithKeyVals(keyvals...)
}

func (client *EthereumClient) isFinalized(txReceipt ETHTxReceipt, confHeight uint64) (bool, error) {
	if txReceipt.BlockNumber == nil {
		return false, fmt.Errorf("block number of tx receipt is nil")
	}

	latestFinalizedBlockCache := client.latestFinalizedBlockCache.Get()
	if latestFinalizedBlockCache != 0 && latestFinalizedBlockCache >= txReceipt.BlockNumber.Uint64() {
		return true, nil
	}

	latestFinalizedBlockHeight, err := client.LatestFinalizedBlockHeight(confHeight)
	if err != nil {
		return false, err
	}

	client.latestFinalizedBlockCache.Set(latestFinalizedBlockHeight)

	// This is a rare case, but it can happen if the block height is not updated in the cache
	if latestFinalizedBlockHeight < txReceipt.BlockNumber.Uint64() {
		return false, nil
	}

	return true, nil
}

// func (c *EthereumClient) HeaderByNumber(ctx context.Context, number uint64) (*Header, error) {
// 	var head *Header
// 	err := c.rpc.CallContext(ctx, &head, "eth_getBlockByNumber", toBlockNumArg(number), false)
// 	if err == nil && head == nil {
// 		err = ethereum.NotFound
// 	}

// 	return head, err
// }
