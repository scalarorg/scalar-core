package exported

import (
	"github.com/scalarorg/scalar-core/x/nexus/exported"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"
)

var (
	// Ethereum defines properties of the Ethereum chain
	Ethereum = exported.Chain{
		Name:                  "ethereum-123",
		SupportsForeignAssets: true,
		KeyType:               tss.Multisig,
		Module:                "evm", // cannot use constant due to import cycle
	}
)
