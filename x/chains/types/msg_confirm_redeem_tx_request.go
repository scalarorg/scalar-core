package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var _ sdk.Msg = &ConfirmRedeemTxRequest{}

func NewConfirmRedeemTxRequest(sender sdk.AccAddress, chain string, txID string) *ConfirmRedeemTxRequest {
	return &ConfirmRedeemTxRequest{
		Sender: sender,
		Chain:  nexus.ChainName(chain),
		TxID:   chains.Hash(common.HexToHash(txID)),
	}
}

func (msg *ConfirmRedeemTxRequest) ValidateBasic() error {
	return nil
}

func (msg *ConfirmRedeemTxRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
