package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/axelarnetwork/axelar-core/utils"
	"github.com/axelarnetwork/axelar-core/utils/key"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

var (
	subspacePrefix = "subspace"
	chainPrefix    = utils.KeyFromStr("chain")
)

// TODO: Implement this keeper

var _ types.BaseKeeper = &BaseKeeper{}

// BaseKeeper implements the base Keeper
type BaseKeeper struct {
	initialized bool
	internalKeeper
}

type internalKeeper struct {
	cdc          codec.BinaryCodec
	storeKey     sdk.StoreKey
	paramsKeeper types.ParamsKeeper
}

// NewKeeper returns a new BTC base keeper
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, paramsKeeper types.ParamsKeeper) *BaseKeeper {
	return &BaseKeeper{
		internalKeeper: internalKeeper{
			cdc:          cdc,
			storeKey:     storeKey,
			paramsKeeper: paramsKeeper,
		},
	}
}

// InitChains initializes all existing BTC chains and their respective param subspaces
func (k *BaseKeeper) InitChains(ctx sdk.Context) {
	if k.initialized {
		panic("chains are already initialized")
	}

	clog.Red("InitChains", "subspacePrefix", subspacePrefix)
	clog.Red("InitChains", "key", key.FromStr(subspacePrefix))

	baseStore := k.getBaseStore(ctx)
	fmt.Printf("baseStore: %+v\n", baseStore)
	iter := baseStore.IteratorNew(key.FromStr(subspacePrefix))
	fmt.Printf("iter: %+v\n", iter)
	defer utils.CloseLogError(iter, k.Logger(ctx))

	clog.Red("iter.Valid()", iter.Valid())

	for ; iter.Valid(); iter.Next() {
		clog.Red("iter.Value()", string(iter.Value()))
		_ = k.createSubspace(ctx, nexus.ChainName(iter.Value()))
	}

	k.initialized = true
}

// CreateChain creates the subspace for a new BTC chain. Returns an error if the chain already exists
func (k BaseKeeper) CreateChain(ctx sdk.Context, params types.Params) (err error) {
	defer func() {
		err = sdkerrors.Wrap(err, "cannot create new BTC chain")
	}()

	if !k.initialized {
		panic("InitChain must be called before chain keepers can be used")
	}

	if err := params.Validate(); err != nil {
		return err
	}
	chainKey := key.FromStr(subspacePrefix).Append(key.FromStr(params.ChainName.String()))
	if k.getBaseStore(ctx).HasNew(chainKey) {
		return fmt.Errorf("chain %s already exists", params.ChainName)
	}

	k.getBaseStore(ctx).SetRawNew(chainKey, []byte(params.ChainName.String()))

	k.createSubspace(ctx, params.ChainName).SetParamSet(ctx, &params)

	return nil
}

// ForChain returns the keeper associated to the given chain
func (k BaseKeeper) ForChain(ctx sdk.Context, chain nexus.ChainName) (types.ChainKeeper, error) {
	if !k.initialized {
		panic("InitChain must be called before chain keepers can be used")
	}

	return k.forChain(ctx, chain)
}

func (k BaseKeeper) forChain(ctx sdk.Context, chain nexus.ChainName) (chainKeeper, error) {
	chainKey := key.FromStr(subspacePrefix).Append(key.From(chain))
	if !k.getBaseStore(ctx).HasNew(chainKey) {
		clog.Red("chainKey", chainKey)
		return chainKeeper{}, fmt.Errorf("unknown chain %s", chain)
	}

	return chainKeeper{
		internalKeeper: k.internalKeeper,
		chain:          chain,
	}, nil
}

func (k BaseKeeper) createSubspace(ctx sdk.Context, chain nexus.ChainName) params.Subspace {
	chainKey := key.FromStr(types.ModuleName).Append(key.From(chain))
	k.Logger(ctx).Debug(fmt.Sprintf("initialized btc subspace %s", chain))
	return k.paramsKeeper.Subspace(chainKey.String()).WithKeyTable(types.KeyTable())
}

// Logger returns a module-specific logger.
func (k internalKeeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k internalKeeper) getBaseStore(ctx sdk.Context) utils.KVStore {
	return utils.NewNormalizedStore(ctx.KVStore(k.storeKey), k.cdc)
}
