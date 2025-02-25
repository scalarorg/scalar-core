package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/protocol/types"
)

type msgServer struct {
	Keeper
	covenant types.CovenantKeeper
}

// NewMsgServerImpl returns a new msg server instance
func NewMsgServerImpl(keeper Keeper, covenant types.CovenantKeeper) types.MsgServer {
	return msgServer{Keeper: keeper, covenant: covenant}
}

func (s msgServer) CreateProtocol(c context.Context, req *types.CreateProtocolRequest) (*types.CreateProtocolResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	custodianGr, ok := s.covenant.GetCustodianGroup(ctx, req.CustodianGroupUid)
	if !ok {
		return nil, fmt.Errorf("custodian group not found")
	}

	protocol := types.Protocol{
		BitcoinPubkey:     req.BitcoinPubkey,
		ScalarAddress:     req.Sender.Bytes(),
		ScalarPubkey:      req.ScalarPubkey,
		Name:              req.Name,
		Tag:               []byte(req.Tag), // ascii
		Attribute:         req.Attribute,
		Status:            types.Pending,
		CustodianGroupUID: custodianGr.UID,
		Avatar:            req.Avatar,
	}

	s.Keeper.SetProtocol(ctx, &protocol)

	return &types.CreateProtocolResponse{
		Protocol: &protocol,
	}, nil
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
