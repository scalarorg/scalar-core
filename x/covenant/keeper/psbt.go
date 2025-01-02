package keeper

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/utils/slices"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

func (k Keeper) CreateAndSignPsbt(ctx sdk.Context, keyID multisig.KeyID, extraData [][]byte, module string, chainName nexus.ChainName, moduleMetadata ...codec.ProtoMarshaler) error {
	if !k.GetCovenantRouter().HasHandler(module) {
		panic(fmt.Errorf("covenant handler not registered for module %s", module))
	}

	key, ok := k.getKey(ctx, keyID)
	if !ok {
		return fmt.Errorf("key %s not found", keyID)
	}

	if key.State != multisig.Active {
		return fmt.Errorf("key %s is not activated yet", keyID)
	}

	params := k.GetParams(ctx)

	expiresAt := ctx.BlockHeight() + params.SigningTimeout
	batchPsbtPayload := slices.Map(extraData, func(payload []byte) types.PsbtPayload {
		return types.PsbtPayload(payload)
	})

	signingSession := types.NewSigningSession(&types.NewSigningSessionParams{
		ID:               k.nextSigID(ctx),
		Key:              key,
		BatchPsbtPayload: batchPsbtPayload,
		ExpiresAt:        expiresAt,
		GracePeriod:      params.SigningGracePeriod,
		Module:           module,
		ModuleMetadata:   moduleMetadata,
	})

	if err := signingSession.ValidateBasic(); err != nil {
		return err
	}

	k.setSigningSession(ctx, signingSession)

	events.Emit(ctx, types.NewCreateAndSigningPsbtStarted(signingSession.GetID(), key, batchPsbtPayload, module, chainName))

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
