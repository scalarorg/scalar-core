package btc

import (
	"strings"
	"sync"

	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
)

//go:generate moq -out ./mock/latest_finalized_block_cache.go -pkg mock . LatestFinalizedBlockCache

// LatestFinalizedBlockCache is a cache for the latest finalized block number for each chain
type LatestFinalizedBlockCache interface {
	// Get returns the latest finalized block number for chain
	Get(chain nexus.ChainName) *uint64
	// Set sets the latest finalized block number for chain, if the given block number is greater than the current latest finalized block number
	Set(chain nexus.ChainName, blockNumber uint64)
}

type latestFinalizedBlockCache struct {
	cache map[string]uint64
	lock  sync.RWMutex
}

func NewLatestFinalizedBlockCache() LatestFinalizedBlockCache {
	return &latestFinalizedBlockCache{
		cache: make(map[string]uint64),
		lock:  sync.RWMutex{},
	}
}

// Get returns the latest finalized block number for chain
func (c *latestFinalizedBlockCache) Get(chain nexus.ChainName) *uint64 {
	c.lock.RLock()
	defer c.lock.RUnlock()

	cachedBlockNumber, ok := c.cache[strings.ToLower(chain.String())]
	if !ok {
		return nil
	}

	return &cachedBlockNumber
}

// Set sets the latest finalized block number for chain, if the given block number is greater than the current latest finalized block number
func (c *latestFinalizedBlockCache) Set(chain nexus.ChainName, blockHeight uint64) {
	c.lock.Lock()
	defer c.lock.Unlock()

	chainName := strings.ToLower(chain.String())

	cachedBlockNumber, ok := c.cache[chainName]
	if !ok || blockHeight > cachedBlockNumber {
		c.cache[chainName] = blockHeight
	}
}
