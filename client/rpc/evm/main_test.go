package evm_test

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/scalarorg/scalar-core/client/rpc/evm"
	"github.com/scalarorg/scalar-core/utils/clog"
)

type Suite struct {
	Client  *evm.Client
	Account struct {
		Address string
		PrivKey *ecdsa.PrivateKey
	}
	ContractAddress common.Address
	ContractAbi     abi.ABI
}

var testSuite *Suite

func SetupTest() *Suite {
	rpcUrl := os.Getenv("SEPOLIA_RPC_URL")
	if rpcUrl == "" {
		panic("SEPOLIA_RPC_URL is not set")
	}

	privKey := os.Getenv("PRIVATE_KEY")
	if privKey == "" {
		panic("PRIVATE_KEY is not set")
	}

	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	if contractAddress == "" {
		panic("CONTRACT_ADDRESS is not set")
	}

	mockClient, err := evm.NewClient(rpcUrl)
	if err != nil {
		panic("Error creating client: " + err.Error())
	}

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		panic("Error converting private key: " + err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	abiJson, err := os.ReadFile("../data/abi.json")
	if err != nil {
		panic("Error reading abi.json: " + err.Error())
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(abiJson)))
	if err != nil {
		panic("Error parsing abi.json: " + err.Error())
	}

	return &Suite{
		Client: mockClient,
		Account: struct {
			Address string
			PrivKey *ecdsa.PrivateKey
		}{
			Address: fromAddress.Hex(),
			PrivKey: privateKey,
		},
		ContractAddress: common.HexToAddress(contractAddress),
		ContractAbi:     contractAbi,
	}
}

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env.test")
	if err != nil {
		panic("Error loading .env file")
	}

	testSuite = SetupTest()

	os.Exit(m.Run())
}

func createAndSignTransaction(ctx context.Context, createInput func() ([]byte, error), value *big.Int) (*types.Transaction, error) {
	input, err := createInput()
	if err != nil {
		return nil, fmt.Errorf("create input: %w", err)
	}

	nonce, err := testSuite.Client.PendingNonceAt(ctx, common.HexToAddress(testSuite.Account.Address))
	if err != nil {
		return nil, err
	}

	gasLimit, err := testSuite.Client.EstimateGas(ctx, ethereum.CallMsg{
		From:  common.HexToAddress(testSuite.Account.Address),
		To:    &testSuite.ContractAddress,
		Data:  input,
		Value: value,
	})
	if err != nil {
		return nil, fmt.Errorf("estimate gas: %w", err)
	}

	gasPrice, err := testSuite.Client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("suggest gas price: %w", err)
	}

	tx := types.NewTransaction(
		nonce,
		testSuite.ContractAddress,
		value,
		gasLimit,
		gasPrice,
		input,
	)

	chainID, err := testSuite.Client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	return types.SignTx(tx, types.NewEIP155Signer(chainID), testSuite.Account.PrivKey)
}

func sendAndWaitForTransaction(ctx context.Context, signedTx *types.Transaction) (*types.Receipt, error) {
	if err := testSuite.Client.SendTransaction(ctx, signedTx); err != nil {
		return nil, err
	}

	clog.Redf("tx sent: %s\n", signedTx.Hash().Hex())

	receipt, err := waitForReceipt(ctx, testSuite.Client.Client, signedTx.Hash())
	if err != nil {
		return nil, err
	}

	clog.Greenf("receipt found: %+v\n", receipt)

	return receipt, nil
}

func waitForReceipt(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	clog.Greenf("waiting for receipt: %s\n", txHash.Hex())
	for {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Second):
			// Wait 1 second before trying again
		}
	}
}
