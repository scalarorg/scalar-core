package keeper

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/bitcoin-vault/ffi/go-vault"
	btc_utils "github.com/scalarorg/bitcoin-vault/go-utils/btc"
	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	go_utils "github.com/scalarorg/bitcoin-vault/go-utils/types"
	"github.com/scalarorg/scalar-core/utils/btc"
	"github.com/scalarorg/scalar-core/utils/clog"
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
	if err != nil {
		return nil, err
	}

	for _, tx := range req.Batch.Txs {
		if err := s.validateAndSaveTokenSent(ctx, keeper, chain.Name, tx, block); err != nil {
			clog.Redf("Failed to validate and save token sent: %v", err)
			continue
		}
	}

	return &types.ConfirmSourceTxsResponseV2{}, nil
}

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
		TxID:  tx.Hash,
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
