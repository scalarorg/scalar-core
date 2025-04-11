package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/key"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/chains/types"
)

func (k ChainKeeper) GetRedeemSession(ctx sdk.Context, custodianGroupUID exported.Hash) (session types.RedeemSession, ok bool) {
	ok = k.getStore(ctx).GetNew(redeemSessionPrefix.Append(key.FromBz(custodianGroupUID.Bytes())), &session)
	return session, ok
}

func (k ChainKeeper) SetRedeemSession(ctx sdk.Context, session *types.RedeemSession) error {
	return k.getStore(ctx).SetNewValidated(redeemSessionPrefix.Append(key.FromBz(session.CustodianGroupUID.Bytes())), session)
}

func (k ChainKeeper) GetRedeemSessions(ctx sdk.Context) []types.RedeemSession {
	iter := k.getStore(ctx).IteratorNew(redeemSessionPrefix)
	defer iter.Close()
	var sessions []types.RedeemSession
	for ; iter.Valid(); iter.Next() {
		var session types.RedeemSession
		iter.UnmarshalValue(&session)
		sessions = append(sessions, session)
	}
	return sessions
}
