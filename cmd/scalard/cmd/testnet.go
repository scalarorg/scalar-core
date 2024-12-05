package cmd

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cometbft/cometbft/p2p"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	hd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkserver "github.com/cosmos/cosmos-sdk/server"
	sdkconfig "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/cmd/scalard/cmd/utils"
	"github.com/scalarorg/scalar-core/types"
	"github.com/tendermint/tendermint/privval"

	permexported "github.com/axelarnetwork/axelar-core/x/permission/exported"
	permtypes "github.com/axelarnetwork/axelar-core/x/permission/types"
	scalartypes "github.com/scalarorg/scalar-core/types"
	"github.com/spf13/cobra"
	tmconfig "github.com/tendermint/tendermint/config"
	tmed25519 "github.com/tendermint/tendermint/crypto/ed25519"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	tmtypes "github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

const (
	DefaultGRPCAddress    = "0.0.0.0:9090"
	DefaultJSONRPCAddress = "0.0.0.0:8545"
)

var (
	flagNodeMnemonic        = "NODE_MNEMONIC"
	flagBroadcasterMnemonic = "BROADCASTER_MNEMONIC"
	flagGovernanceMnemonic  = "GOV_MNEMONIC"
	flagValidatorMnemonic   = "VALIDATOR_MNEMONIC"
	flagBtcPubkey           = "BTC_PUBKEY"
	flagNodeDirPrefix       = "node-dir-prefix"
	flagNumValidators       = "v"
	flagSupportedChains     = "supported-chains"
	flagOutputDir           = "output-dir"
	flagNodeDaemonHome      = "node-daemon-home"
	flagNodeDomain          = "node-domain"
	flagEnableLogging       = "enable-logging"
	flagRPCAddress          = "rpc.address"
	flagAPIAddress          = "api.address"
	flagPrintMnemonic       = "print-mnemonic"
	flagBaseFee             = "base-fee"
	flagMinGasPrice         = "min-gas-price"
	flagGRPCAddress         = "grpc.address"
	flagJSONRPCAddress      = "json-rpc.address"
	flagKeyType             = "key-type"
	moduleNameFeemarket     = "feemarket"
)

