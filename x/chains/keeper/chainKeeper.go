package keeper

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	chainsTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/events"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/key"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	protocol "github.com/scalarorg/scalar-core/x/protocol/exported"
)

var (
	gatewayKey                       = key.FromStr("gateway")
	unsignedBatchIDKey               = key.FromStr("unsigned_command_batch_id")
	latestSignedBatchIDKey           = key.FromStr("latest_signed_command_batch_id")
	tokenMetadataByAssetPrefix       = "token_deployment_by_asset"
	tokenMetadataBySymbolPrefix      = key.FromStr("token_deployment_by_symbol")
	confirmedDepositPrefixDeprecated = "confirmed_deposit" // Deprecated
	burnedDepositPrefixDeprecated    = "burned_deposit"    // Deprecated
	commandBatchPrefix               = "batched_commands"
	commandPrefix                    = "command"
	blockPrefix                      = "block"
	eventPrefix                      = utils.KeyFromStr("event")
	confirmedEventQueueName          = "confirmed_event_queue"
	commandQueueName                 = "cmd_queue"

	burnerAddrPrefix       = key.RegisterStaticKey(types.ModuleName+types.ChainNamespace, 1)
	confirmedDepositPrefix = key.RegisterStaticKey(types.ModuleName+types.ChainNamespace, 2)
	burnedDepositPrefix    = key.RegisterStaticKey(types.ModuleName+types.ChainNamespace, 3)

	confirmedSourceTxPrefix = key.RegisterStaticKey(types.ModuleName+types.ChainNamespace, 4)
	completedSourceTxPrefix = key.RegisterStaticKey(types.ModuleName+types.ChainNamespace, 5)
	redeemSessionPrefix     = key.RegisterStaticKey(types.ModuleName+types.ChainNamespace, 6)
)

var _ types.ChainKeeper = ChainKeeper{}

type ChainKeeper struct {
	internalKeeper
	chain nexus.ChainName
}

func (k ChainKeeper) SetDeposit(ctx sdk.Context, deposit types.ERC20Deposit, state types.DepositStatus) {
	switch state {
	case types.DepositStatus_Confirmed:
		funcs.MustNoErr(
			k.getStore(ctx).SetNewValidated(
				confirmedDepositPrefix.Append(key.FromStr(deposit.TxID.Hex())).Append(key.FromUInt(deposit.LogIndex)), &deposit))
	case types.DepositStatus_Burned:
		funcs.MustNoErr(
			k.getStore(ctx).SetNewValidated(
				burnedDepositPrefix.Append(key.FromStr(deposit.TxID.Hex())).Append(key.FromUInt(deposit.LogIndex)), &deposit))
	default:
		panic("invalid deposit state")
	}
}

func (k ChainKeeper) SetBurnerInfo(ctx sdk.Context, burnerInfo types.BurnerInfo) {
	funcs.MustNoErr(
		k.getStore(ctx).SetNewValidated(burnerAddrPrefix.Append(key.FromStr(burnerInfo.BurnerAddress.Hex())), &burnerInfo))
}

func (k ChainKeeper) GetVotingThreshold(ctx sdk.Context) utils.Threshold {
	return getParam[utils.Threshold](k, ctx, types.KeyVotingThreshold)
}

func (k ChainKeeper) GetRevoteLockingPeriod(ctx sdk.Context) int64 {
	return getParam[int64](k, ctx, types.KeyRevoteLockingPeriod)
}

func (k ChainKeeper) GetPendingCommands(ctx sdk.Context) []types.Command {
	var commands []types.Command
	keys := k.getCommandQueue(ctx).Keys()
	for _, queueKey := range keys {
		var cmd types.Command
		ok := k.getStore(ctx).GetNew(key.FromBz(queueKey.AsKey()), &cmd)
		if ok {
			commands = append(commands, cmd)
		}
	}

	return commands
}

func (k ChainKeeper) GetMinVoterCount(ctx sdk.Context) int64 {
	return getParam[int64](k, ctx, types.KeyMinVoterCount)
}

func (k ChainKeeper) GetDepositsByTxID(ctx sdk.Context, txID exported.Hash, status types.DepositStatus) ([]types.ERC20Deposit, error) {
	var prefix key.Key
	switch status {
	case types.DepositStatus_Confirmed:
		prefix = confirmedDepositPrefix
	case types.DepositStatus_Burned:
		prefix = burnedDepositPrefix
	default:
		return nil, fmt.Errorf("unsupported deposit status %s", status.String())
	}

	iter := k.getStore(ctx).IteratorNew(prefix.Append(key.FromStr(txID.Hex())))
	defer utils.CloseLogError(iter, k.Logger(ctx))

	var deposits []types.ERC20Deposit
	for ; iter.Valid(); iter.Next() {
		var deposit types.ERC20Deposit
		iter.UnmarshalValue(&deposit)

		deposits = append(deposits, deposit)
	}

	return deposits, nil
}

func (k ChainKeeper) GetDeposit(ctx sdk.Context, txID exported.Hash, logIndex uint64) (types.ERC20Deposit, types.DepositStatus, bool) {
	var deposit types.ERC20Deposit

	if k.getStore(ctx).GetNew(confirmedDepositPrefix.Append(key.FromStr(txID.Hex())).Append(key.FromUInt(logIndex)), &deposit) {
		return deposit, types.DepositStatus_Confirmed, true
	}
	if k.getStore(ctx).GetNew(burnedDepositPrefix.Append(key.FromStr(txID.Hex())).Append(key.FromUInt(logIndex)), &deposit) {
		return deposit, types.DepositStatus_Burned, true
	}

	return types.ERC20Deposit{}, 0, false
}

