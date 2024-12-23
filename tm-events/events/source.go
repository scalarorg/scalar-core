package events

import (
	"context"
	"fmt"
	"time"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/scalarorg/scalar-core/utils/log"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/libs/pubsub/query"
	tm "github.com/tendermint/tendermint/types"

	"github.com/scalarorg/scalar-core/utils"
)

//go:generate moq -pkg mock -out ./mock/source.go . BlockSource BlockClient BlockResultClient BlockNotifier

// WebsocketQueueSize is set to the maximum length to prevent Tendermint from closing the connection due to congestion
const WebsocketQueueSize = 32768

type dialOptions struct {
	timeout   time.Duration
	retries   int
	backOff   time.Duration
	keepAlive time.Duration
}

// DialOption for Tendermint connections
type DialOption struct {
	apply func(options dialOptions) dialOptions
}

// Timeout sets the time after which the call to Tendermint is cancelled
func Timeout(timeout time.Duration) DialOption {
	return DialOption{
		apply: func(options dialOptions) dialOptions {
			options.timeout = timeout
			return options
		},
	}
}

// Retries sets the number of times a Tendermint call is retried
func Retries(retries int) DialOption {
	return DialOption{
		apply: func(options dialOptions) dialOptions {
			options.retries = retries
			return options
		},
	}
}

// KeepAlive sets the time after which contact to Tendermint is reestablished if no there is no communication
func KeepAlive(interval time.Duration) DialOption {
	return DialOption{
		apply: func(options dialOptions) dialOptions {
			options.keepAlive = interval
			return options
		},
	}
}

// BackOff sets the time to wait until retrying a failed call to Tendermint
func BackOff(backOff time.Duration) DialOption {
	return DialOption{
		apply: func(options dialOptions) dialOptions {
			options.backOff = backOff
			return options
		},
	}
}

// BlockNotifier notifies the caller of new blocks
type BlockNotifier interface {
	BlockHeights(ctx context.Context) (<-chan int64, <-chan error)
	Done() <-chan struct{}
}

type eventblockNotifier struct {
	client            SubscriptionClient
	query             string
	timeout           time.Duration
	backOff           time.Duration
	retries           int
	keepAliveInterval time.Duration
	done              chan struct{}
}

func newEventBlockNotifier(client SubscriptionClient, options ...DialOption) *eventblockNotifier {
	var opts dialOptions
	for _, option := range options {
		opts = option.apply(opts)
	}
	return &eventblockNotifier{
		client:            client,
		query:             query.MustParse(fmt.Sprintf("%s='%s'", tm.EventTypeKey, tm.EventNewBlockHeader)).String(),
		backOff:           opts.backOff,
		timeout:           opts.timeout,
		retries:           opts.retries,
		keepAliveInterval: opts.keepAlive,
		done:              make(chan struct{}),
	}
}

func (b *eventblockNotifier) BlockHeights(ctx context.Context) (<-chan int64, <-chan error) {
	blocks := make(chan int64)
	errChan := make(chan error, 1)

	var keepAlive context.Context
	var keepAliveCancel context.CancelFunc = func() {}
	go func() {
		defer close(b.done)
		defer func() { keepAliveCancel() }() // the cancel function gets reassigned, so call it indirectly
		defer close(blocks)
		defer log.FromCtx(ctx).Info("stopped listening to events")
		defer b.tryUnsubscribe(context.Background(), b.timeout)

		eventChan, err := b.subscribe(ctx)
		if err != nil {
			errChan <- err
			return
		}

		for {
			keepAlive, keepAliveCancel = ctxWithTimeout(context.Background(), b.keepAliveInterval)
			var blockHeaderEvent tm.EventDataNewBlockHeader
			select {
			case <-keepAlive.Done():
				log.FromCtx(ctx).Debug(fmt.Sprintf("no block event in %.2f seconds", b.keepAliveInterval.Seconds()))
				b.tryUnsubscribe(ctx, b.timeout)
				eventChan, err = b.subscribe(ctx)
				if err != nil {
					errChan <- err
					return
				}

			case event := <-eventChan:
				var ok bool
				blockHeaderEvent, ok = event.Data.(tm.EventDataNewBlockHeader)
				if !ok {
					errChan <- fmt.Errorf("event data is of type %T, expected %T", event.Data, tm.EventDataNewBlockHeader{})
					return
				}
			case <-ctx.Done():
				return
			}

			select {
			case blocks <- blockHeaderEvent.Header.Height:
				break
			case <-ctx.Done():
				return
			}
		}
	}()

	return blocks, errChan
}

