package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/covenant/types"
)

func (k Keeper) GetSigningSessions(ctx sdk.Context) (signingSessions []types.SigningSession, ok bool) {
	return nil, false
}
