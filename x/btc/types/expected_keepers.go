package types

import (
	utils "github.com/axelarnetwork/axelar-core/utils"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	snapshot "github.com/axelarnetwork/axelar-core/x/snapshot/exported"
	vote "github.com/axelarnetwork/axelar-core/x/vote/exported"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

//go:generate moq -out ./mock/expected_keepers.go -pkg mock . Voter Nexus Snapshotter BaseKeeper ChainKeeper Rewarder StakingKeeper SlashingKeeper MultisigKeeper

// BaseKeeper is implemented by this module's base keeper
type BaseKeeper interface {
	Logger(ctx sdk.Context) log.Logger

	CreateChain(ctx sdk.Context, params Params) error

	ForChain(ctx sdk.Context, chain nexus.ChainName) (ChainKeeper, error)
}

// ChainKeeper is implemented by this module's chain keeper
type ChainKeeper interface {
	Logger(ctx sdk.Context) log.Logger

	GetName() nexus.ChainName

	GetParams(ctx sdk.Context) Params

	// GetNetwork(ctx sdk.Context) string

	GetRequiredConfirmationHeight(ctx sdk.Context) uint64
}

// ParamsKeeper represents a global paramstore
type ParamsKeeper interface {
	Subspace(s string) params.Subspace
	GetSubspace(s string) (params.Subspace, bool)
}

// Voter exposes voting functionality
type Voter interface {
	InitializePoll(ctx sdk.Context, pollBuilder vote.PollBuilder) (vote.PollID, error)
}

// Nexus provides functionality to manage cross-chain transfers
type Nexus interface {
	// LinkAddresses(ctx sdk.Context, sender nexus.CrossChainAddress, recipient nexus.CrossChainAddress) error
	// GetRecipient(ctx sdk.Context, sender nexus.CrossChainAddress) (nexus.CrossChainAddress, bool)
	// EnqueueTransfer(ctx sdk.Context, senderChain nexus.Chain, recipient nexus.CrossChainAddress, asset sdk.Coin) (nexus.TransferID, error)
	// EnqueueForTransfer(ctx sdk.Context, sender nexus.CrossChainAddress, amount sdk.Coin) (nexus.TransferID, error)
	// GetTransfersForChainPaginated(ctx sdk.Context, chain nexus.Chain, state nexus.TransferState, pageRequest *query.PageRequest) ([]nexus.CrossChainTransfer, *query.PageResponse, error)
	// ArchivePendingTransfer(ctx sdk.Context, transfer nexus.CrossChainTransfer)
	// SetChain(ctx sdk.Context, chain nexus.Chain)
	// GetChains(ctx sdk.Context) []nexus.Chain
	GetChain(ctx sdk.Context, chain nexus.ChainName) (nexus.Chain, bool)
	// IsAssetRegistered(ctx sdk.Context, chain nexus.Chain, denom string) bool
	// RegisterAsset(ctx sdk.Context, chain nexus.Chain, asset nexus.Asset, limit sdk.Uint, window time.Duration) error
	GetChainMaintainers(ctx sdk.Context, chain nexus.Chain) []sdk.ValAddress
	IsChainActivated(ctx sdk.Context, chain nexus.Chain) bool
	// GetChainByNativeAsset(ctx sdk.Context, asset string) (chain nexus.Chain, ok bool)
	// ComputeTransferFee(ctx sdk.Context, sourceChain nexus.Chain, destinationChain nexus.Chain, asset sdk.Coin) (sdk.Coin, error)
	// AddTransferFee(ctx sdk.Context, coin sdk.Coin)
	// GetChainMaintainerState(ctx sdk.Context, chain nexus.Chain, address sdk.ValAddress) (nexus.MaintainerState, bool)
	// SetChainMaintainerState(ctx sdk.Context, maintainerState nexus.MaintainerState) error
	// RateLimitTransfer(ctx sdk.Context, chain nexus.ChainName, asset sdk.Coin, direction nexus.TransferDirection) error
	// SetNewMessage(ctx sdk.Context, m nexus.GeneralMessage) error
	// GetProcessingMessages(ctx sdk.Context, chain nexus.ChainName, limit int64) []nexus.GeneralMessage
	// SetMessageFailed(ctx sdk.Context, id string) error
	// SetMessageExecuted(ctx sdk.Context, id string) error
	// EnqueueRouteMessage(ctx sdk.Context, id string) error
}

// InitPoller is a minimal interface to start a poll. This must be a type alias instead of a type definition,
// because the concrete implementation of Signer (specifically StartSign) is defined in a different package using another (identical)
// InitPoller interface. Go cannot match the types otherwise
type InitPoller = interface {
	// InitializePoll(ctx sdk.Context, pollBuilder vote.PollBuilder) (vote.PollID, error)
}

// Snapshotter provides access to the snapshot functionality
type Snapshotter interface {
	CreateSnapshot(ctx sdk.Context, candidates []sdk.ValAddress, filterFunc func(snapshot.ValidatorI) bool, weightFunc func(consensusPower sdk.Uint) sdk.Uint, threshold utils.Threshold) (snapshot.Snapshot, error)
	GetProxy(ctx sdk.Context, principal sdk.ValAddress) (addr sdk.AccAddress, active bool)
}

// Rewarder provides reward functionality
type Rewarder interface {
	// GetPool(ctx sdk.Context, name string) reward.RewardPool
}

// StakingKeeper adopts the methods from "github.com/cosmos/cosmos-sdk/x/staking/exported" that are
// actually used by this module
type StakingKeeper interface {
	// PowerReduction(ctx sdk.Context) sdk.Int
}

// SlashingKeeper provides functionality to manage slashing info for a validator
type SlashingKeeper interface {
	IsTombstoned(ctx sdk.Context, consAddr sdk.ConsAddress) bool
}

// MultisigKeeper provides functionality to the multisig module
type MultisigKeeper interface {
	// GetCurrentKeyID(ctx sdk.Context, chainName nexus.ChainName) (multisig.KeyID, bool)
	// GetNextKeyID(ctx sdk.Context, chainName nexus.ChainName) (multisig.KeyID, bool)
	// GetKey(ctx sdk.Context, keyID multisig.KeyID) (multisig.Key, bool)
	// AssignKey(ctx sdk.Context, chainName nexus.ChainName, keyID multisig.KeyID) error
	// RotateKey(ctx sdk.Context, chainName nexus.ChainName) error
	// Sign(ctx sdk.Context, keyID multisig.KeyID, payloadHash multisig.Hash, module string, moduleMetadata ...codec.ProtoMarshaler) error
}
