package exported

import (
	"github.com/scalarorg/scalar-core/x/nexus/exported"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"
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
