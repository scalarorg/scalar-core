package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewConfirmBridgeTxsRequest creates a message of type ConfirmBridgeTxsRequest
func NewConfirmBridgeTxsRequest(sender sdk.AccAddress, chain nexus.ChainName, txIDs []Hash) *ConfirmBridgeTxsRequest {
	return &ConfirmBridgeTxsRequest{
		Sender: sender,
		Chain:  chain,
		TxIDs:  txIDs,
	}
}

func (msg *ConfirmBridgeTxsRequest) ValidateBasic() error {
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

func (m ConfirmBridgeTxsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
