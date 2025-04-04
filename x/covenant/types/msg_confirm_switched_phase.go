package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var _ sdk.Msg = &ConfirmSwitchedPhaseRequest{}

func NewConfirmSwitchedPhaseRequest(sender sdk.AccAddress, chain string, custodianGroupUid string, txID string) *ConfirmSwitchedPhaseRequest {
	return &ConfirmSwitchedPhaseRequest{
		Sender:            sender,
		Chain:             nexus.ChainName(chain),
		CustodianGroupUID: chains.Hash(common.HexToHash(custodianGroupUid)),
		TxID:              chains.Hash(common.HexToHash(txID)),
	}
}

func (msg *ConfirmSwitchedPhaseRequest) ValidateBasic() error {
	return nil
}

func (msg *ConfirmSwitchedPhaseRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
