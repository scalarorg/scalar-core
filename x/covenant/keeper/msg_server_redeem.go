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
	clog.Greenf("UpdateUtxoForRedeemSession, req.ListOfUtxos: %v\n", req.ListOfUtxos)

	err := req.ValidateBasic()
	if err != nil {
		clog.Redf("UpdateUtxoForRedeemSession, ValidateBasic error: %v\n", err)
		return nil, err
	}

	clog.Bluef("UpdateUtxoForRedeemSession, req.Symbol: %s\n", req.Symbol)

	protocol, err := s.protocol.FindProtocolInfoByExternalSymbol(ctx, req.Symbol)
	if err != nil {
		clog.Redf("UpdateUtxoForRedeemSession, FindProtocolInfoByExternalSymbol error: %v\n", err)
		return nil, err
	}

	clog.Magentaf("UpdateUtxoForRedeemSession, protocol: %v\n", protocol)

	if !protocol.IsActivated() {
		clog.Redf("UpdateUtxoForRedeemSession, protocol %s is not activated\n", req.Symbol)
		return nil, fmt.Errorf("protocol %s is not activated", req.Symbol)
	}

	// TODO: Split to the confirmation event then handle in abci,
	// This way is just a temporary solution

	session, ok := s.GetRedeemSession(ctx, req.Symbol)
	if !ok {
		clog.Bluef("Check if session exist\n")
		session = types.DefaultRedeemSession()
		session.LastestBtcBlockHeight = req.BlockHeight
		session.Utxos = req.ListOfUtxos
		session.Symbol = req.Symbol
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

func (s msgServer) UpdateSequenceAndPhaseForRedeemSession(c context.Context, req *types.UpdateSequenceAndPhaseForRedeemSessionRequest) (*types.UpdateSequenceAndPhaseForRedeemSessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := req.ValidateBasic()
	if err != nil {
		clog.Redf("UpdateUtxoForRedeemSession, ValidateBasic error: %v\n", err)
		return nil, err
	}

	clog.Bluef("UpdateUtxoForRedeemSession, req.Symbol: %s\n", req.Symbol)

	protocol, err := s.protocol.FindProtocolInfoByExternalSymbol(ctx, req.Symbol)
	if err != nil {
		clog.Redf("UpdateUtxoForRedeemSession, FindProtocolInfoByExternalSymbol error: %v\n", err)
		return nil, err
	}

	clog.Magentaf("UpdateUtxoForRedeemSession, protocol: %v\n", protocol)

	if !protocol.IsActivated() {
		clog.Redf("UpdateUtxoForRedeemSession, protocol %s is not activated\n", req.Symbol)
		return nil, fmt.Errorf("protocol %s is not activated", req.Symbol)
	}

	// Check if session exist
	// TODO: Split to the confirmation event then handle in abci,
	// This way is just a temporary solution

	session, ok := s.GetRedeemSession(ctx, req.Symbol)
	if !ok {
		clog.Bluef("Check if session exist\n")
		session = types.DefaultRedeemSession()
		session.CurrentPhase = req.Phase
		session.Sequence = req.Sequence
		session.Symbol = req.Symbol
		s.Keeper.AddRedeemSession(ctx, req.Symbol, session)
		return &types.UpdateSequenceAndPhaseForRedeemSessionResponse{
			Session: session,
		}, nil
	}

	clog.Magentaf("UpdateUtxoForRedeemSession, session: %v\n", session)

	if session.Sequence > req.Sequence {
		return nil, fmt.Errorf("sequence is invalid")
	}

	if session.Sequence == req.Sequence && session.CurrentPhase >= req.Phase {
		return nil, fmt.Errorf("phase is invalid")
	}

	session.Sequence = req.Sequence
	session.CurrentPhase = req.Phase
	s.Keeper.AddRedeemSession(ctx, req.Symbol, session)

	return &types.UpdateSequenceAndPhaseForRedeemSessionResponse{
		Session: session,
	}, nil
}
