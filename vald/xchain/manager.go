package xchain

import (
	"context"
	goerrors "errors"
	"fmt"

	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/scalar-core/sdk-utils/broadcast"
	"github.com/scalarorg/scalar-core/utils/log"

	// "fmt"
	// "github.com/btcsuite/btcd/btcjson"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/scalarorg/bitcoin-vault/go-utils/chain"
	// "github.com/scalarorg/scalar-core/utils/monads/results"
	// "github.com/scalarorg/scalar-core/vald/xchain/rpc"
	// "github.com/scalarorg/scalar-core/x/btc/types"
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	// "github.com/scalarorg/scalar-core/sdk-utils/broadcast"
	// "github.com/scalarorg/scalar-core/utils/errors"
	// "github.com/scalarorg/scalar-core/utils/log"
	// "github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/btc/types"
	vote "github.com/scalarorg/scalar-core/x/vote/exported"
)

// ErrNotFinalized is returned when a transaction is not finalized
var ErrNotFinalized = goerrors.New("not finalized")

// ErrTxFailed is returned when a transaction has failed
var ErrTxFailed = goerrors.New("transaction failed")

// Manager manages all communication with Ethereum
type Manager struct {
	rpcs                      map[chain.ChainInfoBytes]Client
	broadcaster               broadcast.Broadcaster
	validator                 sdk.ValAddress
	proxy                     sdk.AccAddress
	latestFinalizedBlockCache LatestFinalizedBlockCache
}

// NewManager returns a new Manager instance
func NewManager(
	clientCtx sdkClient.Context,
	rpcs map[chain.ChainInfoBytes]Client,
	broadcaster broadcast.Broadcaster,
	valAddr sdk.ValAddress,
) *Manager {
	latestFinalizedBlockCache := NewLatestFinalizedBlockCache()
	return &Manager{
		rpcs:                      rpcs,
		broadcaster:               broadcaster,
		validator:                 valAddr,
		proxy:                     clientCtx.FromAddress,
		latestFinalizedBlockCache: latestFinalizedBlockCache,
	}
}

func (mgr Manager) ProcessStakingTxsConfirmation(event *types.EventConfirmStakingTxsStarted) error {
	if !mgr.isParticipantOf(event.Participants) {
		pollIDs := slices.Map(event.PollMappings, func(m types.PollMapping) vote.PollID { return m.PollID })
		mgr.logger("poll_ids", pollIDs).Debug("ignoring gateway txs confirmation poll: not a participant")
		return nil
	}

	mgr.logger("event", event).Debug("processing staking txs confirmation poll")

	client, ok := mgr.rpcs[event.ChainInfo]
	if !ok {
		return fmt.Errorf("rpc client not found for chain %s", event.ChainInfo)
	}

	votes, err := client.ProcessStakingTxsConfirmation(event)
	if err != nil {
		return err
	}

	// txIDs := slices.Map(event.PollMappings, func(m types.PollMapping) types.Hash { return m.TxID })
	// txReceipts, err := mgr.GetTxReceiptsIfFinalized(event.Chain, txIDs, event.ConfirmationHeight)
	// if err != nil {
	// 	return err
	// }

	// var votes []sdk.Msg
	// for i, txReceipt := range txReceipts {
	// 	pollID := event.PollMappings[i].PollID
	// 	txID := event.PollMappings[i].TxID

	// 	logger := mgr.logger("chain", event.Chain, "poll_id", pollID.String(), "tx_id", txID.HexStr())

	// 	if txReceipt.Err() != nil {
	// 		votes = append(votes, voteTypes.NewVoteRequest(mgr.proxy, pollID, types.NewVoteEvents(event.Chain)))

	// 		logger.Infof("broadcasting empty vote for poll %s: %s", pollID.String(), txReceipt.Err().Error())
	// 	} else {
	// 		events := mgr.processStakingTxReceipt(event.Chain, txReceipt.Ok())
	// 		votes = append(votes, voteTypes.NewVoteRequest(mgr.proxy, pollID, types.NewVoteEvents(event.Chain, events...)))

	// 		logger.Infof("broadcasting vote %v for poll %s", events, pollID.String())
	// 	}
	// }

	_, err = mgr.broadcaster.Broadcast(context.TODO(), votes...)

	return err

}

// func (mgr Manager) ProcessUnstakingTxsConfirmation(event *btcTypes.EventConfirmUnstakingTxsStarted) error {

// 	if !mgr.isParticipantOf(event.Participants) {
// 		pollIDs := slices.Map(event.PollMappings, func(m types.PollMapping) vote.PollID { return m.PollID })
// 		mgr.logger("poll_ids", pollIDs).Debug("ignoring gateway txs confirmation poll: not a participant")
// 		return nil
// 	}

// 	mgr.logger("event", event).Debug("processing unstaking txs confirmation poll")

// 	chainInfo := chain.ChainInfo(event.)

// txIDs := slices.Map(event.PollMappings, func(m types.PollMapping) types.Hash { return m.TxID })
// txReceipts, err := mgr.GetTxReceiptsIfFinalized(event.Chain, txIDs, event.ConfirmationHeight)
// if err != nil {
// 	return err
// }

// var votes []sdk.Msg
// for i, txReceipt := range txReceipts {
// 	pollID := event.PollMappings[i].PollID
// 	txID := event.PollMappings[i].TxID

