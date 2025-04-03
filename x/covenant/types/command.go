package types

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	fmt "fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	covExported "github.com/scalarorg/scalar-core/x/covenant/exported"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var (
	stringType       = funcs.Must(abi.NewType("string", "string", nil))
	addressType      = funcs.Must(abi.NewType("address", "address", nil))
	addressesType    = funcs.Must(abi.NewType("address[]", "address[]", nil))
	bytes32Type      = funcs.Must(abi.NewType("bytes32", "bytes32", nil))
	uint8Type        = funcs.Must(abi.NewType("uint8", "uint8", nil))
	uint256Type      = funcs.Must(abi.NewType("uint256", "uint256", nil))
	uint256ArrayType = funcs.Must(abi.NewType("uint256[]", "uint256[]", nil))
	bytes32ArrayType = funcs.Must(abi.NewType("bytes32[]", "bytes32[]", nil))
	stringArrayType  = funcs.Must(abi.NewType("string[]", "string[]", nil))
	bytesArrayType   = funcs.Must(abi.NewType("bytes[]", "bytes[]", nil))
)

const (
	switchPhaseMaxGasCost = 100000
)

type StandaloneCommand struct {
	metadata StandaloneCommandMetadata
	setter   func(batch StandaloneCommandMetadata)
}

// NonExistentStandaloneStandaloneCommand can be used to represent a non-existent command
var NonExistentStandaloneCommand = NewStandaloneCommand(StandaloneCommandMetadata{}, func(StandaloneCommandMetadata) {})

// NewStandaloneCommand returns a new command batch struct
func NewStandaloneCommand(metadata StandaloneCommandMetadata, setter func(batch StandaloneCommandMetadata)) StandaloneCommand {
	return StandaloneCommand{
		metadata: metadata,
		setter:   setter,
	}
}

// GetStatus returns the batch's status
func (b StandaloneCommand) GetStatus() StandaloneCommandStatus {
	return b.metadata.Status
}

// GetData returns the batch's data
func (b StandaloneCommand) GetData() []byte {
	return b.metadata.Data
}

// GetID returns the batch ID
func (b StandaloneCommand) GetID() []byte {
	return b.metadata.ID

}

// GetKeyID returns the batch's key ID
func (b StandaloneCommand) GetKeyID() multisig.KeyID {
	return b.metadata.KeyID

}

// GetSigHash returns the batch's key ID
func (b StandaloneCommand) GetSigHash() exported.Hash {
	return b.metadata.SigHash

}

// GetSignature returns the batch's signature
func (b StandaloneCommand) GetSignature() utils.ValidatedProtoMarshaler {
	if b.metadata.Signature == nil {
		return nil
	}

	return b.metadata.Signature.GetCachedValue().(utils.ValidatedProtoMarshaler)
}

// Is returns true if batched commands is in the given status; false otherwise
func (b StandaloneCommand) Is(status StandaloneCommandStatus) bool {
	return b.metadata.Status == status
}

// SetStatus sets the status for the batch, returning true if the status was updated
func (b *StandaloneCommand) SetStatus(status StandaloneCommandStatus) bool {
	if b.metadata.Status != StandaloneCommandStatusNonExistent && b.metadata.Status != StandaloneCommandStatusSigned {
		b.metadata.Status = status
		b.setter(b.metadata)
		return true
	}

	return false
}

// SetSigned sets the signature and signed status for the batch
func (b *StandaloneCommand) SetSigned(signature utils.ValidatedProtoMarshaler) error {
	if b.metadata.Status != StandaloneCommandStatusSigning {
		return fmt.Errorf("command %s is not being signed", hex.EncodeToString(b.GetID()))
	}

	b.metadata.Status = StandaloneCommandStatusSigned
	sig := funcs.Must(codectypes.NewAnyWithValue(signature))
	b.metadata.Signature = sig

	b.setter(b.metadata)

	return nil
}

func (b *StandaloneCommand) GetCommandID() CommandID {
	return b.metadata.CommandID
}

// NewStandaloneCommandMetadata assembles a StandaloneCommandMetadata struct from the provided arguments
func NewStandaloneCommandMetadata(
	blockHeight int64,
	keyID multisig.KeyID,
	chainID sdk.Int,
	commandID CommandID,
	commandType chainsTypes.CommandType,
	commandParam []byte,
) (StandaloneCommandMetadata, error) {
	data, err := packArguments(chainID, commandID, commandType, commandParam)
	if err != nil {
		return StandaloneCommandMetadata{}, err
	}

	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(blockHeight))

	return StandaloneCommandMetadata{
		ID:        crypto.Keccak256(bz, data),
		CommandID: commandID,
		Data:      data,
		SigHash:   exported.Hash(chainsTypes.GetSignHash(data)),
		Status:    StandaloneCommandStatusSigning,
		KeyID:     keyID,
	}, nil
}

