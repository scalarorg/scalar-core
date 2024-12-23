package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers concrete types on codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&CreateProtocolRequest{}, "protocol/CreateProtocol", nil)
	cdc.RegisterConcrete(&UpdateProtocolRequest{}, "protocol/UpdateProtocol", nil)
	cdc.RegisterConcrete(&AddSupportedChainRequest{}, "protocol/AddSupportedChain", nil)
	cdc.RegisterConcrete(&UpdateSupportedChainRequest{}, "protocol/UpdateSupportedChain", nil)
}

// RegisterInterfaces registers types and interfaces with the given registry
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&CreateProtocolRequest{},
		&UpdateProtocolRequest{},
		&AddSupportedChainRequest{},
		&UpdateSupportedChainRequest{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var amino = codec.NewLegacyAmino()

// ModuleCdc defines the module codec
var ModuleCdc = codec.NewAminoCodec(amino)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
