package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

func (s msgServer) ConfirmStakingTxs(c context.Context, req *types.ConfirmStakingTxsRequest) (*types.ConfirmStakingTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	_ = ctx
	// fmt.Println("btc chain", req.Chain)
	clog.Red("req.Chain", req.Chain)

	// chain, ok := s.nexus.GetChain(ctx, req.Chain)
	// if !ok {
	// 	return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	// }

	// if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
	// 	return nil, err
	// }

	// keeper, err := s.ForChain(ctx, chain.Name)
	// if err != nil {
	// 	return nil, err
	// }

	// snapshot, err := s.createSnapshot(ctx, chain)
	// if err != nil {
	// 	return nil, err
	// }

	// pollMappings, err := s.initializePolls(ctx, chain, snapshot, req.TxIDs)
	// if err != nil {
	// 	return nil, err
	// }

	// events.Emit(ctx, &types.ConfirmStakingTxsStarted{
	// 	PollMappings:       pollMappings,
	// 	Chain:              chain.Name,
	// 	ConfirmationHeight: keeper.GetRequiredConfirmationHeight(ctx),
	// 	Participants:       snapshot.GetParticipantAddresses(),
	// })

	return &types.ConfirmStakingTxsResponse{}, nil
}
