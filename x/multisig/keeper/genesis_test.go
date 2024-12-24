package keeper_test

import (
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/assert"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/scalarorg/scalar-core/app"
	"github.com/scalarorg/scalar-core/testutils/fake"
	"github.com/scalarorg/scalar-core/testutils/rand"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	test "github.com/scalarorg/scalar-core/utils/test"
	chainsKeeper "github.com/scalarorg/scalar-core/x/chains/keeper"
	"github.com/scalarorg/scalar-core/x/multisig"
	"github.com/scalarorg/scalar-core/x/multisig/exported"
	"github.com/scalarorg/scalar-core/x/multisig/exported/testutils"
	"github.com/scalarorg/scalar-core/x/multisig/keeper"
	keeperMock "github.com/scalarorg/scalar-core/x/multisig/keeper/mock"
	"github.com/scalarorg/scalar-core/x/multisig/types"
	"github.com/scalarorg/scalar-core/x/multisig/types/mock"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	reward "github.com/scalarorg/scalar-core/x/reward/exported"
	rewardmock "github.com/scalarorg/scalar-core/x/reward/exported/mock"
	snapshot "github.com/scalarorg/scalar-core/x/snapshot/exported"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
)

func TestInitExportGenesis(t *testing.T) {
	encCfg := app.MakeEncodingConfig()
	chain := chains.Ethereum
	validators := slices.Expand(func(int) snapshot.Participant { return snapshot.NewParticipant(rand.ValAddr(), sdk.OneUint()) }, 10)

	var (
		msgServer   types.MsgServiceServer
		k           keeper.Keeper
		ctx         sdk.Context
		snapshotter *keeperMock.SnapshotterMock
		nexusK      *mock.NexusMock
		rewardK     *mock.RewarderMock
		keyID       exported.KeyID
	)

	setup := func() {
		subspace := params.NewSubspace(encCfg.Codec, encCfg.Amino, sdk.NewKVStoreKey("paramsKey"), sdk.NewKVStoreKey("tparamsKey"), "multisig")
		k = keeper.NewKeeper(encCfg.Codec, sdk.NewKVStoreKey(types.StoreKey), subspace)

		multisigRounter := types.NewSigRouter()
		multisigRounter.AddHandler(chain.Module, chainsKeeper.NewSigHandler(encCfg.Codec, &chainsKeeper.BaseKeeper{}))
		k.SetSigRouter(multisigRounter)

		ctx = rand.Context(fake.NewMultiStore())
		k.InitGenesis(ctx, types.DefaultGenesisState())
		snapshotter = &keeperMock.SnapshotterMock{
			CreateSnapshotFunc: func(sdk.Context, utils.Threshold) (snapshot.Snapshot, error) {
				return snapshot.NewSnapshot(ctx.BlockTime(), ctx.BlockHeight(), validators, sdk.NewUint(10)), nil
			},
		}
		nexusK = &mock.NexusMock{
			GetChainFunc: func(ctx sdk.Context, chainName nexus.ChainName) (nexus.Chain, bool) {
				return chain, chain.GetName().Equals(chainName)
			},
		}
		pool := rewardmock.RewardPoolMock{
			ReleaseRewardsFunc: func(valAddress sdk.ValAddress) error { return nil },
		}
		rewardK = &mock.RewarderMock{
			GetPoolFunc: func(ctx sdk.Context, name string) reward.RewardPool { return &pool },
		}

		msgServer = keeper.NewMsgServer(k, snapshotter, &mock.StakerMock{}, nexusK)
	}

	givenMsgServer := test.Given("a multisig msg server", setup)

	whenKeygenSessionExists := test.When("some keygen session exists", func() {
		msgServer.StartKeygen(sdk.WrapSDKContext(ctx), types.NewStartKeygenRequest(rand.AccAddr(), testutils.KeyID()))
	})

	whenKeyExists := test.When("some key exists", func() {
		keyID = testutils.KeyID()

		msgServer.StartKeygen(sdk.WrapSDKContext(ctx), types.NewStartKeygenRequest(rand.AccAddr(), keyID))
		for _, v := range validators {
			snapshotter.GetOperatorFunc = func(sdk.Context, sdk.AccAddress) sdk.ValAddress { return v.Address }

			sk := funcs.Must(btcec.NewPrivateKey())
			msgServer.SubmitPubKey(sdk.WrapSDKContext(ctx), types.NewSubmitPubKeyRequest(rand.AccAddr(), keyID, sk.PubKey().SerializeCompressed(), ecdsa.Sign(sk, []byte(keyID)).Serialize()))
		}

		multisig.EndBlocker(ctx.WithBlockHeight(ctx.BlockHeight()+types.DefaultParams().KeygenGracePeriod), abci.RequestEndBlock{}, k, rewardK)
	})

	whenSigningSessionExists := test.When("some signing session exists", func() {
		k.Sign(ctx, keyID, rand.Bytes(exported.HashLength), chain.Module)
	})

	whenKeyIsAssigned := test.When("some key is assigned", func() {
		k.AssignKey(ctx, chain.Name, keyID)
	})

	whenKeyIsRotated := test.When("some key is rotated", func() {
		k.RotateKey(ctx, chain.Name)
	})

	t.Run("ExportGenesis", func(t *testing.T) {
		givenMsgServer.
			When2(whenKeygenSessionExists).
			Then("should export", func(t *testing.T) {
				actual := k.ExportGenesis(ctx)

				assert.Len(t, actual.KeygenSessions, 1)
				assert.Len(t, actual.Keys, 0)
				assert.NoError(t, actual.Validate())
			}).
			Run(t)

		givenMsgServer.
			When2(whenKeyExists).
			Then("should export", func(t *testing.T) {
				actual := k.ExportGenesis(ctx)

				assert.Len(t, actual.KeygenSessions, 0)
				assert.Len(t, actual.Keys, 1)
				assert.NoError(t, actual.Validate())
			}).
			Run(t)

		givenMsgServer.
			When2(whenKeyExists).
			When2(whenKeyIsAssigned).
			When2(whenKeyIsRotated).
			When2(whenSigningSessionExists).
			Then("should export", func(t *testing.T) {
				actual := k.ExportGenesis(ctx)

				assert.Len(t, actual.SigningSessions, 1)
				assert.NoError(t, actual.Validate())
			}).
			Run(t)

		givenMsgServer.
			When2(whenKeyExists).
			When2(whenKeyIsAssigned).
			When2(whenKeyIsRotated).
			When2(whenKeyExists).
			When2(whenKeyIsAssigned).
			Then("should export", func(t *testing.T) {
				actual := k.ExportGenesis(ctx)

				assert.Len(t, actual.KeyEpochs, 2)
				assert.NoError(t, actual.Validate())
			}).
			Run(t)
	})

	t.Run("InitGenesis", func(t *testing.T) {
		givenMsgServer.
			When2(whenKeygenSessionExists).
			When2(whenKeyExists).
			When2(whenKeyIsAssigned).
			When2(whenKeyIsRotated).
			When2(whenSigningSessionExists).
			When2(whenKeyExists).
			When2(whenKeyIsAssigned).
			Then("should init", func(t *testing.T) {
				expected := k.ExportGenesis(ctx)
				setup()

				k.InitGenesis(ctx, expected)
				actual := k.ExportGenesis(ctx)

				assert.NoError(t, actual.Validate())
				assert.Equal(t, expected, actual)
				assert.Error(t, k.Sign(ctx, keyID, rand.Bytes(exported.HashLength), chain.Module))
				assert.Error(t, k.AssignKey(ctx, chain.Name, keyID))
				assert.NoError(t, k.RotateKey(ctx, chain.Name))
				assert.NoError(t, k.Sign(ctx, keyID, rand.Bytes(exported.HashLength), chain.Module))
				assert.Len(t, k.ExportGenesis(ctx).SigningSessions, 2)
			}).
			Run(t)

		givenMsgServer.
			When2(whenKeygenSessionExists).
			When2(whenKeyExists).
			When2(whenKeyIsAssigned).
			When2(whenKeyIsRotated).
			When2(whenSigningSessionExists).
			When2(whenKeyExists).
			Then("should init", func(t *testing.T) {
				expected := k.ExportGenesis(ctx)
				setup()

				k.InitGenesis(ctx, expected)
				actual := k.ExportGenesis(ctx)

				assert.NoError(t, actual.Validate())
				assert.Equal(t, expected, actual)
				assert.Error(t, k.Sign(ctx, keyID, rand.Bytes(exported.HashLength), chain.Module))
				assert.NoError(t, k.AssignKey(ctx, chain.Name, keyID))
				assert.NoError(t, k.RotateKey(ctx, chain.Name))
			}).
			Run(t)

		givenMsgServer.
			When2(whenKeygenSessionExists).
			When2(whenKeyExists).
			When2(whenKeyIsAssigned).
			When2(whenKeyIsRotated).
			When2(whenSigningSessionExists).
			Then("should init", func(t *testing.T) {
				expected := k.ExportGenesis(ctx)
				setup()

				k.InitGenesis(ctx, expected)
				actual := k.ExportGenesis(ctx)

				assert.NoError(t, actual.Validate())
				assert.Equal(t, expected, actual)
				assert.NoError(t, k.Sign(ctx, keyID, rand.Bytes(exported.HashLength), chain.Module))
				assert.Error(t, k.AssignKey(ctx, chain.Name, keyID))
				assert.Error(t, k.RotateKey(ctx, chain.Name))
			}).
			Run(t)
	})
}
