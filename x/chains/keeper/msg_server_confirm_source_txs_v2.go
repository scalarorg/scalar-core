package keeper

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	sdk "github.com/cosmos/cosmos-sdk/types"
	vault "github.com/scalarorg/bitcoin-vault/ffi/go"
	btc_utils "github.com/scalarorg/go-common/btc"
	"github.com/scalarorg/go-common/chain"
	go_utils "github.com/scalarorg/go-common/types"
	"github.com/scalarorg/scalar-core/utils/btc"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/utils/slices"
	btcVald "github.com/scalarorg/scalar-core/vald/xchain/btc"
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
			txHashBytes := btc.DoubleSha256(tx.Raw)
			clog.Greenf("Calculated tx hash: %s", hex.EncodeToString(txHashBytes))
			clog.Greenf("Input tx hash: %s", hex.EncodeToString(tx.Hash[:]))
			clog.Greenf("Block merkle root: %s", block.MerkleRoot.Hex())
			err = validateTxProof(txHashBytes, tx.TxIndex, tx.MerklePath, block.MerkleRoot)
			if err != nil {
				s.Logger(ctx).Error("failed to validate tx proof", "error", err)
				continue
			}
			protocolInfo, err := s.protocol.FindProtocolInfoByInternalAddress(ctx, chain.Name, nexus.ChainName(txInfo.DestinationChain), txInfo.DestinationTokenAddress)
			if err != nil {
				s.Logger(ctx).Error("failed to find protocol info by internal address", "error", err)
				continue
			}

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

func validateTxProof(txId []byte, txIndex uint64, merklePath []exported.Hash, blockMerkleRoot exported.Hash) error {
	merkleRoot := btc.GetMerkleRootFromPath(txId, txIndex, slices.Map(merklePath, func(p exported.Hash) []byte {
		return p.Bytes()
	}), true)

	if !bytes.Equal(merkleRoot, blockMerkleRoot.Bytes()) {
		return fmt.Errorf("merkle root mismatch: %s != %s", merkleRoot, blockMerkleRoot.Bytes())
	}

	return nil
}

// TODO: remove this function
func (s msgServer) validateAndSaveTokenSent(ctx sdk.Context, keeper types.ChainKeeper, chain nexus.ChainName, tx *types.TrustedTx, block *types.BlockMetadata) error {
	reader := bytes.NewReader(tx.Raw)
	var msgTx wire.MsgTx
	if err := msgTx.Deserialize(reader); err != nil {
		log.Fatalf("Failed to deserialize tx: %v", err)
	}

	msgTxHash := msgTx.TxHash()

	// in reverse byte order in case of btc
	if msgTxHash.String() != tx.Hash.String() {
		return fmt.Errorf("tx hash mismatch: %s != %s", msgTx.TxHash(), tx.Hash)
	}

	merkleRoot := btc.GetMerkleRootFromPath(msgTxHash.CloneBytes(), tx.TxIndex, slices.Map(tx.MerklePath, func(p exported.Hash) []byte {
		return p.Bytes()
	}), true)

	if !bytes.Equal(merkleRoot, block.MerkleRoot.Bytes()) {
		return fmt.Errorf("merkle root mismatch: %s != %s", merkleRoot, block.MerkleRoot.Bytes())
	}

	// TODO: validate block_hash_chain

	// TODO: add checking when integrate with evm

	chainParams := keeper.GetParams(ctx)

	nwParams := chainParams.Metadata["params"]
	if nwParams == "" {
		return fmt.Errorf("params are required")
	}

	sender, err := btc_utils.ScriptPubKeyToAddress(tx.PrevOutpointScriptPubkey, nwParams)
	if err != nil {
		return err
	}

	tokenSent, err := s.createEventTokenSent(ctx, chain, &msgTx, tx.TxIndex, sender.String(), block.Height)
	if err != nil {
		return err
	}

	event := types.Event{
		Chain: chain,
		Hash:  tx.Hash,
		Event: &types.Event_TokenSent{
			TokenSent: tokenSent,
		},
		Index: uint64(tx.TxIndex),
	}

	if err := keeper.SetConfirmedEvent(ctx, event); err != nil {
		return err
	}

	keeper.EnqueueConfirmedEvent(ctx, event.GetID())

	clog.Greenf("Confirmed event: %s", event.GetID())

	return nil
}

