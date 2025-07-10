package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	btc_utils "github.com/scalarorg/go-common/btc"
	"github.com/scalarorg/scalar-core/utils/btc"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

func (s msgServer) ConfirmSourceTxsV2(c context.Context, req *types.ConfirmSourceTxsRequestV2) (*types.ConfirmSourceTxsResponseV2, error) {
	ctx := sdk.UnwrapSDKContext(c)

	chain, ok := s.nexus.GetChain(ctx, req.Chain)
	if !ok {
		return nil, fmt.Errorf("%s is not a registered chain", req.Chain)
	}

	if err := validateChainActivated(ctx, s.nexus, chain); err != nil {
		return nil, err
	}

	clog.Green("After validateChainActivated", chain)

	keeper, err := s.ForChain(ctx, chain.Name)
	if err != nil {
		return nil, err
	}

	// TODO: support evm
	if !types.IsBitcoinChain(chain.Name) {
		return nil, fmt.Errorf("chain %s is not a btc chain", chain.Name)
	}

	// each tx contains: tx_id, raw, merkle_path, block_hash_chain
	// 1. validate tx_id = hash(raw)
	// 2. get block_hash from keeper then validate the merkle_path
	// 3. TODO: validate block_hash_chain, 6-12 blocks

	block, err := keeper.GetBlock(ctx, req.Batch.BlockHash)
	if err == nil {
		s.Logger(ctx).Info("Block found, start to confirm txs", "block_hash", req.Batch.BlockHash)
		start := time.Now()
		// Get sender address
		chainParams := keeper.GetParams(ctx)

		nwParams := chainParams.Metadata["params"]
		if nwParams == "" {
			return nil, fmt.Errorf("params are required")
		}
		for _, tx := range req.Batch.Txs {
			txInfo, err := btc.ParseTx(tx.Raw)
			if err != nil {
				s.Logger(ctx).Error("failed to parse tx", "error", err)
				continue
			}

			if txInfo.MsgTx.TxHash().String() != tx.Hash.String() {
				s.Logger(ctx).Error("tx hash mismatch", "expected", tx.Hash.String(), "actual", txInfo.MsgTx.TxHash().String())
				continue
			}
			err = btc.ValidateTxProof(txInfo.TxHash.CloneBytes(), tx.TxIndex, tx.MerklePath, block.MerkleRoot)
			if err != nil {
				clog.Greenf("BlockHash: %s, MerkleRoot: %s", block.BlockHash.Hex(), block.MerkleRoot.Hex())
				s.Logger(ctx).Error("failed to validate tx proof", "error", err)
				continue
			}
			protocolInfo, err := s.protocol.FindProtocolInfoByInternalAddress(ctx, chain.Name, nexus.ChainName(txInfo.DestinationChain), txInfo.DestinationTokenAddress)
			if err != nil {
				s.Logger(ctx).Error("failed to find protocol info by internal address", "error", err)
				continue
			}

			clog.Greenf("txInfo: %+v", txInfo)
			clog.Greenf("tx.PrevOutpointScriptPubkey: %x", tx.PrevOutpointScriptPubkey)

			sender, err := btc_utils.ScriptPubKeyToAddress(tx.PrevOutpointScriptPubkey, nwParams)
			if err != nil {
				s.Logger(ctx).Error("Failed to get sender address", "error", err)
				continue
			}
			err = keeper.ProcessConfirmRequestTx(ctx, txInfo, tx.TxIndex, block.Height, protocolInfo.Symbol, sender.String())
			if err != nil {
				s.Logger(ctx).Error("failed to process confirm request batch", "error", err)
				continue
			}
			// if err := s.validateAndSaveTokenSent(ctx, keeper, chain.Name, tx, block); err != nil {
			// 	clog.Redf("Failed to validate and save token sent: %v", err)
			// 	continue
			// }
		}
		s.Logger(ctx).Info("ConfirmSourceTxsV2", "time", time.Since(start), "number of txs", len(req.Batch.Txs))
	} else {
		if !keeper.HasPendingConfirmRequest(ctx, req.Batch.BlockHash) {
			s.Logger(ctx).Info("First time see the block, start to confirm block header", "block_hash", req.Batch.BlockHash.Hex())
			err := s.startConfirmBlock(ctx, keeper, chain, req.Batch)
			if err != nil {
				s.Logger(ctx).Error("Failed to start confirm block", "error", err)
				return nil, err
			}
		} else {
			s.Logger(ctx).Info("Block Confirmation is in processing", "block_hash", req.Batch.BlockHash.Hex())
		}
	}
	return &types.ConfirmSourceTxsResponseV2{}, nil
}

func (s msgServer) startConfirmBlock(ctx sdk.Context, keeper types.ChainKeeper, chain nexus.Chain, batch *types.TrustedTxsByBlock) error {
	//Store request
	pollParticipants, err := s.initializeBlockConfirmPoll(ctx, chain, batch.BlockHash)
	if err != nil {
		s.Logger(ctx).Error("Failed to initialize poll", "error", err)
		return err
	}

	// Store the batch using pollId as the key
	keeper.SetPendingConfirmRequest(ctx, pollParticipants.PollID, batch)

	s.Logger(ctx).Info("Emit event ConfirmNewBlockStarted")
	events.Emit(ctx, &types.ConfirmNewBlockStarted{
		Chain:              chain.Name,
		BlockHash:          batch.BlockHash,
		PreviousBlockHash:  nil,
		ConfirmationHeight: keeper.GetRequiredConfirmationHeight(ctx),
		PollParticipants:   pollParticipants,
	})
	keeper.EnqueueConfirmedEvent(ctx, types.NewEventID(batch.BlockHash, 0))
	return nil
}
