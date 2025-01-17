package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/scalarorg/scalar-core/x/protocol/types"
)

var _ types.QueryServer = Keeper{}

// GovernanceKey returns the multisig governance key
func (k Keeper) Protocols(c context.Context, req *types.ProtocolsRequest) (*types.ProtocolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	protocols, ok := k.findProtocols(ctx, req)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "protocols not found")
	}

	return &types.ProtocolsResponse{
		Protocols: protocols,
	}, nil
}

func (k Keeper) ProtocolAsset(c context.Context, req *types.ProtocolAssetRequest) (*types.ProtocolAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := req.ValidateBasic()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	protocols, ok := k.GetAllProtocols(ctx)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "protocol not found")
	}

	for _, protocol := range protocols {
		if req.SourceChain == protocol.Asset.Chain {
			err := protocol.IsAssetSupported(req.DestinationChain, req.TokenAddress)
			if err != nil {
				k.Logger(ctx).Error("error checking if asset is supported", "error", err)
				continue
			}
			return &types.ProtocolAssetResponse{
				Asset: protocol.Asset,
			}, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "protocol asset not found")
}
