package evm_test

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var unstakeAbi, _ = abi.JSON(strings.NewReader(`[{
	"type": "function",
	"name": "unstake",
	"inputs": [
		{ "name": "_destinationChain", "type": "string" },
		{ "name": "_destinationAddress", "type": "string" },
		{ "name": "_amount", "type": "uint256" },
		{ "name": "_encodedPayload", "type": "bytes" }
	],
	"outputs": [],
	"stateMutability": "nonpayable"
}]`))

func TestUnstake(t *testing.T) {
	ctx := context.Background()
	// Prepare the input parameters
	destinationChain := "bitcoin|4"
	destinationAddress := "tb1q2rwweg2c48y8966qt4fzj0f4zyg9wty7tykzwg"
	amount := big.NewInt(10_000_000)
	encodedPayload := []byte{64, 61, 76, 69, 64}

	// Pack the input data
	input, err := unstakeAbi.Pack("unstake",
		destinationChain,
		destinationAddress,
		amount,
		encodedPayload,
	)
	if err != nil {
		t.Fatal(err)
	}

	// Create transaction data
	nonce, err := testSuite.Client.PendingNonceAt(ctx, common.HexToAddress(testSuite.Account.Address))
	if err != nil {
		t.Fatal(err)
	}

	contractAddress := testSuite.ProtocolAddress

	value := big.NewInt(0) // in wei (0 eth)

	gasLimit, err := testSuite.Client.EstimateGas(ctx, ethereum.CallMsg{
		To:   &contractAddress,
		Data: input,
	})
	if err != nil {
		t.Fatal(err)
	}

	gasPrice, err := testSuite.Client.SuggestGasPrice(ctx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("nonce", nonce)
	fmt.Println("contractAddress", contractAddress)
	fmt.Println("value", value)
	fmt.Println("gasLimit", gasLimit)
	fmt.Println("gasPrice", gasPrice)
	fmt.Println("input", input)

	// Now you can use this data to send a transaction
	tx := types.NewTransaction(
		nonce,
		contractAddress,
		value,
		gasLimit,
		gasPrice,
		input,
	)

	chainID, err := testSuite.Client.NetworkID(ctx)
	if err != nil {
		t.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), testSuite.Account.PrivKey)
	if err != nil {
		t.Fatal(err)
	}

	err = testSuite.Client.SendTransaction(ctx, signedTx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

	receipt, err := testSuite.Client.TransactionReceipt(ctx, signedTx.Hash())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("receipt: %+v", receipt)
}
