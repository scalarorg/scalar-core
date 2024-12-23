package covenant

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/scalarorg/scalar-core/x/covenant/keeper"
	"github.com/scalarorg/scalar-core/x/covenant/types"
)

// NewHandler returns the handler of the EVM module
func NewHandler(k types.CovenantKeeper, v types.Voter, snapshotter types.Snapshotter, staking types.StakingKeeper, slashing types.SlashingKeeper) sdk.Handler {
	server := keeper.NewMsgServerImpl(k, v, snapshotter, staking, slashing)
	h := func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case *types.CreateCustodianRequest:
			res, err := server.CreateCustodian(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.CreateCustodianGroupRequest:
			res, err := server.CreateCustodianGroup(sdk.WrapSDKContext(ctx), msg)
			result, err := sdk.WrapServiceResult(ctx, res, err)
			if err == nil {
				result.Log = fmt.Sprintf("successfully create custodian group")
			}
			return result, err
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
				fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg))

		}
	}

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		res, err := h(ctx, msg)
		if err != nil {
			// if the error is not already a registered error the error message would be obscured, so wrap it in a general registered error
			k.Logger(ctx).Debug(err.Error())
			return nil, err
		}
		if len(res.Log) > 0 {
			k.Logger(ctx).Debug(res.Log)
		}
		return res, nil
	}
}
