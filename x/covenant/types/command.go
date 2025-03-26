package types

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	fmt "fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	sdk "github.com/cosmos/cosmos-sdk/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

type Command struct {
	metadata CommandMetadata
	setter   func(batch CommandMetadata)
}

// NonExistentCommand can be used to represent a non-existent command
var NonExistentCommand = NewCommand(CommandMetadata{}, func(CommandMetadata) {})

// NewCommand returns a new command batch struct
func NewCommand(metadata CommandMetadata, setter func(batch CommandMetadata)) Command {
	return Command{
		metadata: metadata,
		setter:   setter,
	}
}

// GetStatus returns the batch's status
func (b Command) GetStatus() CommandStatus {
	return b.metadata.Status
}

// GetData returns the batch's data
func (b Command) GetData() []byte {
	return b.metadata.Data
}

// GetID returns the batch ID
func (b Command) GetID() []byte {
	return b.metadata.ID

}

// GetKeyID returns the batch's key ID
func (b Command) GetKeyID() multisig.KeyID {
	return b.metadata.KeyID

}

// GetSigHash returns the batch's key ID
func (b Command) GetSigHash() exported.Hash {
	return b.metadata.SigHash

}

// GetSignature returns the batch's signature
func (b Command) GetSignature() utils.ValidatedProtoMarshaler {
	if b.metadata.Signature == nil {
		return nil
	}

	return b.metadata.Signature.GetCachedValue().(utils.ValidatedProtoMarshaler)
}

// Is returns true if batched commands is in the given status; false otherwise
func (b Command) Is(status CommandStatus) bool {
	return b.metadata.Status == status
}

// SetStatus sets the status for the batch, returning true if the status was updated
func (b *Command) SetStatus(status CommandStatus) bool {
	if b.metadata.Status != CommandNonExistent && b.metadata.Status != CommandSigned {
		b.metadata.Status = status
		b.setter(b.metadata)
		return true
	}

	return false
}

// SetSigned sets the signature and signed status for the batch
func (b *Command) SetSigned(signature utils.ValidatedProtoMarshaler) error {
	if b.metadata.Status != CommandSigning {
		return fmt.Errorf("command %s is not being signed", hex.EncodeToString(b.GetID()))
	}

	b.metadata.Status = CommandSigned
	sig := funcs.Must(codectypes.NewAnyWithValue(signature))
	b.metadata.Signature = sig

	b.setter(b.metadata)

	return nil
}

// NewCommandMetadata assembles a CommandMetadata struct from the provided arguments
func NewCommandMetadata(blockHeight int64, chainID sdk.Int, keyID multisig.KeyID, data []byte) (CommandMetadata, error) {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(blockHeight))

	return CommandMetadata{
		ID:      crypto.Keccak256(bz, data),
		Data:    data,
		SigHash: exported.Hash(chainsTypes.GetSignHash(data)),
		Status:  CommandSigning,
		KeyID:   keyID,
	}, nil
}

// ValidateBasic returns an error if the CommandMetadata is not valid
func (m CommandMetadata) ValidateBasic() error {
	switch m.Status {
	case CommandNonExistent:
		return errors.New("command does not exist")
	case CommandSigning, CommandAborted:
		if m.Signature != nil {
			return errors.New("unsigned command must not have a signature")
		}
	case CommandSigned:
		if m.Signature == nil {
			return errors.New("signed command must have a valid signature")
		}

		if err := m.Signature.GetCachedValue().(utils.ValidatedProtoMarshaler).ValidateBasic(); err != nil {
			return err
		}
	}

	if len(m.ID) != 32 {
		return errors.New("command ID must be of length 32")
	}

	if len(m.Data) == 0 {
		return errors.New("batch data must not be empty")
	}

	if m.SigHash.IsZero() {
		return errors.New("batch data hash must not be empty")
	}

	if err := m.KeyID.ValidateBasic(); err != nil {
		return err
	}

	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage
func (m CommandMetadata) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var data codec.ProtoMarshaler

	return unpacker.UnpackAny(m.Signature, &data)
}
