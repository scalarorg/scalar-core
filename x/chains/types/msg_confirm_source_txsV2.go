package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewConfirmSourceTxsRequest creates a message of type ConfirmSourceTxsRequest
func NewConfirmSourceTxsRequestV2(sender sdk.AccAddress, chain nexus.ChainName, batch *TrustedTxsByBlock) *ConfirmSourceTxsRequestV2 {
	return &ConfirmSourceTxsRequestV2{
		Sender: sender,
		Chain:  chain,
		Batch:  batch,
	}
}

func (msg *ConfirmSourceTxsRequestV2) ValidateBasic() error {
	// TODO: validate the txIDs
	if err := sdk.VerifyAddressFormat(msg.Sender); err != nil {
		return err
	}
	if msg.Chain == "" {
		return fmt.Errorf("chain is required")
	}
	if len(msg.Batch.Txs) == 0 {
		return fmt.Errorf("txIDs are required")
	}
	return nil
}

func (m ConfirmSourceTxsRequestV2) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
