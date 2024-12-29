package jobs

import (
	"context"
	"reflect"

	"github.com/gogo/protobuf/proto"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/client/rpc/cosmos"
)

// EventJob represents a job that listens for specific events with generic type handling
type EventJob struct {
	name          string
	query         string
	networkClient *cosmos.NetworkClient
}

// NewEventJob creates a new event job with typed event handling
func NewEventJob(name string, query cosmos.EventQuery, networkClient *cosmos.NetworkClient) *EventJob {
	return &EventJob{
		name:          name,
		query:         query.ToQuery(),
		networkClient: networkClient,
	}
}

// Run implements the Job interface
func RunJob[T proto.Message](job *EventJob, ctx context.Context, handler func(T) error) {
	ch, err := job.networkClient.Subscribe(ctx, job.name, job.query)
	if err != nil {
		log.Error().Err(err).Str("job", job.name).Msg("Failed to subscribe to events")
		return
	}

	log.Info().Str("job", job.name).Str("query", job.query).Msg("Starting event job")

	for {
		select {
		case <-ctx.Done():
			return

		case resultEvent, ok := <-ch:
			if !ok {
				log.Error().Str("job", job.name).Msg("Event channel closed")
				return
			}

			var event T
			eventType := reflect.TypeOf(event).Elem()

			newEvent := reflect.New(eventType).Interface().(T)

			err := cosmos.ParseEvent(resultEvent.Events, newEvent)
			if err != nil {
				log.Error().Err(err).Str("job", job.name).Msg("Failed to parse event")
				continue
			}

			if err := handler(newEvent); err != nil {
				log.Error().Err(err).Str("job", job.name).Msg("Failed to handle event")
			}
		}
	}
}
