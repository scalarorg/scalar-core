package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/gogo/protobuf/proto"

	"github.com/axelarnetwork/axelar-core/x/reward/exported"
)

// RegisterLegacyAminoCodec registers concrete types on codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&LinkRequest{}, "scalarnet/Link", nil)
	cdc.RegisterConcrete(&ConfirmDepositRequest{}, "scalarnet/ConfirmDeposit", nil)
	cdc.RegisterConcrete(&ExecutePendingTransfersRequest{}, "scalarnet/ExecutePendingTransfers", nil)
	cdc.RegisterConcrete(&AddCosmosBasedChainRequest{}, "scalarnet/AddCosmosBasedChain", nil)
	cdc.RegisterConcrete(&RegisterAssetRequest{}, "scalarnet/RegisterAsset", nil)
	cdc.RegisterConcrete(&RouteIBCTransfersRequest{}, "scalarnet/RouteIBCTransfers", nil)
	cdc.RegisterConcrete(&RegisterFeeCollectorRequest{}, "scalarnet/RegisterFeeCollector", nil)
	cdc.RegisterConcrete(&RetryIBCTransferRequest{}, "scalarnet/RetryIBCTransfer", nil)
	cdc.RegisterConcrete(&RouteMessageRequest{}, "scalarnet/RouteMessage", nil)
	cdc.RegisterConcrete(&CallContractRequest{}, "scalarnet/CallContract", nil)

	cdc.RegisterConcrete(&CallContractsProposal{}, "scalarnet/CallContractsProposal", nil)
}

type customRegistry interface {
	RegisterCustomTypeURL(iface interface{}, typeURL string, impl proto.Message)
}

// RegisterInterfaces registers types and interfaces with the given registry
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&LinkRequest{},
		&ConfirmDepositRequest{},
		&ExecutePendingTransfersRequest{},
		&AddCosmosBasedChainRequest{},
		&RegisterAssetRequest{},
		&RouteIBCTransfersRequest{},
		&RegisterFeeCollectorRequest{},
		&RetryIBCTransferRequest{},
		&RouteMessageRequest{},
		&CallContractRequest{},
	)
	registry.RegisterInterface("reward.v1beta1.Refundable",
		(*exported.Refundable)(nil))

	// register renamed messages for old routes
	r, ok := registry.(customRegistry)
	if !ok {
		panic(fmt.Errorf("failed to convert registry type %T", registry))
	}

	r.RegisterCustomTypeURL((*sdk.Msg)(nil), "/axelar.scalarnet.v1beta1.ExecuteMessageRequest", &RouteMessageRequest{})

	registry.RegisterImplementations((*govtypes.Content)(nil),
		&CallContractsProposal{},
	)
}

var amino = codec.NewLegacyAmino()

// ModuleCdc defines the module codec
var ModuleCdc = codec.NewAminoCodec(amino)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
