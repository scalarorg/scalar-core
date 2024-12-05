package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
)

// NewConfirmGatewayTxsRequest creates a message of type ConfirmGatewayTxsRequest
func NewConfirmGatewayTxsRequest(sender sdk.AccAddress, chain nexus.ChainName, txIDs []TxHash) *ConfirmGatewayTxsRequest {
	return &ConfirmGatewayTxsRequest{
		Sender: sender,
		Chain:  chain,
		TxIDs:  txIDs,
	}
}

func (msg *ConfirmGatewayTxsRequest) ValidateBasic() error {
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

func (m ConfirmGatewayTxsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
