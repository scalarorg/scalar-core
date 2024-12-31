package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k Keeper) GetCustodianGroup(ctx sdk.Context) (custodianGroup *types.CustodianGroup, ok bool) {
	return &types.CustodianGroup{}, false
}
