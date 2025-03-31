package btc_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/scalarorg/scalar-core/vald/config"
	"github.com/scalarorg/scalar-core/vald/xchain/btc"
	"github.com/stretchr/testify/require"
)

var (
	mockClient    *rpcclient.Client
	mockBtcClient *btc.BtcClient
)

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestGetTransaction$ github.com/scalarorg/scalar-core/vald/xchain/btc -v -count=1
func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		fmt.Println("failed to load .env.test", err)
	}

	host := os.Getenv("BTC_RPC_HOST")
	user := os.Getenv("BTC_RPC_USER")
	pass := os.Getenv("BTC_RPC_PASS")
	tag := os.Getenv("BTC_SCALAR_TAG")
	version, err := strconv.ParseUint(os.Getenv("BTC_SCALAR_VERSION"), 10, 64)
	if err != nil {
		fmt.Println("failed to parse BTC_SCALAR_VERSION", err)
	}
	fmt.Println(host, user, pass, tag, version)

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

	btcConfig := config.BTCConfig{
		ID:           "bitcoin|4",
		Chain:        "testnet4",
		RPCHost:      host,
		RPCUser:      user,
		RPCPass:      pass,
		Tag:          tag,
		Version:      version,
		DisableTLS:   true,
		HttpPostMode: true,
	}
	rpcClient, err := rpcclient.New(&rpcConfig, nil)
	if err != nil {
		fmt.Println("failed to create mock rpc client", err)
	}
	mockClient = rpcClient
	fmt.Println("btcConfig", btcConfig)
	commonClient, err := btc.NewClient(&btcConfig)
	if err != nil {
		fmt.Println("failed to create mock btc client", err)
	} else {
		mockBtcClient = commonClient.(*btc.BtcClient)
	}
	os.Exit(m.Run())
}

// Try to figure out how long it takes to get the tx receipts if finalized
// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestGetTxReceiptsIfFinalized$ github.com/scalarorg/scalar-core/vald/xchain/btc -v -count=1
func TestGetTxReceiptsIfFinalized(t *testing.T) {
	txids := []common.Hash{
		common.HexToHash("59aca882dc49738b711ec709153c7288d62c97df67b98ebf7e224421d31ae94d"),
		common.HexToHash("b3ba45ad5063307a2765afab98d580d1c24f7216ff9f527188078325c432248b"),
		common.HexToHash("acaeef8d5bdc644a508e529e589ed49093f0521f59c5cb71971c8a42bb0aa841"),
		common.HexToHash("224598e880a71eed260e39ff04a0a409289872ebb07e10bb70b6a33e29532c05"),
		common.HexToHash("6e658115968fe3bd70e2f0a7d84a8a10c6c4dbc30e45bedc26ddf7e2dfa25a98"),
		common.HexToHash("1a14811a706af78f1cdc731804d7f7bc8dea5830aecd93068a8517aaeb6caea3"),
		common.HexToHash("08082505fbe06a3bd085e889853dc68216da670848eedfd8ca4ef904f4bb38cc"),
		common.HexToHash("8bac55022e1b58da39a175acc77de056c1ba4975b5a5a777e16c826946d7040c"),
		common.HexToHash("26a694ebd3dcec1a8a23abf9611bd1dc7eae4d5bdfbb6ed92d97be8c2374c3b9"),
		common.HexToHash("006681af1860770e95083795754a67e17df863a72f22fc8716b0fa9934a4e433"),
	}
	require.NotNil(t, mockBtcClient)
	start := time.Now()
	btcResults, err := mockBtcClient.GetTxReceiptsIfFinalized(txids, 12)
	require.NoError(t, err)
	require.Equal(t, len(btcResults), len(txids))
	for _, result := range btcResults {
		require.NoError(t, result.Err())
	}
	elapsed := time.Since(start)
	fmt.Println("Time taken:", elapsed)
	require.Less(t, elapsed, 3*time.Second)
}

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
