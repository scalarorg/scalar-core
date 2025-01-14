package btc_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/joho/godotenv"
)

var mockClient *rpcclient.Client

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestGetTransaction$ github.com/scalarorg/scalar-core/vald/xchain/btc -v -count=1
func TestGetTransaction(t *testing.T) {

	err := godotenv.Load("../../../.env.test")
	if err != nil {
		t.Fatalf("failed to load .env.test: %s", err)
	}

	host := os.Getenv("BTC_RPC_HOST")
	user := os.Getenv("BTC_RPC_USER")
	pass := os.Getenv("BTC_RPC_PASS")

	fmt.Println(host, user, pass)

	rpcConfig := rpcclient.ConnConfig{
		Host:                 "testnet4.btc.scalar.org",
		User:                 user,
		Pass:                 pass,
		Params:               "testnet3",
		DisableTLS:           true,
		DisableConnectOnNew:  true,
		DisableAutoReconnect: true,
		HTTPPostMode:         true,
	}

	rpcClient, error := rpcclient.New(&rpcConfig, nil)
	if error != nil {
		t.Fatalf("failed to create mock client: %s", error)
	}

	mockClient = rpcClient

	txid := "9640d60c9f53bdca7fe0520a276e5d7f7d33bd07773a2d7c8c462ac64480b5a8"

	chainHash, err := chainhash.NewHashFromStr(txid)
	if err != nil {
		t.Fatalf("failed to create chain hash: %s", err)
	}

	fmt.Printf("%+x\n", chainHash.CloneBytes())

	tx, err := rpcClient.GetRawTransactionVerbose(chainHash)
	if err != nil {
		t.Fatalf("failed to get transaction: %s", err)
	}

	blockHash, err := chainhash.NewHashFromStr(tx.BlockHash)
	if err != nil {
		t.Fatalf("failed to create block hash: %s", err)
	}

	block, err := rpcClient.GetBlockVerbose(blockHash)
	if err != nil {
		t.Fatalf("failed to get block: %s", err)
	}

	blockIndex := -1
	for i, blockTxId := range block.Tx {
		if blockTxId == txid {
			blockIndex = i
			break
		}
	}

	fmt.Println("blockIndex", blockIndex)

}
