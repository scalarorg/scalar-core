package types

import (
	fmt "fmt"

	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/scalarorg/scalar-core/utils"
)

var (
	KeySigningThreshold   = []byte("SigningThreshold")
	KeySigningTimeout     = []byte("SigningTimeout")
	KeySigningGracePeriod = []byte("SigningGracePeriod")
	KeyActiveEpochCount   = []byte("ActiveEpochCount")
)

// KeyTable retrieves a subspace table for the module
func KeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams creates the default genesis parameters
func DefaultParams() *Params {
	return &Params{
		SigningThreshold:   utils.NewThreshold(3, 5),
		SigningTimeout:     10,
		SigningGracePeriod: 1,
		ActiveEpochCount:   5,
	}
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of permission module's parameters.
func (m *Params) ParamSetPairs() params.ParamSetPairs {
	/*
		because the subspace package makes liberal use of pointers to set and get values from the store,
		this method needs to have a pointer receiver AND NewParamSetPair needs to receive the
		parameter values as pointer arguments, otherwise either the internal type reflection panics or the value will not be
		set on the correct Params data struct
	*/

	return params.ParamSetPairs{
		params.NewParamSetPair(KeySigningThreshold, &m.SigningThreshold, validateSigningThreshold),
		params.NewParamSetPair(KeySigningTimeout, &m.SigningTimeout, validateSigningTimeout),
		params.NewParamSetPair(KeySigningGracePeriod, &m.SigningGracePeriod, validateSigningGracePeriod),
		params.NewParamSetPair(KeyActiveEpochCount, &m.ActiveEpochCount, validateActiveEpochCount),
	}
}

// Validate checks the validity of the values of the parameter set
func (m Params) Validate() error {
	if err := validateSigningTimeout(m.SigningTimeout); err != nil {
		return err
	}

	if err := validateSigningGracePeriod(m.SigningGracePeriod); err != nil {
		return err
	}

	if err := validateActiveEpochCount(m.ActiveEpochCount); err != nil {
		return err
	}

	return nil
}

func validateSigningThreshold(i interface{}) error {
	threshold, ok := i.(utils.Threshold)
	if !ok {
		return fmt.Errorf("invalid parameter type for threshold: %T", i)
	}

	if err := threshold.Validate(); err != nil {
		return err
	}

	return nil
}

func validateSigningTimeout(i interface{}) error {
	timeout, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type for timeout: %T", i)
	}

	if timeout <= 0 {
		return fmt.Errorf("timeout must be >0")
	}

	return nil
}

func validateSigningGracePeriod(i interface{}) error {
	gracePeriod, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type for grace period: %T", i)
	}

	if gracePeriod < 0 {
		return fmt.Errorf("grace period must be >=0")
	}

	return nil
}

func validateActiveEpochCount(i interface{}) error {
	activeEpochCount, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type for active epoch count: %T", i)
	}

	if activeEpochCount <= 0 {
		return fmt.Errorf("active epoch count must be >0")
	}

	return nil
}
