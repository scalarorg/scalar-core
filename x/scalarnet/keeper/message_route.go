package keeper

import (
	"fmt"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/scalarorg/scalar-core/x/scalarnet/types"
)

// for IBC execution
const gasCost = storetypes.Gas(1000000)

// NewMessageRoute creates a new message route
func NewMessageRoute(
	ibcK types.IBCKeeper,
	feegrantK types.FeegrantKeeper,
	bankK types.BankKeeper,
	nexusK types.Nexus,
	stakingK types.StakingKeeper,
) nexus.MessageRoute {
	return func(ctx sdk.Context, routingCtx nexus.RoutingContext, msg nexus.GeneralMessage) error {
		if routingCtx.Payload == nil {
			return fmt.Errorf("payload is required for routing messages to a cosmos chain")
		}

		bz, err := types.TranslateMessage(msg, routingCtx.Payload)
		if err != nil {
			return sdkerrors.Wrap(err, "invalid payload")
		}

		asset, err := escrowAssetToMessageSender(ctx, feegrantK, bankK, nexusK, ibcK, stakingK, routingCtx, msg)
		if err != nil {
			return err
		}

		ctx.GasMeter().ConsumeGas(gasCost, "execute-message")

		return ibcK.SendMessage(sdk.WrapSDKContext(ctx), msg.Recipient, asset, string(bz), msg.ID)
	}
}

// all general messages are sent from the Scalar general message sender, so receiver can use the packet sender to authenticate the message
// escrowAssetToMessageSender sends the asset to general msg sender account
func escrowAssetToMessageSender(
	ctx sdk.Context,
	feegrantK types.FeegrantKeeper,
	bankK types.BankKeeper,
	nexusK types.Nexus,
	ibcK types.IBCKeeper,
	stakingK types.StakingKeeper,
	routingCtx nexus.RoutingContext,
	msg nexus.GeneralMessage,
) (sdk.Coin, error) {
	switch msg.Type() {
	case nexus.TypeGeneralMessage:
		// pure general message, take dust amount from sender to satisfy ibc transfer requirements
		asset := sdk.NewCoin(stakingK.BondDenom(ctx), sdk.OneInt())
		sender := routingCtx.Sender

		if !routingCtx.FeeGranter.Empty() {
			req := types.RouteMessageRequest{
				Sender:     routingCtx.Sender,
				ID:         msg.ID,
				Payload:    routingCtx.Payload,
				Feegranter: routingCtx.FeeGranter,
			}
			if err := feegrantK.UseGrantedFees(ctx, routingCtx.FeeGranter, routingCtx.Sender, sdk.NewCoins(asset), []sdk.Msg{&req}); err != nil {
				return sdk.Coin{}, err
			}

			sender = routingCtx.FeeGranter
		}

		return asset, bankK.SendCoins(ctx, sender, types.ScalarIBCAccount, sdk.NewCoins(asset))
	case nexus.TypeGeneralMessageWithToken:
		lockableAsset, err := nexusK.NewLockableAsset(ctx, ibcK, bankK, *msg.Asset)
		if err != nil {
			return sdk.Coin{}, err
		}

		return lockableAsset.GetCoin(ctx), lockableAsset.UnlockTo(ctx, types.ScalarIBCAccount)
	default:
		return sdk.Coin{}, fmt.Errorf("unrecognized message type")
	}
}
