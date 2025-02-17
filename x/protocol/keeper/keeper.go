package keeper

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/x/nexus/exported"
	pexported "github.com/scalarorg/scalar-core/x/protocol/exported"
	"github.com/scalarorg/scalar-core/x/protocol/types"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	protocolPrefix = utils.KeyFromStr("protocol")
)

type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	paramSpace paramtypes.Subspace
}

func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, paramSpace paramtypes.Subspace) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace,
	}
}

// GetParams gets the permission module's parameters
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.paramSpace.GetParamSet(ctx, &p)
	return p
}

// setParams sets the permission module's parameters
func (k Keeper) SetParams(ctx sdk.Context, p types.Params) {
	k.paramSpace.SetParamSet(ctx, &p)
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetProtocol(ctx sdk.Context, protocol *types.Protocol) {
	k.getStore(ctx).Set(protocolPrefix.Append(utils.KeyFromBz(protocol.ScalarAddress)), protocol)
}
func (k Keeper) SetProtocols(ctx sdk.Context, protocols []*types.Protocol) {
	store := k.getStore(ctx)
	for _, protocol := range protocols {
		store.Set(protocolPrefix.Append(utils.KeyFromBz(protocol.ScalarAddress)), protocol)
	}
}
func (k Keeper) GetAllProtocols(ctx sdk.Context) ([]*types.Protocol, bool) {
	store := k.getStore(ctx)
	protocols := []*types.Protocol{}
	iter := store.Iterator(protocolPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		protocol := types.Protocol{}
		iter.UnmarshalValue(&protocol)
		protocols = append(protocols, &protocol)
	}
	return protocols, true
}

func (k Keeper) findProtocols(ctx sdk.Context, req *types.ProtocolsRequest) ([]*types.Protocol, bool) {
	store := k.getStore(ctx)
	protocols := []*types.Protocol{}
	iter := store.Iterator(protocolPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		protocol := types.Protocol{}
		iter.UnmarshalValue(&protocol)
		if isMatch(&protocol, req) {
			protocols = append(protocols, &protocol)
		}
	}
	return protocols, true
}

func (k Keeper) getProtocolByAddress(ctx sdk.Context, address []byte) (*types.Protocol, bool) {
	store := k.getStore(ctx)
	iter := store.Iterator(protocolPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		protocol := types.Protocol{}
		iter.UnmarshalValue(&protocol)
		if bytes.Compare(protocol.ScalarAddress, address) == 0 {
			return &protocol, true
		}
	}
	return nil, false
}

/*
 * In scalar each asset is defined uniquely by its original chain (bitcoin networks: mainnet or testnets) and name.
 * This function finds the protocol that supports the given asset.
 */
func (k Keeper) FindProtocolByExternalSymbol(ctx sdk.Context, originChain exported.ChainName, minorChain exported.ChainName, symbol string) (*pexported.ProtocolInfo, error) {
	//ctx := sdk.UnwrapSDKContext(c)

	protocols, ok := k.GetAllProtocols(ctx)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "protocol not found")
	}
	for _, protocol := range protocols {
		if originChain == protocol.Asset.Chain && symbol == protocol.Asset.Name {
			//Check if the minor chain is supported by the protocol
			for _, chain := range protocol.Chains {
				if chain.Chain == minorChain {
					return protocol.ToProtocolInfo(), nil
				}
			}
		}
	}

	return nil, status.Errorf(codes.NotFound, "protocol with asset %s on the chain %s does not support transfering to the minor chain %s", symbol, originChain, minorChain)
}

func (k Keeper) FindProtocolByInternalAddress(ctx sdk.Context, originChain exported.ChainName, minorChain exported.ChainName, internalAddress string) (*pexported.ProtocolInfo, error) {
	protocols, ok := k.GetAllProtocols(ctx)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "protocol not found")
	}

	for _, protocol := range protocols {
		if originChain == protocol.Asset.Chain {
			err := protocol.IsAssetSupported(minorChain, internalAddress)
			if err != nil {
				k.Logger(ctx).Debug("[WARNING] checking if asset is supported", "error", err)
				continue
			}
			return protocol.ToProtocolInfo(), nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "protocol with origin chain %s does not support transfering to the token address %s on the minor chain %s",
		originChain, internalAddress, minorChain)
}

// Todo: Implement Matching function
func isMatch(protocol *types.Protocol, req *types.ProtocolsRequest) bool {
	if req.Name != "" {
		return protocol.Name == req.Name
	}
	return true
}
func (k Keeper) getStore(ctx sdk.Context) utils.KVStore {
	return utils.NewNormalizedStore(ctx.KVStore(k.storeKey), k.cdc)
}