type initArgs struct {
	algo            string
	chainID         string
	keyringBackend  string
	minGasPrices    string
	nodeDaemonHome  string
	supportedChains string
	nodeDirPrefix   string
	numValidators   int
	outputDir       string
	nodeDomain      string
	baseFee         sdk.Int
	minGasPrice     sdk.Dec
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

type EnvKeys struct {
	NodeMnemonic        string
	ValidatorMnemonic   string
	BroadcasterMnemonic string
	GovernanceMnemonic  string
	BtcPubkey           string
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
	scalard testnet init-files --v 4 --output-dir ./.testnets --node-domain scalarnode --supported-chains=./chains
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
			args.supportedChains, _ = cmd.Flags().GetString(flagSupportedChains)
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
	cmd.Flags().String(flagSupportedChains, "./chains", `Supported chains directory, each chain family is config in a seperated json file under this directory: 
		*./chains/evm.json* stores all evm chain configs ...
		*./chains/btc.json* stores all btc chain configs ...
		`)
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")

	return cmd
}

const nodeDirPerm = 0o755

func readEnvMnemonic(i int) EnvKeys {
	envKeys := EnvKeys{}
	index := strconv.Itoa(i)
	if envKeys.NodeMnemonic = os.Getenv(flagNodeMnemonic + index); envKeys.NodeMnemonic == "" {
		envKeys.NodeMnemonic = os.Getenv(flagNodeMnemonic)
	}
	if envKeys.ValidatorMnemonic = os.Getenv(flagValidatorMnemonic + index); envKeys.ValidatorMnemonic == "" {
		envKeys.ValidatorMnemonic = os.Getenv(flagValidatorMnemonic)
	}
	if envKeys.BroadcasterMnemonic = os.Getenv(flagBroadcasterMnemonic + index); envKeys.BroadcasterMnemonic == "" {
		envKeys.BroadcasterMnemonic = os.Getenv(flagBroadcasterMnemonic)
	}
	if envKeys.GovernanceMnemonic = os.Getenv(flagGovernanceMnemonic + index); envKeys.GovernanceMnemonic == "" {
		envKeys.GovernanceMnemonic = os.Getenv(flagGovernanceMnemonic)
	}
	if envKeys.BtcPubkey = os.Getenv(flagBtcPubkey + index); envKeys.BtcPubkey == "" {
		envKeys.BtcPubkey = os.Getenv(flagBtcPubkey)
	}
	return envKeys
}

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
	fmt.Printf("nodeConfig: %v\n", nodeConfig)
	var (
		validatorInfos []scalartypes.ValidatorInfo
	)
	// generate private keys, node IDs, and initial transactions
	for i := 0; i < args.numValidators; i++ {
		nodeDirName := getNodeDirName(i, args.nodeDirPrefix)
		host, err := getHost(i, args.nodeDomain)
		if err != nil {
			_ = os.RemoveAll(args.outputDir)
			return err
		}
		// Validator index starts from 1
		envMnemonic := readEnvMnemonic(i + 1)
		validatorInfo, err := initValidatorConfig(clientCtx, cmd, nodeConfig, host, nodeDirName, args, envMnemonic, int64((i+1)*(i+1)))
		if err != nil {
			_ = os.RemoveAll(args.outputDir)
			cmd.PrintErrf("failed to initialize validator config: %s", err.Error())
			return err
		}
		validatorInfos = append(validatorInfos, *validatorInfo)
	}
	if err := initGenFiles(clientCtx, mbm,
		args.chainID,
		scalartypes.BaseDenom,
		args.supportedChains,
		validatorInfos,
		args.baseFee,
		args.minGasPrice,
	); err != nil {
		cmd.PrintErrf("failed to initGenFiles: %s", err.Error())
		return err
	}

	err := collectGenFiles(clientCtx, nodeConfig, args.chainID, validatorInfos, args.outputDir, genBalIterator)
	if err != nil {
		cmd.PrintErrf("failed to collect genesis files: %s", err.Error())
		return err
	}

	cmd.PrintErrf("Successfully initialized %d node directories\n", args.numValidators)
	return nil
}

func createPubkeyFromMnemonic(config *tmconfig.Config, mnemonic string, keybase keyring.Keyring, algo keyring.SignatureAlgo, pvKeyName string) (cryptotypes.PubKey, error) {
	privKey := tmed25519.GenPrivKeyFromSecret([]byte(mnemonic))
	var pvKeyFile, pvStateFile string
	if pvKeyName == "" {
		pvKeyFile = config.PrivValidatorKeyFile()
		if err := tmos.EnsureDir(filepath.Dir(pvKeyFile), 0o777); err != nil {
			return nil, err
		}

		pvStateFile = config.PrivValidatorStateFile()
		if err := tmos.EnsureDir(filepath.Dir(pvStateFile), 0o777); err != nil {
			return nil, err
		}
		pvKeyName = types.ValidatorKeyName
	} else {
		pvKeyFile = filepath.Join(config.RootDir, "config", fmt.Sprintf("%s_key.json", pvKeyName))
		pvStateFile = filepath.Join(config.RootDir, "data", fmt.Sprintf("%s_state.json", pvKeyName))
	}
	//Input private key mnemonic
	// err := keybase.ImportPrivKey(pvKeyName, hex.EncodeToString(privKey), keyring.DefaultBIP39Passphrase)
	// if err != nil {
	// 	return nil, err
	// }
	// info, err := keybase.Key(pvKeyName)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Printf("info: %v\n", info)
	// Store private key in keybase
	//Todo: clearify the usage of keybase.NewAccount and tmed25519.GenPrivKeyFromSecret
	keybase.NewAccount(
		pvKeyName,
		mnemonic,
		keyring.DefaultBIP39Passphrase,
		sdk.GetConfig().GetFullBIP44Path(),
		algo,
	)
	filePV := privval.NewFilePV(privKey, pvKeyFile, pvStateFile)
	if err := tmos.EnsureDir(filepath.Dir(pvKeyFile), 0o777); err != nil {
		return nil, err
	}
	fmt.Printf("PrivValidator saved to file: %s\n", pvKeyFile)
	filePV.Save()

	valPubKey, err := cryptocodec.FromTmPubKeyInterface(privKey.PubKey())
	if err != nil {
		return nil, fmt.Errorf("failed to convert tmtypes.Pubkey to cryptotypes.PubKey: %w", err)
	}
	return valPubKey, nil
}
func createNodeID(config *tmconfig.Config) (string, error) {
	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return "", err
	}

	nodeID := string(nodeKey.ID())
	return nodeID, nil
}

