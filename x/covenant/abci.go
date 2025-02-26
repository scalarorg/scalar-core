package covenant

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, _ abci.RequestBeginBlock, bk types.Keeper) {}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, _ abci.RequestEndBlock, bk types.Keeper, rewarder types.Rewarder) ([]abci.ValidatorUpdate, error) {
	clog.Greenf("Covenant EndBlocker, ctx.BlockHeight: %+v", ctx.BlockHeight())
	handleSignings(ctx, bk, rewarder)
	return nil, nil
}

func handleSignings(ctx sdk.Context, k types.Keeper, rewarder types.Rewarder) {
	// we handle sessions that'll expire on the next block,
	// to avoid waiting for an additional block
	for _, signing := range k.GetSigningSessionsByExpiry(ctx, ctx.BlockHeight()+1) {
		clog.Bluef("handleSignings, signing.GetID(): %+v", signing.GetID())
		_ = utils.RunCached(ctx, k, func(cachedCtx sdk.Context) ([]abci.ValidatorUpdate, error) {
			k.DeleteSigningSession(cachedCtx, signing.GetID())
			module := signing.GetModule()

			pool := rewarder.GetPool(cachedCtx, types.ModuleName)
			slices.ForEach(signing.GetMissingParticipants(), pool.ClearRewards)

			if signing.State != exported.Completed {
				events.Emit(cachedCtx, types.NewSigningPsbtExpired(signing.GetID()))
				k.Logger(cachedCtx).Info("signing session expired",
					"sig_id", signing.GetID(),
				)

				funcs.MustNoErr(k.GetCovenantRouter().GetHandler(module).HandleFailed(cachedCtx, signing.GetMetadata()))
				return nil, nil
			}

			// finalize the psbt
			err := FinalizeMultiPsbt(&signing.PsbtMultiSig)
			//serr := signing.PsbtMultiSig.Finalize()
			if err != nil {
				return nil, sdkerrors.Wrap(err, "failed to finalize psbt")
			}

			sig := funcs.Must(signing.Result())

			// TODO: must validate the signature in the submit signature request then release the rewards
			slices.ForEach(sig.GetParticipants(), func(p sdk.ValAddress) { funcs.MustNoErr(pool.ReleaseRewards(p)) })

			if err := k.GetCovenantRouter().GetHandler(module).HandleCompleted(cachedCtx, &sig, signing.GetMetadata()); err != nil {
				return nil, sdkerrors.Wrap(err, "failed to handle completed signature")
			}

			for index, p := range sig.GetParticipants() {
				clog.Greenf("CovenantHandler: HandleCompleted, Participant: %x", p)
				clog.Greenf("CovenantHandler: HandleCompleted, Psbts: %x", sig.GetMultiPsbt()[index].Bytes())
				clog.Greenf("CovenantHandler: HandleCompleted, FinalizedTx: %x", sig.GetFinalizedTxs()[index])
			}

			events.Emit(cachedCtx, types.NewSigningPsbtCompleted(signing.GetID()))
			k.Logger(cachedCtx).Info("signing session completed",
				"sig_id", signing.GetID(),
				"key_id", sig.GetKeyID(),
				"module", module,
			)

			return nil, nil
		})
	}
}

func FinalizeMultiPsbt(p *types.PsbtMultiSig) error {
	var finalizedTxs [][]byte

	var tapScriptSigsMapByEachPsbt = make([]map[string]*exported.TapScriptSigsMap, len(p.MultiPsbt))

	// collect the map for each psbt
	// ParticipantListTapScriptSigs = {
	// "Alice": [sigOfPsbt1, sigOfPsbt2, sigOfPsbt3],
	// "Bob": [sigOfPsbt1, sigOfPsbt2, sigOfPsbt3],
	// "Charlie": [sigOfPsbt1, sigOfPsbt2, sigOfPsbt3],
	//}
	// => output: [
	//    map[Alice:[sigOfPsbt1] Bob:[sigOfPsbt1] Charlie:[sigOfPsbt1]],
	//    map[Alice:[sigOfPsbt2] Bob:[sigOfPsbt2] Charlie:[sigOfPsbt2]],
	//    map[Alice:[sigOfPsbt3] Bob:[sigOfPsbt3] Charlie:[sigOfPsbt3]],
	// ]

	for party, listOfEachParty := range p.ParticipantListTapScriptSigs {
		for index, sig := range listOfEachParty.Inner {
			if tapScriptSigsMapByEachPsbt[index] == nil {
				tapScriptSigsMapByEachPsbt[index] = make(map[string]*exported.TapScriptSigsMap)
			}
			tapScriptSigsMapByEachPsbt[index][party] = sig
		}
	}

	for index, psbt := range p.MultiPsbt {
		psbtBytes := psbt.Bytes()
		var err error

		tapScriptSigsMap := tapScriptSigsMapByEachPsbt[index]

		for _, m := range tapScriptSigsMap {
			raw := m.ToRaw()
			psbtBytes, err = vault.AggregateTapScriptSigs(psbtBytes, raw)
			if err != nil {
				return err
			}
		}
		clog.Greenf("CovenantHandler: Finalize, Psbt: %x", psbtBytes)

		tx, err := vault.FinalizePsbtAndExtractTx(psbtBytes)
		if err != nil {
			return err
		}

		p.FinalizedTxs = append(finalizedTxs, tx)
		p.MultiPsbt[index] = psbtBytes
	}

	return nil
}
