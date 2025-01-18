package keeper

import (
	"bytes"
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
func (k Keeper) SetParams(ctx sdk.Context, p types.Params) {
	k.paramSpace.SetParamSet(ctx, &p)
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetProtocol(ctx sdk.Context, protocol *types.Protocol) {
	k.getStore(ctx).Set(protocolPrefix.Append(utils.KeyFromBz(protocol.ScalarAddress)), protocol)
}
func (k Keeper) SetProtocols(ctx sdk.Context, protocols []*types.Protocol) {
	store := k.getStore(ctx)
	for _, protocol := range protocols {
		store.Set(protocolPrefix.Append(utils.KeyFromBz(protocol.ScalarAddress)), protocol)
	}
}
func (k Keeper) GetAllProtocols(ctx sdk.Context) ([]*types.Protocol, bool) {
	store := k.getStore(ctx)
	protocols := []*types.Protocol{}
	iter := store.Iterator(protocolPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		protocol := types.Protocol{}
		iter.UnmarshalValue(&protocol)
		protocols = append(protocols, &protocol)
	}
	return protocols, true
}

func (k Keeper) findProtocols(ctx sdk.Context, req *types.ProtocolsRequest) ([]*types.Protocol, bool) {
	store := k.getStore(ctx)
	protocols := []*types.Protocol{}
	iter := store.Iterator(protocolPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		protocol := types.Protocol{}
		iter.UnmarshalValue(&protocol)
		if isMatch(&protocol, req) {
			protocols = append(protocols, &protocol)
		}
	}
	return protocols, true
}

func (k Keeper) getProtocolByAddress(ctx sdk.Context, address []byte) (*types.Protocol, bool) {
	store := k.getStore(ctx)
	iter := store.Iterator(protocolPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		protocol := types.Protocol{}
		iter.UnmarshalValue(&protocol)
		if bytes.Compare(protocol.ScalarAddress, address) == 0 {
			return &protocol, true
		}
	}
	return nil, false
}

// Todo: Implement Matching function
func isMatch(protocol *types.Protocol, req *types.ProtocolsRequest) bool {
	match := true
	if req.Address != "" {

	}
	return match
}
func (k Keeper) getStore(ctx sdk.Context) utils.KVStore {
	return utils.NewNormalizedStore(ctx.KVStore(k.storeKey), k.cdc)
}
