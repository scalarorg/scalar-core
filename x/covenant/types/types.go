package types

import (
	"bytes"
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
	"golang.org/x/exp/maps"
)

type Psbt []byte

func (p Psbt) ValidateBasic() error {
	return nil
}

type TapScriptSig []byte

// GetParticipants returns the participants of the given key
func (m Key) GetParticipants() []sdk.ValAddress {
	return sortAddresses(
		slices.Map(maps.Keys(m.PubKeys), func(a string) sdk.ValAddress { return funcs.Must(sdk.ValAddressFromBech32(a)) }),
	)
}

// GetParticipantsWeight returns the total weight of all participants who have submitted their public keys
func (m Key) GetParticipantsWeight() sdk.Uint {
	return slices.Reduce(m.GetParticipants(), sdk.ZeroUint(), func(total sdk.Uint, p sdk.ValAddress) sdk.Uint {
		return total.Add(m.Snapshot.GetParticipantWeight(p))
	})
}

// GetMinPassingWeight returns the minimum amount of weights required for the
// key to sign
func (m Key) GetMinPassingWeight() sdk.Uint {
	return m.Snapshot.CalculateMinPassingWeight(m.SigningThreshold)
}

// GetPubKey returns the public key of the given participant
func (m Key) GetPubKey(p sdk.ValAddress) (exported.PublicKey, bool) {
	pubKey, ok := m.PubKeys[p.String()]

	return pubKey, ok
}

// GetWeight returns the weight of the given participant
func (m Key) GetWeight(p sdk.ValAddress) sdk.Uint {
	return m.Snapshot.GetParticipantWeight(p)
}

// GetHeight returns the height of the key snapshot
func (m Key) GetHeight() int64 {
	return m.Snapshot.Height
}

// GetTimestamp returns the timestamp of the key snapshot
func (m Key) GetTimestamp() time.Time {
	return m.Snapshot.Timestamp
}

// GetBondedWeight returns the bonded weight of the key snapshot
func (m Key) GetBondedWeight() sdk.Uint {
	return m.Snapshot.BondedWeight
}

// ValidateBasic returns an error if the given key is invalid; nil otherwise
func (m Key) ValidateBasic() error {
	for _, pubKey := range m.PubKeys {
		if err := pubKey.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}

func sortAddresses[T sdk.Address](addrs []T) []T {
	sorted := make([]T, len(addrs))
	copy(sorted, addrs)

	sort.SliceStable(sorted, func(i, j int) bool { return bytes.Compare(sorted[i].Bytes(), sorted[j].Bytes()) < 0 })

	return sorted
}
