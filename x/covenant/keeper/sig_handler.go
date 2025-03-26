package keeper

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/funcs"

	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
)

type sigHandler struct {
	cdc    codec.Codec
	keeper types.Keeper
}

// NewSigHandler returns the handler for processing signatures delivered by the multisig module
func NewSigHandler(cdc codec.Codec, keeper types.Keeper) multisig.SigHandler {
	return sigHandler{
		cdc:    cdc,
		keeper: keeper,
	}
}

func (s sigHandler) HandleCompleted(ctx sdk.Context, sig utils.ValidatedProtoMarshaler, moduleMetadata codec.ProtoMarshaler) error {
	sigMetadata := moduleMetadata.(*types.SigMetadata)
	commandBatch, err := s.getCommand(ctx, sigMetadata)
	if err != nil {
		return err
	}

	funcs.MustNoErr(commandBatch.SetSigned(sig))

	events.Emit(ctx, types.NewStandaloneCommandSigned(sigMetadata.Chain, sigMetadata.CommandID))

	return nil
}

func (s sigHandler) HandleFailed(ctx sdk.Context, moduleMetadata codec.ProtoMarshaler) error {
	sigMetadata := moduleMetadata.(*types.SigMetadata)
	cmd, err := s.getCommand(ctx, sigMetadata)
	if err != nil {
		return err
	}

	ok := cmd.SetStatus(types.StandaloneCommandStatusAborted)
	if !ok {
		panic(fmt.Errorf("failed to abort command batch %s", hex.EncodeToString(cmd.GetID())))
	}

	events.Emit(ctx, types.NewStandaloneCommandAborted(sigMetadata.Chain, sigMetadata.CommandID))

	return nil
}

func (s sigHandler) getCommand(ctx sdk.Context, sigMetadata *types.SigMetadata) (types.StandaloneCommand, error) {
	command := s.keeper.GetReserveUTXOCommandByID(ctx, sigMetadata.CommandID)
	if !command.Is(types.StandaloneCommandStatusSigning) {
		return types.StandaloneCommand{}, fmt.Errorf("the command  %s of chain %s is not being signed", hex.EncodeToString(sigMetadata.CommandID), sigMetadata.Chain)
	}

	return command, nil
}
