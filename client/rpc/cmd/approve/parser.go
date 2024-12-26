package approve

import (
	"github.com/scalarorg/scalar-core/client/rpc/utils"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

func ParseDestCallApproved(eventData map[string]interface{}) chainsTypes.DestCallApproved {

	event := chainsTypes.DestCallApproved{}

	if chainID, ok := eventData["chain"].(string); ok {
		event.Chain = nexus.ChainName(chainID)
	}

	if commandID, ok := eventData["command_id"].(string); ok {
		bytes, err := utils.StringArrayToBytes(commandID)
		if err != nil {
			panic(err)
		}

		event.CommandID = chainsTypes.CommandID(bytes)
	}

	if contractAddress, ok := eventData["contract_address"].(string); ok {
		event.ContractAddress = contractAddress
	}

	if destAddress, ok := eventData["destination_chain"].(string); ok {
		event.DestinationChain = nexus.ChainName(destAddress)
	}

	// Parse each field from the event
	if eventID, ok := eventData["event_id"].(string); ok {
		event.EventID = chainsTypes.EventID(eventID)
	}

	if payloadHash, ok := eventData["payload_hash"].(string); ok {
		bytes, err := utils.StringArrayToBytes(payloadHash)
		if err != nil {
			panic(err)
		}

		event.PayloadHash = chainsTypes.Hash(bytes)
	}

	if sender, ok := eventData["sender"].(string); ok {
		event.Sender = sender
	}

	return event
}
