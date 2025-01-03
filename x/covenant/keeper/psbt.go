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

// TODO: Currently, we are mocking the psbt, we need to split it into two events
// CreatingPsbtStarted and SigningPsbtStarted

var mockPsbt = []byte("mock_psbt")

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

	events.Emit(ctx, types.NewSigningPsbtStarted(signingSession.GetID(), key, batchPsbtPayload, module, chainName))

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


// pubkey: "0215da913b3e87b4932b1e1b87d9667c28e7250aa0ed60b3a31095f541e1641488", privkey: "f7981df190cd4e8009a5472adf3d6318dee2290698d2ad723e300fbdf80ea81c"
// pubkey: "02f0f3d9beaf7a3945bcaa147e041ae1d5ca029bde7e40d8251f0783d6ecbe8fb5", privkey: "80c44748c16bd5c50c5ae0fa5bd8280d4c9c6191de2d8be9b3b1f48124df54a6"
// pubkey: "03594e78c0a2968210d9c1550d4ad31b03d5e4b9659cf2f67842483bb3c2bb7811", privkey: "b5ebf3c7f1d78056176afdbe44fbddbc3b8ef7e1ed0331bf5c4f889a6ef9b52a"
// pubkey: "03b59e575cef873ea95273afd55956c84590507200d410e693e4b079a426cc6102", privkey: "f92d44713b18ec56bf387201b0439d8e8ef0731235d487f81c5f3d5f18a52af3"
// pubkey: "03e2d226cfdaec93903c3f3b81a01a81b19137627cb26e621a0afb7bcd6efbcfff", privkey: "5ddbfaf1f19eedc8999526d618fccbf2572b320ef7b0649d7c91ca5279189a50"



// 70736274ff0100a602000000022aab2ff2a776da8dc894306e83562776a664e56ec64d346d61b79c819965394a0000000000fdffffff6bdd8c7e85c6a5599ca62f758ecac1369ebc14fec9569c84678a7aa12a371bcd0000000000fdffffff02a11900000000000016001450dceca158a9c872eb405d52293d351110572c9ee8f10200000000002251207f815abf6dfd78423a708aa8db1c2c906eecac910c035132d342e4988a37b8d5000000000001012ba0860100000000002251207f815abf6dfd78423a708aa8db1c2c906eecac910c035132d342e4988a37b8d5010304000000002215c050929b74c1a04954b78b4b6035e97a5e078a5a0f28ec96d547bfee9ace803ac0ad2015da913b3e87b4932b1e1b87d9667c28e7250aa0ed60b3a31095f541e1641488ac20594e78c0a2968210d9c1550d4ad31b03d5e4b9659cf2f67842483bb3c2bb7811ba20b59e575cef873ea95273afd55956c84590507200d410e693e4b079a426cc6102ba20e2d226cfdaec93903c3f3b81a01a81b19137627cb26e621a0afb7bcd6efbcfffba20f0f3d9beaf7a3945bcaa147e041ae1d5ca029bde7e40d8251f0783d6ecbe8fb5ba53a2c0211615da913b3e87b4932b1e1b87d9667c28e7250aa0ed60b3a31095f541e164148825015a10a5ec729629c6dd863dc28b7162e18f96b00dedd87f158b228428a298bccb000000002116594e78c0a2968210d9c1550d4ad31b03d5e4b9659cf2f67842483bb3c2bb781125015a10a5ec729629c6dd863dc28b7162e18f96b00dedd87f158b228428a298bccb000000002116b59e575cef873ea95273afd55956c84590507200d410e693e4b079a426cc610225015a10a5ec729629c6dd863dc28b7162e18f96b00dedd87f158b228428a298bccb000000002116e2d226cfdaec93903c3f3b81a01a81b19137627cb26e621a0afb7bcd6efbcfff25015a10a5ec729629c6dd863dc28b7162e18f96b00dedd87f158b228428a298bccb000000002116f0f3d9beaf7a3945bcaa147e041ae1d5ca029bde7e40d8251f0783d6ecbe8fb525015a10a5ec729629c6dd863dc28b7162e18f96b00dedd87f158b228428a298bccb0000000001172050929b74c1a04954b78b4b6035e97a5e078a5a0f28ec96d547bfee9ace803ac00118205a10a5ec729629c6dd863dc28b7162e18f96b00dedd87f158b228428a298bccb0001012ba0860100000000002251207f815abf6dfd78423a708aa8db1c2c906eecac910c035132d342e4988a37b8d5010304000000002215c050929b74c1a04954b78b4b6035e97a5e078a5a0f28ec96d547bfee9ace803ac0ad2015da913b3e87b4932b1e1b87d9667c28e7250aa0ed60b3a31095f541e1641488ac20594e78c0a2968210d9c1550d4ad31b03d5e4b9659cf2f67842483bb3c2bb7811ba20b59e575cef873ea95273afd55956c84590507200d410e693e4b079a426cc6102ba20e2d226cfdaec93903c3f3b81a01a81b19137627cb26e621a0afb7bcd6efbcfffba20f0f3d9beaf7a3945bcaa147e041ae1d5ca029bde7e40d8251f0783d6ecbe8fb5ba53a2c0211615da913b3e87b4932b1e1b87d9667c28e7250aa0ed60b3a31095f541e164148825015a10a5ec729629c6dd863dc28b7162e18f96b00dedd87f158b228428a298bccb000000002116594e78c0a2968210d9c1550d4ad31b03d5e4b9659cf2f67842483bb3c2bb781125015a10a5ec729629c6dd863dc28b7162e18f96b00dedd87f158b228428a298bccb000000002116b59e575cef873ea95273afd55956c84590507200d410e693e4b079a426cc610225015a10a5ec729629c6dd863dc28b7162e18f96b00dedd87f158b228428a298bccb000000002116e2d226cfdaec93903c3f3b81a01a81b19137627cb26e621a0afb7bcd6efbcfff25015a10a5ec729629c6dd863dc28b7162e18f96b00dedd87f158b228428a298bccb000000002116f0f3d9beaf7a3945bcaa147e041ae1d5ca029bde7e40d8251f0783d6ecbe8fb525015a10a5ec729629c6dd863dc28b7162e18f96b00dedd87f158b228428a298bccb0000000001172050929b74c1a04954b78b4b6035e97a5e078a5a0f28ec96d547bfee9ace803ac00118205a10a5ec729629c6dd863dc28b7162e18f96b00dedd87f158b228428a298bccb000000