package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m UpdateSequenceAndPhaseForRedeemSessionRequest) ValidateBasic() error {
	if m.Symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}
	if m.Sequence == 0 {
		return fmt.Errorf("sequence cannot be 0")
	}
	if m.Phase == Switching {
		return fmt.Errorf("phase cannot be switching")
	}
	return nil
}

func (m UpdateSequenceAndPhaseForRedeemSessionRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
