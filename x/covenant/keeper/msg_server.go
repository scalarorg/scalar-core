package keeper

import (
	"context"

	"github.com/scalarorg/scalar-core/x/covenant/types"
)

var _ types.MsgServiceServer = msgServer{}

type msgServer struct {
	types.BaseKeeper
	// nexus          types.Nexus
	voter       types.Voter
	snapshotter types.Snapshotter
	staking     types.StakingKeeper
	slashing    types.SlashingKeeper
	// multisigKeeper types.MultisigKeeper
}

// NewMsgServerImpl returns an implementation of the evm MsgServiceServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper types.BaseKeeper, v types.Voter, snap types.Snapshotter, staking types.StakingKeeper, slashing types.SlashingKeeper) types.MsgServiceServer {
	return msgServer{
		BaseKeeper:  keeper,
		voter:       v,
		snapshotter: snap,
		staking:     staking,
		slashing:    slashing,
		// multisigKeeper: multisigKeeper,
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
func (s msgServer) AddCustodianToGroup(context.Context, *types.CustodianToGroupRequest) (*types.CustodianToGroupResponse, error) {
	return &types.CustodianToGroupResponse{}, nil
}

// Remove Custodian from custodian group
// recalculate taproot address when deleting custodian from custodian group
func (s msgServer) RemoveCustodianFromGroup(context.Context, *types.CustodianToGroupRequest) (*types.CustodianToGroupResponse, error) {
	return &types.CustodianToGroupResponse{}, nil
}
