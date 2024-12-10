package btc

import (
	"sync"
)

//go:generate moq -out ./mock/block_height_cache.go -pkg mock . BlockHeightCache

// BlockHeightCache is a cache for the latest finalized block number for each chain
type BlockHeightCache interface {
	// Get returns the latest finalized block number for chain
	Get(blockHash string) *uint64
	// Set sets the latest finalized block number for chain, if the given block number is greater than the current latest finalized block number
	Set(blockHash string, blockHeight uint64)
}

type blockHeightCache struct {
	cache map[string]uint64
	lock  sync.RWMutex
}

func NewBlockHeightCache() BlockHeightCache {
	return &blockHeightCache{
		cache: make(map[string]uint64),
		lock:  sync.RWMutex{},
	}
}

// Get returns the latest finalized block number for chain
func (c *blockHeightCache) Get(blockHash string) *uint64 {
	c.lock.RLock()
	defer c.lock.RUnlock()

	height, ok := c.cache[blockHash]
	if !ok {
		return nil
	}

	return &height
}

// Set sets the latest finalized block number for chain, if the given block number is greater than the current latest finalized block number
func (c *blockHeightCache) Set(blockHash string, blockHeight uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()

	_, ok := c.cache[blockHash]
	if !ok {
		c.cache[blockHash] = blockHeight
	}
}
