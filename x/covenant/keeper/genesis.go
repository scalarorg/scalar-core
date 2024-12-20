package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
)

// InitGenesis initializes the state from a genesis file
func (k BaseKeeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {

}

// ExportGenesis generates a genesis file from the state
func (k BaseKeeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
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
	return types.NewGenesisState(custodians, custodianGroup)
}
