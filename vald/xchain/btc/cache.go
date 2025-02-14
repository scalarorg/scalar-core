package btc

import (
	"sync"

	"github.com/btcsuite/btcd/btcjson"
)

//go:generate moq -out ./mock/block_cache.go -pkg mock . BlockCache

type BlockCache struct {
	// the key is the block hash in reverse order bytes aka rpc
	cache map[string]*btcjson.GetBlockVerboseResult
	lock  sync.RWMutex
}

func NewBlockCache() *BlockCache {
	return &BlockCache{
		cache: make(map[string]*btcjson.GetBlockVerboseResult),
		lock:  sync.RWMutex{},
	}
}

// Get returns the block for the given block hash
func (c *BlockCache) Get(blockHash string) *btcjson.GetBlockVerboseResult {
	c.lock.RLock()
	defer c.lock.RUnlock()

	block, ok := c.cache[blockHash]
	if !ok {
		return nil
	}

	return block
}

// Set sets the block for the given block hash
func (c *BlockCache) Set(blockHash string, block *btcjson.GetBlockVerboseResult) {
	c.lock.Lock()
	defer c.lock.Unlock()

	_, ok := c.cache[blockHash]
	if !ok {
		c.cache[blockHash] = block
	}
}
