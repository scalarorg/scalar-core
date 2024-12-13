// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/snapshot/exported"
	"sync"
)

// Ensure, that ValidatorIMock does implement exported.ValidatorI.
// If this is not the case, regenerate this file with moq.
var _ exported.ValidatorI = &ValidatorIMock{}

// ValidatorIMock is a mock implementation of exported.ValidatorI.
//
//	func TestSomethingThatUsesValidatorI(t *testing.T) {
//
//		// make and configure a mocked exported.ValidatorI
//		mockedValidatorI := &ValidatorIMock{
//			GetConsAddrFunc: func() (github_com_cosmos_cosmos_sdk_types.ConsAddress, error) {
//				panic("mock out the GetConsAddr method")
//			},
//			GetConsensusPowerFunc: func(intMoqParam github_com_cosmos_cosmos_sdk_types.Int) int64 {
//				panic("mock out the GetConsensusPower method")
//			},
//			GetOperatorFunc: func() github_com_cosmos_cosmos_sdk_types.ValAddress {
//				panic("mock out the GetOperator method")
//			},
//			IsBondedFunc: func() bool {
//				panic("mock out the IsBonded method")
//			},
//			IsJailedFunc: func() bool {
//				panic("mock out the IsJailed method")
//			},
//		}
//
//		// use mockedValidatorI in code that requires exported.ValidatorI
//		// and then make assertions.
//
//	}
type ValidatorIMock struct {
	// GetConsAddrFunc mocks the GetConsAddr method.
	GetConsAddrFunc func() (github_com_cosmos_cosmos_sdk_types.ConsAddress, error)

	// GetConsensusPowerFunc mocks the GetConsensusPower method.
	GetConsensusPowerFunc func(intMoqParam github_com_cosmos_cosmos_sdk_types.Int) int64

	// GetOperatorFunc mocks the GetOperator method.
	GetOperatorFunc func() github_com_cosmos_cosmos_sdk_types.ValAddress

	// IsBondedFunc mocks the IsBonded method.
	IsBondedFunc func() bool

	// IsJailedFunc mocks the IsJailed method.
	IsJailedFunc func() bool

	// calls tracks calls to the methods.
	calls struct {
		// GetConsAddr holds details about calls to the GetConsAddr method.
		GetConsAddr []struct {
		}
		// GetConsensusPower holds details about calls to the GetConsensusPower method.
		GetConsensusPower []struct {
			// IntMoqParam is the intMoqParam argument value.
			IntMoqParam github_com_cosmos_cosmos_sdk_types.Int
		}
		// GetOperator holds details about calls to the GetOperator method.
		GetOperator []struct {
		}
		// IsBonded holds details about calls to the IsBonded method.
		IsBonded []struct {
		}
		// IsJailed holds details about calls to the IsJailed method.
		IsJailed []struct {
		}
	}
	lockGetConsAddr       sync.RWMutex
	lockGetConsensusPower sync.RWMutex
	lockGetOperator       sync.RWMutex
	lockIsBonded          sync.RWMutex
	lockIsJailed          sync.RWMutex
}

