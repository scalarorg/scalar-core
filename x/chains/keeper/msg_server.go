package keeper

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	snapshot "github.com/scalarorg/scalar-core/x/snapshot/exported"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"
	vote "github.com/scalarorg/scalar-core/x/vote/exported"
)

var _ types.MsgServiceServer = msgServer{}

type msgServer struct {
	types.BaseKeeper
	nexus       types.Nexus
	snapshotter types.Snapshotter
	slashing    types.SlashingKeeper
	voter       types.Voter
	staking     types.StakingKeeper
	multisig    types.MultisigKeeper
	covenant    types.CovenantKeeper
	protocol    types.ProtocolKeeper
}

type MsgServerConstructArgs struct {
	types.BaseKeeper
	Nexus       types.Nexus
	Voter       types.Voter
	Snapshotter types.Snapshotter
	Staking     types.StakingKeeper
	Slashing    types.SlashingKeeper
	Multisig    types.MultisigKeeper
	Covenant    types.CovenantKeeper
	Protocol    types.ProtocolKeeper
}

func (args MsgServerConstructArgs) Validate() error {
	if args.BaseKeeper == nil {
		return fmt.Errorf("BaseKeeper is nil")
	}

	if args.Nexus == nil {
		return fmt.Errorf("nexus is nil")
	}

	if args.Voter == nil {
		return fmt.Errorf("voter is nil")
	}

	if args.Snapshotter == nil {
		return fmt.Errorf("snapshotter is nil")
	}

	if args.Staking == nil {
		return fmt.Errorf("staking keeper is nil")
	}

	if args.Slashing == nil {
		return fmt.Errorf("slashing keeper is nil")
	}

	if args.Multisig == nil {
		return fmt.Errorf("multisig keeper is nil")
	}

	if args.Covenant == nil {
		return fmt.Errorf("covenant keeper is nil")
	}
	if args.Protocol == nil {
		return fmt.Errorf("protocol keeper is nil")
	}
	return nil
}

func NewMsgServerImpl(args MsgServerConstructArgs) types.MsgServiceServer {
	return msgServer{
		BaseKeeper:  args.BaseKeeper,
		nexus:       args.Nexus,
		voter:       args.Voter,
		snapshotter: args.Snapshotter,
		slashing:    args.Slashing,
		staking:     args.Staking,
		multisig:    args.Multisig,
		covenant:    args.Covenant,
		protocol:    args.Protocol,
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

	candidates := s.nexus.GetChainMaintainers(ctx, chain)
	return s.snapshotter.CreateSnapshot(
		ctx,
		candidates,
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

func (s msgServer) SetGateway(c context.Context, req *types.SetGatewayRequest) (*types.SetGatewayResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	chain, ok := s.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
		return nil, err
	}

	if _, ok := s.multisig.GetCurrentKeyID(ctx, chain.Name); !ok {
		return nil, fmt.Errorf("current key not set for chain %s", chain.Name)
	}

	keeper, err := s.ForChain(ctx, chain.Name)
	if err != nil {
		return nil, err
	}
	if _, ok := keeper.GetGatewayAddress(ctx); ok {
		return nil, fmt.Errorf("%s gateway already set", req.Chain)
	}

	keeper.SetGateway(ctx, req.Address)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EventTypeGateway,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueConfirm),
			sdk.NewAttribute(types.AttributeKeyChain, chain.Name.String()),
			sdk.NewAttribute(types.AttributeKeyAddress, req.Address.Hex()),
		),
	)

	return &types.SetGatewayResponse{}, nil
}

