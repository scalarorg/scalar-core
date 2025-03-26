package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var _ sdk.Msg = &RegisterCustodianGroupRequest{}

func NewRegisterCustodianGroupRequest(sender sdk.AccAddress, chain string, custodianGroupUID chains.Hash) *RegisterCustodianGroupRequest {
	return &RegisterCustodianGroupRequest{
		Sender:            sender,
		Chain:             nexus.ChainName(chain),
		CustodianGroupUID: custodianGroupUID,
	}
}

func (msg *RegisterCustodianGroupRequest) ValidateBasic() error {
	// TODO: validate the txID
	if err := sdk.VerifyAddressFormat(msg.Sender); err != nil {
		return err
	}
	if msg.Chain == "" {
		return fmt.Errorf("chain is required")
	}
	if msg.CustodianGroupUID.IsZero() {
		return fmt.Errorf("custodian group uid is required")
	}

	return nil
}

func (msg *RegisterCustodianGroupRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
