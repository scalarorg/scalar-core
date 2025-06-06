package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/x/chains/types"
)

func (s msgServer) ConfirmSourceTxs(c context.Context, req *types.ConfirmSourceTxsRequest) (*types.ConfirmSourceTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	chain, ok := s.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
		return nil, err
	}

	clog.Green("After validateChainActivated", chain)

	keeper, err := s.ForChain(ctx, chain.Name)
	if err != nil {
		return nil, err
	}

	// TODO: get gateway address if evm chain, gate pool address(p2tr) if btc chain
	// TODO: include these address in the event
	// gatewayAddress, ok := keeper.GetGatewayAddress(ctx)
	// if !ok {
	// 	return nil, fmt.Errorf("gateway address not found")
	// }

	snapshot, err := s.createSnapshot(ctx, chain)
	if err != nil {
		s.Logger(ctx).Error("createSnapshot failed", "error", err)
		return nil, err
	}

	pollMappings, err := s.initializePolls(ctx, chain, snapshot, req.TxIDs)
	if err != nil {
		return nil, err
	}

	events.Emit(ctx, &types.EventConfirmSourceTxsStarted{
		Chain:              chain.Name,
		PollMappings:       pollMappings,
		ConfirmationHeight: keeper.GetRequiredConfirmationHeight(ctx),
		Participants:       snapshot.GetParticipantAddresses(),
	})
	return &types.ConfirmSourceTxsResponse{}, nil
}
