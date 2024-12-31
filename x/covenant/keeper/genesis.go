package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
)

// InitGenesis initializes the state from a genesis file
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {

}

// ExportGenesis generates a genesis file from the state
func (k Keeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	// protocols, ok := k.GetProtocols(ctx)
	// if !ok {
	// 	return types.NewGenesisState([]*types.Protocol{})
	// }
	custodians, ok := k.GetCustodians(ctx)
	if !ok {
		custodians = []*types.Custodian{}
	}
	custodianGroup, ok := k.GetCustodianGroup(ctx)
	if !ok {
		custodianGroup = &types.CustodianGroup{}
	}

	signingSessions, ok := k.GetSigningSessions(ctx)
	if !ok {
		signingSessions = []types.SigningSession{}
	}

	params := k.GetParams(ctx)

	return types.NewGenesisState(&params, signingSessions, custodians, custodianGroup)
}
