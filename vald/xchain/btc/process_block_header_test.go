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
	block86055 := "00000000000000054c3cc30a062c9d106c62bbc0596cf0ce221606d7a7a9c675"
	response, err := mockBtcClient.GetBlockVerboseTx(block86055)
	require.NoError(t, err)
	require.NotNil(t, response)
	block := response.Ok()
	require.NotNil(t, block)

	txs := []string{
		"787d1130d2992ca4954c9a64c6e6cadcc791579f825b8a26017b58607fbf0857",
		"10a67fbf7fc4493340d830dbdcd590d54021c51ab46811602f667e05d4d5013d",
		"d308f024ca222e3020a9ae7252964403f33bf384352b43850cf2a69d7636a3c9",
		"a24b88873b857f18da9799dc4a93f79079e96e0bbad5e969257cb99c9921f631",
		"fb4f11bd1df1bd16ef3bb57283421d88d67a16212e49baac1f15d84053f98aa3",
		"c66eb6352d5eb22ec5f3ed9aab4d839dd06dbbed247cb7ad92f58a2a7f0c63c0",
		"acedf5986ff7ba1406f258df07d197063b4b5b76d7c6c81e006857b349bcfd65",
		"b2b5473e14fd244202d547a26fcc76b9815548ffb7be2d71c527216229389e09",
		"badef8e66d1c945d6d131bc03104ef38304ce0e401e7e88872506f10e645430f",
		"f8859081958a8881c8c00ccfab2d1c1c9f56bfa3ceea32deba41aff6ef238320",
		"e339a921feb378f6464815ce37256deaf36754c777d089fc615d6dc24d41a027",
		"fe65fba4f1a143ab73daa446ed729463714514494d71e0cb0a906aa7c7b29d54",
		"901632bacb37ebc1e42e78d278b1059abe77076e786e202c16af25932e23a478",
		"280d793351aafe695d44884589ffb9b79a133563659be6e4d63a97d4e7754a8a",
		"2dbcc763f98a3ec61bb9ac9091a212860d807f79652e2ccb308b62e30092c195",
		"9a9682f1a1ff0252d2e35d4080278d27b82d67e9818ac24b3e5096ace8905fb4",
	}

	for i, tx := range block.Tx {
		fmt.Printf("Tx[%d]: %s\n", i, tx.Txid)
		require.Equal(t, txs[i], tx.Txid)
	}

	merkleRoot := btc.CalculateMerkleRoot(slices.Map(block.Tx, func(tx btcjson.TxRawResult) []byte {
		txHash, err := chainhash.NewHashFromStr(tx.Txid)
		require.NoError(t, err)
		return txHash.CloneBytes()
	}))

	require.Equal(t, block.MerkleRoot, hex.EncodeToString(btc.ReverseBytes(merkleRoot)))
	fmt.Printf("MerkleRoot: %s\n", block.MerkleRoot)

	path1, err := hex.DecodeString("e339a921feb378f6464815ce37256deaf36754c777d089fc615d6dc24d41a027")
	require.NoError(t, err)
	path2, err := hex.DecodeString("7770206e3cbb3edc5ccba0529dfd8e60db98be4ba76bdb9ad415991581eb7282")
	require.NoError(t, err)
	path3, err := hex.DecodeString("9723833a5344bdc2b23755874c98869eaba9ff58618af309c3f92b0c56c6fb08")
	require.NoError(t, err)
	path4, err := hex.DecodeString("8ca8c2b2a98d8cbd99d6ae4188aebe517dbb6991612ae97f26c9fa37573dcae3")
	require.NoError(t, err)
	root := btc.GetMerkleRootFromPath(merkleRoot, 11, [][]byte{path1, path2, path3, path4}, true)

	fmt.Printf("Root: %s\n", hex.EncodeToString(root))
}

// "tx": [
//     "787d1130d2992ca4954c9a64c6e6cadcc791579f825b8a26017b58607fbf0857",
//     "10a67fbf7fc4493340d830dbdcd590d54021c51ab46811602f667e05d4d5013d",
//     "d308f024ca222e3020a9ae7252964403f33bf384352b43850cf2a69d7636a3c9",
//     "a24b88873b857f18da9799dc4a93f79079e96e0bbad5e969257cb99c9921f631",
//     "fb4f11bd1df1bd16ef3bb57283421d88d67a16212e49baac1f15d84053f98aa3",
//     "c66eb6352d5eb22ec5f3ed9aab4d839dd06dbbed247cb7ad92f58a2a7f0c63c0",
//     "acedf5986ff7ba1406f258df07d197063b4b5b76d7c6c81e006857b349bcfd65",
//     "b2b5473e14fd244202d547a26fcc76b9815548ffb7be2d71c527216229389e09",
//     "badef8e66d1c945d6d131bc03104ef38304ce0e401e7e88872506f10e645430f",
//     "f8859081958a8881c8c00ccfab2d1c1c9f56bfa3ceea32deba41aff6ef238320",
//     "e339a921feb378f6464815ce37256deaf36754c777d089fc615d6dc24d41a027",
//     "fe65fba4f1a143ab73daa446ed729463714514494d71e0cb0a906aa7c7b29d54",
//     "901632bacb37ebc1e42e78d278b1059abe77076e786e202c16af25932e23a478",
//     "280d793351aafe695d44884589ffb9b79a133563659be6e4d63a97d4e7754a8a",
//     "2dbcc763f98a3ec61bb9ac9091a212860d807f79652e2ccb308b62e30092c195",
//     "9a9682f1a1ff0252d2e35d4080278d27b82d67e9818ac24b3e5096ace8905fb4"
//   ]
