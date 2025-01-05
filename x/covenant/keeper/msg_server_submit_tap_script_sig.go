package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/events"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
)

func (s msgServer) SubmitTapScriptSigs(c context.Context, req *types.SubmitTapScriptSigsRequest) (*types.SubmitTapScriptSigsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	signingSession, ok := s.getSigningSession(ctx, req.SigID)
	if !ok {
		return nil, fmt.Errorf("signing session %d not found", req.SigID)
	}

	participant := s.snapshotter.GetOperator(ctx, req.Sender)
	if participant.Empty() {
		return nil, fmt.Errorf("sender %s is not a registered proxy", req.Sender.String())
	}

	if err := signingSession.AddTapScriptSigs(ctx.BlockHeight(), participant, req.TapScriptSigs); err != nil {
		return nil, sdkerrors.Wrap(err, "unable to add signature for signing")
	}

	s.setSigningSession(ctx, signingSession)

	s.Logger(ctx).Debug("new signature submitted",
		"sig_id", signingSession.GetID(),
		"participant", participant.String(),
		"participants_weight", signingSession.GetParticipantsWeight().String(),
		"bonded_weight", signingSession.Key.Snapshot.BondedWeight.String(),
		"signing_threshold", signingSession.Key.SigningThreshold.String(),
		"expires_at", signingSession.ExpiresAt,
	)

	events.Emit(ctx, types.NewTapscriptSigsSubmitted(req.SigID, participant, req.TapScriptSigs))

	return &types.SubmitTapScriptSigsResponse{}, nil
}
