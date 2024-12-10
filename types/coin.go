package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	scalarnet "github.com/scalarorg/scalar-core/x/scalarnet/exported"
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
	// DisplayDenom defines the denomination displayed to users in client applications.
	DisplayDenom = scalarnet.NativeAsset
	// BaseDenom defines to the default denomination used in Scalar (staking, EVM, governance, etc.)
	BaseDenom = "a" + scalarnet.NativeAsset

	PowerReduction    = sdk.NewInt(1e6)
	ValidatorTokens   = sdk.NewInt(1e12)
	BroadcasterTokens = sdk.NewInt(1e15)
	GovTokens         = sdk.NewInt(1e15)
	FaucetTokens      = sdk.NewInt(1e18)
	DelegatorTokens   = sdk.NewInt(1e6)
)