// 	logger := mgr.logger("chain", event.Chain, "poll_id", pollID.String(), "tx_id", txID.HexStr())

// 	if txReceipt.Err() != nil {
// 		votes = append(votes, voteTypes.NewVoteRequest(mgr.proxy, pollID, types.NewVoteEvents(event.Chain)))

// 		logger.Infof("broadcasting empty vote for poll %s: %s", pollID.String(), txReceipt.Err().Error())
// 	} else {
// 		events := mgr.processStakingTxReceipt(event.Chain, txReceipt.Ok())
// 		votes = append(votes, voteTypes.NewVoteRequest(mgr.proxy, pollID, types.NewVoteEvents(event.Chain, events...)))

// 		logger.Infof("broadcasting vote %v for poll %s", events, pollID.String())
// 	}
// }

// 	_, err = mgr.broadcaster.Broadcast(context.TODO(), votes...)

// 	return err

// }

// isParticipantOf checks if the validator is in the poll participants list
func (mgr Manager) isParticipantOf(participants []sdk.ValAddress) bool {
	return slices.Any(participants, func(v sdk.ValAddress) bool { return v.Equals(mgr.validator) })
}

func (mgr Manager) logger(keyvals ...any) log.Logger {
	keyvals = append([]any{"listener", "btc"}, keyvals...)
	return log.WithKeyVals(keyvals...)
}

// func (mgr Manager) isFinalized(chain chain.ChainInfoBytes, txReceipt btcjson.TxRawResult, confHeight uint64) (bool, error) {
// 	client, ok := mgr.rpcs[chain]
// 	if !ok {
// 		return false, fmt.Errorf("rpc client not found for chain %d", chain)
// 	}

// 	_ = client

// blockHeightCache := mgr.blockHeightCache.Get(txReceipt.BlockHash)
// if blockHeightCache == nil {
// 	clog.Redf("block_height_not_found, block_hash: %s, block_height: %d", txReceipt.BlockHash, confHeight)
// 	blockHeight, err := client.GetBlockHeight(txReceipt.BlockHash)
// 	if err != nil {
// 		return false, err
// 	}
// 	mgr.blockHeightCache.Set(txReceipt.BlockHash, blockHeight)
// 	blockHeightCache = &blockHeight
// }

// if mgr.latestFinalizedBlockCache.Get(chain) != nil && *mgr.latestFinalizedBlockCache.Get(chain) >= *blockHeightCache {
// 	return true, nil
// }

// latestFinalizedBlockHeight, err := client.LatestFinalizedBlockHeight(confHeight)
// if err != nil {
// 	return false, err
// }

// mgr.blockHeightCache.Set(txReceipt.BlockHash, latestFinalizedBlockHeight)
// mgr.latestFinalizedBlockCache.Set(chain, latestFinalizedBlockHeight)

// // This is a rare case, but it can happen if the block height is not updated in the cache
// if latestFinalizedBlockHeight < *blockHeightCache {
// 	return false, nil
// }

// 	return true, nil
// }

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
// func (mgr Manager) GetTxReceiptIfFinalized(chain chain.ChainInfoBytes, txID Hash, confHeight uint64) (TxResult, error) {
// 	txReceipts, err := mgr.GetTxReceiptsIfFinalized(chain, []Hash{txID}, confHeight)
// 	if err != nil {
// 		return TxResult{}, err
// 	}

// 	return txReceipts[0], err
// }

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
// func (mgr Manager) GetTxReceiptsIfFinalized(chain chain.ChainInfoBytes, txIDs []Hash, confHeight uint64) ([]TxResult, error) {
// 	client, ok := mgr.rpcs[chain]
// 	if !ok {
// 		return nil, fmt.Errorf("rpc client not found for chain %d", chain)
// 	}

// 	txResults, err := client.GetTransactions(txIDs)
// 	if err != nil {
// 		return nil, sdkerrors.Wrapf(
// 			errors.With(err, "chain", chain, "tx_ids", txIDs),
// 			"cannot get transaction receipts",
// 		)
// 	}

// 	return slices.Map(txResults, func(receipt TxResult) results.Result[TxReceipt] {
// 		return results.Pipe(results.Result[TxReceipt](receipt), func(receipt TxReceipt) results.Result[TxReceipt] {
// 			isFinalized, err := mgr.isFinalized(chain, receipt.Raw, confHeight)
// 			if err != nil {
// 				return results.FromErr[TxReceipt](sdkerrors.Wrapf(errors.With(err, "chain", chain),
// 					"cannot determine if the transaction %s is finalized", receipt.Raw.Txid),
// 				)
// 			}

// 			if !isFinalized {
// 				return results.FromErr[TxReceipt](ErrNotFinalized)
// 			}

// 			if receipt.Raw.Confirmations <= confHeight {
// 				return results.FromErr[TxReceipt](ErrTxFailed)
// 			}

// 			return results.FromOk(receipt)
// 		})
// 	}), nil
// }

// func (mgr *Manager) getTxOut(chain chain.ChainInfoBytes, outpoint wire.OutPoint) (*btcjson.GetTxOutResult, error) {
// 	client, ok := mgr.[chain]
// 	if !ok {
// 		return nil, fmt.Errorf("rpc client not found for chain %d", chain)
// 	}

// 	txOut, err := client.GetTxOut(outpoint)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return txOut, nil
// }
