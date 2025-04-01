package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var _ sdk.Msg = &InitializeUtxoRequest{}

func NewInitializeUtxoRequest(sender sdk.AccAddress, chain nexus.ChainName, blockHeight uint64) *InitializeUtxoRequest {
	return &InitializeUtxoRequest{
		Sender:          sender,
		Chain:           chain,
		BlockCheckpoint: blockHeight,
	}
}

func (msg *InitializeUtxoRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(msg.Sender); err != nil {
		return err
	}
	if msg.Chain == "" {
		return fmt.Errorf("chain is required")
	}
	chainName := nexus.ChainName(msg.Chain)
	if chainName.GetFamily() != nexus.BITCOIN {
		return fmt.Errorf("chain %s is not a bitcoin chain", msg.Chain)
	}
	if msg.BlockCheckpoint <= 0 {
		return fmt.Errorf("block checkpoint is required and must be greater than 0")
	}
	return nil
}

func (msg *InitializeUtxoRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