func (s msgServer) Link(c context.Context, req *types.LinkRequest) (*types.LinkResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	senderChain, ok := s.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if err := validateChainActivated(ctx, s.nexus, senderChain); err != nil {
		return nil, err
	}

	keeper, err := s.ForChain(ctx, senderChain.Name)
	if err != nil {
		return nil, err
	}
	gatewayAddr, ok := keeper.GetGatewayAddress(ctx)
	if !ok {
		return nil, fmt.Errorf("scalar gateway address not set")
	}

	recipientChain, ok := s.nexus.GetChain(ctx, req.RecipientChain)
	if !ok {
		return nil, fmt.Errorf("unknown recipient chain")
	}

	token := keeper.GetERC20TokenByAsset(ctx, req.Asset)
	found := s.nexus.IsAssetRegistered(ctx, recipientChain, req.Asset)
	if !found || !token.Is(types.Confirmed) {
		return nil, fmt.Errorf("asset '%s' not registered for chain '%s'", req.Asset, recipientChain.Name)
	}

	salt := keeper.GenerateSalt(ctx, req.RecipientAddr)
	burnerAddress, err := keeper.GetBurnerAddress(ctx, token, salt, gatewayAddr)
	if err != nil {
		return nil, err
	}

	symbol := token.GetDetails().Symbol
	recipient := nexus.CrossChainAddress{Chain: recipientChain, Address: req.RecipientAddr}

	err = s.nexus.LinkAddresses(ctx,
		nexus.CrossChainAddress{Chain: senderChain, Address: burnerAddress.Hex()},
		recipient)
	if err != nil {
		return nil, fmt.Errorf("could not link addresses: %s", err.Error())
	}

	burnerInfo := types.BurnerInfo{
		BurnerAddress:    burnerAddress,
		TokenAddress:     token.GetAddress(),
		DestinationChain: req.RecipientChain,
		Symbol:           symbol,
		Asset:            req.Asset,
		Salt:             salt,
	}
	keeper.SetBurnerInfo(ctx, burnerInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLink,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeySourceChain, senderChain.Name.String()),
			sdk.NewAttribute(types.AttributeKeyDepositAddress, burnerAddress.Hex()),
			sdk.NewAttribute(types.AttributeKeyDestinationAddress, req.RecipientAddr),
			sdk.NewAttribute(types.AttributeKeyDestinationChain, recipientChain.Name.String()),
			sdk.NewAttribute(types.AttributeKeyTokenAddress, token.GetAddress().Hex()),
			sdk.NewAttribute(types.AttributeKeyAsset, req.Asset),
		),
	)

	s.Logger(ctx).Debug(fmt.Sprintf("successfully linked deposit %s on chain %s to recipient %s on chain %s for asset %s", burnerAddress.Hex(), req.Chain, req.RecipientAddr, req.RecipientChain, req.Asset),
		types.AttributeKeySourceChain, senderChain.Name,
		types.AttributeKeyDepositAddress, burnerAddress.Hex(),
		types.AttributeKeyDestinationChain, recipientChain.Name,
		types.AttributeKeyDestinationAddress, req.RecipientAddr,
		types.AttributeKeyAsset, req.Asset,
	)

	return &types.LinkResponse{DepositAddr: burnerAddress.Hex()}, nil
}

// ConfirmToken handles token deployment confirmation
func (s msgServer) ConfirmToken(c context.Context, req *types.ConfirmTokenRequest) (*types.ConfirmTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	chain, ok := s.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
		return nil, err
	}

	_, ok = s.nexus.GetChain(ctx, req.Asset.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Asset.Chain)
	}

	keeper, err := s.ForChain(ctx, chain.Name)
	if err != nil {
		return nil, err
	}
	token := keeper.GetERC20TokenByAsset(ctx, req.Asset.Symbol)

	if err := token.RecordDeployment(req.TxID); err != nil {
		return nil, err
	}

	pollParticipants, err := s.initializePoll(ctx, chain, req.TxID)
	if err != nil {
		return nil, err
	}

	events.Emit(ctx, &types.ConfirmTokenStarted{
		TxID:               req.TxID,
		Chain:              chain.Name,
		GatewayAddress:     funcs.MustOk(keeper.GetGatewayAddress(ctx)),
		TokenAddress:       token.GetAddress(),
		TokenDetails:       token.GetDetails(),
		ConfirmationHeight: keeper.GetRequiredConfirmationHeight(ctx),
		PollParticipants:   pollParticipants,
	})

	return &types.ConfirmTokenResponse{}, nil
}