// func createKeyFromMnemonic(keybase keyring.Keyring, keyName string, mnemonic string, algo keyring.SignatureAlgo) (cryptotypes.PubKey, string, error) {
// 	info, secret, err := keybase.NewMnemonic(
// 		keyName,
// 		keyring.English,
// 		sdk.GetConfig().GetFullBIP44Path(),
// 		keyring.DefaultBIP39Passphrase,
// 		algo,
// 	)
// 	if err != nil {
// 		return nil, secret, err
// 	}
// 	return info.GetPubKey(), secret, nil
// }

func initValidatorConfig(clientCtx client.Context, cmd *cobra.Command,
	nodeConfig *tmconfig.Config,
	host string,
	nodeDirName string,
	args initArgs,
	envKeys EnvKeys,
	power int64,
) (*scalartypes.ValidatorInfo, error) {
	var err error
	nodeDir := filepath.Join(args.outputDir, nodeDirName, args.nodeDaemonHome)
	nodeConfig.SetRoot(nodeDir)
	if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
		return nil, err
	}
	fmt.Printf("Create validator config in dir %s\n", nodeDir)
	nodeConfig.Moniker = nodeDirName
	validatorInfo := scalartypes.ValidatorInfo{
		Host:      host,
		Moniker:   nodeConfig.Moniker,
		NodeDir:   filepath.Join(nodeDir, "config"),
		GenFile:   nodeConfig.GenesisFile(),
		BtcPubkey: envKeys.BtcPubkey,
	}
	// validatorInfo.NodeID, validatorInfo.ValPubKey, err = genutil.InitializeNodeValidatorFilesFromMnemonic(nodeConfig, envKeys.ValidatorMnemonic)
	validatorInfo.NodeID, err = createNodeID(nodeConfig)
	if err != nil {
		return nil, err
	}
	validatorInfo.SeedAddress = fmt.Sprintf("%s@%s:26656", validatorInfo.NodeID, host)
	// validatorInfo.nodeID, validatorInfo.valPubKey, err = genutil.InitializeNodeValidatorFiles(nodeConfig)
	// if err != nil {
	// 	return nil, err
	// }

	inBuf := bufio.NewReader(cmd.InOrStdin())
	gentxsDir := filepath.Join(args.outputDir, "gentxs")
	nodeConfig.RPC.ListenAddress = "tcp://0.0.0.0:26657"

	// TODO: add ledger support
	kb, err := keyring.New(sdk.KeyringServiceName(), args.keyringBackend, nodeDir, inBuf, defaultOption)
	if err != nil {
		return nil, err
	}

	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(args.algo, keyringAlgos)
	if err != nil {
		return nil, err
	}
	validatorInfo.ValPubKey, err = createPubkeyFromMnemonic(nodeConfig, envKeys.ValidatorMnemonic, kb, algo, "")
	if err != nil {
		return nil, err
	}
	if envKeys.BroadcasterMnemonic != "" {
		broadcasterPubKey, err := createPubkeyFromMnemonic(nodeConfig, envKeys.BroadcasterMnemonic, kb, algo, scalartypes.BroadcasterKeyName)
		if err != nil {
			fmt.Printf("ExtractBroadcaster Err: %s\n", err.Error())
			return nil, err
		}
		validatorInfo.Broadcaster = broadcasterPubKey
	}
	if envKeys.GovernanceMnemonic != "" {
		validatorInfo.GovPubKey, err = createPubkeyFromMnemonic(nodeConfig, envKeys.GovernanceMnemonic, kb, algo, scalartypes.GovKeyName)
		if err != nil {
			fmt.Printf("ExtractGovernance Err: %s\n", err.Error())
			return nil, err
		}
		validatorInfo.MngAccount = permtypes.GovAccount{
			Address: sdk.AccAddress(validatorInfo.GovPubKey.Address()),
			Role:    permexported.ROLE_CHAIN_MANAGEMENT,
		}
	}
	senderKeyName := nodeDirName
	//Generate node key with name {senderKeyName} inthe keyring
	senderAddress, secret, err := testutil.GenerateSaveCoinKey(kb, senderKeyName, envKeys.NodeMnemonic, true, algo)
	fmt.Printf("nodeAddr: %v\n", senderAddress)
	if err != nil {
		_ = os.RemoveAll(args.outputDir)
		return nil, err
	}
	key, err := kb.Key(senderKeyName)
	if err != nil {
		return nil, err
	}
	info := map[string]string{
		"secret":  secret,
		"address": sdk.ValAddress(senderAddress).String(),
		"pubkey":  key.GetPubKey().String(),
	}

	cliPrint, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	// save private key seed words
	if err := utils.WriteFile(fmt.Sprintf("%v.json", "key_seed"), nodeDir, cliPrint); err != nil {
		return nil, err
	}
	validatorInfo.NodeBalance = banktypes.Balance{
		Address: senderAddress.String(),
		Coins:   sdk.Coins{sdk.NewCoin(scalartypes.BaseDenom, scalartypes.NodeTokens)},
	}
	//Use validator address as node address
	// senderAddress := sdk.AccAddress(validatorInfo.Broadcaster.Address())
	// senderKeyName = types.BroadcasterKeyName
	valTokens := sdk.TokensFromConsensusPower(power, scalartypes.ValidatorTokens)
	valCoin := sdk.NewCoin(scalartypes.BaseDenom, valTokens)
	validatorInfo.ValBalance = banktypes.Balance{
		Address: sdk.AccAddress(validatorInfo.ValPubKey.Address()).String(),
		Coins:   sdk.Coins{valCoin},
	}

	validatorInfo.BroadcasterBalance = banktypes.Balance{
		Address: sdk.AccAddress(validatorInfo.Broadcaster.Address()).String(),
		Coins: sdk.Coins{
			sdk.NewCoin(scalartypes.BaseDenom, scalartypes.BroadcasterTokens),
		},
	}
	tmPubKey, err := cryptocodec.ToTmPubKeyInterface(validatorInfo.ValPubKey)
	if err != nil {
		fmt.Printf("ToTmPubKeyInterface Err: %s\n", err.Error())
		return nil, err
	}
	validatorInfo.GenesisValidator = tmtypes.GenesisValidator{
		Name:    nodeDirName,
		Address: tmPubKey.Address(),
		PubKey:  tmPubKey,
		Power:   sdk.NewInt(power).Mul(scalartypes.PowerReduction).Int64(),
	}
	// validatorInfo.NodeAccount = authtypes.NewBaseAccount(nodeAddr, nil, 0, 0)
	// validatorInfo.BroadcasterAccount = authtypes.NewBaseAccount(sdk.AccAddress(validatorInfo.Broadcaster.Address()), validatorInfo.Broadcaster, 0, 0)

	//Create a self delegation message for validator
	createValMsg, err := stakingtypes.NewMsgCreateValidator(
		sdk.ValAddress(senderAddress),
		validatorInfo.ValPubKey,
		valCoin,
		stakingtypes.NewDescription(nodeDirName, "", "", "", ""),
		stakingtypes.NewCommissionRates(sdk.OneDec(), sdk.OneDec(), sdk.OneDec()),
		sdk.OneInt(),
	)
	if err != nil {
		return nil, err
	}

	txBuilder := clientCtx.TxConfig.NewTxBuilder()
	if err := txBuilder.SetMsgs(createValMsg); err != nil {
		return nil, err
	}

	minGasPrice := args.minGasPrice
	if sdk.NewDecFromInt(args.baseFee).GT(args.minGasPrice) {
		minGasPrice = sdk.NewDecFromInt(args.baseFee)
	}

	txBuilder.SetMemo(validatorInfo.SeedAddress)
	txBuilder.SetGasLimit(createValidatorMsgGasLimit)
	txBuilder.SetFeeAmount(sdk.NewCoins(sdk.NewCoin(scalartypes.BaseDenom, minGasPrice.MulInt64(createValidatorMsgGasLimit).Ceil().TruncateInt())))

	txFactory := tx.Factory{}
	txFactory = txFactory.
		WithChainID(args.chainID).
		WithMemo(validatorInfo.SeedAddress).
		WithKeybase(kb).
		WithTxConfig(clientCtx.TxConfig)

	if err := tx.Sign(txFactory, senderKeyName, txBuilder, true); err != nil {
		return nil, err
	}
	// if err := tx.SignWithPrivKey(txFactory, senderKeyName, txBuilder, true); err != nil {
	// 	return nil, err
	// }
	txBz, err := clientCtx.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	if err := utils.WriteFile(fmt.Sprintf("%v.json", nodeDirName), gentxsDir, txBz); err != nil {
		return nil, err
	}
	// Add custom app config
	if err := setCustomAppConfig(cmd); err != nil {
		return nil, err
	}
	//Generate cosmos default app config
	appConfig := sdkconfig.DefaultConfig()
	appConfig.MinGasPrices = args.minGasPrices
	appConfig.API.Enable = true
	appConfig.Telemetry.Enabled = true
	appConfig.Telemetry.PrometheusRetentionTime = 60
	appConfig.Telemetry.EnableHostnameLabel = false
	appConfig.Telemetry.GlobalLabels = [][]string{{"chain_id", args.chainID}}
	sdkconfig.WriteConfigFile(filepath.Join(nodeDir, "config/app.toml"), appConfig)
	// Generate tendermint default config
	configPath := filepath.Join(nodeDir, "config", "config.toml")
	tmconfig.WriteConfigFile(configPath, nodeConfig)
	err = appendBridgeConfig(configPath, args.supportedChains)
	if err != nil {
		log.Error().Err(err).Msg("Failed to append bridge config")
		return nil, err
	}
	// ReadFile(configPath)
	return &validatorInfo, nil
}

