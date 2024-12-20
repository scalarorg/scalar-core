package events_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	tm "github.com/tendermint/tendermint/types"

	testutils "github.com/scalarorg/scalar-core/utils/test"
	"github.com/scalarorg/scalar-core/utils/test/rand"

	"github.com/scalarorg/scalar-core/tm-events/events"
	"github.com/scalarorg/scalar-core/tm-events/events/mock"
)

func TestBlockNotifier_BlockHeights(t *testing.T) {
	repeats := 20

	t.Run("GIVEN only query is responsive THEN sync all blocks", testutils.Func(func(t *testing.T) {

		client := NewClientMock()
		client.SubscribeFunc = func(context.Context, string, string, ...int) (<-chan coretypes.ResultEvent, error) {
			return nil, nil
		}
		start := rand.I64Between(0, 1000000)
		notifier := events.NewBlockNotifier(client, events.Timeout(1*time.Second), events.Retries(1), events.KeepAlive(1*time.Millisecond)).StartingAt(start)

		newBlockCount := rand.I64Between(1, 20)

		receivedBlocks, errChan := notifier.BlockHeights(context.Background())

		client.NextBlock(start + newBlockCount)

		timeout, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		for i := int64(0); i < newBlockCount; i++ {
			select {
			case <-timeout.Done():
				assert.FailNow(t, "timed out", "delivered %d of %d blocks", i+1, newBlockCount+1)
			case err := <-errChan:
				assert.FailNow(t, "returned error", err.Error())
			case receivedBlock := <-receivedBlocks:
				assert.Equal(t, start+i, receivedBlock)
			}
		}
	}).Repeat(repeats))

	t.Run("GIVEN only events are responsive THEN sync all blocks", testutils.Func(func(t *testing.T) {
		client := NewClientMock()
		start := rand.I64Between(0, 1000000)
		notifier := events.NewBlockNotifier(client, events.Timeout(1*time.Millisecond), events.Retries(1), events.KeepAlive(1*time.Millisecond)).StartingAt(start)

		receivedBlocks, errChan := notifier.BlockHeights(context.Background())

		firstBatch := start + rand.I64Between(1, 10)
		secondBatch := firstBatch + rand.I64Between(1, 10)

		client.NextBlock(firstBatch)
		client.NextBlock(secondBatch)

		timeout, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		for i := int64(0); i < secondBatch-start; i++ {
			select {
			case <-timeout.Done():
				assert.FailNow(t, "timed out", "delivered %d of %d blocks", i+1, firstBatch+secondBatch+1)
			case err := <-errChan:
				assert.FailNow(t, "returned error", err.Error())
			case receivedBlock := <-receivedBlocks:
				assert.Equal(t, start+i, receivedBlock)
			}
		}
	}).Repeat(repeats))

	t.Run("GIVEN context is canceled THEN shutdown gracefully", testutils.Func(func(t *testing.T) {
		client := NewClientMock()
		start := rand.I64Between(0, 1000000)
		notifier := events.NewBlockNotifier(client,
			events.Timeout(1*time.Millisecond), events.Retries(1), events.KeepAlive(1*time.Millisecond)).StartingAt(start)

		ctx, cancelMainCtx := context.WithCancel(context.Background())

		receivedBlocks, errChan := notifier.BlockHeights(ctx)

		timeout, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		go func() {
			for i := int64(0); i < rand.I64Between(50, 100); i++ {
				select {
				case <-timeout.Done():
					return
				default:
					client.NextBlock(start + i)
				}
			}
		}()

		for i := int64(0); i < rand.I64Between(1, 50); i++ {
			select {
			case err := <-errChan:
				assert.FailNow(t, "returned error", err.Error())
			case receivedBlock, ok := <-receivedBlocks:
				if !ok {
					assert.FailNow(t, "premature channel close")
				}
				assert.Equal(t, start+i, receivedBlock)
			}
		}

		cancelMainCtx()

		for {
			select {
			case _, ok := <-receivedBlocks:
				if !ok {
					return
				}
			case <-timeout.Done():
				assert.FailNow(t, "channel should have been closed")
			}
		}
	}).Repeat(repeats))

	t.Run("GIVEN subscription fails THEN continue", testutils.Func(func(t *testing.T) {
		client := NewClientMock()
		client.SubscribeFunc = func(context.Context, string, string, ...int) (<-chan coretypes.ResultEvent, error) {
			return nil, fmt.Errorf("some error")
		}
		start := rand.I64Between(0, 1000000)
		notifier := events.NewBlockNotifier(client, events.KeepAlive(1*time.Millisecond)).StartingAt(start)

		blocks, errChan := notifier.BlockHeights(context.Background())

		blockCount := rand.I64Between(1, 10)
		nextBlock := start
		for i := int64(0); i < blockCount; i++ {
			nextBlock += rand.I64Between(0, 200)
			client.NextBlock(nextBlock)
		}

		timeout, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		for i := int64(0); i < blockCount; i++ {
			select {
			case <-errChan:
				assert.FailNow(t, "should not fail")
			case _, ok := <-blocks:
				if !ok {
					assert.FailNow(t, "should not close block channel")
				}
			case <-timeout.Done():
				assert.FailNow(t, "test timed out")
			}
		}
	}).Repeat(repeats))

	t.Run("GIVEN query fails THEN return error", testutils.Func(func(t *testing.T) {
		client := NewClientMock()
		client.LatestSyncInfoFunc = func(context.Context) (*coretypes.SyncInfo, error) {
			return nil, fmt.Errorf("some error")
		}
		start := rand.I64Between(0, 1000000)
		notifier := events.NewBlockNotifier(client, events.KeepAlive(1*time.Millisecond)).StartingAt(start)

		blocks, errChan := notifier.BlockHeights(context.Background())

		blockCount := rand.I64Between(1, 10)
		nextBlock := start
		for i := int64(0); i < blockCount; i++ {
			nextBlock += rand.I64Between(0, 200)
			client.NextBlock(nextBlock)
		}

		timeout, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		for {
			select {
			case err := <-errChan:
				assert.Error(t, err)
				return
			case <-blocks:
			case <-timeout.Done():
				assert.FailNow(t, "test timed out")
			}
		}
	}).Repeat(repeats))

	t.Run("GIVEN start < 0 THEN start at latest block", testutils.Func(func(t *testing.T) {
		client := NewClientMock()
		start := rand.PosI64()
		client.NextBlock(start)

		notifier := events.NewBlockNotifier(client, events.KeepAlive(1*time.Millisecond)).StartingAt(-rand.PosI64())

		blocks, errChan := notifier.BlockHeights(context.Background())

		timeout, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		select {
		case <-errChan:
			assert.FailNow(t, "should not fail")
			return
		case b, ok := <-blocks:
			assert.True(t, ok)
			assert.Equal(t, start, b)
		case <-timeout.Done():
			assert.FailNow(t, "test timed out")
		}
	}).Repeat(repeats))
}

