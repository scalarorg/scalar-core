package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
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

func (s msgServer) UpdateUtxoForRedeemSession(c context.Context, req *types.UpdateUtxoForRedeemSessionRequest) (*types.UpdateUtxoForRedeemSessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	clog.Magentaf("UpdateUtxoForRedeemSession, req: %v\n", req)

	err := req.ValidateBasic()
	if err != nil {
		return nil, err
	}

	protocol, err := s.protocol.FindProtocolInfoByExternalSymbol(ctx, req.Symbol)
	if err != nil {
		clog.Redf("UpdateUtxoForRedeemSession, FindProtocolInfoByExternalSymbol error: %v\n", err)
		return nil, err
	}

	if !protocol.IsActivated() {
		clog.Redf("UpdateUtxoForRedeemSession, protocol %s is not activated\n", req.Symbol)
		return nil, fmt.Errorf("protocol %s is not activated", req.Symbol)
	}

	session, ok := s.GetRedeemSession(ctx, req.Symbol)
	if !ok {
		session = types.DefaultRedeemSession()
		session.LastestBtcBlockHeight = req.BlockHeight
		session.Utxos = req.ListOfUtxos
		s.Keeper.AddRedeemSession(ctx, req.Symbol, session)
		return &types.UpdateUtxoForRedeemSessionResponse{
			Session: session,
		}, nil
	}

	clog.Magentaf("UpdateUtxoForRedeemSession, session: %v\n", session)

	if session.LastestBtcBlockHeight >= req.BlockHeight {
		return nil, fmt.Errorf("utxos already exist")
	}

	session.LastestBtcBlockHeight = req.BlockHeight
	session.Utxos = req.ListOfUtxos
	s.Keeper.AddRedeemSession(ctx, req.Symbol, session)

	return &types.UpdateUtxoForRedeemSessionResponse{
		Session: session,
	}, nil
}