func (b *eventblockNotifier) subscribe(ctx context.Context) (<-chan coretypes.ResultEvent, error) {
	backOff := utils.LinearBackOff(b.backOff)
	for i := 0; i <= b.retries; i++ {
		ctx, cancel := ctxWithTimeout(ctx, b.timeout)
		eventChan, err := b.client.Subscribe(ctx, "", b.query, WebsocketQueueSize)
		cancel()
		if err == nil {
			log.FromCtx(ctx).Debug(fmt.Sprintf("subscribed to query \"%s\"", b.query))
			return eventChan, nil
		}
		log.FromCtx(ctx).Debug(sdkerrors.Wrapf(err, "failed to subscribe to query \"%s\", attempt %d", b.query, i+1).Error())
		time.Sleep(backOff(i))
	}
	return nil, fmt.Errorf("aborting Tendermint block header subscription after %d attemts ", b.retries+1)
}

func (b *eventblockNotifier) tryUnsubscribe(ctx context.Context, timeout time.Duration) {
	ctx, cancel := ctxWithTimeout(ctx, timeout)
	defer cancel()

	if b.client == nil {
		return
	}

	// this unsubscribe is a best-effort action, we still try to continue as usual if it fails, so errors are only logged
	err := b.client.Unsubscribe(ctx, "", b.query)
	if err != nil {
		log.FromCtx(ctx).Info(sdkerrors.Wrapf(err, "could not unsubscribe from query \"%s\"", b.query).Error())
		return
	}
	log.FromCtx(ctx).Debug(fmt.Sprintf("unsubscribed from query \"%s\"", b.query))
}

func (b *eventblockNotifier) Done() chan struct{} {
	<-b.done
	done := make(chan struct{})
	close(done)
	return done
}

type queryBlockNotifier struct {
	client            SyncInfoClient
	retries           int
	timeout           time.Duration
	backOff           time.Duration
	keepAliveInterval time.Duration
}

func newQueryBlockNotifier(client SyncInfoClient, options ...DialOption) *queryBlockNotifier {
	var opts dialOptions
	for _, option := range options {
		opts = option.apply(opts)
	}

	return &queryBlockNotifier{
		client:            client,
		retries:           opts.retries,
		backOff:           opts.backOff,
		timeout:           opts.timeout,
		keepAliveInterval: opts.keepAlive,
	}
}

func (q *queryBlockNotifier) BlockHeights(ctx context.Context) (<-chan int64, <-chan error) {
	blocks := make(chan int64)
	errChan := make(chan error, 1)

	go func() {
		defer close(blocks)
		defer log.FromCtx(ctx).Info("stopped block query")

		keepAlive, keepAliveCancel := ctxWithTimeout(context.Background(), q.keepAliveInterval)
		defer func() { keepAliveCancel() }() // the cancel function might get reassigned, so call it indirectly

		for {
			var latest int64
			var err error
			select {
			case <-keepAlive.Done():
				latest, err = q.latestFromSyncStatus(ctx)
				if err != nil {
					errChan <- err
					return
				}
				keepAlive, keepAliveCancel = ctxWithTimeout(context.Background(), q.keepAliveInterval)
			case <-ctx.Done():
				return
			}

			select {
			case blocks <- latest:
				break
			case <-ctx.Done():
				return
			}
		}
	}()
	return blocks, errChan
}

func (q *queryBlockNotifier) latestFromSyncStatus(ctx context.Context) (int64, error) {
	backOff := utils.LinearBackOff(q.backOff)
	for i := 0; i <= q.retries; i++ {
		ctx, cancel := ctxWithTimeout(ctx, q.timeout)
		syncInfo, err := q.client.LatestSyncInfo(ctx)
		cancel()
		if err == nil {
			return syncInfo.LatestBlockHeight, nil
		}

		log.FromCtx(ctx).Info(sdkerrors.Wrapf(err, "failed to retrieve node status, attempt %d", i+1).Error())
		time.Sleep(backOff(i))
	}
	return 0, fmt.Errorf("aborting sync status retrieval after %d attemts ", q.retries+1)
}

