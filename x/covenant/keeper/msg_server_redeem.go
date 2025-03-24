package keeper

import (
	"bytes"
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/events"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	vote "github.com/scalarorg/scalar-core/x/vote/exported"
)

func (s msgServer) ConfirmRedeemTxs(c context.Context, req *types.ConfirmRedeemTxsRequest) (*types.ConfirmRedeemTxsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	chain, err := s.validateBtcChain(ctx, req.Chain)
	if err != nil {
		return nil, err
	}

	cusGr, ok := s.Keeper.GetCustodianGroup(ctx, req.CustodianGroupUID)
	if !ok {
		return nil, fmt.Errorf("custodian group %s not found", req.CustodianGroupUID)
	}

	chainKeeper, err := s.chains.ForChain(ctx, req.Chain)
	if err != nil {
		return nil, err
	}

	chainParams := chainKeeper.GetParams(ctx)

	nwParams := chainParams.Metadata["params"]
	if nwParams == "" {
		return nil, fmt.Errorf("params is required")
	}

	threshold := chainParams.VotingThreshold

	snapshot, err := s.createSnapshot(ctx, *chain, threshold)
	if err != nil {
		return nil, err
	}

	expiresAt := ctx.BlockHeight() + chainParams.RevoteLockingPeriod

	var data bytes.Buffer
	for _, txID := range req.TxIDs {
		data.Write(txID.Bytes())
	}

	pollID, err := s.voter.InitializePoll(
		ctx,
		vote.NewPollBuilder(types.ModuleName, chainParams.VotingThreshold, snapshot, expiresAt).
			MinVoterCount(chainParams.MinVoterCount).
			RewardPoolName(chain.Name.String()).
			GracePeriod(chainParams.VotingGracePeriod).
			ModuleMetadata(&types.BasicPollMetadata{
				Data:  data.Bytes(),
				Chain: chain.Name,
			}),
	)
	if err != nil {
		return nil, err
	}

	event := &types.ConfirmRedeemTxStarted{
		PollID:             pollID,
		TxIDs:              req.TxIDs,
		Chain:              chain.Name,
		ConfirmationHeight: chainKeeper.GetRequiredConfirmationHeight(ctx),
		Participants:       snapshot.GetParticipantAddresses(),
		CustodianGroupUID:  req.CustodianGroupUID,
		ScriptPubkey:       cusGr.BitcoinPubkey,
		NetworkParams:      nwParams,
	}

	s.Logger(ctx).Info("ConfirmRedeemTxStarted", event)

	events.Emit(ctx, event)

	return &types.ConfirmRedeemTxsResponse{}, nil
}

func (s msgServer) ConfirmSwitchedPhase(ctx context.Context, req *types.ConfirmSwitchedPhaseRequest) (*types.ConfirmSwitchedPhaseResponse, error) {
	return &types.ConfirmSwitchedPhaseResponse{}, nil
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

	reservedTx, err := s.reserveUtxos(ctx, protocol.CustodianGroupUID.Bytes(), req.ReqId, req.Amount)
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

