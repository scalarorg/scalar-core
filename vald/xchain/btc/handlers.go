package btc

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/errors"
	"github.com/scalarorg/scalar-core/utils/monads/results"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/vald/xchain/common"
)

func (client *BtcClient) GetTxReceiptsIfFinalized(txIDs []common.Hash, confHeight uint64) ([]BTCTxResult, error) {
	txResults, err := client.GetTransactions(txIDs)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			errors.With(err, "tx_ids", txIDs),
			"cannot get transaction receipts",
		)
	}

	return slices.Map(txResults, func(receipt BTCTxResult) results.Result[common.TxReceipt] {
		return results.Pipe(results.Result[common.TxReceipt](receipt), func(receipt common.TxReceipt) results.Result[common.TxReceipt] {
			btcReceipt := receipt.(BTCTxReceipt)
			isFinalized, err := client.isFinalized(btcReceipt.Raw, confHeight)
			if err != nil {
				return results.FromErr[common.TxReceipt](sdkerrors.Wrapf(errors.With(err, "tx_id", btcReceipt.Raw.Txid),
					"cannot determine if the transaction %s is finalized", btcReceipt.Raw.Txid),
				)
			}

			if !isFinalized {
				return results.FromErr[common.TxReceipt](common.ErrNotFinalized)
			}

			if btcReceipt.Raw.Confirmations < confHeight {
				clog.Redf("[BTC] tx_id: %s, conf_height: %d, confirmations: %d", btcReceipt.Raw.Txid, confHeight, btcReceipt.Raw.Confirmations)
				return results.FromErr[common.TxReceipt](common.ErrTxFailed)
			}

			return results.FromOk(receipt)
		})
	}), nil
}

func (c *BtcClient) GetTransactions(txIDs []common.Hash) ([]BTCTxResult, error) {
	txs := make([]BTCTxResult, len(txIDs))
	var wg sync.WaitGroup

	for i := range txIDs {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			tx, err := c.GetTransaction(txIDs[index])
			if err != nil {
				txs[index] = BTCTxResult(results.FromErr[common.TxReceipt](err))
			} else {
				txs[index] = tx
			}
		}(i)
	}

	wg.Wait()
	if slices.Any(txs, func(tx BTCTxResult) bool { return tx.IsErr() }) {
		return nil, common.ErrFailedToGetTransactions
	}

	return txs, nil
}

func (c *BtcClient) GetTransaction(txID common.Hash) (BTCTxResult, error) {
	var tx BTCTxReceipt

	// convert to string first to avoid the issue of reversed txid
	chainHash := common.HashToChainHash(txID)
	txResult, err := c.client.GetRawTransactionVerbose(&chainHash)
	if err != nil {
		clog.Cyanf("Failed to get BTC transaction %s: %+v", txID, err)
		return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
	} else {
		txRaw, err := hex.DecodeString(txResult.Hex)
		if err != nil {
			c.logger("failed to decode hex string", "txID", txID, "error", err)
			return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
		}

		msgTx := wire.NewMsgTx(wire.TxVersion)
		err = msgTx.Deserialize(bytes.NewReader(txRaw))
		if err != nil {
			c.logger("failed to parse transaction", "txID", txID, "error", err)
			return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
		}

		prevTxOuts, err := c.GetTxOuts(slices.Map(msgTx.TxIn, func(txIn *wire.TxIn) wire.OutPoint {
			return txIn.PreviousOutPoint
		}))

		if err != nil {
			c.logger("failed to get BTC transaction", "txID", txID, "error", err)
			return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
		}

		blockHash, err := chainhash.NewHashFromStr(txResult.BlockHash)
		if err != nil {
			c.logger("failed to create block hash", "blockHash", txResult.BlockHash, "error", err)
			return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
		}

		block, err := c.client.GetBlockVerbose(blockHash)
		if err != nil {
			c.logger("failed to get block", "blockHash", txResult.BlockHash, "error", err)
			return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
		}

		c.blockCache.Set(txResult.BlockHash, block)

		blockIndex := -1
		for i, txid := range block.Tx {
			if txid == txResult.Txid {
				blockIndex = i
				break
			}
		}

		clog.Redf("[BTC] blockIndex: %d", blockIndex)

		if blockIndex == -1 {
			return BTCTxResult(results.FromErr[common.TxReceipt](fmt.Errorf("transaction not found in block"))), fmt.Errorf("transaction not found in block")
		}

		tx.Raw = txResult
		tx.PrevTxOuts = prevTxOuts
		tx.MsgTx = msgTx
		tx.TransactionIndex = blockIndex
	}

	return results.FromOk[common.TxReceipt](tx), nil
}

func (c *BtcClient) GetTxOuts(outpoints []wire.OutPoint) ([]*btcjson.Vout, error) {
	txOuts := make([]*btcjson.Vout, len(outpoints))
	errChan := make(chan error, len(outpoints))
	var wg sync.WaitGroup

	for i := range outpoints {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			txOut, err := c.GetTxOut(outpoints[index])
			if err != nil || txOut == nil {
				errChan <- err
				return
			}
			txOuts[index] = txOut
		}(i)
	}

	wg.Wait()
	close(errChan)

	if err := <-errChan; err != nil {
		return nil, err
	}
	return txOuts, nil
}

func (c *BtcClient) GetTxOut(outpoint wire.OutPoint) (*btcjson.Vout, error) {
	txResult, err := c.client.GetRawTransactionVerbose(&outpoint.Hash)
	if err != nil {
		return nil, err
	}
	txOut := txResult.Vout[outpoint.Index]
	return &txOut, nil
}

func (c *BtcClient) LatestFinalizedBlockHeight(_ uint64) (uint64, error) {
	info, err := c.getBlockChainInfo()
	if err != nil {
		return 0, err
	}

	return uint64(info.Blocks), nil
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
