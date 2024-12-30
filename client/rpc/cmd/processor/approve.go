package processor

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

func (p *Processor) ProcessDestCallApproved(event *chainsTypes.DestCallApproved) error {
	ctx := context.Background()
	fmt.Println("Processing DestCallApproved event")
	fmt.Println("chain: ", event.Chain)
	fmt.Println("event_id: ", event.EventID)
	fmt.Printf("command_id: %x\n", event.CommandID)
	fmt.Println("sender: ", event.Sender)
	fmt.Println("destination_chain: ", event.DestinationChain)
	fmt.Println("contract_address: ", event.ContractAddress)
	fmt.Printf("payload_hash: %x\n", event.PayloadHash)

	res, err := p.networkClient.SignBTCCommandsRequest(ctx, event.DestinationChain.String())
	if err != nil {
		return err
	}

	log.Info().Msgf("[Processor] [ProcessDestCallApproved] txRes: %v", res)

	return nil
}
