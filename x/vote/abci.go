package vote

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/x/vote/exported"
	"github.com/scalarorg/scalar-core/x/vote/keeper"
	"github.com/scalarorg/scalar-core/x/vote/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(_ sdk.Context, _ abci.RequestBeginBlock, _ keeper.Keeper) {}

func handleCompletedPolls(ctx sdk.Context, k types.Voter) error {
	pollQueue := k.GetPollQueue(ctx)
	hasPollCompleted := func(value codec.ProtoMarshaler) bool {
		return value.(*exported.PollMetadata).State == exported.Completed
	}

	endBlockerLimit := k.GetParams(ctx).EndBlockerLimit
	handledPolls := int64(0)
	var pollMetadata exported.PollMetadata
	for handledPolls < endBlockerLimit && pollQueue.DequeueIf(&pollMetadata, hasPollCompleted) {
		handledPolls++

		pollID := pollMetadata.ID
		poll, ok := k.GetPoll(ctx, pollID)
		if !ok {
			panic(fmt.Errorf("poll %s not found", pollID))
		}

		logger := k.Logger(ctx).With("poll", pollID.String())

		voteHandler := k.GetVoteRouter().GetHandler(poll.GetModule())
		if voteHandler.IsFalsyResult(poll.GetResult()) {
			logger.Debug(fmt.Sprintf("poll %s completed with falsy result: %++v", pollID.String(), poll))
		} else {
			logger.Debug(fmt.Sprintf("poll %s completed with final result: %++v", pollID.String(), poll))
		}
		if err := voteHandler.HandleCompletedPoll(ctx, poll); err != nil {
			return err
		}

		k.DeletePoll(ctx, pollID)
	}
	k.Logger(ctx).Info(fmt.Sprintf("handled completed polls: %d at block %d", handledPolls, ctx.BlockHeight()))
	return nil
}
func handlePollsAtExpiry(ctx sdk.Context, k types.Voter) error {
	pollQueue := k.GetPollQueue(ctx)
	hasPollExpired := func(value codec.ProtoMarshaler) bool {
		return ctx.BlockHeight() >= value.(*exported.PollMetadata).ExpiresAt
	}

	endBlockerLimit := k.GetParams(ctx).EndBlockerLimit
	handledPolls := int64(0)
	var pollMetadata exported.PollMetadata
	for handledPolls < endBlockerLimit && pollQueue.DequeueIf(&pollMetadata, hasPollExpired) {
		handledPolls++

		pollID := pollMetadata.ID
		poll, ok := k.GetPoll(ctx, pollID)
		if !ok {
			panic(fmt.Errorf("poll %s not found", pollID))
		}

		logger := k.Logger(ctx).With("poll", pollID.String())

		voteHandler := k.GetVoteRouter().GetHandler(poll.GetModule())
		pollState := poll.GetState()
		clog.Magentaf(fmt.Sprintf("poll %s state: %++v", pollID.String(), pollState))
		switch pollState {
		case exported.Pending:
			logger.Debug(fmt.Sprintf("poll %s expired", pollID.String()))
			if err := voteHandler.HandleExpiredPoll(ctx, poll); err != nil {
				return err
			}

		case exported.Failed:
			logger.Debug(fmt.Sprintf("poll %s failed", pollID.String()))
			if err := voteHandler.HandleFailedPoll(ctx, poll); err != nil {
				return err
			}

		case exported.Completed:
			if voteHandler.IsFalsyResult(poll.GetResult()) {
				logger.Debug(fmt.Sprintf("poll %s completed with falsy result: %++v", pollID.String(), poll))
			} else {
				logger.Debug(fmt.Sprintf("poll %s completed with final result: %++v", pollID.String(), poll))
			}
			if err := voteHandler.HandleCompletedPoll(ctx, poll); err != nil {
				return err
			}
		default:
			panic(fmt.Errorf("unexpected poll state %s", poll.GetState().String()))
		}

		k.DeletePoll(ctx, pollID)
	}

	return nil
}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, k types.Voter) ([]abci.ValidatorUpdate, error) {
	clog.Yellow("[Vote] Abci Endblocker, BlockHeight: ", ctx.BlockHeight())
	// if err := handleCompletedPolls(ctx, k); err != nil {
	// 	return nil, err
	// }
	if err := handlePollsAtExpiry(ctx, k); err != nil {
		return nil, err
	}

	return nil, nil
}
