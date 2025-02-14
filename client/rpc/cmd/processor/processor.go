package processor

import (
	"github.com/scalarorg/scalar-core/client/rpc/cosmos"
)

type Processor struct {
	networkClient *cosmos.NetworkClient
}

func NewProcessor(networkClient *cosmos.NetworkClient) *Processor {
	return &Processor{networkClient: networkClient}
}
