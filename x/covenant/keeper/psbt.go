package keeper

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/utils/slices"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// TODO: Currently, we are mocking the psbt, we need to split it into two events
// CreatingPsbtStarted and SigningPsbtStarted

func (k Keeper) SignPsbt(ctx sdk.Context, keyID multisig.KeyID, multiPsbt []types.Psbt, module string, chainName nexus.ChainName, moduleMetadata ...codec.ProtoMarshaler) error {
	if !k.GetCovenantRouter().HasHandler(module) {
		panic(fmt.Errorf("covenant handler not registered for module %s", module))
	}
	clog.Yellowf("[CovenantKeeper] [SignPsbt] keyID: %s, module: %s, chainName: %s", keyID, module, chainName)
	key, ok := k.GetKey(ctx, keyID)
	if !ok {
		return fmt.Errorf("key %s not found", keyID)
	}

	if key.State != multisig.Active {
		return fmt.Errorf("key %s is not activated yet", keyID)
	}

	params := k.GetParams(ctx)
	clog.Red("SigningPsbt", "params", params)

	expiresAt := ctx.BlockHeight() + params.SigningTimeout
	signingSession := types.NewSigningSession(&types.NewSigningSessionParams{
		ID:             k.nextSigID(ctx),
		Key:            key,
		MultiPsbt:      multiPsbt,
		ExpiresAt:      expiresAt,
		GracePeriod:    params.SigningGracePeriod,
		Module:         module,
		ModuleMetadata: moduleMetadata,
	})

	if err := signingSession.ValidateBasic(); err != nil {
		return err
	}

	k.setSigningSession(ctx, signingSession)

	events.Emit(ctx, types.NewSigningPsbtStarted(signingSession.GetID(), key, multiPsbt, module, chainName))

	k.Logger(ctx).Info("create and signing psbt started",
		"sig_id", signingSession.GetID(),
		"key_id", key.GetID(),
		"participant_count", len(key.GetPubKeys()),
		"participants", strings.Join(slices.Map(key.GetParticipants(), sdk.ValAddress.String), ", "),
		"participants_weight", key.GetParticipantsWeight().String(),
		"bonded_weight", key.GetSnapshot().BondedWeight.String(),
		"signing_threshold", key.GetSigningThreshold().String(),
		"expires_at", expiresAt,
	)

	return nil
}
