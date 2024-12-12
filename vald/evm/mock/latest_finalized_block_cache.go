// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/scalarorg/scalar-core/vald/evm"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"math/big"
	"sync"
)

// Ensure, that LatestFinalizedBlockCacheMock does implement evm.LatestFinalizedBlockCache.
// If this is not the case, regenerate this file with moq.
var _ evm.LatestFinalizedBlockCache = &LatestFinalizedBlockCacheMock{}

// LatestFinalizedBlockCacheMock is a mock implementation of evm.LatestFinalizedBlockCache.
//
//	func TestSomethingThatUsesLatestFinalizedBlockCache(t *testing.T) {
//
//		// make and configure a mocked evm.LatestFinalizedBlockCache
//		mockedLatestFinalizedBlockCache := &LatestFinalizedBlockCacheMock{
//			GetFunc: func(chain nexus.ChainName) *big.Int {
//				panic("mock out the Get method")
//			},
//			SetFunc: func(chain nexus.ChainName, blockNumber *big.Int)  {
//				panic("mock out the Set method")
//			},
//		}
//
//		// use mockedLatestFinalizedBlockCache in code that requires evm.LatestFinalizedBlockCache
//		// and then make assertions.
//
//	}
type LatestFinalizedBlockCacheMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(chain nexus.ChainName) *big.Int

	// SetFunc mocks the Set method.
	SetFunc func(chain nexus.ChainName, blockNumber *big.Int)

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// Chain is the chain argument value.
			Chain nexus.ChainName
		}
		// Set holds details about calls to the Set method.
		Set []struct {
			// Chain is the chain argument value.
			Chain nexus.ChainName
			// BlockNumber is the blockNumber argument value.
			BlockNumber *big.Int
		}
	}
	lockGet sync.RWMutex
	lockSet sync.RWMutex
}

// Get calls GetFunc.
func (mock *LatestFinalizedBlockCacheMock) Get(chain nexus.ChainName) *big.Int {
	if mock.GetFunc == nil {
		panic("LatestFinalizedBlockCacheMock.GetFunc: method is nil but LatestFinalizedBlockCache.Get was just called")
	}
	callInfo := struct {
		Chain nexus.ChainName
	}{
		Chain: chain,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(chain)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedLatestFinalizedBlockCache.GetCalls())
func (mock *LatestFinalizedBlockCacheMock) GetCalls() []struct {
	Chain nexus.ChainName
} {
	var calls []struct {
		Chain nexus.ChainName
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// Set calls SetFunc.
func (mock *LatestFinalizedBlockCacheMock) Set(chain nexus.ChainName, blockNumber *big.Int) {
	if mock.SetFunc == nil {
		panic("LatestFinalizedBlockCacheMock.SetFunc: method is nil but LatestFinalizedBlockCache.Set was just called")
	}
	callInfo := struct {
		Chain       nexus.ChainName
		BlockNumber *big.Int
	}{
		Chain:       chain,
		BlockNumber: blockNumber,
	}
	mock.lockSet.Lock()
	mock.calls.Set = append(mock.calls.Set, callInfo)
	mock.lockSet.Unlock()
	mock.SetFunc(chain, blockNumber)
}

// SetCalls gets all the calls that were made to Set.
// Check the length with:
//
//	len(mockedLatestFinalizedBlockCache.SetCalls())
func (mock *LatestFinalizedBlockCacheMock) SetCalls() []struct {
	Chain       nexus.ChainName
	BlockNumber *big.Int
} {
	var calls []struct {
		Chain       nexus.ChainName
		BlockNumber *big.Int
	}
	mock.lockSet.RLock()
	calls = mock.calls.Set
	mock.lockSet.RUnlock()
	return calls
}