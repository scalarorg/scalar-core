package scalarnet_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	ibcclient "github.com/cosmos/ibc-go/v4/modules/core/exported"
	"github.com/stretchr/testify/assert"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	appParams "github.com/scalarorg/scalar-core/app/params"
	"github.com/scalarorg/scalar-core/testutils/fake"
	"github.com/scalarorg/scalar-core/testutils/rand"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/math"
	utilsMock "github.com/scalarorg/scalar-core/utils/mock"
	"github.com/scalarorg/scalar-core/utils/slices"
	. "github.com/scalarorg/scalar-core/utils/test"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/scalarorg/scalar-core/x/scalarnet"
	"github.com/scalarorg/scalar-core/x/scalarnet/keeper"
	"github.com/scalarorg/scalar-core/x/scalarnet/types"
	"github.com/scalarorg/scalar-core/x/scalarnet/types/mock"
	scalartestutils "github.com/scalarorg/scalar-core/x/scalarnet/types/testutils"
)

func setupABCI() (sdk.Context, keeper.IBCKeeper, *mock.ChannelKeeperMock, *mock.IBCTransferKeeperMock) {
	ctx := sdk.NewContext(fake.NewMultiStore(), tmproto.Header{}, false, log.TestingLogger())
	encCfg := appParams.MakeEncodingConfig()
	scalarnetSubspace := params.NewSubspace(encCfg.Codec, encCfg.Amino, sdk.NewKVStoreKey("scalarnetKey"), sdk.NewKVStoreKey("tscalarnetKey"), "scalarnet")

	channelK := &mock.ChannelKeeperMock{
		GetNextSequenceSendFunc: func(ctx sdk.Context, portID string, channelID string) (uint64, bool) {
			return uint64(rand.I64Between(1, 999999)), true
		},
	}
	transferK := &mock.IBCTransferKeeperMock{}

	k := keeper.NewKeeper(encCfg.Codec, sdk.NewKVStoreKey("scalarnet"), scalarnetSubspace, channelK, &mock.FeegrantKeeperMock{})
	k.InitGenesis(ctx, types.DefaultGenesisState())
	ibcK := keeper.NewIBCKeeper(k, transferK)
	return ctx, ibcK, channelK, transferK
}

