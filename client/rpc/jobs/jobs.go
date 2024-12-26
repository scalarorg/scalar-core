package jobs

import (
	"context"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/client/rpc/cosmos"
)

// EventJob represents a job that listens for specific events with generic type handling
type EventJob struct {
	name          string
	query         string
	group         string
	networkClient *cosmos.NetworkClient
}

// NewEventJob creates a new event job with typed event handling
func NewEventJob(name, query, group string, networkClient *cosmos.NetworkClient) *EventJob {
	return &EventJob{
		name:          name,
		query:         query,
		group:         group,
		networkClient: networkClient,
	}
}

// Run implements the Job interface
func RunJob[T any](job *EventJob, ctx context.Context, parser func(map[string]interface{}) T, handler func(T) error) {

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

			// Create a map with single values instead of arrays
			eventData := make(map[string]interface{})
			for eventType, attrs := range resultEvent.Events {
				if strings.Contains(eventType, job.group) {
					keys := strings.Split(eventType, ".")
					key := keys[len(keys)-1]
					if len(attrs) > 0 {
						// Handle byte arrays (like command_id) by decoding from hex
						eventData[key] = attrs[0]
					}
				}
			}

			typedEvent := parser(eventData)
			if err := handler(typedEvent); err != nil {
				log.Error().Err(err).Str("job", job.name).Msg("Failed to handle event")
			}
		}
	}
}
