package btc

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/chains/types"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *BtcClient) ProcessNewBlockConfirmation(event *types.ConfirmBtcNewBlockStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	result, err := client.GetBlockVerboseTx(event.BlockHash.String())
	if err != nil {
		return nil, err
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	block := result.Ok()
	if block.Confirmations < int64(event.ConfirmationHeight) {
		return nil, fmt.Errorf("block confirmations are less than confirmation height: %d < %d", block.Confirmations, event.ConfirmationHeight)
	}

	if block.PreviousHash != event.PreviousBlockHash.String() {
		return nil, fmt.Errorf("block previous block hash does not match")
	}

	return []sdk.Msg{voteTypes.NewVoteRequest(proxy, event.PollID, types.NewVoteEvents(event.Chain))}, nil
}
