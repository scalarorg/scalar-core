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

	clog.Red("After validateChainActivated", chain)

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
		return nil, err
	}

	pollMappings, err := s.initializePolls(ctx, chain, snapshot, req.TxIDs)
	if err != nil {
		return nil, err
	}

	event := &types.EventConfirmSourceTxsStarted{
		Chain:              chain.Name,
		PollMappings:       pollMappings,
		ConfirmationHeight: keeper.GetRequiredConfirmationHeight(ctx),
		Participants:       snapshot.GetParticipantAddresses(),
	}

	clog.Green("ConfirmStakingTxsStarted", event)

	events.Emit(ctx, event)

	return &types.ConfirmSourceTxsResponse{}, nil
}