func appendBridgeConfig(configPath string, supportedChainsPath string) error {
	//log.Info().Str("configPath", configPath).Str("supportedChainsPath", supportedChainsPath).Msg("Appending bridge config")
	file, err := os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Error().Err(err).Str("configPath", configPath).Msg("Could not open config file")
		return err
	}

	defer file.Close()

	_, err = file.WriteString(`
#######################################################
###         Bridge Configuration Options            ###
#######################################################
	`)

	if err != nil {
		log.Error().Err(err).Str("configPath", configPath).Msg("Could not write text to config file")
		return err
	}

	if supportedChainsPath != "" {
		// Add evm bridge config
		evmConfigs, err := scalartypes.ParseJsonArrayConfig[scalartypes.EvmNetworkConfig](fmt.Sprintf("%s/evm.json", supportedChainsPath))
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse evm config")
		}
		for _, evmConfig := range evmConfigs {
			//Todo: change bridge config to scalar_bridge_evm if rewrite vald module
			//https://github.com/axelarnetwork/axelar-core/blob/main/vald/config/config.go#L24
			_, err = file.WriteString(fmt.Sprintf(`
[[axelar_bridge_evm]]
id = "%s"
chain_id = %d
rpc_addr = "%s"
start-with-bridge = true
finality_override = "confirmation"
# When using new chains (not Ethereum Mainnet), you may need to set the finality override to "confirmation" to avoid issues with the bridge
# With finality override, scalar will create evm client using ethereum.go, not ethereum_2.go
			`, evmConfig.ID, evmConfig.ChainID, evmConfig.RpcUrl))
		}
		if err != nil {
			log.Error().Err(err).Msg("Failed to write evm bridge config")
		}
		// btcConfigs, err := scalartypes.ParseJsonArrayConfig[scalartypes.BtcNetworkConfig](fmt.Sprintf("%s/btc.json", supportedChainsPath))
		// if err != nil {
		// 	log.Error().Err(err).Msg("Failed to parse btc config")
		// }
		//Add btc bridge config after implement btc module
		// 		for _, btcConfig := range btcConfigs {
		// 			_, err = file.WriteString(fmt.Sprintf(`
		// [[scalar_bridge_btc]]
		// id = "%s"
		// chain_id = %d
		// name = "%s"
		// Host = "%s"
		// Port = %d
		// User = "%s"
		// Pass = "%s"
		// DisableTLS = true
		// DisableConnectOnNew = true
		// DisableAutoReconnect = false
		// HTTPPostMode = true
		// 			`, btcConfig.ID, btcConfig.ChainID, btcConfig.Name, btcConfig.RpcHost, btcConfig.RpcPort, btcConfig.RpcUser, btcConfig.RpcPass))
		// 		}
		// 		if err != nil {
		// 			log.Error().Err(err).Msg("Failed to write btc bridge config")
		// 		}
	}
	return nil
}

