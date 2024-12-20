package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewAddChainRequest is the constructor for NewAddChainRequest
func NewRemoveCustodiaFromGroupRequest(sender sdk.AccAddress, name string) *RemoveCustodianFromGroupRequest {
	return &RemoveCustodianFromGroupRequest{
		Sender: sender,
	}
}

// Route returns the route for this message
func (m RemoveCustodianFromGroupRequest) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m RemoveCustodianFromGroupRequest) Type() string {
	return "CustodianToGroup"
}

// ValidateBasic executes a stateless message validation
func (m RemoveCustodianFromGroupRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	return nil
}

// GetSignBytes returns the message bytes that need to be signed
func (m RemoveCustodianFromGroupRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the set of signers for this message
func (m RemoveCustodianFromGroupRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
