package keeper

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils"
	cov "github.com/scalarorg/scalar-core/x/covenant/exported"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
)

func (k Keeper) CreateCustodian(ctx sdk.Context, params types.Params) (err error) {
	return nil
}
func (k Keeper) GetCustodians(ctx sdk.Context) (custodians []*cov.Custodian, ok bool) {
	return nil, false
}

func (k Keeper) CreateCustodianGroup(ctx sdk.Context, params types.Params) (err error) {
	return nil
}

func (k Keeper) SetCustodian(ctx sdk.Context, custodian *cov.Custodian) {
	k.getStore(ctx).Set(custodianPrefix.Append(utils.KeyFromBz(custodian.BitcoinPubkey)), custodian)
}

func (k Keeper) SetCustodians(ctx sdk.Context, custodians []*cov.Custodian) {
	store := k.getStore(ctx)
	for _, custodian := range custodians {
		store.Set(custodianPrefix.Append(utils.KeyFromBz(custodian.BitcoinPubkey)), custodian)
	}
}

func (k Keeper) GetAllCustodians(ctx sdk.Context) ([]*cov.Custodian, bool) {
	protocols := []*cov.Custodian{}
	iter := k.getStoreIterator(ctx, custodianPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		protocol := cov.Custodian{}
		iter.UnmarshalValue(&protocol)
		protocols = append(protocols, &protocol)
	}
	return protocols, true
}

func (k Keeper) findCustodians(ctx sdk.Context, req *types.CustodiansRequest) ([]*cov.Custodian, bool) {
	custodians := []*cov.Custodian{}
	iter := k.getStoreIterator(ctx, custodianPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		custodian := cov.Custodian{}
		iter.UnmarshalValue(&custodian)
		if isMatchCustodian(&custodian, req) {
			custodians = append(custodians, &custodian)
		}
	}
	return custodians, true
}

func (k Keeper) GetCustodianGroup(ctx sdk.Context, uid string) (custodianGroup *cov.CustodianGroup, ok bool) {
	group := cov.CustodianGroup{}
	ok = k.getStore(ctx).Get(custodianGroupPrefix.Append(utils.KeyFromBz([]byte(uid))), &group)
	return &group, ok
}

func (k Keeper) SetCustodianGroup(ctx sdk.Context, custodianGroup *cov.CustodianGroup) {
	k.getStore(ctx).Set(custodianGroupPrefix.Append(utils.KeyFromBz([]byte(custodianGroup.UID))), custodianGroup)
}

func (k Keeper) SetCustodianGroups(ctx sdk.Context, custodianGroups []*cov.CustodianGroup) {
	store := k.getStore(ctx)
	for _, group := range custodianGroups {
		store.Set(custodianGroupPrefix.Append(utils.KeyFromBz([]byte(group.UID))), group)
	}
}

func (k Keeper) GetAllCustodianGroups(ctx sdk.Context) ([]*cov.CustodianGroup, bool) {
	custodianGroups := []*cov.CustodianGroup{}
	iter := k.getStoreIterator(ctx, custodianGroupPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		custodianGroup := cov.CustodianGroup{}
		iter.UnmarshalValue(&custodianGroup)
		custodianGroups = append(custodianGroups, &custodianGroup)
	}
	return custodianGroups, true
}

func (k Keeper) findCustodianGroups(ctx sdk.Context, req *types.GroupsRequest) ([]*cov.CustodianGroup, bool) {
	custodianGroups := []*cov.CustodianGroup{}
	iter := k.getStoreIterator(ctx, custodianGroupPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		custodianGroup := cov.CustodianGroup{}
		iter.UnmarshalValue(&custodianGroup)
		if isMatchCustodianGroup(&custodianGroup, req) {
			custodianGroups = append(custodianGroups, &custodianGroup)
		}
	}
	return custodianGroups, true
}

// Todo: Implement Matching function
func isMatchCustodian(protocol *cov.Custodian, req *types.CustodiansRequest) bool {
	match := true

	return match
}

// Todo: Implement Matching function
func isMatchCustodianGroup(group *cov.CustodianGroup, req *types.GroupsRequest) bool {
	if req.UID != "" && group.UID != req.UID {
		return false
	}

	return true
}

func (k Keeper) GetCustodianKeys(ctx sdk.Context, groupId string) ([]string, bool) {
	group, ok := k.GetCustodianGroup(ctx, groupId)
	if !ok {
		return nil, ok
	}
	keys := make([]string, len(group.Custodians))
	for i, custodian := range group.Custodians {
		keys[i] = hex.EncodeToString(custodian.BitcoinPubkey)
	}
	return keys, true
}
