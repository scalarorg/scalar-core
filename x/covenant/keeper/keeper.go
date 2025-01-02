package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	"github.com/tendermint/tendermint/libs/log"
)

var (
	custodianPrefix      = utils.KeyFromStr("covenantCustodian")
	custodianGroupPrefix = utils.KeyFromStr("covenantCustodianGroup")
	keyPrefix            = utils.KeyFromStr("key")
	keyEpochPrefix       = utils.KeyFromStr("keyEpoch")
	subspacePrefix       = "subspace"
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

func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.Params{}
}

func (k *Keeper) SetCovenantRouter(router types.CovenantRouter) {
	if k.covRouter != nil {
		panic("router already set")
	}

	k.covRouter = router

	// In order to avoid invalid or non-deterministic behavior, we seal the router immediately
	// to prevent additional handlers from being registered after the keeper is initialized.
	k.covRouter.Seal()
}

// GetCovenantRouter returns the covenant router. If no router was set, it returns a (sealed) router with no handlers
func (k Keeper) GetCovenantRouter() types.CovenantRouter {
	if k.covRouter == nil {
		k.SetCovenantRouter(types.NewCovenantRouter())
	}

	return k.covRouter
}

func (k Keeper) getStore(ctx sdk.Context) utils.KVStore {
	return utils.NewNormalizedStore(ctx.KVStore(k.storeKey), k.cdc)
}

func (k Keeper) CreateAndSignPsbt(ctx sdk.Context, keyID multisig.KeyID, payloadHash multisig.Hash, module string, moduleMetadata ...codec.ProtoMarshaler) error {
	return nil
}
