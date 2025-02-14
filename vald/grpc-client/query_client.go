package grpc_client

import (
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	protocolTypes "github.com/scalarorg/scalar-core/x/protocol/types"
)

type QueryClientManager struct {
	chainsClient   chainsTypes.QueryServiceClient
	protocolClient protocolTypes.QueryClient
}

var queryManager *QueryClientManager

// For our use case, since we're primarily making queries and not doing complex connection management, the simpler version might be sufficient. The gRPC client's built-in features like multiple connections, keepalive, backoff, and retry should handle most edge cases automatically.
func new(ctx sdkClient.Context) (*QueryClientManager, error) {
	chainsClient := chainsTypes.NewQueryServiceClient(ctx)
	protocolClient := protocolTypes.NewQueryClient(ctx)

	return &QueryClientManager{
		chainsClient:   chainsClient,
		protocolClient: protocolClient,
	}, nil
}

func QueryManager() *QueryClientManager {
	if queryManager == nil {
		panic("QueryClientManager is not initialized")
	}
	return queryManager
}

func InitQueryClientManager(ctx sdkClient.Context) error {
	m, err := new(ctx)
	if err != nil {
		return err
	}
	queryManager = m
	return nil
}

func (m *QueryClientManager) GetChainsClient() chainsTypes.QueryServiceClient {
	if m == nil {
		panic("QueryClientManager is not initialized")
	}
	return m.chainsClient
}

func (m *QueryClientManager) GetProtocolClient() protocolTypes.QueryClient {
	if m == nil {
		panic("QueryClientManager is not initialized")
	}
	return m.protocolClient
}
