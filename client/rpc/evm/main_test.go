package evm_test

import (
	"crypto/ecdsa"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/scalarorg/scalar-core/client/rpc/evm"
)

type Suite struct {
	Client  *evm.Client
	Account struct {
		Address string
		PrivKey *ecdsa.PrivateKey
	}
	ProtocolAddress common.Address
}

var testSuite *Suite

func SetupTest() *Suite {
	rpcUrl := os.Getenv("ETHEREUM_RPC_URL")
	if rpcUrl == "" {
		panic("ETHEREUM_RPC_URL is not set")
	}

	privKey := os.Getenv("PRIVATE_KEY")
	if privKey == "" {
		panic("PRIVATE_KEY is not set")
	}

	protocolAddress := os.Getenv("PROTOCOL_ADDRESS")
	if protocolAddress == "" {
		panic("PROTOCOL_ADDRESS is not set")
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

	return &Suite{
		Client: mockClient,
		Account: struct {
			Address string
			PrivKey *ecdsa.PrivateKey
		}{
			Address: fromAddress.Hex(),
			PrivKey: privateKey,
		},
		ProtocolAddress: common.HexToAddress(protocolAddress),
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
