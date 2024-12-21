package covenant

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(sdk.Context, abci.RequestBeginBlock, types.BaseKeeper) {}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, bk types.BaseKeeper) ([]abci.ValidatorUpdate, error) {
	clog.Yellow("COVENANT ABCI START ENDBLOCKER")
	// handleConfirmedEvents(ctx, bk, n, m)
	// handleMessages(ctx, bk, n, m)
	clog.Yellow("COVENANT ABCI FINISH ENDBLOCKER")
	return nil, nil
}
