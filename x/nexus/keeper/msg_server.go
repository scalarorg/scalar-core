package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/scalarorg/scalar-core/x/nexus/types"
	snapshot "github.com/scalarorg/scalar-core/x/snapshot/exported"
)

var _ types.MsgServiceServer = msgServer{}

const allChain = ":all:"
const wasmAsChain = ":wasm:"

type msgServer struct {
	types.Nexus
	snapshotter types.Snapshotter
	slashing    types.SlashingKeeper
	staking     types.StakingKeeper
	scalarnet   types.ScalarnetKeeper
}

// NewMsgServerImpl returns an implementation of the nexus MsgServiceServer interface for the provided Keeper.
func NewMsgServerImpl(k types.Nexus, snapshotter types.Snapshotter, slashing types.SlashingKeeper, staking types.StakingKeeper, scalarnet types.ScalarnetKeeper) types.MsgServiceServer {
	return msgServer{
		Nexus:       k,
		snapshotter: snapshotter,
		slashing:    slashing,
		staking:     staking,
		scalarnet:   scalarnet,
	}
}

func (s msgServer) RegisterChainMaintainer(c context.Context, req *types.RegisterChainMaintainerRequest) (*types.RegisterChainMaintainerResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	validator := s.snapshotter.GetOperator(ctx, req.Sender)
	if validator.Empty() {
		return nil, fmt.Errorf("account %v is not registered as a validator proxy", req.Sender.String())
	}

	valAddr := s.staking.Validator(ctx, validator)
	if valAddr == nil {
		return nil, fmt.Errorf("account %v is not registered as a validator", validator)
	}

	if !valAddr.IsBonded() {
		return nil, fmt.Errorf("validator %v is not bonded", validator)
	}

	for _, chainStr := range req.Chains {
		chain, ok := s.GetChain(ctx, chainStr)
		if !ok {
			s.Logger(ctx).Debug(fmt.Sprintf("'%s' is not a registered chain, skipping maintainer registration", chainStr))
			continue
		}

		if s.scalarnet.IsCosmosChain(ctx, chain.Name) {
			s.Logger(ctx).Debug(fmt.Sprintf("'%s' is a cosmos chain, skipping maintainer registration", chain.Name))
			continue
		}
		if s.IsChainMaintainer(ctx, chain, validator) {
			s.Logger(ctx).Info(fmt.Sprintf("'%s' is already a maintainer for chain '%s'", validator.String(), chain.Name))
			continue
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeChainMaintainer,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueRegister),
				sdk.NewAttribute(types.AttributeKeyChain, chain.Name.String()),
				sdk.NewAttribute(types.AttributeKeyChainMaintainerAddress, validator.String()),
			),
		)

		s.Logger(ctx).Info(fmt.Sprintf("validator %s registered maintainer for chain %s", validator.String(), chain.Name))
		if err := s.AddChainMaintainer(ctx, chain, validator); err != nil {
			return nil, err
		}
	}

	return &types.RegisterChainMaintainerResponse{}, nil
}

func (s msgServer) DeregisterChainMaintainer(c context.Context, req *types.DeregisterChainMaintainerRequest) (*types.DeregisterChainMaintainerResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	validator := s.snapshotter.GetOperator(ctx, req.Sender)
	if validator.Empty() {
		return nil, fmt.Errorf("account %v is not registered as a validator proxy", req.Sender.String())
	}

	for _, chainStr := range req.Chains {
		chain, ok := s.GetChain(ctx, chainStr)
		if !ok {
			return nil, fmt.Errorf("%s is not a registered chain", chainStr)
		}

		if !s.IsChainMaintainer(ctx, chain, validator) {
			continue
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeChainMaintainer,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueDeregister),
				sdk.NewAttribute(types.AttributeKeyChain, chain.Name.String()),
				sdk.NewAttribute(types.AttributeKeyChainMaintainerAddress, validator.String()),
			),
		)

		s.Logger(ctx).Info(fmt.Sprintf("validator %s deregistered maintainer for chain %s", validator.String(), chain.Name))
		if err := s.RemoveChainMaintainer(ctx, chain, validator); err != nil {
			return nil, err
		}
	}

	return &types.DeregisterChainMaintainerResponse{}, nil
}

func (s msgServer) ActivateChain(c context.Context, req *types.ActivateChainRequest) (*types.ActivateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	chains, doWasm := s.findRelevantChains(req.Chains, ctx)
	for _, chain := range chains {
		s.activateChain(ctx, chain)
	}
	if doWasm {
		s.ActivateWasmConnection(ctx)
	}

	return &types.ActivateChainResponse{}, nil
}

