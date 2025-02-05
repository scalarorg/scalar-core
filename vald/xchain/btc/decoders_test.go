package btc

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	vault "github.com/scalarorg/bitcoin-vault/ffi/go-vault"
)

func TestDecodeStakingTransaction(t *testing.T) {
	txHex, _ := hex.DecodeString("02000000000101581e38836e48702b1634848c83514e7bdf4b6c774965a3318bbec9a0b79a65d20100000000fdffffff03102700000000000022512043a34440e92a00fdc7b6f9c38e12f99155e616c4b7e999c3373ba9717ddc8e4f00000000000000003d6a013504531801040100080000000000aa36a714b91e3a8ef862567026d6f376c9f3d6b814ca43371424a1db57fa3ecafcbad91d6ef068439aceeae09058cb0000000000002251200f94f9d9c4c6e39cbef6c708b632173d8007b827936907176e19495c3e355c120140ecc8999be320f19d6393c88b966473f020c6cf21fc851df6358aef5699a7f7c3d83cd82fb636895e0e8fba8c378cc684f802fd4b30f6d763ea8aa3a70c900c6600000000")

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
	txHex, _ := hex.DecodeString("02000000000101a8e432c21d1d275d5a69bdc048575567781474a2e3b896af22f9c34a55d0fbbe0200000000fdffffff030000000000000000416a3f5343414c41520201807472616e73030100000000aa36a7390e831349f34e8a7f323cb7350bf04a021d3c1272d3fa31e9fdd2f2ce195bdf9aba8393a717fe01e8030000000000002251209ec8dc148990200705b97d6ee201362936d7f9ce80926c7a3e1bdb33382aeda6a3d54b000000000016001450dceca158a9c872eb405d52293d351110572c9e0247304402202ec6fcd1dbb786379b47628ea6960bb7cee94c1e1b479a170d8895312d8d430602200da0345a845e63b9e0042b75330709bfc69dd6c58bd3febdf6dcc2f1f0505f730121022ae31ea8709aeda8194ba3e2f7e7e95e680e8b65135c8983c0a298d17bc5350a00000000")

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

}
