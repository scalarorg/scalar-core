package evm_test

import (
	"context"
	"math/big"
	"testing"
)

// go test -v -run TestFaucet ./client/rpc/evm/faucet_test.go
func TestFaucet(t *testing.T) {
	ctx := context.Background()
	// Prepare the input parameters
	amount := big.NewInt(90_000_000_000_000_000)

	tx, err := createAndSignTransaction(ctx, func() ([]byte, error) {
		return testSuite.ContractAbi.Pack("faucet", amount)
	}, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	sendAndWaitForTransaction(ctx, tx)
}
