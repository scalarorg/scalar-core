package jobs

import (
	"reflect"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/client/rpc/cosmos"
	"github.com/scalarorg/scalar-core/x/chains/types"
)

func RunJob2[T proto.Message](rawData map[string][]string, handler func(T) error) {
	var event T
	eventType := reflect.TypeOf(event).Elem()

	// Create a new instance of `T`
	newEvent := reflect.New(eventType).Interface().(T)

	// Parse the event
	err := cosmos.ParseEvent(rawData, newEvent)
	if err != nil {
		panic(err)
	}

	if err := handler(newEvent); err != nil {
		panic(err)
	}
}

func TestRunJob(t *testing.T) {

	var mockEventData = map[string][]string{
		"scalar.chains.v1beta1.DestCallApproved.chain":             []string{"bitcoin|4"},
		"scalar.chains.v1beta1.DestCallApproved.command_id":        []string{"[208,202,25,221,196,160,174,119,189,83,196,188,71,103,47,66,122,159,131,27,169,93,167,74,80,221,239,156,253,73,36,215]"},
		"scalar.chains.v1beta1.DestCallApproved.contract_address":  []string{"0x1F98C06D8734D5A9FF0b53e3294626E62e4d232C"},
		"scalar.chains.v1beta1.DestCallApproved.destination_chain": []string{"evm|11155111"},
		"scalar.chains.v1beta1.DestCallApproved.event_id":          []string{"0x5188eea7ceb9f585f5ba8a2abebb82f9850dd671b6e2926263674af6882fd3f6-0"},
		"scalar.chains.v1beta1.DestCallApproved.payload_hash":      []string{"[145,20,65,62,165,139,208,34,178,219,121,26,79,239,29,140,77,7,161,240,162,55,12,75,244,6,181,67,98,174,165,18]"},
		"scalar.chains.v1beta1.DestCallApproved.sender":            []string{"tb1q2rwweg2c48y8966qt4fzj0f4zyg9wty7tykzwg"},
	}

	RunJob2(mockEventData, func(event *types.DestCallApproved) error {
		log.Info().
			Any("chain", event.Chain.String()).
			Any("event_id", event.EventID).
			Any("command_id", event.CommandID).
			Any("sender", event.Sender).
			Any("destination_chain", event.DestinationChain.String()).
			Any("contract_address", event.ContractAddress).
			Any("payload_hash", event.PayloadHash).
			Msg("Processing DestCallApproved event")
		return nil
	})
}