func (k ChainKeeper) GetConfirmedDepositsPaginated(ctx sdk.Context, pageRequest *query.PageRequest) ([]types.ERC20Deposit, *query.PageResponse, error) {
	var deposits []types.ERC20Deposit

	// TODO: refactor iteration over values using a prefix to avoid collisions
	resp, err := query.Paginate(prefix.NewStore(k.getStore(ctx).KVStore, append(confirmedDepositPrefix.Bytes(), []byte(key.DefaultDelimiter)...)), pageRequest, func(key []byte, value []byte) error {
		var deposit types.ERC20Deposit
		k.cdc.MustUnmarshalLengthPrefixed(value, &deposit)
		deposits = append(deposits, deposit)
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return deposits, resp, nil
}

func (k ChainKeeper) GetName() nexus.ChainName {
	return k.chain
}

// GetParams gets the evm module's parameters
func (k ChainKeeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.getSubspace().GetParamSet(ctx, &p)
	return p
}

func (k ChainKeeper) getCommandsGasLimit(ctx sdk.Context) uint32 {
	return getParam[uint32](k, ctx, types.KeyCommandsGasLimit)
}

func (k ChainKeeper) getConfirmedSourceTxs(ctx sdk.Context) []types.SourceTx {
	var sourceTxs []types.SourceTx
	iter := k.getStore(ctx).IteratorNew(confirmedSourceTxPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))

	for ; iter.Valid(); iter.Next() {
		var sourceTx types.SourceTx
		iter.UnmarshalValue(&sourceTx)
		sourceTxs = append(sourceTxs, sourceTx)
	}

	return sourceTxs
}

func (k ChainKeeper) getCommandBatchesMetadata(ctx sdk.Context) []types.CommandBatchMetadata {
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

func (k ChainKeeper) getEvents(ctx sdk.Context) []types.Event {
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

func (k ChainKeeper) GetChainID(ctx sdk.Context) (sdk.Int, bool) {
	chainId := getParam[sdk.Int](k, ctx, types.KeyChainID)
	if chainId.IsNil() {
		return sdk.Int{}, false
	}
	return chainId, true
}

func (k ChainKeeper) GetMetadata(ctx sdk.Context) map[string]string {
	return getParam[map[string]string](k, ctx, types.KeyMetadata)
}

func (k ChainKeeper) GetRequiredConfirmationHeight(ctx sdk.Context) uint64 {
	return getParam[uint64](k, ctx, types.KeyConfirmationHeight)
}

func (k ChainKeeper) SetBlock(ctx sdk.Context, meta types.BlockMetadata) {
	newMeta := types.BlockMetadata{
		BlockHash:  meta.BlockHash,
		MerkleRoot: meta.MerkleRoot,
		Height:     meta.Height,
	}

	funcs.MustNoErr(
		k.getStore(ctx).SetNewValidated(key.FromStr(blockPrefix).Append(key.FromUInt(meta.Height)), &newMeta))
}

func (k ChainKeeper) GetCurrentBlock(ctx sdk.Context) (*types.BlockMetadata, error) {
	iter := k.getStore(ctx).ReverseIterator(utils.KeyFromStr(blockPrefix))
	defer utils.CloseLogError(iter, k.Logger(ctx))

	if !iter.Valid() {
		return nil, fmt.Errorf("no block metadata found")
	}

	var current types.BlockMetadata
	iter.UnmarshalValue(&current)

	if iter.Next(); iter.Valid() {
		var prev types.BlockMetadata
		iter.UnmarshalValue(&prev)
		current.PreviousBlockHash = &prev.BlockHash
	}

	return &current, nil
}

func (k ChainKeeper) GetBlock(ctx sdk.Context, hash exported.Hash) (*types.BlockMetadata, error) {
	iter := k.getStore(ctx).ReverseIterator(utils.KeyFromStr(blockPrefix))
	defer utils.CloseLogError(iter, k.Logger(ctx))

	if !iter.Valid() {
		return nil, fmt.Errorf("no block metadata found")
	}

	for ; iter.Valid(); iter.Next() {
		var block types.BlockMetadata
		iter.UnmarshalValue(&block)
		if block.BlockHash == hash {
			return &block, nil
		}
	}

	return nil, fmt.Errorf("block metadata not found")
}

func getParam[T any](k ChainKeeper, ctx sdk.Context, paramKey []byte) T {
	var value T
	k.getSubspace().Get(ctx, paramKey, &value)
	return value
}

func (k ChainKeeper) getSubspace() params.Subspace {
	chainKey := key.FromStr(types.ModuleName).Append(key.From(k.chain))
	subspace, ok := k.paramsKeeper.GetSubspace(chainKey.String())
	if !ok {
		panic(fmt.Sprintf("subspace for chain %s does not exist", k.chain))
	}
	return subspace
}

func (k ChainKeeper) getStore(ctx sdk.Context) utils.KVStore {
	pre := string(chainPrefix.Append(utils.LowerCaseKey(k.chain.String())).AsKey()) + "_"
	return utils.NewNormalizedStore(prefix.NewStore(ctx.KVStore(k.storeKey), []byte(pre)), k.cdc)
}

func (k ChainKeeper) validateCommandQueueState(state utils.QueueState, queueName ...string) error {
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

func (k ChainKeeper) getCommandQueue(ctx sdk.Context) utils.BlockHeightKVQueue {
	return utils.NewBlockHeightKVQueue(
		commandQueueName,
		k.getStore(ctx),
		ctx.BlockHeight(),
		k.Logger(ctx),
	)
}

func (k ChainKeeper) SetSourceTx(ctx sdk.Context, sourceTx types.SourceTx, state types.SourceTxStatus) {
	switch state {
	case types.SourceTxStatus_Confirmed:
		funcs.MustNoErr(
			k.getStore(ctx).SetNewValidated(
				confirmedSourceTxPrefix.Append(key.FromStr(sourceTx.TxID.Hex())).Append(key.FromUInt(sourceTx.LogIndex)), &sourceTx))
	case types.SourceTxStatus_Completed:
		funcs.MustNoErr(
			k.getStore(ctx).SetNewValidated(
				completedSourceTxPrefix.Append(key.FromStr(sourceTx.TxID.Hex())).Append(key.FromUInt(sourceTx.LogIndex)), &sourceTx))
	default:
		panic("invalid source tx state")
	}
}

func (k ChainKeeper) setCommandBatchMetadata(ctx sdk.Context, meta types.CommandBatchMetadata) {
	funcs.MustNoErr(
		k.getStore(ctx).SetNewValidated(key.FromStr(commandBatchPrefix).Append(key.FromBz(meta.ID)), &meta))
}

func (k ChainKeeper) setUnsignedCommandBatchID(ctx sdk.Context, id []byte) {
	k.getStore(ctx).SetRawNew(unsignedBatchIDKey, id)
}

func (k ChainKeeper) SetLatestSignedCommandBatchID(ctx sdk.Context, id []byte) {
	k.getStore(ctx).SetRawNew(latestSignedBatchIDKey, id)
}

func (k ChainKeeper) setLatestBatchMetadata(ctx sdk.Context, batch types.CommandBatchMetadata) {
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

func (k ChainKeeper) setEvent(ctx sdk.Context, event types.Event) {
	funcs.MustNoErr(
		k.getStore(ctx).SetNewValidated(key.FromBz(getEventKey(event.GetID()).AsKey()), &event))
}

// validateConfirmedEventQueueState checks if the keys of the given map have the correct format to be imported as confirmed event state.
func (k ChainKeeper) validateConfirmedEventQueueState(state utils.QueueState, queueName ...string) error {
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
func (k ChainKeeper) GetConfirmedEventQueue(ctx sdk.Context) utils.KVQueue {
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

// GetEvent returns the event for the given event ID
func (k ChainKeeper) GetEvent(ctx sdk.Context, eventID types.EventID) (event types.Event, ok bool) {
	k.getStore(ctx).GetNew(key.FromBz(getEventKey(eventID).AsKey()), &event)

	return event, event.Status != types.EventNonExistent
}

func (k ChainKeeper) SetConfirmedEvent(ctx sdk.Context, event types.Event) error {
	eventID := event.GetID()
	if _, ok := k.GetEvent(ctx, eventID); ok {
		return fmt.Errorf("event %s is already confirmed", eventID)
	}

	event.Status = types.EventConfirmed
	k.setEvent(ctx, event)

	events.Emit(ctx, &types.ChainEventConfirmed{
		Chain:   event.Chain,
		EventID: event.GetID(),
		Type:    event.GetEventType(),
	})

	return nil
}

// SetEventCompleted sets the event as completed
func (k ChainKeeper) SetEventCompleted(ctx sdk.Context, eventID types.EventID) error {
	event, ok := k.GetEvent(ctx, eventID)
	if !ok || event.Status != types.EventConfirmed {
		return fmt.Errorf("event %s is not confirmed", eventID)
	}

	event.Status = types.EventCompleted
	k.setEvent(ctx, event)

	events.Emit(ctx,
		&types.ChainEventCompleted{
			Chain:   event.Chain,
			EventID: event.GetID(),
			Type:    event.GetEventType(),
		})

	return nil
}

// SetEventFailed sets the event as invalid
func (k ChainKeeper) SetEventFailed(ctx sdk.Context, eventID types.EventID) error {
	event, ok := k.GetEvent(ctx, eventID)
	if !ok || event.Status != types.EventConfirmed {
		return fmt.Errorf("event %s is not confirmed", eventID)
	}

	event.Status = types.EventFailed
	k.setEvent(ctx, event)

	k.Logger(ctx).Debug("failed handling event",
		"chain", event.Chain,
		"eventID", event.GetID(),
	)

	events.Emit(ctx,
		&types.ChainEventFailed{
			Chain:   event.Chain,
			EventID: event.GetID(),
			Type:    event.GetEventType(),
		})

	return nil
}

func (k ChainKeeper) EnqueueCommand(ctx sdk.Context, command types.Command) error {
	if k.getStore(ctx).HasNew(key.FromStr(commandPrefix).Append(key.FromStr(command.ID.Hex()))) {
		return fmt.Errorf("command %s already exists", command.ID.Hex())
	}

	k.getCommandQueue(ctx).Enqueue(utils.LowerCaseKey(commandPrefix).AppendStr(command.ID.Hex()), &command)
	return nil
}

func (k ChainKeeper) CreateERC20Token(ctx sdk.Context, asset string, details nexus.TokenDetails, address types.Address) (types.ERC20Token, error) {
	metadata, err := k.initTokenMetadata(ctx, asset, details, address)
	if err != nil {
		return types.NilToken, err
	}

	return types.CreateERC20Token(func(m types.ERC20TokenMetadata) {
		k.setTokenMetadata(ctx, m)
	}, metadata), nil
}

func (k ChainKeeper) initTokenMetadata(ctx sdk.Context, asset string, details nexus.TokenDetails, address types.Address) (types.ERC20TokenMetadata, error) {
	if err := details.Validate(); err != nil {
		return types.ERC20TokenMetadata{}, err
	}

	// perform a few checks now, so that it is impossible to get errors later
	if token := k.GetERC20TokenByAsset(ctx, asset); !token.Is(types.NonExistent) {
		return types.ERC20TokenMetadata{}, fmt.Errorf("token for asset '%s' already set", asset)
	}

	if token := k.GetERC20TokenBySymbol(ctx, details.Symbol); !token.Is(types.NonExistent) {
		return types.ERC20TokenMetadata{}, fmt.Errorf("token with symbol '%s' already set", details.Symbol)
	}

	chainID := k.getSigner(ctx).ChainID()

	burnerCode := k.GetBurnerByteCode(ctx)

	if !address.IsZeroAddress() {
		meta := types.ERC20TokenMetadata{
			Asset:        asset,
			Details:      details,
			TokenAddress: address,
			ChainID:      sdk.NewIntFromBigInt(chainID),
			Status:       types.Initialized,
			IsExternal:   true,
			BurnerCode:   nil,
		}
		k.setTokenMetadata(ctx, meta)

		return meta, nil
	}

	gatewayAddr, found := k.GetGatewayAddress(ctx)
	if !found {
		return types.ERC20TokenMetadata{}, fmt.Errorf("scalar gateway address for chain '%s' not set", k.chain)
	}

	tokenAddr, err := k.getTokenAddress(ctx, details, gatewayAddr)
	if err != nil {
		return types.ERC20TokenMetadata{}, err
	}

	// all good
	meta := types.ERC20TokenMetadata{
		Asset:        asset,
		Details:      details,
		TokenAddress: tokenAddr,
		ChainID:      sdk.NewIntFromBigInt(chainID),
		Status:       types.Initialized,
		IsExternal:   false,
		BurnerCode:   burnerCode,
	}
	k.setTokenMetadata(ctx, meta)

	return meta, nil
}

func (k ChainKeeper) getTokenAddress(ctx sdk.Context, details nexus.TokenDetails, gatewayAddr types.Address) (types.Address, error) {
	var saltToken [32]byte
	copy(saltToken[:], crypto.Keccak256Hash([]byte(details.Symbol)).Bytes())

	uint8Type, err := abi.NewType("uint8", "uint8", nil)
	if err != nil {
		return types.Address{}, err
	}

	uint256Type, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return types.Address{}, err
	}

	stringType, err := abi.NewType("string", "string", nil)
	if err != nil {
		return types.Address{}, err
	}

	arguments := abi.Arguments{{Type: stringType}, {Type: stringType}, {Type: uint8Type}, {Type: uint256Type}}
	packed, err := arguments.Pack(details.TokenName, details.Symbol, details.Decimals, details.Capacity.BigInt())
	if err != nil {
		return types.Address{}, err
	}

	bytecode := k.GetTokenByteCode(ctx)
	tokenInitCode := append(bytecode, packed...)
	tokenInitCodeHash := crypto.Keccak256Hash(tokenInitCode)

	tokenAddr := types.Address(crypto.CreateAddress2(common.Address(gatewayAddr), saltToken, tokenInitCodeHash.Bytes()))
	return tokenAddr, nil
}

func (k ChainKeeper) GetTokenByteCode(ctx sdk.Context) []byte {
	return getParam[[]byte](k, ctx, types.KeyToken)
}

func (k ChainKeeper) GetGatewayAddress(ctx sdk.Context) (types.Address, bool) {
	if gateway := k.getGateway(ctx); !gateway.Address.IsZeroAddress() {
		return gateway.Address, true
	}

	return types.Address{}, false
}

func (k ChainKeeper) getGateway(ctx sdk.Context) types.Gateway {
	var gateway types.Gateway
	k.getStore(ctx).GetNew(gatewayKey, &gateway)

	return gateway
}

func (k ChainKeeper) GetBurnerByteCode(ctx sdk.Context) []byte {
	return getParam[[]byte](k, ctx, types.KeyBurnable)
}

func (k ChainKeeper) getSigner(ctx sdk.Context) chainsTypes.EIP155Signer {

	chainID, found := k.GetChainID(ctx)

	// both chain, subspace, and network must be valid if the chain keeper was instantiated,
	// so a nil value here must be a catastrophic failure
	if !found {
		panic(fmt.Sprintf("could not find chain ID for network '%s'", chainID))
	}
	return chainsTypes.NewEIP155Signer(chainID.BigInt())
}

func (k ChainKeeper) setTokenMetadata(ctx sdk.Context, meta types.ERC20TokenMetadata) {
	// lookup by asset
	funcs.MustNoErr(
		k.getStore(ctx).SetNewValidated(key.FromStr(tokenMetadataByAssetPrefix).Append(key.FromStr(meta.Asset)), &meta))

	// lookup by symbol
	funcs.MustNoErr(
		k.getStore(ctx).SetNewValidated(tokenMetadataBySymbolPrefix.Append(key.FromStr(meta.Details.Symbol)), &meta))
}

func (k ChainKeeper) getTokenMetadataByAsset(ctx sdk.Context, asset string) (types.ERC20TokenMetadata, bool) {
	var result types.ERC20TokenMetadata
	found := k.getStore(ctx).GetNew(key.FromStr(tokenMetadataByAssetPrefix).Append(key.FromStr(asset)), &result)

	return result, found
}

func (k ChainKeeper) getTokenMetadataBySymbol(ctx sdk.Context, symbol string) (types.ERC20TokenMetadata, bool) {
	var result types.ERC20TokenMetadata
	found := k.getStore(ctx).GetNew(tokenMetadataBySymbolPrefix.Append(key.FromStr(symbol)), &result)

	return result, found
}

// GetERC20TokenByAddress finds a token's information by its address
func (k ChainKeeper) GetERC20TokenByAddress(ctx sdk.Context, address types.Address) types.ERC20Token {
	for _, tokenMetadata := range k.getTokensMetadata(ctx) {
		if tokenMetadata.TokenAddress == address {
			return types.CreateERC20Token(func(m types.ERC20TokenMetadata) {
				k.setTokenMetadata(ctx, m)
			}, tokenMetadata)
		}
	}

	return types.ERC20Token{}
}

func (k ChainKeeper) GetTokens(ctx sdk.Context) []types.ERC20Token {
	tokensMetadata := k.getTokensMetadata(ctx)
	tokens := make([]types.ERC20Token, len(tokensMetadata))

	for i, tokenMetadata := range tokensMetadata {
		tokens[i] = types.CreateERC20Token(func(m types.ERC20TokenMetadata) {
			k.setTokenMetadata(ctx, m)
		}, tokenMetadata)
	}

	return tokens
}

func (k ChainKeeper) getTokensMetadata(ctx sdk.Context) []types.ERC20TokenMetadata {
	iter := k.getStore(ctx).Iterator(utils.LowerCaseKey(tokenMetadataByAssetPrefix))
	defer utils.CloseLogError(iter, k.Logger(ctx))

	var tokens []types.ERC20TokenMetadata
	for ; iter.Valid(); iter.Next() {
		var token types.ERC20TokenMetadata
		iter.UnmarshalValue(&token)
		tokens = append(tokens, token)
	}
	return tokens
}

func (k ChainKeeper) GetERC20TokenByAsset(ctx sdk.Context, asset string) types.ERC20Token {
	metadata, ok := k.getTokenMetadataByAsset(ctx, asset)
	if !ok {
		return types.NilToken
	}

	return types.CreateERC20Token(func(m types.ERC20TokenMetadata) {
		k.setTokenMetadata(ctx, m)
	}, metadata)
}

// GetERC20TokenBySymbol returns the erc20 token by symbol
func (k ChainKeeper) GetERC20TokenBySymbol(ctx sdk.Context, symbol string) types.ERC20Token {
	metadata, ok := k.getTokenMetadataBySymbol(ctx, symbol)
	if !ok {
		return types.NilToken
	}

	return types.CreateERC20Token(func(m types.ERC20TokenMetadata) {
		k.setTokenMetadata(ctx, m)
	}, metadata)
}

// CreateNewBatchToSign creates a new batch of commands to be signed
func (k ChainKeeper) CreateNewBatchToSign(ctx sdk.Context) (types.CommandBatch, error) {
	chain := k.GetName()
	if types.IsBitcoinChain(chain) {
		return k.createNewBtcUpcBatchToSign(ctx)
	}
	return k.createNewBatchToSign(ctx)
}

func (k ChainKeeper) CreateNewBtcPoolingBatchToSign(ctx sdk.Context, chain nexus.ChainName, pk []byte) (types.CommandBatch, error) {
	if !types.IsBitcoinChain(chain) {
		return types.CommandBatch{}, fmt.Errorf("pooling is only supported for bitcoin chains")
	}

	firstCmd, ok := k.getFirstBtcPoolingCommand(ctx, pk)
	if !ok {
		return types.CommandBatch{}, nil
	}

	return k.createNewBtcBatchFollowCmd(ctx, firstCmd)
}

func (k ChainKeeper) HasBtcPoolingCommands(ctx sdk.Context, pk []byte) bool {
	if !types.IsBitcoinChain(k.GetName()) {
		clog.Redf("chain %s is not a bitcoin chain", k.GetName())
		return false
	}
	clog.Greenf("[ChainKeeper] [HasBtcPoolingCommands] chain: %+v", k.GetName())
	keys := k.getCommandQueue(ctx).Keys()
	clog.Redf("[ChainKeeper] [HasBtcPoolingCommands] keys: %+v", keys)
	for _, queueKey := range keys {
		var cmd types.Command
		ok := k.getStore(ctx).GetNew(key.FromBz(queueKey.AsKey()), &cmd)
		if ok {
			clog.Redf("[ChainKeeper] [HasBtcPoolingCommands] cmd: %+v", cmd)
		}
	}
	key := protocol.GetBTCKeyID(protocol.LIQUIDITY_MODEL_POOL, pk)
	// Loop through the command queue and check if there are any commands with the given key
	firstCmdFilter := func(value codec.ProtoMarshaler) bool {
		cmd, ok := value.(*types.Command)
		clog.Magentaf("[ChainKeeper] filter firstcmd, ok: %v, cmd.KeyID: %v, key: %v", ok, cmd.KeyID.String(), key.String())
		clog.Redf("[ChainKeeper] cmd: %v, %v", cmd, ok)
		return ok && cmd.KeyID.String() == key.String()
	}
	var firstCmd types.Command

	ok := k.getCommandQueue(ctx).FindUntil(&firstCmd, firstCmdFilter)

	//Do not dequeue the first command like in the function getFirstBtcPoolingCommand
	//ok := k.getCommandQueue(ctx).DequeueUntil(&firstCmd, firstCmdFilter)
	clog.Redf("[ChainKeeper] [HasBtcPoolingCommands] chain: %+v, firstCmd: %v, %v", k.GetName(), firstCmd, ok)
	return ok
}

func (k ChainKeeper) getFirstBtcPoolingCommand(ctx sdk.Context, pk []byte) (*types.Command, bool) {
	key := protocol.GetBTCKeyID(protocol.LIQUIDITY_MODEL_POOL, pk)
	clog.Magentaf("[ChainKeeper] getFirstBtcPoolingCommand, Key: %s", key)
	firstCmdFilter := func(value codec.ProtoMarshaler) bool {
		cmd, ok := value.(*types.Command)
		clog.Magentaf("[ChainKeeper] filter firstcmd, ok: %v, cmd.KeyID: %v, key: %v", ok, cmd.KeyID.String(), key.String())
		clog.Redf("[ChainKeeper] cmd: %v, %v", cmd, ok)
		return ok && cmd.KeyID.String() == key.String()
	}

	var firstCmd types.Command
	ok := k.getCommandQueue(ctx).DequeueUntil(&firstCmd, firstCmdFilter)
	clog.Redf("firstCmd: %v, %v", firstCmd, ok)
	return &firstCmd, ok
}

func (k ChainKeeper) createNewBtcUpcBatchToSign(ctx sdk.Context) (types.CommandBatch, error) {
	prefix := protocol.GetBTCKeyIDPrefix(protocol.LIQUIDITY_MODEL_UPC)
	firstCmdFilter := func(value codec.ProtoMarshaler) bool {
		cmd, ok := value.(*types.Command)
		return ok && strings.HasPrefix(cmd.KeyID.String(), prefix)
	}

	var firstCmd types.Command

	ok := k.getCommandQueue(ctx).DequeueUntil(&firstCmd, firstCmdFilter)
	if !ok {
		return types.CommandBatch{}, nil
	}

	return k.createNewBtcBatchFollowCmd(ctx, &firstCmd)
}

func (k ChainKeeper) createNewBtcBatchFollowCmd(ctx sdk.Context, cmd *types.Command) (types.CommandBatch, error) {
	chainID := sdk.NewIntFromBigInt(k.getSigner(ctx).ChainID())
	keyID := cmd.KeyID

	filter := func(value codec.ProtoMarshaler) bool {
		c, ok := value.(*types.Command)
		return ok && c.KeyID == keyID
	}

	commands := []types.Command{cmd.Clone()}
	for {
		var cmd types.Command
		ok := k.getCommandQueue(ctx).DequeueIf(&cmd, filter)
		if !ok {
			break
		}

		clog.Magentaf("[keeper] [createNewBtcBatchFollowCmd] command: %+v", cmd)

		// Note: Becareful with cmd.Clone() if you want to add more fields to the command, please update this function
		//TODO: change to commands = append(commands, cmd)
		clonedCmd := cmd.Clone()
		commands = append(commands, clonedCmd)
	}

	commandBatch, err := types.NewCommandBatchMetadata(ctx.BlockHeight(), chainID, keyID, commands)
	if err != nil {
		return types.CommandBatch{}, err
	}

	latest := k.GetLatestCommandBatch(ctx)
	if !latest.Is(types.BatchSigned) && !latest.Is(types.BatchNonExistent) {
		return types.CommandBatch{}, fmt.Errorf("latest command batch %s is still being processed", hex.EncodeToString(latest.GetID()))
	}

	commandBatch.PrevBatchedCommandsID = latest.GetID()
	k.setCommandBatchMetadata(ctx, commandBatch)
	k.setUnsignedCommandBatchID(ctx, commandBatch.ID)

	setter := func(m types.CommandBatchMetadata) {
		k.setCommandBatchMetadata(ctx, m)
	}
	clog.Greenf("[ChainKeeper] [createNewBtcBatchFollowCmd] commandBatch: %+v", commandBatch)
	return types.NewCommandBatch(commandBatch, setter), nil
}

func (k ChainKeeper) createNewBatchToSign(ctx sdk.Context) (types.CommandBatch, error) {
	var firstCmd types.Command
	ok := k.getCommandQueue(ctx).Dequeue(&firstCmd)
	if !ok {
		return types.CommandBatch{}, nil
	}

	chainID := sdk.NewIntFromBigInt(k.getSigner(ctx).ChainID())
	gasLimit := k.getCommandsGasLimit(ctx)
	gasCost := firstCmd.MaxGasCost
	keyID := firstCmd.KeyID

	filter := func(value codec.ProtoMarshaler) bool {
		cmd, ok := value.(*types.Command)
		gasCost += cmd.MaxGasCost
		// Note: This is used to limit the number of commands in the batch
		return ok && cmd.KeyID == keyID && gasCost <= gasLimit
	}

	commands := []types.Command{firstCmd.Clone()}
	for {
		var cmd types.Command
		ok := k.getCommandQueue(ctx).DequeueIf(&cmd, filter)
		if !ok {
			break
		}

		clog.Magentaf("[keeper] [createNewBatchToSign] command: %+v", cmd)

		// Note: Becareful with cmd.Clone() if you want to add more fields to the command, please update this function
		clonedCmd := cmd.Clone()
		commands = append(commands, clonedCmd)
	}

	commandBatch, err := types.NewCommandBatchMetadata(ctx.BlockHeight(), chainID, keyID, commands)
	if err != nil {
		return types.CommandBatch{}, err
	}

	latest := k.GetLatestCommandBatch(ctx)
	if !latest.Is(types.BatchSigned) && !latest.Is(types.BatchNonExistent) {
		return types.CommandBatch{}, fmt.Errorf("latest command batch %s is still being processed", hex.EncodeToString(latest.GetID()))
	}

	commandBatch.PrevBatchedCommandsID = latest.GetID()
	k.setCommandBatchMetadata(ctx, commandBatch)
	k.setUnsignedCommandBatchID(ctx, commandBatch.ID)

	setter := func(m types.CommandBatchMetadata) {
		k.setCommandBatchMetadata(ctx, m)
	}
	return types.NewCommandBatch(commandBatch, setter), nil
}

func (k ChainKeeper) GetLatestCommandBatch(ctx sdk.Context) types.CommandBatch {
	if batch := k.getLatestCommandBatchMetadata(ctx); batch.Status != types.BatchNonExistent {
		setter := func(m types.CommandBatchMetadata) {
			k.setCommandBatchMetadata(ctx, m)
		}
		return types.NewCommandBatch(batch, setter)
	}

	return types.NonExistentCommand
}

func (k ChainKeeper) GetLatestBtcPoolingBatch(ctx sdk.Context) *types.CommandBatch {
	chain := k.GetName()
	if !types.IsBitcoinChain(chain) {
		return nil
	}
	prefix := protocol.GetBTCKeyIDPrefix(protocol.LIQUIDITY_MODEL_POOL)

	var md *types.CommandBatchMetadata
	iter := k.getStore(ctx).Iterator(utils.LowerCaseKey(commandBatchPrefix))
	defer utils.CloseLogError(iter, k.Logger(ctx))
	for ; iter.Valid(); iter.Next() {
		var batch types.CommandBatchMetadata
		iter.UnmarshalValue(&batch)
		if batch.Status == types.BatchNonExistent {
			continue
		}
		if strings.HasPrefix(batch.KeyID.String(), prefix) {
			md = &batch
		}
	}

	if md == nil {
		return nil
	}

	setter := func(m types.CommandBatchMetadata) {
		k.setCommandBatchMetadata(ctx, m)
	}

	batch := types.NewCommandBatch(*md, setter)
	return &batch
}

func (k ChainKeeper) getLatestCommandBatchMetadata(ctx sdk.Context) types.CommandBatchMetadata {
	if batch := k.getUnsignedCommandBatch(ctx); batch.Status != types.BatchNonExistent {
		return batch
	}

	if id := k.getLatestSignedCommandBatchID(ctx); id != nil {
		return k.getCommandBatchMetadata(ctx, id)
	}
	return types.CommandBatchMetadata{Status: types.BatchNonExistent}
}

func (k ChainKeeper) getLatestSignedCommandBatchID(ctx sdk.Context) []byte {
	return k.getStore(ctx).GetRawNew(latestSignedBatchIDKey)
}

func (k ChainKeeper) getCommandBatchMetadata(ctx sdk.Context, id []byte) types.CommandBatchMetadata {
	var batch types.CommandBatchMetadata
	k.getStore(ctx).GetNew(key.FromStr(commandBatchPrefix).Append(key.FromBz(id)), &batch)
	return batch
}

func (k ChainKeeper) getUnsignedCommandBatch(ctx sdk.Context) types.CommandBatchMetadata {
	if id := k.getStore(ctx).GetRawNew(unsignedBatchIDKey); id != nil {
		return k.getCommandBatchMetadata(ctx, id)
	}

	return types.CommandBatchMetadata{}
}

func (k ChainKeeper) DeleteDeposit(ctx sdk.Context, deposit types.ERC20Deposit) {
	k.getStore(ctx).DeleteNew(
		confirmedDepositPrefix.Append(key.FromStr(deposit.TxID.Hex())).Append(key.FromUInt(deposit.LogIndex)))
	k.getStore(ctx).DeleteNew(
		burnedDepositPrefix.Append(key.FromStr(deposit.TxID.Hex())).Append(key.FromUInt(deposit.LogIndex)))
}

// DeleteUnsignedCommandBatchID deletes the unsigned command batch ID
func (k ChainKeeper) DeleteUnsignedCommandBatchID(ctx sdk.Context) {
	k.getStore(ctx).DeleteNew(unsignedBatchIDKey)
}

// SetGateway sets the gateway
func (k ChainKeeper) SetGateway(ctx sdk.Context, address types.Address) {
	k.setGateway(ctx, types.Gateway{Address: address})
}

func (k ChainKeeper) setGateway(ctx sdk.Context, gateway types.Gateway) {
	// TODO: remove this guard clause once genesis state can have nil Gateway
	if gateway.Address.IsZeroAddress() {
		return
	}

	funcs.MustNoErr(
		k.getStore(ctx).SetNewValidated(gatewayKey, &gateway))
}

// EnqueueConfirmedEvent enqueues the confirmed event
func (k ChainKeeper) EnqueueConfirmedEvent(ctx sdk.Context, id types.EventID) error {
	event, ok := k.GetEvent(ctx, id)
	if !ok {
		return fmt.Errorf("event %s does not exist", id)
	}
	if event.Status != types.EventConfirmed {
		return fmt.Errorf("event %s is not confirmed", id)
	}

	switch event.GetEvent().(type) {
	// the missing Event_ContractCall is no longer allowed to be enqueued in the
	// EVM module, it must be routed through the nexus module instead
	case *types.Event_ContractCallWithToken,
		*types.Event_TokenSent,
		*types.Event_Transfer,
		*types.Event_TokenDeployed,
		*types.Event_MultisigOperatorshipTransferred,
		*types.Event_SourceTxConfirmationEvent,
		*types.Event_RedeemToken:
		k.GetConfirmedEventQueue(ctx).Enqueue(getEventKey(id), &event)
	default:
		return fmt.Errorf("unsupported event type %T", event)
	}

	return nil
}

func (k ChainKeeper) GenerateSalt(ctx sdk.Context, recipient string) exported.Hash {
	nonce := utils.GetNonce(ctx.HeaderHash(), ctx.BlockGasMeter())
	bz := []byte(recipient)
	bz = append(bz, nonce[:]...)
	salt := exported.Hash(common.BytesToHash(crypto.Keccak256Hash(bz).Bytes()))
	return salt
}

func (k ChainKeeper) GetBatchByID(ctx sdk.Context, id []byte) types.CommandBatch {
	batch := k.getCommandBatchMetadata(ctx, id)

	setter := func(m types.CommandBatchMetadata) {
		k.setCommandBatchMetadata(ctx, m)
	}

	return types.NewCommandBatch(batch, setter)
}

func (k ChainKeeper) GetBurnerAddress(ctx sdk.Context, token types.ERC20Token, salt exported.Hash, gatewayAddr types.Address) (types.Address, error) {
	var tokenBurnerCodeHash exported.Hash
	if token.IsExternal() {
		// always use the latest burner byte code for external token
		burnerCode := k.GetBurnerByteCode(ctx)
		tokenBurnerCodeHash = exported.Hash(crypto.Keccak256Hash(burnerCode))
	} else {
		tokenBurnerCodeHash = funcs.MustOk(token.GetBurnerCodeHash())
	}

	var initCodeHash exported.Hash
	switch tokenBurnerCodeHash.Hex() {
	case types.BurnerCodeHashV1:
		addressType, err := abi.NewType("address", "address", nil)
		if err != nil {
			return types.Address{}, err
		}

		bytes32Type, err := abi.NewType("bytes32", "bytes32", nil)
		if err != nil {
			return types.Address{}, err
		}

		arguments := abi.Arguments{{Type: addressType}, {Type: bytes32Type}}
		params, err := arguments.Pack(token.GetAddress(), salt)
		if err != nil {
			return types.Address{}, err
		}

		initCodeHash = exported.Hash(crypto.Keccak256Hash(append(token.GetBurnerCode(), params...)))
	case types.BurnerCodeHashV2, types.BurnerCodeHashV3, types.BurnerCodeHashV4, types.BurnerCodeHashV5, types.BurnerCodeHashV6:
		initCodeHash = tokenBurnerCodeHash
	default:
		return types.Address{}, fmt.Errorf("unsupported burner code with hash %s for chain %s", tokenBurnerCodeHash.Hex(), k.chain)
	}

	return types.Address(crypto.CreateAddress2(common.Address(gatewayAddr), salt, initCodeHash.Bytes())), nil
}

func (k ChainKeeper) GetBurnerInfo(ctx sdk.Context, burnerAddr types.Address) *types.BurnerInfo {
	var result types.BurnerInfo
	if !k.getStore(ctx).GetNew(burnerAddrPrefix.Append(key.FromStr(burnerAddr.Hex())), &result) {
		return nil
	}

	return &result
}

func (k ChainKeeper) getBurnerInfos(ctx sdk.Context) []types.BurnerInfo {
	iter := k.getStore(ctx).IteratorNew(burnerAddrPrefix)
	defer utils.CloseLogError(iter, k.Logger(ctx))

	var burners []types.BurnerInfo
	for ; iter.Valid(); iter.Next() {
		var burner types.BurnerInfo
		iter.UnmarshalValue(&burner)
		burners = append(burners, burner)
	}

	return burners
}

func (k ChainKeeper) GetCommand(ctx sdk.Context, id types.CommandID) (types.Command, bool) {
	var cmd types.Command
	found := k.getStore(ctx).GetNew(key.FromStr(commandPrefix).Append(key.FromStr(id.Hex())), &cmd)

	return cmd, found
}
