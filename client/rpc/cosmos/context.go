package cosmos

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/client/rpc/codec"
	"github.com/scalarorg/scalar-core/client/rpc/config"
)

// var _ grpc.ClientConn = &client.Context{}

func CreateClientContext() (*client.Context, error) {
	clientCtx := client.Context{
		ChainID: config.GlobalConfig.ID,
	}
	if config.GlobalConfig.RPCUrl != "" {
		log.Info().Msgf("Create rpcClient using RPC URL: %s", config.GlobalConfig.RPCUrl)
		clientCtx = clientCtx.WithNodeURI(config.GlobalConfig.RPCUrl)
		rpcClient, err := client.NewClientFromNode(config.GlobalConfig.RPCUrl)
		if err != nil {
			return nil, fmt.Errorf("failed to create RPC client: %w", err)
		}
		clientCtx = clientCtx.WithClient(rpcClient)
	}
	if config.GlobalConfig.Mnemonic != "" {
		_, addr, err := CreateAccountFromMnemonic(config.GlobalConfig.Mnemonic, "")
		if err != nil {
			return nil, fmt.Errorf("failed to create account from mnemonic: %w", err)
		}
		clientCtx = clientCtx.WithFromAddress(addr)
	}
	clientCtx = clientCtx.WithCodec(codec.GetProtoCodec())
	clientCtx = clientCtx.WithOutputFormat("json")
	return &clientCtx, nil
}

type ClientContextOption func(*client.Context) error

func WithRpcClientCtx(rpcUrl string) ClientContextOption {
	return func(c *client.Context) error {
		rpcClient, err := client.NewClientFromNode(rpcUrl)
		if err != nil {
			return fmt.Errorf("failed to create RPC client: %w", err)
		}
		clientCtx := *c
		clientCtx = clientCtx.WithNodeURI(rpcUrl)
		clientCtx = clientCtx.WithClient(rpcClient)
		*c = clientCtx
		return nil
	}
}

func CreateClientContextWithOptions(opts ...ClientContextOption) (*client.Context, error) {
	clientCtx := client.Context{
		ChainID: config.GlobalConfig.ID,
	}
	for _, opt := range opts {
		err := opt(&clientCtx)
		if err != nil {
			return nil, err
		}
	}
	clientCtx = clientCtx.WithCodec(codec.GetProtoCodec())
	clientCtx = clientCtx.WithOutputFormat("json")
	return &clientCtx, nil
}
