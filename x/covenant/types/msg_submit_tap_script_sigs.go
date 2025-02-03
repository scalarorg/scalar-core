package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
)

var _ sdk.Msg = &SubmitTapScriptSigsRequest{}

// NewSubmitTapScriptSigRequest constructor for SubmitTapScriptSigRequest
func NewSubmitTapScriptSigsRequest(sender sdk.AccAddress, sigID uint64, tapScriptSigs *exported.TapScriptSigsMap) *SubmitTapScriptSigsRequest {
	return &SubmitTapScriptSigsRequest{
		Sender:           sender,
		SigID:            sigID,
		TapScriptSigsMap: tapScriptSigs,
	}
}

// ValidateBasic implements the sdk.Msg interface.
func (m SubmitTapScriptSigsRequest) ValidateBasic() error {
	clog.Magentaf("ValidateBasic for SubmitTapScriptSigsRequest, m: %+v", m)
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if m.TapScriptSigsMap == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "tap script sigs is nil")
	}

	if m.TapScriptSigsMap == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "tap script sigs is nil")
	}

	if len(m.TapScriptSigsMap.Inner) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "tap script sigs is empty")
	}

	for _, tapScriptList := range m.TapScriptSigsMap.Inner {
		for _, tapScriptSig := range tapScriptList.Sigs.List {
			if err := tapScriptSig.ValidateBasic(); err != nil {
				return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
			}
		}
	}

	return nil
}

// GetSigners implements the sdk.Msg interface
func (m SubmitTapScriptSigsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
