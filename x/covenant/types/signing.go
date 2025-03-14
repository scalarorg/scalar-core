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
	MultiPsbt      []Psbt
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
			KeyID:        params.Key.ID,
			MultiPsbt:    params.MultiPsbt,
			FinalizedTxs: make([][]byte, len(params.MultiPsbt)),
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

		if m.GetNumberOfSignedPsbtParticipants().LT(m.GetMinPassingSignedParticipants()) {
			return fmt.Errorf("completed signing session must have completed multi signature")
		}
	default:
		return fmt.Errorf("unexpected state %s", m.GetState())
	}

	for addr, sigs := range m.PsbtMultiSig.ParticipantListTapScriptSigs {
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

// AddListOfTapScriptSigs adds the given tapscript sigs to the signing session
func (m *SigningSession) AddListOfTapScriptSigs(blockHeight int64, participant sdk.ValAddress, list []*exported.TapScriptSigsMap) error {
	if len(list) == 0 {
		return fmt.Errorf("no signatures submitted")
	}

	if m.PsbtMultiSig.ParticipantListTapScriptSigs == nil {
		m.PsbtMultiSig.ParticipantListTapScriptSigs = make(map[string]*exported.ListOfTapScriptSigsMap)
	}

	if m.isExpired(blockHeight) {
		return fmt.Errorf("signing session %d has expired", m.GetID())
	}

	if _, ok := m.Key.PubKeys[participant.String()]; !ok {
		return fmt.Errorf("%s is not a participant of signing %d", participant.String(), m.GetID())
	}

	if _, ok := m.PsbtMultiSig.ParticipantListTapScriptSigs[participant.String()]; ok {
		return fmt.Errorf("participant %s already submitted its signature for signing %d", participant.String(), m.GetID())
	}

	// TODO: Implement signature verification
	clog.Yellow("AddListOfTapScriptSigs, TODO: Implement signature verification for tapscriptsig\n")
	// if !sig.Verify(m.PsbtMultiSig.PayloadHash, m.Key.PubKeys[participant.String()]) {
	// 	return fmt.Errorf("invalid signature received from participant %s for signing %d", participant.String(), m.GetID())
	// }

	// TODO: use isWithinGracePeriod to determine who is signed to reward
	//if m.GetState() == exported.Completed && !m.isWithinGracePeriod(blockHeight) {
	if m.GetState() == exported.Completed {
		return fmt.Errorf("psbts were completed in session %d", m.GetID())
	}

	if len(m.PsbtMultiSig.MultiPsbt) != len(list) {
		return fmt.Errorf("number of psbt does not match")
	}

	m.addSig(participant, list)

	total := m.GetNumberOfSignedPsbtParticipants()
	min := m.GetMinPassingSignedParticipants()

	clog.Greenf(">>> [AddListOfTapScriptSigs], total: %+v", total)
	clog.Greenf(">>> [AddListOfTapScriptSigs], min: %+v", min)

	if m.GetState() != exported.Completed && total.GTE(min) {
		m.CompletedAt = blockHeight
		m.State = exported.Completed
	}

	return nil
}

// GetMissingParticipants returns all participants who failed to submit their signatures
func (m SigningSession) GetMissingParticipants() []sdk.ValAddress {
	participants := m.Key.GetParticipants()

	return slices.Filter(participants, func(p sdk.ValAddress) bool {
		_, ok := m.PsbtMultiSig.ParticipantListTapScriptSigs[p.String()]

		return !ok
	})
}

// Result returns the generated multi signature if the session is completed and the multi signature is valid
func (m SigningSession) Result() (PsbtMultiSig, error) {
	if m.GetState() != exported.Completed {
		return PsbtMultiSig{}, fmt.Errorf("signing %d is not completed yet", m.GetID())
	}

	if m.GetNumberOfSignedPsbtParticipants().LT(m.GetMinPassingSignedParticipants()) {
		panic(fmt.Errorf("multi sig is not completed yet"))
	}
	funcs.MustNoErr(m.PsbtMultiSig.ValidateBasic())

	return m.PsbtMultiSig, nil
}

func (m SigningSession) GetMinPassingSignedParticipants() sdk.Uint {
	return sdk.NewUint(uint64(m.Key.SigningThreshold.Numerator))
}

func (m SigningSession) GetNumberOfSignedPsbtParticipants() sdk.Uint {
	return sdk.NewUint(uint64(len(m.PsbtMultiSig.ParticipantListTapScriptSigs)))
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

func (m *SigningSession) addSig(participant sdk.ValAddress, sigs []*exported.TapScriptSigsMap) {
	clog.Redf("addSig, participant: %+s", participant.String())
	clog.Redf("addSig, sigs: %+v", sigs)
	m.PsbtMultiSig.ParticipantListTapScriptSigs[participant.String()] = &exported.ListOfTapScriptSigsMap{
		Inner: sigs,
	}
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
	for _, psbt := range m.MultiPsbt {
		clog.Magenta("validate m.Psbt, psbt: ", psbt)
		if err := psbt.ValidateBasic(); err != nil {
			return err
		}
	}

	clog.Magenta("validate m.ParticipantTapScriptSigs, len: ", len(m.ParticipantListTapScriptSigs))

	signatureSeen := make(map[string]bool, len(m.ParticipantListTapScriptSigs))
	numSigs := -1
	// TODO: numSigs just the number of inputs signed, not the number of sigs, we need to go through the sigs and count the number of sigs by map
	for address, listOfSigs := range m.ParticipantListTapScriptSigs {
		if numSigs == -1 {
			numSigs = listOfSigs.Size()
		}
		if numSigs != listOfSigs.Size() {
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

		for _, sig := range listOfSigs.Inner {
			clog.Magenta("validate m.TapScriptSigs, sig: ", sig)
			if err := sig.ValidateBasic(); err != nil {
				return err
			}
		}
	}

	return nil
}

// GetSignature returns the ECDSA signature of the given participant
func (m PsbtMultiSig) GetListOfTapScriptSigsMap(p sdk.ValAddress) ([]*exported.TapScriptSigsMap, bool) {
	sigs, ok := m.ParticipantListTapScriptSigs[p.String()]
	if !ok {
		return nil, false
	}

	return sigs.Inner, true
}

// GetParticipants returns the participants of the given multi sig
func (m PsbtMultiSig) GetParticipants() []sdk.ValAddress {
	return multisigTypes.SortAddresses(
		slices.Map(maps.Keys(m.ParticipantListTapScriptSigs), func(a string) sdk.ValAddress { return funcs.Must(sdk.ValAddressFromBech32(a)) }),
	)
}

// verify tap script sig

// func verifyTapScriptSigs(sigs *exported.TapScriptSigList, pubKeyHex string, psbt Psbt) error {
//     // Parse the public key
//     pubKeyBytes, err := hex.DecodeString(pubKeyHex)
//     if err != nil {
//         return fmt.Errorf("invalid public key hex: %w", err)
//     }

//     pubKey, err := schnorr.ParsePubKey(pubKeyBytes)
//     if err != nil {
//         return fmt.Errorf("invalid public key: %w", err)
//     }

//     // Verify each signature
//     for i, sig := range sigs.TapScriptSigs {
//         if err := verifyTapScriptSig(sig, pubKey, psbt, i); err != nil {
//             return fmt.Errorf("invalid signature at index %d: %w", i, err)
//         }
//     }

//     return nil
// }

// func verifyTapScriptSig(sig *exported.TapScriptSig, pubKey *btcec.PublicKey, psbt Psbt, inputIndex int) error {
//     // Parse the signature
//     schnorrSig, err := schnorr.ParseSignature(sig.Signature)
//     if err != nil {
//         return fmt.Errorf("invalid schnorr signature: %w", err)
//     }

//     // Parse leaf hash
//     leafHash, err := txscript.NewTapLeaf(sig.LeafHash)
//     if err != nil {
//         return fmt.Errorf("invalid leaf hash: %w", err)
//     }

//     // Compute sighash for verification
//     sighash, err := psbt.ComputeTaprootSighash(inputIndex, leafHash)
//     if err != nil {
//         return fmt.Errorf("failed to compute sighash: %w", err)
//     }

//     // Verify the signature
//     if !schnorrSig.Verify(sighash, pubKey) {
//         return fmt.Errorf("signature verification failed")
//     }

//     return nil
// }

// ComputeTaprootSighash computes the sighash for taproot signature verification
// func (p Psbt) ComputeTaprootSighash(inputIndex int, leafHash txscript.TapLeaf) ([]byte, error) {
//     // Parse the PSBT
//     btcPsbt, err := p.ToBtcPsbt()
//     if err != nil {
//         return nil, fmt.Errorf("failed to parse PSBT: %w", err)
//     }

//     // Get the input
//     if inputIndex >= len(btcPsbt.Inputs) {
//         return nil, fmt.Errorf("input index out of range")
//     }
//     input := btcPsbt.Inputs[inputIndex]

//     // Compute sighash
//     sighash, err := input.TaprootSighash(
//         btcPsbt.UnsignedTx,
//         txscript.NewTxSigHashes(btcPsbt.UnsignedTx),
//         leafHash,
//     )
//     if err != nil {
//         return nil, fmt.Errorf("failed to compute sighash: %w", err)
//     }

//     return sighash, nil
// }
