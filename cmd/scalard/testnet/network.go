package testnet

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/app"
	"github.com/scalarorg/scalar-core/app/params"
	"github.com/scalarorg/scalar-core/config"
	scalartypes "github.com/scalarorg/scalar-core/types"
	scalarnet "github.com/scalarorg/scalar-core/x/scalarnet/exported"
	tmconfig "github.com/tendermint/tendermint/config"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/node"
	tmclient "github.com/tendermint/tendermint/rpc/client"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/grpc"
)

func DefaultOption(options *keyring.Options) {
	options.SupportedAlgos = keyring.SigningAlgoList{hd.Secp256k1}
	options.SupportedAlgosLedger = keyring.SigningAlgoList{hd.Secp256k1}
}

type GenesisState map[string]json.RawMessage

var (
	minGasPrice = "0.007" + scalarnet.NativeAsset
	wasmDirFlag = "wasm-dir"
)

// package-wide network lock to only allow one test network at a time
var lock = new(sync.Mutex)

// AppConstructor defines a function which accepts a network configuration and
// creates an ABCI Application to provide to Tendermint.
type AppConstructor = func(val Validator) servertypes.Application

// Config defines the necessary configuration used to bootstrap and start an
// in-process local testing network.
type Config struct {
	KeyringOptions    []keyring.Option // keyring configuration options
	Codec             codec.Codec
	LegacyAmino       *codec.LegacyAmino // TODO: Remove!
	InterfaceRegistry codectypes.InterfaceRegistry
	TxConfig          client.TxConfig
	AccountRetriever  client.AccountRetriever
	AppConstructor    AppConstructor // the ABCI application constructor
	GenesisState      GenesisState   // custom gensis state to provide
	TimeoutCommit     time.Duration  // the consensus commitment timeout
	AccountTokens     sdk.Int        // the amount of unique validator tokens (e.g. 1000node0)
	StakingTokens     sdk.Int        // the amount of tokens each validator has available to stake
	BondedTokens      sdk.Int        // the amount of tokens each validator stakes
	NumValidators     int            // the total number of validators to create and bond
	ChainID           string         // the network chain-id
	BondDenom         string         // the staking bond denomination
	MinGasPrices      string         // the minimum gas prices each validator will accept
	PruningStrategy   string         // the pruning strategy each validator will have
	SigningAlgo       string         // signing algorithm for keys
	RPCAddress        string         // RPC listen address (including port)
	JSONRPCAddress    string         // JSON-RPC listen address (including port)
	APIAddress        string         // REST API listen address (including port)
	GRPCAddress       string         // GRPC server listen address (including port)
	EnableTMLogging   bool           // enable Tendermint logging to STDOUT
	CleanupDir        bool           // remove base temporary directory during cleanup
	PrintMnemonic     bool           // print the mnemonic of first validator as log output for testing
}

// DefaultConfig returns a sane default configuration suitable for nearly all
// testing requirements.
func DefaultConfig() Config {
	encCfg := app.MakeEncodingConfig()
	chainID := fmt.Sprintf("scalar_%d-1", tmrand.Int63n(9999999999999)+1)
	return Config{
		Codec:             encCfg.Codec,
		TxConfig:          encCfg.TxConfig,
		LegacyAmino:       encCfg.Amino,
		InterfaceRegistry: encCfg.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor:    NewAppConstructor(encCfg, chainID),
		GenesisState:      app.GetModuleBasics().DefaultGenesis(encCfg.Codec),
		TimeoutCommit:     3 * time.Second,
		ChainID:           chainID,
		BondDenom:         scalarnet.NativeAsset,
		MinGasPrices:      fmt.Sprintf("0.007%s", scalarnet.BaseAsset),
		AccountTokens:     sdk.TokensFromConsensusPower(10, scalartypes.PowerReduction),
		StakingTokens:     sdk.TokensFromConsensusPower(5, scalartypes.PowerReduction),
		BondedTokens:      sdk.TokensFromConsensusPower(1, scalartypes.PowerReduction),
		CleanupDir:        true,
		SigningAlgo:       string(hd.Secp256k1Type),
		KeyringOptions:    []keyring.Option{DefaultOption},
		PrintMnemonic:     false,
	}
}