// GetConsAddr calls GetConsAddrFunc.
func (mock *ValidatorIMock) GetConsAddr() (github_com_cosmos_cosmos_sdk_types.ConsAddress, error) {
	if mock.GetConsAddrFunc == nil {
		panic("ValidatorIMock.GetConsAddrFunc: method is nil but ValidatorI.GetConsAddr was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetConsAddr.Lock()
	mock.calls.GetConsAddr = append(mock.calls.GetConsAddr, callInfo)
	mock.lockGetConsAddr.Unlock()
	return mock.GetConsAddrFunc()
}

// GetConsAddrCalls gets all the calls that were made to GetConsAddr.
// Check the length with:
//
//	len(mockedValidatorI.GetConsAddrCalls())
func (mock *ValidatorIMock) GetConsAddrCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetConsAddr.RLock()
	calls = mock.calls.GetConsAddr
	mock.lockGetConsAddr.RUnlock()
	return calls
}

// GetConsensusPower calls GetConsensusPowerFunc.
func (mock *ValidatorIMock) GetConsensusPower(intMoqParam github_com_cosmos_cosmos_sdk_types.Int) int64 {
	if mock.GetConsensusPowerFunc == nil {
		panic("ValidatorIMock.GetConsensusPowerFunc: method is nil but ValidatorI.GetConsensusPower was just called")
	}
	callInfo := struct {
		IntMoqParam github_com_cosmos_cosmos_sdk_types.Int
	}{
		IntMoqParam: intMoqParam,
	}
	mock.lockGetConsensusPower.Lock()
	mock.calls.GetConsensusPower = append(mock.calls.GetConsensusPower, callInfo)
	mock.lockGetConsensusPower.Unlock()
	return mock.GetConsensusPowerFunc(intMoqParam)
}

// GetConsensusPowerCalls gets all the calls that were made to GetConsensusPower.
// Check the length with:
//
//	len(mockedValidatorI.GetConsensusPowerCalls())
func (mock *ValidatorIMock) GetConsensusPowerCalls() []struct {
	IntMoqParam github_com_cosmos_cosmos_sdk_types.Int
} {
	var calls []struct {
		IntMoqParam github_com_cosmos_cosmos_sdk_types.Int
	}
	mock.lockGetConsensusPower.RLock()
	calls = mock.calls.GetConsensusPower
	mock.lockGetConsensusPower.RUnlock()
	return calls
}

// GetOperator calls GetOperatorFunc.
func (mock *ValidatorIMock) GetOperator() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if mock.GetOperatorFunc == nil {
		panic("ValidatorIMock.GetOperatorFunc: method is nil but ValidatorI.GetOperator was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetOperator.Lock()
	mock.calls.GetOperator = append(mock.calls.GetOperator, callInfo)
	mock.lockGetOperator.Unlock()
	return mock.GetOperatorFunc()
}

// GetOperatorCalls gets all the calls that were made to GetOperator.
// Check the length with:
//
//	len(mockedValidatorI.GetOperatorCalls())
func (mock *ValidatorIMock) GetOperatorCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetOperator.RLock()
	calls = mock.calls.GetOperator
	mock.lockGetOperator.RUnlock()
	return calls
}

// IsBonded calls IsBondedFunc.
func (mock *ValidatorIMock) IsBonded() bool {
	if mock.IsBondedFunc == nil {
		panic("ValidatorIMock.IsBondedFunc: method is nil but ValidatorI.IsBonded was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsBonded.Lock()
	mock.calls.IsBonded = append(mock.calls.IsBonded, callInfo)
	mock.lockIsBonded.Unlock()
	return mock.IsBondedFunc()
}

// IsBondedCalls gets all the calls that were made to IsBonded.
// Check the length with:
//
//	len(mockedValidatorI.IsBondedCalls())
func (mock *ValidatorIMock) IsBondedCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsBonded.RLock()
	calls = mock.calls.IsBonded
	mock.lockIsBonded.RUnlock()
	return calls
}

// IsJailed calls IsJailedFunc.
func (mock *ValidatorIMock) IsJailed() bool {
	if mock.IsJailedFunc == nil {
		panic("ValidatorIMock.IsJailedFunc: method is nil but ValidatorI.IsJailed was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsJailed.Lock()
	mock.calls.IsJailed = append(mock.calls.IsJailed, callInfo)
	mock.lockIsJailed.Unlock()
	return mock.IsJailedFunc()
}

// IsJailedCalls gets all the calls that were made to IsJailed.
// Check the length with:
//
//	len(mockedValidatorI.IsJailedCalls())
func (mock *ValidatorIMock) IsJailedCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsJailed.RLock()
	calls = mock.calls.IsJailed
	mock.lockIsJailed.RUnlock()
	return calls
}