// SyncInfoClient can query the node's sync info
type SyncInfoClient interface {
	LatestSyncInfo(ctx context.Context) (*coretypes.SyncInfo, error)
}

// SubscriptionClient subscribes to and unsubscribes from Tendermint events
type SubscriptionClient interface {
	Subscribe(ctx context.Context, subscriber, query string, outCapacity ...int) (out <-chan coretypes.ResultEvent, err error)
	Unsubscribe(ctx context.Context, subscriber, query string) error
}

// BlockClient is both SyncInfoClient and SubscriptionClient
type BlockClient interface {
	SyncInfoClient
	SubscriptionClient
}

var _ BlockNotifier = &Notifier{}

// Notifier can notify a consumer about new blocks
type Notifier struct {
	client         BlockClient
	queryNotifier  *queryBlockNotifier
	eventsNotifier *eventblockNotifier
	start          int64
	timeout        time.Duration
}

// Done returns a channel that is closed when the Notifier has completed cleanup
func (b *Notifier) Done() <-chan struct{} {
	return b.eventsNotifier.Done()
}

// NewBlockNotifier returns a new BlockNotifier instance.
// It accepts dial options that configure the request management behaviour.
func NewBlockNotifier(client BlockClient, options ...DialOption) *Notifier {
	var opts dialOptions
	for _, option := range options {
		opts = option.apply(opts)
	}

	return &Notifier{
		client:         client,
		timeout:        opts.timeout,
		eventsNotifier: newEventBlockNotifier(client, options...),
		queryNotifier:  newQueryBlockNotifier(client, options...),
	}
}

// StartingAt sets the start block from which to receive notifications
func (b *Notifier) StartingAt(block int64) *Notifier {
	if block > 0 {
		b.start = block
	}
	return b
}

func (b *Notifier) getLatestBlockHeight() (int64, error) {
	ctx, cancel := ctxWithTimeout(context.Background(), b.timeout)
	defer cancel()
	info, err := b.client.LatestSyncInfo(ctx)
	if err != nil {
		return 0, err
	}

	return info.LatestBlockHeight, nil
}

// BlockHeights returns a channel with the block heights from the beginning of the chain to all newly discovered blocks.
// Optionally, starts at the given start block.
func (b *Notifier) BlockHeights(ctx context.Context) (<-chan int64, <-chan error) {
	ctx = log.Append(ctx, "listener", "block notifier")

	errChan := make(chan error, 1)
	var shutdown context.CancelFunc
	ctx, shutdown = context.WithCancel(ctx)

	// if no start block has been set explicitly try to fetch the latest block height
	if b.start == 0 {
		b.start = -1
		height, err := b.getLatestBlockHeight()
		if err == nil {
			b.start = height
		}
	}

	blocksFromQuery, queryErrs := b.queryNotifier.BlockHeights(ctx)
	blocksFromEvents, eventErrs := b.eventsNotifier.BlockHeights(ctx)

	go handleErrors(ctx, shutdown, eventErrs, queryErrs, errChan)

	log.FromCtx(ctx).Info(fmt.Sprintf("syncing blocks starting with block %d", b.start))

	blocks := make(chan int64)
	go b.pipeLatestBlock(ctx, blocksFromQuery, blocksFromEvents, blocks)

	return blocks, errChan
}

func handleErrors(ctx context.Context, shutdown context.CancelFunc, eventErrs <-chan error, queryErrs <-chan error, errChan chan error) {
	defer shutdown()

	for {
		// the query notifier is more reliable but slower, so we still continue on as long as it's available
		select {
		case <-ctx.Done():
			return
		case err := <-queryErrs:
			errChan <- err
			return
		case err := <-eventErrs:
			log.FromCtx(ctx).Error(sdkerrors.Wrapf(err, "cannot receive new blocks from events, falling back on querying actively for blocks").Error())
		}
	}
}

