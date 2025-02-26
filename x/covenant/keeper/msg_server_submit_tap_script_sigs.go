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

	if err := signingSession.AddListOfTapScriptSigs(ctx.BlockHeight(), participant, req.ListOfTapScriptSigsMap); err != nil {
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

	events.Emit(ctx, types.NewTapscriptSigsSubmitted(req.SigID, participant, req.ListOfTapScriptSigsMap))

	return &types.SubmitTapScriptSigsResponse{}, nil
}

func logSigningSession(m types.SigningSession) {
	clog.Greenf("AddListOfTapScriptSigs, signing session, ID: %+v", m.ID)
	clog.Greenf("AddListOfTapScriptSigs, signing session, State: %+v", m.State)
	clog.Greenf("AddListOfTapScriptSigs, signing session, Key: %+v", m.Key)
	clog.Greenf("AddListOfTapScriptSigs, signing session, ExpiresAt: %+v", m.ExpiresAt)
	clog.Greenf("AddListOfTapScriptSigs, signing session, CompletedAt: %+v", m.CompletedAt)
	clog.Greenf("AddListOfTapScriptSigs, signing session, GracePeriod: %+v", m.GracePeriod)
	clog.Greenf("AddListOfTapScriptSigs, signing session, Module: %+v", m.Module)
	clog.Greenf("AddListOfTapScriptSigs, signing session, ModuleMetadata: %+v", m.ModuleMetadata)

	clog.Redf("AddListOfTapScriptSigs, PsbtMultiSig.KeyID: %+v", m.PsbtMultiSig.KeyID)
	// clog.Redf("AddListOfTapScriptSigs, PsbtMultiSig.Psbt: %+x", m.PsbtMultiSig.MultiPsbt[])
	for participant, list := range m.PsbtMultiSig.ParticipantListTapScriptSigs {
		for _, sigsMap := range list.Inner {
			clog.Redf("Participant: %s", participant)
			for _, tapScriptList := range sigsMap.Inner {
				for _, tapScriptSig := range tapScriptList.Sigs.List {
					clog.Redf("TapScriptSig, KeyXOnly: %+v", tapScriptSig.KeyXOnly)
					clog.Redf("TapScriptSig, Signature: %+v", tapScriptSig.Signature)
					clog.Redf("TapScriptSig, LeafHash: %+v", tapScriptSig.LeafHash)
				}
			}
		}
	}
}
