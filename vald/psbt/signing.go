package psbt

import (
	"bytes"
	"context"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	go_utils "github.com/scalarorg/bitcoin-vault/go-utils/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	grpc_client "github.com/scalarorg/scalar-core/vald/grpc-client"
	"github.com/scalarorg/scalar-core/x/chains/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	covenantTypes "github.com/scalarorg/scalar-core/x/covenant/types"
)

// TODO: Validate psbt inputs whether they are available in the btc chain

// ProcessSigningPsbtStarted handles event signing psbt started
func (mgr *Mgr) ProcessSigningPsbtStarted(event *covenantTypes.SigningPsbtStarted) error {
	if !types.IsBitcoinChain(event.Chain) {
		return nil
	}

	mgrParticipant := mgr.valAddr.String()

	pubKey, ok := event.PubKeys[mgrParticipant]
	if !ok {
		clog.Redf("ProcessSigningPsbtStarted/pubKey not found for %s, got event.PubKeys: %+v", mgrParticipant, event.PubKeys)
		return nil
	}

	if !mgr.validatePubKey(pubKey) {
		clog.Redf("ProcessSigningPsbtStarted/invalid pubKey for %s", mgrParticipant)
		return fmt.Errorf("invalid pubKey")
	}

	// TODO: validate the psbt inputs whether they are available in the btc chain

	keyUID := fmt.Sprintf("%s_%d", event.GetKeyID().String(), 0)
	partyUID := mgr.valAddr.String()

	chainParams, err := grpc_client.QueryManager().GetChainsClient().Params(context.Background(), &chainsTypes.ParamsRequest{
		Chain: event.Chain.String(),
	})
	if err != nil {
		return err
	}

	// clog.Greenf("ProcessSigningPsbtStarted/chainParams: %+v", chainParams)
	// chainInfoBytes, err := scalarUtils.ChainInfoBytesFromID(event.Chain.String())
	// if err != nil {
	// 	return err
	// }

	// client, ok := mgr.rpcs[chainInfoBytes]
	// if !ok {
	// 	return fmt.Errorf("client not found for chain %s", event.Chain.String())
	// }

	// if err := mgr.ValidatePsbt(client, event.Psbt); err != nil {
	// 	return err
	// }

	mapOfTapScriptSigs, err := mgr.sign(keyUID, event.Psbt, go_utils.NetworkKind(chainParams.Params.NetworkKind))
	if err != nil {
		clog.Redf("ProcessSigningPsbtStarted/sign error: %v", err)
		return err
	}

	for i, tapScriptSig := range mapOfTapScriptSigs.Inner {
		clog.Yellowf("ProcessSigningPsbtStarted, tapScriptSig[%d]: %+v", i, tapScriptSig)
	}

	clog.Greenf("operator %s sending signature for signing %d", partyUID, event.GetSigID())

	msg := covenantTypes.NewSubmitTapScriptSigsRequest(mgr.ctx.FromAddress, event.GetSigID(), mapOfTapScriptSigs)

	clog.Greenf("SubmitTapScriptSigsRequest: %+v", msg)
	if _, err := mgr.b.Broadcast(context.Background(), msg); err != nil {
		return sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing submit signature message")
	}

	return nil
}

func (mgr *Mgr) validatePubKey(pubKey []byte) bool {
	if len(pubKey) != 33 {
		return false
	}

	return bytes.Equal(pubKey, mgr.pubKey.SerializeCompressed())
}
