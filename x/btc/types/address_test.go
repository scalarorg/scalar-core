package types_test

import (
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

func TestScriptPubKeyToAddress(t *testing.T) {
	tests := []struct {
		scriptPubKey string
		address      string
		chain        *chaincfg.Params
	}{
		{
			scriptPubKey: "51200f94f9d9c4c6e39cbef6c708b632173d8007b827936907176e19495c3e355c12",
			address:      "tb1pp720nkwycm3ee0hkcuytvvsh8kqq0wp8jd5sw9mwr9y4c034tsfqkjptxh",
			chain:        &chaincfg.TestNet3Params,
		},
		{
			scriptPubKey: "5120bea84c0a37dfb00534d247ed2bfdfdc94f173067307a6f1cda2f0837b684eaaa",
			address:      "tb1ph65ycz3hm7cq2dxjglkjhl0ae983wvr8xpax78x69uyr0d5ya24qlr2tja",
			chain:        &chaincfg.TestNet3Params,
		},
		{
			scriptPubKey: "001450dceca158a9c872eb405d52293d351110572c9e",
			address:      "tb1q2rwweg2c48y8966qt4fzj0f4zyg9wty7tykzwg",
			chain:        &chaincfg.TestNet3Params,
		},
		{
			scriptPubKey: "512095033d48b6029174ed3ba21390756c56e90c41eeeef41c172c81d1d09a167cda",
			address:      "bcrt1pj5pn6j9kq2ghfmfm5gfeqatv2m5scs0wam6pc9evs8gapxsk0ndqzagfze",
			chain:        &chaincfg.RegressionNetParams,
		},
	}

	for _, test := range tests {
		scriptPubKey, err := hex.DecodeString(test.scriptPubKey)
		if err != nil {
			t.Fatal(err)
		}

		address, err := types.ScriptPubKeyToAddress(scriptPubKey, test.chain)
		if err != nil {
			t.Fatal(err)
		}

		if address.String() != test.address {
			t.Fatalf("expected address %s, got %s", test.address, address.String())
		}
	}
}