// DeactivateChain handles deactivate chains in case of emergencies
func (s msgServer) DeactivateChain(c context.Context, req *types.DeactivateChainRequest) (*types.DeactivateChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	chains, doWasm := s.findRelevantChains(req.Chains, ctx)

	for _, chain := range chains {
		s.deactivateChain(ctx, chain)
	}
	if doWasm {
		s.DeactivateWasmConnection(ctx)
	}

	return &types.DeactivateChainResponse{}, nil
}

func (s msgServer) findRelevantChains(chainNames []exported.ChainName, ctx sdk.Context) (chains []exported.Chain, doWasm bool) {
	doAllChains := strings.ToLower(chainNames[0].String()) == allChain
	doWasm = doAllChains || slices.Any(chainNames, chainIsWasm)

	if doAllChains {
		chains = s.GetChains(ctx)
	} else {
		for _, chain := range chainNames {
			chain, ok := s.GetChain(ctx, chain)
			if !ok {
				continue
			}
			chains = append(chains, chain)
		}
	}

	return chains, doWasm
}

func chainIsWasm(chain exported.ChainName) bool {
	return strings.ToLower(chain.String()) == wasmAsChain
}

func (s msgServer) activateChain(ctx sdk.Context, chain exported.Chain) {
	if s.IsChainActivated(ctx, chain) {
		s.Logger(ctx).Info(fmt.Sprintf("chain %s already activated", chain.Name))
		return
	}

	// no chain maintainer for cosmos chains
	if !s.scalarnet.IsCosmosChain(ctx, chain.Name) && !s.isActivationThresholdMet(ctx, chain) {
		return
	}

	s.Nexus.ActivateChain(ctx, chain)

	s.Logger(ctx).Info(fmt.Sprintf("chain %s activated", chain.Name))
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChain,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueActivated),
			sdk.NewAttribute(types.AttributeKeyChain, chain.Name.String()),
		),
	)
}

func (s msgServer) deactivateChain(ctx sdk.Context, chain exported.Chain) {
	if !s.IsChainActivated(ctx, chain) {
		s.Logger(ctx).Info(fmt.Sprintf("chain %s already deactivated", chain.Name))
		return
	}

	s.Nexus.DeactivateChain(ctx, chain)

	s.Logger(ctx).Info(fmt.Sprintf("chain %s deactivated", chain.Name))
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChain,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueDeactivated),
			sdk.NewAttribute(types.AttributeKeyChain, chain.Name.String()),
		),
	)
}

func (s msgServer) isActivationThresholdMet(ctx sdk.Context, chain exported.Chain) bool {
	isTombstoned := func(v snapshot.ValidatorI) bool {
		consAdd, err := v.GetConsAddr()
		if err != nil {
			return true
		}

		return s.slashing.IsTombstoned(ctx, consAdd)
	}

	isProxyActive := func(v snapshot.ValidatorI) bool {
		_, isActive := s.snapshotter.GetProxy(ctx, v.GetOperator())

		return isActive
	}

	filter := funcs.And(
		snapshot.ValidatorI.IsBonded,
		funcs.Not(snapshot.ValidatorI.IsJailed),
		funcs.Not(isTombstoned),
		isProxyActive,
	)

	params := s.Nexus.GetParams(ctx)

	_, err := s.snapshotter.CreateSnapshot(
		ctx,
		s.Nexus.GetChainMaintainers(ctx, chain),
		filter,
		snapshot.QuadraticWeightFunc,
		params.ChainActivationThreshold,
	)
	if err != nil {
		s.Logger(ctx).Info(fmt.Sprintf("activation threshold is not met for %s due to: %s", chain.Name, err.Error()))
	}

	return err == nil
}

func (s msgServer) RegisterAssetFee(c context.Context, req *types.RegisterAssetFeeRequest) (*types.RegisterAssetFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	chain, ok := s.GetChain(ctx, req.FeeInfo.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.FeeInfo.Chain)
	}

	if err := s.RegisterFee(ctx, chain, req.FeeInfo); err != nil {
		return nil, err
	}

	s.Logger(ctx).Info(fmt.Sprintf("registered fee info for asset %s on chain %s", req.FeeInfo.Asset, chain.Name), types.AttributeKeyChain, chain.Name, types.AttributeKeyAsset, req.FeeInfo.Asset)

	return &types.RegisterAssetFeeResponse{}, nil
}

// SetTransferRateLimit handles setting the transfer rate limit for an asset on a chain
func (s msgServer) SetTransferRateLimit(c context.Context, req *types.SetTransferRateLimitRequest) (*types.SetTransferRateLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if err := s.SetRateLimit(ctx, req.Chain, req.Limit, req.Window); err != nil {
		return nil, err
	}

	return &types.SetTransferRateLimitResponse{}, nil
}
