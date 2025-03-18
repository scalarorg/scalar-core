package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var _ sdk.Msg = &ConfirmSwitchedPhaseRequest{}

func NewConfirmSwitchedPhaseRequest(sender sdk.AccAddress, chain string, txID string) *ConfirmSwitchedPhaseRequest {
	return &ConfirmSwitchedPhaseRequest{
		Sender: sender,
		Chain:  nexus.ChainName(chain),
		TxID:   Hash(common.HexToHash(txID)),
	}
}

func (msg *ConfirmSwitchedPhaseRequest) ValidateBasic() error {
	return nil
}

func (msg *ConfirmSwitchedPhaseRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
