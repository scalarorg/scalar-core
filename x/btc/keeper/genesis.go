package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/scalarorg/scalar-core/x/btc/types"
)

// InitGenesis initializes the state from a genesis file
func (k BaseKeeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	// TODO: add btc genesis
}

// ExportGenesis generates a genesis file from the state
func (k BaseKeeper) ExportGenesis(ctx sdk.Context) types.GenesisState {
	return *types.NewGenesisState(k.getParams(ctx), k.getVaultTag(ctx), k.getVaultVersion(ctx))
}

func (k BaseKeeper) getParams(ctx sdk.Context) types.Params {
	// TODO: add btc chains params
	return types.Params{}
}

func (k BaseKeeper) getVaultTag(ctx sdk.Context) types.VaultTag {
	// TODO: add btc vault tag
	return types.VaultTag{}
}

func (k BaseKeeper) getVaultVersion(ctx sdk.Context) types.VaultVersion {
	// TODO: add btc vault version
	return types.VaultVersion{}
}
