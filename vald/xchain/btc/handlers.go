package btc

import (
	"bytes"
	"encoding/hex"
	"sync"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/errors"
	"github.com/scalarorg/scalar-core/utils/monads/results"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/vald/xchain/common"
)

func (client *BtcClient) GetTxReceiptsIfFinalized(txIDs []common.Hash, confHeight uint64) ([]BTCTxResult, error) {
	start := time.Now()
	txResults, err := client.GetTransactions(txIDs)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			errors.With(err, "tx_ids", txIDs),
			"cannot get transaction receipts",
		)
	}
	elapsed := time.Since(start)
	log.Info().Str("elapsed", elapsed.String()).Int("tx_count", len(txIDs)).Msg("GetTransactions")
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
	for i, txID := range txIDs {
		wg.Add(1)
		go func(index int, txID common.Hash) {
			defer wg.Done()
			tx, err := c.GetTransaction(txID)
			if err != nil {
				txs[index] = BTCTxResult(results.FromErr[common.TxReceipt](err))
			} else {
				txs[index] = tx
			}
		}(i, txID)
	}
	wg.Wait()

	if slices.Any(txs, func(tx BTCTxResult) bool { return tx.IsErr() }) {
		return nil, common.ErrFailedToGetTransactions
	}
	//Set txIndex for each tx
	for index, rawTx := range txs {
		if rawTx.IsErr() {
			continue
		}
		txReceipt := rawTx.Ok().(BTCTxReceipt)
		block, _ := c.blockCache.GetBlock(txReceipt.Raw.BlockHash)
		// block := mapBlocks[txReceipt.Raw.BlockHash]
		if block == nil {
			continue
		}
		wg.Add(1)
		go func(index int, rawReceipt *BTCTxReceipt, block *btcjson.GetBlockVerboseTxResult) {
			defer wg.Done()
			txResult, err := c.createBtcTxResult(rawReceipt, block)
			if err != nil {
				return
			}
			txs[index] = txResult
		}(index, &txReceipt, block)
	}
	wg.Wait()
	return txs, nil
}
func (c *BtcClient) GetRawTx(txID common.Hash) (results.Result[BTCTxReceipt], error) {
	// convert to string first to avoid the issue of reversed txid
	var tx BTCTxReceipt

	// convert to string first to avoid the issue of reversed txid
	chainHash := common.HashToChainHash(txID)
	//Try get from cache first
	txWithIndex := c.blockCache.GetTx(chainHash.String())
	if txWithIndex != nil {
		log.Info().Any("txWithIndex", txWithIndex).Str("txID", txID.Hex()).Msg("GetTransaction from cache")
		tx.Raw = txWithIndex.TxRawResult
		tx.TransactionIndex = txWithIndex.Index
	} else {
		//Get from rpc
		txResult, err := c.client.GetRawTransactionVerbose(&chainHash)
		if err != nil {
			clog.Cyanf("Failed to get BTC transaction %s: %+v", txID, err)
			return results.FromErr[BTCTxReceipt](err), err
		}
		tx.Raw = txResult
	}
	return results.FromOk(tx), nil
}

func (c *BtcClient) GetBlockVerboseTx(hashStr string) (results.Result[*btcjson.GetBlockVerboseTxResult], error) {
	block, _ := c.blockCache.GetBlock(hashStr)
	if block != nil {
		return results.FromOk(block), nil
	}
	blockHash, err := chainhash.NewHashFromStr(hashStr)
	if err != nil {
		return results.FromErr[*btcjson.GetBlockVerboseTxResult](err), err
	}
	block, err = c.client.GetBlockVerboseTx(blockHash)
	if err != nil {
		return results.FromErr[*btcjson.GetBlockVerboseTxResult](err), err
	}
	c.blockCache.SetBlock(hashStr, block)
	return results.FromOk(block), nil
}

