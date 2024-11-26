package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	hd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkserver "github.com/cosmos/cosmos-sdk/server"
	sdkconfig "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/scalarorg/scalar-core/cmd/scalard/cmd/utils"
	"github.com/scalarorg/scalar-core/types"
	scalartypes "github.com/scalarorg/scalar-core/types"
	"github.com/spf13/cobra"
	tmconfig "github.com/tendermint/tendermint/config"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	tmtypes "github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

const (
	DefaultGRPCAddress    = "0.0.0.0:9090"
	DefaultJSONRPCAddress = "0.0.0.0:8545"
)

var (
	flagNodeDirPrefix   = "node-dir-prefix"
	flagNumValidators   = "v"
	flagOutputDir       = "output-dir"
	flagNodeDaemonHome  = "node-daemon-home"
	flagNodeDomain      = "node-domain"
	flagEnableLogging   = "enable-logging"
	flagRPCAddress      = "rpc.address"
	flagAPIAddress      = "api.address"
	flagPrintMnemonic   = "print-mnemonic"
	flagBaseFee         = "base-fee"
	flagMinGasPrice     = "min-gas-price"
	flagGRPCAddress     = "grpc.address"
	flagJSONRPCAddress  = "json-rpc.address"
	flagKeyType         = "key-type"
	moduleNameFeemarket = "feemarket"
)

type initArgs struct {
	algo           string
	chainID        string
	keyringBackend string
	minGasPrices   string
	nodeDaemonHome string
	nodeDirPrefix  string
	numValidators  int
	outputDir      string
	nodeDomain     string
	baseFee        sdk.Int
	minGasPrice    sdk.Dec
}

type startArgs struct {
	algo           string
	apiAddress     string
	chainID        string
	grpcAddress    string
	minGasPrices   string
	outputDir      string
	rpcAddress     string
	jsonrpcAddress string
	numValidators  int
	enableLogging  bool
	printMnemonic  bool
}

func defaultOption(options *keyring.Options) {
	options.SupportedAlgos = keyring.SigningAlgoList{hd.Secp256k1}
	options.SupportedAlgosLedger = keyring.SigningAlgoList{hd.Secp256k1}
}

// createValidatorMsgGasLimit is the gas limit used in the MsgCreateValidator included in genesis transactions.
// This transaction consumes approximately 220,000 gas when executed in the genesis block.
const createValidatorMsgGasLimit = 250_000

