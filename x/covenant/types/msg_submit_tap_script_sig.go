package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
)

var _ sdk.Msg = &SubmitTapScriptSigRequest{}

// NewSubmitTapScriptSigRequest constructor for SubmitTapScriptSigRequest
func NewSubmitTapScriptSigRequest(sender sdk.AccAddress, sigID uint64, tapScriptSig *exported.TapScriptSig) *SubmitTapScriptSigRequest {
	return &SubmitTapScriptSigRequest{
		Sender:       sender,
		SigID:        sigID,
		TapScriptSig: tapScriptSig,
	}
}

// ValidateBasic implements the sdk.Msg interface.
func (m SubmitTapScriptSigRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if err := m.TapScriptSig.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return nil
}

// GetSigners implements the sdk.Msg interface
func (m SubmitTapScriptSigRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
