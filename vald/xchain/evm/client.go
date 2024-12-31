package evm

import (
	"context"
	"fmt"
	"strings"

	"github.com/scalarorg/scalar-core/utils/monads/results"
	"github.com/scalarorg/scalar-core/vald/xchain/common"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type EthereumClient struct {
	*ethclient.Client
	rpc                       *rpc.Client
	latestFinalizedBlockCache common.LatestFinalizedBlockCache
}

type ETHTxReceipt = types.Receipt

type ETHTxResult = results.Result[common.TxReceipt]

var _ common.Client = &EthereumClient{}

// func NewClient(url string, override FinalityOverride) (Client, error) {
// 	rpc, err := rpc.DialContext(context.Background(), url)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ethereumClient, err := NewEthereumClient(ethclient.NewClient(rpc), rpc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if override == Confirmation {
// 		return ethereumClient, nil
// 	}

// 	if ethereum2Client, err := NewEthereum2Client(ethereumClient); err == nil {
// 		return ethereum2Client, nil
// 	}

// 	return ethereumClient, nil
// }

func NewClient(url, finalityOverride string) (common.Client, error) {
	rpcClient, err := rpc.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}

	if !strings.EqualFold(finalityOverride, Confirmation.String()) {
		return nil, fmt.Errorf("no override for ethereum client, only confirmation is supported")
	}

	ethClient := ethclient.NewClient(rpcClient)

	if _, err := ethClient.BlockNumber(context.Background()); err != nil {
		return nil, err
	}

	latestFinalizedBlockCache := common.NewLatestFinalizedBlockCache()

	client := &EthereumClient{
		rpc:                       rpcClient,
		Client:                    ethClient,
		latestFinalizedBlockCache: latestFinalizedBlockCache,
	}

	return client, nil
}

func (c *EthereumClient) Close() {
	c.rpc.Close()
	c.Client.Close()
}
