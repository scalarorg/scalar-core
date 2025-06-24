package btc_test

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/chains/types"
	votes "github.com/scalarorg/scalar-core/x/vote/exported"
	"github.com/stretchr/testify/require"
)

func reverseBytes(data []byte) []byte {
	result := make([]byte, len(data))
	for i, b := range data {
		result[len(data)-1-i] = b
	}
	return result
}

func doubleSha256(data []byte) []byte {
	firstHash := sha256.Sum256(data)
	secondHash := sha256.Sum256(firstHash[:])
	return secondHash[:]
}
func buildMerkleTree(txHashes [][]byte) []byte {
	if len(txHashes) == 0 {
		return nil
	}
	if len(txHashes) == 1 {
		return txHashes[0]
	}
	// Double SHA256 each transaction hash
	hashes := make([][]byte, len(txHashes))
	for i, txHash := range txHashes {
		hashes[i] = doubleSha256(txHash)
	}
	for len(hashes) > 1 {
		nextLevel := make([][]byte, 0, (len(hashes)+1)/2)
		for i := 0; i < len(hashes); i += 2 {
			if i+1 < len(hashes) {
				combined := append(hashes[i], hashes[i+1]...)
				nextLevel = append(nextLevel, doubleSha256(combined))
			} else {
				combined := append(hashes[i], hashes[i]...)
				nextLevel = append(nextLevel, doubleSha256(combined))
			}
		}
		hashes = nextLevel
	}
	return hashes[0]
}

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestProcessNewBlockConfirmation$ github.com/scalarorg/scalar-core/vald/xchain/btc -v -count=1
func TestProcessNewBlockConfirmation(t *testing.T) {

	prevBlockHash := exported.Hash(common.HexToHash("00000000000000030642a17ebbc7ebb4215b50618983087427c0d2632fd1eeaa").Bytes())

	vote, err := mockBtcClient.ProcessNewBlockConfirmation(&types.ConfirmNewBlockStarted{
		Chain: "btc",
		PollParticipants: votes.PollParticipants{
			PollID: "1",
			Participants: []sdk.ValAddress{
				sdk.ValAddress(common.HexToAddress("0x1").Bytes()),
			},
		},
		BlockHash:          exported.Hash(common.HexToHash("0000000009cc1e47605fdd3ecd7bfedbed161ae82afd555e4cd760c8d9eecb3d").Bytes()),
		PreviousBlockHash:  &prevBlockHash,
		ConfirmationHeight: 2,
	}, []byte{})

	require.NoError(t, err)
	require.NotNil(t, vote)
}

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestGetBlockVerboseTx$ github.com/scalarorg/scalar-core/vald/xchain/btc -v -count=1

func TestGetBlockVerboseTx(t *testing.T) {
	block86055 := "00000000000000054c3cc30a062c9d106c62bbc0596cf0ce221606d7a7a9c675"
	response, err := mockBtcClient.GetBlockVerboseTx(block86055)
	require.NoError(t, err)
	require.NotNil(t, response)
	block := response.Ok()
	require.NotNil(t, block)
	fmt.Printf("Block height: %d\n", block.Height)
	fmt.Printf("Block hash: %s\n", block.Hash)
	fmt.Printf("Block previous hash: %s\n", block.PreviousHash)

	txHashes := [][]byte{}

	for i, tx := range block.Tx {
		txHex := tx.Hex
		// Convert hex string into raw bytes
		txBytes, err := hex.DecodeString(txHex)
		if err != nil {
			panic(err)
		}

		// Calculate double SHA256
		txHash := doubleSha256(txBytes)
		fmt.Printf("Tx[%d]: %s\n", i, txHex)
		fmt.Printf("Tx[%d]: %s; calculated hash: %s\n", i, tx.Hash, hex.EncodeToString(txHash))

		txHashBytes, err := hex.DecodeString(tx.Hash)
		require.NoError(t, err)
		txHashBytes = reverseBytes(txHashBytes)
		txHashes = append(txHashes, txHashBytes)
	}
	merkleRoot := buildMerkleTree(txHashes)
	fmt.Printf("MerkleRoot: %s\n", block.MerkleRoot)
	fmt.Printf("Calculated MerkleRoot: %s\n", hex.EncodeToString(merkleRoot))
}