// NewAppConstructor returns a new Evmos AppConstructor
func NewAppConstructor(encodingCfg params.EncodingConfig, chainID string) AppConstructor {
	return func(val Validator) servertypes.Application {
		appOpts := NewAppOptionsWithFlagHome(val.Ctx.Config.RootDir)
		return app.NewScalarApp(
			val.Ctx.Logger,         //loger
			dbm.NewMemDB(),         //db
			nil,                    //tracerStore
			true,                   //loadLatest
			make(map[int64]bool),   //skipUpgradeHeights
			val.Ctx.Config.RootDir, //homeDir
			"",                     //wasmDir
			0,                      //invCheckPeriod
			encodingCfg,            //encodingConfig
			appOpts,                //appOpts
			[]wasm.Option{},        //wasmOpts
			// simutils.NewAppOptionsWithFlagHome(val.Ctx.Config.RootDir), //appOpts
			// baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			// baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
			// baseapp.SetChainID(chainID),
		)
	}
}

type (
	// Network defines a local in-process testing network using SimApp. It can be
	// configured to start any number of validators, each with its own RPC and API
	// clients. Typically, this test network would be used in client and integration
	// testing where user input is expected.
	//
	// Note, due to Tendermint constraints in regards to RPC functionality, there
	// may only be one test network running at a time. Thus, any caller must be
	// sure to Cleanup after testing is finished in order to allow other tests
	// to create networks. In addition, only the first validator will have a valid
	// RPC and API server/client.
	Network struct {
		Logger     Logger
		BaseDir    string
		Validators []*Validator

		Config Config
	}

	// Validator defines an in-process Tendermint validator node. Through this object,
	// a client can make RPC and API calls and interact with any client command
	// or handler.
	Validator struct {
		AppConfig     *tmconfig.Config
		ClientCtx     client.Context
		Ctx           *server.Context
		Dir           string
		NodeID        string
		PubKey        cryptotypes.PubKey
		Moniker       string
		APIAddress    string
		RPCAddress    string
		P2PAddress    string
		Address       sdk.AccAddress
		ValAddress    sdk.ValAddress
		RPCClient     tmclient.Client
		JSONRPCClient *ethclient.Client

		tmNode      *node.Node
		api         *api.Server
		grpc        *grpc.Server
		grpcWeb     *http.Server
		jsonrpc     *http.Server
		jsonrpcDone chan struct{}
	}
)

