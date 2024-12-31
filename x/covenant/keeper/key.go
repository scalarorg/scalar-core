package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// GetKey returns the key of the given ID
func (k Keeper) GetKey(ctx sdk.Context, keyID exported.KeyID) (exported.Key, bool) {
	// var key types.Key
	// ok := k.getStore(ctx).Get(keyPrefix.Append(utils.LowerCaseKey(keyID.String())), &key)
	// if !ok {
	// 	return nil, false
	// }

	// return &key, true
	return nil, false
}

// SetKey sets the given key
func (k Keeper) SetKey(ctx sdk.Context, key types.Key) {
	// k.setKey(ctx, key)

	// participants := key.GetParticipants()
	// events.Emit(ctx, types.NewKeygenCompleted(key.ID))
	// k.Logger(ctx).Info("setting key",
	// 	"key_id", key.ID,
	// 	"participant_count", len(participants),
	// 	"participants", strings.Join(slices.Map(participants, sdk.ValAddress.String), ", "),
	// 	"participants_weight", key.GetParticipantsWeight().String(),
	// 	"bonded_weight", key.Snapshot.BondedWeight.String(),
	// 	"signing_threshold", key.SigningThreshold.String(),
	// )
}

func (k Keeper) GetCurrentKeyID(ctx sdk.Context, chainName nexus.ChainName) (exported.KeyID, bool) {
	return "", false
}
