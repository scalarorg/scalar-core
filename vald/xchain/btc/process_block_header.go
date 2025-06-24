package btc

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcjson"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/btc"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/chains/types"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *BtcClient) CalculateMerkleRoot(block *btcjson.GetBlockVerboseTxResult) ([]byte, error) {
	txHashes := [][]byte{}
	for _, tx := range block.Tx {
		txHash, err := hex.DecodeString(tx.Hash)
		if err != nil {
			return nil, err
		}
		txHashes = append(txHashes, btc.DoubleSha256(txHash))
	}
	return btc.CalculateMerkleRoot(txHashes), nil
}

func (client *BtcClient) ProcessNewBlockConfirmation(event *types.ConfirmNewBlockStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
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

	if event.PreviousBlockHash != nil && block.PreviousHash != event.PreviousBlockHash.String() {
		return nil, fmt.Errorf("block previous block hash does not match")
	}

	merkleRootBytes, err := client.CalculateMerkleRoot(block)
	if err != nil {
		return nil, err
	}
	calculatedMerkleRoot := exported.HashFromBytes(merkleRootBytes)
	clog.Greenf("Calculated merkle root: %s", calculatedMerkleRoot.Hex())
	clog.Greenf("Block merkle root: %s", block.MerkleRoot)
	// merkleRoot, err := exported.HashFromHex(block.MerkleRoot)
	// if err != nil {
	// 	return nil, err
	// }

	voteEvent := types.NewVoteEvents(event.Chain, types.Event{
		Chain: event.Chain,
		Hash:  event.BlockHash,
		Event: &types.Event_NewBlockConfirmed{
			NewBlockConfirmed: &types.EventNewBlockConfirmed{
				BlockHash:         event.BlockHash,
				PreviousBlockHash: event.PreviousBlockHash,
				MerkleRoot:        calculatedMerkleRoot,
				BlockHeight:       uint64(block.Height),
			},
		},
		Index: uint64(block.Height),
	})

	clog.Greenf("New block confirmed: %+v", voteEvent)

	return []sdk.Msg{voteTypes.NewVoteRequest(proxy, event.PollID, voteEvent)}, nil
}
