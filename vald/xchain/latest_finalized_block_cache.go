package xchain

import (
	"sync"

	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
)

//go:generate moq -out ./mock/latest_finalized_block_cache.go -pkg mock . LatestFinalizedBlockCache

// LatestFinalizedBlockCache is a cache for the latest finalized block number for each chain
type LatestFinalizedBlockCache interface {
	// Get returns the latest finalized block number for chain
	Get(chainInfo chain.ChainInfoBytes) *uint64
	// Set sets the latest finalized block number for chain, if the given block number is greater than the current latest finalized block number
	Set(chainInfo chain.ChainInfoBytes, blockNumber uint64)
}

type latestFinalizedBlockCache struct {
	cache map[chain.ChainInfoBytes]uint64
	lock  sync.RWMutex
}

func NewLatestFinalizedBlockCache() LatestFinalizedBlockCache {
	return &latestFinalizedBlockCache{
		cache: make(map[chain.ChainInfoBytes]uint64),
		lock:  sync.RWMutex{},
	}
}

// Get returns the latest finalized block number for chain
func (c *latestFinalizedBlockCache) Get(chainInfo chain.ChainInfoBytes) *uint64 {
	c.lock.RLock()
	defer c.lock.RUnlock()

	cachedBlockNumber, ok := c.cache[chainInfo]
	if !ok {
		return nil
	}

	return &cachedBlockNumber
}

// Set sets the latest finalized block number for chain, if the given block number is greater than the current latest finalized block number
func (c *latestFinalizedBlockCache) Set(chainInfo chain.ChainInfoBytes, blockHeight uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()

	cachedBlockNumber, ok := c.cache[chainInfo]
	if !ok || blockHeight > cachedBlockNumber {
		c.cache[chainInfo] = blockHeight
	}
}
