package evm

import (
	"context"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/errors"
	"github.com/scalarorg/scalar-core/utils/monads/results"
	"github.com/scalarorg/scalar-core/utils/slices"
	xcommon "github.com/scalarorg/scalar-core/vald/xchain/common"
)

func (client *EthereumClient) GetTxReceiptsIfFinalized(txIDs []xcommon.Hash, confHeight uint64) ([]ETHTxResult, error) {
	txResults, err := client.GetTransactions(txIDs)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			errors.With(err, "tx_ids", txIDs),
			"cannot get transaction receipts",
		)
	}

	return slices.Map(txResults, func(receipt ETHTxResult) results.Result[xcommon.TxReceipt] {
		return results.Pipe(results.Result[xcommon.TxReceipt](receipt), func(receipt xcommon.TxReceipt) results.Result[xcommon.TxReceipt] {
			ethReceipt := receipt.(ETHTxReceipt)
			isFinalized, err := client.isFinalized(ethReceipt, confHeight)
			if err != nil {
				return results.FromErr[xcommon.TxReceipt](sdkerrors.Wrapf(errors.With(err, "tx_id", ethReceipt.TxHash.Hex()),
					"cannot determine if the transaction %s is finalized", ethReceipt.TxHash.Hex()),
				)
			}

			if !isFinalized {
				return results.FromErr[xcommon.TxReceipt](xcommon.ErrNotFinalized)
			}

			if ethReceipt.Status != types.ReceiptStatusSuccessful {
				return results.FromErr[xcommon.TxReceipt](xcommon.ErrTxFailed)
			}

			return results.FromOk(receipt)
		})
	}), nil
}

func (c *EthereumClient) GetTransactions(txIDs []xcommon.Hash) ([]ETHTxResult, error) {
	ctx := context.Background()
	batch := slices.Map(txIDs, func(txHash xcommon.Hash) rpc.BatchElem {
		var receipt *types.Receipt
		return rpc.BatchElem{
			Method: "eth_getTransactionReceipt",
			Args:   []interface{}{common.Hash(txHash)},
			Result: &receipt,
		}
	})

	clog.Redf("[ETH] batch: %+v", batch)

	if err := c.rpc.BatchCallContext(ctx, batch); err != nil {
		return nil, fmt.Errorf("unable to send batch request: %v", err)
	}

	return slices.Map(batch, func(elem rpc.BatchElem) ETHTxResult {
		if elem.Error != nil {
			return ETHTxResult(results.FromErr[xcommon.TxReceipt](elem.Error))
		}

		receipt := elem.Result.(**ETHTxReceipt)
		if *receipt == nil {
			return ETHTxResult(results.FromErr[xcommon.TxReceipt](ethereum.NotFound))
		}

		return ETHTxResult(results.FromOk(xcommon.TxReceipt(**receipt)))
	}), nil
}

func (c *EthereumClient) GetTransaction(txID xcommon.Hash) (ETHTxResult, error) {
	ctx := context.Background()
	receipt := &types.Receipt{}

	if err := c.rpc.CallContext(ctx, receipt, "eth_getTransactionReceipt", txID); err != nil {
		return ETHTxResult(results.FromErr[xcommon.TxReceipt](err)), err
	}

	return ETHTxResult(results.FromOk(xcommon.TxReceipt(*receipt))), nil
}

func (c *EthereumClient) LatestFinalizedBlockHeight(_ uint64) (uint64, error) {
	blockNumber, err := c.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}

	return blockNumber, nil
}
