package exported

import (
	tss "github.com/axelarnetwork/axelar-core/x/tss/exported"
	"github.com/scalarorg/scalar-core/x/nexus/exported"
)

var (
	// Bitcoin defines properties of the Bitcoin chain
	Bitcoin = exported.Chain{
		Name:                  "Bitcoin",
		SupportsForeignAssets: true,
		KeyType:               tss.Multisig,
		Module:                "btc", // cannot use constant due to import cycle
	}
)
