package btc

import (
	"sync"
	"time"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/rs/zerolog/log"
)

//go:generate moq -out ./mock/block_cache.go -pkg mock . BlockCache
type TxResultWithIndex struct {
	*btcjson.TxRawResult
	Index int
}

const (
	ChainInfoCacheDuration = 1 * time.Minute
	BlockTime              = 12 * time.Minute
)

type BlockCacheStatus int

const (
	BlockCacheStatusUndefined BlockCacheStatus = iota
	BlockCacheStatusFetching
	BlockCacheStatusStored
)

type ChainInfoWithTimestamp struct {
	*btcjson.GetBlockChainInfoResult
	Timestamp time.Time //Request time
}

type BlockCache interface {
	GetBlock(blockHash string) (*btcjson.GetBlockVerboseTxResult, BlockCacheStatus)
	SetBlock(blockHash string, block *btcjson.GetBlockVerboseTxResult)
	GetTx(txHash string) *TxResultWithIndex
	SetTx(txHash string, tx *TxResultWithIndex)
	GetChainInfo() *btcjson.GetBlockChainInfoResult
	SetChainInfo(chainInfo *btcjson.GetBlockChainInfoResult)
	GetBlockStatus(blockHash string) BlockCacheStatus
	SetBlockStatus(blockHash string, status BlockCacheStatus)
}

// TODO: Add rocksdb cache
type blockCache struct {
	// the key is the block hash in reverse order bytes aka rpc
	blocks        map[string]*btcjson.GetBlockVerboseTxResult
	blockLock     sync.RWMutex
	txs           map[string]*TxResultWithIndex
	txLock        sync.RWMutex
	chainInfo     *ChainInfoWithTimestamp //Latest finalized block height
	chainInfoLock sync.RWMutex
	blockStatus   sync.Map
}

func NewBlockCache() *blockCache {
	return &blockCache{
		blocks:    make(map[string]*btcjson.GetBlockVerboseTxResult),
		txs:       make(map[string]*TxResultWithIndex),
		blockLock: sync.RWMutex{},
		txLock:    sync.RWMutex{},
	}
}

// Get returns the block for the given block hash
func (c *blockCache) GetBlock(blockHash string) (*btcjson.GetBlockVerboseTxResult, BlockCacheStatus) {
	c.blockLock.RLock()
	defer c.blockLock.RUnlock()

	block, ok := c.blocks[blockHash]
	if !ok {
		return nil, c.GetBlockStatus(blockHash)
	}
	return block, BlockCacheStatusStored
}

// Set sets the block for the given block hash
func (c *blockCache) SetBlock(blockHash string, block *btcjson.GetBlockVerboseTxResult) {
	c.blockLock.Lock()
	defer c.blockLock.Unlock()
	log.Info().Str("blockHash", blockHash).Msg("SetBlock to cache")
	_, ok := c.blocks[blockHash]
	if !ok {
		c.blocks[blockHash] = block
	}
	c.SetBlockStatus(blockHash, BlockCacheStatusStored)
	c.txLock.Lock()
	defer c.txLock.Unlock()
	//Loop through the txs and set them
	for index, tx := range block.Tx {
		txHash, err := chainhash.NewHashFromStr(tx.Txid)
		if err != nil {
			log.Error().Str("txid", tx.Txid).Msg("failed to convert txid to chainhash")
			continue
		}
		c.txs[txHash.String()] = &TxResultWithIndex{
			TxRawResult: &tx,
			Index:       index,
		}
	}
}

func (c *blockCache) GetTx(txHash string) *TxResultWithIndex {
	c.txLock.RLock()
	defer c.txLock.RUnlock()

	tx, ok := c.txs[txHash]
	if !ok {
		return nil
	}

	return tx
}

func (c *blockCache) SetTx(txHash string, tx *TxResultWithIndex) {
	c.txLock.Lock()
	defer c.txLock.Unlock()

	c.txs[txHash] = tx
}

func (c *blockCache) GetChainInfo() *btcjson.GetBlockChainInfoResult {
	c.chainInfoLock.RLock()
	defer c.chainInfoLock.RUnlock()
	if c.chainInfo == nil {
		return nil
	}
	//Todo: This condition is not best option for Checking if the chain info is outdated
	if time.Since(c.chainInfo.Timestamp) > ChainInfoCacheDuration {
		return nil
	}
	return c.chainInfo.GetBlockChainInfoResult
}

func (c *blockCache) SetChainInfo(chainInfo *btcjson.GetBlockChainInfoResult) {
	c.chainInfoLock.Lock()
	defer c.chainInfoLock.Unlock()

	c.chainInfo = &ChainInfoWithTimestamp{
		GetBlockChainInfoResult: chainInfo,
		Timestamp:               time.Now(),
	}
}

func (c *blockCache) GetBlockStatus(blockHash string) BlockCacheStatus {
	status, ok := c.blockStatus.Load(blockHash)
	if !ok {
		return BlockCacheStatusUndefined
	}
	return status.(BlockCacheStatus)
}

func (c *blockCache) SetBlockStatus(blockHash string, status BlockCacheStatus) {
	c.blockStatus.Store(blockHash, status)
}
