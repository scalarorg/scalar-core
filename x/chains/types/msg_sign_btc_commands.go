package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/scalarorg/scalar-core/utils"
	covenant "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewSignBtcCommandsRequest creates a message of type SignBtcCommandsRequest
func NewSignBtcCommandsRequest(sender sdk.AccAddress, chain string) *SignBtcCommandsRequest {
	return &SignBtcCommandsRequest{
		Sender: sender,
		Chain:  nexus.ChainName(utils.NormalizeString(chain)),
	}
}

// Route implements sdk.Msg
func (m SignBtcCommandsRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (m SignBtcCommandsRequest) Type() string {
	return "SignBtcCommands"
}

// ValidateBasic implements sdk.Msg
func (m SignBtcCommandsRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if err := m.Chain.Validate(); err != nil {
		return sdkerrors.Wrap(err, "invalid chain")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (m SignBtcCommandsRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (m SignBtcCommandsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}

func NewSignPsbtCommandRequest(sender sdk.AccAddress, chain string, psbt covenant.Psbt) *SignPsbtCommandRequest {
	return &SignPsbtCommandRequest{
		Sender: sender,
		Chain:  nexus.ChainName(utils.NormalizeString(chain)),
		Psbt:   psbt,
	}
}

// Route implements sdk.Msg
func (m SignPsbtCommandRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (m SignPsbtCommandRequest) Type() string {
	return "SignPsbtCommand"
}

// ValidateBasic implements sdk.Msg
func (m SignPsbtCommandRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if err := m.Chain.Validate(); err != nil {
		return sdkerrors.Wrap(err, "invalid chain")
	}
	//Todo:
	return nil
}

// GetSignBytes implements sdk.Msg
func (m SignPsbtCommandRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (m SignPsbtCommandRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
