package exported

import (
	"github.com/axelarnetwork/axelar-core/x/nexus/exported"
	tss "github.com/axelarnetwork/axelar-core/x/tss/exported"
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
