package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/x/protocol/exported"
)

func NewUpdateSupportedChainRequest(sender sdk.AccAddress, chainFamily string, chainId uint64, status exported.Status) *UpdateSupportedChainRequest {
	return &UpdateSupportedChainRequest{
		Sender:      sender,
		ChainFamily: chainFamily,
		ChainId:     chainId,
		Status:      status,
	}
}

// Route returns the route for this message
func (m UpdateSupportedChainRequest) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m UpdateSupportedChainRequest) Type() string {
	return "AddSupportedChain"
}

// ValidateBasic executes a stateless message validation
func (m UpdateSupportedChainRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	return nil
}

// GetSignBytes returns the message bytes that need to be signed
func (m UpdateSupportedChainRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the set of signers for this message
func (m UpdateSupportedChainRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
