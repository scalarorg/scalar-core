package jobs

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/client/rpc/cosmos"
)

// EventJob represents a job that listens for specific events with generic type handling
type EventJob struct {
	name          string
	query         string
	networkClient *cosmos.NetworkClient
	handler       func(proto.Message) error
}

// NewEventJob creates a new event job with typed event handling
func NewEventJob(name, query string, networkClient *cosmos.NetworkClient, handler func(proto.Message) error) *EventJob {
	return &EventJob{
		name:          name,
		query:         query,
		networkClient: networkClient,
		handler:       handler,
	}
}

// Run implements the Job interface
func (j *EventJob) Run(ctx context.Context) error {
	ch, err := j.networkClient.Subscribe(ctx, j.name, j.query)
	if err != nil {
		return fmt.Errorf("failed to subscribe to events for job %s: %w", j.name, err)
	}

	log.Info().Str("job", j.name).Msg("Starting event job")

	for {
		select {
		case <-ctx.Done():
			return nil

		case resultEvent, ok := <-ch:
			if !ok {
				return fmt.Errorf("event channel closed for job %s", j.name)
			}

			// Debug log the raw event
			log.Debug().
				Str("job", j.name).
				Interface("event", resultEvent).
				Msg("Received event")

			// Process each event in the result
			for eventType, attrs := range resultEvent.Events {
				if eventType == j.name {
					_ = attrs
					log.Info().Str("event", eventType).Msg("Received event")
					// parsedEvent := funcs.Must(sdk.ParseTypedEvent(abci.Event{
					// 	Type:       eventType,
					// 	Attributes: attrs,
					// })).(T)

					// if err := j.handler(parsedEvent); err != nil {
					// 	log.Error().Err(err).
					// 		Str("job", j.name).
					// 		Msg("Failed to handle event")
					// }
				}
			}
		}
	}
}
