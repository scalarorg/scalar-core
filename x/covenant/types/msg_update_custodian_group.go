package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewAddChainRequest is the constructor for NewAddChainRequest
func NewUpdateCustodianGroupRequest(sender sdk.AccAddress, name string) *UpdateCustodianGroupRequest {
	return &UpdateCustodianGroupRequest{
		Sender: sender,
	}
}

// Route returns the route for this message
func (m UpdateCustodianGroupRequest) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m UpdateCustodianGroupRequest) Type() string {
	return "UpdateCustodianGroup"
}

// ValidateBasic executes a stateless message validation
func (m UpdateCustodianGroupRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	return nil
}

// GetSignBytes returns the message bytes that need to be signed
func (m UpdateCustodianGroupRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the set of signers for this message
func (m UpdateCustodianGroupRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
