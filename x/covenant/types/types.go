package types

import (
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
)

type Psbt []byte

func (p Psbt) ValidateBasic() error {
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
