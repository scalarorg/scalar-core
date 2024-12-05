package keeper

import (
	"context"
	"fmt"

	"github.com/axelarnetwork/axelar-core/utils/events"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

var _ types.MsgServiceServer = msgServer{}

type msgServer struct {
	types.BaseKeeper
	nexus       types.Nexus
	snapshotter types.Snapshotter
	slashing    types.SlashingKeeper
	voter       types.Voter
}

type MsgServerConstructArgs struct {
	types.BaseKeeper
	Nexus       types.Nexus
	Snapshotter types.Snapshotter
	Slashing    types.SlashingKeeper
}

func NewMsgServerImpl(args MsgServerConstructArgs) types.MsgServiceServer {
	return msgServer{
		BaseKeeper:  args.BaseKeeper,
		nexus:       args.Nexus,
		snapshotter: args.Snapshotter,
		slashing:    args.Slashing,
	}
}

func (s msgServer) ConfirmGatewayTxs(c context.Context, req *types.ConfirmGatewayTxsRequest) (*types.ConfirmGatewayTxsResponse, error) {
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

	snapshot, err := s.createSnapshot(ctx, chain)
	if err != nil {
		return nil, err
	}

	pollMappings, err := s.initializePolls(ctx, chain, snapshot, req.TxIDs)
	if err != nil {
		return nil, err
	}

	events.Emit(ctx, &types.ConfirmGatewayTxsStarted{
		PollMappings:       pollMappings,
		Chain:              chain.Name,
		ConfirmationHeight: keeper.GetRequiredConfirmationHeight(ctx),
		Participants:       snapshot.GetParticipantAddresses(),
	})

	return &types.ConfirmGatewayTxsResponse{}, nil
}
