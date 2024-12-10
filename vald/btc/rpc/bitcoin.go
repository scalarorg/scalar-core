package rpc

import (
	"sync"

	"github.com/axelarnetwork/utils/log"
	"github.com/axelarnetwork/utils/monads/results"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

type BtcClient struct {
	client *rpcclient.Client
	cfg    *rpcclient.ConnConfig
}

func NewBtcClient(cfg *rpcclient.ConnConfig) (*BtcClient, error) {
	client, err := rpcclient.New(cfg, nil)
	if err != nil {
		return nil, err
	}

	return &BtcClient{
		client: client,
		cfg:    cfg,
	}, nil
}

func (c *BtcClient) logger(keyvals ...any) log.Logger {
	keyvals = append([]any{"rpc", "btc"}, keyvals...)
	return log.WithKeyVals(keyvals...)
}

func (c *BtcClient) Close() {
	c.client.Shutdown()
}

type TxReceipt struct {
	Data btcjson.TxRawResult
}

func (c *BtcClient) GetTransaction(txID types.Hash) (TxReceipt, error) {
	var tx TxReceipt
	// txBytes := txID.Bytes()
	// txBytesReverse := make([]byte, len(txBytes))
	// for i, b := range txBytes {
	// 	txBytesReverse[len(txBytes)-1-i] = b
	// }

	// txHash, err := chainhash.NewHash(txBytesReverse)

	// if err != nil {
	// 	c.logger.Errorf("failed to create BTC chainhash from txID", "txID", txID, "error", err)
	// 	return tx, err
	// }

	txMetadata, err := c.client.GetRawTransactionVerbose(txID.IntoRef())
	if err != nil {
		c.logger("failed to get BTC transaction", "txID", txID, "error", err)
	} else {
		tx.Data = *txMetadata
	}

	return tx, err
}

func (c *BtcClient) GetTransactions(txIDs []types.Hash) ([]TxResult, error) {
	txs := make([]TxResult, len(txIDs))
	var wg sync.WaitGroup

	for i := range txIDs {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			tx, err := c.GetTransaction(txIDs[index])

			var txResult TxResult
			if err != nil {
				txResult = TxResult(results.FromErr[TxReceipt](err))
			} else {
				txResult = TxResult(results.FromOk(tx))
			}

			txs[index] = txResult

		}(i)
	}

	wg.Wait()
	return txs, nil
}

func (c *BtcClient) LatestFinalizedBlockHeight(_ uint64) (uint64, error) {
	_, height, err := c.client.GetBestBlock()
	if err != nil {
		return 0, err
	}

	return uint64(height), nil
}

func (c *BtcClient) GetBlockHeight(blockHash string) (uint64, error) {
	chainhashBlockHash, err := chainhash.NewHashFromStr(blockHash)
	if err != nil {
		return 0, err
	}

	block, err := c.client.GetBlockHeaderVerbose(chainhashBlockHash)
	if err != nil {
		return 0, err
	}

	return uint64(block.Height), nil
}
