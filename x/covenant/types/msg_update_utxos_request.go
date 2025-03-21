package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

var _ sdk.Msg = &UpdateUtxoListsRequest{}

func NewUpdateUtxoListsRequest(sender sdk.AccAddress, chain string, height uint64, taprootPubkey []byte) *UpdateUtxoListsRequest {
	return &UpdateUtxoListsRequest{
		Sender:        sender,
		Chain:         nexus.ChainName(chain),
		Height:        height,
		TaprootPubkey: taprootPubkey,
	}
}

func (msg *UpdateUtxoListsRequest) ValidateBasic() error {
	return nil
}

func (msg *UpdateUtxoListsRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
