package exported

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/scalarorg/scalar-core/utils"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
)

//go:generate moq -out ./mock/types.go -pkg mock . CovenantHandler Key MultiSig

// PsbtMultiSig provides an interface to work with the multi sig
type PsbtMultiSig interface {
	GetTapScriptSig(p sdk.ValAddress) (TapScriptSig, bool)
	GetPsbt() []byte
	GetKeyID() multisig.KeyID
	ValidateBasic() error
}

// CovenantHandler defines the interface for the requesting module to implement in
// order to handle the different results of signing session
type CovenantHandler interface {
	HandleCompleted(ctx sdk.Context, sig utils.ValidatedProtoMarshaler, moduleMetadata codec.ProtoMarshaler) error
	HandleFailed(ctx sdk.Context, moduleMetadata codec.ProtoMarshaler) error
}

// key id length range bounds dictated by tofnd
const (
	KeyIDLengthMin  = 4
	KeyIDLengthMax  = 256
	KeyXOnlyLength  = 32
	LeafHashLength  = 32
	SignatureLength = 64
)

type KeyXOnly [KeyXOnlyLength]byte

func (k KeyXOnly) ValidateBasic() error {
	if len(k) != KeyXOnlyLength {
		return fmt.Errorf("key x only length %d not in range [%d,%d]", len(k), KeyXOnlyLength, KeyXOnlyLength)
	}

	return nil
}

func (k KeyXOnly) Size() int {
	return KeyXOnlyLength
}

func (k KeyXOnly) MarshalTo(dAtA []byte) (int, error) {
	copy(dAtA, k[:])
	return KeyXOnlyLength, nil
}

func (k KeyXOnly) Unmarshal(dAtA []byte) error {
	copy(k[:], dAtA)
	return nil
}

type LeafHash [LeafHashLength]byte

func (l LeafHash) ValidateBasic() error {
	if len(l) != LeafHashLength {
		return fmt.Errorf("leaf hash length %d not in range [%d,%d]", len(l), LeafHashLength, LeafHashLength)
	}

	return nil
}

func (l LeafHash) Size() int {
	return LeafHashLength
}

func (l LeafHash) MarshalTo(dAtA []byte) (int, error) {
	copy(dAtA, l[:])
	return LeafHashLength, nil
}

func (l LeafHash) Unmarshal(dAtA []byte) error {
	copy(l[:], dAtA)
	return nil
}

type Signature [SignatureLength]byte

func (s Signature) ValidateBasic() error {
	if len(s) != SignatureLength {
		return fmt.Errorf("signature length %d not in range [%d,%d]", len(s), SignatureLength, SignatureLength)
	}

	return nil
}

func (s Signature) Size() int {
	return SignatureLength
}

func (s Signature) MarshalTo(dAtA []byte) (int, error) {
	copy(dAtA, s[:])
	return SignatureLength, nil
}

func (s Signature) Unmarshal(dAtA []byte) error {
	copy(s[:], dAtA)
	return nil
}

func (t TapScriptSig) ValidateBasic() error {
	err := t.KeyXOnly.ValidateBasic()
	if err != nil {
		return err
	}

	err = t.LeafHash.ValidateBasic()
	if err != nil {
		return err
	}

	err = t.Signature.ValidateBasic()
	if err != nil {
		return err
	}

	return nil
}