func addTestnetFlagsToCmd(cmd *cobra.Command) {
	cmd.Flags().Int(flagNumValidators, 4, "Number of validators to initialize the testnet with")
	cmd.Flags().StringP(flagOutputDir, "o", "./.testnets", "Directory to store initialization data for the testnet")
	cmd.Flags().String(flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().String(sdkserver.FlagMinGasPrices, fmt.Sprintf("0.000006%s", scalartypes.BaseDenom), "Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01photino,0.001stake)")
	cmd.Flags().String(flagKeyType, string(hd.Secp256k1Type), "Key signing algorithm to generate keys for")
	cmd.Flags().String(flagBaseFee, strconv.Itoa(params.InitialBaseFee), "The params base_fee in the feemarket module in geneis")
	cmd.Flags().String(flagMinGasPrice, "0", "The params min_gas_price in the feemarket module in geneis")
}

// NewTestnetCmd creates a root testnet command with subcommands to run an in-process testnet or initialize
// validator configuration files for running a multi-validator testnet in a separate process
func NewTestnetCmd(mbm module.BasicManager, genBalIterator banktypes.GenesisBalancesIterator) *cobra.Command {
	testnetCmd := &cobra.Command{
		Use:                        "testnet",
		Short:                      "subcommands for starting or configuring local testnets",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	testnetCmd.AddCommand(testnetInitFilesCmd(mbm, genBalIterator))
	//testnetCmd.AddCommand(testnetStartCmd())

	return testnetCmd
}

// get cmd to initialize all files for tendermint testnet and application
func testnetInitFilesCmd(mbm module.BasicManager, genBalIterator banktypes.GenesisBalancesIterator) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init-files",
		Short: "Initialize config directories & files for a multi-validator testnet running locally via separate processes (e.g. Docker Compose or similar)",
		Long: `init-files will setup "v" number of directories and populate each with
necessary files (private validator, genesis, config, etc.) for running "v" validator nodes.

Booting up a network with these validator folders is intended to be used with Docker Compose,
or a similar setup where each node has a manually configurable IP address.

Note, strict routability for addresses is turned off in the config file.

Example:
	scalard testnet init-files --v 4 --output-dir ./.testnets --node-domain scalarnode
	`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			serverCtx := sdkserver.GetServerContextFromCmd(cmd)

			args := initArgs{}
			args.outputDir, _ = cmd.Flags().GetString(flagOutputDir)
			args.keyringBackend, _ = cmd.Flags().GetString(flags.FlagKeyringBackend)
			args.chainID, _ = cmd.Flags().GetString(flags.FlagChainID)
			args.minGasPrices, _ = cmd.Flags().GetString(sdkserver.FlagMinGasPrices)
			args.nodeDirPrefix, _ = cmd.Flags().GetString(flagNodeDirPrefix)
			args.nodeDaemonHome, _ = cmd.Flags().GetString(flagNodeDaemonHome)
			args.nodeDomain, _ = cmd.Flags().GetString(flagNodeDomain)
			args.numValidators, _ = cmd.Flags().GetInt(flagNumValidators)
			args.algo, _ = cmd.Flags().GetString(flagKeyType)
			baseFee, _ := cmd.Flags().GetString(flagBaseFee)
			minGasPrice, _ := cmd.Flags().GetString(flagMinGasPrice)

			var ok bool
			args.baseFee, ok = sdk.NewIntFromString(baseFee)
			if !ok || args.baseFee.LT(sdk.ZeroInt()) {
				return fmt.Errorf("invalid value for --base-fee. expected a int number greater than or equal to 0 but got %s", baseFee)
			}

			args.minGasPrice, err = sdk.NewDecFromStr(minGasPrice)
			if err != nil {
				return fmt.Errorf("invalid value for --min-gas-price. expected a int or decimal greater than or equal to 0 but got %s and err %s", minGasPrice, err.Error())
			}
			if args.minGasPrice.LT(sdk.ZeroDec()) {
				return fmt.Errorf("invalid value for --min-gas-price. expected a int or decimal greater than or equal to 0 but got an negative number %s", minGasPrice)
			}

			return initTestnetFiles(clientCtx, cmd, serverCtx.Config, mbm, genBalIterator, args)
		},
	}

	addTestnetFlagsToCmd(cmd)
	cmd.Flags().String(flagNodeDirPrefix, "node", "Prefix the directory name for each node with (node results in node1, node2, ...)")
	cmd.Flags().String(flagNodeDaemonHome, "scalard", "Home directory of the node's daemon configuration")
	cmd.Flags().String(flagNodeDomain, "scalarnode", `Node domain: 
		*scalarnode* results in persistent peers list ID0@scalarnode1:46656, ID1@scalarnode2:46656, ...
		*192.168.0.1* results in persistent peers list ID0@192.168.0.11:46656, ID1@192.168.0.12:46656, ...
		`)

	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")

	return cmd
}

const nodeDirPerm = 0o755

