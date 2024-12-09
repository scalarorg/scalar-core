package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/axelarnetwork/axelar-core/utils"
	"github.com/axelarnetwork/axelar-core/utils/key"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	"github.com/axelarnetwork/utils/funcs"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/scalarorg/scalar-core/x/btc/types"
)

var (
	commandQueueName   = "command_queue"
	commandBatchPrefix = "batched_commands"

	unsignedBatchIDKey     = key.FromStr("unsigned_command_batch_id")
	latestSignedBatchIDKey = key.FromStr("latest_signed_command_batch_id")

	eventPrefix = utils.KeyFromStr("event")

	confirmedEventQueueName = "confirmed_event_queue"

	destinationRecipientAddressPrefix = key.RegisterStaticKey(types.ModuleName+types.ChainNamespace, 1)
	confirmedStakingTxPrefix          = key.RegisterStaticKey(types.ModuleName+types.ChainNamespace, 2)
	completedStakingTxPrefix          = key.RegisterStaticKey(types.ModuleName+types.ChainNamespace, 3)
)

var _ types.ChainKeeper = chainKeeper{}

type chainKeeper struct {
	internalKeeper
	chain nexus.ChainName
}

func (k chainKeeper) GetName() nexus.ChainName {
	return k.chain
}

// GetParams gets the evm module's parameters
func (k chainKeeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.getSubspace().GetParamSet(ctx, &p)
	return p
}

func (k chainKeeper) getConfirmedStakingTxs(ctx sdk.Context) []types.StakingTx {

	var stakingTxs []types.StakingTx
	iter := k.getStore(ctx).IteratorNew(confirmedStakingTxPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))

	for ; iter.Valid(); iter.Next() {
		var stakingTx types.StakingTx
		iter.UnmarshalValue(&stakingTx)
		stakingTxs = append(stakingTxs, stakingTx)
	}

	return stakingTxs
}

func (k chainKeeper) getCommandBatchesMetadata(ctx sdk.Context) []types.CommandBatchMetadata {
	iter := k.getStore(ctx).Iterator(utils.KeyFromStr(commandBatchPrefix))
	defer utils.CloseLogError(iter, k.Logger(ctx))

	var batches []types.CommandBatchMetadata
	for ; iter.Valid(); iter.Next() {
		var batch types.CommandBatchMetadata
		iter.UnmarshalValue(&batch)
		batches = append(batches, batch)
	}

	return batches
}

func (k chainKeeper) getEvents(ctx sdk.Context) []types.Event {
	iter := k.getStore(ctx).Iterator(eventPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))

	var events []types.Event
	for ; iter.Valid(); iter.Next() {
		var event types.Event
		iter.UnmarshalValue(&event)
		events = append(events, event)
	}

	return events
}

// func (k chainKeeper) GetChainID(ctx sdk.Context) (sdk.Int, bool) {
// 	network := k.GetNetwork(ctx)
// 	return k.GetChainIDByNetwork(ctx, network)
// }

// // GetNetwork returns the EVM network Axelar-Core is expected to connect to
// func (k chainKeeper) GetNetwork(ctx sdk.Context) string {
// 	return getParam[string](k, ctx, types.KeyNetwork)
// }

func (k chainKeeper) GetRequiredConfirmationHeight(ctx sdk.Context) uint64 {
	return getParam[uint64](k, ctx, types.KeyConfirmationHeight)
}

func getParam[T any](k chainKeeper, ctx sdk.Context, paramKey []byte) T {
	var value T
	k.getSubspace().Get(ctx, paramKey, &value)
	return value
}

func (k chainKeeper) getSubspace() params.Subspace {
	chainKey := key.FromStr(types.ModuleName).Append(key.From(k.chain))
	subspace, ok := k.paramsKeeper.GetSubspace(chainKey.String())
	if !ok {
		panic(fmt.Sprintf("subspace for chain %s does not exist", k.chain))
	}
	return subspace
}

func (k chainKeeper) getStore(ctx sdk.Context) utils.KVStore {
	pre := string(chainPrefix.Append(utils.LowerCaseKey(k.chain.String())).AsKey()) + "_"
	return utils.NewNormalizedStore(prefix.NewStore(ctx.KVStore(k.storeKey), []byte(pre)), k.cdc)
}