// ConfirmDeposit handles deposit confirmations
func (s msgServer) ConfirmDeposit(c context.Context, req *types.ConfirmDepositRequest) (*types.ConfirmDepositResponse, error) {
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
	gatewayAddr, ok := keeper.GetGatewayAddress(ctx)
	if !ok {
		return nil, fmt.Errorf("gateway address not set for chain %s", chain.Name)
	}

	burnerInfo := keeper.GetBurnerInfo(ctx, req.BurnerAddress)
	if burnerInfo == nil {
		return nil, fmt.Errorf("no burner info found for address %s", req.BurnerAddress.Hex())
	}

	token := keeper.GetERC20TokenByAsset(ctx, burnerInfo.Asset)
	if !token.Is(types.Confirmed) {
		return nil, fmt.Errorf("token %s is not confirmed on %s", token.GetAsset(), chain.Name)
	}

	burnerAddress, err := keeper.GetBurnerAddress(ctx, token, burnerInfo.Salt, gatewayAddr)
	if err != nil {
		return nil, err
	}

	if burnerAddress != req.BurnerAddress {
		return nil, fmt.Errorf("provided burner address %s doesn't match expected address %s", req.BurnerAddress.Hex(), burnerAddress.Hex())
	}

	pollParticipants, err := s.initializePoll(ctx, chain, req.TxID)
	if err != nil {
		return nil, err
	}

	height := keeper.GetRequiredConfirmationHeight(ctx)
	events.Emit(ctx, &types.ConfirmDepositStarted{
		TxID:               req.TxID,
		Chain:              chain.Name,
		DepositAddress:     req.BurnerAddress,
		TokenAddress:       burnerInfo.TokenAddress,
		ConfirmationHeight: height,
		PollParticipants:   pollParticipants,
		Asset:              burnerInfo.Asset,
	})

	return &types.ConfirmDepositResponse{}, nil
}

// ConfirmTransferKey handles transfer operatorship confirmation
func (s msgServer) ConfirmTransferKey(c context.Context, req *types.ConfirmTransferKeyRequest) (*types.ConfirmTransferKeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	chain, ok := s.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
		return nil, err
	}

	if _, ok := s.multisig.GetNextKeyID(ctx, chain.Name); !ok {
		return nil, fmt.Errorf("next key for chain %s not set yet", chain.Name)
	}

	keeper, err := s.ForChain(ctx, chain.Name)
	if err != nil {
		return nil, err
	}

	gatewayAddr, ok := keeper.GetGatewayAddress(ctx)
	if !ok {
		return nil, fmt.Errorf("scalar gateway address not set")
	}

	pollParticipants, err := s.initializePoll(ctx, chain, req.TxID)
	if err != nil {
		return nil, err
	}

	params := keeper.GetParams(ctx)
	events.Emit(ctx, types.NewConfirmKeyTransferStarted(chain.Name, req.TxID, gatewayAddr, params.ConfirmationHeight, pollParticipants))

	return &types.ConfirmTransferKeyResponse{}, nil
}

