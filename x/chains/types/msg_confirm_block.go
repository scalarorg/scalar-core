package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/scalarorg/scalar-core/x/chains/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// NewConfirmBlock creates a message of type ConfirmSourceTxsRequest
func NewConfirmBlock(sender sdk.AccAddress, chain nexus.ChainName, blockHash exported.Hash) *ConfirmBlockRequest {
	return &ConfirmBlockRequest{
		Sender:    sender,
		Chain:     chain,
		BlockHash: blockHash,
	}
}

func (msg *ConfirmBlockRequest) ValidateBasic() error {
	// TODO: validate the txIDs
	if err := sdk.VerifyAddressFormat(msg.Sender); err != nil {
		return err
	}
	if msg.Chain == "" {
		return fmt.Errorf("chain is required")
	}
	if msg.BlockHash.IsZero() {
		return fmt.Errorf("block hash is required")
	}
	return nil
}

func (m ConfirmBlockRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
