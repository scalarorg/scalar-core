package common

import (
	"sync"
)

//go:generate moq -out ./mock/latest_finalized_block_cache.go -pkg mock . LatestFinalizedBlockCache

// LatestFinalizedBlockCache is a cache for the latest finalized block number for each chain
type LatestFinalizedBlockCache interface {
	// Get returns the latest finalized block number for chain
	Get() uint64
	// Set sets the latest finalized block number for chain, if the given block number is greater than the current latest finalized block number
	Set(blockNumber uint64)
}

type latestFinalizedBlockCache struct {
	cache uint64
	lock  sync.RWMutex
}

func NewLatestFinalizedBlockCache() LatestFinalizedBlockCache {
	return &latestFinalizedBlockCache{
		cache: 0,
		lock:  sync.RWMutex{},
	}
}

// Get returns the latest finalized block number for chain
func (c *latestFinalizedBlockCache) Get() uint64 {
	c.lock.RLock()
	defer c.lock.RUnlock()

	cachedBlockNumber := c.cache
	if cachedBlockNumber == 0 {
		return 0
	}

	return cachedBlockNumber
}

// Set sets the latest finalized block number for chain, if the given block number is greater than the current latest finalized block number
func (c *latestFinalizedBlockCache) Set(blockHeight uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if blockHeight > c.cache {
		c.cache = blockHeight
	}
}
