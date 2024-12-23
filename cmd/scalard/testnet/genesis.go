package testnet

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

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
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/utils"
	evmtypes "github.com/scalarorg/scalar-core/x/chains/evm/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	covenanttypes "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	nexustypes "github.com/scalarorg/scalar-core/x/nexus/types"
	permissionexported "github.com/scalarorg/scalar-core/x/permission/exported"
	permissiontypes "github.com/scalarorg/scalar-core/x/permission/types"
	protocoltypes "github.com/scalarorg/scalar-core/x/protocol/types"
	snapshottypes "github.com/scalarorg/scalar-core/x/snapshot/types"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	ValidatorKeyName   = "priv_validator"
	BroadcasterKeyName = "broadcaster"
	ScalarKeyName      = "scalar"
	GovKeyName         = "govenance"
	FaucetKeyName      = "faucet"
)

type GenesisState map[string]json.RawMessage

type ScalarProtocol struct {
	ScalarPubKey  cryptotypes.PubKey
	ScalarBalance banktypes.Balance
	EvmAddress    string
}

type ValidatorInfo struct {
	Host        string
	Moniker     string
	NodeID      string
	NodeDir     string
	SeedAddress string
	RPCPort     int

	ValPubKey  cryptotypes.PubKey
	ValBalance banktypes.Balance

	ValNodePubKey  cryptotypes.PubKey
	ValNodeBalance banktypes.Balance
	//Balance of broadcaster
	Broadcaster        cryptotypes.PubKey
	BroadcasterBalance banktypes.Balance

	GovPubKey  cryptotypes.PubKey
	GovBalance banktypes.Balance

	FaucetPubKey  cryptotypes.PubKey
	FaucetBalance banktypes.Balance

	GenesisValidator tmtypes.GenesisValidator
	BtcPubkey        string
	GenFile          string
}

