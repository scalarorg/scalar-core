package testnet

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cometbft/cometbft/p2p"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	hd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdksecp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkserver "github.com/cosmos/cosmos-sdk/server"
	sdkconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/go-bip39"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/params"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/cmd/scalard/cmd/utils"
	"github.com/scalarorg/scalar-core/vald/config"
	scalarnetexported "github.com/scalarorg/scalar-core/x/scalarnet/exported"
	"github.com/tendermint/tendermint/privval"

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
	flagScalarMnemonic      = "SCALAR_MNEMONIC"
	flagProtocolBtcPriKey   = "PROTOCOL_BTC_PRIKEY"
	flagValidatorMnemonic   = "VALIDATOR_MNEMONIC"
	flagBroadcasterMnemonic = "BROADCASTER_MNEMONIC"
	flagGovernanceMnemonic  = "GOV_MNEMONIC"
	flagFaucetMnemonic      = "FAUCET_MNEMONIC"
	flagBtcPubkey           = "BTC_PUBKEY"
	flagBtcPrivkey          = "BTC_PRIVKEY"
	flagNodeDirPrefix       = "node-dir-prefix"
	flagNumValidators       = "v"
	flagNumProtocols        = "p"
	flagConfigPath          = "config-path"
	flagOutputDir           = "output-dir"
	flagBaseDir             = "base-dir"
	flagTimeout             = "timeout"
	flagBlockHeight         = "block-height"
	flagNodeDaemonHome      = "node-daemon-home"
	flagNodeDomain          = "node-domain"
	flagPortOffset          = "port-offset"
	flagBaseFee             = "base-fee"
	flagMinGasPrice         = "min-gas-price"
	flagKeyType             = "key-type"
	flagEnvFile             = "env-file"
	flagTag                 = "tag"
	flagVersion             = "version"
	flagEnableLogging       = "enable-logging"
	flagRPCAddress          = "rpc.address"
	flagAPIAddress          = "api.address"
	flagPrintMnemonic       = "print-mnemonic"
	flagGRPCAddress         = "grpc.address"
	flagJSONRPCAddress      = "json-rpc.address"
)

type initArgs struct {
	algo           string
	chainID        string
	tag            []byte
	version        uint32
	keyringBackend string
	minGasPrices   string
	nodeDaemonHome string
	configPath     string
	nodeDirPrefix  string
	numValidators  int
	numProtocols   int
	outputDir      string
	nodeDomain     string
	portOffset     int
	baseFee        sdk.Int
	minGasPrice    sdk.Dec
}

type startArgs struct {
	baseDir       string
	numValidators int
	timeout       int
	blockHeight   int
}

type EnvKeys struct {
	ValidatorMnemonic   string
	BroadcasterMnemonic string
	GovernanceMnemonic  string
	FaucetMnemonic      string
	BtcPubkey           string
	BtcPrivkey          string
}

// createValidatorMsgGasLimit is the gas limit used in the MsgCreateValidator included in genesis transactions.
// This transaction consumes approximately 220,000 gas when executed in the genesis block.
const createValidatorMsgGasLimit = 250_000
const GenesisAsset = scalarnetexported.BaseAsset

var (
	ValidatorCoin   = sdk.NewCoin(GenesisAsset, sdk.NewIntWithDecimal(1, sdk.Precision).Mul(ValidatorStaking))
	BroadcasterCoin = sdk.NewCoin(GenesisAsset, sdk.NewIntWithDecimal(1, sdk.Precision).Mul(BroadcasterTokens))
	ScalarCoin      = sdk.NewCoin(GenesisAsset, sdk.NewIntWithDecimal(1, sdk.Precision).Mul(ScalarTokens))
	GovCoin         = sdk.NewCoin(GenesisAsset, sdk.NewIntWithDecimal(1, sdk.Precision).Mul(GovTokens))
	FaucetCoin      = sdk.NewCoin(GenesisAsset, sdk.NewIntWithDecimal(1, sdk.Precision).Mul(FaucetTokens))
)

