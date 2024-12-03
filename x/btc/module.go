package btc

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/scalarorg/scalar-core/x/btc/keeper"
	"github.com/scalarorg/scalar-core/x/btc/types"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic implements module.AppModuleBasic
type AppModuleBasic struct {
}

func (AppModuleBasic) Name() string {
	return types.ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	// types.RegisterLegacyAminoCodec(cdc)
}

func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// types.RegisterInterfaces(registry)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, genesisState json.RawMessage) error {
	var state types.GenesisState
	if err := cdc.UnmarshalJSON(genesisState, &state); err != nil {
		return err
	}
	return state.Validate()
}

func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
	// rest.RegisterRoutes(clientCtx, rtr)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryServiceHandlerClient(context.Background(), mux, types.NewQueryServiceClient(clientCtx)); err != nil {
		panic(err)
	}
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	// return cli.GetTxCmd()
	return nil
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	// return cli.GetQueryCmd(types.QuerierRoute)
	return nil
}

type AppModule struct {
	AppModuleBasic
	keeper *keeper.BaseKeeper
}

func NewAppModule(keeper *keeper.BaseKeeper) AppModule {
	return AppModule{
		keeper: keeper,
	}
}

func (am AppModule) ConsensusVersion() uint64 {
	return 6
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	cdc.MustUnmarshalJSON(gs, &genState)
	am.keeper.InitGenesis(ctx, genState)

	return []abci.ValidatorUpdate{}
}

func (am AppModule) LegacyQuerierHandler(legacyQuerierRouter *codec.LegacyAmino) sdk.Querier {
	return keeper.NewQuerier(am.keeper)
}

func (AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

// RegisterServices registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	// TODO: implement
}

func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {
	// No invariants yet

	// TODO: implement
}

func (am AppModule) Route() sdk.Route {
	// return sdk.NewRoute(types.RouterKey, NewHandler(am.keeper))
	return sdk.Route{}
}

func (am AppModule) BeginBlock(ctx sdk.Context) error {
	return nil
}

func (am AppModule) EndBlock(ctx sdk.Context) error {
	return nil
}
