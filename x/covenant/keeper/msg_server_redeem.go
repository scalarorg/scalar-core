package keeper

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/bitcoin-vault/go-utils/btc"
	"github.com/scalarorg/scalar-core/utils/events"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
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
		return nil, fmt.Errorf("custodian group %s not found", req.CustodianGroupUID.Hex())
	}

	chainKeeper, err := s.chains.ForChain(ctx, req.Chain)
	if err != nil {
		return nil, err
	}

	chainParams := chainKeeper.GetParams(ctx)

	s.Keeper.Logger(ctx).Info("ConfirmRedeemTxs", "chainParams", chainParams)

	nwParams := chainParams.Metadata["params"]
	if nwParams == "" {
		return nil, fmt.Errorf("[ConfirmRedeemTxs] params is required")
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

	s.Keeper.Logger(ctx).Info(fmt.Sprintf("ConfirmRedeemTxStarted: %++v", event))

	events.Emit(ctx, event)

	return &types.ConfirmRedeemTxsResponse{}, nil
}

func (s msgServer) ConfirmSwitchedPhase(c context.Context, req *types.ConfirmSwitchedPhaseRequest) (*types.ConfirmSwitchedPhaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	chain, err := s.validateEvmChain(ctx, req.Chain)
	if err != nil {
		return nil, err
	}
	_, ok := s.Keeper.GetCustodianGroup(ctx, req.CustodianGroupUID)
	if !ok {
		s.Keeper.Logger(ctx).Error(fmt.Sprintf("custodian group %s not found", req.CustodianGroupUID.Hex()))
		return nil, fmt.Errorf("custodian group %s not found", req.CustodianGroupUID.Hex())
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

	expiresAt := ctx.BlockHeight() + chainParams.RevoteLockingPeriod

	pollID, err := s.voter.InitializePoll(
		ctx,
		vote.NewPollBuilder(types.ModuleName, chainParams.VotingThreshold, snapshot, expiresAt).
			MinVoterCount(chainParams.MinVoterCount).
			RewardPoolName(chain.Name.String()).
			GracePeriod(chainParams.VotingGracePeriod).
			ModuleMetadata(&types.BasicPollMetadata{
				Data:  req.TxID.Bytes(),
				Chain: chain.Name,
			}),
	)
	if err != nil {
		return nil, err
	}

	event := &types.ConfirmSwitchedPhaseStarted{
		PollID:             pollID,
		TxID:               req.TxID,
		Chain:              chain.Name,
		ConfirmationHeight: chainKeeper.GetRequiredConfirmationHeight(ctx),
		Participants:       snapshot.GetParticipantAddresses(),
		CustodianGroupUID:  req.CustodianGroupUID,
	}

	s.Keeper.Logger(ctx).Info(
		fmt.Sprintf("ConfirmSwitchedPhaseStarted: pollID: %s, txid: %s, chain: %s, custodian group: %s",
			event.PollID,
			hex.EncodeToString(event.TxID[:]),
			event.Chain,
			event.CustodianGroupUID.Hex(),
		),
	)

	events.Emit(ctx, event)
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
func (s msgServer) validateEvmChain(ctx sdk.Context, chain nexus.ChainName) (*nexus.Chain, error) {
	c, ok := s.nexus.GetChain(ctx, chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", chain)
	}
	if err := validateChainActivated(ctx, s.nexus, c); err != nil {
		return nil, err
	}
	if !chainsTypes.IsEvmChain(chain) {
		return nil, fmt.Errorf("chain %s is not a EVM chain", chain)
	}

	return &c, nil
}

func (s msgServer) ReserveRedeemUtxo(c context.Context, req *types.ReserveRedeemUtxoRequest) (*types.ReserveRedeemUtxoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	//Validate request
	sourceChainName := nexus.ChainName(req.SourceChain)
	sourceChain, ok := s.nexus.GetChain(ctx, sourceChainName)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.SourceChain)
	}

	destChainName := nexus.ChainName(req.DestChain)
	destChain, ok := s.nexus.GetChain(ctx, destChainName)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.DestChain)
	}

	if !chainsTypes.IsEvmChain(sourceChain.Name) {
		return nil, fmt.Errorf("source chain %s is not a EVM chain", sourceChain.Name)
	}

	if !chainsTypes.IsBitcoinChain(destChain.Name) {
		return nil, fmt.Errorf("destination chain %s is not a bitcoin chain", destChain.Name)
	}

	if err := validateChainActivated(ctx, s.nexus, sourceChain); err != nil {
		return nil, err
	}

	protocol, err := s.protocol.FindProtocolInfoByExternalSymbol(ctx, req.Symbol)
	if err != nil {
		return nil, err
	}

	// Start signing session for reserve redeem utxos
	keyID, ok := s.multisig.GetCurrentKeyID(ctx, nexus.ChainName(req.SourceChain))
	if !ok {
		return nil, fmt.Errorf("could not find key ID for '%s'", req.SourceChain)
	}
	// Get current session
	session, ok := s.Keeper.GetRedeemSession(ctx, protocol.CustodianGroupUID)
	if !ok {
		return nil, fmt.Errorf("could not find redeem session for '%s'", protocol.CustodianGroupUID.Hex())
	}
	log.Info().Any("RedeemSession", session).
		Bool("IsSwitching", session.IsSwitching).
		Any("CurrentPhase", session.CurrentPhase).
		Hex("CustodianGroupUid", protocol.CustodianGroupUID[:]).
		Msg("[x/covernant] reserveUtxos found redeem session")
	if session.CurrentPhase != exported.Preparing {
		return nil, fmt.Errorf("redeem session is not in preparing phase")
	}

	if session.IsSwitching {
		return nil, fmt.Errorf("redeem session is in switching")
	}
	// Create redeem payload for evm tx
	ck, err := s.chains.ForChain(ctx, destChainName)
	if err != nil {
		return nil, fmt.Errorf("ChainKeeper not found for chain %s", destChainName)
	}
	chainParams := ck.GetParams(ctx)
	params, dataHash, err := s.Keeper.CreateRedeemParams(ctx, req, protocol.CustodianGroupUID, &chainParams, session.Sequence)
	if err != nil {
		return nil, err
	}

	chainID, err := req.SourceChain.GetChainID()
	if err != nil {
		return nil, err
	}
	// s.Keeper.Logger(ctx).Info("ReserveRedeemUtxoStarted", "amount", req.Amount, "chain", sourceChain.Name, "sender", req.Sender)
	command, err := s.createReserveRedeemUtxoStandaloneCommand(ctx, keyID, *chainID, *dataHash, params)
	if err != nil {
		return nil, fmt.Errorf("could not create reserve utxo command")
	}

	//TODO: check signing process
	if err := s.multisig.Sign(
		ctx,
		command.GetKeyID(),
		command.GetSigHash().Bytes(),
		types.ModuleName,
		types.NewSigMetadata(types.SigCommand, sourceChain.Name, command.GetID()),
	); err != nil {
		return nil, err
	}

	s.Keeper.Logger(ctx).Info("ReserveRedeemUtxoStarted", "amount", req.Amount, "chain", sourceChain.Name, "sender", req.Sender)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeReserveRedeemUtxo,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueStart),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAmount, strconv.Itoa(int(req.Amount))),
			sdk.NewAttribute(types.AttributeKeyChain, sourceChain.Name.String()),
			sdk.NewAttribute(sdk.AttributeKeySender, req.Sender.String()),
			sdk.NewAttribute(types.AttributeKeyDataHash, hex.EncodeToString((*dataHash)[:])),
			sdk.NewAttribute(types.AttributeCommandId, hex.EncodeToString(command.GetID())),
		),
	)

	return &types.ReserveRedeemUtxoResponse{}, nil
}

