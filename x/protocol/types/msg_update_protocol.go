package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewUpdateProtocolRequest(sender sdk.AccAddress, name string) *UpdateProtocolRequest {
	return &UpdateProtocolRequest{
		Sender: sender,
		Name:   name,
	}
}

// Route returns the route for this message
func (m UpdateProtocolRequest) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m UpdateProtocolRequest) Type() string {
	return "UpdateProtocol"
}

// ValidateBasic executes a stateless message validation
func (m UpdateProtocolRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	return nil
}

// GetSignBytes returns the message bytes that need to be signed
func (m UpdateProtocolRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the set of signers for this message
func (m UpdateProtocolRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
