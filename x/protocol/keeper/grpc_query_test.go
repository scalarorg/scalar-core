package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/scalarorg/scalar-core/app"
	"github.com/scalarorg/scalar-core/testutils/fake"
	. "github.com/scalarorg/scalar-core/utils/test"
	"github.com/scalarorg/scalar-core/x/protocol/keeper"
	"github.com/scalarorg/scalar-core/x/protocol/types"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestGrpcQuery(t *testing.T) {
	cfg := app.MakeEncodingConfig()
	var (
		k              keeper.Keeper
		ctx            sdk.Context
		initialGenesis *types.GenesisState
		protocols      []*types.Protocol
	)

	Given("a keeper",
		func() {
			subspace := paramstypes.NewSubspace(cfg.Codec, cfg.Amino, sdk.NewKVStoreKey("paramsKey"), sdk.NewKVStoreKey("tparamsKey"), "protocol")
			k = keeper.NewKeeper(cfg.Codec, sdk.NewKVStoreKey(types.StoreKey), subspace)
		}).
		Given("a state that is initialized",
			func() {
				protocols := []*types.Protocol{
					types.DefaultProtocol(),
				}
				initialGenesis = types.NewGenesisState(protocols)
				assert.NoError(t, initialGenesis.Validate())

				ctx = sdk.NewContext(fake.NewMultiStore(), tmproto.Header{}, false, log.TestingLogger())
				k.InitGenesis(ctx, initialGenesis)
			}).
		When("querying the protocol",
			func() {
				req := &types.ProtocolsRequest{}
				resp, err := k.Protocols(sdk.WrapSDKContext(ctx), req)
				assert.NotNil(t, resp)
				assert.Nil(t, err)
				protocols = resp.Protocols
			}).
		Then("return the expected key",
			func(t *testing.T) {
				assert.Equal(t, initialGenesis.Protocols, protocols)
			}).Run(t, 10)

}
