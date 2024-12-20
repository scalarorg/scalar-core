package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/x/protocol/types"
	"github.com/tendermint/tendermint/libs/log"
)

var (
	protocolPrefix = utils.KeyFromStr("protocol")
)

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	paramSpace paramtypes.Subspace
}

func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, paramSpace paramtypes.Subspace) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace,
	}
}

// GetParams gets the permission module's parameters
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.paramSpace.GetParamSet(ctx, &p)
	return p
}

// setParams sets the permission module's parameters
func (k Keeper) setParams(ctx sdk.Context, p types.Params) {
	k.paramSpace.SetParamSet(ctx, &p)
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetProtocols(ctx sdk.Context) ([]*types.Protocol, bool) {
	store := ctx.KVStore(k.storeKey)
	protocols := []*types.Protocol{}
	iter := sdk.KVStorePrefixIterator(store, protocolPrefix.AsKey())
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		protocol := types.Protocol{}
		k.cdc.MustUnmarshal(iter.Value(), &protocol)
		protocols = append(protocols, &protocol)
	}
	return protocols, true
}
