package types_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/scalarorg/scalar-core/x/btc/types"
)

func TestUnmarshalBtcChain(t *testing.T) {
	content, err := os.ReadFile("testdata/btc.json")
	if err != nil {
		t.Fatal(err)
	}
	var result []types.BTCConfig
	if err := json.Unmarshal(content, &result); err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)
}
