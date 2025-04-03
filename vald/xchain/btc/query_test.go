package btc_test

import (
	"testing"

	"github.com/scalarorg/scalar-core/vald/xchain/btc"
	"gotest.tools/assert"
)

type UtxoStatus struct {
	Confirmed   bool   `json:"confirmed"`
	BlockHeight uint64 `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   uint64 `json:"block_time"`
}

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestSortUTXOsByBlockHeight$ github.com/scalarorg/scalar-core/vald/xchain/btc -v -count=1

func TestSortUTXOsByBlockHeight(t *testing.T) {
	utxos := []btc.MempoolUtxo{
		{
			Txid: "12dc",
			Vout: 0,
			Status: UtxoStatus{
				Confirmed:   true,
				BlockHeight: 2,
				BlockHash:   "block_hash",
				BlockTime:   100000000,
			},
			Value: 100000000,
		},
		{
			Txid: "a132",
			Vout: 0,
			Status: UtxoStatus{
				Confirmed:   true,
				BlockHeight: 2,
				BlockHash:   "block_hash",
				BlockTime:   100000000,
			},
			Value: 100000000,
		},
		{
			Txid: "b6h7",
			Vout: 0,
			Status: UtxoStatus{
				Confirmed:   true,
				BlockHeight: 1,
				BlockHash:   "block_hash",
				BlockTime:   100000000,
			},
		},
	}
	sortedUTXOs := btc.SortUTXOsByBlockHeight(utxos)
	assert.Equal(t, sortedUTXOs[0].Txid, "b6h7")
	assert.Equal(t, sortedUTXOs[1].Txid, "12dc")
	assert.Equal(t, sortedUTXOs[2].Txid, "a132")
}