// New creates a new Network for integration tests or in-process testnets run via the CLI
func New(l Logger, ioReader io.Reader, baseDir string, cfg Config) (*Network, error) {
	// only one caller/test can create and use a network at a time
	log.Debug().Msg("acquiring test network lock")
	lock.Lock()

	network := &Network{
		Logger:     l,
		BaseDir:    baseDir,
		Validators: make([]*Validator, cfg.NumValidators),
		Config:     cfg,
	}

	l.Logf("preparing test network with chain-id \"%s\"\n", cfg.ChainID)

	// generate private keys, node IDs, and initial transactions
	for i := 0; i < cfg.NumValidators; i++ {

		appCfg := tmconfig.DefaultConfig()
		// appCfg.Pruning = cfg.PruningStrategy
		// appCfg.MinGasPrices = cfg.MinGasPrices
		// appCfg.API.Enable = true
		// appCfg.API.Swagger = false
		// appCfg.Telemetry.Enabled = false
		// appCfg.Telemetry.GlobalLabels = [][]string{{"chain_id", cfg.ChainID}}

		nodeMoniker := fmt.Sprintf("node%d", i+1)
		valRootDir := filepath.Join(network.BaseDir, nodeMoniker)
		ctx := server.NewDefaultContext()
		tmCfg := ctx.Config
		tmCfg.RootDir = filepath.Join(valRootDir, "scalard")
		tmCfg.P2P.RootDir = filepath.Join(valRootDir, "scalard")
		tmCfg.Moniker = nodeMoniker
		tmCfg.Consensus.TimeoutCommit = cfg.TimeoutCommit

		// Only allow the first validator to expose an RPC, API and gRPC
		// server/client due to Tendermint in-process constraints.
		// tmCfg.RPC.ListenAddress = ""
		// appCfg.GRPC.Enable = false
		// appCfg.GRPCWeb.Enable = false
		clientDir := filepath.Join(network.BaseDir, nodeMoniker, "scalard")
		ctx.Viper.AddConfigPath(filepath.Join(valRootDir, "scalard/config"))
		//Read seeds.toml
		seeds, err := config.ReadSeeds(ctx.Viper)
		if err != nil {
			return nil, err
		}
		log.Debug().Msgf("Seeds: %v", seeds)
		tmCfg = config.MergeSeeds(tmCfg, seeds)
		log.Debug().Msgf("P2P Seeds: %v", tmCfg.P2P.Seeds)
		ctx.Viper.SetConfigType("toml")
		ctx.Viper.SetConfigName("config")
		err = ctx.Viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
		// Read into the configuration whatever data the viper instance has for it.
		// This may come from the configuration file above but also any of the other
		// sources viper uses.
		if err := ctx.Viper.Unmarshal(&tmCfg); err != nil {
			return nil, err
		}

		//fmt.Println("App config: %v", tmCfg.RPC.ListenAddress)

		kb, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, clientDir, bufio.NewReader(ioReader), DefaultOption)
		if err != nil {
			return nil, err
		}

		// keyringAlgos, _ := kb.SupportedAlgorithms()
		// algo, err := keyring.NewSigningAlgoFromString(cfg.SigningAlgo, keyringAlgos)
		// if err != nil {
		// 	return nil, err
		// }

		clientCtx := client.Context{}.
			WithKeyringDir(clientDir).
			WithKeyring(kb).
			WithHomeDir(tmCfg.RootDir).
			WithChainID(cfg.ChainID).
			WithInterfaceRegistry(cfg.InterfaceRegistry).
			WithCodec(cfg.Codec).
			WithLegacyAmino(cfg.LegacyAmino).
			WithTxConfig(cfg.TxConfig).
			WithAccountRetriever(cfg.AccountRetriever)

		network.Validators[i] = &Validator{
			AppConfig: appCfg,
			ClientCtx: clientCtx,
			Ctx:       ctx,
			Dir:       filepath.Join(valRootDir, "scalard"),
			// NodeID: nodeID,
			// PubKey:     pubKey,
			Moniker:    nodeMoniker,
			RPCAddress: tmCfg.RPC.ListenAddress,
			// P2PAddress: tmCfg.P2P.ListenAddress,
			// APIAddress: apiAddr,
			// Address:    addr,
			// ValAddress: sdk.ValAddress(addr),
		}
	}

	l.Log("starting test network...")
	for _, v := range network.Validators {
		err := startInProcess(cfg, v)
		if err != nil {
			return nil, err
		}
	}

	l.Log("started test network")

	// Ensure we cleanup incase any test was abruptly halted (e.g. SIGINT) as any
	// defer in a test would not be called.
	// server.TrapSignal(network.Cleanup)

	return network, nil
}

// LatestHeight returns the latest height of the network or an error if the
// query fails or no validators exist.
func (n *Network) LatestHeight() (int64, error) {
	if len(n.Validators) == 0 {
		return 0, errors.New("no validators available")
	}

	status, err := n.Validators[0].RPCClient.Status(context.Background())
	if err != nil {
		return 0, err
	}

	return status.SyncInfo.LatestBlockHeight, nil
}

// WaitForHeight performs a blocking check where it waits for a block to be
// committed after a given block. If that height is not reached within a timeout,
// an error is returned. Regardless, the latest height queried is returned.
func (n *Network) WaitForHeight(h int64) (int64, error) {
	return n.WaitForHeightWithTimeout(h, 10*time.Second)
}