// initTestnetFiles initializes testnet files for a testnet to be run in a separate process
func initTestnetFiles(
	clientCtx client.Context,
	cmd *cobra.Command,
	nodeConfig *tmconfig.Config,
	mbm module.BasicManager,
	genBalIterator banktypes.GenesisBalancesIterator,
	args initArgs,
) error {
	if args.chainID == "" {
		args.chainID = fmt.Sprintf("scalar_%d-1", tmrand.Int63n(9999999999999)+1)
	}

	nodeIDs := make([]string, args.numValidators)
	valPubKeys := make([]cryptotypes.PubKey, args.numValidators)

	var (
		genAccounts []authtypes.GenesisAccount
		genBalances []banktypes.Balance
		genFiles    []string
	)

	inBuf := bufio.NewReader(cmd.InOrStdin())
	// generate private keys, node IDs, and initial transactions
	for i := 0; i < args.numValidators; i++ {
		nodeDirName := getNodeDirName(i, args.nodeDirPrefix)
		nodeDir := filepath.Join(args.outputDir, nodeDirName, args.nodeDaemonHome)
		gentxsDir := filepath.Join(args.outputDir, "gentxs")

		nodeConfig.SetRoot(nodeDir)
		nodeConfig.RPC.ListenAddress = "tcp://0.0.0.0:26657"

		if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
			_ = os.RemoveAll(args.outputDir)
			return err
		}

		nodeConfig.Moniker = nodeDirName

		host, err := getHost(i, args.nodeDomain)
		if err != nil {
			_ = os.RemoveAll(args.outputDir)
			return err
		}

		nodeIDs[i], valPubKeys[i], err = genutil.InitializeNodeValidatorFiles(nodeConfig)
		if err != nil {
			_ = os.RemoveAll(args.outputDir)
			return err
		}

		memo := fmt.Sprintf("%s@%s:26656", nodeIDs[i], host)
		genFiles = append(genFiles, nodeConfig.GenesisFile())
		// TODO: add ledger support
		kb, err := keyring.New(sdk.KeyringServiceName(), args.keyringBackend, nodeDir, inBuf, defaultOption)
		if err != nil {
			return err
		}

		keyringAlgos, _ := kb.SupportedAlgorithms()
		algo, err := keyring.NewSigningAlgoFromString(args.algo, keyringAlgos)
		if err != nil {
			return err
		}

		addr, secret, err := testutil.GenerateSaveCoinKey(kb, nodeDirName, "", true, algo)
		if err != nil {
			_ = os.RemoveAll(args.outputDir)
			return err
		}

		info := map[string]string{"secret": secret}

		cliPrint, err := json.Marshal(info)
		if err != nil {
			return err
		}

		// save private key seed words
		if err := utils.WriteFile(fmt.Sprintf("%v.json", "key_seed"), nodeDir, cliPrint); err != nil {
			return err
		}
		accStakingTokens := sdk.TokensFromConsensusPower(5000, scalartypes.PowerReduction)
		coins := sdk.Coins{
			sdk.NewCoin(scalartypes.BaseDenom, accStakingTokens),
		}

		genBalances = append(genBalances, banktypes.Balance{Address: addr.String(), Coins: coins.Sort()})
		genAccounts = append(genAccounts, authtypes.NewBaseAccount(addr, nil, 0, 0))

		valTokens := sdk.TokensFromConsensusPower(100, scalartypes.PowerReduction)
		createValMsg, err := stakingtypes.NewMsgCreateValidator(
			sdk.ValAddress(addr),
			valPubKeys[i],
			sdk.NewCoin(scalartypes.BaseDenom, valTokens),
			stakingtypes.NewDescription(nodeDirName, "", "", "", ""),
			stakingtypes.NewCommissionRates(sdk.OneDec(), sdk.OneDec(), sdk.OneDec()),
			sdk.OneInt(),
		)
		if err != nil {
			return err
		}

		txBuilder := clientCtx.TxConfig.NewTxBuilder()
		if err := txBuilder.SetMsgs(createValMsg); err != nil {
			return err
		}

		minGasPrice := args.minGasPrice
		if sdk.NewDecFromInt(args.baseFee).GT(args.minGasPrice) {
			minGasPrice = sdk.NewDecFromInt(args.baseFee)
		}

		txBuilder.SetMemo(memo)
		txBuilder.SetGasLimit(createValidatorMsgGasLimit)
		txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewCoin(scalartypes.BaseDenom, minGasPrice.MulInt64(createValidatorMsgGasLimit).Ceil().TruncateInt())))

		txFactory := tx.Factory{}
		txFactory = txFactory.
			WithChainID(args.chainID).
			WithMemo(memo).
			WithKeybase(kb).
			WithTxConfig(clientCtx.TxConfig)

		if err := tx.Sign(txFactory, nodeDirName, txBuilder, true); err != nil {
			return err
		}

		txBz, err := clientCtx.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
		if err != nil {
			return err
		}

		if err := utils.WriteFile(fmt.Sprintf("%v.json", nodeDirName), gentxsDir, txBz); err != nil {
			return err
		}
		// Add custom app config
		if err := setCustomAppConfig(cmd); err != nil {
			return err
		}
		//Generate default app config
		appConfig := sdkconfig.DefaultConfig()
		appConfig.MinGasPrices = args.minGasPrices
		appConfig.API.Enable = true
		appConfig.Telemetry.Enabled = true
		appConfig.Telemetry.PrometheusRetentionTime = 60
		appConfig.Telemetry.EnableHostnameLabel = false
		appConfig.Telemetry.GlobalLabels = [][]string{{"chain_id", args.chainID}}
		sdkconfig.WriteConfigFile(filepath.Join(nodeDir, "config/app.toml"), appConfig)
	}

	if err := initGenFiles(clientCtx, mbm, args.chainID, scalartypes.BaseDenom, genAccounts, genBalances, genFiles, args.numValidators, args.baseFee, args.minGasPrice); err != nil {
		return err
	}

	err := collectGenFiles(
		clientCtx, nodeConfig, args.chainID, nodeIDs, valPubKeys, args.numValidators,
		args.outputDir, args.nodeDirPrefix, args.nodeDaemonHome, genBalIterator,
	)
	if err != nil {
		return err
	}

	cmd.PrintErrf("Successfully initialized %d node directories\n", args.numValidators)
	return nil
}
func setCustomAppConfig(cmd *cobra.Command) error {
	//T odo: add custom app config if needed
	// customAppTemplate, customAppConfig := createCustomAppConfig(scalartypes.BaseDenom)
	customAppTemplate, customAppConfig := sdkconfig.DefaultConfigTemplate, sdkconfig.DefaultConfig()
	sdkconfig.SetConfigTemplate(customAppTemplate)

	return sdkserver.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig)
}
func initGenFiles(
	clientCtx client.Context,
	mbm module.BasicManager,
	chainID,
	coinDenom string,
	genAccounts []authtypes.GenesisAccount,
	genBalances []banktypes.Balance,
	genFiles []string,
	numValidators int,
	baseFee sdk.Int,
	minGasPrice sdk.Dec,
) error {
	appGenState, err := types.GenerateGenesis(clientCtx, mbm, coinDenom, genAccounts, genBalances)
	if err != nil {
		return err
	}

	appGenStateJSON, err := json.MarshalIndent(appGenState, "", "  ")
	if err != nil {
		return err
	}

	genDoc := tmtypes.GenesisDoc{
		ChainID:    chainID,
		AppState:   appGenStateJSON,
		Validators: nil,
	}

	// generate empty genesis files for each validator and save
	for i := 0; i < numValidators; i++ {
		if err := genDoc.SaveAs(genFiles[i]); err != nil {
			return err
		}
	}
	return nil
}

