package exported

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	utiltypes "github.com/scalarorg/bitcoin-vault/go-utils/types"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/slices"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
)

//go:generate moq -out ./mock/types.go -pkg mock . CovenantHandler Key MultiSig

// PsbtMultiSig provides an interface to work with the multi sig
type PsbtMultiSig interface {
	GetKeyID() multisig.KeyID
	ValidateBasic() error
	GetFinalizedTx() []byte
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

var EmptyKeyXOnly = KeyXOnly{}

func (k KeyXOnly) Bytes() [KeyXOnlyLength]byte {
	return k
}

func (k KeyXOnly) ValidateBasic() error {
	if k == EmptyKeyXOnly {
		return fmt.Errorf("key x only is empty")
	}
	return nil
}

func (k KeyXOnly) Size() int {
	return KeyXOnlyLength
}

func (k *KeyXOnly) MarshalTo(dAtA []byte) (int, error) {
	copy(dAtA, k[:])
	return KeyXOnlyLength, nil
}

func (k *KeyXOnly) Unmarshal(dAtA []byte) error {
	if len(dAtA) != KeyXOnlyLength {
		return fmt.Errorf("invalid data length: expected %d, got %d", KeyXOnlyLength, len(dAtA))
	}
	copy(k[:], dAtA)
	return nil
}

type LeafHash [LeafHashLength]byte

var EmptyLeafHash = LeafHash{}

func (l LeafHash) Bytes() [LeafHashLength]byte {
	return l
}

func (l LeafHash) ValidateBasic() error {
	if len(l) != LeafHashLength {
		return fmt.Errorf("leaf hash length %d not in range [%d,%d]", len(l), LeafHashLength, LeafHashLength)
	}

	if l == EmptyLeafHash {
		return fmt.Errorf("leaf hash is empty")
	}

	return nil
}

func (l LeafHash) Size() int {
	return LeafHashLength
}

func (l *LeafHash) MarshalTo(dAtA []byte) (int, error) {
	copy(dAtA, l[:])
	return LeafHashLength, nil
}

func (l *LeafHash) Unmarshal(dAtA []byte) error {
	if len(dAtA) != LeafHashLength {
		return fmt.Errorf("invalid data length: expected %d, got %d", LeafHashLength, len(dAtA))
	}
	copy(l[:], dAtA)
	return nil
}

type Signature [SignatureLength]byte

var EmptySignature = Signature{}

func (s Signature) Bytes() [SignatureLength]byte {
	return s
}

func (s Signature) ValidateBasic() error {
	if len(s) != SignatureLength {
		return fmt.Errorf("signature length %d not in range [%d,%d]", len(s), SignatureLength, SignatureLength)
	}

	if s == EmptySignature {
		return fmt.Errorf("signature is empty")
	}

	return nil
}

func (s Signature) Size() int {
	return SignatureLength
}

func (s *Signature) MarshalTo(dAtA []byte) (int, error) {
	copy(dAtA, s[:])
	return SignatureLength, nil
}

func (s *Signature) Unmarshal(dAtA []byte) error {
	if len(dAtA) != SignatureLength {
		return fmt.Errorf("invalid data length: expected %d, got %d", SignatureLength, len(dAtA))
	}
	copy(s[:], dAtA)
	return nil
}

func (t *TapScriptSig) ValidateBasic() error {
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

var EmptyTapScriptSigsList = TapScriptSigsList{}

var EmptyTapScriptSigsMap = TapScriptSigsMap{
	Inner: make(map[uint64]*TapScriptSigsList),
}

func (t *TapScriptSigsMap) ValidateBasic() error {
	for _, sig := range t.Inner {
		for _, sig := range sig.List {
			if err := sig.ValidateBasic(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *TapScriptSigsMap) ToRaw() utiltypes.TapScriptSigsMap {
	raw := make(utiltypes.TapScriptSigsMap)
	for inputIndex, tapScriptList := range t.Inner {
		raw[inputIndex] = []utiltypes.TapScriptSig{}
		for _, tapScriptSig := range tapScriptList.List {
			raw[inputIndex] = append(raw[inputIndex], utiltypes.TapScriptSig{
				KeyXOnly:  tapScriptSig.KeyXOnly.Bytes(),
				LeafHash:  tapScriptSig.LeafHash.Bytes(),
				Signature: tapScriptSig.Signature.Bytes(),
			})
		}
	}
	return raw
}

func (t *TapScriptSigsMap) FromRaw(raw utiltypes.TapScriptSigsMap) *TapScriptSigsMap {
	mapOfTapScriptSigs := make(map[uint64]*TapScriptSigsList)
	for inputIndex, tapScriptSigs := range raw {
		mapOfTapScriptSigs[inputIndex] = &TapScriptSigsList{
			List: slices.Map(tapScriptSigs, func(t utiltypes.TapScriptSig) *TapScriptSig {
				keyXOnly := KeyXOnly(t.KeyXOnly)
				signature := Signature(t.Signature)
				leafHash := LeafHash(t.LeafHash)
				return &TapScriptSig{
					KeyXOnly:  &keyXOnly,
					Signature: &signature,
					LeafHash:  &leafHash,
				}
			}),
		}
	}

	t.Inner = mapOfTapScriptSigs
	return t
}