// WaitForHeightWithTimeout is the same as WaitForHeight except the caller can
// provide a custom timeout.
func (n *Network) WaitForHeightWithTimeout(h int64, t time.Duration) (int64, error) {
	ticker := time.NewTicker(time.Second)
	timeout := time.After(t)

	if len(n.Validators) == 0 {
		return 0, errors.New("no validators available")
	}

	var latestHeight int64
	val := n.Validators[0]

	for {
		select {
		case <-timeout:
			ticker.Stop()
			return latestHeight, errors.New("timeout exceeded waiting for block")
		case <-ticker.C:
			status, err := val.RPCClient.Status(context.Background())
			if err == nil && status != nil {
				latestHeight = status.SyncInfo.LatestBlockHeight
				if latestHeight >= h {
					return latestHeight, nil
				}
			}
		}
	}
}

// WaitForNextBlock waits for the next block to be committed, returning an error
// upon failure.
func (n *Network) WaitForNextBlock() error {
	lastBlock, err := n.LatestHeight()
	if err != nil {
		return err
	}

	_, err = n.WaitForHeight(lastBlock + 1)
	if err != nil {
		return err
	}

	return err
}

// Cleanup removes the root testing (temporary) directory and stops both the
// Tendermint and API services. It allows other callers to create and start
// test networks. This method must be called when a test is finished, typically
// in a defer.
func (n *Network) Cleanup() {
	defer func() {
		lock.Unlock()
		n.Logger.Log("released test network lock")
	}()

	n.Logger.Log("cleaning up test network...")

	for _, v := range n.Validators {
		if v.tmNode != nil && v.tmNode.IsRunning() {
			_ = v.tmNode.Stop()
		}

		if v.api != nil {
			_ = v.api.Close()
		}

		if v.grpc != nil {
			v.grpc.Stop()
			if v.grpcWeb != nil {
				_ = v.grpcWeb.Close()
			}
		}

		if v.jsonrpc != nil {
			shutdownCtx, cancelFn := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancelFn()

			if err := v.jsonrpc.Shutdown(shutdownCtx); err != nil {
				v.tmNode.Logger.Error("HTTP server shutdown produced a warning", "error", err.Error())
			} else {
				v.tmNode.Logger.Info("HTTP server shut down, waiting 5 sec")
				select {
				case <-time.Tick(5 * time.Second):
				case <-v.jsonrpcDone:
				}
			}
		}
	}

	if n.Config.CleanupDir {
		_ = os.RemoveAll(n.BaseDir)
	}

	n.Logger.Log("finished cleaning up test network")
}

// printMnemonic prints a provided mnemonic seed phrase on a network logger
// for debugging and manual testing
func printMnemonic(l Logger, secret string) {
	lines := []string{
		"THIS MNEMONIC IS FOR TESTING PURPOSES ONLY",
		"DO NOT USE IN PRODUCTION",
		"",
		strings.Join(strings.Fields(secret)[0:8], " "),
		strings.Join(strings.Fields(secret)[8:16], " "),
		strings.Join(strings.Fields(secret)[16:24], " "),
	}

	lineLengths := make([]int, len(lines))
	for i, line := range lines {
		lineLengths[i] = len(line)
	}

	maxLineLength := 0
	for _, lineLen := range lineLengths {
		if lineLen > maxLineLength {
			maxLineLength = lineLen
		}
	}

	l.Log("\n")
	l.Log(strings.Repeat("+", maxLineLength+8))
	for _, line := range lines {
		l.Logf("++  %s  ++\n", centerText(line, maxLineLength))
	}
	l.Log(strings.Repeat("+", maxLineLength+8))
	l.Log("\n")
}

// centerText centers text across a fixed width, filling either side with whitespace buffers
func centerText(text string, width int) string {
	textLen := len(text)
	leftBuffer := strings.Repeat(" ", (width-textLen)/2)
	rightBuffer := strings.Repeat(" ", (width-textLen)/2+(width-textLen)%2)

	return fmt.Sprintf("%s%s%s", leftBuffer, text, rightBuffer)
}