func setCustomAppConfig(cmd *cobra.Command) error {
	// Todo: add custom app config if needed
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
	supportedChainsPath string,
	validatorInfos []scalartypes.ValidatorInfo,
	baseFee sdk.Int,
	minGasPrice sdk.Dec,
) error {
	appGenState, err := scalartypes.GenerateGenesis(clientCtx, mbm, coinDenom, validatorInfos, supportedChainsPath)
	if err != nil {
		fmt.Printf("GenerateGenesis err: %s\n", err.Error())
		return err
	}

	appGenStateJSON, err := json.MarshalIndent(appGenState, "", "  ")
	if err != nil {
		fmt.Printf("MarshalIndent err: %s\n", err.Error())
		return err
	}
	validators := make([]tmtypes.GenesisValidator, len(validatorInfos))
	for i, validatorInfo := range validatorInfos {
		validators[i] = validatorInfo.GenesisValidator
		fmt.Printf("Validator: power: %d; pubkey: %v, address: %s\n", validators[i].Power, hex.EncodeToString(validators[i].PubKey.Bytes()), sdk.AccAddress(validatorInfo.ValPubKey.Address()).String())
	}
	genDoc := tmtypes.GenesisDoc{
		ChainID:    chainID,
		AppState:   appGenStateJSON,
		Validators: validators,
	}

	// generate empty genesis files for each validator and save
	for i := 0; i < len(validatorInfos); i++ {
		if err := genDoc.SaveAs(validatorInfos[i].GenFile); err != nil {
			fmt.Printf("Save genDoc Err: %s\n", err.Error())
			return err
		} else {
			fmt.Printf("genDoc successfully generated to %s\n", validatorInfos[i].GenFile)
		}
		//Write seed file
		seeds := []string{}
		for index, validatorInfo := range validatorInfos {
			if validatorInfo.NodeID != validatorInfos[i].NodeID {
				seed := fmt.Sprintf(`[[seed]]
name = "validator%d"
address = "%s"
`, index+1, validatorInfo.SeedAddress)
				seeds = append(seeds, seed)
			}
		}
		utils.WriteFile("seed.toml", validatorInfos[i].NodeDir, []byte(strings.Join(seeds, "\n")))

	}
	return nil
}

