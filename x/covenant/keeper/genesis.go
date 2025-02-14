package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/slices"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
)

// InitGenesis initializes the state from a genesis file
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	k.setParams(ctx, state.Params)
	slices.ForEach(state.SigningSessions, withContext(ctx, k.setSigningSession))

	k.setSigningSessionCount(ctx, uint64(len(state.SigningSessions)))

	k.SetCustodians(ctx, state.Custodians)
	k.SetCustodianGroups(ctx, state.Groups)
}

// ExportGenesis generates a genesis file from the state
func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	custodians, ok := k.GetAllCustodians(ctx)
	if !ok {
		custodians = []*types.Custodian{}
	}
	custodianGroups, ok := k.GetAllCustodianGroups(ctx)
	if !ok {
		custodianGroups = []*types.CustodianGroup{}
	}

	signingSessions, ok := k.GetSigningSessions(ctx)
	if !ok {
		signingSessions = []types.SigningSession{}
	}

	params := k.GetParams(ctx)

	return types.NewGenesisState(&params, signingSessions, custodians, custodianGroups)
}

func withContext[T any](ctx sdk.Context, fn func(sdk.Context, T)) func(T) {
	return func(t T) {
		fn(ctx, t)
	}
}
