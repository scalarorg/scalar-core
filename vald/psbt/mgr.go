package psbt

import (
	"fmt"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"

	btcRpcClient "github.com/btcsuite/btcd/rpcclient"
	"github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/bitcoin-vault/go-utils/types"
	"github.com/scalarorg/scalar-core/sdk-utils/broadcast"
	"github.com/scalarorg/scalar-core/utils/clog"
	covenant "github.com/scalarorg/scalar-core/x/covenant/exported"
	covenantTypes "github.com/scalarorg/scalar-core/x/covenant/types"
)

// Mgr represents an object that manages all communication with the psbt process
type Mgr struct {
	rpcs    map[chain.ChainInfoBytes]*btcRpcClient.Client
	ctx     sdkclient.Context
	valAddr sdk.ValAddress
	b       broadcast.Broadcaster
	privKey *secp256k1.PrivateKey
	pubKey  *secp256k1.PublicKey
}

// NewMgr is the constructor of mgr
func NewMgr(rpcs map[chain.ChainInfoBytes]*btcRpcClient.Client, ctx sdkclient.Context, valAddr sdk.ValAddress, b broadcast.Broadcaster, privKeyBytes []byte) *Mgr {

	if len(privKeyBytes) != 32 {
		panic("invalid private key length, got: " + fmt.Sprintf("%x", privKeyBytes))
	}

	privKey := secp256k1.PrivKeyFromBytes(privKeyBytes)

	pubKey := privKey.PubKey()

	return &Mgr{
		rpcs:    rpcs,
		ctx:     ctx,
		valAddr: valAddr,
		b:       b,
		privKey: privKey,
		pubKey:  pubKey,
	}
}

func (mgr Mgr) isParticipant(p sdk.ValAddress) bool {
	return mgr.valAddr.Equals(p)
}

func (mgr Mgr) sign(keyUID string, psbt covenantTypes.Psbt, networkKind types.NetworkKind) (*covenant.TapScriptSigsMap, error) {
	if !mgr.validateKeyID(keyUID) {
		return nil, fmt.Errorf("invalid keyID")
	}

	clog.Greenf("signing psbt with keyID: %s", keyUID)
	clog.Greenf("signing psbt with networkKind: %v", networkKind)
	clog.Greenf("signing psbt with privKey: %x", mgr.privKey.Serialize())
	clog.Greenf("signing psbt with PSBT: %x", psbt)

	tapScriptSigs, err := vault.SignPsbtAndCollectSigs(
		psbt.Bytes(),
		mgr.privKey.Serialize(),
		networkKind,
	)
	if err != nil {
		clog.Redf("[PsbtMgr] failed to sign PSBT: %s", err)
		return nil, err
	}

	clog.Greenf("signing psbt with tapScriptSigs: %+v\n", tapScriptSigs)

	mapOfTapScriptSigs := covenant.NewTapScriptSigsMapFromRaw(tapScriptSigs)

	clog.Greenf("signing psbt with mapOfTapScriptSigs: %+v\n", mapOfTapScriptSigs)

	return mapOfTapScriptSigs, nil
}

func (mgr Mgr) validateKeyID(keyID string) bool {
	// TODO: validate keyID
	return true
}
