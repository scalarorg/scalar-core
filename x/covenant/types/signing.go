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
	exported "github.com/scalarorg/scalar-core/x/covenant/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
)

var _ codectypes.UnpackInterfacesMessage = SigningSession{}

type NewSigningSessionParams struct {
	ID             uint64
	Key            multisigTypes.Key
	ExpiresAt      int64
	GracePeriod    int64
	Module         string
	ModuleMetadata []codec.ProtoMarshaler
	Psbt           Psbt
}

// NewSigningSession is the contructor for signing session
func NewSigningSession(params *NewSigningSessionParams) SigningSession {
	var moduleMetadata *codectypes.Any
	if len(params.ModuleMetadata) > 0 {
		moduleMetadata = funcs.Must(codectypes.NewAnyWithValue(params.ModuleMetadata[0]))
	}

	return SigningSession{
		ID: params.ID,
		PsbtMultiSig: PsbtMultiSig{
			KeyID: params.Key.ID,
			Psbt:  params.Psbt,
		},
		State:          exported.Pending,
		Key:            params.Key,
		ExpiresAt:      params.ExpiresAt,
		GracePeriod:    params.GracePeriod,
		Module:         params.Module,
		ModuleMetadata: moduleMetadata,
	}
}

// ValidateBasic returns an error if the given signing session is invalid; nil otherwise
func (m SigningSession) ValidateBasic() error {
	if err := m.Key.ValidateBasic(); err != nil {
		return err
	}

	if m.Key.ID != m.PsbtMultiSig.KeyID {
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

		if err := m.PsbtMultiSig.ValidateBasic(); err != nil {
			return err
		}

		if m.GetParticipantsWeight().LT(m.Key.GetMinPassingWeight()) {
			return fmt.Errorf("completed signing session must have completed multi signature")
		}
	default:
		return fmt.Errorf("unexpected state %s", m.GetState())
	}

	for addr, sigs := range m.PsbtMultiSig.ParticipantTapScriptSigs {
		pubKey, ok := m.Key.PubKeys[addr]
		if !ok {
			return fmt.Errorf("participant %s does not have public key submitted", addr)
		}

		// TODO: Implement signature verification
		_ = sigs
		_ = pubKey
		clog.Redf("ValidateBasic, TODO: Implement signature verification for tapscriptsig")
		// if !sig.Verify(m.PsbtMultiSig.TapScriptHash, pubKey) {
		// 	return fmt.Errorf("signature does not match the public key")
		// }
	}

	return nil
}

// AddTapScriptSigs adds the given tapscript sigs to the signing session
func (m *SigningSession) AddTapScriptSigs(blockHeight int64, participant sdk.ValAddress, inputSigs *exported.TapScriptSigList) error {
	if m.PsbtMultiSig.ParticipantTapScriptSigs == nil {
		m.PsbtMultiSig.ParticipantTapScriptSigs = make(map[string]*exported.TapScriptSigList)
	}

	if m.isExpired(blockHeight) {
		return fmt.Errorf("signing session %d has expired", m.GetID())
	}

	if _, ok := m.Key.PubKeys[participant.String()]; !ok {
		return fmt.Errorf("%s is not a participant of signing %d", participant.String(), m.GetID())
	}

	if _, ok := m.PsbtMultiSig.ParticipantTapScriptSigs[participant.String()]; ok {
		return fmt.Errorf("participant %s already submitted its signature for signing %d", participant.String(), m.GetID())
	}

	// TODO: Implement signature verification
	clog.Yellow("AddTapScriptSigs, TODO: Implement signature verification for tapscriptsig\n")
	// if !sig.Verify(m.PsbtMultiSig.PayloadHash, m.Key.PubKeys[participant.String()]) {
	// 	return fmt.Errorf("invalid signature received from participant %s for signing %d", participant.String(), m.GetID())
	// }

	if m.GetState() == exported.Completed && !m.isWithinGracePeriod(blockHeight) {
		return fmt.Errorf("signing session %d has closed", m.GetID())
	}

	m.addSig(participant, inputSigs)

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
		_, ok := m.PsbtMultiSig.ParticipantTapScriptSigs[p.String()]

		return !ok
	})
}

// Result returns the generated multi signature if the session is completed and the multi signature is valid
func (m SigningSession) Result() (PsbtMultiSig, error) {
	if m.GetState() != exported.Completed {
		return PsbtMultiSig{}, fmt.Errorf("signing %d is not completed yet", m.GetID())
	}

	if m.GetParticipantsWeight().LT(m.Key.GetMinPassingWeight()) {
		panic(fmt.Errorf("multi sig is not completed yet"))
	}
	funcs.MustNoErr(m.PsbtMultiSig.ValidateBasic())

	return m.PsbtMultiSig, nil
}

// GetParticipantsWeight returns the total weights of the participants
func (m SigningSession) GetParticipantsWeight() sdk.Uint {
	return slices.Reduce(m.PsbtMultiSig.GetParticipants(), sdk.ZeroUint(), func(total sdk.Uint, p sdk.ValAddress) sdk.Uint {
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

func (m *SigningSession) addSig(participant sdk.ValAddress, sigs *exported.TapScriptSigList) {
	clog.Redf("addSig, participant: %+s", participant.String())
	clog.Redf("addSig, sigs: %+v", sigs)
	m.PsbtMultiSig.ParticipantTapScriptSigs[participant.String()] = sigs
}

func (m SigningSession) isWithinGracePeriod(blockHeight int64) bool {
	return blockHeight <= m.CompletedAt+m.GracePeriod
}

func (m SigningSession) isExpired(blockHeight int64) bool {
	return blockHeight >= m.ExpiresAt
}

// ValidateBasic returns an error if the given sig is invalid; nil otherwise
func (m PsbtMultiSig) ValidateBasic() error {
	clog.Magenta("validate m.KeyID", m.KeyID)
	if err := m.KeyID.ValidateBasic(); err != nil {
		return err
	}

	clog.Magenta("validate m.Psbt")
	if err := m.Psbt.ValidateBasic(); err != nil {
		return err
	}

	clog.Magenta("validate m.ParticipantTapScriptSigs, len: ", len(m.ParticipantTapScriptSigs))

	signatureSeen := make(map[string]bool, len(m.ParticipantTapScriptSigs))
	numSigs := -1
	for address, sigs := range m.ParticipantTapScriptSigs {
		if numSigs == -1 {
			numSigs = sigs.Size()
		}
		if numSigs != sigs.Size() {
			return fmt.Errorf("participant %s has different number of signatures", address)
		}
		if signatureSeen[address] {
			return fmt.Errorf("duplicate signature seen")
		}
		signatureSeen[address] = true

		if _, err := sdk.ValAddressFromBech32(address); err != nil {
			clog.Magenta("validate m.TapScriptSigs vald address, address: ", address)
			return err
		}

		if err := sigs.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}

// GetSignature returns the ECDSA signature of the given participant
func (m PsbtMultiSig) GetTapScriptSigs(p sdk.ValAddress) (*exported.TapScriptSigList, bool) {
	sigs, ok := m.ParticipantTapScriptSigs[p.String()]
	if !ok {
		return nil, false
	}

	return sigs, true
}

// GetParticipants returns the participants of the given multi sig
func (m PsbtMultiSig) GetParticipants() []sdk.ValAddress {
	return multisigTypes.SortAddresses(
		slices.Map(maps.Keys(m.ParticipantTapScriptSigs), func(a string) sdk.ValAddress { return funcs.Must(sdk.ValAddressFromBech32(a)) }),
	)
}
