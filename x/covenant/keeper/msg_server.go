package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	snapshot "github.com/scalarorg/scalar-core/x/snapshot/exported"
)

type msgServer struct {
	Keeper
	snapshotter types.Snapshotter
	staker      types.StakingKeeper
	slashing    types.SlashingKeeper
	nexus       types.Nexus
}

var _ types.MsgServiceServer = msgServer{}

type MsgServerConstructArgs struct {
	Keeper
	Snapshotter types.Snapshotter
	Staker      types.StakingKeeper
	Slashing    types.SlashingKeeper
	Nexus       types.Nexus
}

// NewMsgServerImpl returns an implementation of the evm MsgServiceServer interface
// for the provided Keeper.
func NewMsgServerImpl(arg *MsgServerConstructArgs) types.MsgServiceServer {
	return msgServer{
		Keeper:      arg.Keeper,
		snapshotter: arg.Snapshotter,
		staker:      arg.Staker,
		slashing:    arg.Slashing,
		nexus:       arg.Nexus,
	}
}

// Create custodian
func (s msgServer) CreateCustodian(context.Context, *types.CreateCustodianRequest) (*types.CreateCustodianResponse, error) {
	return &types.CreateCustodianResponse{}, nil
}

// Update custodian
func (s msgServer) UpdateCustodian(context.Context, *types.UpdateCustodianRequest) (*types.UpdateCustodianResponse, error) {
	return &types.UpdateCustodianResponse{}, nil
}

// Create custodian group
func (s msgServer) CreateCustodianGroup(context.Context, *types.CreateCustodianGroupRequest) (*types.CreateCustodianGroupResponse, error) {
	return &types.CreateCustodianGroupResponse{}, nil
}

// Update Custodian group
func (s msgServer) UpdateCustodianGroup(context.Context, *types.UpdateCustodianGroupRequest) (*types.UpdateCustodianGroupResponse, error) {
	return &types.UpdateCustodianGroupResponse{}, nil
}

// Add Custodian to custodian group
// recalculate taproot pubkey when adding custodian to custodian group
func (s msgServer) AddCustodianToGroup(context.Context, *types.AddCustodianToGroupRequest) (*types.CustodianToGroupResponse, error) {
	return &types.CustodianToGroupResponse{}, nil
}

// Remove Custodian from custodian group
// recalculate taproot address when deleting custodian from custodian group
func (s msgServer) RemoveCustodianFromGroup(context.Context, *types.RemoveCustodianFromGroupRequest) (*types.CustodianToGroupResponse, error) {
	return &types.CustodianToGroupResponse{}, nil
}

// Rotate key for each custodian group
func (s msgServer) RotateKey(c context.Context, req *types.RotateKeyRequest) (*types.RotateKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	chain, ok := s.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("unknown chain")
	}
	custodianGroups, ok := s.Keeper.GetAllCustodianGroups(ctx)
	if !ok {
		return nil, fmt.Errorf("no custodian groups found")
	}
	for _, group := range custodianGroups {
		threshold := utils.Threshold{
			Numerator:   int64(group.Quorum),
			Denominator: int64(len(group.Custodians)),
		}
		snapshot, err := s.createSnapshot(ctx, chain, threshold)
		if err != nil {
			return nil, err
		}
		key := group.CreateKey(ctx, snapshot, threshold)
		if err := s.Keeper.RotateKey(ctx, req.Chain, key); err != nil {
			return nil, sdkerrors.Wrap(err, "failed to rotate the next key")
		}
	}
	return &types.RotateKeyResponse{}, nil
}

func (s msgServer) createSnapshot(ctx sdk.Context, chain nexus.Chain, threshold utils.Threshold) (snapshot.Snapshot, error) {
	candidates := s.nexus.GetChainMaintainers(ctx, chain)

	return s.snapshotter.CreateSnapshot(
		ctx,
		candidates,
		excludeJailedOrTombstoned(ctx, s.slashing, s.snapshotter),
		snapshot.QuadraticWeightFunc,
		threshold,
	)
}

func excludeJailedOrTombstoned(ctx sdk.Context, slashing types.SlashingKeeper, snapshotter types.Snapshotter) func(v snapshot.ValidatorI) bool {
	isTombstoned := func(v snapshot.ValidatorI) bool {
		consAdd, err := v.GetConsAddr()
		if err != nil {
			return true
		}

		return slashing.IsTombstoned(ctx, consAdd)
	}

	isProxyActive := func(v snapshot.ValidatorI) bool {
		_, isActive := snapshotter.GetProxy(ctx, v.GetOperator())

		return isActive
	}

	return funcs.And(
		snapshot.ValidatorI.IsBonded,
		funcs.Not(snapshot.ValidatorI.IsJailed),
		funcs.Not(isTombstoned),
		isProxyActive,
	)
}