func (s msgServer) createReserveRedeemUtxoStandaloneCommand(
	ctx sdk.Context,
	keyID multisig.KeyID,
	chainID sdk.Int,
	commandID types.CommandID,
	commandParam []byte,
) (*types.StandaloneCommand, error) {
	md, err := types.NewStandaloneCommandMetadata(
		ctx.BlockHeight(),
		keyID,
		chainID,
		commandID,
		chainsTypes.COMMAND_TYPE_REDEEM_TOKEN,
		commandParam,
	)

	if err != nil {
		return nil, err
	}

	s.Keeper.SetStandaloneCommandMetadata(ctx, md, reserveUtxoCommandPrefix)
	s.Keeper.SetUnsignedStandaloneCommandID(ctx, md.ID)

	setter := func(m types.StandaloneCommandMetadata) {
		s.Keeper.SetStandaloneCommandMetadata(ctx, m, reserveUtxoCommandPrefix)
	}

	cmd := types.NewStandaloneCommand(md, setter)

	return &cmd, nil
}

func (s msgServer) InitializeUtxo(c context.Context, req *types.InitializeUtxoRequest) (*types.InitializeUtxoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	s.Keeper.Logger(ctx).Info(fmt.Sprintf("[InitializeUtxo] start handling initializingUtxoRequest for chain %s", req.Chain.String()))
	chain, err := s.validateBtcChain(ctx, req.Chain)
	if err != nil {
		return nil, err
	}

	// cusGr, ok := s.Keeper.GetCustodianGroup(ctx, chains.Hash(req.CustodianGroupUID))
	// if !ok {
	// 	return nil, fmt.Errorf("custodian group %s not found", req.CustodianGroupUID)
	// }

	chainKeeper, err := s.chains.ForChain(ctx, req.Chain)
	if err != nil {
		return nil, err
	}

	chainParams := chainKeeper.GetParams(ctx)
	nwParams := chainParams.Metadata["params"]
	if nwParams == "" {
		return nil, fmt.Errorf("[InitializeUtxo] params is required")
	}
	cusGrs, _ := s.Keeper.GetAllCustodianGroups(ctx)
	for _, cusGr := range cusGrs {
		taprootAddress, err := btc.ScriptPubKeyToAddress(cusGr.BitcoinPubkey, nwParams)
		if err != nil {
			return nil, err
		}
		s.Keeper.Logger(ctx).Info(fmt.Sprintf("[InitializeUtxo] custodian group %s with taproot address %s", cusGr.UID.Hex(), taprootAddress.String()))
		threshold := chainParams.VotingThreshold

		snapshot, err := s.createSnapshot(ctx, *chain, threshold)
		if err != nil {
			return nil, err
		}

		expiresAt := ctx.BlockHeight() + chainParams.RevoteLockingPeriod

		pollID, err := s.voter.InitializePoll(
			ctx,
			vote.NewPollBuilder(types.ModuleName, chainParams.VotingThreshold, snapshot, expiresAt).
				MinVoterCount(chainParams.MinVoterCount).
				RewardPoolName(req.Chain.String()).
				GracePeriod(chainParams.VotingGracePeriod).
				ModuleMetadata(&types.BasicPollMetadata{
					Chain: req.Chain,
				}),
		)
		if err != nil {
			return nil, err
		}

		event := &types.IntializeUtxoSnapshotStarted{
			PollID:             pollID,
			Chain:              req.Chain,
			ConfirmationHeight: chainKeeper.GetRequiredConfirmationHeight(ctx),
			Participants:       snapshot.GetParticipantAddresses(),
			CustodianGroupUID:  cusGr.UID,
			Address:            taprootAddress.String(),
			BlockCheckpoint:    req.BlockCheckpoint,
		}

		events.Emit(ctx, event)
	}
	return nil, nil
}
