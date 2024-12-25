package grpc_client

import (
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/scalarorg/scalar-core/x/chains/types"
)

type QueryClientManager struct {
	client types.QueryServiceClient
}

var QueryManager *QueryClientManager

// For our use case, since we're primarily making queries and not doing complex connection management, the simpler version might be sufficient. The gRPC client's built-in features like multiple connections, keepalive, backoff, and retry should handle most edge cases automatically.
func new(ctx sdkClient.Context) (*QueryClientManager, error) {
	client := types.NewQueryServiceClient(ctx)

	return &QueryClientManager{
		client: client,
	}, nil
}

func InitQueryClientManager(ctx sdkClient.Context) error {
	m, err := new(ctx)
	if err != nil {
		return err
	}
	QueryManager = m
	return nil
}

func (m *QueryClientManager) GetClient() types.QueryServiceClient {
	if m == nil {
		panic("QueryClientManager is not initialized")
	}
	return m.client
}
