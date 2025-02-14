package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/events"

	"github.com/scalarorg/scalar-core/x/covenant/types"
	multisig "github.com/scalarorg/scalar-core/x/multisig/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// RotateKey rotates to the given chain's next key
func (k Keeper) RotateKey(ctx sdk.Context,
	chainName nexus.ChainName,
	key multisig.Key,
) error {

	k.SetKey(ctx, key)

	events.Emit(ctx, types.NewKeyRotated(chainName, key.GetID()))
	k.Logger(ctx).Info("new key rotated",
		"chain", chainName,
		"keyID", key.GetID(),
	)

	return nil
}
