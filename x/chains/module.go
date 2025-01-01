package chains

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
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/x/chains/client/cli"
	"github.com/scalarorg/scalar-core/x/chains/keeper"
	"github.com/scalarorg/scalar-core/x/chains/types"
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

func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	state := types.DefaultGenesisState()
	return cdc.MustMarshalJSON(&state)
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, genesisState json.RawMessage) error {
	var state types.GenesisState
	if err := cdc.UnmarshalJSON(genesisState, &state); err != nil {
		return err
	}
	return state.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module. Just maps handlers to the client.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryServiceHandlerClient(context.Background(), mux, types.NewQueryServiceClient(clientCtx)); err != nil {
		panic(err)
	}
}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.QuerierRoute)
}

type AppModule struct {
	AppModuleBasic
	keeper      *keeper.BaseKeeper
	voter       types.Voter
	nexus       types.Nexus
	snapshotter types.Snapshotter
	slashing    types.SlashingKeeper
	staking     types.StakingKeeper
	multisig    types.MultisigKeeper
	covenant    types.CovenantKeeper
	protocol    types.ProtocolKeeper
}

func NewAppModule(keeper *keeper.BaseKeeper,
	voter types.Voter,
	nexus types.Nexus,
	snapshotter types.Snapshotter,
	slashing types.SlashingKeeper,
	staking types.StakingKeeper,
	multisig types.MultisigKeeper,
	covenant types.CovenantKeeper,
	protocol types.ProtocolKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
		voter:          voter,
		nexus:          nexus,
		snapshotter:    snapshotter,
		slashing:       slashing,
		staking:        staking,
		multisig:       multisig,
		covenant:       covenant,
		protocol:       protocol,
	}
}

// RegisterServices registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	params := keeper.MsgServerConstructArgs{
		BaseKeeper:  am.keeper,
		Nexus:       am.nexus,
		Voter:       am.voter,
		Snapshotter: am.snapshotter,
		Slashing:    am.slashing,
		Multisig:    am.multisig,
		Covenant:    am.covenant,
		Protocol:    am.protocol,
	}
	msgServer := keeper.NewMsgServerImpl(params)
	types.RegisterMsgServiceServer(cfg.MsgServer(), msgServer)

	queryServer := keeper.NewGRPCQuerier(am.keeper, am.nexus, am.multisig)
	types.RegisterQueryServiceServer(cfg.QueryServer(), queryServer)

	// TODO: add migration
}

// BeginBlock executes all state transitions this module requires at the beginning of each new block
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	BeginBlocker(ctx, req, am.keeper)
}

// EndBlock executes all state transitions this module requires at the end of each new block
func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
	return utils.RunCached(ctx, am.keeper, func(ctx sdk.Context) ([]abci.ValidatorUpdate, error) {
		return EndBlocker(ctx, req, am.keeper, am.nexus, am.multisig)
	})
}
func (am AppModule) ConsensusVersion() uint64 {
	return 1
}

func (AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	cdc.MustUnmarshalJSON(gs, &genState)
	am.keeper.InitGenesis(ctx, genState)

	return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	state := types.DefaultGenesisState()
	return cdc.MustMarshalJSON(&state)
}

func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {
	// No invariants yet

	// TODO: implement
}

/**

Deprecated methods

**/

func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
}

func (am AppModule) Route() sdk.Route {
	return sdk.Route{}
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
}

func (am AppModule) LegacyQuerierHandler(legacyQuerierRouter *codec.LegacyAmino) sdk.Querier {
	return nil
}
