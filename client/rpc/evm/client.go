package evm

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/scalarorg/scalar-core/utils/clog"
)

type Client struct {
	*ethclient.Client
}

func NewClient(url string) (*Client, error) {
	client, err := ethclient.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	clog.Bluef("[EVM] blockNumber: %+v", blockNumber)

	return &Client{Client: client}, nil
}
