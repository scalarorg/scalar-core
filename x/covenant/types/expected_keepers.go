package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/scalarorg/scalar-core/utils"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	covenant "github.com/scalarorg/scalar-core/x/covenant/exported"
	exported "github.com/scalarorg/scalar-core/x/covenant/exported"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	mtypes "github.com/scalarorg/scalar-core/x/multisig/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	protocol "github.com/scalarorg/scalar-core/x/protocol/exported"
	reward "github.com/scalarorg/scalar-core/x/reward/exported"
	snapshot "github.com/scalarorg/scalar-core/x/snapshot/exported"
	vote "github.com/scalarorg/scalar-core/x/vote/exported"
)

// Keeper provides keeper functionality of this module
//
//go:generate moq -pkg mock -out ./mock/expected_keepers.go . Keeper Snapshotter Staker Slasher Rewarder Nexus
type Keeper interface {
	Logger(ctx sdk.Context) log.Logger
	GetParams(ctx sdk.Context) (params Params)

	CreateCustodian(ctx sdk.Context, params Params) (err error)
	GetCustodians(ctx sdk.Context) (custodians []*covenant.Custodian, ok bool)
	CreateCustodianGroup(ctx sdk.Context, params Params) (err error)
	GetAllCustodianGroups(ctx sdk.Context) (custodianGroups []*covenant.CustodianGroup, ok bool)
	GetCustodianGroup(ctx sdk.Context, groupId string) (custodianGroup *covenant.CustodianGroup, ok bool)

	//GetCurrentKeyID(ctx sdk.Context, chainName nexus.ChainName) (multisig.KeyID, bool)
	GetKey(ctx sdk.Context, keyID multisig.KeyID) (mtypes.Key, bool)
	SetKey(ctx sdk.Context, key mtypes.Key)
	SetCovenantRouter(router CovenantRouter)

	GetSigningSessions(ctx sdk.Context) (signingSessions []SigningSession, ok bool)

	GetSigningSessionsByExpiry(ctx sdk.Context, expiry int64) []SigningSession
	DeleteSigningSession(ctx sdk.Context, id uint64)
	GetCovenantRouter() CovenantRouter

	SignPsbt(ctx sdk.Context, keyID multisig.KeyID, multiPsbt []exported.Psbt, module string, chainName nexus.ChainName, moduleMetadata ...codec.ProtoMarshaler) error
}

// Snapshotter provides snapshot keeper functionality
type Snapshotter interface {
	CreateSnapshot(
		ctx sdk.Context,
		candidates []sdk.ValAddress,
		filterFunc func(snapshot.ValidatorI) bool,
		weightFunc func(consensusPower sdk.Uint) sdk.Uint,
		threshold utils.Threshold,
	) (snapshot.Snapshot, error)
	GetProxy(ctx sdk.Context, operator sdk.ValAddress) (addr sdk.AccAddress, active bool)
	GetOperator(ctx sdk.Context, proxy sdk.AccAddress) sdk.ValAddress
}

// Staker provides staking keeper functionality
type StakingKeeper interface {
	GetBondedValidatorsByPower(ctx sdk.Context) []stakingTypes.Validator
}

// Slasher provides slashing keeper functionality
type SlashingKeeper interface {
	IsTombstoned(ctx sdk.Context, consAddr sdk.ConsAddress) bool
}

// Rewarder provides reward keeper functionality
type Rewarder interface {
	GetPool(ctx sdk.Context, name string) reward.RewardPool
}

// Nexus provides nexus keeper functionality
type Nexus interface {
	GetChain(ctx sdk.Context, chain nexus.ChainName) (nexus.Chain, bool)
	GetChains(ctx sdk.Context) []nexus.Chain
	GetChainMaintainers(ctx sdk.Context, chain nexus.Chain) []sdk.ValAddress
	IsChainActivated(ctx sdk.Context, chain nexus.Chain) bool
}

// MultisigKeeper provides functionality to the multisig module
type MultisigKeeper interface {
	GetCurrentKeyID(ctx sdk.Context, chainName nexus.ChainName) (multisig.KeyID, bool)
	GetNextKeyID(ctx sdk.Context, chainName nexus.ChainName) (multisig.KeyID, bool)
	GetKey(ctx sdk.Context, keyID multisig.KeyID) (multisig.Key, bool)
	AssignKey(ctx sdk.Context, chainName nexus.ChainName, keyID multisig.KeyID) error
	RotateKey(ctx sdk.Context, chainName nexus.ChainName) error
	Sign(ctx sdk.Context, keyID multisig.KeyID, payloadHash multisig.Hash, module string, moduleMetadata ...codec.ProtoMarshaler) error
}

type ProtocolKeeper interface {
	FindProtocolInfoByExternalSymbol(ctx sdk.Context, symbol string) (*protocol.ProtocolInfo, error)
}

type BaseKeeper interface {
	ForChain(ctx sdk.Context, chain nexus.ChainName) (chainsTypes.ChainKeeper, error)
}

type Voter interface {
	InitializePoll(ctx sdk.Context, pollBuilder vote.PollBuilder) (vote.PollID, error)
}