func (s msgServer) CreateDeployToken(c context.Context, req *types.CreateDeployTokenRequest) (*types.CreateDeployTokenResponse, error) {
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

	protocol, err := s.protocol.FindProtocolInfoByExternalSymbol(ctx, req.TokenSymbol)
	if err != nil {
		return nil, fmt.Errorf("failed to find protocol info by symbol %s: %w", req.TokenSymbol, err)
	}

	switch req.Address.IsZeroAddress() {
	case true:
		originChain, found := s.nexus.GetChain(ctx, protocol.OriginChain)
		if !found {
			return nil, fmt.Errorf("%s is not a registered chain", protocol.OriginChain)
		}

		if !s.nexus.IsAssetRegistered(ctx, originChain, req.TokenSymbol) {
			return nil, fmt.Errorf("asset %s is not registered on the origin chain %s", req.TokenSymbol, originChain.Name)
		}
	case false:
		for _, c := range s.nexus.GetChains(ctx) {
			if s.nexus.IsAssetRegistered(ctx, c, req.TokenSymbol) {
				return nil, fmt.Errorf("asset %s already registered on chain %s", req.TokenSymbol, c.Name)
			}
		}

		for _, token := range keeper.GetTokens(ctx) {
			if bytes.Equal(token.GetAddress().Bytes(), req.Address.Bytes()) {
				return nil, fmt.Errorf("token %s already created for chain %s", token.GetAddress().Hex(), chain.Name)
			}
		}
	}

	keyID, ok := s.multisig.GetCurrentKeyID(ctx, chain.Name)
	if !ok {
		return nil, fmt.Errorf("current key not set for chain %s", chain.Name)
	}

	token, err := keeper.CreateERC20Token(ctx, req.TokenSymbol, *protocol.TokenDetails, req.Address)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to initialize token %s(%s) for chain %s", protocol.TokenDetails.TokenName, protocol.TokenDetails.Symbol, chain.Name)
	}

	cmd, err := token.CreateDeployCommand(keyID, protocol.TokenDailyMintLimit)
	if err != nil {
		return nil, err
	}

	if err := keeper.EnqueueCommand(ctx, cmd); err != nil {
		return nil, err
	}

	var mintLimit = protocol.TokenDailyMintLimit

	if mintLimit.IsZero() {
		mintLimit = utils.MaxUint
	}
	if err = s.nexus.RegisterAsset(ctx, chain, nexus.NewAsset(req.TokenSymbol, false), mintLimit, types.DefaultRateLimitWindow); err != nil {
		return nil, err
	}

	return &types.CreateDeployTokenResponse{}, nil
}

func (s msgServer) CreateBurnTokens(c context.Context, req *types.CreateBurnTokensRequest) (*types.CreateBurnTokensResponse, error) {
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
	transferLimit := keeper.GetParams(ctx).TransferLimit
	pageRequest := &query.PageRequest{
		Key:        nil,
		Offset:     0,
		Limit:      transferLimit,
		CountTotal: false,
		Reverse:    false,
	}
	deposits, _, err := keeper.GetConfirmedDepositsPaginated(ctx, pageRequest)
	if err != nil {
		return nil, err
	}
	if len(deposits) == 0 {
		return &types.CreateBurnTokensResponse{}, nil
	}

	chainID, ok := keeper.GetChainID(ctx)
	if !ok {
		return nil, fmt.Errorf("could not find chain ID for '%s'", chain.Name)
	}

	keyID, ok := s.multisig.GetCurrentKeyID(ctx, chain.Name)
	if !ok {
		return nil, fmt.Errorf("current key not set for chain %s", chain.Name)
	}

	seen := map[string]bool{}
	for _, deposit := range deposits {
		keeper.DeleteDeposit(ctx, deposit)
		keeper.SetDeposit(ctx, deposit, types.DepositStatus_Burned)

		burnerAddressHex := deposit.BurnerAddress.Hex()

		if seen[burnerAddressHex] {
			continue
		}

		burnerInfo := keeper.GetBurnerInfo(ctx, deposit.BurnerAddress)
		if burnerInfo == nil {
			return nil, fmt.Errorf("no burner info found for address %s", burnerAddressHex)
		}

		token := keeper.GetERC20TokenByAsset(ctx, burnerInfo.Asset)
		if !token.Is(types.Confirmed) {
			return nil, fmt.Errorf("token %s is not confirmed on %s", token.GetAsset(), chain.Name)
		}

		cmd := types.NewBurnTokenCommand(chainID, multisig.KeyID(keyID), ctx.BlockHeight(), *burnerInfo, token.IsExternal())
		if err := keeper.EnqueueCommand(ctx, cmd); err != nil {
			return nil, err
		}

		events.Emit(ctx, &types.BurnCommand{
			Chain:            chain.Name,
			CommandID:        cmd.ID,
			DestinationChain: deposit.DestinationChain,
			DepositAddress:   deposit.BurnerAddress.Hex(),
			Asset:            token.GetAsset(),
		})

		seen[burnerAddressHex] = true
	}

	return &types.CreateBurnTokensResponse{}, nil
}

