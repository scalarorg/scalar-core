package types

import (
	"encoding/hex"
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/clog"
	exported "github.com/scalarorg/scalar-core/x/covenant/exported"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	snapshot "github.com/scalarorg/scalar-core/x/snapshot/exported"
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

var DefaultParticipantTapScriptSigs = make(map[string]*exported.TapScriptSigsMap)

func (g *CustodianGroup) CreateKey(ctx sdk.Context, snapshot snapshot.Snapshot, threshold utils.Threshold) multisigTypes.Key {
	pubKeys := map[string]multisig.PublicKey{}
	for _, custodian := range g.Custodians {
		pubKeys[custodian.ValAddress] = custodian.BitcoinPubkey
	}
	key := multisigTypes.Key{
		ID:               multisig.KeyID(hex.EncodeToString(g.BitcoinPubkey)),
		Snapshot:         snapshot,
		PubKeys:          pubKeys,
		SigningThreshold: threshold,
		State:            multisig.Active,
	}
	return key
}

func NewCustodianGroup(name string, bitcoinPubkey []byte, quorum uint32, description string, custodians []*Custodian) *CustodianGroup {
	return &CustodianGroup{
		UID:           uuid.NewString(),
		Name:          name,
		BitcoinPubkey: bitcoinPubkey,
		Quorum:        quorum,
		Status:        Pending,
		Description:   description,
		Custodians:    custodians,
	}
}
