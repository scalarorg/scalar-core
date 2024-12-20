package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewAddSupportedChainRequest(sender sdk.AccAddress, chain *SupportedChain) *AddSupportedChainRequest {
	return &AddSupportedChainRequest{
		Sender: sender,
		Chain:  chain,
	}
}

// Route returns the route for this message
func (m AddSupportedChainRequest) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m AddSupportedChainRequest) Type() string {
	return "AddSupportedChain"
}

// ValidateBasic executes a stateless message validation
func (m AddSupportedChainRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	return nil
}

// GetSignBytes returns the message bytes that need to be signed
func (m AddSupportedChainRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the set of signers for this message
func (m AddSupportedChainRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
