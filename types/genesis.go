package types

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	//btctypes "github.com/scalarorg/scalar-core/x/btc/types"
	"github.com/axelarnetwork/axelar-core/utils"
	evmtypes "github.com/axelarnetwork/axelar-core/x/evm/types"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	permissionexported "github.com/axelarnetwork/axelar-core/x/permission/exported"
	permissiontypes "github.com/axelarnetwork/axelar-core/x/permission/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	tmtypes "github.com/tendermint/tendermint/types"
)

type GenesisState map[string]json.RawMessage

type EvmNetworkConfig struct {
	ChainID    uint64        `json:"chainId"`
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Gateway    string        `json:"gateway"`
	Finality   int           `json:"finality"`
	LastBlock  uint64        `json:"lastBlock"`
	GasLimit   uint64        `json:"gasLimit"`
	BlockTime  time.Duration `json:"blockTime"` //Timeout im ms for pending txs
	MaxRetry   int           `json:"maxRetry"`
	RetryDelay time.Duration `json:"retryDelay"`
	TxTimeout  time.Duration `json:"txTimeout"` //Timeout for send txs (~3s)
	RpcUrl     string        `json:"rpcUrl"`
}

type BtcNetworkConfig struct {
	ChainID    uint64        `json:"chainId"`
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Gateway    string        `json:"gateway"` //Taproot address
	Finality   int           `json:"finality"`
	LastBlock  uint64        `json:"lastBlock"`
	GasLimit   uint64        `json:"gasLimit"`
	BlockTime  time.Duration `json:"blockTime"` //Timeout im ms for pending txs
	MaxRetry   int           `json:"maxRetry"`
	RetryDelay time.Duration `json:"retryDelay"`
	TxTimeout  time.Duration `json:"txTimeout"` //Timeout for send txs (~3s)
	RpcHost    string        `json:"rpcHost"`
	RpcPort    int           `json:"rpcPort"`
	RpcUser    string        `json:"rpcUser"`
	RpcPass    string        `json:"rpcPass"`
}

type ValidatorInfo struct {
	NodeID           string
	NodeDir          string
	SeedAddress      string
	ValPubKey        cryptotypes.PubKey
	Balances         []banktypes.Balance
	Accounts         []*authtypes.BaseAccount
	MngAccount       permissiontypes.GovAccount
	GovPubKey        cryptotypes.PubKey
	Broadcaster      cryptotypes.PubKey
	GenesisValidator tmtypes.GenesisValidator
	GenFile          string
}

