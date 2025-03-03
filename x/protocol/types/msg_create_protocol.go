package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	types "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/scalar-core/x/protocol/exported"
)

func NewCreateProtocolRequest(sender sdk.AccAddress, name string, bitcoinPubkey []byte, tag string, attributes *exported.ProtocolAttributes, custodianGroupUid string, avatar []byte, asset *types.Asset) *CreateProtocolRequest {
	return &CreateProtocolRequest{
		Sender:            sender,
		Name:              name,
		BitcoinPubkey:     bitcoinPubkey,
		Tag:               tag,
		Attributes:        attributes,
		CustodianGroupUid: custodianGroupUid,
		Avatar:            avatar,
		Asset:             asset,
	}
}

// Route returns the route for this message
func (m CreateProtocolRequest) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m CreateProtocolRequest) Type() string {
	return "CreateProtocol"
}

// ValidateBasic executes a stateless message validation
func (m *CreateProtocolRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if len(m.BitcoinPubkey) != 33 {
		return fmt.Errorf("bitcoin pubkey must be 33 bytes")
	}

	if len(m.Name) > 64 {
		return fmt.Errorf("name must be less than 64 bytes")
	}

	if len(m.Tag) > 0 && len(m.Tag) > 64 {
		return fmt.Errorf("tag must be less than 64 bytes")
	}

	if len(m.CustodianGroupUid) > 64 {
		return fmt.Errorf("custodian group uid must be less than 64 bytes")
	}

	// TODO: validate asset name and chain is supported

	return nil
}

// GetSignBytes returns the message bytes that need to be signed
func (m CreateProtocolRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the set of signers for this message
func (m CreateProtocolRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
