package psbt_test

import (
	"encoding/hex"
	"os"
	"testing"

	covenant "github.com/scalarorg/scalar-core/x/covenant/exported"
)

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestValidatePsbt$ github.com/scalarorg/scalar-core/vald/psbt -v -count=1
func TestValidatePsbt(t *testing.T) {

	psbtBytes, err := hex.DecodeString(os.Getenv("MOCK_PSBT"))
	if err != nil {
		t.Fatalf("failed to decode PSBT: %s", err)
	}

	err = mockMgr.ValidatePsbt(nil, covenant.Psbt(psbtBytes))
	if err != nil {
		t.Fatalf("failed to validate PSBT: %s", err)
	}
}
