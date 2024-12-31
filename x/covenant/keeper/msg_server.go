package keeper

import (
	"context"

	"github.com/scalarorg/scalar-core/x/covenant/types"
)

type msgServer struct {
	Keeper
	snapshotter types.Snapshotter
	staker      types.Staker
	nexus       types.Nexus
}

var _ types.MsgServiceServer = msgServer{}

type MsgServerConstructArgs struct {
	Keeper
	Snapshotter types.Snapshotter
	Staker      types.Staker
	Nexus       types.Nexus
}

// NewMsgServerImpl returns an implementation of the evm MsgServiceServer interface
// for the provided Keeper.
func NewMsgServerImpl(arg *MsgServerConstructArgs) types.MsgServiceServer {
	return msgServer{
		Keeper:      arg.Keeper,
		snapshotter: arg.Snapshotter,
		staker:      arg.Staker,
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
