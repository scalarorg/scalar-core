package exported

import (
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	tss "github.com/scalarorg/scalar-core/x/tss/exported"
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
