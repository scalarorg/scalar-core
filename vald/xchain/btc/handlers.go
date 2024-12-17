package btc

import (
	"bytes"
	"encoding/hex"
	"sync"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/errors"
	"github.com/scalarorg/scalar-core/utils/monads/results"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/vald/xchain"
	"github.com/scalarorg/scalar-core/x/btc/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *BtcClient) ProcessStakingTxsConfirmation(event *types.EventConfirmStakingTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	client.logger("event", event).Debug("processing staking txs confirmation poll")
	txIDs := slices.Map(event.PollMappings, func(m types.PollMapping) xchain.Hash { return m.TxID })
	txReceipts, err := client.GetTxReceiptsIfFinalized(txIDs, event.ConfirmationHeight)
	if err != nil {
		return nil, err
	}

	var votes []sdk.Msg
	for i, txReceipt := range txReceipts {
		pollID := event.PollMappings[i].PollID
		txID := event.PollMappings[i].TxID

		logger := client.logger("chain", event.Chain.String(), "poll_id", pollID.String(), "tx_id", txID.HexStr())

		if txReceipt.Err() != nil {
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain)))

			logger.Infof("broadcasting empty vote for poll %s: %s", pollID.String(), txReceipt.Err().Error())
		} else {
			events := client.processStakingTxReceipt(event.Chain, txReceipt.Ok().(BTCTxReceipt))
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain, events...)))

			logger.Infof("broadcasting vote %v for poll %s", events, pollID.String())
		}
	}

	return votes, nil
}

func (client *BtcClient) processStakingTxReceipt(chain nexus.ChainName, receipt BTCTxReceipt) []types.Event {

	var events []types.Event

	btcEvent, err := client.decodeStakingTransaction(&receipt)
	if err != nil {
		client.logger().Debug(sdkerrors.Wrap(err, "decode event ContractCall failed").Error())
	}

	if err := btcEvent.ValidateBasic(); err != nil {
		client.logger().Debug(sdkerrors.Wrap(err, "invalid event ContractCall").Error())
	}

	txID, err := types.HashFromHexStr(receipt.Raw.Txid)
	if err != nil {
		client.logger().Debug(sdkerrors.Wrap(err, "invalid tx id").Error())
	}

	events = append(events, types.Event{
		Chain: chain,
		TxID:  *txID,
		Event: &types.Event_StakingTx{
			StakingTx: &btcEvent,
		},
		Index: 0, // TODO: fix this hardcoded index, this is used to identify the staking tx in the event
	})

	return events
}

func (client *BtcClient) ProcessUnstakingTxsConfirmation(event *types.EventConfirmUnstakingTxsStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	panic("not implemented")
}

func (client *BtcClient) GetTxReceiptsIfFinalized(txIDs []xchain.Hash, confHeight uint64) ([]BTCTxResult, error) {
	txResults, err := client.GetTransactions(txIDs)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			errors.With(err, "tx_ids", txIDs),
			"cannot get transaction receipts",
		)
	}

	return slices.Map(txResults, func(receipt BTCTxResult) results.Result[xchain.TxReceipt] {
		return results.Pipe(results.Result[xchain.TxReceipt](receipt), func(receipt xchain.TxReceipt) results.Result[xchain.TxReceipt] {
			btcReceipt := receipt.(BTCTxReceipt)
			isFinalized, err := client.isFinalized(btcReceipt.Raw, confHeight)
			if err != nil {
				return results.FromErr[xchain.TxReceipt](sdkerrors.Wrapf(errors.With(err, "tx_id", btcReceipt.Raw.Txid),
					"cannot determine if the transaction %s is finalized", btcReceipt.Raw.Txid),
				)
			}

			if !isFinalized {
				return results.FromErr[xchain.TxReceipt](xchain.ErrNotFinalized)
			}

			if btcReceipt.Raw.Confirmations <= confHeight {
				return results.FromErr[xchain.TxReceipt](xchain.ErrTxFailed)
			}

			return results.FromOk(receipt)
		})
	}), nil
}

func (c *BtcClient) GetTransactions(txIDs []xchain.Hash) ([]BTCTxResult, error) {
	txs := make([]BTCTxResult, len(txIDs))
	var wg sync.WaitGroup

	for i := range txIDs {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			tx, err := c.GetTransaction(txIDs[index])
			if err != nil {
				txs[index] = BTCTxResult(results.FromErr[xchain.TxReceipt](err))
			} else {
				txs[index] = tx
			}
		}(i)
	}

	wg.Wait()
	if slices.Any(txs, func(tx BTCTxResult) bool { return tx.IsErr() }) {
		return nil, xchain.ErrFailedToGetTransactions
	}

	return txs, nil
}

func (c *BtcClient) GetTransaction(txID xchain.Hash) (BTCTxResult, error) {
	var tx BTCTxReceipt
	chainHash := chainhash.Hash(txID)
	txMetadata, err := c.client.GetRawTransactionVerbose(&chainHash)
	if err != nil {
		c.logger("failed to get BTC transaction", "txID", txID, "error", err)
		return BTCTxResult(results.FromErr[xchain.TxReceipt](err)), err
	} else {
		txRaw, err := hex.DecodeString(txMetadata.Hex)
		if err != nil {
			c.logger("failed to decode hex string", "txID", txID, "error", err)
			return BTCTxResult(results.FromErr[xchain.TxReceipt](err)), err
		}

		msgTx := wire.NewMsgTx(wire.TxVersion)
		err = msgTx.Deserialize(bytes.NewReader(txRaw))
		if err != nil {
			c.logger("failed to parse transaction", "txID", txID, "error", err)
			return BTCTxResult(results.FromErr[xchain.TxReceipt](err)), err
		}

		prevTxOuts, err := c.GetTxOuts(slices.Map(msgTx.TxIn, func(txIn *wire.TxIn) wire.OutPoint {
			return txIn.PreviousOutPoint
		}))
		if err != nil {
			c.logger("failed to get BTC transaction", "txID", txID, "error", err)
			return BTCTxResult(results.FromErr[xchain.TxReceipt](err)), err
		}

		tx.Raw = *txMetadata
		tx.PrevTxOuts = prevTxOuts
		tx.MsgTx = msgTx
	}

	return results.FromOk[xchain.TxReceipt](tx), nil
}

func (c *BtcClient) GetTxOuts(outpoints []wire.OutPoint) ([]*btcjson.GetTxOutResult, error) {
	txOuts := make([]*btcjson.GetTxOutResult, len(outpoints))
	var wg sync.WaitGroup

	for i := range outpoints {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			txOut, err := c.GetTxOut(outpoints[index])
			if err != nil {
				txOuts[index] = nil
			} else {
				txOuts[index] = txOut
			}
		}(i)
	}

	wg.Wait()
	return txOuts, nil
}

func (c *BtcClient) GetTxOut(outpoint wire.OutPoint) (*btcjson.GetTxOutResult, error) {
	return c.client.GetTxOut(&outpoint.Hash, outpoint.Index, false)
}

func (c *BtcClient) LatestFinalizedBlockHeight(_ uint64) (uint64, error) {
	_, height, err := c.client.GetBestBlock()
	if err != nil {
		return 0, err
	}

	return uint64(height), nil
}

func (c *BtcClient) GetBlockHeight(blockHash string) (uint64, error) {
	chainhashBlockHash, err := chainhash.NewHashFromStr(blockHash)
	if err != nil {
		return 0, err
	}

	block, err := c.client.GetBlockHeaderVerbose(chainhashBlockHash)
	if err != nil {
		return 0, err
	}

	return uint64(block.Height), nil
}
