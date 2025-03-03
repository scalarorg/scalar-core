package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
)

var _ sdk.Msg = &SubmitTapScriptSigsRequest{}

// NewSubmitTapScriptSigRequest constructor for SubmitTapScriptSigRequest
func NewSubmitTapScriptSigsRequest(sender sdk.AccAddress, sigID uint64, list []*exported.TapScriptSigsMap) *SubmitTapScriptSigsRequest {
	return &SubmitTapScriptSigsRequest{
		Sender:                 sender,
		SigID:                  sigID,
		ListOfTapScriptSigsMap: list,
	}
}

// ValidateBasic implements the sdk.Msg interface.
func (request *SubmitTapScriptSigsRequest) ValidateBasic() error {
	clog.Magentaf("ValidateBasic for SubmitTapScriptSigsRequest, m: %+v", request)
	if err := sdk.VerifyAddressFormat(request.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if len(request.ListOfTapScriptSigsMap) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "list of tap script sigs is empty")
	}

	for _, m := range request.ListOfTapScriptSigsMap {
		if m == nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "tap script sigs is nil")
		}

		if len(m.Inner) == 0 {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "tap script sigs is empty")
		}

		for _, tapScriptList := range m.Inner {
			for _, tapScriptSig := range tapScriptList.Sigs.List {
				if err := tapScriptSig.ValidateBasic(); err != nil {
					return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
				}
			}
		}
	}

	return nil
}

// GetSigners implements the sdk.Msg interface
func (m SubmitTapScriptSigsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
