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
	nexustypes "github.com/axelarnetwork/axelar-core/x/nexus/types"
	permissionexported "github.com/axelarnetwork/axelar-core/x/permission/exported"
	permissiontypes "github.com/axelarnetwork/axelar-core/x/permission/types"
	tss "github.com/axelarnetwork/axelar-core/x/tss/exported"
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
	covenanttypes "github.com/scalarorg/scalar-core/x/covenant/types"
	protocoltypes "github.com/scalarorg/scalar-core/x/protocol/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	ValidatorKeyName   = "priv_validator"
	BroadcasterKeyName = "broadcaster"
	GovKeyName         = "govenance"
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
	Host        string
	Moniker     string
	NodeID      string
	NodeDir     string
	SeedAddress string
	ValPubKey   cryptotypes.PubKey
	//Balance of validator
	ValBalance banktypes.Balance
	//Balance of broadcaster
	BroadcasterBalance banktypes.Balance
	NodeBalance        banktypes.Balance
	//BroadcasterAccount *authtypes.BaseAccount
	MngAccount       permissiontypes.GovAccount
	GovPubKey        cryptotypes.PubKey
	Broadcaster      cryptotypes.PubKey
	GenesisValidator tmtypes.GenesisValidator
	BtcPubkey        string
	GenFile          string
}

