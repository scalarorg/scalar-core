package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var _ sdk.Msg = &UpdateNewBtcBlockRequest{}

func NewUpdateNewBtcBlockRequest(sender sdk.AccAddress, chain string, height uint64) *UpdateNewBtcBlockRequest {
	return &UpdateNewBtcBlockRequest{
		Sender: sender,
		Chain:  nexus.ChainName(chain),
		Height: height,
	}
}

func (msg *UpdateNewBtcBlockRequest) ValidateBasic() error {
	return nil
}

func (msg *UpdateNewBtcBlockRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
