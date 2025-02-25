package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Sender            github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
// 	BitcoinPubkey     []byte                                        `protobuf:"bytes,2,opt,name=bitcoin_pubkey,json=bitcoinPubkey,proto3" json:"bitcoin_pubkey,omitempty"`
// 	ScalarPubkey      []byte                                        `protobuf:"bytes,3,opt,name=scalar_pubkey,json=scalarPubkey,proto3" json:"scalar_pubkey,omitempty"`
// 	Name              string                                        `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
// 	Tag               string                                        `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`
// 	Attribute         *ProtocolAttribute                            `protobuf:"bytes,6,opt,name=attribute,proto3" json:"attribute,omitempty"`
// 	CustodianGroupUid string                                        `protobuf:"bytes,7,opt,name=custodian_group_uid,json=custodianGroupUid,proto3" json:"custodian_group_uid,omitempty"`
// 	Avatar            []byte                                        `protobuf:"bytes,8,opt,name=avatar,proto3" json:"avatar,omitempty"`

func NewCreateProtocolRequest(sender sdk.AccAddress, name string, bitcoinPubkey []byte, scalarPubkey []byte, tag string, attribute *ProtocolAttribute, custodianGroupUid string, avatar []byte) *CreateProtocolRequest {
	return &CreateProtocolRequest{
		Sender:            sender,
		Name:              name,
		BitcoinPubkey:     bitcoinPubkey,
		ScalarPubkey:      scalarPubkey,
		Tag:               tag,
		Attribute:         attribute,
		CustodianGroupUid: custodianGroupUid,
		Avatar:            avatar,
	}
}

// Route returns the route for this message
func (m CreateProtocolRequest) Route() string {
	return RouterKey
}

// Type returns the type of the message
func (m CreateProtocolRequest) Type() string {
	return "CreateProtocol"
}

// ValidateBasic executes a stateless message validation
func (m CreateProtocolRequest) ValidateBasic() error {
	if err := sdk.VerifyAddressFormat(m.Sender); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, sdkerrors.Wrap(err, "sender").Error())
	}

	return nil
}

// GetSignBytes returns the message bytes that need to be signed
func (m CreateProtocolRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the set of signers for this message
func (m CreateProtocolRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Sender}
}
