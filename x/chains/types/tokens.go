package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// ERC20Token represents an ERC20 token and its respective state
type ERC20Token struct {
	metadata ERC20TokenMetadata
	setMeta  func(meta ERC20TokenMetadata)
}

// CreateERC20Token returns an ERC20Token struct
func CreateERC20Token(setter func(meta ERC20TokenMetadata), meta ERC20TokenMetadata) ERC20Token {
	token := ERC20Token{
		metadata: meta,
		setMeta:  setter,
	}

	return token
}

// GetAsset returns the asset name
func (t ERC20Token) GetAsset() string {
	return t.metadata.Asset
}

// GetTxID returns the tx ID set with StartConfirmation
func (t ERC20Token) GetTxID() Hash {
	return t.metadata.TxHash
}

// GetDetails returns the details of the token
func (t ERC20Token) GetDetails() TokenDetails {
	return t.metadata.Details
}

// Is returns true if the given status matches the token's status
func (t ERC20Token) Is(status Status) bool {
	// this special case check is needed, because 0 & x == 0 is true for any x
	if status == NonExistent {
		return t.metadata.Status == NonExistent
	}
	return status&t.metadata.Status == status
}

// IsExternal returns true if the given token is external; false otherwise
func (t ERC20Token) IsExternal() bool {
	return t.metadata.IsExternal
}

// GetBurnerCode returns the version of the burner the token is deployed with
func (t ERC20Token) GetBurnerCode() []byte {
	return t.metadata.BurnerCode
}

// GetBurnerCodeHash returns the version of the burner the token is deployed with if it exists
func (t ERC20Token) GetBurnerCodeHash() (Hash, bool) {
	if t.metadata.BurnerCode == nil {
		return Hash{}, false
	}

	return Hash(crypto.Keccak256Hash(t.metadata.BurnerCode)), true
}

// CreateDeployCommand returns a token deployment command for the token
func (t *ERC20Token) CreateDeployCommand(keyID multisig.KeyID, dailyMintLimit sdk.Uint) (Command, error) {
	switch {
	case t.Is(NonExistent):
		return Command{}, fmt.Errorf("token %s non-existent", t.GetAsset())
	case t.Is(Confirmed):
		return Command{}, fmt.Errorf("token %s already confirmed", t.GetAsset())
	}
	if err := keyID.ValidateBasic(); err != nil {
		return Command{}, err
	}

	if t.IsExternal() {
		return NewDeployTokenCommand(
			t.metadata.ChainID,
			keyID,
			t.GetAsset(),
			t.metadata.Details,
			t.GetAddress(),
			dailyMintLimit,
		), nil
	}

	return NewDeployTokenCommand(
		t.metadata.ChainID,
		keyID,
		t.GetAsset(),
		t.metadata.Details,
		ZeroAddress,
		dailyMintLimit,
	), nil
}

// CreateMintCommand returns a mint deployment command for the token
func (t *ERC20Token) CreateMintCommand(keyID multisig.KeyID, transfer nexus.CrossChainTransfer) (Command, error) {
	if !t.Is(Confirmed) {
		return Command{}, fmt.Errorf("token %s not confirmed (current status: %s)",
			t.metadata.Asset, t.metadata.Status.String())
	}
	if err := keyID.ValidateBasic(); err != nil {
		return Command{}, err
	}

	return NewMintTokenCommand(
		keyID,
		transfer.ID,
		t.metadata.Details.Symbol,
		common.HexToAddress(transfer.Recipient.Address),
		transfer.Asset.Amount.BigInt(),
	), nil
}

// GetAddress returns the token's address
func (t ERC20Token) GetAddress() Address {
	return t.metadata.TokenAddress

}

// RecordDeployment signals that the token confirmation is underway for the given tx ID
func (t *ERC20Token) RecordDeployment(txID Hash) error {
	switch {
	case t.Is(NonExistent):
		return fmt.Errorf("token %s non-existent", t.metadata.Asset)
	case t.Is(Confirmed):
		return fmt.Errorf("token %s already confirmed", t.metadata.Asset)
	}

	t.metadata.TxHash = txID
	t.metadata.Status |= Pending
	t.setMeta(t.metadata)

	return nil
}

// RejectDeployment reverts the token state back to Initialized
func (t *ERC20Token) RejectDeployment() error {
	switch {
	case t.Is(NonExistent):
		return fmt.Errorf("token %s non-existent", t.metadata.Asset)
	case !t.Is(Pending):
		return fmt.Errorf("token %s not waiting confirmation (current status: %s)", t.metadata.Asset, t.metadata.Status.String())
	}

	t.metadata.Status = Initialized
	t.metadata.TxHash = Hash{}
	t.setMeta(t.metadata)
	return nil
}

// ConfirmDeployment signals that the token was successfully confirmed
func (t *ERC20Token) ConfirmDeployment() error {
	switch {
	case t.Is(NonExistent):
		return fmt.Errorf("token %s non-existent", t.metadata.Asset)
	case !t.Is(Pending):
		return fmt.Errorf("token %s not waiting confirmation (current status: %s)", t.metadata.Asset, t.metadata.Status.String())
	}

	t.metadata.Status = Confirmed
	t.setMeta(t.metadata)

	return nil
}

// NilToken is a nil erc20 token
var NilToken = ERC20Token{}
