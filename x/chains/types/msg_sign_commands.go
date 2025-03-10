package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/scalarorg/scalar-core/utils"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewSignCommandsRequest creates a message of type SignCommandsRequest
func NewSignCommandsRequest(sender sdk.AccAddress, chain string) *SignCommandsRequest {
	return &SignCommandsRequest{
		Sender: sender,
		Chain:  nexus.ChainName(utils.NormalizeString(chain)),
	}
}

// Route implements sdk.Msg
func (m SignCommandsRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (m SignCommandsRequest) Type() string {
	return "SignCommands"
}

// ValidateBasic implements sdk.Msg
func (m SignCommandsRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if err := m.Chain.Validate(); err != nil {
		return sdkerrors.Wrap(err, "invalid chain")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (m SignCommandsRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (m SignCommandsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
