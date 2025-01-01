package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
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

	GetCurrentKeyID(ctx sdk.Context, chainName nexus.ChainName) (exported.KeyID, bool)
	GetKey(ctx sdk.Context, keyID exported.KeyID) (exported.Key, bool)
	SetKey(ctx sdk.Context, key Key)
	GetSigRouter() SigRouter

	GetSigningSessions(ctx sdk.Context) (signingSessions []SigningSession, ok bool)
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
type Staker interface {
	GetBondedValidatorsByPower(ctx sdk.Context) []stakingTypes.Validator
}

// Slasher provides slashing keeper functionality
type Slasher interface {
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
}
