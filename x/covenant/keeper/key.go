package keeper

import (
	"encoding/hex"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/slices"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	snapshot "github.com/scalarorg/scalar-core/x/snapshot/exported"
)

const (
	MOCK_CURRENT_KEY_ID = "mock|123456789"
)

// GetKey returns the key of the given ID
func (k Keeper) GetKey(ctx sdk.Context, keyID multisig.KeyID) (multisig.Key, bool) {
	// For Transactional model, the keyID is the custodians pubkey + txid
	id := keyID.String()
	buffer, err := hex.DecodeString(id)
	if err != nil {
		return nil, false
	}
	if len(buffer) > 32 {
		len := len(buffer) - 32
		id = hex.EncodeToString(buffer[:len])
		k.Logger(ctx).Info("Transactional model keyID", "fullFeyID", keyID, "originKeyId", id)
	}
	var key multisigTypes.Key
	ok := k.getStore(ctx).Get(keyPrefix.Append(utils.LowerCaseKey(id)), &key)
	if !ok {
		return nil, false
	}

	return &key, true
}

func (k Keeper) GetAllKeys(ctx sdk.Context) []multisigTypes.Key {
	store := k.getStore(ctx)
	keys := []multisigTypes.Key{}
	iter := store.Iterator(keyPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		key := multisigTypes.Key{}
		iter.UnmarshalValue(&key)
		keys = append(keys, key)
	}
	return keys
}

// SetKey sets the given key
func (k Keeper) SetKey(ctx sdk.Context, key multisigTypes.Key) {
	k.getStore(ctx).Set(keyPrefix.Append(utils.LowerCaseKey(key.ID.String())), &key)

	// TODO: FIX THIS PROBLEM
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

// func (k Keeper) GetCurrentKeyID(ctx sdk.Context, chainName nexus.ChainName) (multisig.KeyID, bool) {
// 	return MOCK_CURRENT_KEY_ID, true
// }

func (k Keeper) getKey(ctx sdk.Context, id multisig.KeyID) (key multisigTypes.Key, ok bool) {
	//return k._MustRemoveFakeKey(ctx), true
	return key, k.getStore(ctx).Get(keyPrefix.Append(utils.LowerCaseKey(id.String())), &key)
}

// ID               github_com_scalarorg_scalar_core_x_multisig_exported.KeyID                `protobuf:"bytes,1,opt,name=id,proto3,casttype=github.com/scalarorg/scalar-core/x/multisig/exported.KeyID" json:"id,omitempty"`
// 	Snapshot         exported.Snapshot                                                         `protobuf:"bytes,2,opt,name=snapshot,proto3" json:"snapshot"`
// 	PubKeys          map[string]github_com_scalarorg_scalar_core_x_multisig_exported.PublicKey `protobuf:"bytes,3,rep,name=pub_keys,json=pubKeys,proto3,castvalue=github.com/scalarorg/scalar-core/x/multisig/exported.PublicKey" json:"pub_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`

func (k Keeper) _MustRemoveFakeKey(ctx sdk.Context) multisigTypes.Key {
	mockPubKey := map[string]string{
		"scalarvaloper1u8ennvwsshneu4nvec38e3jvcxmppf7lfq3pf5": "0215da913b3e87b4932b1e1b87d9667c28e7250aa0ed60b3a31095f541e1641488",
		"scalarvaloper1qyhnraq6dwpl69heem9qd4phrd946lct6xdgqg": "02f0f3d9beaf7a3945bcaa147e041ae1d5ca029bde7e40d8251f0783d6ecbe8fb5",
		"scalarvaloper1tc9q2ngj29kllv4kfrqjh6gj4hrsp6p35zpyep": "03594e78c0a2968210d9c1550d4ad31b03d5e4b9659cf2f67842483bb3c2bb7811",
		"scalarvaloper105wsknjutftqksqtugr83avrlsxdz0tvc96yhk": "03b59e575cef873ea95273afd55956c84590507200d410e693e4b079a426cc6102",
	}

	pubKeys := make(map[string]multisig.PublicKey)
	participants := map[string]snapshot.Participant{}

	for address, pubkey := range mockPubKey {
		var pk multisig.PublicKey
		pk, err := pk.FromHex(pubkey)
		if err != nil {
			panic(err)
		}
		pubKeys[address] = pk

		addr, err := sdk.ValAddressFromBech32(address)
		if err != nil {
			panic(err)
		}
		participants[address] = snapshot.Participant{
			Address: addr,
			Weight:  sdk.NewUint(100),
		}
	}

	return multisigTypes.Key{
		ID: MOCK_CURRENT_KEY_ID,
		Snapshot: snapshot.Snapshot{
			Timestamp:    ctx.BlockTime(),
			Height:       ctx.BlockHeight(),
			Participants: participants,
			BondedWeight: sdk.NewUint(400),
		},
		PubKeys: pubKeys,
		SigningThreshold: utils.Threshold{
			Numerator:   3,
			Denominator: 4,
		},
		State: multisig.Active,
	}
}
