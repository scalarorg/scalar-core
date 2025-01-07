package testnet

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"path"
	"strings"
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
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/utils"
	covenanttypes "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	nexustypes "github.com/scalarorg/scalar-core/x/nexus/types"
	permissionexported "github.com/scalarorg/scalar-core/x/permission/exported"
	permissiontypes "github.com/scalarorg/scalar-core/x/permission/types"
	protocoltypes "github.com/scalarorg/scalar-core/x/protocol/types"
	snapshottypes "github.com/scalarorg/scalar-core/x/snapshot/types"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"

	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

// DefaultProtocol returns the default chains for a genesis state
func DefaultProtocol(scalarProtocol ScalarProtocol, tokenInfo Token, custodianGroup covenanttypes.CustodianGroup) protocoltypes.Protocol {
	log.Debug().Any("TokenInfo", tokenInfo).Msg("Create defaultProtocol")
	// chains := make([]*protocoltypes.SupportedChain, len(tokenInfos))
	// for i, tokenInfo := range tokenInfos {
	// 	tokenAddress := chainsTypes.Address(common.HexToAddress(tokenInfo.TokenAddress))
	// 	log.Debug().Any("TokenAddress", tokenAddress).Msg("Parse Tokenaddress")
	// 	params := chainsTypes.Params{
	// 		Chain:       nexus.ChainName(tokenInfo.ID),
	// 		ChainID:     sdk.NewInt(tokenInfo.ChainID),
	// 		NetworkKind: chainsTypes.Testnet,
	// 	}
	// 	supportedChain := protocoltypes.SupportedChain{
	// 		Chain:       nexus.ChainName(tokenInfo.ID):      &params,
	// 		Address: tokenInfo.ProtocolAddress,
	// 	}
	// 	if types.IsEvmChain(nexus.ChainName(tokenInfo.ID)) {
	// 		token := chainsTypes.ERC20TokenMetadata{
	// 			Asset:   tokenInfo.Asset,
	// 			ChainID: sdk.NewInt(tokenInfo.ChainID),
	// 			// TxHash:  chainsTypes.Hash(chainsTypes.ZeroHash),
	// 			// TokenAddress: tokenAddress,
	// 			Status: chainsTypes.Confirmed,
	// 			Details: chainsTypes.TokenDetails{
	// 				TokenName: tokenInfo.Name,
	// 				Symbol:    tokenInfo.Symbol,
	// 				Decimals:  tokenInfo.Decimals,
	// 				Capacity:  sdk.NewInt(tokenInfo.Capacity),
	// 			},
	// 		}
	// 		supportedChain.Token = &protocoltypes.SupportedChain_Erc20{
	// 			Erc20: &token,
	// 		}
	// 		chains[i] = &supportedChain
	// 	} else if types.IsEvmChain(nexus.ChainName(tokenInfo.ID)) {
	// 		supportedChain.Token = &protocoltypes.SupportedChain_Btc{
	// 			Btc: &btctypes.BtcToken{},
	// 		}
	// 		chains[i] = &supportedChain
	// 	}

	// }
	attributes := protocoltypes.ProtocolAttribute{
		Model: protocoltypes.Pooling,
	}
	supportedChains := []*protocoltypes.SupportedChain{}
	for _, chain := range tokenInfo.Deployments {
		supportedChains = append(supportedChains, &protocoltypes.SupportedChain{
			Chain:   nexus.ChainName(chain.ID),
			Name:    chain.Name,
			Address: chain.TokenAddress,
		})
	}
	protocol := protocoltypes.Protocol{
		Pubkey:         scalarProtocol.ScalarPubKey.Bytes(),
		Address:        sdk.AccAddress(scalarProtocol.ScalarPubKey.Address()),
		Attribute:      &attributes,
		Name:           protocoltypes.DefaultProtocolName,
		Tag:            "pools",
		Status:         protocoltypes.Activated,
		CustodianGroup: &custodianGroup,
		Asset:          &chainsTypes.Asset{Chain: nexus.ChainName(tokenInfo.ID), Name: tokenInfo.Asset},
		Chains:         supportedChains,
	}
	return protocol
}
func GenerateGenesis(clientCtx client.Context,
	mbm module.BasicManager,
	coinDenom string,
	validatorInfos []ValidatorInfo,
	scalarProtocol ScalarProtocol,
	args initArgs,
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
	nexusGenState := generateNexusGenesis(args.chains, validatorAddrs, coinDenom)
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
	// scalar chains' module
	if err := GenerateSupportedChains(clientCtx, args.chains, appGenState); err != nil {
		log.Error().Err(err).Msg("Failed to generate supported chains")
	}
	//Covenant
	custodians := make([]*covenanttypes.Custodian, len(validatorInfos))
	custodianGroup := covenanttypes.CustodianGroup{
		Uid:         "mock|123456789",
		Name:        "scalar",
		Custodians:  custodians,
		Quorum:      3,
		Status:      covenanttypes.Activated,
		Description: "Default custodial group, which contains all custodians",
	}
	for i, validator := range validatorInfos {
		btcPrivKey, err := hex.DecodeString(validator.AdditionalKeys.BtcPrivKey)
		if err != nil {
			return appGenState, err
		}
		if len(btcPrivKey) != 32 {
			return appGenState, fmt.Errorf("invalid private key length")
		}

		privKey := secp256k1.PrivKeyFromBytes(btcPrivKey)

		custodians[i] = &covenanttypes.Custodian{
			Name:      validator.Host,
			Status:    covenanttypes.Activated,
			BtcPubkey: privKey.PubKey().SerializeCompressed(),
		}
	}

	defaultCovenantState := covenanttypes.DefaultGenesisState()
	covnantGenState := covenanttypes.NewGenesisState(&defaultCovenantState.Params, defaultCovenantState.SigningSessions, custodians, []*covenanttypes.CustodianGroup{&custodianGroup})
	appGenState[covenanttypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&covnantGenState)
	//Protocol
	protocolGenState, err := generateProtocolGenesis(scalarProtocol, custodianGroup, args.tokens)
	if err == nil {
		appGenState[protocoltypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(protocolGenState)
	}

	// var feemarketGenState feemarkettypes.GenesisState
	// clientCtx.Codec.MustUnmarshalJSON(appGenState[feemarkettypes.ModuleName], &feemarketGenState)

	// feemarketGenState.Params.BaseFee = baseFee
	// feemarketGenState.Params.MinGasPrice = minGasPrice
	// appGenState[feemarkettypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&feemarketGenState)

	return appGenState, nil
}
func generateProtocolGenesis(scalarProtocol ScalarProtocol, custodianGroup covenanttypes.CustodianGroup, tokensPath string) (*protocoltypes.GenesisState, error) {
	// evmTokenPath := path.Join(tokensPath, "evm.json")
	// log.Debug().Msgf("Read token config in the path %s", evmTokenPath)
	// tokenInfos, err := ParseJsonArrayConfig[Token](evmTokenPath)
	// if err != nil {
	// 	return nil, err
	// }
	btcTokenPath := path.Join(tokensPath, "btc.json")
	log.Debug().Msgf("Read token config in the path %s", btcTokenPath)
	tokenInfos, err := ParseJsonArrayConfig[Token](btcTokenPath)
	if err != nil {
		return nil, err
	}
	if len(tokenInfos) == 0 {
		log.Error().Msgf("Missing token infos in path %s", btcTokenPath)
	}
	log.Debug().Any("TokenInfo", tokenInfos).Msgf("Successfull parsed token config")
	protocol := DefaultProtocol(scalarProtocol, tokenInfos[0], custodianGroup)
	protocolGenState := protocoltypes.NewGenesisState([]*protocoltypes.Protocol{&protocol})
	return protocolGenState, nil
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
	// 	commission := chainsTypesNewCommissionWithTime(rate, maxRate, maxChangeRate, time.Now())
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
			chainName := nexus.ChainName(chainConfig.ID)
			nexusGenState.Chains = append(nexusGenState.Chains, nexus.Chain{
				Name:                  chainName,
				SupportsForeignAssets: true,
				KeyType:               tss.Multisig,
				Module:                chainsTypes.ModuleName,
			})
			chainState := nexustypes.ChainState{
				Chain: nexus.Chain{
					Name:                  chainName,
					SupportsForeignAssets: true,
					KeyType:               tss.Multisig,
					Module:                chainsTypes.ModuleName,
				},
				Activated:        true,
				Assets:           []nexus.Asset{nexus.NewAsset(coinDenom, true)},
				MaintainerStates: make([]nexustypes.MaintainerState, len(validatorAddrs)),
			}
			if chainsTypes.IsBitcoinChain(chainName) {
				chainState.Assets = append(chainState.Assets, nexus.NewAsset("sBtc", false))
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

		chainsState := chainsTypes.DefaultGenesisState()
		for _, chainConfig := range chainConfigs {
			//Check if chainName is already in the genesis state
			addChain := true
			chainName := nexus.ChainName(chainConfig.ID)
			for _, chain := range chainsState.Chains {
				if chain.Params.Chain == chainName {
					addChain = false
					log.Debug().Msgf("chain name %s already exists", chainConfig.ID)
				}
			}
			if addChain {
				var gateway *chainsTypes.Gateway
				if chainConfig.Gateway != "" {
					gwHex, err := getByteAddress(chainConfig.Gateway)
					fmt.Printf("Gateway %s, decoded %v", chainConfig.Gateway, gwHex)
					if err == nil && len(gwHex) == 20 {
						gateway = &chainsTypes.Gateway{
							Address: chainsTypes.Address(gwHex),
						}
					}
				}
				chain := chainsTypes.GenesisState_Chain{
					Params:              chainsTypes.DefaultChainParams(sdk.NewInt(int64(chainConfig.ChainID)), nexus.ChainName(chainConfig.ID), chainConfig.NetworkKind, chainConfig.Metadata),
					CommandQueue:        utils.QueueState{},
					ConfirmedEventQueue: utils.QueueState{},
					ConfirmedSourceTxs:  make([]chainsTypes.SourceTx, 0),
					CommandBatches:      make([]chainsTypes.CommandBatchMetadata, 0),
					Gateway:             gateway,
					Events:              make([]chainsTypes.Event, 0),
					Tokens:              []chainsTypes.ERC20TokenMetadata{},
				}
				if chainsTypes.IsBitcoinChain(chainName) {
					// Add default sBtc
					sBtc := createDefaultSbtc()
					chain.Tokens = append(chain.Tokens, sBtc)
				}
				chainsState.Chains = append(chainsState.Chains, chain)
			}
		}
		genesisState[chainsTypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&chainsState)
	}
	return nil
}
func createDefaultSbtc() chainsTypes.ERC20TokenMetadata {
	return chainsTypes.ERC20TokenMetadata{
		Asset:   "sBtc",
		ChainID: sdk.NewInt(4),
		Details: chainsTypes.TokenDetails{
			TokenName: "BtcTestnet4",
			Symbol:    "sBtc",
			Decimals:  8,
			Capacity:  sdk.NewInt(0),
		},
		TokenAddress: chainsTypes.Address(common.HexToAddress("ffffffffffffffffffffffffffffffffffffffff")),
		TxHash:       chainsTypes.ZeroHash,
		Status:       chainsTypes.Confirmed,
		IsExternal:   true,
		BurnerCode:   nil, //burner code for external tokens must be nil
	}
}
func getByteAddress(hexString string) ([]byte, error) {
	if strings.HasPrefix(hexString, "0x") {
		return hex.DecodeString(hexString[2:])
	} else {
		return hex.DecodeString(hexString)
	}
}
