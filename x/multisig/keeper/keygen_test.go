package keeper_test

// import (
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/stretchr/testify/assert"

// 	"github.com/scalarorg/scalar-core/testutils/rand"
// 	"github.com/scalarorg/scalar-core/utils/testutils"
// 	"github.com/scalarorg/scalar-core/app/params"
// 	"github.com/scalarorg/scalar-core/x/multisig/keeper"
// 	"github.com/scalarorg/scalar-core/x/multisig/types"
// )

// func TestKeygenOptOut(t *testing.T) {
// 	encCfg := params.MakeEncodingConfig()
// 	k := keeper.NewKeeper(encCfg.Codec, sdk.NewKVStoreKey(types.StoreKey), testutils.NewSubspace(encCfg))

// 	participant := rand.AccAddr()

// 	ctx := testutils.NewContext()
// 	assert.False(t, k.HasOptedOut(ctx, participant))
// 	k.KeygenOptOut(ctx, participant)
// 	assert.True(t, k.HasOptedOut(ctx, participant))
// 	k.KeygenOptIn(ctx, participant)
// 	assert.False(t, k.HasOptedOut(ctx, participant))
// 	k.KeygenOptOut(ctx, participant)
// 	assert.True(t, k.HasOptedOut(ctx, participant))
// }
