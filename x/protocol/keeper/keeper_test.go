package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/scalarorg/scalar-core/app"
	"github.com/scalarorg/scalar-core/testutils/fake"
	"github.com/scalarorg/scalar-core/x/protocol/keeper"
	types "github.com/scalarorg/scalar-core/x/protocol/types"
)

func TestKeeper_GetRole_nil_Address_Return_Unrestricted(t *testing.T) {
	encCfg := app.MakeEncodingConfig()
	key := sdk.NewKVStoreKey("protocol")
	subspace := paramstypes.NewSubspace(encCfg.Codec, encCfg.Amino, key, sdk.NewKVStoreKey("trewardKey"), "reward")
	k := keeper.NewKeeper(encCfg.Codec, key, subspace)

	ctx := sdk.NewContext(fake.NewMultiStore(), sdk.Context{}.BlockHeader(), false, log.TestingLogger())
	protocols, ok := k.GetProtocols(ctx)
	assert.True(t, ok)
	assert.Equal(t, protocols, []*types.Protocol{})
}
