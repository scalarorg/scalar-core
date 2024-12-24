package keeper_test

// import (
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/tendermint/tendermint/libs/log"
// 	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

// 	"github.com/scalarorg/scalar-core/testutils/fake"
// 	"github.com/scalarorg/scalar-core/testutils/rand"
// 	chains "github.com/scalarorg/scalar-core/x/chains/exported"
// 	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
// 	"github.com/scalarorg/scalar-core/x/scalarnet/exported"
// 	"github.com/scalarorg/scalar-core/x/scalarnet/keeper"
// 	"github.com/scalarorg/scalar-core/x/scalarnet/types"
// 	"github.com/scalarorg/scalar-core/x/scalarnet/types/mock"
// )

// func TestAddressValidator(t *testing.T) {
// 	ctx := sdk.NewContext(fake.NewMultiStore(), tmproto.Header{}, false, log.TestingLogger())
// 	prefixes := map[string]string{
// 		"Scalarnet": "scalar",
// 		"terra":     "terra",
// 		"osmosis":   "osmo",
// 	}
// 	scalarnetK := &mock.BaseKeeperMock{
// 		GetCosmosChainByNameFunc: func(ctx sdk.Context, chain nexus.ChainName) (types.CosmosChain, bool) {
// 			prefix, ok := prefixes[chain.String()]
// 			if !ok {
// 				return types.CosmosChain{}, false
// 			}
// 			return types.CosmosChain{Name: chain, AddrPrefix: prefix}, true
// 		},
// 	}

// 	sdk.GetConfig().SetBech32PrefixForAccount("", "")

// 	validator := keeper.NewAddressValidator(scalarnetK)
// 	assert.NotNil(t, validator)

// 	scalarAddr := "scalar1t66w8cazua870wu7t2hsffndmy2qy2v556ymndnczs83qpz2h45sq6lq9w"

// 	addr := nexus.CrossChainAddress{Chain: exported.Scalarnet, Address: scalarAddr}
// 	assert.NoError(t, validator(ctx, addr))

// 	addr = nexus.CrossChainAddress{Chain: nexus.Chain{Name: "terra", Module: types.ModuleName}, Address: "terra18zhnqjv70v0d2f8v0s5lape0gr5ua94eqkk8ex"}
// 	assert.NoError(t, validator(ctx, addr))

// 	addr = nexus.CrossChainAddress{Chain: nexus.Chain{Name: "osmosis", Module: types.ModuleName}, Address: "osmo18zhnqjv70v0d2f8v0s5lape0gr5ua94ewflhd5"}
// 	assert.NoError(t, validator(ctx, addr))

// 	longAddr := sdk.AccAddress(rand.Bytes(255))
// 	addr = nexus.CrossChainAddress{Chain: exported.Scalarnet, Address: longAddr.String()}
// 	assert.NoError(t, validator(ctx, addr))

// 	shortAddr := sdk.AccAddress(rand.Bytes(1))
// 	addr = nexus.CrossChainAddress{Chain: exported.Scalarnet, Address: shortAddr.String()}
// 	assert.NoError(t, validator(ctx, addr))

// 	addr = nexus.CrossChainAddress{Chain: exported.Scalarnet, Address: "0x68B93045fe7D8794a7cAF327e7f855CD6Cd03BB8"}
// 	assert.Error(t, validator(ctx, addr))

// 	addr = nexus.CrossChainAddress{Chain: exported.Scalarnet, Address: ""}
// 	assert.Error(t, validator(ctx, addr))

// 	addr = nexus.CrossChainAddress{Chain: exported.Scalarnet, Address: "scalar1t66w8cazua870wu7t2hsffndmy2qy2v556ymndnczs83qpz2h45sq6lq9v"}
// 	assert.ErrorContains(t, validator(ctx, addr), "invalid checksum")

// 	addr = nexus.CrossChainAddress{Chain: exported.Scalarnet, Address: "osmo18zhnqjv70v0d2f8v0s5lape0gr5ua94ewflhd5"}
// 	assert.ErrorContains(t, validator(ctx, addr), "invalid Bech32 prefix")

// 	shortAddr = rand.Bytes(0)
// 	addr = nexus.CrossChainAddress{Chain: exported.Scalarnet, Address: shortAddr.String()}
// 	assert.ErrorContains(t, validator(ctx, addr), "non empty address")

// 	longAddr = rand.Bytes(256)
// 	addr = nexus.CrossChainAddress{Chain: exported.Scalarnet, Address: longAddr.String()}
// 	assert.ErrorContains(t, validator(ctx, addr), "max length")

// 	addr = nexus.CrossChainAddress{Chain: nexus.Chain{Name: "secret", Module: types.ModuleName}, Address: "secret18zhnqjv70v0d2f8v0s5lape0gr5ua94eyhcwx6"}
// 	assert.ErrorContains(t, validator(ctx, addr), "no known prefix")

// 	addr = nexus.CrossChainAddress{Chain: evm.Ethereum, Address: "0x68B93045fe7D8794a7cAF327e7f855CD6Cd03BB8"}
// 	assert.ErrorContains(t, validator(ctx, addr), "no known prefix")

// 	addr = nexus.CrossChainAddress{Chain: exported.Scalarnet, Address: rand.AccAddr().String()}
// 	assert.NoError(t, validator(ctx, addr))

// 	addr = nexus.CrossChainAddress{Chain: nexus.Chain{Name: "osmosis", Module: types.ModuleName}, Address: "osmo18zhnqjv70v0d2f8v0s5lape0gr5ua94ewflhd5"}
// 	assert.NoError(t, validator(ctx, addr))
// }
