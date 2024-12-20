package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/scalarorg/scalar-core/x/chains/types"
)

var _ types.QueryServiceServer = Querier{}

// Querier implements the grpc querier
type Querier struct {
	types.BaseKeeper
}

// NewGRPCQuerier returns a new Querier
func NewGRPCQuerier(k types.BaseKeeper) Querier {
	return Querier{
		BaseKeeper: k,
	}
}

// BatchedCommands implements the batched commands query
// If BatchedCommandsResponse.Id is set, it returns the latest batched commands with the specified id.
// Otherwise returns the latest batched commands.
func (q Querier) BatchedCommands(c context.Context, req *types.BatchedCommandsRequest) (*types.BatchedCommandsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	_ = ctx

	q.Logger(ctx).Debug("Not implemented batched commands query")

	// ck, err := q.keeper.ForChain(ctx, nexustypes.ChainName(req.Chain))
	// if err != nil {
	// 	return nil, status.Error(codes.NotFound, sdkerrors.Wrap(types.ErrEVM, fmt.Sprintf("%s is not a registered chain", req.Chain)).Error())
	// }

	// var commandBatch types.CommandBatch
	// switch req.Id {
	// case "":
	// 	commandBatch = ck.GetLatestCommandBatch(ctx)
	// 	if commandBatch.Is(types.BatchNonExistent) {
	// 		return nil, status.Error(codes.NotFound, sdkerrors.Wrap(types.ErrEVM, fmt.Sprintf("could not get the latest batched commands for chain %s", req.Chain)).Error())
	// 	}
	// default:
	// 	commandBatchID, err := utils.HexDecode(req.Id)
	// 	if err != nil {
	// 		return nil, status.Error(codes.InvalidArgument, sdkerrors.Wrap(types.ErrEVM, fmt.Sprintf("invalid batched commands ID: %v", err)).Error())
	// 	}

	// 	commandBatch = ck.GetBatchByID(ctx, commandBatchID)
	// 	if commandBatch.Is(types.BatchNonExistent) {
	// 		return nil, status.Error(codes.NotFound, sdkerrors.Wrap(types.ErrEVM, fmt.Sprintf("batched commands with ID %s not found", req.Id)).Error())
	// 	}
	// }

	// resp, err := commandBatchToResp(ctx, commandBatch, q.multisig)
	// if err != nil {
	// 	return nil, status.Error(codes.NotFound, err.Error())
	// }

	return &types.BatchedCommandsResponse{}, nil
}
