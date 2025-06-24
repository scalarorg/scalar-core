package btc_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/scalarorg/scalar-core/utils/btc"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/stretchr/testify/require"
)

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestGetBlockVerboseTx$ github.com/scalarorg/scalar-core/vald/xchain/btc -v -count=1

func TestGetBlockVerboseTx(t *testing.T) {
	investigatedBlock := "000000000e3a834f586d716682aad42e58028ad6358f957c879327a03ef18d35"
	response, err := mockBtcClient.GetBlockVerboseTx(investigatedBlock)
	require.NoError(t, err)
	require.NotNil(t, response)
	block := response.Ok()
	require.NotNil(t, block)

	for i, tx := range block.Tx {
		fmt.Printf("Tx[%d]: %s\n", i, tx.Txid)
	}

	hashes := slices.Map(block.Tx, func(tx btcjson.TxRawResult) []byte {
		txHash, err := chainhash.NewHashFromStr(tx.Txid)
		require.NoError(t, err)
		return txHash.CloneBytes()
	})

	paths := GetMerklePath(hashes, 52)
	for i, p := range paths {
		fmt.Printf("MerklePath[%d]: %s\n", i, hex.EncodeToString(p))
	}

	root := btc.CalculateMerkleRoot(hashes)
	root2 := btc.GetMerkleRootFromPath(hashes[52], 52, paths, false)

	fmt.Printf("Root: %s\n", hex.EncodeToString(root))
	fmt.Printf("Root2: %s\n", hex.EncodeToString(root2))

}

func BuildMerkleTree(hashes [][]byte) [][][]byte {
	var levels [][][]byte
	current := hashes
	levels = append(levels, current)
	for len(current) > 1 {
		var nextLevel [][]byte
		for i := 0; i < len(current); i += 2 {
			left := current[i]
			var right []byte
			if i+1 < len(current) {
				right = current[i+1]
			} else {
				right = current[i]
			}
			nextLevel = append(nextLevel, btc.DoubleSha256(append(left, right...)))
		}
		levels = append(levels, nextLevel)
		current = nextLevel
	}
	return levels
}

func GetMerklePath(hashes [][]byte, txIndex int) [][]byte {
	levels := BuildMerkleTree(hashes)
	return GetMerklePathFromMerkleTree(levels, txIndex)
}

// GetMerklePath returns the Merkle path for txIndex using the precomputed tree.
func GetMerklePathFromMerkleTree(levels [][][]byte, txIndex int) [][]byte {
	var path [][]byte
	idx := txIndex
	for level := 0; level < len(levels)-1; level++ {
		levelHashes := levels[level]
		var siblingIdx int
		if idx%2 == 0 {
			siblingIdx = idx + 1
			if siblingIdx >= len(levelHashes) {
				siblingIdx = idx // duplicate if odd
			}
		} else {
			siblingIdx = idx - 1
		}
		path = append(path, levelHashes[siblingIdx])
		idx /= 2
	}
	return path
}
