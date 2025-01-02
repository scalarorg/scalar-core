package keeper

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/funcs"

	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/x/chains/types"
	covenant "github.com/scalarorg/scalar-core/x/covenant/exported"
)

type covHandler struct {
	cdc    codec.Codec
	keeper types.BaseKeeper
}

// NewCovenantHandler returns the handler for processing signatures delivered by the covenant module
func NewCovenantHandler(cdc codec.Codec, keeper types.BaseKeeper) covenant.CovenantHandler {
	return covHandler{
		cdc:    cdc,
		keeper: keeper,
	}
}

func (s covHandler) HandleCompleted(ctx sdk.Context, sig utils.ValidatedProtoMarshaler, moduleMetadata codec.ProtoMarshaler) error {
	sigMetadata := moduleMetadata.(*types.SigMetadata)
	commandBatch, err := s.getCommandBatch(ctx, sigMetadata)
	if err != nil {
		return err
	}

	funcs.MustNoErr(commandBatch.SetSigned(sig))

	events.Emit(ctx, types.NewCommandBatchSigned(sigMetadata.Chain, sigMetadata.CommandBatchID))

	return nil
}

func (s covHandler) HandleFailed(ctx sdk.Context, moduleMetadata codec.ProtoMarshaler) error {
	sigMetadata := moduleMetadata.(*types.SigMetadata)
	commandBatch, err := s.getCommandBatch(ctx, sigMetadata)
	if err != nil {
		return err
	}

	ok := commandBatch.SetStatus(types.BatchAborted)
	if !ok {
		panic(fmt.Errorf("failed to abort command batch %s", hex.EncodeToString(commandBatch.GetID())))
	}

	events.Emit(ctx, types.NewCommandBatchAborted(sigMetadata.Chain, sigMetadata.CommandBatchID))

	return nil
}

func (s covHandler) getCommandBatch(ctx sdk.Context, sigMetadata *types.SigMetadata) (types.CommandBatch, error) {
	ck, err := s.keeper.ForChain(ctx, sigMetadata.Chain)
	if err != nil {
		return types.CommandBatch{}, fmt.Errorf("chain %s does not exist as a chain", sigMetadata.Chain)
	}

	commandBatch := ck.GetBatchByID(ctx, sigMetadata.CommandBatchID)
	if !commandBatch.Is(types.BatchSigning) {
		return types.CommandBatch{}, fmt.Errorf("the command batch %s of chain %s is not being signed", hex.EncodeToString(sigMetadata.CommandBatchID), sigMetadata.Chain)
	}

	return commandBatch, nil
}
