package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/scalarorg/scalar-core/utils"
	covenant "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewSignBTCCommandsRequest creates a message of type SignBTCCommandsRequest
func NewSignBTCCommandsRequest(sender sdk.AccAddress, chain string, psbt covenant.Psbt) *SignBTCCommandsRequest {
	return &SignBTCCommandsRequest{
		Sender: sender,
		Chain:  nexus.ChainName(utils.NormalizeString(chain)),
		Psbt:   psbt,
	}
}

// Route implements sdk.Msg
func (m SignBTCCommandsRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (m SignBTCCommandsRequest) Type() string {
	return "SignBTCCommands"
}

// ValidateBasic implements sdk.Msg
func (m SignBTCCommandsRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if err := m.Chain.Validate(); err != nil {
		return sdkerrors.Wrap(err, "invalid chain")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (m SignBTCCommandsRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (m SignBTCCommandsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
