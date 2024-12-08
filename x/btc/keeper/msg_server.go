package keeper

import (
	"fmt"

	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	"github.com/axelarnetwork/utils/funcs"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/btc/types"

	snapshot "github.com/axelarnetwork/axelar-core/x/snapshot/exported"
	vote "github.com/axelarnetwork/axelar-core/x/vote/exported"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (args MsgServerConstructArgs) Validate() error {
	if args.BaseKeeper == nil {
		return fmt.Errorf("BaseKeeper is nil")
	}

	if args.Slashing == nil {
		return fmt.Errorf("Slashing keeper is nil")
	}

	if args.Snapshotter == nil {
		return fmt.Errorf("Snapshotter is nil")
	}

	if args.Nexus == nil {
		return fmt.Errorf("Nexus is nil")
	}

	return nil
}

func NewMsgServerImpl(args MsgServerConstructArgs) types.MsgServiceServer {
	return msgServer{
		BaseKeeper:  args.BaseKeeper,
		nexus:       args.Nexus,
		snapshotter: args.Snapshotter,
		slashing:    args.Slashing,
	}
}

func validateChainActivated(ctx sdk.Context, nexus types.Nexus, chain nexus.Chain) error {
	if !nexus.IsChainActivated(ctx, chain) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("chain %s is not activated yet", chain.Name))
	}

	return nil
}

func (s msgServer) createSnapshot(ctx sdk.Context, chain nexus.Chain) (snapshot.Snapshot, error) {
	keeper, err := s.ForChain(ctx, chain.Name)
	if err != nil {
		return snapshot.Snapshot{}, err
	}
	params := keeper.GetParams(ctx)

	return s.snapshotter.CreateSnapshot(
		ctx,
		s.nexus.GetChainMaintainers(ctx, chain),
		excludeJailedOrTombstoned(ctx, s.slashing, s.snapshotter),
		snapshot.QuadraticWeightFunc,
		params.VotingThreshold,
	)
}

func excludeJailedOrTombstoned(ctx sdk.Context, slashing types.SlashingKeeper, snapshotter types.Snapshotter) func(v snapshot.ValidatorI) bool {
	isTombstoned := func(v snapshot.ValidatorI) bool {
		consAdd, err := v.GetConsAddr()
		if err != nil {
			return true
		}

		return slashing.IsTombstoned(ctx, consAdd)
	}

	isProxyActive := func(v snapshot.ValidatorI) bool {
		_, isActive := snapshotter.GetProxy(ctx, v.GetOperator())

		return isActive
	}

	return funcs.And(
		snapshot.ValidatorI.IsBonded,
		funcs.Not(snapshot.ValidatorI.IsJailed),
		funcs.Not(isTombstoned),
		isProxyActive,
	)
}

func (s msgServer) initializePolls(ctx sdk.Context, chain nexus.Chain, snapshot snapshot.Snapshot, txIDs []types.Hash) ([]types.PollMapping, error) {
	keeper, err := s.ForChain(ctx, chain.Name)
	if err != nil {
		return nil, err
	}

	params := keeper.GetParams(ctx)
	expiresAt := ctx.BlockHeight() + params.RevoteLockingPeriod

	pollMappings := make([]types.PollMapping, len(txIDs))
	for i, txID := range txIDs {
		pollID, err := s.voter.InitializePoll(
			ctx,
			vote.NewPollBuilder(types.ModuleName, params.VotingThreshold, snapshot, expiresAt).
				MinVoterCount(params.MinVoterCount).
				RewardPoolName(chain.Name.String()).
				GracePeriod(keeper.GetParams(ctx).VotingGracePeriod).
				ModuleMetadata(&types.PollMetadata{
					Chain: chain.Name,
					TxID:  txID,
				}),
		)
		if err != nil {
			return nil, err
		}

		pollMappings[i] = types.PollMapping{
			TxID:   txID,
			PollID: pollID,
		}
	}

	return pollMappings, nil
}
