package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewConfirmSourceTxsRequest creates a message of type ConfirmSourceTxsRequest
func NewConfirmSourceTxsRequest(sender sdk.AccAddress, chain nexus.ChainName, txIDs []Hash) *ConfirmSourceTxsRequest {
	return &ConfirmSourceTxsRequest{
		Sender: sender,
		Chain:  chain,
		TxIDs:  txIDs,
	}
}

func (msg *ConfirmSourceTxsRequest) ValidateBasic() error {
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

func (m ConfirmSourceTxsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
