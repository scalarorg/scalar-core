package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/scalarorg/scalar-core/utils"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewConfirmDestTxsRequest creates a message of type ConfirmDestTxsRequest
func NewConfirmDestTxsRequest(sender sdk.AccAddress, chain nexus.ChainName, txIDs []Hash) *ConfirmDestTxsRequest {
	return &ConfirmDestTxsRequest{
		Sender: sender,
		Chain:  chain,
		TxIDs:  txIDs,
	}
}

func (msg *ConfirmDestTxsRequest) ValidateBasic() error {
	// TODO: validate the txIDs
	if err := sdk.VerifyAddressFormat(msg.Sender); err != nil {
		return err
	}
	_, err := utils.ChainInfoBytesFromString(string(msg.Chain))
	if err != nil {
		return fmt.Errorf("invalid chain: %v", err)
	}

	if len(msg.TxIDs) == 0 {
		return fmt.Errorf("txIDs are required")
	}
	return nil
}

func (m ConfirmDestTxsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
