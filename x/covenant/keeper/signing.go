package keeper

import (
	"fmt"

	"github.com/scalarorg/scalar-core/utils/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	gogoprototypes "github.com/gogo/protobuf/types"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
)

func (k Keeper) GetSigningSessions(ctx sdk.Context) (signingSessions []types.SigningSession, ok bool) {
	return nil, false
}

func (k Keeper) setSigningSession(ctx sdk.Context, signing types.SigningSession) {
	// the deletion is necessary because we may update it to a different location depending on the current state of the session
	k.getStore(ctx).Delete(expirySigningPrefix.Append(utils.KeyFromInt(signing.ExpiresAt)).Append(utils.KeyFromInt(signing.GetID())))

	k.getStore(ctx).Set(getSigningSessionExpiryKey(signing), &gogoprototypes.UInt64Value{Value: signing.GetID()})

	k.getStore(ctx).Set(getSigningSessionKey(signing.GetID()), &signing)
}

func (k Keeper) getSigningSession(ctx sdk.Context, id uint64) (signing types.SigningSession, ok bool) {
	return signing, k.getStore(ctx).Get(getSigningSessionKey(id), &signing)
}

func getSigningSessionExpiryKey(signing types.SigningSession) utils.Key {
	expiry := signing.ExpiresAt
	if signing.State == exported.Completed {
		expiry = math.Min(signing.ExpiresAt, signing.CompletedAt+signing.GracePeriod+1)
	}

	return expirySigningPrefix.Append(utils.KeyFromInt(expiry)).Append(utils.KeyFromInt(signing.GetID()))
}

func getSigningSessionKey(id uint64) utils.Key {
	return signingPrefix.Append(utils.KeyFromInt(id))
}

func (k Keeper) getSigningSessions(ctx sdk.Context) (signingSessions []types.SigningSession) {
	iter := k.getStore(ctx).Iterator(signingPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))

	for ; iter.Valid(); iter.Next() {
		var signingSession types.SigningSession
		iter.UnmarshalValue(&signingSession)

		signingSessions = append(signingSessions, signingSession)
	}

	return signingSessions
}

func (k Keeper) setSigningSessionCount(ctx sdk.Context, count uint64) {
	k.getStore(ctx).Set(signingSessionCountKey, &gogoprototypes.UInt64Value{Value: count})
}

func (k Keeper) nextSigID(ctx sdk.Context) uint64 {
	var val gogoprototypes.UInt64Value
	k.getStore(ctx).Get(signingSessionCountKey, &val)
	defer k.getStore(ctx).Set(signingSessionCountKey, &gogoprototypes.UInt64Value{Value: val.Value + 1})

	return val.Value
}

// DeleteSigningSession deletes the signing session with the given ID
func (k Keeper) DeleteSigningSession(ctx sdk.Context, id uint64) {
	signing, ok := k.getSigningSession(ctx, id)
	if !ok {
		return
	}

	k.getStore(ctx).Delete(getSigningSessionExpiryKey(signing))
	k.getStore(ctx).Delete(getSigningSessionKey(id))
}

// GetSigningSessionsByExpiry returns all signing sessions that either expires at
// or goes out of the grace period at the given block height
func (k Keeper) GetSigningSessionsByExpiry(ctx sdk.Context, expiry int64) []types.SigningSession {
	var results []types.SigningSession

	iter := k.getStore(ctx).Iterator(expirySigningPrefix.Append(utils.KeyFromInt(expiry)))
	defer utils.CloseLogError(iter, k.Logger(ctx))

	for ; iter.Valid(); iter.Next() {
		var value gogoprototypes.UInt64Value
		iter.UnmarshalValue(&value)

		sigID := value.Value
		result, ok := k.getSigningSession(ctx, sigID)
		if !ok {
			panic(fmt.Errorf("signing session %d not found", sigID))
		}

		results = append(results, result)
	}

	return results
}
