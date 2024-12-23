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

	protocols, ok := k.GetProtocols(ctx)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "protocols not found")
	}

	return &types.ProtocolsResponse{
		Protocols: protocols,
	}, nil
}