func collectGenFiles(
	clientCtx client.Context, nodeConfig *tmconfig.Config, chainID string,
	validatorInfos []scalartypes.ValidatorInfo,
	outputDir string, genBalIterator banktypes.GenesisBalancesIterator,
) error {
	var appState json.RawMessage
	genTime := tmtime.Now()
	validators := make([]tmtypes.GenesisValidator, len(validatorInfos))
	for i, validatorInfo := range validatorInfos {
		validators[i] = validatorInfo.GenesisValidator
	}
	for _, validatorInfo := range validatorInfos {
		gentxsDir := filepath.Join(outputDir, "gentxs")
		initCfg := genutiltypes.NewInitConfig(chainID, gentxsDir, validatorInfo.NodeID, validatorInfo.ValPubKey)

		genDoc, err := tmtypes.GenesisDocFromFile(validatorInfo.GenFile)
		if err != nil {
			fmt.Printf("GenesisDocFromFile Err: %s\n", err.Error())
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
		// overwrite each validator's genesis file to have a canonical genesis time
		if err := genutil.ExportGenesisFileWithTime(validatorInfo.GenFile, chainID, validators, appState, genTime); err != nil {
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

func ReadFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Error().Err(err).Msg("failed reading data from file")
	}
	fmt.Printf("\nLength: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
	fmt.Printf("\nError: %v", err)
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
