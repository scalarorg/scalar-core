package keeper

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
)

func (k Keeper) CreateCustodian(ctx sdk.Context, params types.Params) (err error) {
	return nil
}
func (k Keeper) GetCustodians(ctx sdk.Context) (custodians []*types.Custodian, ok bool) {
	return nil, false
}

func (k Keeper) CreateCustodianGroup(ctx sdk.Context, params types.Params) (err error) {
	return nil
}

func (k Keeper) SetCustodian(ctx sdk.Context, custodian *types.Custodian) {
	k.getStore(ctx).Set(custodianPrefix.Append(utils.KeyFromBz(custodian.BtcPubkey)), custodian)
}
func (k Keeper) SetCustodians(ctx sdk.Context, custodians []*types.Custodian) {
	store := k.getStore(ctx)
	for _, custodian := range custodians {
		store.Set(custodianPrefix.Append(utils.KeyFromBz(custodian.BtcPubkey)), custodian)
	}
}
func (k Keeper) GetAllCustodians(ctx sdk.Context) ([]*types.Custodian, bool) {
	protocols := []*types.Custodian{}
	iter := k.getStoreIterator(ctx, custodianPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		protocol := types.Custodian{}
		iter.UnmarshalValue(&protocol)
		protocols = append(protocols, &protocol)
	}
	return protocols, true
}

func (k Keeper) findCustodians(ctx sdk.Context, req *types.CustodiansRequest) ([]*types.Custodian, bool) {
	custodians := []*types.Custodian{}
	iter := k.getStoreIterator(ctx, custodianPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		custodian := types.Custodian{}
		iter.UnmarshalValue(&custodian)
		if isMatchCustodian(&custodian, req) {
			custodians = append(custodians, &custodian)
		}
	}
	return custodians, true
}

func (k Keeper) SetCustodianGroup(ctx sdk.Context, custodianGroup *types.CustodianGroup) {
	k.getStore(ctx).Set(custodianGroupPrefix.Append(utils.KeyFromBz([]byte(custodianGroup.Uid))), custodianGroup)
}
func (k Keeper) SetCustodianGroups(ctx sdk.Context, custodianGroups []*types.CustodianGroup) {
	store := k.getStore(ctx)
	for _, group := range custodianGroups {
		store.Set(custodianGroupPrefix.Append(utils.KeyFromBz([]byte(group.Uid))), group)
	}
}
func (k Keeper) GetAllCustodianGroups(ctx sdk.Context) ([]*types.CustodianGroup, bool) {
	custodianGroups := []*types.CustodianGroup{}
	iter := k.getStoreIterator(ctx, custodianGroupPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		custodianGroup := types.CustodianGroup{}
		iter.UnmarshalValue(&custodianGroup)
		custodianGroups = append(custodianGroups, &custodianGroup)
	}
	return custodianGroups, true
}

func (k Keeper) findCustodianGroups(ctx sdk.Context, req *types.GroupsRequest) ([]*types.CustodianGroup, bool) {
	custodianGroups := []*types.CustodianGroup{}
	iter := k.getStoreIterator(ctx, custodianGroupPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		custodianGroup := types.CustodianGroup{}
		iter.UnmarshalValue(&custodianGroup)
		if isMatchCustodianGroup(&custodianGroup, req) {
			custodianGroups = append(custodianGroups, &custodianGroup)
		}
	}
	return custodianGroups, true
}

// Todo: Implement Matching function
func isMatchCustodian(protocol *types.Custodian, req *types.CustodiansRequest) bool {
	match := true

	return match
}

// Todo: Implement Matching function
func isMatchCustodianGroup(protocol *types.CustodianGroup, req *types.GroupsRequest) bool {
	match := true

	return match
}
func (k Keeper) GetCustodianGroup(ctx sdk.Context, groupId string) (custodianGroup *types.CustodianGroup, ok bool) {
	return &types.CustodianGroup{}, false
}

func (k Keeper) GetCustodianKeys(ctx sdk.Context, groupId string) ([]string, bool) {
	group, ok := k.GetCustodianGroup(ctx, groupId)
	if !ok {
		return nil, ok
	}
	keys := make([]string, len(group.Custodians))
	for i, custodian := range group.Custodians {
		keys[i] = hex.EncodeToString(custodian.BtcPubkey)
	}
	return keys, true
}
