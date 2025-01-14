package psbt

import (
	"bytes"
	"fmt"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"

	"github.com/btcsuite/btcd/btcutil/psbt"
	btcRpcClient "github.com/btcsuite/btcd/rpcclient"
	"github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/scalar-core/sdk-utils/broadcast"
	"github.com/scalarorg/scalar-core/utils/clog"
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

func (mgr Mgr) sign(keyUID string, psbt covenantTypes.Psbt, networkKind vault.NetworkKind) (*covenant.TapScriptSigList, error) {
	if !mgr.validateKeyID(keyUID) {
		return nil, fmt.Errorf("invalid keyID")
	}

	tapScriptSigs, err := vault.SignPsbtAndCollectSigs(
		psbt,
		mgr.privKey.Serialize(),
		networkKind,
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

func (mgr Mgr) validatePsbt(p covenantTypes.Psbt, chainInfoBytes chain.ChainInfoBytes) error {
	psbt, err := psbt.NewFromRawBytes(bytes.NewReader(p), false)
	if err != nil {
		return fmt.Errorf("failed to parse PSBT: %w", err)
	}

	if len(psbt.Inputs) == 0 {
		return fmt.Errorf("PSBT has no inputs")
	}

	clog.Redf("psbt.Inputs: %+v", psbt.Inputs)
	clog.Redf("psbt.UnsignedTx.TxIn: %+v", psbt.UnsignedTx.TxIn)

	// TODO: use go routines to validate each input

	// For each PSBT input, verify it exists as a UTXO
	for i, input := range psbt.Inputs {
		txIn := psbt.UnsignedTx.TxIn[i]

		// Create outpoint hash string
		txHash := txIn.PreviousOutPoint.Hash.String()

		// Get the transaction from the Bitcoin node
		tx, err := mgr.rpcs[chainInfoBytes].GetRawTransaction(&txIn.PreviousOutPoint.Hash)
		if err != nil {
			return fmt.Errorf("input %d: failed to fetch transaction %s: %w", i, txHash, err)
		}

		// Check if the referenced output index exists
		if int(txIn.PreviousOutPoint.Index) >= len(tx.MsgTx().TxOut) {
			return fmt.Errorf("input %d: invalid output index %d for transaction %s",
				i, txIn.PreviousOutPoint.Index, txHash)
		}

		// Verify the UTXO is unspent using GetTxOut
		txOut, err := mgr.rpcs[chainInfoBytes].GetTxOut(&txIn.PreviousOutPoint.Hash, txIn.PreviousOutPoint.Index, true)
		if err != nil {
			return fmt.Errorf("input %d: failed to check UTXO status: %w", i, err)
		}
		if txOut == nil {
			return fmt.Errorf("input %d: UTXO is already spent", i)
		}

		// Verify witness utxo if present
		if input.WitnessUtxo != nil {
			outputIndex := txIn.PreviousOutPoint.Index
			if input.WitnessUtxo.Value != tx.MsgTx().TxOut[outputIndex].Value {
				return fmt.Errorf("input %d: witness UTXO amount mismatch", i)
			}
			if !bytes.Equal(input.WitnessUtxo.PkScript, tx.MsgTx().TxOut[outputIndex].PkScript) {
				return fmt.Errorf("input %d: witness UTXO script mismatch", i)
			}
		}
	}

	return nil
}
