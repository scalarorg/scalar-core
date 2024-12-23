package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/x/chains/types"
)

func (s msgServer) ConfirmDestTxs(c context.Context, req *types.ConfirmDestTxsRequest) (*types.ConfirmDestTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	chain, ok := s.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
		return nil, err
	}

	clog.Red("After validateChainActivated", chain)

	keeper, err := s.ForChain(ctx, chain.Name)
	if err != nil {
		return nil, err
	}

	snapshot, err := s.createSnapshot(ctx, chain)
	if err != nil {
		return nil, err
	}

	pollMappings, err := s.initializePolls(ctx, chain, snapshot, req.TxIDs)
	if err != nil {
		return nil, err
	}

	event := &types.EventConfirmDestTxsStarted{
		Chain:              chain.Name,
		PollMappings:       pollMappings,
		ConfirmationHeight: keeper.GetRequiredConfirmationHeight(ctx),
		Participants:       snapshot.GetParticipantAddresses(),
	}

	events.Emit(ctx, event)

	return &types.ConfirmDestTxsResponse{}, nil
}