func (s msgServer) createEventTokenSent(ctx sdk.Context, eventChain nexus.ChainName, tx *wire.MsgTx, txIndex uint64, sender string, blockHeight uint64) (*types.EventTokenSent, error) {
	if tx == nil {
		return nil, fmt.Errorf("tx is nil")
	}

	if len(tx.TxOut) < btcVald.MinNumberOfOutputs {
		return nil, btcVald.ErrInvalidTxOutCount
	}

	txHash := tx.TxHash()
	txId := exported.HashFromBytes(txHash.CloneBytes())

	eventId := types.NewEventID(txId, txIndex)

	embeddedDataTxOut := tx.TxOut[btcVald.EmbeddedDataOutputIndex]
	if embeddedDataTxOut == nil || embeddedDataTxOut.PkScript == nil || embeddedDataTxOut.PkScript[0] != txscript.OP_RETURN {
		return nil, btcVald.ErrInvalidOpReturn
	}

	output, err := vault.ParseVaultEmbeddedData(embeddedDataTxOut.PkScript)
	if err != nil || output == nil {
		return nil, btcVald.ErrInvalidOpReturnData
	}

	if output.TransactionType != go_utils.TransactionTypeLocking {
		return nil, btcVald.ErrInvalidTransactionType
	}

	var stakingAmount int64 = tx.TxOut[btcVald.LockingOutputIndex].Value
	var scriptPubkey []byte = tx.TxOut[btcVald.LockingOutputIndex].PkScript
	destinationChain := chain.NewChainInfoFromBytes(output.DestinationChain)
	if destinationChain == nil {
		return nil, btcVald.ErrInvalidDestinationChain
	}

	var destinationRecipientAddress types.Address
	err = destinationRecipientAddress.Unmarshal(output.DestinationRecipientAddress)
	if err != nil {
		return nil, err
	}

	protocol, err := s.protocol.FindProtocolInfoByInternalAddress(ctx, eventChain, nexus.ChainName(destinationChain.ToBytes().String()), hex.EncodeToString(output.DestinationTokenAddress))
	if err != nil {
		return nil, err
	}

	return &types.EventTokenSent{
		EventID:            eventId,
		Sender:             sender,
		Chain:              eventChain,
		TransferID:         nexus.TransferID(1),
		DestinationChain:   nexus.ChainName(destinationChain.ToBytes().String()),
		DestinationAddress: types.Address(destinationRecipientAddress).Hex(),
		Asset:              sdk.NewCoin(protocol.Symbol, sdk.NewInt(stakingAmount)),
		ScriptPubkey:       scriptPubkey,
		Vout:               uint32(btcVald.LockingOutputIndex),
		BlockHeight:        blockHeight,
	}, nil
}

// ProcessConfirmedBlock handles the confirmation of a block and processes the stored batch
// func (s msgServer) ProcessConfirmedBlock(ctx sdk.Context, keeper types.ChainKeeper, chain nexus.ChainName, blockHash exported.Hash) error {
// 	// Find all pending confirm requests for this block hash
// 	batches := keeper.FindPendingConfirmRequestsByBlockHash(ctx, blockHash)
// 	if len(batches) == 0 {
// 		s.Logger(ctx).Info("No pending batches found for block", "block_hash", blockHash)
// 		return nil
// 	}

// 	s.Logger(ctx).Info("Processing confirmed block batches", "block_hash", blockHash, "batch_count", len(batches))

// 	// Process all batches for this block
// 	start := time.Now()
// 	for pollID, batch := range batches {
// 		if err := keeper.ProcessConfirmRequestBatch(ctx, batch, pollID); err != nil {
// 			s.Logger(ctx).Error("Failed to process confirm request batch", "error", err, "block_hash", batch.BlockHash, "poll_id", pollID)
// 			continue
// 		}
// 	}

// 	s.Logger(ctx).Info("Processed confirmed block batches",
// 		"time", time.Since(start),
// 		"batch_count", len(batches),
// 		"block_hash", blockHash)

// 	return nil
// }
