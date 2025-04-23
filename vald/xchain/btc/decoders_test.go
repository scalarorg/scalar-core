package btc

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	vault "github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	go_utils "github.com/scalarorg/bitcoin-vault/go-utils/types"
)

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestDecodeBridgeTx$ github.com/scalarorg/scalar-core/vald/xchain/btc -v -count=1
func TestDecodeBridgeTx(t *testing.T) {
	txHex, _ := hex.DecodeString("857e5c908ea91abca4612ad7cfb0b033e93a1efd281e843b8c376713c111a1a6")
	length := len(txHex)
	for i := 0; i < length/2; i++ {
		txHex[i], txHex[length-i-1] = txHex[length-i-1], txHex[i]
	}
	fmt.Println(hex.EncodeToString(txHex))

}
func TestDecodeStakingTransaction(t *testing.T) {
	txHex, _ := hex.DecodeString("020000000001011c668837803b7fd523f65bede28e3cc8ee2067d6e07733b3ea66925a2471d9f60200000000ffffffff030000000000000000416a3f5343414c4152030140706f6f6c73030100000000aa36a72ca3698a551a57169e73b0b2566a106ddec1b7b60cd77f1f0eb0f5477dc5f5d6e6a7200df94e4f120d260000000000002251207f815abf6dfd78423a708aa8db1c2c906eecac910c035132d342e4988a37b8d5052a010000000000160014526e83947f82190b5feb41f2131b025f4215396002473044022011578a153e478d27c86b153857c5c71818be8c8a3b8b959ac8fb82c1f3c421c9022039531d005c7c85d5069dd232f53229c91bb62821a641f2e5b10e1d2a191d68280121037083a18df56ab26ff0195687c5e99644597c550fa14f33df0ddbe466d08daf8b00000000")

	// Parse the transaction
	msgTx := wire.NewMsgTx(wire.TxVersion)
	err := msgTx.Deserialize(bytes.NewReader(txHex))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("msgTx: %+v", msgTx)

	t.Logf("msgTx.PrevOut: %+v", msgTx.TxIn[0].PreviousOutPoint)
	t.Logf("msgTx.PrevOut: %+v", msgTx.TxIn[0].PreviousOutPoint.String())
}

func TestDecodeEmbeddedDataTransaction(t *testing.T) {
	txHex, _ := hex.DecodeString("020000000001011c668837803b7fd523f65bede28e3cc8ee2067d6e07733b3ea66925a2471d9f60200000000ffffffff030000000000000000416a3f5343414c4152030140706f6f6c73030100000000aa36a72ca3698a551a57169e73b0b2566a106ddec1b7b60cd77f1f0eb0f5477dc5f5d6e6a7200df94e4f120d260000000000002251207f815abf6dfd78423a708aa8db1c2c906eecac910c035132d342e4988a37b8d5052a010000000000160014526e83947f82190b5feb41f2131b025f4215396002473044022011578a153e478d27c86b153857c5c71818be8c8a3b8b959ac8fb82c1f3c421c9022039531d005c7c85d5069dd232f53229c91bb62821a641f2e5b10e1d2a191d68280121037083a18df56ab26ff0195687c5e99644597c550fa14f33df0ddbe466d08daf8b00000000")

	// Parse the transaction
	msgTx := wire.NewMsgTx(wire.TxVersion)
	err := msgTx.Deserialize(bytes.NewReader(txHex))
	if err != nil {
		t.Fatal(err)
	}
	embeddedDataTxOut := msgTx.TxOut[EmbeddedDataOutputIndex]
	if embeddedDataTxOut == nil || embeddedDataTxOut.PkScript == nil || embeddedDataTxOut.PkScript[0] != txscript.OP_RETURN {
		t.Fatal("invalid op return")
	}

	output, err := vault.ParseVaultEmbeddedData(embeddedDataTxOut.PkScript)
	if err != nil || output == nil {
		t.Fatal("invalid op return data")
	}

	t.Logf("output: %+v", output)

	if output.TransactionType != go_utils.TransactionTypeLocking {
		t.Fatalf("invalid transaction type")
	}
}
