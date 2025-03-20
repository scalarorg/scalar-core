package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
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

	reservedTx, err := s.reserveUtxos(ctx, req.Symbol, req.ReqId, req.Amount)
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

func (s msgServer) ConfirmSwitchedPhase(ctx context.Context, req *types.ConfirmSwitchedPhaseRequest) (*types.ConfirmSwitchedPhaseResponse, error) {
	return &types.ConfirmSwitchedPhaseResponse{}, nil
}

func (s msgServer) UpdateUtxoForRedeemSession(c context.Context, req *types.UpdateUtxoForRedeemSessionRequest) (*types.UpdateUtxoForRedeemSessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	clog.Magentaf("UpdateUtxoForRedeemSession, req: %v\n", req)
	clog.Greenf("UpdateUtxoForRedeemSession, req.ListOfUtxos: %v\n", req.ListOfUtxos)

	err := req.ValidateBasic()
	if err != nil {
		clog.Redf("UpdateUtxoForRedeemSession, ValidateBasic error: %v\n", err)
		return nil, err
	}

	clog.Bluef("UpdateUtxoForRedeemSession, req.Symbol: %s\n", req.Symbol)

	protocol, err := s.protocol.FindProtocolInfoByExternalSymbol(ctx, req.Symbol)
	if err != nil {
		clog.Redf("UpdateUtxoForRedeemSession, FindProtocolInfoByExternalSymbol error: %v\n", err)
		return nil, err
	}

	clog.Magentaf("UpdateUtxoForRedeemSession, protocol: %v\n", protocol)

	if !protocol.IsActivated() {
		clog.Redf("UpdateUtxoForRedeemSession, protocol %s is not activated\n", req.Symbol)
		return nil, fmt.Errorf("protocol %s is not activated", req.Symbol)
	}

	// TODO: Split to the confirmation event then handle in abci,
	// This way is just a temporary solution

	session, ok := s.GetRedeemSession(ctx, req.Symbol)
	if !ok {
		clog.Bluef("Check if session exist\n")
		session = types.DefaultRedeemSession()
		session.LastestBtcBlockHeight = req.BlockHeight
		session.Utxos = req.ListOfUtxos
		session.Symbol = req.Symbol
		s.Keeper.AddRedeemSession(ctx, req.Symbol, session)
		return &types.UpdateUtxoForRedeemSessionResponse{
			Session: session,
		}, nil
	}

	clog.Magentaf("UpdateUtxoForRedeemSession, session: %v\n", session)

	if session.LastestBtcBlockHeight >= req.BlockHeight {
		return nil, fmt.Errorf("utxos already exist")
	}

	session.LastestBtcBlockHeight = req.BlockHeight
	session.Utxos = req.ListOfUtxos
	s.Keeper.AddRedeemSession(ctx, req.Symbol, session)

	return &types.UpdateUtxoForRedeemSessionResponse{
		Session: session,
	}, nil
}

func (s msgServer) UpdateSequenceAndPhaseForRedeemSession(c context.Context, req *types.UpdateSequenceAndPhaseForRedeemSessionRequest) (*types.UpdateSequenceAndPhaseForRedeemSessionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := req.ValidateBasic()
	if err != nil {
		clog.Redf("UpdateUtxoForRedeemSession, ValidateBasic error: %v\n", err)
		return nil, err
	}

	clog.Bluef("UpdateUtxoForRedeemSession, req.Symbol: %s\n", req.Symbol)

	protocol, err := s.protocol.FindProtocolInfoByExternalSymbol(ctx, req.Symbol)
	if err != nil {
		clog.Redf("UpdateUtxoForRedeemSession, FindProtocolInfoByExternalSymbol error: %v\n", err)
		return nil, err
	}

	clog.Magentaf("UpdateUtxoForRedeemSession, protocol: %v\n", protocol)

	if !protocol.IsActivated() {
		clog.Redf("UpdateUtxoForRedeemSession, protocol %s is not activated\n", req.Symbol)
		return nil, fmt.Errorf("protocol %s is not activated", req.Symbol)
	}

	// Check if session exist
	// TODO: Split to the confirmation event then handle in abci,
	// This way is just a temporary solution

	session, ok := s.GetRedeemSession(ctx, req.Symbol)
	if !ok {
		clog.Bluef("Check if session exist\n")
		session = types.DefaultRedeemSession()
		session.CurrentPhase = req.Phase
		session.Sequence = req.Sequence
		session.Symbol = req.Symbol
		s.Keeper.AddRedeemSession(ctx, req.Symbol, session)
		return &types.UpdateSequenceAndPhaseForRedeemSessionResponse{
			Session: session,
		}, nil
	}

	clog.Magentaf("UpdateUtxoForRedeemSession, session: %v\n", session)

	if session.Sequence > req.Sequence {
		return nil, fmt.Errorf("sequence is invalid")
	}

	if session.Sequence == req.Sequence && session.CurrentPhase >= req.Phase {
		return nil, fmt.Errorf("phase is invalid")
	}

	session.Sequence = req.Sequence
	session.CurrentPhase = req.Phase
	s.Keeper.AddRedeemSession(ctx, req.Symbol, session)

	return &types.UpdateSequenceAndPhaseForRedeemSessionResponse{
		Session: session,
	}, nil
}

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

func (s msgServer) ConfirmRedeemTx(c context.Context, req *types.ConfirmRedeemTxRequest) (*types.ConfirmRedeemTxResponse, error) {
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
