package rpc_test

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/scalarorg/relayers/pkg/clients/cosmos"
	"github.com/scalarorg/scalar-core/client/rpc"
)

const (
	chainNameBtcTestnet4 = "bitcoin-testnet5555"
)

var (
	CosmosNetworkConfig = cosmos.CosmosNetworkConfig{
		ChainID:       111,
		ID:            "scalar-testnet-1",
		Name:          "scalar-testnet-1",
		Denom:         "scalar",
		RPCUrl:        "http://localhost:26657",
		GasPrice:      0.001,
		LCDUrl:        "http://localhost:2317",
		WSUrl:         "ws://localhost:26657/websocket",
		MaxRetries:    3,
		RetryInterval: int64(1000),
		Mnemonic:      "4f24b74abf6b780d4b04836b44f33cc74d83d0a229d29c21ff072dbde9493958",
		BroadcastMode: "async",
	}
)

var (
	mockClientCtx     *client.Context
	mockQueryClient   *cosmos.QueryClient
	mockNetworkClient *cosmos.NetworkClient
)

func TestMain(m *testing.M) {
	txConfig := tx.NewTxConfig(rpc.GetProtoCodec(), []signing.SignMode{signing.SignMode_SIGN_MODE_DIRECT})
	var err error

	mockClientCtx, err = cosmos.CreateClientContextWithOptions(&CosmosNetworkConfig, cosmos.WithRpcClientCtx(CosmosNetworkConfig.RPCUrl))
	if err != nil {
		panic(err)
	}

	mockQueryClient = cosmos.NewQueryClient(mockClientCtx)

	privKey, addr, err := rpc.CreateAccountFromKey(CosmosNetworkConfig.Mnemonic)
	if err != nil {
		panic(err)
	}

	mockNetworkClient, err = cosmos.NewNetworkClientWithOptions(&CosmosNetworkConfig, mockQueryClient, txConfig, cosmos.WithRpcClient(mockClientCtx.Client), cosmos.WithQueryClient(mockQueryClient), cosmos.WithAccount(privKey, addr), cosmos.WithTxConfig(txConfig))
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
