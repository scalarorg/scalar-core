package exported

import (
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"
)

const (
	// ModuleName exposes scalarnet module name
	ModuleName = "Scalarnet"
	// NativeAsset is the native asset on ScalarNet
	NativeAsset = "scal"

	BaseAsset = "ascal"
)

var (
	Scalarnet = nexus.Chain{
		Name:                  "cosmos|73475",
		SupportsForeignAssets: true,
		KeyType:               tss.None,
		Module:                ModuleName,
	}
)
