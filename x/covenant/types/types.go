package types

import (
	"encoding/hex"
	fmt "fmt"

	"github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	exported "github.com/scalarorg/scalar-core/x/covenant/exported"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
)

type Psbt []byte

func (p Psbt) Bytes() []byte {
	return p
}

var EmptyPsbt = []byte{}

func PsbtFromHex(h string) (Psbt, error) {
	psbt, err := hex.DecodeString(h)
	if err != nil {
		return nil, err
	}
	return psbt, nil
}

func (p Psbt) ValidateBasic() error {
	// TODO: validate psbt format by btcd-lib.packet
	clog.Yellow("!! TODO: validate psbt", "psbt", p)
	return nil
}

type PsbtPayload []byte

func (p PsbtPayload) ValidateBasic() error {
	if len(p) == 0 {
		return fmt.Errorf("can't be empty")
	}
	return nil
}

type TapScriptSig []byte

func (g CustodianGroup) CreateKey() multisigTypes.Key {
	pubKeys := map[string]multisig.PublicKey{}
	for _, custodian := range g.Custodians {
		pubKeys[custodian.Name] = custodian.BtcPubkey
	}
	key := multisigTypes.Key{
		ID:      multisig.KeyID(g.Uid),
		PubKeys: pubKeys,
	}
	return key
}

var DefaultParticipantTapScriptSigs = make(map[string]*exported.TapScriptSigList)

func (p *PsbtMultiSig) Finalize() error {
	psbtBytes := p.Psbt.Bytes()
	var err error
	for _, list := range p.ParticipantTapScriptSigs {
		inputTapscriptSigs := slices.Map(list.TapScriptSigs, func(sig *exported.TapScriptSig) vault.TapScriptSig {
			keyXOnly := sig.KeyXOnly.Bytes()
			leafHash := sig.LeafHash.Bytes()
			signature := sig.Signature.Bytes()
			return vault.TapScriptSig{
				KeyXOnly:  keyXOnly,
				LeafHash:  leafHash,
				Signature: signature,
			}
		})
		psbtBytes, err = vault.AggregateTapScriptSigs(psbtBytes, inputTapscriptSigs)
		if err != nil {
			return err
		}
	}
	clog.Greenf("CovenantHandler: Finalize, Psbt: %x", psbtBytes)

	tx, err := vault.FinalizePsbtAndExtractTx(psbtBytes)
	if err != nil {
		return err
	}

	p.FinalizedTx = tx
	p.Psbt = psbtBytes
	return nil
}
