package evm

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	grpc_client "github.com/scalarorg/scalar-core/vald/grpc-client"
	"github.com/scalarorg/scalar-core/vald/xchain/common"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	covTypes "github.com/scalarorg/scalar-core/x/covenant/types"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *EthereumClient) ProcessSwitchedPhaseConfirmation(event *covTypes.ConfirmSwitchedPhaseStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	txReceipt, _ := client.GetTransaction(common.Hash(event.TxID))

	if txReceipt.Err() != nil {
		// This ensures all of the txs are valid
		return nil, fmt.Errorf("exist error in txReceipt: %s, txID: %s", txReceipt.Err().Error(), event.TxID)
	}
	eventId := client.processConfirmSwitchedPhaseTxReceipt(txReceipt.Ok().(ETHTxReceipt))

	queryClient := grpc_client.QueryManager().GetChainsClient()
	chainID := event.Chain

	chainParams, err := queryClient.Params(context.Background(), &chainsTypes.ParamsRequest{
		Chain: string(chainID),
	})
	if err != nil {
		return nil, fmt.Errorf("error getting chain metadata: %w", err)
	}

	chainMetadata := chainParams.Params.Metadata

	params := chainMetadata["params"]
	if params == "" {
		return nil, fmt.Errorf("params is required")
	}

	voteEvent := covTypes.NewVoteEvents(event.Chain, covTypes.Event{
		Chain: event.Chain,
		Hash:  nil,
		Event: &covTypes.Event_SwitchedPhaseConfirmed{
			SwitchedPhaseConfirmed: &covTypes.SwitchedPhaseConfirmed{
				EventID: *eventId,
			},
		},
		Index: 0,
	})

	events := covTypes.Event{}
	votes := []sdk.Msg{
		voteTypes.NewVoteRequest(proxy, event.PollID, voteEvent),
	}
	clog.Redf("broadcasting vote %v for poll %s", events, event.PollID.String())

	return votes, nil
}

func (client *EthereumClient) processConfirmSwitchedPhaseTxReceipt(receipt ETHTxReceipt) *chainsTypes.EventID {
	// TODO: 🛑 validate the btc protocol address from the event

	// TODO: Validate the session sequence, custodian_group_id in OP_RETURN

	// Note: TxID is the reversed-order hash of the txid aka RPC TxID, aka Mempool TxID
	txID, err := chains.HashFromHex(receipt.TxHash.Hex())
	if err != nil {
		client.logger().Error(sdkerrors.Wrap(err, "invalid tx id").Error())
		return nil
	}

	eventId := chainsTypes.NewEventID(txID, uint64(receipt.TransactionIndex))

	return &eventId
}
