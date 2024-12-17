package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	fmt "fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	utils "github.com/scalarorg/scalar-core/utils"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

func IsBTCChain(chain nexus.Chain) bool {
	return chain.Module == ModuleName
}

type VaultTag [6]byte

func (VaultTag) Size() int {
	return 6
}

func (v VaultTag) MarshalTo(data []byte) (int, error) {
	return copy(data, v[:]), nil
}

func (v VaultTag) Unmarshal(data []byte) error {
	copy(v[:], data)
	return nil
}

func TagFromAscii(str string) VaultTag {
	var tag VaultTag
	copy(tag[:], []byte(str))
	return tag
}

func (v VaultTag) HexStr() string {
	return hex.EncodeToString(v[:])
}

type VaultVersion [1]byte

func VersionFromInt(version int) VaultVersion {
	var v VaultVersion
	v[0] = byte(version)
	return v
}

func (VaultVersion) Size() int {
	return 1
}

func (v VaultVersion) MarshalTo(data []byte) (int, error) {
	return copy(data, v[:]), nil
}

func (v VaultVersion) Unmarshal(data []byte) error {
	copy(v[:], data)
	return nil
}

type Hash chainhash.Hash

var ZeroHash = Hash{}

func FromChainHash(hash chainhash.Hash) Hash {
	return Hash(hash)
}

func (h Hash) Into() chainhash.Hash {
	return chainhash.Hash(h)
}

func (h Hash) IntoRef() *chainhash.Hash {
	return (*chainhash.Hash)(&h)
}

func (h Hash) Bytes() []byte {
	return h[:]
}

func HashFromBytes(data []byte) (Hash, error) {
	if len(data) != chainhash.HashSize {
		return Hash{}, fmt.Errorf("invalid hash length")
	}
	return Hash(chainhash.Hash(data)), nil
}

func (h Hash) IsZero() bool {
	return bytes.Equal(h.Bytes(), ZeroHash.Bytes())
}

func (h Hash) Size() int {
	return chainhash.HashSize
}

func (h Hash) Marshal() ([]byte, error) {
	return h[:], nil
}

func (h Hash) MarshalTo(data []byte) (int, error) {
	bytesCopied := copy(data, h[:])
	if bytesCopied != chainhash.HashSize {
		return 0, fmt.Errorf("failed to marshal TxHash: expected %d bytes, got %d", chainhash.HashSize, bytesCopied)
	}
	return bytesCopied, nil
}

func (h *Hash) Unmarshal(data []byte) error {
	hash, err := chainhash.NewHash(data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal TxHash: %w", err)
	}
	*h = Hash(*hash)
	return nil
}

func HashFromHexStr(hexStr string) (*Hash, error) {
	hash, err := chainhash.NewHashFromStr(hexStr)
	if err != nil {
		return nil, err
	}
	return (*Hash)(hash), nil
}

func (h Hash) HexStr() string {
	return (*chainhash.Hash)(&h).String()
}

func (nk NetworkKind) Validate() error {
	if nk != Mainnet && nk != Testnet {
		return fmt.Errorf("invalid network kind: %d", nk)
	}
	return nil
}

func (nk NetworkKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(nk.String())
}

func (nk *NetworkKind) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return nk.FromString(s)
}

func (nk *NetworkKind) FromString(s string) error {
	num, ok := NetworkKind_value[s]
	if !ok {
		return fmt.Errorf("invalid network kind: %s", s)
	}
	*nk = NetworkKind(num)
	return nil
}

func (nk *NetworkKind) UnmarshalText(text []byte) error {
	return nk.FromString(string(text))
}

const commandIDSize = 32

// CommandID represents the unique command identifier
type CommandID [commandIDSize]byte

// NewCommandID is the constructor for CommandID
func NewCommandID(data []byte, chainID sdk.Int) CommandID {
	var commandID CommandID
	copy(commandID[:], crypto.Keccak256(append(data, chainID.BigInt().Bytes()...))[:commandIDSize])

	return commandID
}

// CommandIDFromTransferID converts a TransferID into a CommandID
func CommandIDFromTransferID(id nexus.TransferID) CommandID {
	var commandID CommandID
	idBz := id.Bytes()

	copy(commandID[:], common.LeftPadBytes(idBz[:], commandIDSize))

	return commandID
}

