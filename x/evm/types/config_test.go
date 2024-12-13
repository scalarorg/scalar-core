package types_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/scalarorg/scalar-core/x/evm/types"
)

func TestUnmarshalEvmChain(t *testing.T) {
	content, err := os.ReadFile("testdata/evm.json")
	if err != nil {
		t.Fatal(err)
	}
	var result []types.EVMConfig
	if err := json.Unmarshal(content, &result); err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)
}