// ValidateBasic returns an error if the StandaloneCommandMetadata is not valid
func (m StandaloneCommandMetadata) ValidateBasic() error {
	switch m.Status {
	case StandaloneCommandStatusNonExistent:
		return errors.New("command does not exist")
	case StandaloneCommandStatusSigning, StandaloneCommandStatusAborted:
		if m.Signature != nil {
			return errors.New("unsigned command must not have a signature")
		}
	case StandaloneCommandStatusSigned:
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
func (m StandaloneCommandMetadata) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var data codec.ProtoMarshaler

	return unpacker.UnpackAny(m.Signature, &data)
}

// NewStandaloneCommandSigned returns a new StandaloneCommandSigned instance
func NewStandaloneCommandSigned(chain nexus.ChainName, ID []byte) *StandaloneCommandSigned {
	return &StandaloneCommandSigned{Chain: chain, CommandID: ID}
}

// NewStandaloneCommandAborted returns a new StandaloneCommandAborted instance
func NewStandaloneCommandAborted(chain nexus.ChainName, ID []byte) *StandaloneCommandAborted {
	return &StandaloneCommandAborted{Chain: chain, CommandID: ID}
}

const commandIDSize = 32

// CommandID represents the unique command identifier
type CommandID [commandIDSize]byte

func NewCommandID(id []byte) CommandID {
	var commandID CommandID
	copy(commandID[:], id)
	return commandID
}

// Hex returns the hex representation of command ID
func (c CommandID) Hex() string {
	return hex.EncodeToString(c[:])
}

// Size implements codec.ProtoMarshaler
func (c CommandID) Size() int {
	return commandIDSize
}

// Marshal implements codec.ProtoMarshaler
func (c CommandID) Marshal() ([]byte, error) {
	return c[:], nil
}

// MarshalTo implements codec.ProtoMarshaler
func (c CommandID) MarshalTo(data []byte) (n int, err error) {
	bytesCopied := copy(data, c[:])
	if bytesCopied != commandIDSize {
		return 0, fmt.Errorf("expected data size to be %d, actual %d", commandIDSize, len(data))
	}

	return commandIDSize, nil
}

// Unmarshal implements codec.ProtoMarshaler
func (c *CommandID) Unmarshal(data []byte) error {
	bytesCopied := copy(c[:], data)
	if bytesCopied != commandIDSize {
		return fmt.Errorf("expected data size to be %d, actual %d", commandIDSize, len(data))
	}

	return c.ValidateBasic()
}

// ValidateBasic returns an error if the given command ID is invalid
func (c CommandID) ValidateBasic() error {
	return nil
}

func packArguments(chainID sdk.Int, id CommandID, typ chainsTypes.CommandType, param []byte) ([]byte, error) {
	arguments := abi.Arguments{{Type: uint256Type}, {Type: bytes32ArrayType}, {Type: stringArrayType}, {Type: bytesArrayType}}
	result, err := arguments.Pack(
		chainID.BigInt(),
		[]CommandID{id},
		[]string{typ.String()},
		[][]byte{param},
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewSwitchPhaseCommandWithExpiredSessioin(
	session *ExpiredEvmSession,
	chainID sdk.Int,
	keyID multisig.KeyID,
	newPhase covExported.Phase,
) chainsTypes.Command {
	chainbz := make([]byte, 8)
	binary.BigEndian.PutUint64(chainbz, uint64(chainID.Uint64()))

	id := append(session.CustodianGroupUID.Bytes(), chainbz...)
	id = append(id, byte(session.Sequence))
	id = append(id, byte(session.CurrentPhase))

	cmd := chainsTypes.Command{
		ID:         chainsTypes.NewCommandID(id, chainID),
		Type:       chainsTypes.COMMAND_TYPE_REDEEM_TOKEN,
		Params:     CreateSwitchPhasePayload(session.CustodianGroupUID, newPhase),
		Payload:    []byte{},
		KeyID:      keyID,
		MaxGasCost: uint32(switchPhaseMaxGasCost),
	}

	return cmd
}

func CreateSwitchPhasePayload(CustodianGroupUID exported.Hash, newPhase covExported.Phase) []byte {
	var cusID [32]byte
	copy(cusID[:], CustodianGroupUID.Bytes())
	payload := funcs.Must(chainsTypes.SwitchPhaseArguments.Pack(uint8(newPhase), cusID))
	return payload
}
