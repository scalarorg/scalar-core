package types_test

// import (
// 	"encoding/json"
// 	"os"
// 	"testing"

// 	"github.com/scalarorg/scalar-core/x/chains/types"
// )

// func TestUnmarshalBtcChain(t *testing.T) {
// 	content, err := os.ReadFile("testdata/btc.json")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	var result []types.ChainConfig
// 	if err := json.Unmarshal(content, &result); err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Logf("%+v", result)
// }

// // func TestUnmarshalBtcChainText(t *testing.T) {
// // 	var chain types.ChainConfig
// // 	if err := chain.UnmarshalText([]byte("mainnet")); err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	t.Logf("%+v", chain)
// // }