func TestEndBlocker(t *testing.T) {
	var (
		bk                   *mock.BaseKeeperMock
		transferK            *mock.IBCTransferKeeperMock
		channelK             *mock.ChannelKeeperMock
		ibcK                 keeper.IBCKeeper
		transferQueue        *utilsMock.KVQueueMock
		queueSize            int
		queueIdx             int
		ibcTransferErrors    int
		transferLimit        int
		panicOnTransferError bool
		ctx                  sdk.Context
	)

	repeats := 20

	givenTransferQueue := Given("transfer queue", func() {
		ctx, ibcK, channelK, transferK = setupABCI()

		bk = &mock.BaseKeeperMock{
			LoggerFunc:                func(ctx sdk.Context) log.Logger { return log.NewNopLogger() },
			GetIBCTransferQueueFunc:   func(ctx sdk.Context) utils.KVQueue { return transferQueue },
			GetRouteTimeoutWindowFunc: func(ctx sdk.Context) uint64 { return 10 },
			GetEndBlockerLimitFunc:    func(ctx sdk.Context) uint64 { return 1000 },
			SetTransferFailedFunc: func(sdk.Context, nexus.TransferID) error {
				return nil
			},
		}

		transferK.SendTransferFunc = func(sdk.Context, string, string, sdk.Coin, sdk.AccAddress, string, clienttypes.Height, uint64) error {
			if queueIdx <= ibcTransferErrors {
				if panicOnTransferError {
					panic("panicked on transfer")
				}

				return fmt.Errorf("failed to send transfer")
			}

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					ibcchanneltypes.EventTypeSendPacket,
				))
			return nil
		}

		channelK.GetChannelClientStateFunc = func(ctx sdk.Context, portID string, channelID string) (string, ibcclient.ClientState, error) {
			return "07-tendermint-0", scalartestutils.ClientState(), nil
		}

		transferQueue = &utilsMock.KVQueueMock{
			IsEmptyFunc: func() bool {
				return queueIdx == queueSize
			},
			DequeueFunc: func(value codec.ProtoMarshaler) bool {
				if queueIdx == queueSize {
					return false
				}

				transfer := scalartestutils.RandomIBCTransfer()
				bz, _ := transfer.Marshal()
				if err := value.Unmarshal(bz); err != nil {
					panic(err)
				}
				queueIdx++

				return true
			},
		}

		queueIdx = 0
		ibcTransferErrors = 0
		panicOnTransferError = false
	})

	givenTransferQueue.
		When("queue is empty", func() {
			queueSize = 0
		}).
		Then("should do nothing", func(t *testing.T) {
			_, err := scalarnet.EndBlocker(ctx, abci.RequestEndBlock{Height: ctx.BlockHeight()}, bk, ibcK)
			assert.NoError(t, err)
			assert.Equal(t, len(transferQueue.DequeueCalls()), 0)
			assert.Equal(t, len(transferK.SendTransferCalls()), 0)
		}).
		Run(t, repeats)

	givenTransferQueue.
		When("given a queue size", func() {
			queueSize = int(rand.I64Between(50, 200))
		}).
		Then("should init ibc transfers", func(t *testing.T) {
			_, err := scalarnet.EndBlocker(ctx, abci.RequestEndBlock{Height: ctx.BlockHeight()}, bk, ibcK)
			assert.NoError(t, err)
			assert.Equal(t, queueSize, len(transferQueue.DequeueCalls()))
			assert.Equal(t, queueSize, len(transferK.SendTransferCalls()))
			assert.Equal(t, queueSize, slices.Reduce(ctx.EventManager().Events().ToABCIEvents(), 0, func(c int, e abci.Event) int {
				if e.Type == ibcchanneltypes.EventTypeSendPacket {
					c++
				}
				return c
			}))
		}).
		Run(t, repeats)

	givenTransferQueue.
		When("given a queue size", func() {
			queueSize = int(rand.I64Between(50, 200))
		}).
		When("there is a transfer limit", func() {
			transferLimit = int(rand.I64Between(0, 200))
			bk.GetEndBlockerLimitFunc = func(ctx sdk.Context) uint64 { return uint64(transferLimit) }
		}).
		Then("should init ibc transfers", func(t *testing.T) {
			numTransfers := math.Min(queueSize, transferLimit)
			_, err := scalarnet.EndBlocker(ctx, abci.RequestEndBlock{Height: ctx.BlockHeight()}, bk, ibcK)
			assert.NoError(t, err)
			assert.Equal(t, numTransfers, len(transferQueue.DequeueCalls()))
			assert.Equal(t, numTransfers, len(transferK.SendTransferCalls()))
			assert.Equal(t, numTransfers, slices.Reduce(ctx.EventManager().Events().ToABCIEvents(), 0, func(c int, e abci.Event) int {
				if e.Type == ibcchanneltypes.EventTypeSendPacket {
					c++
				}
				return c
			}))
		}).
		Run(t, repeats)

	givenTransferQueue.
		When("given a queue size and some ibc transfer errors", func() {
			queueSize = int(rand.I64Between(50, 200))
			ibcTransferErrors = int(rand.I64Between(1, int64(queueSize)) + 1)
		}).
		Then("should set failed transfers", func(t *testing.T) {
			_, err := scalarnet.EndBlocker(ctx, abci.RequestEndBlock{Height: ctx.BlockHeight()}, bk, ibcK)
			assert.NoError(t, err)
			assert.Equal(t, queueSize, len(transferQueue.DequeueCalls()))
			assert.Equal(t, queueSize, len(transferK.SendTransferCalls()))
			assert.Equal(t, queueSize-ibcTransferErrors, slices.Reduce(ctx.EventManager().Events().ToABCIEvents(), 0, func(c int, e abci.Event) int {
				if e.Type == ibcchanneltypes.EventTypeSendPacket {
					c++
				}
				return c
			}))
			assert.Equal(t, ibcTransferErrors, len(bk.SetTransferFailedCalls()))
		}).
		Run(t, repeats)

	givenTransferQueue.
		When("given a queue size and some ibc transfer panics", func() {
			queueSize = int(rand.I64Between(50, 200))
			ibcTransferErrors = int(rand.I64Between(1, int64(queueSize)) + 1)
			panicOnTransferError = true
		}).
		Then("should set transfers failed", func(t *testing.T) {
			_, err := scalarnet.EndBlocker(ctx, abci.RequestEndBlock{Height: ctx.BlockHeight()}, bk, ibcK)
			assert.NoError(t, err)
			assert.Equal(t, queueSize, len(transferQueue.DequeueCalls()))
			assert.Equal(t, queueSize, len(transferK.SendTransferCalls()))
			assert.Equal(t, queueSize-ibcTransferErrors, slices.Reduce(ctx.EventManager().Events().ToABCIEvents(), 0, func(c int, e abci.Event) int {
				if e.Type == ibcchanneltypes.EventTypeSendPacket {
					c++
				}
				return c
			}))
			assert.Equal(t, ibcTransferErrors, len(bk.SetTransferFailedCalls()))
		}).
		Run(t, repeats)
}
