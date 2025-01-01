package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServiceServer = Querier{}

// Querier implements the grpc querier
type Querier struct {
	keeper *Keeper
}

// NewGRPCQuerier returns a new Querier
func NewGRPCQuerier(k *Keeper) Querier {
	return Querier{
		keeper: k,
	}
}

// Get custodians
func (q Querier) Custodians(c context.Context, req *types.CustodiansRequest) (*types.CustodiansResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	custodians, ok := q.keeper.findCustodians(ctx, req)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "custodians not found")
	}

	return &types.CustodiansResponse{
		Custodians: custodians,
	}, nil
}

// Get custodian groups
func (q Querier) CustodianGroups(c context.Context, req *types.CustodianGroupsRequest) (*types.CustodianGroupsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	groups, ok := q.keeper.findCustodianGroups(ctx, req)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "custodian groups not found")
	}

	return &types.CustodianGroupsResponse{
		Groups: groups,
	}, nil
}

// Params returns the params of the module
func (q Querier) Params(context.Context, *types.ParamsRequest) (*types.ParamsResponse, error) {
	return nil, nil
}
