package exported

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils"
	"golang.org/x/crypto/sha3"

	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	snapshot "github.com/scalarorg/scalar-core/x/snapshot/exported"
)

const (
	DefaultCustodianName = "scalarv33"
)

func DefaultCustodian() *Custodian {
	custodian := &Custodian{
		Name:   DefaultCustodianName,
		Status: Custodian_Activated,
	}
	return custodian
}

func DefaultCustodianGroup() *CustodianGroup {
	return &CustodianGroup{
		UID:  CalculateUID(DefaultCustodianName),
		Name: DefaultCustodianName,
	}
}

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

	uid := CalculateUID(name)
	return &CustodianGroup{
		UID:           uid,
		Name:          name,
		BitcoinPubkey: bitcoinPubkey,
		Quorum:        quorum,
		Status:        Custodian_Pending,
		Description:   description,
		Custodians:    custodians,
	}
}

func CalculateUID(name string) [32]byte {
	//uid := sha256.Sum256([]byte(name))
	return sha3.Sum256([]byte(name))
}