// HexToCommandID decodes a hex representation of a CommandID
func HexToCommandID(id string) (CommandID, error) {
	bz, err := utils.HexDecode(id)
	if err != nil {
		return CommandID{}, err
	}

	var commandID CommandID
	copy(commandID[:], bz)

	return commandID, commandID.ValidateBasic()
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

func (m *StakingTx) ValidateBasic() error {
	if err := sdk.ValidateDenom(m.Asset); err != nil {
		return sdkerrors.Wrap(err, "invalid asset")
	}

	if err := m.DestinationChain.Validate(); err != nil {
		return sdkerrors.Wrap(err, "invalid destination chain")
	}

	if m.Amount.IsZero() {
		return fmt.Errorf("amount must be >0")
	}

	return nil
}

func (m CommandBatchMetadata) ValidateBasic() error {
	switch m.Status {
	case BatchNonExistent:
		return errors.New("batch does not exist")
	case BatchSigning, BatchAborted:
		if m.Signature != nil {
			return errors.New("unsigned batch must not have a signature")
		}
	case BatchSigned:
		if m.Signature == nil {
			return errors.New("signed batch must have a valid signature")
		}

		if err := m.Signature.GetCachedValue().(utils.ValidatedProtoMarshaler).ValidateBasic(); err != nil {
			return err
		}
	}

	if len(m.ID) != 32 {
		return errors.New("batch ID must be of length 32")
	}

	if len(m.CommandIDs) == 0 {
		return errors.New("command IDs must not be empty")
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

	if len(m.PrevBatchedCommandsID) != 0 && len(m.PrevBatchedCommandsID) != 32 {
		return errors.New("previous batch ID must either be nil or of length 32")
	}

	return nil
}

// EventID ensures a correctly formatted event ID
type EventID string

// NewEventID returns a new event ID
func NewEventID(txID Hash, index uint64) EventID {
	return EventID(fmt.Sprintf("%s-%d", txID.HexStr(), index))
}

// Validate returns an error, if the event ID is not in format of txID-index
func (id EventID) Validate() error {
	if err := utils.ValidateString(string(id)); err != nil {
		return err
	}

	arr := strings.Split(string(id), "-")
	if len(arr) != 2 {
		return fmt.Errorf("event ID should be in foramt of txID-index")
	}

	bz, err := hexutil.Decode(arr[0])
	if err != nil {
		return sdkerrors.Wrap(err, "invalid tx hash hex encoding")
	}

	if len(bz) != common.HashLength {
		return fmt.Errorf("invalid tx hash length")
	}

	_, err = strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid index")
	}

	return nil
}

// GetID returns an unique ID for the event
func (m Event) GetID() EventID {
	return NewEventID(m.TxID, m.Index)
}

// ValidateBasic returns an error if the event is invalid
func (m Event) ValidateBasic() error {
	if err := m.Chain.Validate(); err != nil {
		return sdkerrors.Wrap(err, "invalid source chain")
	}

	if m.TxID.IsZero() {
		return fmt.Errorf("invalid tx id")
	}

	// TODO: validate event type

	return nil
}

// NewVoteEvents is the constructor for vote events
func NewVoteEvents(chain nexus.ChainName, events ...Event) *VoteEvents {
	return &VoteEvents{
		Chain:  chain,
		Events: events,
	}
}

// ValidateBasic does stateless validation of the object
func (m VoteEvents) ValidateBasic() error {
	if err := m.Chain.Validate(); err != nil {
		return err
	}

	for _, event := range m.Events {
		if err := event.ValidateBasic(); err != nil {
			return err
		}

		if event.Chain != m.Chain {
			return fmt.Errorf("events are not from the same source chain")
		}
	}

	return nil
}

func (m *EventStakingTx) ValidateBasic() error {
	// TODO: validate event
	return nil
}

func getType(val interface{}) string {
	t := reflect.TypeOf(val)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

// GetEventType returns the type for the event
func (m Event) GetEventType() string {
	return getType(m.GetEvent())
}
