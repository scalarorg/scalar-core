package exported

import (
	tss "github.com/axelarnetwork/axelar-core/x/tss/exported"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
)

const (
	// ModuleName exposes scalarnet module name
	ModuleName = "Scalarnet"
)

var (
	// NativeAsset is the native asset on ScalarNet
	NativeAsset = "scal"

	Scalarnet = nexus.Chain{
		Name:                  "ScalarNet",
		SupportsForeignAssets: true,
		KeyType:               tss.None,
		Module:                ModuleName,
	}
)