func collectGenFiles(
	clientCtx client.Context, nodeConfig *tmconfig.Config, chainID string,
	nodeIDs []string, valPubKeys []cryptotypes.PubKey, numValidators int,
	outputDir, nodeDirPrefix, nodeDaemonHome string, genBalIterator banktypes.GenesisBalancesIterator,
) error {
	var appState json.RawMessage
	genTime := tmtime.Now()

	for i := 0; i < numValidators; i++ {
		nodeDirName := getNodeDirName(i, nodeDirPrefix)
		nodeDir := filepath.Join(outputDir, nodeDirName, nodeDaemonHome)
		gentxsDir := filepath.Join(outputDir, "gentxs")
		nodeConfig.Moniker = nodeDirName

		nodeConfig.SetRoot(nodeDir)

		nodeID, valPubKey := nodeIDs[i], valPubKeys[i]
		initCfg := genutiltypes.NewInitConfig(chainID, gentxsDir, nodeID, valPubKey)

		genDoc, err := tmtypes.GenesisDocFromFile(nodeConfig.GenesisFile())
		if err != nil {
			return err
		}

		nodeAppState, err := genutil.GenAppStateFromConfig(
			clientCtx.Codec, clientCtx.TxConfig,
			nodeConfig, initCfg, *genDoc, genBalIterator)
		if err != nil {
			return err
		}

		if appState == nil {
			// set the canonical application state (they should not differ)
			appState = nodeAppState
		}

		genFile := nodeConfig.GenesisFile()

		// overwrite each validator's genesis file to have a canonical genesis time
		if err := genutil.ExportGenesisFileWithTime(genFile, chainID, nil, appState, genTime); err != nil {
			return err
		}
	}

	return nil
}
func getNodeDirName(i int, nodeDirPrefix string) string {
	return fmt.Sprintf("%s%d", nodeDirPrefix, i+1)
}
func getHost(i int, nodeDomain string) (host string, err error) {
	if len(nodeDomain) == 0 {
		host, err = sdkserver.ExternalIP()
		if err != nil {
			return "", err
		}
		return host, nil
	}
	return fmt.Sprintf("%s%d", nodeDomain, i+1), nil
}
func getIP(i int, startingIPAddr string) (ip string, err error) {
	if len(startingIPAddr) == 0 {
		ip, err = sdkserver.ExternalIP()
		if err != nil {
			return "", err
		}
		return ip, nil
	}
	return calculateIP(startingIPAddr, i)
}

