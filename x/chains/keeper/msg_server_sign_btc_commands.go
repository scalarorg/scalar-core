package keeper

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/bitcoin-vault/go-utils/encode"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/x/chains/types"
)

func (s msgServer) SignBtcCommand(c context.Context, req *types.SignBtcCommandsRequest) (*types.SignCommandsResponse, error) {
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
	// getCommandBatchToSign get latest command batch or create new one with common keyId as first enqueued commnand
	// Other commands with different keyId will be process in next cycle
	// Todo: Handle simultaneous command batches with different keyId
	commandBatch, err := getCommandBatchToSign(ctx, keeper)
	if err != nil {
		return nil, err
	}

	if len(commandBatch.GetCommandIDs()) == 0 {
		return &types.SignCommandsResponse{CommandCount: 0, BatchedCommandsID: nil}, nil
	}

	clog.Yellowf("[keeper] [msg_server_sign_btc_commands] psbt is empty, try extract psbt from first command")
	if len(commandBatch.GetCommandIDs()) > 1 {
		clog.Yellowf("[keeper] [msg_server_sign_btc_commands] more than one command in batch, only support one command for now")
		return nil, fmt.Errorf("more than one command in batch, only support one command for now")
	}
	commandId := commandBatch.GetCommandIDs()[0]
	command, ok := keeper.GetCommand(ctx, commandId)
	if !ok {
		return nil, fmt.Errorf("command %s not found", commandId)
	}
	payload, err := encode.DecodeContractCallWithTokenPayload(command.Payload)
	if err != nil {
		return nil, err
	}
	if payload.PayloadType != encode.ContractCallWithTokenPayloadType_UPC {
		return nil, fmt.Errorf("command %s is not a contract call with token in UPC model", commandId)
	}

	clog.Redf("SignBtcCommand, [PSBT]>: %x", payload.UPC.Psbt)
	if err := s.covenant.SignPsbt(
		ctx,
		commandBatch.GetKeyID(),
		payload.UPC.Psbt,
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

	return &types.SignCommandsResponse{CommandCount: uint32(len(commandBatch.GetCommandIDs())), BatchedCommandsID: commandBatch.GetID()}, nil
}