func NewBlockHeaderEvent(blockHeight int64) coretypes.ResultEvent {
	return coretypes.ResultEvent{Data: tm.EventDataNewBlockHeader{Header: tm.Header{Height: blockHeight}}}
}

type clientMock struct {
	*mock.BlockClientMock
	LatestBlock     int64
	LatestBlockTime time.Time
	newBlocks       chan coretypes.ResultEvent
}

func NewClientMock() *clientMock {
	client := &clientMock{
		newBlocks:       make(chan coretypes.ResultEvent, 1000),
		LatestBlock:     0,
		LatestBlockTime: time.Unix(0, 0),
	}

	subscriptionCtx, subscriptionCancel := context.WithCancel(context.Background())
	blockClientMock := &mock.BlockClientMock{
		LatestSyncInfoFunc: func(context.Context) (*coretypes.SyncInfo, error) {
			return &coretypes.SyncInfo{LatestBlockHeight: client.LatestBlock, LatestBlockTime: client.LatestBlockTime}, nil
		},
		SubscribeFunc: func(_ context.Context, _ string, _ string, out ...int) (<-chan coretypes.ResultEvent, error) {
			eventChan := make(chan coretypes.ResultEvent, out[0])

			go func(ctx context.Context) {
				for block := range client.newBlocks {
					select {
					case eventChan <- block:
						continue
					case <-ctx.Done():
						return
					}
				}
			}(subscriptionCtx)

			return eventChan, nil
		},
		UnsubscribeFunc: func(context.Context, string, string) error {
			subscriptionCancel()
			subscriptionCtx, subscriptionCancel = context.WithCancel(context.Background())
			return nil
		},
	}

	client.BlockClientMock = blockClientMock
	return client
}

func (c *clientMock) NextBlock(height int64) {
	c.LatestBlock = height
	c.LatestBlockTime = time.Now()
	c.newBlocks <- NewBlockHeaderEvent(height)
}
