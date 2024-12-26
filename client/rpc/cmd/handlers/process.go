package handlers

import (
	"fmt"

	"github.com/rs/zerolog/log"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

func ProcessDestCallApproved(event *chainsTypes.DestCallApproved) error {
	fmt.Println("chain: ", event.Chain)
	fmt.Println("event_id: ", event.EventID)
	fmt.Println("command_id: ", event.CommandID)
	fmt.Println("sender: ", event.Sender)
	fmt.Println("destination_chain: ", event.DestinationChain)
	fmt.Println("contract_address: ", event.ContractAddress)
	fmt.Println("payload_hash: ", event.PayloadHash)

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
}
