package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/types"
)

// Bech32 prefixes
var (
	AccountAddressPrefix   = "scalar"
	AccountPubKeyPrefix    = AccountAddressPrefix + sdk.PrefixPublic
	ValidatorAddressPrefix = AccountAddressPrefix + sdk.PrefixValidator + sdk.PrefixOperator
	ValidatorPubKeyPrefix  = AccountAddressPrefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	ConsNodeAddressPrefix  = AccountAddressPrefix + sdk.PrefixValidator + sdk.PrefixConsensus
	ConsNodePubKeyPrefix   = AccountAddressPrefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

// SetConfig sets the prefix config for the bech32 encoding
func SetConfig() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	config.Seal()
}

// RegisterDenoms registers the base and display denominations to the SDK.
func RegisterDenoms() {
	if err := sdk.RegisterDenom(types.DisplayDenom, sdk.OneDec()); err != nil {
		panic(err)
	}

	if err := sdk.RegisterDenom(types.BaseDenom, sdk.NewDecWithPrec(1, types.BaseDenomUnit)); err != nil {
		panic(err)
	}
}
