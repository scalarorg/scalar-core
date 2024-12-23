package testnet

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// BaseDenomUnit defines the base denomination unit for Evmos.
	// 1 evmos = 1x10^{BaseDenomUnit} aevmos
	BaseDenomUnit = 18
	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20

	PurposeValidator     = 0
	PurposeBroadcaster   = 1
	PurposeGovernance    = 2
	PurposeFaucetAccount = 3
	CoinTypeMainnet      = 0
	CoinTypeTestnet      = 1
)

var (
	ValidatorTokens   = sdk.NewInt(1e9)
	ValidatorStaking  = sdk.NewInt(1e6)
	BroadcasterTokens = sdk.NewInt(1e6)
	ScalarTokens      = sdk.NewInt(1e9)
	GovTokens         = sdk.NewInt(1e9)
	FaucetTokens      = sdk.NewInt(1e9)
	PowerReduction    = sdk.NewInt(1e6)
)
