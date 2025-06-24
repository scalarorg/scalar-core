package btc_test

import (
	"encoding/hex"
	"testing"

	"github.com/scalarorg/scalar-core/utils/btc"
	"github.com/stretchr/testify/assert"
)

func TestGetMerkleRootFromPath(t *testing.T) {
	// Test case 1: Basic test with known values
	t.Run("basic test", func(t *testing.T) {
		txid, _ := hex.DecodeString("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")
		txIndex := uint64(0)
		merklePath := [][]byte{
			[]byte("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"),
			[]byte("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"),
		}
		blockMerkleRoot := []byte("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")

		merkleRoot := btc.GetMerkleRootFromPath(txid, txIndex, merklePath, true)
		assert.Equal(t, blockMerkleRoot, merkleRoot)
	})

	// Test case 2: Real Bitcoin data from the input
	t.Run("real bitcoin data", func(t *testing.T) {
		// Parse the Merkle path from the input data
		merkleHashes := []string{
			"58938104f16ff39dc514aabf90ab57e0b1ac316e5e8a5418e5f62fba11fa2d29",
			"e7aa30270019f140b89d7461c2cdd3ce97a16388fc9c437829df6bdae223c69a",
			"5967973b59be8e243a71fa53690aad1f52cdf4c4a97dfea828eb3c3779107c0a",
			"a5a4e23726f95048e83a577627af4c9a99c294deb42b62d35365e2e8b77cd547",
			"5ec0b6dfe4ae6ff556c5960497344886bb820a0049a769f07b8eed841e3f9e2c",
			"c068ab902a16456ed3c954a2c65fced9dae84836c7eeb5c6c9728b7f5ff8b69a",
			"2993d208a4d59959577ef29c42bdd7013a05f2195cca4afc9334f8f7b3b7e02f",
			"c66900615f48f09321daf6569e4e48deb0c3f70ae14dd038708452e2a2a3a34e",
			"7c96cfe1b5eba2e82ebeb857279a5456a4128f62f0c70657a58586639ec079b1",
			"fd7a0408c21fe05080652b05321e159c36e9c5403840bb1840739a212be9cc3a",
			"1b890ea19c704743dd5b8beecef35080499bed6d21a811ebfccb9260fedc2ac8",
			"184f97eec1a4a5079c02e50aea156cab716b3e97f1ec1b02ee43d1bc0f62717c",
			"e6155e1fa015736ce678c5f8c015c9aea554e590a877fc08154ea9601bcfc2d8",
		}

		// Convert hex strings to byte arrays
		merklePath := make([][]byte, len(merkleHashes))
		for i, hash := range merkleHashes {
			merklePath[i], _ = hex.DecodeString(hash)
		}

		// Expected merkle root from the input data
		expectedRoot, _ := hex.DecodeString("b7f2da6c0bae9acb9158f55d8c1b845df9a6d9ff9987b452d1fbbfff2592b4e7")

		// We need a transaction ID and index - using placeholder for now
		// In a real scenario, these would come from the actual transaction
		txid, _ := hex.DecodeString("04fde533acb663bc0b72f049e8c9d5f53ecf53da7b7d8cf15214558b0a0d2f29")
		txIndex := uint64(61) // Position from the input data

		merkleRoot := btc.GetMerkleRootFromPath(txid, txIndex, merklePath, true)
		assert.Equal(t, expectedRoot, merkleRoot)
	})

	// Test case 3: Test with reverse=false
	t.Run("without reverse", func(t *testing.T) {
		txid, _ := hex.DecodeString("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")
		txIndex := uint64(0)
		merklePath := [][]byte{
			[]byte("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"),
		}

		merkleRoot := btc.GetMerkleRootFromPath(txid, txIndex, merklePath, false)
		assert.NotNil(t, merkleRoot)
		assert.Len(t, merkleRoot, 32)
	})

	// Test case 4: Test with odd transaction index
	t.Run("odd transaction index", func(t *testing.T) {
		txid, _ := hex.DecodeString("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")
		txIndex := uint64(1)
		merklePath := [][]byte{
			[]byte("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"),
		}

		merkleRoot := btc.GetMerkleRootFromPath(txid, txIndex, merklePath, true)
		assert.NotNil(t, merkleRoot)
		assert.Len(t, merkleRoot, 32)
	})

	// Test case 5: Test with empty merkle path
	t.Run("empty merkle path", func(t *testing.T) {
		txid, _ := hex.DecodeString("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")
		txIndex := uint64(0)
		merklePath := [][]byte{}

		merkleRoot := btc.GetMerkleRootFromPath(txid, txIndex, merklePath, true)
		assert.Equal(t, txid, merkleRoot)
	})
}
