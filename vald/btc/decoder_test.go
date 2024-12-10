package btc

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/wire"
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
