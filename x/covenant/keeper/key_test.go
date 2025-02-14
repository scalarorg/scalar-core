package keeper_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gotest.tools/assert"
)

const (
	AccountAddressPrefix = "scalar"
	BaseAsset            = "ascal"
)

var (
	AccountPubKeyPrefix    = AccountAddressPrefix + types.PrefixPublic
	ValidatorAddressPrefix = AccountAddressPrefix + types.PrefixValidator + types.PrefixOperator
	ValidatorPubKeyPrefix  = AccountAddressPrefix + types.PrefixValidator + types.PrefixOperator + types.PrefixPublic
	ConsNodeAddressPrefix  = AccountAddressPrefix + types.PrefixValidator + types.PrefixConsensus
	ConsNodePubKeyPrefix   = AccountAddressPrefix + types.PrefixValidator + types.PrefixConsensus + types.PrefixPublic
)

func setCosmosAccountPrefix() {
	config := types.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
}

func TestDecodeBench32Address(t *testing.T) {
	setCosmosAccountPrefix()
	addr, err := sdk.ValAddressFromBech32("scalarvaloper1u8ennvwsshneu4nvec38e3jvcxmppf7lfq3pf5")
	assert.NilError(t, err)

	fmt.Println(addr.String())
}
