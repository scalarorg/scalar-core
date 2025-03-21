package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/events"
	chainsExported "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	snapshot "github.com/scalarorg/scalar-core/x/snapshot/exported"
	vote "github.com/scalarorg/scalar-core/x/vote/exported"
)

func (s msgServer) ReserveRedeemUtxo(c context.Context, req *types.ReserveRedeemUtxoRequest) (*types.ReserveRedeemUtxoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	//Validate request
	chain, ok := s.nexus.GetChain(ctx, nexus.ChainName(req.Chain))
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if !chainsTypes.IsEvmChain(chain.Name) {
		return nil, fmt.Errorf("chain %s is not a EVM chain", chain.Name)
	}

	if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
		return nil, err
	}

	protocol, err := s.protocol.FindProtocolInfoByExternalSymbol(ctx, req.Symbol)
	if err != nil {
		return nil, err
	}

	reservedTx, err := s.reserveUtxos(ctx, protocol.CustodiansGroupUID, req.ReqId, req.Amount)
	if err != nil {
		return nil, err
	}

	// Create redeem payload for evm tx
	payload, err := s.createRedeemPayload(ctx, req.Chain.String(), req.Address, req.Symbol, req.Amount, reservedTx)
	if err != nil {
		return nil, err
	}
	// Start signing session for reserve redeem utxos
	keyID, ok := s.multisig.GetCurrentKeyID(ctx, nexus.ChainName(req.Chain))
	if !ok {
		return nil, fmt.Errorf("could not find key ID for '%s'", req.Chain)
	}
	//Todo: check signing process
	if err := s.multisig.Sign(
		ctx,
		keyID,
		payload,
		types.ModuleName,
		chainsTypes.NewSigMetadata(chainsTypes.SigTx, chain.Name, []byte(req.ReqId)),
	); err != nil {
		return nil, err
	}

	logger := s.Logger(ctx)
	logger.Info("ReserveRedeemUtxoStarted", "reqId", req.ReqId, "amount", req.Amount, "chain", chain.Name, "sender", req.Sender)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			chainsTypes.EventTypeSign,
			sdk.NewAttribute(sdk.AttributeKeyAction, chainsTypes.AttributeValueStart),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAmount, string(req.Amount)),
			sdk.NewAttribute("chain", chain.Name.String()),
			sdk.NewAttribute(sdk.AttributeKeySender, req.Sender.String()),
		),
	)

	return &types.ReserveRedeemUtxoResponse{}, nil
}

func (s msgServer) ConfirmRedeemTxs(c context.Context, req *types.ConfirmRedeemTxsRequest) (*types.ConfirmRedeemTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	chain, err := s.validateBtcChain(ctx, req.Chain)
	if err != nil {
		return nil, err
	}

	chainKeeper, err := s.chains.ForChain(ctx, req.Chain)
	if err != nil {
		return nil, err
	}

	chainParams := chainKeeper.GetParams(ctx)
	threshold := chainParams.VotingThreshold

	snapshot, err := s.createSnapshot(ctx, *chain, threshold)
	if err != nil {
		return nil, err
	}

	pollMappings, err := s.initializePolls(ctx, *chain, snapshot, req.TxIDs, chainParams)
	if err != nil {
		return nil, err
	}

	event := &types.ConfirmRedeemTxStarted{
		Chain:              chain.Name,
		PollMappings:       pollMappings,
		ConfirmationHeight: chainKeeper.GetRequiredConfirmationHeight(ctx),
		Participants:       snapshot.GetParticipantAddresses(),
	}

	s.Logger(ctx).Info("ConfirmRedeemTxStarted", event)

	events.Emit(ctx, event)

	return &types.ConfirmRedeemTxsResponse{}, nil
}

func (s msgServer) ConfirmSwitchedPhase(ctx context.Context, req *types.ConfirmSwitchedPhaseRequest) (*types.ConfirmSwitchedPhaseResponse, error) {
	return &types.ConfirmSwitchedPhaseResponse{}, nil
}

func (s msgServer) initializePolls(ctx sdk.Context, chain nexus.Chain, snapshot snapshot.Snapshot, txIDs []chainsExported.Hash, params chainsTypes.Params) ([]chainsTypes.PollMapping, error) {

	expiresAt := ctx.BlockHeight() + params.RevoteLockingPeriod

	pollMappings := make([]chainsTypes.PollMapping, len(txIDs))
	for i, txID := range txIDs {
		pollID, err := s.voter.InitializePoll(
			ctx,
			vote.NewPollBuilder(types.ModuleName, params.VotingThreshold, snapshot, expiresAt).
				MinVoterCount(params.MinVoterCount).
				RewardPoolName(chain.Name.String()).
				GracePeriod(params.VotingGracePeriod).
				ModuleMetadata(&chainsTypes.PollMetadata{
					Chain: chain.Name,
					TxID:  txID,
				}),
		)
		if err != nil {
			return nil, err
		}

		pollMappings[i] = chainsTypes.PollMapping{
			TxID:   txID,
			PollID: pollID,
		}
	}

	return pollMappings, nil
}

func (s msgServer) validateBtcChain(ctx sdk.Context, chain nexus.ChainName) (*nexus.Chain, error) {
	c, ok := s.nexus.GetChain(ctx, chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", chain)
	}

	if err := validateChainActivated(ctx, s.nexus, c); err != nil {
		return nil, err
	}

	if !chainsTypes.IsBitcoinChain(chain) {
		return nil, fmt.Errorf("chain %s is not a bitcoin chain", chain)
	}

	return &c, nil
}
