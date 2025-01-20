package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/scalarorg/scalar-core/utils"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	multisigTypes "github.com/scalarorg/scalar-core/x/multisig/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	reward "github.com/scalarorg/scalar-core/x/reward/exported"
	snapshot "github.com/scalarorg/scalar-core/x/snapshot/exported"
)

// Keeper provides keeper functionality of this module
//
//go:generate moq -pkg mock -out ./mock/expected_keepers.go . Keeper Snapshotter Staker Slasher Rewarder Nexus
type Keeper interface {
	Logger(ctx sdk.Context) log.Logger
	GetParams(ctx sdk.Context) (params Params)

	CreateCustodian(ctx sdk.Context, params Params) (err error)
	GetCustodians(ctx sdk.Context) (custodians []*Custodian, ok bool)
	CreateCustodianGroup(ctx sdk.Context, params Params) (err error)
	GetAllCustodianGroups(ctx sdk.Context) (custodianGroups []*CustodianGroup, ok bool)
	GetCustodianGroup(ctx sdk.Context, groupId string) (custodianGroup *CustodianGroup, ok bool)

	//GetCurrentKeyID(ctx sdk.Context, chainName nexus.ChainName) (multisig.KeyID, bool)
	GetKey(ctx sdk.Context, keyID multisig.KeyID) (multisig.Key, bool)
	SetKey(ctx sdk.Context, key multisigTypes.Key)
	SetCovenantRouter(router CovenantRouter)

	GetSigningSessions(ctx sdk.Context) (signingSessions []SigningSession, ok bool)

	GetSigningSessionsByExpiry(ctx sdk.Context, expiry int64) []SigningSession
	DeleteSigningSession(ctx sdk.Context, id uint64)
	GetCovenantRouter() CovenantRouter

	SignPsbt(ctx sdk.Context, keyID multisig.KeyID, psbt Psbt, module string, chainName nexus.ChainName, moduleMetadata ...codec.ProtoMarshaler) error
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
}