func GenerateGenesis(clientCtx client.Context,
	mbm module.BasicManager,
	coinDenom string,
	validatorInfo []ValidatorInfo,
	supportedChainsPath string,
) (GenesisState, error) {
	appGenState := mbm.DefaultGenesis(clientCtx.Codec)
	genAccounts := []authtypes.GenesisAccount{}
	genBalances := []banktypes.Balance{}
	for _, info := range validatorInfo {
		for _, account := range info.Accounts {
			genAccounts = append(genAccounts, account)
		}
		genBalances = append(genBalances, info.Balances...)
	}
	// set the accounts in the genesis state
	var authGenState authtypes.GenesisState
	clientCtx.Codec.MustUnmarshalJSON(appGenState[authtypes.ModuleName], &authGenState)
	accounts, err := authtypes.PackAccounts(genAccounts)
	if err != nil {
		return appGenState, err
	}

	authGenState.Accounts = accounts
	appGenState[authtypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&authGenState)

	// set the balances in the genesis state
	bankGenState := banktypes.DefaultGenesisState()
	bankGenState.Balances = genBalances
	appGenState[banktypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(bankGenState)

	//set crisis constant fee coin denom
	crisisGenState := crisistypes.DefaultGenesisState()
	crisisGenState.ConstantFee.Denom = coinDenom
	appGenState[crisistypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(crisisGenState)

	// var feemarketGenState feemarkettypes.GenesisState
	// clientCtx.Codec.MustUnmarshalJSON(appGenState[feemarkettypes.ModuleName], &feemarketGenState)

	// feemarketGenState.Params.BaseFee = baseFee
	// feemarketGenState.Params.MinGasPrice = minGasPrice
	// appGenState[feemarkettypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&feemarketGenState)

	//Set gov params
	govGenState := govtypes.DefaultGenesisState()
	govGenState.DepositParams.MinDeposit[0].Denom = coinDenom
	govGenState.DepositParams.MaxDepositPeriod = 30 * time.Second
	govGenState.VotingParams.VotingPeriod = 30 * time.Second
	govGenState.TallyParams.Quorum = sdk.NewDec(50)
	//govGenState.TallyParams.Threshold = sdk.NewDec(50)
	appGenState[govtypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(govGenState)

	//mint genesis coins denom
	mintGenState := minttypes.DefaultGenesisState()
	mintGenState.Params.MintDenom = coinDenom
	appGenState[minttypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(mintGenState)
	//pemission module
	permissionGenState := permissiontypes.DefaultGenesisState()
	govMngAccounts := make([]permissiontypes.GovAccount, len(validatorInfo))
	govControlAccounts := make([]permissiontypes.GovAccount, len(validatorInfo))
	govPubKeys := make([]cryptotypes.PubKey, len(validatorInfo))
	for i, info := range validatorInfo {
		govPubKeys[i] = info.GovPubKey
		if info.GovPubKey == nil {
			return appGenState, fmt.Errorf("gov pubkey is nil")
		}
		govMngAccounts[i] = info.MngAccount
		govControlAccounts[i] = permissiontypes.GovAccount{
			Address: sdk.AccAddress(info.GovPubKey.Address()),
			Role:    permissionexported.ROLE_ACCESS_CONTROL,
		}
	}

	permissionGenState.GovAccounts = append(permissionGenState.GovAccounts, govControlAccounts...)
	permissionGenState.GovAccounts = append(permissionGenState.GovAccounts, govMngAccounts...)
	permissionGenState.GovernanceKey = multisig.NewLegacyAminoPubKey(1, govPubKeys)
	appGenState[permissiontypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(permissionGenState)
	//set staking params
	stakingGenState := stakingtypes.DefaultGenesisState()
	stakingGenState.Params.BondDenom = coinDenom
	appGenState[stakingtypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(stakingGenState)

	if err := GenerateSupportedChains(clientCtx, supportedChainsPath, appGenState); err != nil {
		log.Error().Err(err).Msg("Failed to generate supported chains")
	}
	// //evm chains
	// var evmGenState evmtypes.GenesisState
	// clientCtx.Codec.MustUnmarshalJSON(appGenState[evmtypes.ModuleName], &evmGenState)
	// //btc chains
	// var btcGenState btctypes.GenesisState
	// clientCtx.Codec.MustUnmarshalJSON(appGenState[btctypes.ModuleName], &btcGenState)

	// var feemarketGenState feemarkettypes.GenesisState
	// clientCtx.Codec.MustUnmarshalJSON(appGenState[feemarkettypes.ModuleName], &feemarketGenState)

	// feemarketGenState.Params.BaseFee = baseFee
	// feemarketGenState.Params.MinGasPrice = minGasPrice
	// appGenState[feemarkettypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&feemarketGenState)

	return appGenState, nil
}

func GenerateSupportedChains(clientCtx client.Context, supportedChainsPath string, genesisState map[string]json.RawMessage) error {
	if supportedChainsPath != "" {
		evmConfigs, err := ParseJsonArrayConfig[EvmNetworkConfig](fmt.Sprintf("%s/evm.json", supportedChainsPath))
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse evm config")
		}
		evmGenState := evmtypes.DefaultGenesisState()
		for _, evmConfig := range evmConfigs {
			//Todo: get token code and burnable
			params := evmtypes.Params{
				Chain:               nexus.ChainName(evmConfig.ID),
				ConfirmationHeight:  1,
				Network:             evmConfig.ID,
				TokenCode:           []byte{},
				Burnable:            []byte{},
				RevoteLockingPeriod: 50,
				Networks: []evmtypes.NetworkInfo{
					{
						Name: evmConfig.ID,
						Id:   sdk.NewIntFromUint64(evmConfig.ChainID),
					},
				},
				VotingThreshold:   utils.Threshold{Numerator: 51, Denominator: 100},
				VotingGracePeriod: 3,
				MinVoterCount:     1,
				CommandsGasLimit:  5000000,
				EndBlockerLimit:   50,
				TransferLimit:     50,
			}
			evmGenState.Chains = append(evmGenState.Chains, evmtypes.GenesisState_Chain{
				Params:              params,
				CommandQueue:        utils.QueueState{},
				ConfirmedEventQueue: utils.QueueState{},
				Gateway: evmtypes.Gateway{
					Address: evmtypes.Address(common.HexToAddress(evmConfig.Gateway)),
				},
			})
		}
		genesisState[evmtypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&evmGenState)
		// btcConfigs, err := ParseJsonArrayConfig[evmtypes.EVMConfig](fmt.Sprintf("%s/btc.json", supportedChainsPath))
		// if err != nil {
		// 	log.Error().Err(err).Msg("Failed to parse btc config")
		// }
	}
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
