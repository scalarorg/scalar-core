package types

import (
	fmt "fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/exp/maps"

	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
)

var _ codectypes.UnpackInterfacesMessage = SigningSession{}

// NewSigningSession is the contructor for signing session
func NewSigningSession(id uint64, key Key, psbt Psbt, expiresAt int64, gracePeriod int64, module string, moduleMetadataProto ...codec.ProtoMarshaler) SigningSession {
	var moduleMetadata *codectypes.Any
	if len(moduleMetadataProto) > 0 {
		moduleMetadata = funcs.Must(codectypes.NewAnyWithValue(moduleMetadataProto[0]))
	}

	return SigningSession{
		ID: id,
		MultiSig: MultiSig{
			KeyID: key.ID,
			Psbt:  psbt,
		},
		State:          exported.Pending,
		Key:            key,
		ExpiresAt:      expiresAt,
		GracePeriod:    gracePeriod,
		Module:         module,
		ModuleMetadata: moduleMetadata,
	}
}

// ValidateBasic returns an error if the given signing session is invalid; nil otherwise
func (m SigningSession) ValidateBasic() error {
	if err := m.MultiSig.ValidateBasic(); err != nil {
		return err
	}

	if err := m.Key.ValidateBasic(); err != nil {
		return err
	}

	if m.Key.ID != m.MultiSig.KeyID {
		return fmt.Errorf("key ID mismatch")
	}

	if m.ExpiresAt <= 0 {
		return fmt.Errorf("expires at must be >0")
	}

	if m.CompletedAt >= m.ExpiresAt {
		return fmt.Errorf("completed at must be < expires at")
	}

	if m.GracePeriod < 0 {
		return fmt.Errorf("grace period must be >=0")
	}

	if err := utils.ValidateString(m.Module); err != nil {
		return err
	}

	switch m.GetState() {
	case exported.Pending:
		if m.CompletedAt != 0 {
			return fmt.Errorf("pending signing session must not have completed at set")
		}
	case exported.Completed:
		if m.CompletedAt <= 0 {
			return fmt.Errorf("completed signing session must have completed at set")
		}

		if m.GetParticipantsWeight().LT(m.Key.GetMinPassingWeight()) {
			return fmt.Errorf("completed signing session must have completed multi signature")
		}
	default:
		return fmt.Errorf("unexpected state %s", m.GetState())
	}

	for addr, sig := range m.MultiSig.TapScriptSigs {
		pubKey, ok := m.Key.PubKeys[addr]
		if !ok {
			return fmt.Errorf("participant %s does not have public key submitted", addr)
		}

		// TODO: Implement signature verification
		_ = sig
		_ = pubKey
		clog.Redf("TODO: Implement signature verification for tapscriptsig")
		// if !sig.Verify(m.MultiSig.TapScriptHash, pubKey) {
		// 	return fmt.Errorf("signature does not match the public key")
		// }
	}

	return nil
}

// AddSig adds a new signature for the given participant into the signing session
func (m *SigningSession) AddSig(blockHeight int64, participant sdk.ValAddress, sig exported.TapScriptSig) error {
	if m.MultiSig.TapScriptSigs == nil {
		m.MultiSig.TapScriptSigs = make(map[string]*exported.TapScriptSig)
	}

	if m.isExpired(blockHeight) {
		return fmt.Errorf("signing session %d has expired", m.GetID())
	}

	if _, ok := m.Key.PubKeys[participant.String()]; !ok {
		return fmt.Errorf("%s is not a participant of signing %d", participant.String(), m.GetID())
	}

	if _, ok := m.MultiSig.TapScriptSigs[participant.String()]; ok {
		return fmt.Errorf("participant %s already submitted its signature for signing %d", participant.String(), m.GetID())
	}

	// TODO: Implement signature verification
	_ = sig
	_ = m.Key.PubKeys[participant.String()]
	clog.Redf("TODO: Implement signature verification for tapscriptsig")
	// if !sig.Verify(m.MultiSig.PayloadHash, m.Key.PubKeys[participant.String()]) {
	// 	return fmt.Errorf("invalid signature received from participant %s for signing %d", participant.String(), m.GetID())
	// }

	if m.GetState() == exported.Completed && !m.isWithinGracePeriod(blockHeight) {
		return fmt.Errorf("signing session %d has closed", m.GetID())
	}

	m.addSig(participant, sig)

	if m.GetState() != exported.Completed && m.GetParticipantsWeight().GTE(m.Key.GetMinPassingWeight()) {
		m.CompletedAt = blockHeight
		m.State = exported.Completed
	}

	return nil
}

