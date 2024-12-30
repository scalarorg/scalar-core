package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogoprototypes "github.com/gogo/protobuf/types"
)

// RegisterInterfaces registers types and interfaces with the given registry
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&LinkRequest{},
		&ConfirmTokenRequest{},
		&ConfirmDepositRequest{},
		&ConfirmTransferKeyRequest{},
		&CreatePendingTransfersRequest{},
		&CreateDeployTokenRequest{},
		&CreateBurnTokensRequest{},
		&CreateTransferOperatorshipRequest{},
		&SignCommandsRequest{},
		&SignBTCCommandsRequest{},
		&AddChainRequest{},
		&SetGatewayRequest{},
		&RetryFailedEventRequest{},
		&ConfirmSourceTxsRequest{},
	)
	registry.RegisterImplementations((*codec.ProtoMarshaler)(nil),
		&gogoprototypes.BoolValue{},
		&SigMetadata{},
		&Event{},
		&VoteEvents{},
		&PollMetadata{},
	)
}

var amino = codec.NewLegacyAmino()

// ModuleCdc defines the module codec
var ModuleCdc = codec.NewAminoCodec(amino)

func init() {
	// RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
