package btc

import (
	"context"

	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	vote "github.com/axelarnetwork/axelar-core/x/vote/exported"
	voteTypes "github.com/axelarnetwork/axelar-core/x/vote/types"
	"github.com/axelarnetwork/utils/slices"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/vald/btc/rpc"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

func (mgr Mgr) ProcessStakingTxsConfirmation(event *types.EventConfirmStakingTxsStarted) error {

	if !mgr.isParticipantOf(event.Participants) {
		pollIDs := slices.Map(event.PollMappings, func(m types.PollMapping) vote.PollID { return m.PollID })
		mgr.logger("poll_ids", pollIDs).Debug("ignoring gateway txs confirmation poll: not a participant")
		return nil
	}

	mgr.logger("event", event).Debug("processing staking txs confirmation poll")

	txIDs := slices.Map(event.PollMappings, func(m types.PollMapping) types.Hash { return m.TxID })
	txReceipts, err := mgr.GetTxReceiptsIfFinalized(event.Chain, txIDs, event.ConfirmationHeight)
	if err != nil {
		return err
	}

	var votes []sdk.Msg
	for i, txReceipt := range txReceipts {
		pollID := event.PollMappings[i].PollID
		txID := event.PollMappings[i].TxID

		logger := mgr.logger("chain", event.Chain, "poll_id", pollID.String(), "tx_id", txID.HexStr())

		if txReceipt.Err() != nil {
			votes = append(votes, voteTypes.NewVoteRequest(mgr.proxy, pollID, types.NewVoteEvents(event.Chain)))

			logger.Infof("broadcasting empty vote for poll %s: %s", pollID.String(), txReceipt.Err().Error())
		} else {
			events := mgr.processStakingTxReceipt(event.Chain, txReceipt.Ok())
			votes = append(votes, voteTypes.NewVoteRequest(mgr.proxy, pollID, types.NewVoteEvents(event.Chain, events...)))

			logger.Infof("broadcasting vote %v for poll %s", events, pollID.String())
		}
	}

	_, err = mgr.broadcaster.Broadcast(context.TODO(), votes...)

	return err

}

func (mgr Mgr) processStakingTxReceipt(chain nexus.ChainName, receipt rpc.TxReceipt) []types.Event {

	var events []types.Event

	btcEvent, err := DecodeStakingTransaction(&receipt)
	if err != nil {
		mgr.logger().Debug(sdkerrors.Wrap(err, "decode event ContractCall failed").Error())
	}

	if err := btcEvent.ValidateBasic(); err != nil {
		mgr.logger().Debug(sdkerrors.Wrap(err, "invalid event ContractCall").Error())
	}

	txID, err := types.HashFromHexStr(receipt.Data.Txid)
	if err != nil {
		mgr.logger().Debug(sdkerrors.Wrap(err, "invalid tx id").Error())
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
