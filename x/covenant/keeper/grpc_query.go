package keeper

import (
	"context"

	"github.com/scalarorg/scalar-core/x/covenant/types"
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
func (q Querier) GetCustodians(context.Context, *types.CustodiansRequest) (*types.CustodiansResponse, error) {
	return nil, nil
}

// Get custodian groups
func (q Querier) CustodianGroups(context.Context, *types.CustodianGroupsRequest) (*types.CustodianGroupsResponse, error) {
	return nil, nil
}

// Params returns the params of the module
func (q Querier) Params(context.Context, *types.ParamsRequest) (*types.ParamsResponse, error) {
	return nil, nil
}
