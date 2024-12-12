package multisig

import "github.com/scalarorg/scalar-core/x/tss/tofnd"

//go:generate moq -pkg mock -out ./mock/rpcClient.go . Client

// Client defines the interface of a grpc client to communicate with tofnd Multisig service
type Client interface {
	tofnd.MultisigClient
}
