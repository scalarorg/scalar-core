package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewConfirmStakingTxsRequest creates a message of type ConfirmStakingTxsRequest
func NewConfirmStakingTxsRequest(sender sdk.AccAddress, chain nexus.ChainName, txIDs []Hash) *ConfirmStakingTxsRequest {
	return &ConfirmStakingTxsRequest{
		Sender: sender,
		Chain:  chain,
		TxIDs:  txIDs,
	}
}

func (msg *ConfirmStakingTxsRequest) ValidateBasic() error {
	// TODO: validate the txIDs
	if err := sdk.VerifyAddressFormat(msg.Sender); err != nil {
		return err
	}
	if msg.Chain == "" {
		return fmt.Errorf("chain is required")
	}
	if len(msg.TxIDs) == 0 {
		return fmt.Errorf("txIDs are required")
	}
	return nil
}

func (m ConfirmStakingTxsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