func (s msgServer) CreatePendingTransfers(c context.Context, req *types.CreatePendingTransfersRequest) (*types.CreatePendingTransfersResponse, error) {
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
	transferLimit := keeper.GetParams(ctx).TransferLimit
	pageRequest := &query.PageRequest{
		Key:        nil,
		Offset:     0,
		Limit:      transferLimit,
		CountTotal: false,
		Reverse:    false,
	}
	pendingTransfers, _, err := s.nexus.GetTransfersForChainPaginated(ctx, chain, nexus.Pending, pageRequest)
	if err != nil {
		return nil, err
	}

	if len(pendingTransfers) == 0 {
		s.Logger(ctx).Debug("no pending transfers found")
		return &types.CreatePendingTransfersResponse{}, nil
	}

	keyID, ok := s.multisig.GetCurrentKeyID(ctx, chain.Name)
	if !ok {
		return nil, fmt.Errorf("current key not set for chain %s", chain.Name)
	}

	for _, transfer := range pendingTransfers {
		token := keeper.GetERC20TokenByAsset(ctx, transfer.Asset.Denom)
		if !token.Is(types.Confirmed) {
			s.Logger(ctx).Debug(fmt.Sprintf("token %s is not confirmed on %s", token.GetAsset(), chain.Name))
			continue
		}

		cmd, err := token.CreateMintCommand(multisig.KeyID(keyID), transfer)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "failed create mint-token command for transfer %d", transfer.ID)
		}

		s.Logger(ctx).Info(fmt.Sprintf("minting %s to recipient %s on %s with transfer ID %s and command ID %s", transfer.Asset.String(), transfer.Recipient.Address, transfer.Recipient.Chain.Name, transfer.ID.String(), cmd.ID.Hex()),
			types.AttributeKeyDestinationChain, transfer.Recipient.Chain.Name,
			types.AttributeKeyDestinationAddress, transfer.Recipient.Address,
			sdk.AttributeKeyAmount, transfer.Asset.String(),
			types.AttributeKeyAsset, transfer.Asset.Denom,
			types.AttributeKeyTransferID, transfer.ID.String(),
			types.AttributeKeyCommandsID, cmd.ID.Hex(),
		)

		if err := keeper.EnqueueCommand(ctx, cmd); err != nil {
			return nil, err
		}

		events.Emit(ctx, &types.MintCommand{
			Chain:              chain.Name,
			TransferID:         transfer.ID,
			CommandID:          cmd.ID,
			DestinationChain:   transfer.Recipient.Chain.Name,
			DestinationAddress: transfer.Recipient.Address,
			Asset:              transfer.Asset,
		})

		s.nexus.ArchivePendingTransfer(ctx, transfer)
	}

	return &types.CreatePendingTransfersResponse{}, nil
}

func (s msgServer) CreateTransferOperatorship(c context.Context, req *types.CreateTransferOperatorshipRequest) (*types.CreateTransferOperatorshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	keeper, err := s.ForChain(ctx, req.Chain)
	if err != nil {
		return nil, err
	}

	if _, ok := keeper.GetGatewayAddress(ctx); !ok {
		return nil, fmt.Errorf("scalar gateway address not set")
	}

	cmd, err := s.createTransferKeyCommand(ctx, keeper, req.Chain, multisig.KeyID(req.KeyID))
	if err != nil {
		return nil, err
	}

	if err := keeper.EnqueueCommand(ctx, cmd); err != nil {
		return nil, err
	}

	return &types.CreateTransferOperatorshipResponse{}, nil
}

func (s msgServer) createTransferKeyCommand(ctx sdk.Context, keeper types.ChainKeeper, chainStr nexus.ChainName, nextKeyID multisig.KeyID) (types.Command, error) {
	chain, ok := s.nexus.GetChain(ctx, chainStr)
	if !ok {
		return types.Command{}, fmt.Errorf("%s is not a registered chain", chainStr)
	}

	if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
		return types.Command{}, err
	}

	chainID, ok := keeper.GetChainID(ctx)
	if !ok {
		return types.Command{}, fmt.Errorf("could not find chain ID for '%s'", chainStr)
	}

	if _, ok := s.multisig.GetNextKeyID(ctx, chain.Name); ok {
		return types.Command{}, sdkerrors.Wrapf(types.ErrRotationInProgress, "finish rotating to next key for chain %s first", chain.Name)
	}

	if err := s.multisig.AssignKey(ctx, chain.Name, nextKeyID); err != nil {
		return types.Command{}, err
	}

	keyID, ok := s.multisig.GetCurrentKeyID(ctx, chain.Name)
	if !ok {
		return types.Command{}, fmt.Errorf("current key not set for chain %s", chain.Name)
	}

	nextKey, ok := s.multisig.GetKey(ctx, nextKeyID)
	if !ok {
		return types.Command{}, fmt.Errorf("could not find threshold key '%s'", nextKeyID)
	}

	return types.NewMultisigTransferCommand(chainID, keyID, nextKey), nil
}

