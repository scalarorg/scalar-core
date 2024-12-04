package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// DisplayDenom defines the denomination displayed to users in client applications.
	DisplayDenom = "scal"
	// BaseDenom defines to the default denomination used in Scalar (staking, EVM, governance, etc.)
	BaseDenom = "ascal"
	// BaseDenomUnit defines the base denomination unit for Evmos.
	// 1 evmos = 1x10^{BaseDenomUnit} aevmos
	BaseDenomUnit = 18
	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// var PowerReduction = sdk.NewInt(1e0 * params.InitialBaseFee)
var PowerReduction = sdk.NewInt(1e6)

var NodeTokens = sdk.NewInt(1e15)
var ValidatorTokens = sdk.NewInt(1e12)
var BroadcasterTokens = sdk.NewInt(1e9)
var DelegatorTokens = sdk.NewInt(1e6)
