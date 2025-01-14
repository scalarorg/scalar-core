package psbt_test

import (
	"encoding/hex"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	psbtMgr "github.com/scalarorg/scalar-core/vald/psbt"
	covenantTypes "github.com/scalarorg/scalar-core/x/covenant/types"
)

var mockMgr = &psbtMgr.Mgr{}

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestValidatePsbt$ github.com/scalarorg/scalar-core/vald/psbt -v -count=1
func TestValidatePsbt(t *testing.T) {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Fatal("Error loading .env.test file")
	}

	// rpcConfig := rpcclient.ConnConfig{
	// 	Host: os.Getenv("BTC_RPC_HOST"),
	// 	User: os.Getenv("BTC_RPC_USER"),
	// 	Pass: os.Getenv("BTC_RPC_PASS"),
	// }

	// rpcClient, error := rpcclient.New(&rpcConfig, nil)
	// if error != nil {
	// 	t.Fatalf("failed to create mock client: %s", error)
	// }

	psbtBytes, err := hex.DecodeString(os.Getenv("MOCK_PSBT"))
	if err != nil {
		t.Fatalf("failed to decode PSBT: %s", err)
	}

	err = mockMgr.ValidatePsbt(nil, covenantTypes.Psbt(psbtBytes))
	if err != nil {
		t.Fatalf("failed to validate PSBT: %s", err)
	}
}
