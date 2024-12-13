package rpc

import (
	"bytes"
	"encoding/hex"
	"errors"
	"sync"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/scalarorg/scalar-core/utils/log"
	"github.com/scalarorg/scalar-core/utils/monads/results"
	"github.com/scalarorg/scalar-core/utils/slices"
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
	Raw        btcjson.TxRawResult
	PrevTxOuts []*btcjson.GetTxOutResult
	MsgTx      *wire.MsgTx
}

func (c *BtcClient) GetTransaction(txID types.Hash) TxResult {
	var tx TxReceipt
	txMetadata, err := c.client.GetRawTransactionVerbose(txID.IntoRef())
	if err != nil {
		c.logger("failed to get BTC transaction", "txID", txID, "error", err)
		return TxResult(results.FromErr[TxReceipt](err))
	} else {
		txRaw, err := hex.DecodeString(txMetadata.Hex)
		if err != nil {
			c.logger("failed to decode hex string", "txID", txID, "error", err)
			return TxResult(results.FromErr[TxReceipt](err))
		}

		msgTx := wire.NewMsgTx(wire.TxVersion)
		err = msgTx.Deserialize(bytes.NewReader(txRaw))
		if err != nil {
			c.logger("failed to parse transaction", "txID", txID, "error", err)
			return TxResult(results.FromErr[TxReceipt](err))
		}

		prevTxOuts, err := c.GetTxOuts(slices.Map(msgTx.TxIn, func(txIn *wire.TxIn) wire.OutPoint {
			return txIn.PreviousOutPoint
		}))
		if err != nil {
			c.logger("failed to get BTC transaction", "txID", txID, "error", err)
		}

		tx.Raw = *txMetadata
		tx.PrevTxOuts = prevTxOuts
		tx.MsgTx = msgTx
	}

	return TxResult(results.FromOk(tx))
}

func (c *BtcClient) GetTransactions(txIDs []types.Hash) ([]TxResult, error) {
	txs := make([]TxResult, len(txIDs))
	var wg sync.WaitGroup

	for i := range txIDs {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			txs[index] = c.GetTransaction(txIDs[index])
		}(i)
	}

	wg.Wait()
	if slices.Any(txs, func(tx TxResult) bool { return tx.IsErr() }) {
		return nil, errors.New("failed to get transactions")
	}

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

func (c *BtcClient) GetTxOut(outpoint wire.OutPoint) (*btcjson.GetTxOutResult, error) {
	return c.client.GetTxOut(&outpoint.Hash, outpoint.Index, false)
}

func (c *BtcClient) GetTxOuts(outpoints []wire.OutPoint) ([]*btcjson.GetTxOutResult, error) {
	txOuts := make([]*btcjson.GetTxOutResult, len(outpoints))
	var wg sync.WaitGroup

	for i := range outpoints {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			txOut, err := c.GetTxOut(outpoints[index])
			if err != nil {
				txOuts[index] = nil
			} else {
				txOuts[index] = txOut
			}
		}(i)
	}

	wg.Wait()
	return txOuts, nil
}