// GetMissingParticipants returns all participants who failed to submit their signatures
func (m SigningSession) GetMissingParticipants() []sdk.ValAddress {
	participants := m.Key.GetParticipants()

	return slices.Filter(participants, func(p sdk.ValAddress) bool {
		_, ok := m.MultiSig.TapScriptSigs[p.String()]

		return !ok
	})
}

// Result returns the generated multi signature if the session is completed and the multi signature is valid
func (m SigningSession) Result() (MultiSig, error) {
	if m.GetState() != exported.Completed {
		return MultiSig{}, fmt.Errorf("signing %d is not completed yet", m.GetID())
	}

	if m.GetParticipantsWeight().LT(m.Key.GetMinPassingWeight()) {
		panic(fmt.Errorf("multi sig is not completed yet"))
	}
	funcs.MustNoErr(m.MultiSig.ValidateBasic())

	return m.MultiSig, nil
}

// GetParticipantsWeight returns the total weights of the participants
func (m SigningSession) GetParticipantsWeight() sdk.Uint {
	return slices.Reduce(m.MultiSig.GetParticipants(), sdk.ZeroUint(), func(total sdk.Uint, p sdk.ValAddress) sdk.Uint {
		return total.Add(m.Key.Snapshot.GetParticipantWeight(p))
	})
}

// UnpackInterfaces implements UnpackInterfacesMessage
func (m SigningSession) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var data codec.ProtoMarshaler

	return unpacker.UnpackAny(m.ModuleMetadata, &data)
}

// GetMetadata returns the unpacked module metadata
func (m SigningSession) GetMetadata() codec.ProtoMarshaler {
	if m.ModuleMetadata == nil {
		return nil
	}

	return m.ModuleMetadata.GetCachedValue().(codec.ProtoMarshaler)
}

func (m *SigningSession) addSig(participant sdk.ValAddress, sig exported.TapScriptSig) {
	m.MultiSig.TapScriptSigs[participant.String()] = &sig
}

func (m SigningSession) isWithinGracePeriod(blockHeight int64) bool {
	return blockHeight <= m.CompletedAt+m.GracePeriod
}

func (m SigningSession) isExpired(blockHeight int64) bool {
	return blockHeight >= m.ExpiresAt
}

// ValidateBasic returns an error if the given sig is invalid; nil otherwise
func (m MultiSig) ValidateBasic() error {
	if err := m.KeyID.ValidateBasic(); err != nil {
		return err
	}

	if err := m.Psbt.ValidateBasic(); err != nil {
		return err
	}

	signatureSeen := make(map[string]bool, len(m.TapScriptSigs))
	for address, sig := range m.TapScriptSigs {
		sigHex := sig.String()
		if signatureSeen[sigHex] {
			return fmt.Errorf("duplicate signature seen")
		}
		signatureSeen[sigHex] = true

		if _, err := sdk.ValAddressFromBech32(address); err != nil {
			return err
		}

		if err := sig.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}

// GetSignature returns the ECDSA signature of the given participant
func (m MultiSig) GetTapScriptSig(p sdk.ValAddress) (*exported.TapScriptSig, bool) {
	sig, ok := m.TapScriptSigs[p.String()]
	if !ok {
		return nil, false
	}

	return sig, true
}

// GetParticipants returns the participants of the given multi sig
func (m MultiSig) GetParticipants() []sdk.ValAddress {
	return sortAddresses(
		slices.Map(maps.Keys(m.TapScriptSigs), func(a string) sdk.ValAddress { return funcs.Must(sdk.ValAddressFromBech32(a)) }),
	)
}