func addTestnetFlagsToCmd(cmd *cobra.Command) {
	cmd.Flags().Int(flagNumValidators, 4, "Number of validators to initialize the testnet with")
	cmd.Flags().StringP(flagOutputDir, "o", "./.testnets", "Directory to store initialization data for the testnet")
	cmd.Flags().String(flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().String(sdkserver.FlagMinGasPrices, fmt.Sprintf("0.007%s", scalarnetexported.BaseAsset), "Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01photino,0.001stake)")
	cmd.Flags().String(flagKeyType, string(hd.Secp256k1Type), "Key signing algorithm to generate keys for")
	cmd.Flags().String(flagBaseFee, strconv.Itoa(params.InitialBaseFee), "The params base_fee in the feemarket module in geneis")
	cmd.Flags().String(flagMinGasPrice, "0.001", "The params min_gas_price in the feemarket module in geneis")
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
	testnetCmd.AddCommand(testnetStartCmd())

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
	scalard testnet init-files --v 4 --output-dir ./.testnets --node-domain scalarnode --supported-chains=./chains --env-file=.env.local
	`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// Get the env file flag
			if envFile, _ := cmd.Flags().GetString(flagEnvFile); envFile != "" {
				if err := loadEnvFile(envFile); err != nil {
					return err
				}
			}
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
			args.numProtocols, _ = cmd.Flags().GetInt(flagNumProtocols)
			args.portOffset, _ = cmd.Flags().GetInt(flagPortOffset)
			args.configPath, _ = cmd.Flags().GetString(flagConfigPath)
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
			readEnvs(&args)
			//End Test keyring
			return initTestnetFiles(clientCtx, cmd, serverCtx.Config, mbm, genBalIterator, args)
		},
	}

	addTestnetFlagsToCmd(cmd)
	cmd.Flags().Int(flagPortOffset, 0, "Port offset for the testnet")
	cmd.Flags().String(flagNodeDirPrefix, "node", "Prefix the directory name for each node with (node results in node1, node2, ...)")
	cmd.Flags().String(flagNodeDaemonHome, "scalard", "Home directory of the node's daemon configuration")
	cmd.Flags().String(flagNodeDomain, "", `Node domain: 
		*scalarnode* results in persistent peers list ID0@scalarnode1:46656, ID1@scalarnode2:46656, ...
		*192.168.0.1* results in persistent peers list ID0@192.168.0.11:46656, ID1@192.168.0.12:46656, ...
		`)
	cmd.Flags().String(flagConfigPath, "./configs", `Configs directory, keep supported chains in each family file such as evm.json or btc.json: 
		*./chains/evm.json* stores all evm chain configs ...
		*./chains/btc.json* stores all btc chain configs ...
		*./protocols.json* stores all protocol configs ...
		*./tokens/evm.json* stores all evm erc20 token configs ...
		*./tokens/btc.json* stores all btc token configs ...
		`)
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	cmd.Flags().String(flagEnvFile, "", "Path to environment file to load (optional)")

	return cmd
}
func readEnvs(args *initArgs) {
	args.tag = []byte(os.Getenv("TAG"))
	version, err := strconv.ParseUint(os.Getenv("VERSION"), 10, 32)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse version")
	}
	args.version = uint32(version)
}

// get cmd to start multi validator in-process testnet
func testnetStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Launch an in-process multi-validator testnet",
		Long: `testnet will launch an in-process multi-validator testnet,
and generate "v" directories, populated with necessary validator configuration files
(private validator, genesis, config, etc.).

Example:
	scalard testnet start --v 4 --base-dir ./.testnets
	`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			args := startArgs{}
			args.baseDir, _ = cmd.Flags().GetString(flagBaseDir)
			args.numValidators, _ = cmd.Flags().GetInt(flagNumValidators)
			args.timeout, _ = cmd.Flags().GetInt(flagTimeout)
			args.blockHeight, _ = cmd.Flags().GetInt(flagBlockHeight)
			return startTestnet(cmd, args)
		},
	}

	addTestnetFlagsToCmd(cmd)
	cmd.Flags().String(flagBaseDir, "./.testnets", "the base directory to store the testnet")
	cmd.Flags().Int(flagTimeout, 1800, "The testnet run time. Default is 1800 seconds")
	cmd.Flags().Int(flagBlockHeight, 100, "The block height to stop the testnet")
	return cmd
}

const nodeDirPerm = 0o755

func readEnvKeys(index int) EnvKeys {
	validatorMnemonic := getNonQuoteEnv(flagValidatorMnemonic + strconv.Itoa(index))
	if validatorMnemonic == "" {
		validatorMnemonic = getNonQuoteEnv(flagValidatorMnemonic)
	}
	btcPubKey := os.Getenv(flagBtcPubkey + strconv.Itoa(index))
	if btcPubKey == "" {
		btcPubKey = os.Getenv(flagBtcPubkey)
	}

	btcPrivkey := os.Getenv(flagBtcPrivkey + strconv.Itoa(index))
	if btcPrivkey == "" {
		btcPrivkey = os.Getenv(flagBtcPrivkey)
	}
	envKeys := EnvKeys{
		ValidatorMnemonic:   validatorMnemonic,
		BroadcasterMnemonic: getNonQuoteEnv(flagBroadcasterMnemonic),
		GovernanceMnemonic:  getNonQuoteEnv(flagGovernanceMnemonic),
		FaucetMnemonic:      getNonQuoteEnv(flagFaucetMnemonic),
		BtcPubkey:           btcPubKey,
		BtcPrivkey:          btcPrivkey,
	}
	log.Debug().Any("EnvKeys", envKeys).Msg("Environment variables")
	return envKeys
}

// Get environment variable then remove start and end quote '"'
func getNonQuoteEnv(key string) string {
	value := os.Getenv(key)
	if value[0] == '"' {
		value = value[1:]
	}
	if i := len(value) - 1; value[i] == '"' {
		value = value[:i]
	}
	return value
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
	var (
		validatorInfos []ValidatorInfo
	)
	relayer := initRelayer(args)

	protocols := initProtocols(args)

	// generate private keys, node IDs, and initial transactions
	for i := 0; i < args.numValidators; i++ {
		envKeys := readEnvKeys(i + 1)
		nodeDirName := getNodeDirName(i, args.nodeDirPrefix)
		host, err := getHost(i, args.nodeDomain)
		if err != nil {
			_ = os.RemoveAll(args.outputDir)
			return err
		}
		// Validator index starts from 1
		validatorInfo, err := initValidatorConfig(clientCtx, cmd, nodeConfig, host, nodeDirName, args, envKeys, i)
		if err != nil {
			_ = os.RemoveAll(args.outputDir)
			cmd.PrintErrf("failed to initialize validator config: %s", err.Error())
			return err
		}
		validatorInfos = append(validatorInfos, *validatorInfo)
	}

	if err := generateFiles(clientCtx, mbm, nodeConfig, relayer, validatorInfos, protocols, args, genBalIterator); err != nil {
		cmd.PrintErrf("failed to initGenFiles: %s", err.Error())
		return err
	}

	// if err := initGenFiles(clientCtx, mbm,
	// 	nodeConfig,
	// 	GenesisAsset,
	// 	validatorInfos,
	// 	args,
	// ); err != nil {
	// 	cmd.PrintErrf("failed to initGenFiles: %s", err.Error())
	// 	return err
	// }

	// err := collectGenFiles(clientCtx, nodeConfig, args.chainID, validatorInfos, args.outputDir, genBalIterator)
	// if err != nil {
	// 	cmd.PrintErrf("failed to collect genesis files: %s", err.Error())
	// 	return err
	// }

	cmd.PrintErrf("Successfully initialized %d node directories\n", args.numValidators)
	return nil
}
func initRelayer(args initArgs) ScalarRelayer {
	scalarMnemonic := getNonQuoteEnv(flagScalarMnemonic)
	privKey, address, err := createScalarAccount(scalarMnemonic)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create scalar account")
	}
	return ScalarRelayer{
		PubKey: privKey.PubKey(),
		Balance: banktypes.Balance{
			Address: address.String(),
			Coins:   sdk.Coins{ScalarCoin},
		},
	}
}
func initProtocols(args initArgs) []Protocol {
	protocolConfigs, err := ParseJsonArrayConfig[ProtocolConfig](fmt.Sprintf("%s/protocols.json", args.configPath))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse chains config")
	}
	kb, algo, err := createKeyring(bufio.NewReader(os.Stdin), args, "/tmp")
	if err != nil {
		log.Error().Err(err).Msg("Failed to create keyring")
	}
	protocols := make([]Protocol, len(protocolConfigs))
	for i, config := range protocolConfigs {
		name := fmt.Sprintf("protocol-%d", i)
		protocols[i] = Protocol{
			Tag:            config.Tag,
			LiquidityModel: config.LiquidityModel,
		}
		if config.ScalarMnemonic != "" {
			//Create privKey and address of protocol by keyring algorithm
			pubkey, address, err := generateAccount(kb, algo, name, config.ScalarMnemonic, "")
			if err != nil {
				log.Debug().Err(err).Msg("Create scalar account with error")
			}
			protocols[i].PubKey = pubkey
			// privKey, address, err := createScalarAccount(config.ScalarMnemonic)
			// if err != nil {
			// 	log.Debug().Err(err).Msg("Create scalar account with error")
			// }
			// protocols[i].PubKey = privKey.PubKey()

			protocols[i].Balance = banktypes.Balance{
				Address: address.String(),
				Coins:   sdk.Coins{ScalarCoin},
			}
			log.Debug().Str("ProtocolMnemonic", config.ScalarMnemonic).Str("Account", address.String()).Msg("ScalarAccount")
		}
		if config.BitcoinPrivKey != "" {
			privKey := secp256k1.PrivKeyFromBytes([]byte(config.BitcoinPrivKey))
			protocols[i].BitcoinPubKey = privKey.PubKey().SerializeCompressed()
		}
	}
	return protocols
}

func createPubkeyFromSecret(config *tmconfig.Config, secret []byte, pvKeyName string) (cryptotypes.PubKey, error) {
	privKey := tmed25519.GenPrivKeyFromSecret(secret)
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
		pvKeyName = ValidatorKeyName
	} else {
		pvKeyFile = filepath.Join(config.RootDir, "config", fmt.Sprintf("%s_key.json", pvKeyName))
		pvStateFile = filepath.Join(config.RootDir, "data", fmt.Sprintf("%s_state.json", pvKeyName))
	}

	filePV := privval.NewFilePV(privKey, pvKeyFile, pvStateFile)
	if err := tmos.EnsureDir(filepath.Dir(pvKeyFile), 0o777); err != nil {
		return nil, err
	}
	filePV.Save()
	fmt.Printf("Private key saved to file: %s\n", pvKeyFile)
	valPubKey, err := cryptocodec.FromTmPubKeyInterface(privKey.PubKey())
	if err != nil {
		return nil, fmt.Errorf("failed to convert tmtypes.Pubkey to cryptotypes.PubKey: %w", err)
	}
	infoPath := filepath.Join(config.RootDir, fmt.Sprintf("%s.json", pvKeyName))
	log.Info().Msgf("Private key saved to file: %s\n", infoPath)
	storeValidatorInfo(valPubKey, pvKeyName, config.RootDir)
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

func createKeyringAccountFromMnemonic(keybase keyring.Keyring,
	keyName string,
	mnemonic string,
	algo keyring.SignatureAlgo,
	bip44Path string,
) (cryptotypes.PubKey, sdk.AccAddress, error) {
	if bip44Path == "" {
		bip44Path = "m/44'/118'/0'/0/0"
	}
	info, err := keybase.NewAccount(
		keyName,
		mnemonic,
		bip44Path,
		keyring.DefaultBIP39Passphrase,
		algo,
	)
	if err != nil {
		log.Error().Err(err).Str("Mnemonic", mnemonic).Str("bip44path", bip44Path).Msg("[createKeyringAccountFromMnemonic] NewAccount error")
		info, err = keybase.Key(keyName)
		if err != nil {
			log.Error().Err(err).Msg("[createKeyringAccountFromMnemonic] Get existing key error")
			return nil, nil, err
		}
	}

	ko, err := keyring.MkAccKeyOutput(info)
	if err != nil {
		log.Error().Err(err).Msg("[createKeyringAccountFromMnemonic] MkAccKeyOutput")
		return nil, nil, err
	}
	log.Debug().Str("keyName", keyName).Msgf("MkAccKeyOutput: %v", ko)
	return info.GetPubKey(), info.GetAddress(), nil
}
func storeValidatorAddress(valAddress sdk.ValAddress, keyName string, nodeDir string) error {
	if err := utils.WriteFile(fmt.Sprintf("%v.json", keyName), nodeDir, []byte(valAddress.String())); err != nil {
		return err
	}
	return nil
}
func storeValidatorInfo(pubkey cryptotypes.PubKey, keyName string, nodeDir string) error {
	address := sdk.ValAddress(pubkey.Address())
	info := map[string]string{
		"address": address.String(),
		"pubkey":  pubkey.String(),
	}

	cliPrint, err := json.Marshal(info)
	if err != nil {
		return err
	}
	// save private key seed words
	if err := utils.WriteFile(fmt.Sprintf("%v.json", keyName), nodeDir, cliPrint); err != nil {
		return err
	}
	return nil
}
func createKeyring(inBuf *bufio.Reader, args initArgs, nodeDir string) (keyring.Keyring, keyring.SignatureAlgo, error) {
	log.Debug().Str("keyringBackend", args.keyringBackend).Str("nodeDir", nodeDir).Str("keyringServiceName", sdk.KeyringServiceName()).Msg("Create keyring")
	kb, err := keyring.New(sdk.KeyringServiceName(), args.keyringBackend, nodeDir, inBuf, DefaultOption)
	if err != nil {
		return nil, nil, err
	}

	keyringAlgos, _ := kb.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(args.algo, keyringAlgos)
	if err != nil {
		return nil, nil, err
	}
	return kb, algo, nil
}
func genFaucet(kb keyring.Keyring, mnemonic string, algo keyring.SignatureAlgo, tokenAmount sdk.Int) (*banktypes.Balance, error) {
	if mnemonic != "" {
		bip44Path := fmt.Sprintf("m/%d'/%d'/0'/0/0", PurposeFaucetAccount, 0)
		_, address, err := createKeyringAccountFromMnemonic(kb,
			BroadcasterKeyName,
			mnemonic,
			algo,
			bip44Path,
		)
		if err != nil {
			log.Error().Err(err).Msg("[getFaucet] Create keyring account from mnemonic")
			return nil, err
		}
		return &banktypes.Balance{
			Address: address.String(),
			Coins:   sdk.Coins{sdk.NewCoin(GenesisAsset, tokenAmount)},
		}, nil
	}
	return nil, nil
}
func initValidatorConfig(clientCtx client.Context, cmd *cobra.Command,
	nodeConfig *tmconfig.Config,
	host string,
	nodeDirName string,
	args initArgs,
	envKeys EnvKeys,
	index int, //index starts from 0
) (*ValidatorInfo, error) {
	var err error
	nodeDir := filepath.Join(args.outputDir, nodeDirName, args.nodeDaemonHome)
	nodeConfig.SetRoot(nodeDir)
	if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
		return nil, err
	}
	fmt.Printf("Create validator config in dir %s\n", nodeDir)
	nodeConfig.Moniker = nodeDirName
	// TaiVV: 24 dec 12
	// generate ed25519 validator pubkey using random mnemonic
	// If set input mnemonic the, need explicitly write the priv_validator_key.json
	valNodeID, valNodePubKey, err := genutil.InitializeNodeValidatorFiles(nodeConfig)
	log.Debug().Str("valNodeID", valNodeID).Str("valNodePubKey", valNodePubKey.String()).Msg("InitializeNodeValidatorFilesFromMnemonic")
	//nodeID, err := createNodeID(nodeConfig)
	if err != nil {
		return nil, err
	}
	validatorInfo := ValidatorInfo{
		Host:          host,
		NodeID:        valNodeID,
		SeedAddress:   fmt.Sprintf("%s@%s:%d", valNodeID, host, 26656+index*args.portOffset),
		RPCPort:       26657 + index*args.portOffset,
		Moniker:       nodeDirName,
		NodeDir:       nodeDir,
		GenFile:       nodeConfig.GenesisFile(),
		ValNodePubKey: valNodePubKey,
		ValNodeBalance: banktypes.Balance{
			Address: sdk.AccAddress(valNodePubKey.Address()).String(),
			Coins:   sdk.Coins{ValidatorCoin},
		},
		AdditionalKeys: config.AdditionalKeys{
			BtcPrivKey: envKeys.BtcPrivkey,
		},
	}

	gentxsDir := filepath.Join(args.outputDir, "gentxs")
	// TODO: add ledger support
	kb, algo, err := createKeyring(bufio.NewReader(cmd.InOrStdin()), args, nodeDir)
	if err != nil {
		return nil, err
	}
	//This account is used to sign the MsgCreateValidator
	valPubKey, valAddress, err := createKeyringAccountFromMnemonic(kb,
		ValidatorKeyName,
		envKeys.ValidatorMnemonic,
		algo,
		fmt.Sprintf("m/%d'/%d'/0'/0/0",
			PurposeValidator, uint32(index)))
	if err != nil {
		log.Error().Err(err).Msg("[initValidatorConfig] Create validator account from Mnemonic")
		key, err := kb.Key(ValidatorKeyName)
		if err != nil {
			log.Error().Err(err).Msg("[initValidatorConfig] Get faucet keyring account")
			return nil, err
		}
		valPubKey = key.GetPubKey()
		valAddress = key.GetAddress()
	}
	storeValidatorAddress(sdk.ValAddress(valPubKey.Address()), "validator_address", nodeDir)
	validatorInfo.ValPubKey = valPubKey
	validatorInfo.ValBalance = banktypes.Balance{
		Address: valAddress.String(),
		Coins:   sdk.Coins{ValidatorCoin},
	}
	if envKeys.BroadcasterMnemonic != "" {
		//broadcasterPubKey, err := createPubkeyFromMnemonic(nodeConfig, envKeys.BroadcasterMnemonic, kb, algo, BroadcasterKeyName)
		bip44Path := fmt.Sprintf("m/%d'/%d'/0'/0/0", PurposeBroadcaster, uint32(index))
		pubkey, address, err := generateAccount(kb, algo, BroadcasterKeyName, envKeys.BroadcasterMnemonic, bip44Path)
		if err != nil {
			log.Debug().Err(err).Msg("Generate account error")
		}
		validatorInfo.Broadcaster = pubkey
		validatorInfo.BroadcasterBalance = banktypes.Balance{
			Address: address.String(),
			Coins:   sdk.Coins{BroadcasterCoin},
		}
	}
	if envKeys.GovernanceMnemonic != "" {
		//validatorInfo.GovPubKey, err = createPubkeyFromMnemonic(nodeConfig, envKeys.GovernanceMnemonic, kb, algo, GovKeyName)
		bip44Path := fmt.Sprintf("m/%d'/%d'/0'/0/0", PurposeGovernance, uint32(index))
		pubkey, address, err := generateAccount(kb, algo, GovKeyName, envKeys.GovernanceMnemonic, bip44Path)
		if err != nil {
			log.Debug().Err(err).Msg("Generate account error")
		}
		validatorInfo.GovPubKey = pubkey
		validatorInfo.GovBalance = banktypes.Balance{
			Address: address.String(),
			Coins:   sdk.Coins{FaucetCoin},
		}
	}
	if envKeys.FaucetMnemonic != "" {
		pubKey, address, err := generateAccount(kb, algo,
			FaucetKeyName,
			envKeys.FaucetMnemonic,
			fmt.Sprintf("m/%d'/%d'/0'/0/0", PurposeFaucetAccount, uint32(index)),
		)
		if err != nil {
			log.Debug().Err(err).Msg("Generate account error")
		}
		validatorInfo.FaucetPubKey = pubKey
		validatorInfo.FaucetBalance = banktypes.Balance{
			Address: address.String(),
			Coins:   sdk.Coins{FaucetCoin},
		}
	}

	valPower := int64((index + 1) * (index + 1) * 1000000)
	stakingPower := sdk.NewCoin(GenesisAsset, sdk.TokensFromConsensusPower(valPower, ValidatorStaking))
	senderKeyName := ValidatorKeyName
	tmPubKey, err := cryptocodec.ToTmPubKeyInterface(validatorInfo.ValNodePubKey)
	if err != nil {
		fmt.Printf("ToTmPubKeyInterface Err: %s\n", err.Error())
		return nil, err
	}
	validatorInfo.GenesisValidator = tmtypes.GenesisValidator{
		Name:    nodeDirName,
		Address: tmPubKey.Address(),
		PubKey:  tmPubKey,
		Power:   valPower,
		//Power:   sdk.NewInt(valPower).Mul(PowerReduction).Int64(),
	}
	// validatorInfo.NodeAccount = authtypes.NewBaseAccount(nodeAddr, nil, 0, 0)
	// validatorInfo.BroadcasterAccount = authtypes.NewBaseAccount(sdk.AccAddress(validatorInfo.Broadcaster.Address()), validatorInfo.Broadcaster, 0, 0)
	//Create a self delegation message for validator
	createValMsg, err := stakingtypes.NewMsgCreateValidator(
		sdk.ValAddress(valAddress),
		validatorInfo.ValNodePubKey,
		stakingPower,
		stakingtypes.NewDescription(nodeDirName, validatorInfo.NodeID, "", "", ""),
		stakingtypes.NewCommissionRates(sdk.OneDec(), sdk.OneDec(), sdk.OneDec()),
		sdk.OneInt(),
	)
	log.Debug().
		Str("valPubKey", validatorInfo.ValNodePubKey.String()).
		Str("valAddress", valAddress.String()).
		Str("deligatorAddress", sdk.AccAddress(valAddress).String()).
		Str("valStaking", stakingPower.String()).
		Str("senderKeyName", senderKeyName).Msg("MsgCreateValidator")
	if err != nil {
		return nil, err
	}

	txBuilder := clientCtx.TxConfig.NewTxBuilder()
	if err := txBuilder.SetMsgs(createValMsg); err != nil {
		return nil, err
	}

	minGasPrice := args.minGasPrice
	// if sdk.NewDecFromInt(args.baseFee).GT(args.minGasPrice) {
	// 	minGasPrice = sdk.NewDecFromInt(args.baseFee)
	// }
	feeAmount := sdk.NewCoin(scalarnetexported.BaseAsset, minGasPrice.MulInt64(createValidatorMsgGasLimit).Ceil().TruncateInt())
	txBuilder.SetMemo(validatorInfo.SeedAddress)
	txBuilder.SetGasLimit(createValidatorMsgGasLimit)
	txBuilder.SetFeeAmount(sdk.NewCoins(feeAmount))
	log.Debug().
		Str("minGasPrice", minGasPrice.String()).
		Str("createValidatorMsgGasLimit", strconv.FormatInt(createValidatorMsgGasLimit, 10)).
		Any("feeAmount", feeAmount).Msg("SetFeeAmount")
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
	return &validatorInfo, nil
}

// Create scalar account for using in relayer
func createScalarAccount(mnemonic string) (*sdksecp256k1.PrivKey, sdk.AccAddress, error) {
	// Derive the seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "")
	path := "m/44'/118'/0'/0/0"
	// Create master key and derive the private key
	// Using "m/44'/118'/0'/0/0" for Cosmos
	master, ch := hd.ComputeMastersFromSeed(seed)
	privKeyBytes, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		return nil, nil, err
	}
	// Create private key and get address
	privKey := &sdksecp256k1.PrivKey{Key: privKeyBytes}
	addr := sdk.AccAddress(privKey.PubKey().Address())
	return privKey, addr, nil
}

func generateAccount(kr keyring.Keyring, algo keyring.SignatureAlgo, keyName string, mnemonic string, bip44Path string) (cryptotypes.PubKey, sdk.AccAddress, error) {
	pubkey, address, err := createKeyringAccountFromMnemonic(kr,
		keyName,
		mnemonic,
		algo,
		bip44Path,
	)
	//For testnet
	// inMemkr := keyring.NewInMemory()
	// info, err := inMemkr.NewAccount(keyName, mnemonic, "", bip44Path, algo)
	log.Debug().
		Str("keyName", keyName).
		Str("mnemonic", mnemonic).
		Str("bip44Path", bip44Path).
		Str("krAddress", address.String()).
		Str("pubkeyAddress", sdk.AccAddress(pubkey.Address()).String()).Msg("Keyring account created")
	if err != nil {
		log.Error().Err(err).Msg("[initValidatorConfig] Create governance keyring account from mnemonic")
		key, err := kr.Key(keyName)
		if err != nil {
			log.Error().Err(err).Msg("[initValidatorConfig] Get governance keyring account")
			return nil, nil, err
		}
		address = key.GetAddress()
		pubkey = key.GetPubKey()
	}
	return pubkey, address, nil
}
func generateFiles(clientCtx client.Context, mbm module.BasicManager, nodeConfig *tmconfig.Config,
	relayer ScalarRelayer, validatorInfos []ValidatorInfo, protocols []Protocol, args initArgs, genBalIterator banktypes.GenesisBalancesIterator,
) error {
	var appRawState json.RawMessage
	var err error
	validators := make([]tmtypes.GenesisValidator, len(validatorInfos))
	for i, validatorInfo := range validatorInfos {
		validators[i] = validatorInfo.GenesisValidator
	}
	appGenState, err := GenerateGenesis(clientCtx, mbm, GenesisAsset, relayer, validatorInfos, protocols, args)
	if err != nil {
		fmt.Printf("GenerateGenesis err: %s\n", err.Error())
		return err
	}

	appRawState, err = json.MarshalIndent(appGenState, "", "  ")
	if err != nil {
		fmt.Printf("MarshalIndent err: %s\n", err.Error())
		return err
	}
	genDoc := tmtypes.GenesisDoc{
		GenesisTime: tmtime.Now(),
		ChainID:     args.chainID,
		AppState:    appRawState,
		Validators:  validators,
	}
	for i, validatorInfo := range validatorInfos {
		gentxsDir := filepath.Join(args.outputDir, "gentxs")
		initCfg := genutiltypes.NewInitConfig(args.chainID, gentxsDir, validatorInfo.NodeID, validatorInfo.ValPubKey)
		//Seed file
		seeds, seedAddrs := createSeeds(validatorInfos, validatorInfo)
		utils.WriteFile("seeds.toml", filepath.Join(validatorInfos[i].NodeDir, "config"), []byte(strings.Join(seeds, "\n")))
		nodeConfig.Moniker = validatorInfos[i].Moniker
		nodeConfig.SetRoot(validatorInfos[i].NodeDir)
		nodeConfig.P2P.Seeds = strings.Join(seedAddrs, ",")
		setConfigParams(nodeConfig, validatorInfo, i, args)
		_, err := genutil.GenAppStateFromConfig(
			clientCtx.Codec, clientCtx.TxConfig,
			nodeConfig, initCfg, genDoc, genBalIterator)
		if err != nil {
			return err
		}
		createAppConfig(validatorInfo.NodeDir, args, i)
		configPath := filepath.Join(validatorInfo.NodeDir, "config/config.toml")
		appendBridgeConfig(configPath, args.configPath)
		appendAdditionalKeys(configPath, validatorInfo)
	}

	return nil
}
func createSeeds(validatorInfos []ValidatorInfo, validator ValidatorInfo) ([]string, []string) {
	seeds := []string{}
	seedAddrs := []string{}
	for index, validatorInfo := range validatorInfos {
		if validatorInfo.NodeID != validator.NodeID {
			seed := fmt.Sprintf(`[[seed]]
name = "validator%d"
address = "%s"
`, index+1, validatorInfo.SeedAddress)
			seeds = append(seeds, seed)
			seedAddrs = append(seedAddrs, validatorInfo.SeedAddress)
		}
	}
	return seeds, seedAddrs
}
func setConfigParams(nodeConfig *tmconfig.Config, validator ValidatorInfo, index int, args initArgs) {
	nodeConfig.ProxyApp = fmt.Sprintf("tcp://127.0.0.1:%d", 26658+index*args.portOffset)
	nodeConfig.PrivValidatorListenAddr = ""
	nodeConfig.RPC.ListenAddress = fmt.Sprintf("tcp://0.0.0.0:%d", 26657+index*args.portOffset)
	nodeConfig.RPC.PprofListenAddress = fmt.Sprintf("0.0.0.0:%d", 6060+index*args.portOffset)

	nodeConfig.P2P.ListenAddress = fmt.Sprintf("tcp://0.0.0.0:%d", 26656+index*args.portOffset)
	nodeConfig.P2P.ExternalAddress = fmt.Sprintf("%s:%d", validator.Host, 26656+index*args.portOffset)
}
func createAppConfig(nodeDir string, args initArgs, index int) {
	appConfig := sdkconfig.DefaultConfig()
	appConfig.MinGasPrices = args.minGasPrices
	appConfig.API.Enable = true
	appConfig.API.Address = fmt.Sprintf("tcp://0.0.0.0:%d", 1317+index*args.portOffset)
	appConfig.API.Swagger = true
	appConfig.API.EnableUnsafeCORS = true
	appConfig.GRPC.Address = fmt.Sprintf("0.0.0.0:%d", 9090+index*args.portOffset)
	appConfig.GRPCWeb.Address = fmt.Sprintf("0.0.0.0:%d", 9091+index*args.portOffset)
	appConfig.Telemetry.Enabled = true
	appConfig.Telemetry.PrometheusRetentionTime = 60
	appConfig.Telemetry.EnableHostnameLabel = false
	appConfig.Telemetry.GlobalLabels = [][]string{{"chain_id", args.chainID}}
	sdkconfig.WriteConfigFile(filepath.Join(nodeDir, "config/app.toml"), appConfig)
}
func createConfigFiles(nodeConfig *tmconfig.Config, nodeDir string, args initArgs, host string, index int) error {
	//Generate cosmos default app config
	appConfig := sdkconfig.DefaultConfig()
	appConfig.MinGasPrices = args.minGasPrices
	appConfig.API.Enable = true
	appConfig.API.Address = fmt.Sprintf("tcp://0.0.0.0:%d", 1317+index*args.portOffset)
	appConfig.GRPC.Address = fmt.Sprintf("0.0.0.0:%d", 9090+index*args.portOffset)
	appConfig.GRPCWeb.Address = fmt.Sprintf("0.0.0.0:%d", 9091+index*args.portOffset)
	appConfig.Telemetry.Enabled = true
	appConfig.Telemetry.PrometheusRetentionTime = 60
	appConfig.Telemetry.EnableHostnameLabel = false
	appConfig.Telemetry.GlobalLabels = [][]string{{"chain_id", args.chainID}}
	sdkconfig.WriteConfigFile(filepath.Join(nodeDir, "config/app.toml"), appConfig)
	// Generate tendermint default config
	//Set port to different values for each validator
	nodeConfig.ProxyApp = fmt.Sprintf("tcp://127.0.0.1:%d", 26658+index*args.portOffset)
	nodeConfig.PrivValidatorListenAddr = ""
	nodeConfig.RPC.ListenAddress = fmt.Sprintf("tcp://0.0.0.0:%d", 26657+index*args.portOffset)
	nodeConfig.RPC.PprofListenAddress = fmt.Sprintf("0.0.0.0:%d", 6060+index*args.portOffset)

	nodeConfig.P2P.ListenAddress = fmt.Sprintf("tcp://0.0.0.0:%d", 26656+index*args.portOffset)
	nodeConfig.P2P.ExternalAddress = fmt.Sprintf("%s:%d", host, 26656+index*args.portOffset)
	//configPath := filepath.Join(nodeDir, "config/config.toml")
	//
	//tmconfig.WriteConfigFile(configPath, nodeConfig)
	// err := appendBridgeConfig(configPath, args.supportedChains)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Failed to append bridge config")
	// }
	return nil
}
func appendBridgeConfig(bridgeConfigPath string, configPath string) error {
	//log.Info().Str("configPath", configPath).Str("supportedChainsPath", supportedChainsPath).Msg("Appending bridge config")
	file, err := os.OpenFile(bridgeConfigPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Error().Err(err).Str("configPath", bridgeConfigPath).Msg("Could not open config file")
		return err
	}

	defer file.Close()

	_, err = file.WriteString(`
#######################################################
###         Bridge Configuration Options            ###
#######################################################
	`)

	if err != nil {
		log.Error().Err(err).Str("configPath", bridgeConfigPath).Msg("Could not write text to config file")
		return err
	}

	if configPath != "" {
		// Add evm bridge config
		evmConfigs, err := ParseJsonArrayConfig[config.EVMConfig](fmt.Sprintf("%s/chains/evm.json", configPath))
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse evm config")
		}
		for _, evmConfig := range evmConfigs {
			//Todo: change bridge config to scalar_bridge_evm if rewrite vald module
			//https://github.com/scalarorg/scalar-core/blob/main/vald/config/config.go#L24
			_, err = file.WriteString(fmt.Sprintf(`
[[scalar_bridge_evm]]
id = "%s"
rpc_addr = "%s"
with_bridge = true
finality_override = "confirmation"
# When using new chains (not Ethereum Mainnet), you may need to set the finality override to "confirmation" to avoid issues with the bridge
# With finality override, scalar will create evm client using ethereum.go, not ethereum_2.go
			`, evmConfig.ID, evmConfig.RPCAddr))
		}
		if err != nil {
			log.Error().Err(err).Msg("Failed to write evm bridge config")
		}

		btcConfigs, err := ParseJsonArrayConfig[config.BTCConfig](fmt.Sprintf("%s/chains/btc.json", configPath))
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse btc config")
		}

		// Add btc bridge config after implement btc module
		for _, btcConfig := range btcConfigs {
			_, err = file.WriteString(fmt.Sprintf(`
[[scalar_bridge_btc]]
id = "%s"
chain = "%s"
tag = "%s"
version = %d
with_bridge = true
rpc_host = "%s"
rpc_port = %d
rpc_user = "%s"
rpc_pass = "%s"
disable_tls = true
disable_connect_on_new = true
disable_auto_reconnect = false
http_post_mode = true
					`, btcConfig.ID, btcConfig.Chain, btcConfig.Tag, btcConfig.Version, btcConfig.RPCHost, btcConfig.RPCPort, btcConfig.RPCUser, btcConfig.RPCPass))
		}
		if err != nil {
			log.Error().Err(err).Msg("Failed to write btc bridge config")
		}
	}
	return nil
}

func appendAdditionalKeys(configPath string, validatorInfo ValidatorInfo) error {
	file, err := os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Error().Err(err).Str("configPath", configPath).Msg("Could not open config file")
		return err
	}

	defer file.Close()

	_, err = file.WriteString(`
#######################################################
###         Additional Keys            ###
#######################################################
	`)

	if err != nil {
		log.Error().Err(err).Str("configPath", configPath).Msg("Could not write text to config file")
		return err
	}

	_, err = file.WriteString(fmt.Sprintf(`
[additional_keys]
btc_priv_key = "%s"
`, validatorInfo.AdditionalKeys.BtcPrivKey))
	if err != nil {
		log.Error().Err(err).Str("configPath", configPath).Msg("Could not write text to config file")
		return err
	}
	return nil

}

func setCustomAppConfig(cmd *cobra.Command) error {
	// Todo: add custom app config if needed
	// customAppTemplate, customAppConfig := createCustomAppConfig(BaseDenom)
	customAppTemplate, customAppConfig := sdkconfig.DefaultConfigTemplate, sdkconfig.DefaultConfig()
	sdkconfig.SetConfigTemplate(customAppTemplate)

	return sdkserver.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig)
}
func initGenFiles(
	clientCtx client.Context,
	mbm module.BasicManager,
	nodeConfig *tmconfig.Config,
	coinDenom string,
	relayer ScalarRelayer,
	validatorInfos []ValidatorInfo,
	protocols []Protocol,
	args initArgs,
) error {
	appGenState, err := GenerateGenesis(clientCtx, mbm, coinDenom, relayer, validatorInfos, protocols, args)
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
		fmt.Printf("Validator: power: %d; pubkey: %v, address: %s\n", validators[i].Power, validators[i].PubKey, sdk.AccAddress(validatorInfo.ValPubKey.Address()).String())
	}
	genDoc := tmtypes.GenesisDoc{
		ChainID:    args.chainID,
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
		seedAddrs := []string{}
		for index, validatorInfo := range validatorInfos {
			if validatorInfo.NodeID != validatorInfos[i].NodeID {
				seed := fmt.Sprintf(`[[seed]]
name = "validator%d"
address = "%s"
`, index+1, validatorInfo.SeedAddress)
				seeds = append(seeds, seed)
				seedAddrs = append(seedAddrs, validatorInfo.SeedAddress)
			}
		}

		nodeConfig.Moniker = validatorInfos[i].Moniker
		nodeConfig.SetRoot(validatorInfos[i].NodeDir)
		nodeConfig.P2P.Seeds = strings.Join(seedAddrs, ",")
		if err := createConfigFiles(nodeConfig, validatorInfos[i].NodeDir, args, validatorInfos[i].Host, i); err != nil {
			return err
		}

		utils.WriteFile("seeds.toml", filepath.Join(validatorInfos[i].NodeDir, "config"), []byte(strings.Join(seeds, "\n")))
	}
	return nil
}

func collectGenFiles(
	clientCtx client.Context, nodeConfig *tmconfig.Config, chainID string,
	validatorInfos []ValidatorInfo,
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
		nodeConfig.RPC.ListenAddress = fmt.Sprintf("tcp://0.0.0.0:%d", validatorInfo.RPCPort)
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
		host = "127.0.0.1"
		return host, nil
	}
	return fmt.Sprintf("%s%d", nodeDomain, i+1), nil
}

// func getIP(i int, startingIPAddr string) (ip string, err error) {
// 	if len(startingIPAddr) == 0 {
// 		ip, err = sdkserver.ExternalIP()
// 		if err != nil {
// 			return "", err
// 		}
// 		return ip, nil
// 	}
// 	return calculateIP(startingIPAddr, i)
// }

// func calculateIP(ip string, i int) (string, error) {
// 	ipv4 := net.ParseIP(ip).To4()
// 	if ipv4 == nil {
// 		return "", fmt.Errorf("%v: non ipv4 address", ip)
// 	}

// 	for j := 0; j < i; j++ {
// 		ipv4[3]++
// 	}

// 	return ipv4.String(), nil
// }

func ReadFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Error().Err(err).Msg("failed reading data from file")
	}
	fmt.Printf("\nLength: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
	fmt.Printf("\nError: %v", err)
}

// startTestnet starts an in-process testnet
func startTestnet(cmd *cobra.Command, args startArgs) error {
	networkConfig := DefaultConfig()
	networkConfig.NumValidators = args.numValidators
	// Default networkConfig.ChainID is random, and we should only override it if chainID provided
	// is non-empty
	// if args.chainID != "" {
	// 	networkConfig.ChainID = args.chainID
	// }
	// networkConfig.SigningAlgo = args.algo
	// networkConfig.MinGasPrices = args.minGasPrices
	// networkConfig.EnableTMLogging = args.enableLogging
	// networkConfig.RPCAddress = args.rpcAddress
	// networkConfig.APIAddress = args.apiAddress
	// networkConfig.GRPCAddress = args.grpcAddress
	// networkConfig.JSONRPCAddress = args.jsonrpcAddress
	// networkConfig.PrintMnemonic = args.printMnemonic
	networkLogger := NewCLILogger(cmd)

	baseDir := args.baseDir
	testnet, err := New(networkLogger, cmd.InOrStdin(), baseDir, networkConfig)
	if err != nil {
		return err
	}

	_, err = testnet.WaitForHeightWithTimeout(int64(args.blockHeight), time.Duration(args.timeout)*time.Second)
	if err != nil {
		return err
	}

	cmd.Println("press the Enter Key to terminate")
	_, err = fmt.Scanln() // wait for Enter Key
	if err != nil {
		return err
	}
	// testnet.Cleanup()

	return nil
}

func loadEnvFile(envFile string) error {
	if envFile == "" {
		return nil // No env file specified, skip loading
	}

	data, err := os.ReadFile(envFile)
	if err != nil {
		return fmt.Errorf("failed to read env file %s: %w", envFile, err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		// Remove quotes if present
		value = strings.Trim(value, `"'`)

		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("failed to set env variable %s: %w", key, err)
		}
	}

	return nil
}
