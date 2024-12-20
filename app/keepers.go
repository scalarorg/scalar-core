package app

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibctransferkeeper "github.com/cosmos/ibc-go/v4/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v4/modules/core/02-client"
	ibcclienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	ibchost "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"
	"github.com/scalarorg/scalar-core/x/scalarnet"
	"golang.org/x/mod/semver"

	scalarParams "github.com/scalarorg/scalar-core/app/params"
	"github.com/scalarorg/scalar-core/utils/maps"
	bankKeeper "github.com/scalarorg/scalar-core/x/bank/keeper"
	evmKeeper "github.com/scalarorg/scalar-core/x/evm/keeper"
	evmTypes "github.com/scalarorg/scalar-core/x/evm/types"
	multisigKeeper "github.com/scalarorg/scalar-core/x/multisig/keeper"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	nexusKeeper "github.com/scalarorg/scalar-core/x/nexus/keeper"
	nexusTypes "github.com/scalarorg/scalar-core/x/nexus/types"
	permissionKeeper "github.com/scalarorg/scalar-core/x/permission/keeper"
	permissionTypes "github.com/scalarorg/scalar-core/x/permission/types"
	rewardKeeper "github.com/scalarorg/scalar-core/x/reward/keeper"
	rewardTypes "github.com/scalarorg/scalar-core/x/reward/types"
	scalarnetKeeper "github.com/scalarorg/scalar-core/x/scalarnet/keeper"
	scalarnetTypes "github.com/scalarorg/scalar-core/x/scalarnet/types"
	snapKeeper "github.com/scalarorg/scalar-core/x/snapshot/keeper"
	snapTypes "github.com/scalarorg/scalar-core/x/snapshot/types"
	tssKeeper "github.com/scalarorg/scalar-core/x/tss/keeper"
	tssTypes "github.com/scalarorg/scalar-core/x/tss/types"
	voteKeeper "github.com/scalarorg/scalar-core/x/vote/keeper"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"

	chainsKeeper "github.com/scalarorg/scalar-core/x/chains/keeper"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

type KeeperCache struct {
	repository map[string]any
}

func NewKeeperCache() *KeeperCache {
	return &KeeperCache{
		repository: make(map[string]any),
	}
}

func (k *KeeperCache) getSubspace(moduleName string) paramstypes.Subspace {
	paramsK := GetKeeper[paramskeeper.Keeper](k)
	subspace, ok := paramsK.GetSubspace(moduleName)
	if !ok {
		panic(fmt.Sprintf("subspace %s not found", moduleName))
	}
	return subspace
}

func GetKeeper[T any](k *KeeperCache) *T {
	if reflect.TypeOf(*new(T)).Kind() == reflect.Ptr {
		panic(fmt.Sprintf("the generic parameter for %s cannot be a reference type", fullTypeName[T]()))
	}
	key := fullTypeName[T]()
	keeper, ok := k.repository[key].(*T)
	if !ok {
		panic(fmt.Sprintf("keeper %s not found", key))
	}
	return keeper
}

func SetKeeper[T any](k *KeeperCache, keeper T) {
	if reflect.TypeOf(keeper).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("keeper %s must be a reference type", fullTypeName[T]()))
	}

	k.repository[fullTypeName[T]()] = keeper
}

func fullTypeName[T any]() string {
	keeperType := reflect.TypeOf(*new(T))

	if keeperType.Kind() == reflect.Ptr {
		keeperType = keeperType.Elem()
	}

	return keeperType.PkgPath() + "." + keeperType.Name()
}

func initParamsKeeper(encodingConfig scalarParams.EncodingConfig, key, tkey sdk.StoreKey) *paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(encodingConfig.Codec, encodingConfig.Amino, key, tkey)

	paramsKeeper.Subspace(bam.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable())

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(wasm.ModuleName)
	paramsKeeper.Subspace(snapTypes.ModuleName)
	paramsKeeper.Subspace(multisigTypes.ModuleName)
	paramsKeeper.Subspace(tssTypes.ModuleName)
	paramsKeeper.Subspace(nexusTypes.ModuleName)
	paramsKeeper.Subspace(scalarnetTypes.ModuleName)
	paramsKeeper.Subspace(rewardTypes.ModuleName)
	paramsKeeper.Subspace(voteTypes.ModuleName)
	paramsKeeper.Subspace(permissionTypes.ModuleName)

	return &paramsKeeper
}

func initStakingKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *stakingkeeper.Keeper {
	stakingK := stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		GetKeeper[authkeeper.AccountKeeper](keepers),
		GetKeeper[bankkeeper.BaseKeeper](keepers),
		keepers.getSubspace(stakingtypes.ModuleName),
	)
	return &stakingK
}

func initWasmKeeper(encodingConfig scalarParams.EncodingConfig, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache, bApp *bam.BaseApp, appOpts types.AppOptions, wasmOpts []wasm.Option, wasmDir string) *wasm.Keeper {
	wasmConfig := mustReadWasmConfig(appOpts)
	nexusK := GetKeeper[nexusKeeper.Keeper](keepers)

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	wasmOpts = append(
		wasmOpts,
		wasmkeeper.WithMessageHandlerDecorator(
			func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
				encoders := wasm.DefaultEncoders(encodingConfig.Codec, GetKeeper[ibctransferkeeper.Keeper](keepers))
				encoders.Custom = nexusKeeper.EncodeRoutingMessage

				return WithAnteHandlers(
					encoders,
					initMessageAnteDecorators(encodingConfig, keepers),
					// for security reasons we disallow some msg types that can be used for arbitrary calls
					wasmkeeper.NewMessageHandlerChain(NewMsgTypeBlacklistMessenger(), old, nexusKeeper.NewMessenger(nexusK)))
			}),
		wasmkeeper.WithWasmEngineDecorator(func(old wasmtypes.WasmerEngine) wasmtypes.WasmerEngine {
			return nexusKeeper.NewWasmerEngine(old, nexusK)
		}),
		wasmkeeper.WithQueryPlugins(NewQueryPlugins(nexusK)),
	)

	scopedWasmK := GetKeeper[capabilitykeeper.Keeper](keepers).ScopeToModule(wasm.ModuleName)
	ibcKeeper := GetKeeper[ibckeeper.Keeper](keepers)
	wasmK := wasm.NewKeeper(
		encodingConfig.Codec,
		keys[wasm.StoreKey],
		keepers.getSubspace(wasm.ModuleName),
		GetKeeper[authkeeper.AccountKeeper](keepers),
		GetKeeper[bankkeeper.BaseKeeper](keepers),
		GetKeeper[stakingkeeper.Keeper](keepers),
		GetKeeper[distrkeeper.Keeper](keepers),
		ibcKeeper.ChannelKeeper,
		ibcKeeper.ChannelKeeper,
		&ibcKeeper.PortKeeper,
		scopedWasmK,
		GetKeeper[ibctransferkeeper.Keeper](keepers),
		bApp.MsgServiceRouter(),
		bApp.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		WasmCapabilities,
		wasmOpts...,
	)

	return &wasmK
}

func initGovernanceKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *govkeeper.Keeper {
	// Add governance proposal hooks
	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(*GetKeeper[paramskeeper.Keeper](keepers))).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(*GetKeeper[distrkeeper.Keeper](keepers))).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(*GetKeeper[upgradekeeper.Keeper](keepers))).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(GetKeeper[ibckeeper.Keeper](keepers).ClientKeeper)).
		AddRoute(scalarnetTypes.RouterKey, scalarnet.NewProposalHandler(*GetKeeper[scalarnetKeeper.Keeper](keepers), GetKeeper[nexusKeeper.Keeper](keepers), GetKeeper[authkeeper.AccountKeeper](keepers)))

	if IsWasmEnabled() {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(GetKeeper[wasm.Keeper](keepers), wasm.EnableAllProposals))
	}

	govK := govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], keepers.getSubspace(govtypes.ModuleName), GetKeeper[authkeeper.AccountKeeper](keepers), GetKeeper[bankkeeper.BaseKeeper](keepers),
		GetKeeper[stakingkeeper.Keeper](keepers), govRouter,
	)

	scalarnetK := GetKeeper[scalarnetKeeper.Keeper](keepers)
	govK.SetHooks(govtypes.NewMultiGovHooks(scalarnetK.Hooks(GetKeeper[nexusKeeper.Keeper](keepers), govK)))
	return &govK
}

func initPermissionKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *permissionKeeper.Keeper {
	permissionK := permissionKeeper.NewKeeper(appCodec, keys[permissionTypes.StoreKey], keepers.getSubspace(permissionTypes.ModuleName))
	return &permissionK
}

func initVoteKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *voteKeeper.Keeper {
	voteRouter := voteTypes.NewRouter()
	// voteRouter.AddHandler(
	// 	evmTypes.ModuleName,
	// 	evmKeeper.NewVoteHandler(
	// 		appCodec,
	// 		GetKeeper[evmKeeper.BaseKeeper](keepers),
	// 		GetKeeper[nexusKeeper.Keeper](keepers),
	// 		GetKeeper[rewardKeeper.Keeper](keepers),
	// 	),
	// )

	voteRouter.AddHandler(
		chainsTypes.ModuleName,
		chainsKeeper.NewVoteHandler(
			appCodec,
			GetKeeper[chainsKeeper.BaseKeeper](keepers),
			GetKeeper[nexusKeeper.Keeper](keepers),
			GetKeeper[rewardKeeper.Keeper](keepers),
		),
	)

	voteK := voteKeeper.NewKeeper(
		appCodec,
		keys[voteTypes.StoreKey],
		keepers.getSubspace(voteTypes.ModuleName),
		GetKeeper[snapKeeper.Keeper](keepers),
		GetKeeper[stakingkeeper.Keeper](keepers),
		GetKeeper[rewardKeeper.Keeper](keepers),
	)
	voteK.SetVoteRouter(voteRouter)
	return &voteK
}

func initSnapshotKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *snapKeeper.Keeper {
	snapK := snapKeeper.NewKeeper(
		appCodec,
		keys[snapTypes.StoreKey],
		keepers.getSubspace(snapTypes.ModuleName),
		GetKeeper[stakingkeeper.Keeper](keepers),
		bankKeeper.NewBankKeeper(GetKeeper[bankkeeper.BaseKeeper](keepers)),
		GetKeeper[slashingkeeper.Keeper](keepers),
	)
	return &snapK
}

func initTssKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *tssKeeper.Keeper {
	tssK := tssKeeper.NewKeeper(appCodec, keys[tssTypes.StoreKey], keepers.getSubspace(tssTypes.ModuleName))
	return &tssK
}

func initMultisigKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *multisigKeeper.Keeper {
	multisigRouter := multisigTypes.NewSigRouter()
	// multisigRouter.AddHandler(evmTypes.ModuleName, evmKeeper.NewSigHandler(appCodec, GetKeeper[evmKeeper.BaseKeeper](keepers)))

	// TODO: Add chains handler
	// multisigRouter.AddHandler(chainsTypes.ModuleName, chainsKeeper.NewSigHandler(appCodec, GetKeeper[chainsKeeper.BaseKeeper](keepers)))

	multisigK := multisigKeeper.NewKeeper(appCodec, keys[multisigTypes.StoreKey], keepers.getSubspace(multisigTypes.ModuleName))
	multisigK.SetSigRouter(multisigRouter)
	return &multisigK
}

func initRewardKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *rewardKeeper.Keeper {
	rewardK := rewardKeeper.NewKeeper(
		appCodec,
		keys[rewardTypes.StoreKey],
		keepers.getSubspace(rewardTypes.ModuleName),
		bankKeeper.NewBankKeeper(GetKeeper[bankkeeper.BaseKeeper](keepers)),
		GetKeeper[distrkeeper.Keeper](keepers),
		GetKeeper[stakingkeeper.Keeper](keepers),
	)
	return &rewardK
}

func initIBCKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *ibckeeper.Keeper {
	scopedIBCK := GetKeeper[capabilitykeeper.Keeper](keepers).ScopeToModule(ibchost.ModuleName)
	return ibckeeper.NewKeeper(
		appCodec,
		keys[ibchost.StoreKey],
		keepers.getSubspace(ibchost.ModuleName),
		GetKeeper[stakingkeeper.Keeper](keepers),
		GetKeeper[upgradekeeper.Keeper](keepers),
		scopedIBCK,
	)
}

func initIBCTransferKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache, ics4Wrapper ibctransfertypes.ICS4Wrapper) *ibctransferkeeper.Keeper {
	scopedTransferK := GetKeeper[capabilitykeeper.Keeper](keepers).ScopeToModule(ibctransfertypes.ModuleName)
	transferK := ibctransferkeeper.NewKeeper(
		appCodec, keys[ibctransfertypes.StoreKey], keepers.getSubspace(ibctransfertypes.ModuleName),
		// Use the IBC middleware stack
		ics4Wrapper,
		GetKeeper[ibckeeper.Keeper](keepers).ChannelKeeper, &GetKeeper[ibckeeper.Keeper](keepers).PortKeeper,
		GetKeeper[authkeeper.AccountKeeper](keepers), GetKeeper[bankkeeper.BaseKeeper](keepers), scopedTransferK,
	)
	return &transferK
}

func initWasmContractKeeper(keepers *KeeperCache) *wasmkeeper.PermissionedKeeper {
	return wasmkeeper.NewDefaultPermissionKeeper(GetKeeper[wasm.Keeper](keepers))
}

func initScalarIBCKeeper(keepers *KeeperCache) *scalarnetKeeper.IBCKeeper {
	ibcK := scalarnetKeeper.NewIBCKeeper(*GetKeeper[scalarnetKeeper.Keeper](keepers), GetKeeper[ibctransferkeeper.Keeper](keepers))
	return &ibcK
}

func initscalarnetKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *scalarnetKeeper.Keeper {
	scalarnetK := scalarnetKeeper.NewKeeper(
		appCodec,
		keys[scalarnetTypes.StoreKey],
		keepers.getSubspace(scalarnetTypes.ModuleName),
		GetKeeper[ibckeeper.Keeper](keepers).ChannelKeeper,
		GetKeeper[feegrantkeeper.Keeper](keepers),
	)
	return &scalarnetK
}

func initEvmKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *evmKeeper.BaseKeeper {
	return evmKeeper.NewKeeper(appCodec, keys[evmTypes.StoreKey], GetKeeper[paramskeeper.Keeper](keepers))
}

func initChainsKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *chainsKeeper.BaseKeeper {
	return chainsKeeper.NewKeeper(appCodec, keys[chainsTypes.StoreKey], GetKeeper[paramskeeper.Keeper](keepers))
}

func initNexusKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *nexusKeeper.Keeper {
	// setting validator will finalize all by sealing it
	// no more validators can be added
	addressValidators := nexusTypes.NewAddressValidators().
		// AddAddressValidator(evmTypes.ModuleName, evmKeeper.NewAddressValidator()).
		AddAddressValidator(chainsTypes.ModuleName, chainsKeeper.NewAddressValidator()).
		AddAddressValidator(scalarnetTypes.ModuleName, scalarnetKeeper.NewAddressValidator(GetKeeper[scalarnetKeeper.Keeper](keepers)))

	addressValidators.Seal()

	nexusK := nexusKeeper.NewKeeper(appCodec, keys[nexusTypes.StoreKey], keepers.getSubspace(nexusTypes.ModuleName))
	nexusK.SetAddressValidators(addressValidators)

	return &nexusK
}

func initCapabilityKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, memKeys map[string]*sdk.MemoryStoreKey) *capabilitykeeper.Keeper {
	return capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
}

func initFeegrantKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *feegrantkeeper.Keeper {
	feegrantK := feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], GetKeeper[authkeeper.AccountKeeper](keepers))
	return &feegrantK
}

func initEvidenceKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *evidencekeeper.Keeper {
	return evidencekeeper.NewKeeper(appCodec, keys[evidencetypes.StoreKey], GetKeeper[stakingkeeper.Keeper](keepers), GetKeeper[slashingkeeper.Keeper](keepers))
}

func initUpgradeKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, skipUpgradeHeights map[int64]bool, homePath string, bApp *bam.BaseApp) *upgradekeeper.Keeper {
	upgradeK := upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, bApp)
	return &upgradeK
}

func upgradeName(version string) string {
	if !strings.HasPrefix(version, "v") {
		version = fmt.Sprintf("v%s", version)
	}

	fmt.Println("upgradeName", version)
	fmt.Println("semver.MajorMinor", semver.MajorMinor(version))

	name := semver.MajorMinor(version)
	if name == "" {
		panic(fmt.Errorf("invalid app version %s", version))
	}
	return name
}

func initCrisisKeeper(keepers *KeeperCache, invCheckPeriod uint) *crisiskeeper.Keeper {
	crisisK := crisiskeeper.NewKeeper(
		keepers.getSubspace(crisistypes.ModuleName),
		invCheckPeriod,
		GetKeeper[bankkeeper.BaseKeeper](keepers),
		authtypes.FeeCollectorName,
	)
	return &crisisK
}

func initSlashingKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *slashingkeeper.Keeper {
	slashK := slashingkeeper.NewKeeper(appCodec, keys[slashingtypes.StoreKey], GetKeeper[stakingkeeper.Keeper](keepers), keepers.getSubspace(slashingtypes.ModuleName))
	return &slashK
}

func initDistributionKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache, moduleAccPerms map[string][]string) *distrkeeper.Keeper {
	distrK := distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		keepers.getSubspace(distrtypes.ModuleName),
		GetKeeper[authkeeper.AccountKeeper](keepers),
		GetKeeper[bankkeeper.BaseKeeper](keepers),
		GetKeeper[stakingkeeper.Keeper](keepers),
		authtypes.FeeCollectorName,
		moduleAccountAddrs(moduleAccPerms),
	)
	return &distrK
}

func initMintKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache) *mintkeeper.Keeper {
	mintK := mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		keepers.getSubspace(minttypes.ModuleName),
		GetKeeper[stakingkeeper.Keeper](keepers),
		GetKeeper[authkeeper.AccountKeeper](keepers),
		GetKeeper[bankkeeper.BaseKeeper](keepers),
		authtypes.FeeCollectorName,
	)
	return &mintK
}

func initBankKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache, moduleAccPerms map[string][]string) *bankkeeper.BaseKeeper {
	bankK := bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		GetKeeper[authkeeper.AccountKeeper](keepers),
		keepers.getSubspace(banktypes.ModuleName),
		maps.Filter(moduleAccountAddrs(moduleAccPerms), func(addr string, _ bool) bool {
			// we do not rely on internal balance tracking for invariance checks in the nexus module
			// (https://github.com/cosmos/cosmos-sdk/issues/12825 for more details on the purpose of the blocked list),
			// but the nexus module account must be able to send or receive coins to mint/burn them for cross-chain transfers,
			// so we exclude this address from the blocked list
			return addr != authtypes.NewModuleAddress(nexusTypes.ModuleName).String()
		}),
	)
	return &bankK
}

func initAccountKeeper(appCodec codec.Codec, keys map[string]*sdk.KVStoreKey, keepers *KeeperCache, moduleAccPerms map[string][]string) *authkeeper.AccountKeeper {
	authK := authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		keepers.getSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		moduleAccPerms,
	)

	return &authK
}

// moduleAccountAddrs returns all the app's module account addresses.
func moduleAccountAddrs(moduleAccPerms map[string][]string) map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range moduleAccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}
