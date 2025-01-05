package types

import (
	"encoding/hex"
	fmt "fmt"

	"github.com/scalarorg/scalar-core/utils/clog"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
)

type Psbt []byte

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