// DefaultProtocol returns the default chains for a genesis state
func DefaultProtocol(scalarProtocol ScalarProtocol, custodianGroup covenanttypes.CustodianGroup) protocoltypes.Protocol {
	token := evmtypes.ERC20TokenMetadata{
		Asset:   "pBtc",
		ChainID: sdk.NewInt(1115511),
		//TokenAddress: evmtypes.Address(common.HexToAddress("0x25F23D37861210cdc3c694112cFa64bBca6D7143")),
		Status: evmtypes.Confirmed,
		Details: evmtypes.TokenDetails{
			TokenName: "pBtc",
			Symbol:    "pBtc",
			Decimals:  8,
			Capacity:  sdk.NewInt(100000000),
		},
	}

	params := chainsTypes.Params{
		Chain: "evm|11155111",
	}
	chain := protocoltypes.SupportedChain{
		Params: &params,
		Token: &protocoltypes.SupportedChain_Erc20{
			Erc20: &token,
		},
		Address: scalarProtocol.EvmAddress,
	}
	protocol := protocoltypes.Protocol{
		Pubkey:         scalarProtocol.ScalarPubKey.Bytes(),
		Address:        sdk.AccAddress(scalarProtocol.ScalarPubKey.Address()),
		Name:           protocoltypes.DefaultProtocolName,
		Tag:            "pools",
		Status:         protocoltypes.Activated,
		CustodianGroup: &custodianGroup,
		Chains:         []*protocoltypes.SupportedChain{&chain},
	}
	return protocol
}
func GenerateGenesis(clientCtx client.Context,
	mbm module.BasicManager,
	coinDenom string,
	validatorInfos []ValidatorInfo,
	scalarProtocol ScalarProtocol,
	supportedChainsPath string,
) (GenesisState, error) {
	appGenState := mbm.DefaultGenesis(clientCtx.Codec)
	genBalances := []banktypes.Balance{scalarProtocol.ScalarBalance}
	genAccounts := []authtypes.GenesisAccount{authtypes.NewBaseAccount(sdk.AccAddress(scalarProtocol.ScalarPubKey.Address()), scalarProtocol.ScalarPubKey, 0, 0)}
	//allValAmount := sdk.NewCoins()
	proxyValidators := []snapshottypes.ProxiedValidator{}

	for _, info := range validatorInfos {
		valBalances := sdk.NewCoins()
		//Validator balance must be set and greater than deligation amount
		genBalances = append(genBalances, info.ValNodeBalance)
		genAccounts = append(genAccounts, authtypes.NewBaseAccount(sdk.AccAddress(info.ValNodePubKey.Address()), info.ValNodePubKey, 0, 0))

		genBalances = append(genBalances, info.ValBalance)
		genAccounts = append(genAccounts, authtypes.NewBaseAccount(sdk.AccAddress(info.ValPubKey.Address()), info.ValPubKey, 0, 0))

		genBalances = append(genBalances, info.BroadcasterBalance)
		genAccounts = append(genAccounts, authtypes.NewBaseAccount(sdk.AccAddress(info.Broadcaster.Address()), info.Broadcaster, 0, 0))

		genBalances = append(genBalances, info.GovBalance)
		genAccounts = append(genAccounts, authtypes.NewBaseAccount(sdk.AccAddress(info.GovPubKey.Address()), info.GovPubKey, 0, 0))

		genBalances = append(genBalances, info.FaucetBalance)
		genAccounts = append(genAccounts, authtypes.NewBaseAccount(sdk.AccAddress(info.FaucetPubKey.Address()), info.FaucetPubKey, 0, 0))
		valAddr := sdk.ValAddress(info.ValPubKey.Address())
		proxyValidators = append(proxyValidators, snapshottypes.ProxiedValidator{
			Validator: valAddr,
			Proxy:     sdk.AccAddress(info.Broadcaster.Address()),
			Active:    true,
		})
		valBalances = valBalances.Add(info.ValNodeBalance.Coins...).
			Add(info.BroadcasterBalance.Coins...).
			Add(info.GovBalance.Coins...).
			Add(info.FaucetBalance.Coins...)
		fmt.Printf("valBalances %v\n", valBalances)
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
	totalSupply := sdk.NewCoins()
	for _, balance := range genBalances {
		fmt.Printf("address %s, balance %v\n", balance.Address, balance.Coins)
		totalSupply = totalSupply.Add(balance.Coins...)
	}
	// set the balances in the genesis state
	bankGenState := banktypes.DefaultGenesisState()
	bankGenState.Balances = genBalances
	bankGenState.Supply = totalSupply
	appGenState[banktypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(bankGenState)
	//Snapshoter
	snapshotGenState := snapshottypes.DefaultGenesisState()
	snapshotGenState.ProxiedValidators = proxyValidators
	appGenState[snapshottypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(snapshotGenState)
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
	// nexus
	validatorAddrs := []sdk.ValAddress{}
	for _, proxyValidator := range proxyValidators {
		validatorAddrs = append(validatorAddrs, proxyValidator.Validator)
	}
	nexusGenState := generateNexusGenesis(supportedChainsPath, validatorAddrs, coinDenom)
	appGenState[nexustypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(nexusGenState)
	//pemission module
	permissionGenState := permissiontypes.DefaultGenesisState()
	govAccounts := []permissiontypes.GovAccount{}
	govPubKeys := []cryptotypes.PubKey{}
	for _, info := range validatorInfos {
		govPubKeys = append(govPubKeys, info.GovPubKey, info.Broadcaster)
		if info.GovPubKey == nil {
			return appGenState, fmt.Errorf("gov pubkey is nil")
		}
		govAccounts = append(govAccounts,
			permissiontypes.GovAccount{
				Address: sdk.AccAddress(info.GovPubKey.Address()),
				Role:    permissionexported.ROLE_CHAIN_MANAGEMENT,
			},
			permissiontypes.GovAccount{
				Address: sdk.AccAddress(info.Broadcaster.Address()),
				Role:    permissionexported.ROLE_CHAIN_MANAGEMENT,
			},
		)

	}
	//Use first governance account as role access control.
	//Scalar allow an unique account with role access control
	govAccounts = append(govAccounts, permissiontypes.GovAccount{
		Address: sdk.AccAddress(govPubKeys[0].Address()),
		Role:    permissionexported.ROLE_ACCESS_CONTROL,
	})

	permissionGenState.GovAccounts = append(permissionGenState.GovAccounts, govAccounts...)
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
	custodians := make([]*covenanttypes.Custodian, len(validatorInfos))
	custodianGroup := covenanttypes.CustodianGroup{
		Uid:         "1c92b906-d5f8-477d-98c7-0d70d94ebb36",
		Name:        "scalar",
		Custodians:  custodians,
		Quorum:      3,
		Status:      covenanttypes.Activated,
		Description: "Default custodial group, which contains all custodians",
	}
	for i, validator := range validatorInfos {
		btcPubkey, err := hex.DecodeString(validator.BtcPubkey)
		if err != nil {
			return appGenState, err
		}
		fmt.Printf("% x", btcPubkey)
		custodians[i] = &covenanttypes.Custodian{
			Name:      validator.Host,
			Status:    covenanttypes.Activated,
			BtcPubkey: btcPubkey,
		}
	}
	covnantGenState := covenanttypes.NewGenesisState(custodians, &custodianGroup)
	appGenState[covenanttypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&covnantGenState)
	//Protocol
	protocol := DefaultProtocol(scalarProtocol, custodianGroup)
	protocolGenState := protocoltypes.NewGenesisState([]*protocoltypes.Protocol{&protocol})
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
	// 	validator.Status = stakingtypes.Bonded
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
	// 	//fmt.Printf("delegation %v, validator %v", delegation, validator)
	// 	stakingGenState.Delegations = append(stakingGenState.Delegations, delegation)
	// }
	return stakingGenState
}
func generateNexusGenesis(supportedChainsPath string, validatorAddrs []sdk.ValAddress, coinDenom string) *nexustypes.GenesisState {
	nexusGenState := nexustypes.DefaultGenesisState()
	if supportedChainsPath != "" {
		chainConfigs, err := ParseJsonArrayConfig[chainsTypes.ChainConfig](fmt.Sprintf("%s/chains.json", supportedChainsPath))
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse chains config")
		}
		for _, chainConfig := range chainConfigs {
			err := chainConfig.ValidateBasic()
			if err != nil {
				log.Error().Err(err).Msg("Failed to validate chains config")
			}
			nexusGenState.Chains = append(nexusGenState.Chains, nexus.Chain{
				Name:                  nexus.ChainName(chainConfig.ID),
				SupportsForeignAssets: true,
				KeyType:               tss.Multisig,
				Module:                chainsTypes.ModuleName,
			})
			chainState := nexustypes.ChainState{
				Chain: nexus.Chain{
					Name:                  nexus.ChainName(chainConfig.ID),
					SupportsForeignAssets: true,
					KeyType:               tss.Multisig,
					Module:                chainsTypes.ModuleName,
				},
				Activated:        true,
				Assets:           []nexus.Asset{nexus.NewAsset(coinDenom, true)},
				MaintainerStates: make([]nexustypes.MaintainerState, len(validatorAddrs)),
			}
			for i, valAddr := range validatorAddrs {
				chainState.MaintainerStates[i] = *nexustypes.NewMaintainerState(nexus.ChainName(chainConfig.ID), valAddr)
			}
			nexusGenState.ChainStates = append(nexusGenState.ChainStates, chainState)
		}
	}
	return nexusGenState
}
func GenerateSupportedChains(clientCtx client.Context, supportedChainsPath string, genesisState map[string]json.RawMessage) error {
	if supportedChainsPath != "" {
		chainConfigs, err := ParseJsonArrayConfig[chainsTypes.ChainConfig](fmt.Sprintf("%s/chains.json", supportedChainsPath))
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse chains config")
		}
		state := chainsTypes.DefaultGenesisState()
		for _, chainConfig := range chainConfigs {
			params := chainsTypes.Params{
				ChainId:             sdk.NewInt(int64(chainConfig.ChainID)),
				Chain:               nexus.ChainName(chainConfig.ID),
				ConfirmationHeight:  2,
				NetworkKind:         chainConfig.NetworkKind,
				RevoteLockingPeriod: 50,
				VotingThreshold:     utils.Threshold{Numerator: 51, Denominator: 100},
				MinVoterCount:       1,
				VotingGracePeriod:   50,
				EndBlockerLimit:     50,
				TransferLimit:       1000,
			}
			//Check if chainName is already in the genesis state
			addChain := true
			for _, chain := range state.Chains {
				if chain.Params.Chain == nexus.ChainName(chainConfig.ID) {
					addChain = false
					log.Debug().Msgf("chain name %s already exists", chainConfig.ID)
				}
			}
			if addChain {
				state.Chains = append(state.Chains, chainsTypes.GenesisState_Chain{
					Params:              params,
					CommandQueue:        utils.QueueState{},
					ConfirmedEventQueue: utils.QueueState{},
					ConfirmedStakingTxs: make([]chainsTypes.StakingTx, 0),
					CommandBatches:      make([]chainsTypes.CommandBatchMetadata, 0),
					Events:              make([]chainsTypes.Event, 0),
				})
			}
		}
		genesisState[chainsTypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&state)
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
