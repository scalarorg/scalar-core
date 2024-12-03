package btc

import (
	"cosmossdk.io/api/tendermint/abci"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, _ types.BaseKeeper) {}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, _ types.BaseKeeper, _ types.NexusKeeper, _ types.MultisigKeeper) ([]abci.ValidatorUpdate, error) {
	// handleConfirmedEvents(ctx, bk, n, m)
	// handleMessages(ctx, bk, n, m)

	return nil, nil
}
