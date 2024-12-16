package btc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/vald/xchain"
	"github.com/scalarorg/scalar-core/x/btc/types"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *BtcClient) ProcessStakingTxsConfirmation(event *types.EventConfirmStakingTxsStarted, proxy sdk.AccAddress) error {
	client.logger("event", event).Debug("processing staking txs confirmation poll")

	txIDs := slices.Map(event.PollMappings, func(m types.PollMapping) xchain.Hash { return m.TxID })
	txReceipts, err := client.GetTxReceiptsIfFinalized(txIDs, event.ConfirmationHeight)
	if err != nil {
		return err
	}

	var votes []sdk.Msg
	for i, txReceipt := range txReceipts {
		pollID := event.PollMappings[i].PollID
		txID := event.PollMappings[i].TxID

		logger := client.logger("chain_info", event.ChainInfo.String(), "poll_id", pollID.String(), "tx_id", txID.HexStr())

		if txReceipt.Err() != nil {
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain)))

			logger.Infof("broadcasting empty vote for poll %s: %s", pollID.String(), txReceipt.Err().Error())
		} else {
			events := client.processStakingTxReceipt(event.Chain, txReceipt.Ok())
			votes = append(votes, voteTypes.NewVoteRequest(proxy, pollID, types.NewVoteEvents(event.Chain, events...)))

			logger.Infof("broadcasting vote %v for poll %s", events, pollID.String())
		}
	}

	return err
}

func (client *BtcClient) GetTxReceiptsIfFinalized(txIDs []xchain.Hash, confHeight uint64) ([]xchain.TxResult, error) {

	return nil, nil
}

// func (mgr Manager) processStakingTxReceipt(chain chain.ChainInfoBytes, receipt rpc.TxReceipt) []types.Event {

// 	var events []types.Event

// 	btcEvent, err := mgr.decodeStakingTransaction(&receipt)
// 	if err != nil {
// 		mgr.logger().Debug(sdkerrors.Wrap(err, "decode event ContractCall failed").Error())
// 	}

// 	if err := btcEvent.ValidateBasic(); err != nil {
// 		mgr.logger().Debug(sdkerrors.Wrap(err, "invalid event ContractCall").Error())
// 	}

// 	txID, err := types.HashFromHexStr(receipt.Raw.Txid)
// 	if err != nil {
// 		mgr.logger().Debug(sdkerrors.Wrap(err, "invalid tx id").Error())
// 	}

// 	events = append(events, types.Event{
// 		Chain: chain,
// 		TxID:  *txID,
// 		Event: &types.Event_StakingTx{
// 			StakingTx: &btcEvent,
// 		},
// 		Index: 0, // TODO: fix this hardcoded index, this is used to identify the staking tx in the event
// 	})

// 	return events
// }