func (k chainKeeper) validateCommandQueueState(state utils.QueueState, queueName ...string) error {
	if err := state.ValidateBasic(queueName...); err != nil {
		return err
	}

	for _, item := range state.Items {
		var command types.Command
		if err := k.cdc.UnmarshalLengthPrefixed(item.Value, &command); err != nil {
			return err
		}

		if err := command.KeyID.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}

func (k chainKeeper) getCommandQueue(ctx sdk.Context) utils.BlockHeightKVQueue {
	return utils.NewBlockHeightKVQueue(
		commandQueueName,
		k.getStore(ctx),
		ctx.BlockHeight(),
		k.Logger(ctx),
	)
}

func (k chainKeeper) SetStakingTx(ctx sdk.Context, stakingTx types.StakingTx, state types.StakingTxStatus) {
	switch state {
	case types.StakingTxStatus_Confirmed:
		funcs.MustNoErr(
			k.getStore(ctx).SetNewValidated(
				confirmedStakingTxPrefix.Append(key.FromStr(stakingTx.TxID.HexStr())).Append(key.FromUInt(stakingTx.LogIndex)), &stakingTx))
	case types.StakingTxStatus_Completed:
		funcs.MustNoErr(
			k.getStore(ctx).SetNewValidated(
				completedStakingTxPrefix.Append(key.FromStr(stakingTx.TxID.HexStr())).Append(key.FromUInt(stakingTx.LogIndex)), &stakingTx))
	default:
		panic("invalid deposit state")
	}
}

func (k chainKeeper) setCommandBatchMetadata(ctx sdk.Context, meta types.CommandBatchMetadata) {
	funcs.MustNoErr(
		k.getStore(ctx).SetNewValidated(key.FromStr(commandBatchPrefix).Append(key.FromBz(meta.ID)), &meta))
}

func (k chainKeeper) setUnsignedCommandBatchID(ctx sdk.Context, id []byte) {
	k.getStore(ctx).SetRawNew(unsignedBatchIDKey, id)
}

func (k chainKeeper) SetLatestSignedCommandBatchID(ctx sdk.Context, id []byte) {
	k.getStore(ctx).SetRawNew(latestSignedBatchIDKey, id)
}

func (k chainKeeper) setLatestBatchMetadata(ctx sdk.Context, batch types.CommandBatchMetadata) {
	switch batch.Status {
	case types.BatchNonExistent:
		return
	case types.BatchSigning, types.BatchAborted:
		k.setUnsignedCommandBatchID(ctx, batch.ID)
	case types.BatchSigned:
		k.SetLatestSignedCommandBatchID(ctx, batch.ID)
	default:
		panic(fmt.Sprintf("batch status %s is not handled", batch.Status.String()))
	}
}

func getEventKey(eventID types.EventID) utils.Key {
	return eventPrefix.Append(utils.LowerCaseKey(string(eventID)))
}

func (k chainKeeper) setEvent(ctx sdk.Context, event types.Event) {
	funcs.MustNoErr(
		k.getStore(ctx).SetNewValidated(key.FromBz(getEventKey(event.GetID()).AsKey()), &event))
}

// validateConfirmedEventQueueState checks if the keys of the given map have the correct format to be imported as confirmed event state.
func (k chainKeeper) validateConfirmedEventQueueState(state utils.QueueState, queueName ...string) error {
	if err := state.ValidateBasic(queueName...); err != nil {
		return err
	}

	for _, item := range state.Items {
		var event types.Event
		if err := k.cdc.UnmarshalLengthPrefixed(item.Value, &event); err != nil {
			return err
		}

		if err := event.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}

// GetConfirmedEventQueue returns a queue of all the confirmed events
func (k chainKeeper) GetConfirmedEventQueue(ctx sdk.Context) utils.KVQueue {
	blockHeightBz := make([]byte, 8)
	binary.BigEndian.PutUint64(blockHeightBz, uint64(ctx.BlockHeight()))

	return utils.NewGeneralKVQueue(
		confirmedEventQueueName,
		k.getStore(ctx),
		k.Logger(ctx),
		func(value codec.ProtoMarshaler) utils.Key {
			event := value.(*types.Event)

			indexBz := make([]byte, 8)
			binary.BigEndian.PutUint64(indexBz, event.Index)

			return utils.KeyFromBz(blockHeightBz).
				Append(utils.KeyFromBz(event.TxID.Bytes())).
				Append(utils.KeyFromBz(indexBz))
		},
	)
}
