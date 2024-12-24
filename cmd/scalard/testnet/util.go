package testnet

import (
	"encoding/json"
	"os"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/rs/zerolog/log"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	pvm "github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
	"github.com/tendermint/tendermint/rpc/client/local"
)

// AppOptionsMap is a stub implementing AppOptions which can get data from a map
type AppOptionsMap map[string]interface{}

func (m AppOptionsMap) Get(key string) interface{} {
	v, ok := m[key]
	if !ok {
		return interface{}(nil)
	}

	return v
}

func NewAppOptionsWithFlagHome(homePath string) servertypes.AppOptions {
	return AppOptionsMap{
		flags.FlagHome: homePath,
	}
}
func generateKey(keyNme string, keyAlgo string, homePath string) error {
	addKeyCmd := keys.AddKeyCommand()
	addKeyCmd.SetArgs([]string{keyNme, "--algo", keyAlgo, "--keyring-backend", "test", "--home", homePath, "--output", "json"})
	err := addKeyCmd.Execute()
	if err != nil {
		log.Error().Err(err).Msg("[generateKey] Add key")
	}
	return err
}
func startInProcess(cfg Config, val *Validator) error {
	logger := val.Ctx.Logger
	tmCfg := val.Ctx.Config
	tmCfg.Instrumentation.Prometheus = false

	if err := val.AppConfig.ValidateBasic(); err != nil {
		return err
	}

	nodeKey, err := p2p.LoadOrGenNodeKey(tmCfg.NodeKeyFile())
	if err != nil {
		return err
	}

	app := cfg.AppConstructor(*val)

	genDocProvider := node.DefaultGenesisDocProviderFunc(tmCfg)
	log.Debug().Msgf("Address book file: %v", tmCfg.P2P.AddrBookFile())
	tmNode, err := node.NewNode(
		tmCfg,
		pvm.LoadOrGenFilePV(tmCfg.PrivValidatorKeyFile(), tmCfg.PrivValidatorStateFile()),
		nodeKey,
		proxy.NewLocalClientCreator(app),
		genDocProvider,
		node.DefaultDBProvider,
		node.DefaultMetricsProvider(tmCfg.Instrumentation),
		logger.With("module", val.Moniker),
	)
	if err != nil {
		return err
	}

	if err := tmNode.Start(); err != nil {
		return err
	}

	val.tmNode = tmNode

	if val.RPCAddress != "" {
		val.RPCClient = local.New(tmNode)
	}

	//We'll need a RPC client if the validator exposes a gRPC or REST endpoint.
	if val.APIAddress != "" || val.AppConfig.RPC.ListenAddress != "" {
		val.ClientCtx = val.ClientCtx.
			WithClient(val.RPCClient)

		// Add the tx service in the gRPC router.
		app.RegisterTxService(val.ClientCtx)

		// Add the tendermint queries service in the gRPC router.
		app.RegisterTendermintService(val.ClientCtx)
	}

	// if val.AppConfig.API.Enable && val.APIAddress != "" {
	// 	apiSrv := api.New(val.ClientCtx, logger.With("module", "api-server"))
	// 	app.RegisterAPIRoutes(apiSrv, val.AppConfig.API)

	// 	errCh := make(chan error)

	// 	go func() {
	// 		if err := apiSrv.Start(val.AppConfig.Config); err != nil {
	// 			errCh <- err
	// 		}
	// 	}()

	// 	select {
	// 	case err := <-errCh:
	// 		return err
	// 	case <-time.After(srvtypes.ServerStartTime): // assume server started successfully
	// 	}

	// 	val.api = apiSrv
	// }

	// if val.AppConfig.GRPC.Enable {
	// 	grpcSrv, err := servergrpc.StartGRPCServer(val.ClientCtx, app, val.AppConfig.GRPC)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	val.grpc = grpcSrv

	// 	if val.AppConfig.GRPCWeb.Enable {
	// 		val.grpcWeb, err = servergrpc.StartGRPCWeb(grpcSrv, val.AppConfig.Config)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	// if val.AppConfig.JSONRPC.Enable && val.AppConfig.JSONRPC.Address != "" {
	// 	if val.Ctx == nil || val.Ctx.Viper == nil {
	// 		return fmt.Errorf("validator %s context is nil", val.Moniker)
	// 	}

	// 	tmEndpoint := "/websocket"
	// 	tmRPCAddr := fmt.Sprintf("tcp://%s", val.AppConfig.GRPC.Address)

	// 	val.jsonrpc, val.jsonrpcDone, err = server.StartJSONRPC(val.Ctx, val.ClientCtx, tmRPCAddr, tmEndpoint, val.AppConfig, nil)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	address := fmt.Sprintf("http://%s", val.AppConfig.JSONRPC.Address)

	// 	val.JSONRPCClient, err = ethclient.Dial(address)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to dial JSON-RPC at %s: %w", val.AppConfig.JSONRPC.Address, err)
	// 	}
	// }

	return nil
}

func ParseJsonArrayConfig[T any](filePath string) ([]T, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read file")
		return nil, err
	}
	var result []T
	if err := json.Unmarshal(content, &result); err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal json")
		return nil, err
	}

	return result, nil
}

func ParseJsonConfig[T any](filePath string) (*T, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var result T
	if err := json.Unmarshal(content, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
