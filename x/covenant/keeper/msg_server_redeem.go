package keeper

import (
	"context"

	types "github.com/scalarorg/scalar-core/x/covenant/types"
)

func (s msgServer) UpdateNewBtcBlock(ctx context.Context, req *types.UpdateNewBtcBlockRequest) (*types.UpdateNewBtcBlockResponse, error) {
	return &types.UpdateNewBtcBlockResponse{}, nil
}

func (s msgServer) ConfirmRedeemTx(ctx context.Context, req *types.ConfirmRedeemTxRequest) (*types.ConfirmRedeemTxResponse, error) {
	return &types.ConfirmRedeemTxResponse{}, nil
}

func (s msgServer) ConfirmSwitchedPhase(ctx context.Context, req *types.ConfirmSwitchedPhaseRequest) (*types.ConfirmSwitchedPhaseResponse, error) {
	return &types.ConfirmSwitchedPhaseResponse{}, nil
}
