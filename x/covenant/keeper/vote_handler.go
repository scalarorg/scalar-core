package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/funcs"

	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	vote "github.com/scalarorg/scalar-core/x/vote/exported"
)

var _ vote.VoteHandler = &voteHandler{}

type voteHandler struct {
	cdc      codec.Codec
	keeper   types.Keeper
	nexus    types.Nexus
	rewarder types.Rewarder
}

// NewVoteHandler returns the handler for processing vote delivered by the vote module
func NewVoteHandler(cdc codec.Codec, keeper types.Keeper, nexus types.Nexus, rewarder types.Rewarder) vote.VoteHandler {
	return voteHandler{
		cdc:      cdc,
		keeper:   keeper,
		nexus:    nexus,
		rewarder: rewarder,
	}
}

func (v voteHandler) HandleFailedPoll(ctx sdk.Context, poll vote.Poll) error {
	md := mustGetMetadata(poll)
	events.Emit(ctx, &types.BasicPollFailed{
		Data:   md.Data,
		PollID: poll.GetID(),
		Chain:  md.Chain,
	})

	return nil
}

func (v voteHandler) IsFalsyResult(result codec.ProtoMarshaler) bool {
	return len(result.(*types.VoteEvents).Events) == 0
}

func (v voteHandler) HandleExpiredPoll(ctx sdk.Context, poll vote.Poll) error {
	rewardPoolName, ok := poll.GetRewardPoolName()
	if !ok {
		return fmt.Errorf("reward pool not set for poll %s", poll.GetID().String())
	}

	md := mustGetMetadata(poll)
	rewardPool := v.rewarder.GetPool(ctx, rewardPoolName)
	chain, ok := v.nexus.GetChain(ctx, md.Chain)
	if !ok {
		return fmt.Errorf("%s is not a registered chain", md.Chain)
	}
	// Penalize voters who failed to vote
	for _, voter := range poll.GetVoters() {
		hasVoted := poll.HasVoted(voter)
		if maintainerState, ok := v.nexus.GetChainMaintainerState(ctx, chain, voter); ok {
			maintainerState.MarkMissingVote(!hasVoted)
			funcs.MustNoErr(v.nexus.SetChainMaintainerState(ctx, maintainerState))

			msg := fmt.Sprintf("marked voter %s behaviour", voter.String())
			clog.Red("[Covenant] HandleExpiredPoll", msg)
			v.keeper.Logger(ctx).Debug(msg,
				"voter", voter.String(),
				"missing_vote", !hasVoted,
				"poll", poll.GetID().String(),
			)
		}

		if !hasVoted {
			rewardPool.ClearRewards(voter)
			msg := fmt.Sprintf("penalized voter %s due to timeout", voter.String())
			clog.Red("HandleExpiredPoll", msg)
			v.keeper.Logger(ctx).Debug(msg,
				"voter", voter.String(),
				"poll", poll.GetID().String())
		}
	}

	events.Emit(ctx, &types.BasicPollExpired{
		Data:   md.Data,
		Chain:  md.Chain,
		PollID: poll.GetID(),
	})

	return nil
}

