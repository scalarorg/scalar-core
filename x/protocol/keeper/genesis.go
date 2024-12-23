package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/scalarorg/scalar-core/x/protocol/types"
)

// InitGenesis initializes the state from a genesis file
func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	k.SetProtocols(ctx, state.Protocols)
}

// ExportGenesis returns the reward module's genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	protocols, ok := k.GetAllProtocols(ctx)
	if !ok {
		return types.NewGenesisState(
			[]*types.Protocol{},
		)
	}

	return types.NewGenesisState(
		protocols,
	)
}
