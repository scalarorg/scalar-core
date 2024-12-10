package btc

import (
	goerrors "errors"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/btcjson"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/vald/btc/rpc"
	"github.com/scalarorg/scalar-core/x/btc/types"

	"github.com/axelarnetwork/axelar-core/sdk-utils/broadcast"
	"github.com/axelarnetwork/axelar-core/utils/errors"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	"github.com/axelarnetwork/utils/log"
	"github.com/axelarnetwork/utils/monads/results"
	"github.com/axelarnetwork/utils/slices"
	sdkClient "github.com/cosmos/cosmos-sdk/client"
)

// ErrNotFinalized is returned when a transaction is not finalized
var ErrNotFinalized = goerrors.New("not finalized")

// ErrTxFailed is returned when a transaction has failed
var ErrTxFailed = goerrors.New("transaction failed")

// Mgr manages all communication with Ethereum
type Mgr struct {
	rpcs                      map[string]rpc.Client
	broadcaster               broadcast.Broadcaster
	validator                 sdk.ValAddress
	proxy                     sdk.AccAddress
	latestFinalizedBlockCache LatestFinalizedBlockCache
	blockHeightCache          BlockHeightCache
}

// NewMgr returns a new Mgr instance
func NewMgr(
	clientCtx sdkClient.Context,
	rpcs map[string]rpc.Client,
	broadcaster broadcast.Broadcaster,
	valAddr sdk.ValAddress,
	latestFinalizedBlockCache LatestFinalizedBlockCache,
	blockHeightCache BlockHeightCache,
) *Mgr {
	return &Mgr{
		rpcs:                      rpcs,
		broadcaster:               broadcaster,
		validator:                 valAddr,
		proxy:                     clientCtx.FromAddress,
		latestFinalizedBlockCache: latestFinalizedBlockCache,
		blockHeightCache:          blockHeightCache,
	}
}

func (mgr Mgr) logger(keyvals ...any) log.Logger {
	keyvals = append([]any{"listener", "btc"}, keyvals...)
	return log.WithKeyVals(keyvals...)
}

func (mgr Mgr) isFinalized(chain nexus.ChainName, txReceipt btcjson.TxRawResult, confHeight uint64) (bool, error) {
	client, ok := mgr.rpcs[strings.ToLower(chain.String())]
	if !ok {
		return false, fmt.Errorf("rpc client not found for chain %s", chain.String())
	}

	blockHeightCache := mgr.blockHeightCache.Get(txReceipt.BlockHash)
	if blockHeightCache == nil {
		clog.Redf("block_height_not_found, block_hash: %s, block_height: %d", txReceipt.BlockHash, confHeight)
		blockHeight, err := client.GetBlockHeight(txReceipt.BlockHash)
		if err != nil {
			return false, err
		}
		mgr.blockHeightCache.Set(txReceipt.BlockHash, blockHeight)
		blockHeightCache = &blockHeight
	}

	if mgr.latestFinalizedBlockCache.Get(chain) != nil && *mgr.latestFinalizedBlockCache.Get(chain) >= *blockHeightCache {
		return true, nil
	}

	latestFinalizedBlockHeight, err := client.LatestFinalizedBlockHeight(confHeight)
	if err != nil {
		return false, err
	}

	mgr.blockHeightCache.Set(txReceipt.BlockHash, latestFinalizedBlockHeight)
	mgr.latestFinalizedBlockCache.Set(chain, latestFinalizedBlockHeight)

	// This is a rare case, but it can happen if the block height is not updated in the cache
	if latestFinalizedBlockHeight < *blockHeightCache {
		return false, nil
	}

	return true, nil
}

// GetTxReceiptReceiptIfFinalized retrieves receipt for provided transaction ID.
//
// # Result is
//
// - Ok(receipt) if the transaction is finalized and successful
//
// - Err(ethereum.NotFound) if the transaction is not found
//
// - Err(ErrTxFailed) if the transaction is finalized but failed
//
// - Err(ErrNotFinalized) if the transaction is not finalized
//
// - Err(err) otherwise
func (mgr Mgr) GetTxReceiptIfFinalized(chain nexus.ChainName, txID types.Hash, confHeight uint64) (results.Result[rpc.TxReceipt], error) {
	txReceipts, err := mgr.GetTxReceiptsIfFinalized(chain, []types.Hash{txID}, confHeight)
	if err != nil {
		return results.Result[rpc.TxReceipt]{}, err
	}

	return txReceipts[0], err
}

// GetTxReceiptsIfFinalized retrieves receipts for provided transaction IDs.
//
// # Individual result is
//
// - Ok(receipt) if the transaction is finalized and successful
//
// - Err(ethereum.NotFound) if the transaction is not found
//
// - Err(ErrTxFailed) if the transaction is finalized but failed
//
// - Err(ErrNotFinalized) if the transaction is not finalized
//
// - Err(err) otherwise
func (mgr Mgr) GetTxReceiptsIfFinalized(chain nexus.ChainName, txIDs []types.Hash, confHeight uint64) ([]results.Result[rpc.TxReceipt], error) {
	client, ok := mgr.rpcs[strings.ToLower(chain.String())]
	if !ok {
		return nil, fmt.Errorf("rpc client not found for chain %s", chain.String())
	}

	receipts, err := client.GetTransactions(txIDs)
	if err != nil {
		return slices.Map(txIDs, func(_ types.Hash) results.Result[rpc.TxReceipt] {
			return results.FromErr[rpc.TxReceipt](
				sdkerrors.Wrapf(
					errors.With(err, "chain", chain.String(), "tx_ids", txIDs),
					"cannot get transaction receipts"),
			)
		}), nil
	}

	return slices.Map(receipts, func(receipt rpc.TxResult) results.Result[rpc.TxReceipt] {
		return results.Pipe(results.Result[rpc.TxReceipt](receipt), func(receipt rpc.TxReceipt) results.Result[rpc.TxReceipt] {

			isFinalized, err := mgr.isFinalized(chain, receipt.Data, confHeight)
			if err != nil {
				return results.FromErr[rpc.TxReceipt](sdkerrors.Wrapf(errors.With(err, "chain", chain.String()),
					"cannot determine if the transaction %s is finalized", receipt.Data.Txid),
				)
			}

			if !isFinalized {
				return results.FromErr[rpc.TxReceipt](ErrNotFinalized)
			}

			if receipt.Data.Confirmations <= confHeight {
				return results.FromErr[rpc.TxReceipt](ErrTxFailed)
			}

			return results.FromOk[rpc.TxReceipt](receipt)
		})
	}), nil
}

// isParticipantOf checks if the validator is in the poll participants list
func (mgr Mgr) isParticipantOf(participants []sdk.ValAddress) bool {
	return slices.Any(participants, func(v sdk.ValAddress) bool { return v.Equals(mgr.validator) })
}
