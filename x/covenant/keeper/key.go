package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/slices"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// GetKey returns the key of the given ID
func (k Keeper) GetKey(ctx sdk.Context, keyID multisig.KeyID) (multisig.Key, bool) {
	var key multisigTypes.Key
	ok := k.getStore(ctx).Get(keyPrefix.Append(utils.LowerCaseKey(keyID.String())), &key)
	if !ok {
		return nil, false
	}

	return &key, true
}

func (k Keeper) GetAllKeys(ctx sdk.Context) []multisig.Key {
	store := k.getStore(ctx)
	keys := []multisig.Key{}
	iter := store.Iterator(keyPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		key := multisigTypes.Key{}
		iter.UnmarshalValue(&key)
		keys = append(keys, &key)
	}
	return keys
}

// SetKey sets the given key
func (k Keeper) SetKey(ctx sdk.Context, key multisigTypes.Key) {
	k.getStore(ctx).Set(keyPrefix.Append(utils.LowerCaseKey(key.ID.String())), &key)

	participants := key.GetParticipants()
	k.Logger(ctx).Info("setting key",
		"key_id", key.ID,
		"participant_count", len(participants),
		"participants", strings.Join(slices.Map(participants, sdk.ValAddress.String), ", "),
		"participants_weight", key.GetParticipantsWeight().String(),
		"bonded_weight", key.Snapshot.BondedWeight.String(),
		"signing_threshold", key.SigningThreshold.String(),
	)
}

func (k Keeper) GetCurrentKeyID(ctx sdk.Context, chainName nexus.ChainName) (multisig.KeyID, bool) {
	return "", false
}
