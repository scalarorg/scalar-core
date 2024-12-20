package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogoprototypes "github.com/gogo/protobuf/types"
)

// RegisterInterfaces registers types and interfaces with the given registry
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&ConfirmSourceTxsRequest{},
		&ConfirmDestTxsRequest{},
	)

	registry.RegisterImplementations((*codec.ProtoMarshaler)(nil),
		&gogoprototypes.BoolValue{},
		&Event{},
		&VoteEvents{},
		&PollMetadata{},
	)
}
