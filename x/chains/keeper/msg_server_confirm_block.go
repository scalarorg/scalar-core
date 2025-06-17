package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/x/chains/types"
)

func (s msgServer) ConfirmBlock(c context.Context, req *types.ConfirmBlockRequest) (*types.ConfirmBlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	chain, ok := s.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
		return nil, err
	}

	keeper, err := s.ForChain(ctx, chain.Name)
	if err != nil {
		return nil, err
	}

	currentBlockHash, err := keeper.GetCurrentBlock(ctx)
	if err != nil {
		return nil, err
	}

	pollParticipants, err := s.initializePoll(ctx, chain, req.BlockHash)
	if err != nil {
		return nil, err
	}

	events.Emit(ctx, &types.ConfirmNewBlockStarted{
		Chain:              chain.Name,
		BlockHash:          req.BlockHash,
		PreviousBlockHash:  currentBlockHash.PreviousBlockHash,
		ConfirmationHeight: keeper.GetRequiredConfirmationHeight(ctx),
		PollParticipants:   pollParticipants,
	})
	return &types.ConfirmBlockResponse{}, nil
}