func getCommandBatchToSign(ctx sdk.Context, keeper types.ChainKeeper) (types.CommandBatch, error) {
	latest := keeper.GetLatestCommandBatch(ctx)

	switch latest.GetStatus() {
	case types.BatchSigning:
		return types.CommandBatch{}, sdkerrors.Wrapf(types.ErrSignCommandsInProgress, "command batch '%s'", hex.EncodeToString(latest.GetID()))
	case types.BatchAborted:
		return latest, nil
	default:
		return keeper.CreateNewBatchToSign(ctx)
	}
}

func (s msgServer) AddChain(c context.Context, req *types.AddChainRequest) (*types.AddChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if _, found := s.nexus.GetChain(ctx, req.Name); found {
		return nil, fmt.Errorf("chain '%s' is already registered", req.Name)
	}

	if err := req.Params.Validate(); err != nil {
		return nil, err
	}

	chain := nexus.Chain{Name: req.Name, SupportsForeignAssets: true, KeyType: tss.Multisig, Module: types.ModuleName}
	s.nexus.SetChain(ctx, chain)

	if err := s.CreateChain(ctx, req.Params); err != nil {
		return nil, err
	}
	events.Emit(ctx, &types.ChainAdded{Chain: req.Name})

	return &types.AddChainResponse{}, nil
}

func (s msgServer) RetryFailedEvent(c context.Context, req *types.RetryFailedEventRequest) (*types.RetryFailedEventResponse, error) {
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

	event, ok := keeper.GetEvent(ctx, req.EventID)
	if !ok {
		return nil, fmt.Errorf("event %s not found for chain %s", req.EventID, req.Chain)
	}

	if event.Status != types.EventFailed {
		return nil, fmt.Errorf("event %s is not a failed event", req.EventID)
	}

	event.Status = types.EventConfirmed
	keeper.GetConfirmedEventQueue(ctx).Enqueue(getEventKey(req.EventID), &event)

	s.Logger(ctx).Info(
		"re-queued failed event",
		types.AttributeKeyChain, chain.Name,
		"eventID", event.GetID(),
	)

	events.Emit(ctx, &types.ChainEventRetryFailed{
		Chain:   event.Chain,
		EventID: event.GetID(),
		Type:    event.GetEventType(),
	})

	return &types.RetryFailedEventResponse{}, nil
}

func (s msgServer) CreateSnapshot(ctx sdk.Context, chain nexus.Chain) (snapshot.Snapshot, error) {
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

func (s msgServer) initializePoll(ctx sdk.Context, chain nexus.Chain, txID types.Hash) (vote.PollParticipants, error) {
	keeper, err := s.ForChain(ctx, chain.Name)
	if err != nil {
		return vote.PollParticipants{}, err
	}

	params := keeper.GetParams(ctx)
	snap, err := s.CreateSnapshot(ctx, chain)
	if err != nil {
		return vote.PollParticipants{}, err
	}

	pollID, err := s.voter.InitializePoll(
		ctx,
		vote.NewPollBuilder(types.ModuleName, params.VotingThreshold, snap, ctx.BlockHeight()+params.RevoteLockingPeriod).
			MinVoterCount(params.MinVoterCount).
			RewardPoolName(chain.Name.String()).
			GracePeriod(keeper.GetParams(ctx).VotingGracePeriod).
			ModuleMetadata(&types.PollMetadata{
				Chain: chain.Name,
				TxID:  txID,
			}),
	)

	return vote.PollParticipants{
		PollID:       pollID,
		Participants: snap.GetParticipantAddresses(),
	}, err
}
