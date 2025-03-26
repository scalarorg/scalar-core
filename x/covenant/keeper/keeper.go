package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/key"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/tendermint/tendermint/libs/log"
)

const eventsQueue = "events_queue"

var (
	signingPrefix        = utils.KeyFromInt(1)
	keyPrefix            = utils.KeyFromInt(2)
	expirySigningPrefix  = utils.KeyFromInt(3)
	custodianPrefix      = utils.KeyFromInt(4)
	custodianGroupPrefix = utils.KeyFromInt(5)
	redeemSessionPrefix  = utils.KeyFromInt(6)
	utxoSnapshotPrefix   = utils.KeyFromInt(7)

	signingSessionCountKey = utils.KeyFromInt(100)

	reserveUtxoCommandPrefix = key.FromStr("reserve_utxo_command")
	unsignedStandaloneIDKey  = key.FromStr("unsigned_standalone_command_id")
)

var _ types.Keeper = &Keeper{}

// BaseKeeper implements the base Keeper

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	paramSpace paramtypes.Subspace
	covRouter  types.CovenantRouter
}

// NewKeeper returns a new EVM base keeper
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, paramSpace paramtypes.Subspace) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace.WithKeyTable(types.KeyTable()),
	}
}

// Logger returns a module-specific logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) setParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

func (k Keeper) getStore(ctx sdk.Context) utils.KVStore {
	return utils.NewNormalizedStore(ctx.KVStore(k.storeKey), k.cdc)
}

func (k Keeper) getStoreIterator(ctx sdk.Context, prefix utils.StringKey) utils.Iterator {
	store := k.getStore(ctx)
	return store.Iterator(prefix)
}

func (k Keeper) GetEventsQueue(ctx sdk.Context) utils.BlockHeightKVQueue {
	return utils.NewBlockHeightKVQueue(
		eventsQueue,
		k.getStore(ctx),
		ctx.BlockHeight(),
		k.Logger(ctx),
	)
}

func (k Keeper) EnqueueEvent(ctx sdk.Context, event *types.Event) error {
	// TODO: Critical - fix the key of thevent

	if k.getStore(ctx).HasNew(key.FromStr(event.Chain.String()).Append(key.FromStr(event.Hash.Hex()))) {
		return fmt.Errorf("event %s already exists", *event.Hash)
	}

	k.GetEventsQueue(ctx).Enqueue(utils.LowerCaseKey(event.Chain.String()).AppendStr(event.Hash.Hex()), event)
	return nil
}