func calculateIP(ip string, i int) (string, error) {
	ipv4 := net.ParseIP(ip).To4()
	if ipv4 == nil {
		return "", fmt.Errorf("%v: non ipv4 address", ip)
	}

	for j := 0; j < i; j++ {
		ipv4[3]++
	}

	return ipv4.String(), nil
}

// // get cmd to start multi validator in-process testnet
// func testnetStartCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "start",
// 		Short: "Launch an in-process multi-validator testnet",
// 		Long: `testnet will launch an in-process multi-validator testnet,
// and generate "v" directories, populated with necessary validator configuration files
// (private validator, genesis, config, etc.).

// Example:
// 	scalard testnet --v 4 --output-dir ./.testnets
// 	`,
// 		RunE: func(cmd *cobra.Command, _ []string) error {
// 			args := startArgs{}
// 			args.outputDir, _ = cmd.Flags().GetString(flagOutputDir)
// 			args.chainID, _ = cmd.Flags().GetString(flags.FlagChainID)
// 			args.minGasPrices, _ = cmd.Flags().GetString(sdkserver.FlagMinGasPrices)
// 			args.numValidators, _ = cmd.Flags().GetInt(flagNumValidators)
// 			args.algo, _ = cmd.Flags().GetString(flagKeyType)
// 			args.enableLogging, _ = cmd.Flags().GetBool(flagEnableLogging)
// 			args.rpcAddress, _ = cmd.Flags().GetString(flagRPCAddress)
// 			args.apiAddress, _ = cmd.Flags().GetString(flagAPIAddress)
// 			args.grpcAddress, _ = cmd.Flags().GetString(flagGRPCAddress)
// 			args.jsonrpcAddress, _ = cmd.Flags().GetString(flagJSONRPCAddress)
// 			args.printMnemonic, _ = cmd.Flags().GetBool(flagPrintMnemonic)

// 			return startTestnet(cmd, args)
// 		},
// 	}

// 	addTestnetFlagsToCmd(cmd)
// 	cmd.Flags().Bool(flagEnableLogging, false, "Enable INFO logging of tendermint validator nodes")
// 	cmd.Flags().String(flagRPCAddress, "tcp://0.0.0.0:26657", "the RPC address to listen on")
// 	cmd.Flags().String(flagAPIAddress, "tcp://0.0.0.0:1317", "the address to listen on for REST API")
// 	cmd.Flags().String(flagGRPCAddress, DefaultGRPCAddress, "the gRPC server address to listen on")
// 	cmd.Flags().String(flagJSONRPCAddress, DefaultJSONRPCAddress, "the JSON-RPC server address to listen on")
// 	cmd.Flags().Bool(flagPrintMnemonic, true, "print mnemonic of first validator to stdout for manual testing")
// 	return cmd
// }

// // startTestnet starts an in-process testnet
// func startTestnet(cmd *cobra.Command, args startArgs) error {
// 	networkConfig := network.DefaultConfig()

// 	// Default networkConfig.ChainID is random, and we should only override it if chainID provided
// 	// is non-empty
// 	if args.chainID != "" {
// 		networkConfig.ChainID = args.chainID
// 	}
// 	networkConfig.SigningAlgo = args.algo
// 	networkConfig.MinGasPrices = args.minGasPrices
// 	networkConfig.NumValidators = args.numValidators
// 	networkConfig.EnableTMLogging = args.enableLogging
// 	networkConfig.RPCAddress = args.rpcAddress
// 	networkConfig.APIAddress = args.apiAddress
// 	networkConfig.GRPCAddress = args.grpcAddress
// 	networkConfig.JSONRPCAddress = args.jsonrpcAddress
// 	networkConfig.PrintMnemonic = args.printMnemonic
// 	networkLogger := network.NewCLILogger(cmd)

// 	baseDir := fmt.Sprintf("%s/%s", args.outputDir, networkConfig.ChainID)
// 	if _, err := os.Stat(baseDir); !os.IsNotExist(err) {
// 		return fmt.Errorf(
// 			"testnests directory already exists for chain-id '%s': %s, please remove or select a new --chain-id",
// 			networkConfig.ChainID, baseDir)
// 	}

// 	testnet, err := network.New(networkLogger, baseDir, networkConfig)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = testnet.WaitForHeight(1)
// 	if err != nil {
// 		return err
// 	}

// 	cmd.Println("press the Enter Key to terminate")
// 	_, err = fmt.Scanln() // wait for Enter Key
// 	if err != nil {
// 		return err
// 	}
// 	testnet.Cleanup()

// 	return nil
// }
