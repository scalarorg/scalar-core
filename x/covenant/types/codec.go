package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	gogoprototypes "github.com/gogo/protobuf/types"
	reward "github.com/scalarorg/scalar-core/x/reward/exported"
)

// RegisterLegacyAminoCodec registers concrete types on codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&CreateCustodianRequest{}, "covenant/CreateCustodian", nil)
	cdc.RegisterConcrete(&CreateCustodianGroupRequest{}, "covenant/CreateCustodianGroup", nil)
}

// RegisterInterfaces registers types and interfaces with the given registry
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&CreateCustodianRequest{},
		&CreateCustodianGroupRequest{},
		&SubmitTapScriptSigsRequest{},
	)
	registry.RegisterImplementations((*codec.ProtoMarshaler)(nil),
		&gogoprototypes.BoolValue{},
	)

	registry.RegisterImplementations((*reward.Refundable)(nil),
		&SubmitTapScriptSigsRequest{},
	)

	registry.RegisterImplementations((*codec.ProtoMarshaler)(nil),
		&PsbtMultiSig{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_MsgService_serviceDesc)
}

var amino = codec.NewLegacyAmino()

// ModuleCdc defines the module codec
var ModuleCdc = codec.NewAminoCodec(amino)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
