package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m UpdateUtxoForRedeemSessionRequest) ValidateBasic() error {
	if m.Symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}
	if len(m.ListOfUtxos) == 0 {
		return fmt.Errorf("list of utxos cannot be nil")
	}
	if m.BlockHeight == 0 {
		return fmt.Errorf("block height cannot be 0")
	}
	return nil
}

func (m UpdateUtxoForRedeemSessionRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
