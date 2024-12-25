package evm_test

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/scalarorg/bitcoin-vault/go-utils/encode"
	utilsTypes "github.com/scalarorg/bitcoin-vault/go-utils/types"
)

var (
	lockingScript, _ = hex.DecodeString("001450dceca158a9c872eb405d52293d351110572c9e")
	amount           = big.NewInt(10_000_000)
	feeOpts          = utilsTypes.MinimumFee.Bytes()

	destinationChain           = "bitcoin|4"
	destinationContractAddress = "0x0000000000000000000000000000000000000000"
)

// go test -timeout 10m -run ^TestTransferRemote$ github.com/scalarorg/scalar-core/client/rpc/evm -v -count=1

func TestTransferRemote(t *testing.T) {
	ctx := context.Background()
	// Prepare the input parameters

	payload, _, err := encode.CalculateTransferRemoteMetadataPayloadHash(amount.Uint64(), lockingScript, feeOpts[:])
	if err != nil {
		t.Fatalf("Error calculating unstaking payload hash: %v", err)
	}

	value := big.NewInt(1)

	tx, err := createAndSignTransaction(ctx, func() ([]byte, error) {
		return testSuite.ContractAbi.Pack("transferRemote",
			destinationChain,
			common.HexToAddress(destinationContractAddress),
			amount,
			payload,
		)
	}, value)
	if err != nil {
		t.Fatal(err)
	}

	sendAndWaitForTransaction(ctx, tx)
}
