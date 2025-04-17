package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/key"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	covenant "github.com/scalarorg/scalar-core/x/covenant/exported"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	mtypes "github.com/scalarorg/scalar-core/x/multisig/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	protocol "github.com/scalarorg/scalar-core/x/protocol/exported"
	reward "github.com/scalarorg/scalar-core/x/reward/exported"
	scalarnetTypes "github.com/scalarorg/scalar-core/x/scalarnet/types"
	snapshot "github.com/scalarorg/scalar-core/x/snapshot/exported"
	vote "github.com/scalarorg/scalar-core/x/vote/exported"
)

// Keeper provides keeper functionality of this module
//
//go:generate moq -pkg mock -out ./mock/expected_keepers.go . Keeper Snapshotter StakingKeeper SlashingKeeper Rewarder Nexus MultisigKeeper ProtocolKeeper BaseKeeper Voter ScalarnetKeeper
type Keeper interface {
	Logger(ctx sdk.Context) log.Logger
	GetParams(ctx sdk.Context) (params Params)

	CreateCustodian(ctx sdk.Context, params Params) (err error)
	GetCustodians(ctx sdk.Context) (custodians []*covenant.Custodian, ok bool)
	CreateCustodianGroup(ctx sdk.Context, params Params) (err error)
	GetAllCustodianGroups(ctx sdk.Context) (custodianGroups []*covenant.CustodianGroup, ok bool)
	GetCustodianGroup(ctx sdk.Context, groupId chains.Hash) (custodianGroup *covenant.CustodianGroup, ok bool)

	//GetCurrentKeyID(ctx sdk.Context, chainName nexus.ChainName) (multisig.KeyID, bool)
	GetKey(ctx sdk.Context, keyID multisig.KeyID) (mtypes.Key, bool)
	SetKey(ctx sdk.Context, key mtypes.Key)
	SetCovenantRouter(router CovenantRouter)

	GetSigningSessions(ctx sdk.Context) (signingSessions []SigningSession, ok bool)

	GetSigningSessionsByExpiry(ctx sdk.Context, expiry int64) []SigningSession
	DeleteSigningSession(ctx sdk.Context, id uint64)
	GetCovenantRouter() CovenantRouter

	SignPsbt(ctx sdk.Context, keyID multisig.KeyID, multiPsbt []covenant.Psbt, module string, chainName nexus.ChainName, moduleMetadata ...codec.ProtoMarshaler) error

	GetEventsQueue(ctx sdk.Context) utils.BlockHeightKVQueue
	EnqueueEvent(ctx sdk.Context, event *Event) error

	GetRedeemSession(ctx sdk.Context, custodianGroupUID chains.Hash) (*RedeemSession, bool)
	SetRedeemSession(ctx sdk.Context, redeemSession *RedeemSession)
	// GetRedeemSessionByExpiry(ctx sdk.Context, expiry int64) []RedeemSession
	SetSwitchingForRedeemSession(ctx sdk.Context, custodianGroupUID chains.Hash) error
	UpdatePreparingToExecuting(ctx sdk.Context, custodianGroupUID chains.Hash) error
	UpdateExecutingToPreparing(ctx sdk.Context, custodianGroupUID chains.Hash, sequence uint64) error
	SetUtxoSnapshot(ctx sdk.Context, utxoSnapshot *UTXOSnapshot) error

	GetReserveUTXOCommandByID(ctx sdk.Context, id []byte) StandaloneCommand

	RotateKey(ctx sdk.Context, chainName nexus.ChainName, key mtypes.Key) error
	GetSigningSession(ctx sdk.Context, id uint64) (signing SigningSession, ok bool)
	SetSigningSession(ctx sdk.Context, signing SigningSession)
	CreateRedeemParams(ctx sdk.Context, req *ReserveRedeemUtxoRequest, custodianGrUID chains.Hash, sequence uint64) ([]byte, *CommandID, error)

	SetStandaloneCommandMetadata(ctx sdk.Context, meta StandaloneCommandMetadata, prefix key.Key)
	SetUnsignedStandaloneCommandID(ctx sdk.Context, id []byte)
	RenewRedeemSession(ctx sdk.Context, custodianGroupUID chains.Hash) error
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
	GetChainMaintainerState(ctx sdk.Context, chain nexus.Chain, address sdk.ValAddress) (nexus.MaintainerState, bool)
	GetChainMaintainersByChainName(ctx sdk.Context, chainName nexus.ChainName) []sdk.ValAddress
	SetChainMaintainerState(ctx sdk.Context, maintainerState nexus.MaintainerState) error
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
	FindProtocolInfoByCustodianGroupUID(ctx sdk.Context, custodianGroupUIDs [][]byte) []*protocol.ProtocolInfo
	FindProtocolInfoByExternalSymbol(ctx sdk.Context, symbol string) (*protocol.ProtocolInfo, error)
}

type BaseKeeper interface {
	ForChain(ctx sdk.Context, chain nexus.ChainName) (chainsTypes.ChainKeeper, error)
	GetLatestCommandBatchForChain(ctx sdk.Context, chainName nexus.ChainName) chainsTypes.CommandBatch
	CreateNewBtcPoolingBatchToSign(ctx sdk.Context, chainName nexus.ChainName, pk []byte) (chainsTypes.CommandBatch, error)
	GetLatestBtcPoolingBatchForChain(ctx sdk.Context, chain nexus.ChainName) *chainsTypes.CommandBatch
}

type Voter interface {
	InitializePoll(ctx sdk.Context, pollBuilder vote.PollBuilder) (vote.PollID, error)
}

type ScalarnetKeeper interface {
	GetParams(ctx sdk.Context) (params scalarnetTypes.Params)
}
