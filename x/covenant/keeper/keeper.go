package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/key"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/tendermint/tendermint/libs/log"
)

var (
	chainPrefix    = utils.KeyFromStr("chain")
	subspacePrefix = "subspace"
)

var _ types.CovenantKeeper = &Keeper{}

// BaseKeeper implements the base Keeper
type Keeper struct {
	initialized bool
	internalKeeper
}

type internalKeeper struct {
	cdc          codec.BinaryCodec
	storeKey     sdk.StoreKey
	paramsKeeper types.ParamsKeeper
}

// NewKeeper returns a new EVM base keeper
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, paramsKeeper types.ParamsKeeper) Keeper {
	return Keeper{
		internalKeeper: internalKeeper{
			cdc:          cdc,
			storeKey:     storeKey,
			paramsKeeper: paramsKeeper,
		},
	}
}

// InitChains initializes all existing EVM chains and their respective param subspaces
func (k *Keeper) InitChains(ctx sdk.Context) {
	if k.initialized {
		panic("chains are already initialized")
	}

	iter := k.getBaseStore(ctx).IteratorNew(key.FromStr(subspacePrefix))
	defer utils.CloseLogError(iter, k.Logger(ctx))

	// for ; iter.Valid(); iter.Next() {
	// 	_ = k.createSubspace(ctx, nexus.ChainName(iter.Value()))
	// }

	k.initialized = true
}
func (k Keeper) CreateCustodian(ctx sdk.Context, params types.Params) (err error) {
	defer func() {
		err = sdkerrors.Wrap(err, "cannot create new EVM chain")
	}()

	if !k.initialized {
		panic("InitChain must be called before chain keepers can be used")
	}

	if err := params.Validate(); err != nil {
		return err
	}
	// chainKey := key.FromStr(subspacePrefix).Append(key.FromStr(params.Chain.String()))
	// if k.getBaseStore(ctx).HasNew(chainKey) {
	// 	return fmt.Errorf("chain %s already exists", params.Chain)
	// }

	// k.getBaseStore(ctx).SetRawNew(chainKey, []byte(params.Chain))

	// k.createSubspace(ctx, params.Chain).SetParamSet(ctx, &params)
	return nil
}
func (k Keeper) GetCustodians(ctx sdk.Context) (custodians []*types.Custodian, ok bool) {
	return nil, false
}
func (k Keeper) CreateCustodianGroup(ctx sdk.Context, params types.Params) (err error) {
	defer func() {
		err = sdkerrors.Wrap(err, "cannot create new EVM chain")
	}()

	return nil
}
func (k Keeper) GetCustodianGroup(ctx sdk.Context) (custodianGroup *types.CustodianGroup, ok bool) {
	return &types.CustodianGroup{}, false
}
func (k Keeper) createSubspace(ctx sdk.Context, chain nexus.ChainName) params.Subspace {
	chainKey := key.FromStr(types.ModuleName).Append(key.From(chain))
	k.Logger(ctx).Debug(fmt.Sprintf("initialized evm subspace %s", chain))
	return k.paramsKeeper.Subspace(chainKey.String()).WithKeyTable(types.KeyTable())
}

// Logger returns a module-specific logger.
func (k internalKeeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k internalKeeper) getBaseStore(ctx sdk.Context) utils.KVStore {
	return utils.NewNormalizedStore(ctx.KVStore(k.storeKey), k.cdc)
}
