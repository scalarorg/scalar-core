package keeper

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/x/chains/types"
)

func (s msgServer) SignBTCCommands(c context.Context, req *types.SignBTCCommandsRequest) (*types.SignBTCCommandsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	chain, ok := s.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if !types.IsBitcoinChain(chain.Name) {
		return nil, fmt.Errorf("chain %s is not a BTC chain", chain.Name)
	}

	if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
		return nil, err
	}

	keeper, err := s.ForChain(ctx, chain.Name)
	if err != nil {
		return nil, err
	}

	if _, ok := keeper.GetChainID(ctx); !ok {
		return nil, fmt.Errorf("could not find chain ID for '%s'", chain.Name)
	}

	commandBatch, err := getCommandBatchToSign(ctx, keeper)
	if err != nil {
		return nil, err
	}

	if len(commandBatch.GetCommandIDs()) == 0 {
		return &types.SignBTCCommandsResponse{CommandCount: 0, BatchedCommandsID: nil}, nil
	}

	// TODO: validate the psbt with the commands: check the outputs map 1-1 with the command payloads, check the amount of inputs is greater than the amount of outputs, check the format of psbt by btcd-lib.packet
	// use psbt.ValidateBasic()

	if err := s.covenant.SignPsbt(
		ctx,
		commandBatch.GetKeyID(),
		req.Psbt,
		types.ModuleName,
		chain.Name,
		types.NewSigMetadata(types.SigCommand, chain.Name, commandBatch.GetID()),
	); err != nil {
		return nil, err
	}

	if !commandBatch.SetStatus(types.BatchSigning) {
		return nil, fmt.Errorf("failed setting status of command batch %s to be signing", hex.EncodeToString(commandBatch.GetID()))
	}

	clog.Yellowf("[keeper] [msg_server_sign_btc_commands] commandBatch: %+v", commandBatch)

	batchedCommandsIDHex := hex.EncodeToString(commandBatch.GetID())
	commandList := types.CommandIDsToStrings(commandBatch.GetCommandIDs())
	for _, commandID := range commandList {
		s.Logger(ctx).Info(
			fmt.Sprintf("signing command %s in batch %s for chain %s using key %s", commandID, batchedCommandsIDHex, chain.Name, string(commandBatch.GetKeyID())),
			types.AttributeKeyChain, chain.Name,
			types.AttributeKeyKeyID, string(commandBatch.GetKeyID()),
			"commandBatchID", batchedCommandsIDHex,
			"commandID", commandID,
		)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSign,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.AttributeValueStart),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeyChain, chain.Name.String()),
			sdk.NewAttribute(sdk.AttributeKeySender, req.Sender.String()),
			sdk.NewAttribute(types.AttributeKeyBatchedCommandsID, batchedCommandsIDHex),
			sdk.NewAttribute(types.AttributeKeyCommandsIDs, strings.Join(commandList, ",")),
		),
	)

	return &types.SignBTCCommandsResponse{CommandCount: uint32(len(commandBatch.GetCommandIDs())), BatchedCommandsID: commandBatch.GetID()}, nil
}
