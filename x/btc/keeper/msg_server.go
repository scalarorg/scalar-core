package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

var _ types.MsgServiceServer = msgServer{}

type msgServer struct {
	types.BaseKeeper
}

func NewMsgServerImpl(keeper types.BaseKeeper) types.MsgServiceServer {
	return msgServer{
		BaseKeeper: keeper,
	}
}

func (s msgServer) ConfirmGatewayTxs(c context.Context, req *types.ConfirmGatewayTxsRequest) (*types.ConfirmGatewayTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	_ = ctx
	s.Logger(ctx).Info("ConfirmGatewayTxs", "req", req)
	return nil, nil
}
