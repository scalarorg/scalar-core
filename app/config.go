package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	scalarnet "github.com/scalarorg/scalar-core/x/scalarnet/exported"
)

const (
	// DefaultGasPrice is default gas price for scalar transactions
	DefaultGasPrice = 20
)

// RegisterDenoms registers the base and display denominations to the SDK.
func RegisterDenoms() {
	if err := sdk.RegisterDenom(scalarnet.NativeAsset, sdk.ZeroDec()); err != nil {
		panic(err)
	}

	if err := sdk.RegisterDenom(scalarnet.BaseAsset, sdk.NewDecWithPrec(1, sdk.Precision)); err != nil {
		panic(err)
	}
}
