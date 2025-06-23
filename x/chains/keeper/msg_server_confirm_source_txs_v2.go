package keeper

import (
	"bytes"
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	btc_utils "github.com/scalarorg/go-common/btc"
	"github.com/scalarorg/scalar-core/utils/btc"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/chains/exported"
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

	block, err := keeper.GetBlock(ctx, req.Batch.BlockHash)
	if err == nil {
		// Block found, process transactions
		s.Logger(ctx).Info("Block found, start to confirm txs", "block_hash", req.Batch.BlockHash)
		start := time.Now()
		chainParams := keeper.GetParams(ctx)
		nwParams := chainParams.Metadata["params"]
		if nwParams == "" {
			return nil, fmt.Errorf("params are required")
		}
		for _, tx := range req.Batch.Txs {
			err := s.processSourceTxV2(ctx, keeper, chain.Name, tx, block, nwParams)
			if err != nil {
				s.Logger(ctx).Error("failed to process source tx", "error", err)
			}
		}
		s.Logger(ctx).Info("ConfirmSourceTxsV2", "time", time.Since(start), "number of txs", len(req.Batch.Txs))
		return &types.ConfirmSourceTxsResponseV2{}, nil
	}

	// Block not found, handle pending confirm request
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

	event := &types.ConfirmNewBlockStarted{
		Chain:              chain.Name,
		BlockHash:          batch.BlockHash,
		PreviousBlockHash:  nil,
		ConfirmationHeight: keeper.GetRequiredConfirmationHeight(ctx),
		PollParticipants:   pollParticipants,
	}
	s.Logger(ctx).Info("Emit event ConfirmNewBlockStarted", "block_hash", batch.BlockHash.Hex())
	clog.Greenf("ConfirmNewBlockStarted: %+v", event)

	events.Emit(ctx, event)
	return nil
}

func validateTxProof(txId []byte, txIndex uint64, merklePath []exported.Hash, blockMerkleRoot exported.Hash) error {
	merkleRoot := btc.GetMerkleRootFromPath(txId, txIndex, slices.Map(merklePath, func(p exported.Hash) []byte {
		return p.Bytes()
	}), true)

	if !bytes.Equal(merkleRoot, blockMerkleRoot.Bytes()) {
		return fmt.Errorf("merkle root mismatch: %s != %s", merkleRoot, blockMerkleRoot.Bytes())
	}

	return nil
}

// processSourceTx handles validation and processing of a single source transaction.
func (s msgServer) processSourceTxV2(ctx sdk.Context, keeper types.ChainKeeper, chainName nexus.ChainName, tx *types.TrustedTx, block *types.BlockMetadata, nwParams string) error {
	txInfo, err := btc.ParseTx(tx.Raw)
	if err != nil {
		s.Logger(ctx).Error("failed to parse tx", "error", err)
		return err
	}

	if txInfo.MsgTx.TxHash().String() != tx.Hash.String() {
		s.Logger(ctx).Error("tx hash mismatch", "expected", tx.Hash.String(), "actual", txInfo.MsgTx.TxHash().String())
		return fmt.Errorf("tx hash mismatch: %s != %s", tx.Hash.String(), txInfo.MsgTx.TxHash().String())
	}

	err = validateTxProof(txInfo.TxID, tx.TxIndex, tx.MerklePath, block.MerkleRoot)
	if err != nil {
		s.Logger(ctx).Error("failed to validate tx proof", "error", err)
		return err
	}

	protocolInfo, err := s.protocol.FindProtocolInfoByInternalAddress(ctx, chainName, nexus.ChainName(txInfo.DestinationChain), txInfo.DestinationTokenAddress)
	if err != nil {
		s.Logger(ctx).Error("failed to find protocol info by internal address", "error", err)
		return err
	}

	sender, err := btc_utils.ScriptPubKeyToAddress(tx.PrevOutpointScriptPubkey, nwParams)
	if err != nil {
		s.Logger(ctx).Error("Failed to get sender address", "error", err)
		return err
	}

	err = keeper.ProcessConfirmRequestTx(ctx, txInfo, tx.TxIndex, block.Height, protocolInfo.Symbol, sender.String())
	if err != nil {
		s.Logger(ctx).Error("failed to process confirm request batch", "error", err)
		return err
	}

	return nil
}
