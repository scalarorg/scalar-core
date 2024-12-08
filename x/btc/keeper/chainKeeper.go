package keeper

import (
	"fmt"

	"github.com/axelarnetwork/axelar-core/utils/key"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

var _ types.ChainKeeper = chainKeeper{}

type chainKeeper struct {
	internalKeeper
	chain nexus.ChainName
}

func (k chainKeeper) GetName() nexus.ChainName {
	return k.chain
}

// GetParams gets the evm module's parameters
func (k chainKeeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.getSubspace().GetParamSet(ctx, &p)
	return p
}

// func (k chainKeeper) GetChainID(ctx sdk.Context) (sdk.Int, bool) {
// 	network := k.GetNetwork(ctx)
// 	return k.GetChainIDByNetwork(ctx, network)
// }

// // GetNetwork returns the EVM network Axelar-Core is expected to connect to
// func (k chainKeeper) GetNetwork(ctx sdk.Context) string {
// 	return getParam[string](k, ctx, types.KeyNetwork)
// }

func (k chainKeeper) GetRequiredConfirmationHeight(ctx sdk.Context) uint64 {
	return getParam[uint64](k, ctx, types.KeyConfirmationHeight)
}

func getParam[T any](k chainKeeper, ctx sdk.Context, paramKey []byte) T {
	var value T
	k.getSubspace().Get(ctx, paramKey, &value)
	return value
}

func (k chainKeeper) getSubspace() params.Subspace {
	chainKey := key.FromStr(types.ModuleName).Append(key.From(k.chain))
	subspace, ok := k.paramsKeeper.GetSubspace(chainKey.String())
	if !ok {
		panic(fmt.Sprintf("subspace for chain %s does not exist", k.chain))
	}
	return subspace
}
