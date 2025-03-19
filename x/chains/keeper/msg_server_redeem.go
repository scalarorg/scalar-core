package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/events"
	chainsExported "github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/chains/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

func (s msgServer) UpdateUtxoLists(c context.Context, req *types.UpdateUtxoListsRequest) (*types.UpdateUtxoListsResponse, error) {
	//Todo: check if the RedeemSession phase is Preparing then update the utxos list
	ctx := sdk.UnwrapSDKContext(c)
	chain, ok := s.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if !chainsTypes.IsBitcoinChain(chain.Name) {
		return nil, fmt.Errorf("chain %s is not a bitcoin chain", chain.Name)
	}

	poll, err := s.initializePoll(ctx, chain, chainsExported.ZeroHash)
	if err != nil {
		return nil, err
	}
	event := &types.UpdateUtxoListsStarted{
		Chain:        chain.Name,
		PollID:       poll.PollID,
		Participants: poll.Participants,
	}

	s.Logger(ctx).Info("UpdateUtxoListsStarted", event)

	events.Emit(ctx, event)

	return &types.UpdateUtxoListsResponse{}, nil
}

func (s msgServer) ConfirmRedeemTx(c context.Context, req *chainsTypes.ConfirmRedeemTxRequest) (*chainsTypes.ConfirmRedeemTxResponse, error) {
	//Todo: request validator process confirm redeen tx is executed with >=12 confirmations on the bitcoin chain
	ctx := sdk.UnwrapSDKContext(c)

	chain, ok := s.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
		return nil, err
	}

	if !chainsTypes.IsBitcoinChain(chain.Name) {
		return nil, fmt.Errorf("chain %s is not a bitcoin chain", chain.Name)
	}

	keeper, err := s.ForChain(ctx, chain.Name)
	if err != nil {
		return nil, err
	}

	snapshot, err := s.createSnapshot(ctx, chain)
	if err != nil {
		return nil, err
	}

	pollMappings, err := s.initializePolls(ctx, chain, snapshot, []chainsExported.Hash{req.TxID})
	if err != nil {
		return nil, err
	}

	event := &types.ConfirmRedeemTxStarted{
		Chain:              chain.Name,
		PollMappings:       pollMappings,
		ConfirmationHeight: keeper.GetRequiredConfirmationHeight(ctx),
		Participants:       snapshot.GetParticipantAddresses(),
	}

	s.Logger(ctx).Info("ConfirmRedeemTxStarted", event)

	events.Emit(ctx, event)

	return &types.ConfirmRedeemTxResponse{}, nil
}