func (v voteHandler) HandleCompletedPoll(ctx sdk.Context, poll vote.Poll) error {

	clog.Red("[Covenant] HandleCompletedPoll", "poll", poll.GetID().String())

	voteEvents := poll.GetResult().(*types.VoteEvents)

	chain, ok := v.nexus.GetChain(ctx, voteEvents.Chain)
	if !ok {
		return fmt.Errorf("%s is not a registered chain", voteEvents.Chain)
	}

	rewardPoolName, ok := poll.GetRewardPoolName()
	if !ok {
		return fmt.Errorf("reward pool not set for poll %s", poll.GetID().String())
	}

	rewardPool := v.rewarder.GetPool(ctx, rewardPoolName)

	for _, voter := range poll.GetVoters() {
		maintainerState, ok := v.nexus.GetChainMaintainerState(ctx, chain, voter)
		if !ok {
			continue // voter is no longer a chain maintainer, so recording the state is irrelevant
		}

		hasVoted := poll.HasVoted(voter)
		hasVotedIncorrectly := hasVoted && !poll.HasVotedCorrectly(voter)

		maintainerState.MarkMissingVote(!hasVoted)
		maintainerState.MarkIncorrectVote(hasVotedIncorrectly)
		funcs.MustNoErr(v.nexus.SetChainMaintainerState(ctx, maintainerState))

		msg := fmt.Sprintf("marked voter %s behaviour", voter.String())
		clog.Red("HandleCompletedPoll", msg)
		v.keeper.Logger(ctx).Debug(msg,
			"voter", voter.String(),
			"missing_vote", !hasVoted,
			"incorrect_vote", hasVotedIncorrectly,
			"poll", poll.GetID().String(),
		)

		switch {
		case hasVotedIncorrectly, !hasVoted:
			rewardPool.ClearRewards(voter)
			msg := fmt.Sprintf("penalized voter %s due to incorrect vote or missing vote", voter.String())
			clog.Red("HandleCompletedPoll", msg)
			v.keeper.Logger(ctx).Debug(msg,
				"voter", voter.String(),
				"poll", poll.GetID().String())
		default:
			if err := rewardPool.ReleaseRewards(voter); err != nil {
				return err
			}
			msg := fmt.Sprintf("released rewards for voter %s", voter.String())
			clog.Red("[Covenant] HandleCompletedPoll", msg)
			v.keeper.Logger(ctx).Debug(msg,
				"voter", voter.String(),
				"poll", poll.GetID().String())
		}
	}

	md := mustGetMetadata(poll)
	if v.IsFalsyResult(voteEvents) {
		events.Emit(ctx, &types.BasicPollNoEventsConfirmed{
			Data:   md.Data,
			Chain:  md.Chain,
			PollID: poll.GetID(),
		})
	}

	event := &types.BasicPollCompleted{
		Data:   md.Data,
		Chain:  md.Chain,
		PollID: poll.GetID(),
	}

	clog.Red("[Covenant] Poll Completed Event", event)

	events.Emit(ctx, event)

	return nil
}

func (v voteHandler) HandleResult(ctx sdk.Context, result codec.ProtoMarshaler) error {
	voteEvents := result.(*types.VoteEvents)

	if v.IsFalsyResult(result) {
		return nil
	}

	for _, event := range voteEvents.Events {
		if err := v.handleEvent(ctx, event, v.keeper); err != nil {
			return err
		}
	}

	return nil
}

func (v voteHandler) handleEvent(ctx sdk.Context, event types.Event, k types.Keeper) error {

	eventType := event.GetEvent()
	switch eventType.(type) {
	case *types.Event_RedeemTxsConfirmed:
		if err := k.EnqueueEvent(ctx, &event); err != nil {
			k.Logger(ctx).Error("failed to set pending redeem command", "error", err)
			return err
		}
	case *types.Event_SwitchedPhaseConfirmed:
		if err := k.EnqueueEvent(ctx, &event); err != nil {
			k.Logger(ctx).Error("failed to set pending switched phase command", "error", err)
			return err
		}
	case *types.Event_IntializeUtxoSnapshotCompleted:
		if err := k.EnqueueEvent(ctx, &event); err != nil {
			k.Logger(ctx).Error("failed to set pending initialize utxo snapshot command", "error", err)
			return err
		}

	default:
		return fmt.Errorf("unrecognized event type %T", eventType)
	}

	return nil
}

func mustGetMetadata(poll vote.Poll) types.BasicPollMetadata {
	md := funcs.MustOk(poll.GetMetaData())
	metadata, ok := md.(*types.BasicPollMetadata)
	if !ok {
		panic(fmt.Sprintf("poll metadata should be of type %T", &types.BasicPollMetadata{}))
	}
	return *metadata
}
