package keeper

import (
	"context"

	"github.com/scalarorg/scalar-core/x/protocol/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns a new msg server instance
func NewMsgServerImpl(keeper Keeper) types.MsgServiceServer {
	return msgServer{Keeper: keeper}
}

func (s msgServer) CreateProtocol(c context.Context, req *types.CreateProtocolRequest) (*types.CreateProtocolResponse, error) {
	// ctx := sdk.UnwrapSDKContext(c)

	// if _, ok := s.getGovAccount(ctx, req.GovernanceKey.Address().Bytes()); ok {
	// 	return nil, fmt.Errorf("account is already registered with a role")
	// }

	// s.setGovernanceKey(ctx, req.GovernanceKey)
	// // delete the existing governance account address
	// s.deleteGovAccount(ctx, req.Sender)

	// s.setGovAccount(ctx, types.NewGovAccount(req.GovernanceKey.Address().Bytes(), exported.ROLE_ACCESS_CONTROL))

	return &types.CreateProtocolResponse{}, nil
}

// RegisterController handles register a controller account
func (s msgServer) UpdateProtocol(c context.Context, req *types.UpdateProtocolRequest) (*types.UpdateProtocolResponse, error) {
	// ctx := sdk.UnwrapSDKContext(c)

	// if _, ok := s.getGovAccount(ctx, req.Controller); ok {
	// 	return nil, fmt.Errorf("account is already registered with a role")
	// }

	// s.setGovAccount(ctx, types.NewGovAccount(req.Controller, exported.ROLE_CHAIN_MANAGEMENT))

	return &types.UpdateProtocolResponse{}, nil
}

// DeregisterController handles delete a controller account
func (s msgServer) AddSupportedChain(c context.Context, req *types.AddSupportedChainRequest) (*types.AddSupportedChainResponse, error) {
	// ctx := sdk.UnwrapSDKContext(c)

	// if s.GetRole(ctx, req.Controller) == exported.ROLE_CHAIN_MANAGEMENT {
	// 	s.deleteGovAccount(ctx, req.Controller)
	// }

	return &types.AddSupportedChainResponse{}, nil
}

func (s msgServer) UpdateSupportedChain(c context.Context, req *types.UpdateSupportedChainRequest) (*types.UpdateSupportedChainResponse, error) {
	return &types.UpdateSupportedChainResponse{}, nil
}
