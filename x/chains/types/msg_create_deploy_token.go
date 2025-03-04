package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/scalarorg/scalar-core/utils"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewCreateDeployTokenRequest is the constructor for CreateDeployTokenRequest
func NewCreateDeployTokenRequest(sender sdk.AccAddress, chain string, tokenSymbol string, tokenName string, address Address) *CreateDeployTokenRequest {
	return &CreateDeployTokenRequest{
		Sender:           sender,
		Chain:            nexus.ChainName(utils.NormalizeString(chain)),
		TokenSymbol:      tokenSymbol,
		AliasedTokenName: tokenName,
		Address:          address,
	}
}

// Route implements sdk.Msg
func (m CreateDeployTokenRequest) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (m CreateDeployTokenRequest) Type() string {
	return "CreateDeployToken"
}

// GetSignBytes  implements sdk.Msg
func (m CreateDeployTokenRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (m CreateDeployTokenRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}

// ValidateBasic implements sdk.Msg
func (m CreateDeployTokenRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	if err := m.Chain.Validate(); err != nil {
		return sdkerrors.Wrap(err, "invalid chain")
	}

	// if err := m.Asset.Validate(); err != nil {
	// 	return err
	// }

	// switch m.Address.IsZeroAddress() {
	// case true:
	// 	if m.Chain.Equals(m.Asset.Chain) {
	// 		return fmt.Errorf("cannot deploy token on the origin chain")
	// 	}
	// case false:
	// 	if !m.Chain.Equals(m.Asset.Chain) {
	// 		return fmt.Errorf("cannot link token on a different chain")
	// 	}
	// }

	// if err := m.TokenDetails.Validate(); err != nil {
	// 	return err
	// }

	// if _, err := sdk.ParseUint(m.DailyMintLimit); err != nil {
	// 	return err
	// }

	return nil
}