func (b *Notifier) pipeLatestBlock(ctx context.Context, fromQuery <-chan int64, fromEvents <-chan int64, blockHeights chan int64) {
	defer close(blockHeights)
	defer log.FromCtx(ctx).Info("stopped block sync")

	pendingBlockHeight := b.start
	latestBlock := int64(0)
	for {
		for {
			select {
			case <-ctx.Done():
				return
			case block := <-fromQuery:
				if latestBlock > block || pendingBlockHeight > block {
					continue
				}
				latestBlock = block
			case block := <-fromEvents:
				if latestBlock > block || pendingBlockHeight > block {
					continue
				}
				latestBlock = block
			}
			break
		}

		log.FromCtx(ctx).Debug(fmt.Sprintf("found latest block %d", latestBlock))

		for pendingBlockHeight <= latestBlock {
			select {
			case <-ctx.Done():
				return
			case blockHeights <- pendingBlockHeight:
				pendingBlockHeight++
			}
		}
	}
}

// BlockSource returns all block results sequentially
type BlockSource interface {
	BlockResults(ctx context.Context) (<-chan *coretypes.ResultBlockResults, <-chan error)
	Done() <-chan struct{}
}

type blockSource struct {
	notifier BlockNotifier
	client   BlockResultClient
	shutdown context.CancelFunc

	// dial options
	retries int
	backOff time.Duration
	timeout time.Duration
}

// BlockResultClient can query for the block results of a specific block
type BlockResultClient interface {
	BlockResults(ctx context.Context, height *int64) (*coretypes.ResultBlockResults, error)
}

// NewBlockSource returns a new BlockSource instance.
// It accepts dial options that configure the request management behaviour.
func NewBlockSource(client BlockResultClient, notifier BlockNotifier, options ...DialOption) BlockSource {
	var opts dialOptions
	for _, option := range options {
		opts = option.apply(opts)
	}

	return &blockSource{
		client:   client,
		notifier: notifier,
		retries:  opts.retries,
		backOff:  opts.backOff,
		timeout:  opts.timeout,
	}
}

func (b *blockSource) Done() <-chan struct{} {
	return b.notifier.Done()
}

// BlockResults returns a channel of block results. Blocks are pushed into the channel sequentially as they are discovered
func (b *blockSource) BlockResults(ctx context.Context) (<-chan *coretypes.ResultBlockResults, <-chan error) {
	ctx = log.Append(ctx, "listener", "block results")

	// either the notifier or the block source could push an error at the same time, so we need to make sure we don't block
	errChan := make(chan error, 2)
	// notifier is an external dependency, so it's lifetime should be managed by the calling process, i.e. the given context, as well
	blockHeights, notifyErrs := b.notifier.BlockHeights(ctx)

	ctx, b.shutdown = context.WithCancel(ctx)
	go func() {
		defer b.shutdown()

		select {
		case err := <-notifyErrs:
			errChan <- err
			return
		case <-ctx.Done():
			return
		}
	}()

	blockResults := make(chan *coretypes.ResultBlockResults)
	go b.streamBlockResults(ctx, blockHeights, blockResults, errChan)

	return blockResults, errChan
}

func (b *blockSource) streamBlockResults(ctx context.Context, blockHeights <-chan int64, blocks chan *coretypes.ResultBlockResults, errChan chan error) {
	defer close(blocks)
	defer b.shutdown()

	for {
		select {
		case height, ok := <-blockHeights:
			if !ok {
				errChan <- fmt.Errorf("cannot detect new blocks anymore")
				return
			}
			result, err := b.fetchBlockResults(ctx, &height)
			if err != nil {
				errChan <- err
				return
			}

			blocks <- result
		case <-ctx.Done():
			return
		}
	}
}

func (b *blockSource) fetchBlockResults(ctx context.Context, height *int64) (*coretypes.ResultBlockResults, error) {
	// if the client connection or subscription setup fails then retry
	backOff := utils.LinearBackOff(b.backOff)
	for i := 0; i <= b.retries; i++ {
		ctx, cancel := ctxWithTimeout(ctx, b.timeout)

		res, err := b.client.BlockResults(ctx, height)
		cancel()
		if err == nil {
			log.FromCtx(ctx).Debugf("fetched result for block height %d", *height)
			return res, nil
		}
		log.FromCtx(ctx).Debug(sdkerrors.Wrapf(err, "failed to result for block height %d, attempt %d", *height, i+1).Error())
		time.Sleep(backOff(i))
	}
	return nil, fmt.Errorf("aborting block result fetching after %d attemts ", b.retries+1)
}

func ctxWithTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	if timeout == 0 {
		return context.WithCancel(ctx)
	}
	return context.WithTimeout(ctx, timeout)
}
