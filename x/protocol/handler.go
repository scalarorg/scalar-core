package protocol

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/scalarorg/scalar-core/x/protocol/keeper"
	"github.com/scalarorg/scalar-core/x/protocol/types"
)

// NewHandler returns the handler of the Cosmos module
func NewHandler(k keeper.Keeper, covenant types.CovenantKeeper, permission types.PermissionKeeper) sdk.Handler {
	server := keeper.NewMsgServerImpl(k, covenant, permission)
	h := func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case *types.CreateProtocolRequest:
			res, err := server.CreateProtocol(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)

			return result, err
		case *types.UpdateProtocolRequest:
			res, err := server.UpdateProtocol(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)

			return result, err
		case *types.AddSupportedChainRequest:
			res, err := server.AddSupportedChain(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)

			return result, err
		case *types.UpdateSupportedChainRequest:
			res, err := server.UpdateSupportedChain(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)

			return result, err
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
				fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg))
		}
	}

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		res, err := h(ctx, msg)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrProtocol, err.Error())
		}
		return res, nil
	}
}
