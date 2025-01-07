package psbt

import (
	"fmt"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"

	btcRpcClient "github.com/btcsuite/btcd/rpcclient"
	"github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/scalar-core/sdk-utils/broadcast"
	"github.com/scalarorg/scalar-core/utils/slices"
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

func (mgr Mgr) sign(keyUID string, psbt covenantTypes.Psbt) (*covenant.TapScriptSigList, error) {
	if !mgr.validateKeyID(keyUID) {
		return nil, fmt.Errorf("invalid keyID")
	}

	privkey := mgr.privKey

	tapScriptSigs, err := vault.SignPsbtAndCollectSigs(
		psbt,
		privkey.Serialize(),
		vault.NetworkKindTestnet, // TODO: call to chain to get the network kind
	)
	if err != nil {
		return nil, err
	}

	return &covenant.TapScriptSigList{
		TapScriptSigs: slices.Map(tapScriptSigs, func(t vault.TapScriptSig) *covenant.TapScriptSig {
			keyXOnly := covenant.KeyXOnly(t.KeyXOnly)
			signature := covenant.Signature(t.Signature)
			leafHash := covenant.LeafHash(t.LeafHash)
			return &covenant.TapScriptSig{
				KeyXOnly:  &keyXOnly,
				LeafHash:  &leafHash,
				Signature: &signature,
			}
		}),
	}, nil
}

func (mgr Mgr) validateKeyID(keyID string) bool {
	// TODO: validate keyID
	return true
}
