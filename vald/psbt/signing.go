package psbt

import (
	"bytes"
	"context"
	"fmt"
	"sync"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	go_utils "github.com/scalarorg/bitcoin-vault/go-utils/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	grpc_client "github.com/scalarorg/scalar-core/vald/grpc-client"
	"github.com/scalarorg/scalar-core/x/chains/types"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
	covenantTypes "github.com/scalarorg/scalar-core/x/covenant/types"
)

type signResult struct {
	index int
	sigs  *exported.TapScriptSigsMap
}

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

	multiPsbt := event.GetMultiPsbt()
	if multiPsbt == nil {
		return fmt.Errorf("multiPsbt is nil")
	}

	if !mgr.validateKeyID(keyUID) {
		return fmt.Errorf("invalid keyID")
	}

	n := len(multiPsbt)
	resultChan := make(chan signResult, n)
	errChan := make(chan error, n)
	orderedResults := make([]*exported.TapScriptSigsMap, n)

	var wg sync.WaitGroup
	for i, psbt := range multiPsbt {
		wg.Add(1)
		go func(index int, p covenantTypes.Psbt) {
			defer wg.Done()

			mapOfTapScriptSigs, err := mgr.sign(keyUID, p, go_utils.NetworkKind(chainParams.Params.NetworkKind))
			if err != nil {
				clog.Redf("ProcessSigningPsbtStarted/sign error: %v", err)
				errChan <- err
				return
			}

			for i, tapScriptSig := range mapOfTapScriptSigs.Inner {
				clog.Yellowf("ProcessSigningPsbtStarted, tapScriptSig[%d]: %+v", i, tapScriptSig)
			}

			clog.Greenf("operator %s sending signature for signing %d", partyUID, event.GetSigID())

			resultChan <- signResult{
				index: index,
				sigs:  mapOfTapScriptSigs,
			}
		}(i, psbt)
	}

	go func() {
		wg.Wait()
		close(errChan)
		close(resultChan)
	}()

	for i := 0; i < n; i++ {
		select {
		case err, ok := <-errChan:
			if ok && err != nil {
				return err
			}
		case result, ok := <-resultChan:
			if ok {
				if result.sigs == nil {
					return fmt.Errorf("result.sigs is nil")
				}
				clog.Greenf("ProcessSigningPsbtStarted, result: %+v", result)
				orderedResults[result.index] = result.sigs
			}
		}
	}

	msg := covenantTypes.NewSubmitTapScriptSigsRequest(mgr.ctx.FromAddress, event.GetSigID(), orderedResults)

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