func (c *BtcClient) createBtcTxResult(rawReceipt *BTCTxReceipt, block *btcjson.GetBlockVerboseTxResult) (BTCTxResult, error) {
	var txReceipt BTCTxReceipt
	txReceipt.Raw = rawReceipt.Raw
	for i, tx := range block.Tx {
		if tx.Txid == rawReceipt.Raw.Txid {
			txReceipt.TransactionIndex = i
			break
		}
	}
	//Try to set msgTx
	txRaw, err := hex.DecodeString(rawReceipt.Raw.Hex)
	if err != nil {
		c.logger("failed to decode hex string", "txID", rawReceipt.Raw.Txid, "error", err)
		return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
	}

	msgTx := wire.NewMsgTx(wire.TxVersion)
	err = msgTx.Deserialize(bytes.NewReader(txRaw))
	if err != nil {
		c.logger("failed to parse transaction", "txID", rawReceipt.Raw.Txid, "error", err)
		return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
	}
	txReceipt.MsgTx = msgTx
	//Try to set prevTxOuts
	txReceipt.BlockHeight = &block.Height
	txReceipt.PrevTxOuts, err = c.GetTxOuts(slices.Map(msgTx.TxIn, func(txIn *wire.TxIn) wire.OutPoint {
		return txIn.PreviousOutPoint
	}))
	if err != nil {
		c.logger("failed to get BTC transaction", "txID", rawReceipt.Raw.Txid, "error", err)
		return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
	}
	return results.FromOk[common.TxReceipt](txReceipt), nil
}
func (c *BtcClient) GetTransaction(txID common.Hash) (BTCTxResult, error) {
	var tx BTCTxReceipt
	var block *btcjson.GetBlockVerboseTxResult
	// convert to string first to avoid the issue of reversed txid
	chainHash := common.HashToChainHash(txID)
	//Try get from cache first
	txWithIndex := c.blockCache.GetTx(txID.Hex())
	if txWithIndex != nil {
		log.Info().Any("txWithIndex", txWithIndex).Msg("GetTransaction from cache")
		tx.Raw = txWithIndex.TxRawResult
		tx.TransactionIndex = txWithIndex.Index
	} else {
		//Get from rpc
		txResult, err := c.client.GetRawTransactionVerbose(&chainHash)
		if err != nil {
			clog.Cyanf("Failed to get BTC transaction %s: %+v", txID, err)
			return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
		}
		tx.Raw = txResult
		//Try to get block then write to cache
		var status BlockCacheStatus
		block, status = c.blockCache.GetBlock(txResult.BlockHash)
		if block == nil {
			switch status {
			case BlockCacheStatusUndefined:
				c.blockCache.SetBlockStatus(txResult.BlockHash, BlockCacheStatusFetching)
				blockHash, err := chainhash.NewHashFromStr(txResult.BlockHash)
				if err != nil {
					clog.Cyanf("Failed to get BTC block hash %s: %+v", txResult.BlockHash, err)
					return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
				}
				start := time.Now()
				block, err = c.client.GetBlockVerboseTx(blockHash)
				elapsed := time.Since(start)
				log.Info().Str("elapsed", elapsed.String()).Msg("GetBlockVerboseTx from rpc")
				if err != nil {
					c.logger("failed to get block", "blockHash", txResult.BlockHash, "error", err)
					return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
				}
				c.blockCache.SetBlock(txResult.BlockHash, block)
			case BlockCacheStatusFetching:
				//wait for the block to be fetched
				log.Info().Str("blockHash", txResult.BlockHash).Msg("Block is fetching in other go routine")
				for c.blockCache.GetBlockStatus(txResult.BlockHash) == BlockCacheStatusFetching {
					time.Sleep(100 * time.Millisecond)
				}
				block, status = c.blockCache.GetBlock(txResult.BlockHash)
				if block != nil && status == BlockCacheStatusStored {
					log.Info().Str("blockHash", txResult.BlockHash).Msg("Block found in cache")
				} else {
					log.Error().Str("blockHash", txResult.BlockHash).Msg("Block not found in cache")
				}
			case BlockCacheStatusStored:
				//log.Info().Str("blockHash", txResult.BlockHash).Msg("Block is already in cache")
			}
		}
	}
	// //Try to set msgTx
	// txRaw, err := hex.DecodeString(tx.Raw.Hex)
	// if err != nil {
	// 	c.logger("failed to decode hex string", "txID", txID, "error", err)
	// 	return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
	// }

	// msgTx := wire.NewMsgTx(wire.TxVersion)
	// err = msgTx.Deserialize(bytes.NewReader(txRaw))
	// if err != nil {
	// 	c.logger("failed to parse transaction", "txID", txID, "error", err)
	// 	return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
	// }
	// tx.MsgTx = msgTx
	// //Try to set prevTxOuts

	// tx.PrevTxOuts, err = c.GetTxOuts(slices.Map(msgTx.TxIn, func(txIn *wire.TxIn) wire.OutPoint {
	// 	return txIn.PreviousOutPoint
	// }))
	// if err != nil {
	// 	c.logger("failed to get BTC transaction", "txID", txID, "error", err)
	// 	return BTCTxResult(results.FromErr[common.TxReceipt](err)), err
	// }

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
	//Try get from cache first
	txWithIndex := c.blockCache.GetTx(outpoint.Hash.String())
	if txWithIndex != nil {
		txOut := txWithIndex.TxRawResult.Vout[outpoint.Index]
		return &txOut, nil
	}
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