// DefaultProtocol returns the default chains for a genesis state
func DefaultProtocol() protocoltypes.Protocol {
	token := evmtypes.ERC20TokenMetadata{
		Asset:        "pBtc",
		ChainID:      sdk.NewInt(1115511),
		TokenAddress: evmtypes.Address(common.HexToAddress("0x5f214989a5f49ab3c56fd5003c2858e24959c018")),
		Status:       evmtypes.Confirmed,
		Details: evmtypes.TokenDetails{
			TokenName: "pBtc",
			Symbol:    "pBtc",
			Decimals:  8,
			Capacity:  sdk.NewInt(100000000),
		},
	}
	protocol := protocoltypes.Protocol{
		Name:          protocoltypes.DefaultProtocolName,
		CovenantGroup: covenanttypes.DefaultCovenantGroupName,
		Tokens: []evmtypes.ERC20TokenMetadata{
			token,
		},
	}
	return protocol
}
func GenerateGenesis(clientCtx client.Context,
	mbm module.BasicManager,
	coinDenom string,
	validatorInfos []ValidatorInfo,
	supportedChainsPath string,
) (GenesisState, error) {
	appGenState := mbm.DefaultGenesis(clientCtx.Codec)
	genAccounts := []authtypes.GenesisAccount{}
	genBalances := []banktypes.Balance{}
	unbondedPoolAmount := sdk.NewCoins()
	for _, info := range validatorInfos {
		//Validator balance must be set and greater than deligation amount
		genBalances = append(genBalances, banktypes.Balance{
			Address: sdk.AccAddress(info.ValPubKey.Address()).String(),
			Coins:   info.ValBalance.Coins,
		})
		genAccounts = append(genAccounts, authtypes.NewBaseAccount(sdk.AccAddress(info.ValPubKey.Address()), info.ValPubKey, 0, 0))

		genAccounts = append(genAccounts, authtypes.NewBaseAccount(sdk.AccAddress(info.Broadcaster.Address()), info.Broadcaster, 0, 0))
		genBalances = append(genBalances, info.BroadcasterBalance)

		genBalances = append(genBalances, info.NodeBalance)
		genAccounts = append(genAccounts, authtypes.NewBaseAccount(info.NodeBalance.GetAddress(), nil, 0, 0))
		unbondedPoolAmount = unbondedPoolAmount.Add(info.ValBalance.Coins...)
	}
	//Not bonded module accounts
	// macc := authtypes.NewEmptyModuleAccount(stakingtypes.NotBondedPoolName)
	// unbondedPoolBalance := banktypes.Balance{
	// 	Address: macc.GetAddress().String(),
	// 	Coins:   unbondedPoolAmount,
	// }
	// genAccounts = append(genAccounts, macc)
	// // fmt.Printf("unbondedPoolBalance %v, totalSupply %v", unbondedPoolBalance, totalSupply)
	// genBalances = append(genBalances, unbondedPoolBalance)
	// set the accounts in the genesis state
	var authGenState authtypes.GenesisState
	clientCtx.Codec.MustUnmarshalJSON(appGenState[authtypes.ModuleName], &authGenState)
	accounts, err := authtypes.PackAccounts(genAccounts)
	if err != nil {
		return appGenState, err
	}

	authGenState.Accounts = accounts
	appGenState[authtypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&authGenState)
	totalSupply := sdk.NewCoins()
	for _, balance := range genBalances {
		totalSupply = totalSupply.Add(balance.Coins...)
	}
	// set the balances in the genesis state
	bankGenState := banktypes.DefaultGenesisState()
	bankGenState.Balances = genBalances
	bankGenState.Supply = totalSupply
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
	//axelar nexus
	nexusGenState := generateNexusGenesis(supportedChainsPath, validatorInfos, coinDenom)
	appGenState[nexustypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(nexusGenState)
	//pemission module
	permissionGenState := permissiontypes.DefaultGenesisState()
	govMngAccounts := make([]permissiontypes.GovAccount, len(validatorInfos))
	govControlAccounts := make([]permissiontypes.GovAccount, len(validatorInfos))
	govPubKeys := make([]cryptotypes.PubKey, len(validatorInfos))
	for i, info := range validatorInfos {
		govPubKeys[i] = info.GovPubKey
		if info.GovPubKey == nil {
			return appGenState, fmt.Errorf("gov pubkey is nil")
		}
		govMngAccounts[i] = info.MngAccount
		govControlAccounts[i] = permissiontypes.GovAccount{
			Address: sdk.AccAddress(info.GovPubKey.Address()),
			Role:    permissionexported.ROLE_CHAIN_MANAGEMENT,
		}
	}

	govControlAccounts = append(govControlAccounts, permissiontypes.GovAccount{
		Address: sdk.AccAddress(govPubKeys[0].Address()),
		Role:    permissionexported.ROLE_ACCESS_CONTROL,
	})

	permissionGenState.GovAccounts = append(permissionGenState.GovAccounts, govControlAccounts...)
	permissionGenState.GovAccounts = append(permissionGenState.GovAccounts, govMngAccounts...)
	permissionGenState.GovernanceKey = multisig.NewLegacyAminoPubKey(1, govPubKeys)
	appGenState[permissiontypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(permissionGenState)
	//set staking params
	stakingGenState := generateStakingGenesis(coinDenom, validatorInfos)
	appGenState[stakingtypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(stakingGenState)
	// supported chains
	if err := GenerateSupportedChains(clientCtx, supportedChainsPath, appGenState); err != nil {
		log.Error().Err(err).Msg("Failed to generate supported chains")
	}
	//Covenant
	covenants := make([]covenanttypes.Covenant, len(validatorInfos))
	covenantGroup := covenanttypes.CovenantGroup{
		Name:      "scalar",
		Covenants: covenants,
	}
	for i, validator := range validatorInfos {
		covenants[i] = covenanttypes.Covenant{
			Name:      validator.Host,
			Btcpubkey: validator.BtcPubkey,
		}
	}
	covnantGenState := covenanttypes.NewGenesisState(covenants, covenantGroup)
	appGenState[covenanttypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(covnantGenState)
	//Protocol
	protocolGenState := protocoltypes.NewGenesisState(DefaultProtocol())
	appGenState[protocoltypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(protocolGenState)
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
func generateStakingGenesis(coinDenom string, validatorInfos []ValidatorInfo) *stakingtypes.GenesisState {
	stakingGenState := stakingtypes.DefaultGenesisState()
	stakingGenState.Params.BondDenom = coinDenom
	// Create call create_validator from client
	// This is set by execute CreateValidator txs
	// for _, validatorInfo := range validatorInfos {
	// 	validator, err := stakingtypes.NewValidator(
	// 		sdk.ValAddress(validatorInfo.ValPubKey.Address()),
	// 		validatorInfo.ValPubKey,
	// 		stakingtypes.Description{
	// 			Identity: validatorInfo.NodeID,
	// 			Moniker:  validatorInfo.Host,
	// 		},
	// 	)
	// 	validator.Tokens = validatorInfo.ValBalance.Coins[0].Amount
	// 	validator.Status = stakingtypes.Unbonded
	// 	validator.MinSelfDelegation = sdk.OneInt()
	// 	validator.DelegatorShares = sdk.NewDecFromInt(DelegatorTokens)
	// 	if err != nil {
	// 		log.Error().Err(err).Msg("Failed to generate staking validator")
	// 	}
	// 	rate, _ := sdk.NewDecFromStr("0.1")
	// 	maxRate, _ := sdk.NewDecFromStr("0.2")
	// 	maxChangeRate, _ := sdk.NewDecFromStr("0.01")
	// 	commission := types.NewCommissionWithTime(rate, maxRate, maxChangeRate, time.Now())
	// 	validator, err = validator.SetInitialCommission(commission)

	// 	if err != nil {
	// 		continue
	// 	}
	// 	stakingGenState.Validators = append(stakingGenState.Validators, validator)
	// 	delegation := stakingtypes.Delegation{
	// 		DelegatorAddress: sdk.AccAddress(validatorInfo.ValPubKey.Address()).String(),
	// 		ValidatorAddress: sdk.ValAddress(validatorInfo.ValPubKey.Address()).String(),
	// 		Shares:           validator.DelegatorShares,
	// 	}
	// 	fmt.Printf("delegation %v, validator %v", delegation, validator)
	// 	stakingGenState.Delegations = append(stakingGenState.Delegations, delegation)
	// }
	return stakingGenState
}
func generateNexusGenesis(supportedChainsPath string, validatorInfos []ValidatorInfo, coinDenom string) *nexustypes.GenesisState {
	nexusGenState := nexustypes.DefaultGenesisState()
	if supportedChainsPath != "" {
		evmConfigs, err := ParseJsonArrayConfig[EvmNetworkConfig](fmt.Sprintf("%s/evm.json", supportedChainsPath))
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse evm config")
		}
		for _, evmConfig := range evmConfigs {
			nexusGenState.Chains = append(nexusGenState.Chains, nexus.Chain{
				Name:                  nexus.ChainName(evmConfig.Name),
				SupportsForeignAssets: true,
				KeyType:               tss.Multisig,
				Module:                evmtypes.ModuleName,
			})
			chainState := nexustypes.ChainState{
				Chain: nexus.Chain{
					Name:                  nexus.ChainName(evmConfig.Name),
					SupportsForeignAssets: true,
					KeyType:               tss.Multisig,
					Module:                evmtypes.ModuleName,
				},
				Activated:        true,
				Assets:           []nexus.Asset{nexus.NewAsset(coinDenom, true)},
				MaintainerStates: make([]nexustypes.MaintainerState, len(validatorInfos)),
			}
			for i, validator := range validatorInfos {
				chainState.MaintainerStates[i] = nexustypes.MaintainerState{
					Address: sdk.ValAddress(validator.ValPubKey.Bytes()),
					Chain:   nexus.ChainName(evmConfig.Name),
				}
			}
			nexusGenState.ChainStates = append(nexusGenState.ChainStates, chainState)
		}
		btcConfigs, err := ParseJsonArrayConfig[BtcNetworkConfig](fmt.Sprintf("%s/btc.json", supportedChainsPath))
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse btc config")
		}
		for _, btcConfig := range btcConfigs {
			fmt.Printf("btcConfig %v\n", btcConfig)
			// nexusGenState.Chains = append(nexusGenState.Chains, nexus.Chain{
			// 	Name:                  nexus.ChainName(btcConfig.Name),
			// 	SupportsForeignAssets: true,
			// 	KeyType:               tss.Multisig,
			// 	Module:                btctypes.ModuleName,
			// })
			// chainState := nexustypes.ChainState{
			// 	Chain: nexus.Chain{
			// 		Name:                  nexus.ChainName(btcConfig.Name),
			// 		SupportsForeignAssets: true,
			// 		KeyType:               tss.Multisig,
			// 		Module:                btctypes.ModuleName,
			// 	},
			// 	Activated:        true,
			// 	Assets:           []nexus.Asset{nexus.NewAsset(coinDenom, true)},
			// 	MaintainerStates: make([]nexustypes.MaintainerState, len(validatorInfos)),
			// }
			// for i, validator := range validatorInfos {
			// 	chainState.MaintainerStates[i] = nexustypes.MaintainerState{
			// 		Address: sdk.ValAddress(validator.ValPubKey.Bytes()),
			// 		Chain:   nexus.ChainName(btcConfig.Name),
			// 	}
			// }
			// nexusGenState.ChainStates = append(nexusGenState.ChainStates, chainState)
		}
	}
	return nexusGenState
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
			tokenCode, err := utils.HexDecode(evmtypes.Token)
			if err != nil {
				return err
			}
			bzToken, err := utils.HexDecode(evmtypes.Burnable)
			if err != nil {
				return err
			}
			params := evmtypes.Params{
				Chain:               nexus.ChainName(evmConfig.ID),
				ConfirmationHeight:  1,
				Network:             evmConfig.ID,
				TokenCode:           tokenCode,
				Burnable:            bzToken,
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
