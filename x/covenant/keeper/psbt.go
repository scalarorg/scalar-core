package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
)

func (k Keeper) CreateAndSignPsbt(ctx sdk.Context, keyID multisig.KeyID, payloadHash multisig.Hash, module string, moduleMetadata ...codec.ProtoMarshaler) error {
	clog.Green("Creating and signing PSBT", "keyID", keyID, "payloadHash", payloadHash, "module", module, "moduleMetadata", moduleMetadata)

	if !k.GetCovenantRouter().HasHandler(module) {
		panic(fmt.Errorf("covenant handler not registered for module %s", module))
	}

	// key, ok := k.getKey(ctx, keyID)
	// if !ok {
	// 	return fmt.Errorf("key %s not found", keyID)
	// }
	// if key.State != exported.Active {
	// 	return fmt.Errorf("key %s is not activated yet", keyID)
	// }

	// params := k.GetParams(ctx)

	// expiresAt := ctx.BlockHeight() + params.SigningTimeout
	// signingSession := types.NewSigningSession(k.nextSigID(ctx), key, payloadHash, expiresAt, params.SigningGracePeriod, module, moduleMetadata...)
	// if err := signingSession.ValidateBasic(); err != nil {
	// 	return err
	// }

	// k.setSigningSession(ctx, signingSession)

	// events.Emit(ctx, types.NewSigningStarted(signingSession.GetID(), key, payloadHash[:], module))
	// k.Logger(ctx).Info("signing session started",
	// 	"sig_id", signingSession.GetID(),
	// 	"key_id", key.GetID(),
	// 	"participant_count", len(key.GetPubKeys()),
	// 	"participants", strings.Join(slices.Map(key.GetParticipants(), sdk.ValAddress.String), ", "),
	// 	"participants_weight", key.GetParticipantsWeight().String(),
	// 	"bonded_weight", key.GetSnapshot().BondedWeight.String(),
	// 	"signing_threshold", key.GetSigningThreshold().String(),
	// 	"expires_at", expiresAt,
	// )

	return nil
}
