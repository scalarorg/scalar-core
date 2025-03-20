package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
)

func (k Keeper) AddRedeemSession(ctx sdk.Context, tokenSymbol string, session *types.RedeemSession) {
	k.getStore(ctx).Set(redeemSessionPrefix.Append(utils.KeyFromStr(tokenSymbol)), session)
}

func (k Keeper) GetRedeemSession(ctx sdk.Context, tokenSymbol string) (*types.RedeemSession, bool) {
	session := types.RedeemSession{}
	ok := k.getStore(ctx).Get(redeemSessionPrefix.Append(utils.KeyFromStr(tokenSymbol)), &session)
	return &session, ok
}
