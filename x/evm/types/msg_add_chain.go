package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/scalarorg/scalar-core/utils"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"
)

// NewAddChainRequest is the constructor for NewAddChainRequest
func NewAddChainRequest(sender sdk.AccAddress, name string, params Params) *AddChainRequest {
	return &AddChainRequest{
		Sender: sender,
		Name:   nexus.ChainName(utils.NormalizeString(name)),
		Params: params,
	}
}

// Route returns the route for this message
func (m AddChainRequest) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m AddChainRequest) Type() string {
	return "AddChain"
}

// ValidateBasic executes a stateless message validation
func (m AddChainRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	chain := nexus.Chain{
		Name:                  m.Name,
		SupportsForeignAssets: true,
		KeyType:               tss.Multisig,
		Module:                ModuleName,
	}

	if err := chain.Validate(); err != nil {
		return fmt.Errorf("invalid chain spec: %v", err)
	}

	if err := m.Params.Validate(); err != nil {
		return fmt.Errorf("invalid EVM param: %v", err)
	}

	if !m.Name.Equals(m.Params.Chain) {
		return fmt.Errorf("chain mismatch: chain name is %s, parameters chain is %s", m.Name, m.Params.Chain)
	}

	return nil
}

// GetSignBytes returns the message bytes that need to be signed
func (m AddChainRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the set of signers for this message
func (m AddChainRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
