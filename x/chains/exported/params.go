package exported

import (
	"github.com/scalarorg/scalar-core/x/nexus/exported"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"
)

var (
	// Bitcoin defines properties of the Bitcoin chain
	Bitcoin = exported.Chain{
		Name:                  "bitcoin|1",
		SupportsForeignAssets: true,
		KeyType:               tss.Multisig,
		Module:                "chains",
	}

	Ethereum = exported.Chain{
		Name:                  "evm|1",
		SupportsForeignAssets: true,
		KeyType:               tss.Multisig,
		Module:                "chains",
	}
)
