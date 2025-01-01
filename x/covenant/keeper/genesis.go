package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
)

// InitGenesis initializes the state from a genesis file
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
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
