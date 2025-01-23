package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
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

	if err := signingSession.AddTapScriptSigs(ctx.BlockHeight(), participant, req.TapScriptSigsMap); err != nil {
		return nil, sdkerrors.Wrap(err, "unable to add signature for signing")
	}

	logSigningSession(signingSession)

	s.setSigningSession(ctx, signingSession)

	s.Logger(ctx).Info("new signature submitted",
		"sig_id", signingSession.GetID(),
		"participant", participant.String(),
		"participants_weight", signingSession.GetParticipantsWeight().String(),
		"bonded_weight", signingSession.Key.Snapshot.BondedWeight.String(),
		"signing_threshold", signingSession.Key.SigningThreshold.String(),
		"expires_at", signingSession.ExpiresAt,
	)

	events.Emit(ctx, types.NewTapscriptSigsSubmitted(req.SigID, participant, req.TapScriptSigsMap))

	return &types.SubmitTapScriptSigsResponse{}, nil
}

func logSigningSession(m types.SigningSession) {
	clog.Greenf("AddTapScriptSigs, signing session, ID: %+v", m.ID)
	clog.Greenf("AddTapScriptSigs, signing session, State: %+v", m.State)
	clog.Greenf("AddTapScriptSigs, signing session, Key: %+v", m.Key)
	clog.Greenf("AddTapScriptSigs, signing session, ExpiresAt: %+v", m.ExpiresAt)
	clog.Greenf("AddTapScriptSigs, signing session, CompletedAt: %+v", m.CompletedAt)
	clog.Greenf("AddTapScriptSigs, signing session, GracePeriod: %+v", m.GracePeriod)
	clog.Greenf("AddTapScriptSigs, signing session, Module: %+v", m.Module)
	clog.Greenf("AddTapScriptSigs, signing session, ModuleMetadata: %+v", m.ModuleMetadata)

	clog.Redf("AddTapScriptSigs, PsbtMultiSig.KeyID: %+v", m.PsbtMultiSig.KeyID)
	clog.Redf("AddTapScriptSigs, PsbtMultiSig.Psbt: %+x", m.PsbtMultiSig.Psbt)
	for participant, list := range m.PsbtMultiSig.ParticipantTapScriptSigs {
		clog.Redf("Participant: %s", participant)
		for _, tapScriptList := range list.Inner {
			for _, tapScriptSig := range tapScriptList.List {
				clog.Redf("TapScriptSig, KeyXOnly: %+v", tapScriptSig.KeyXOnly)
				clog.Redf("TapScriptSig, Signature: %+v", tapScriptSig.Signature)
				clog.Redf("TapScriptSig, LeafHash: %+v", tapScriptSig.LeafHash)
			}
		}
	}
}
