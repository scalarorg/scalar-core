package exported

import (
	tss "github.com/axelarnetwork/axelar-core/x/tss/exported"
	common "github.com/scalarorg/scalar-core/x/common/exported"
)

const (
	// ModuleName exposes scalarnet module name
	ModuleName = "Scalarnet"
)

var (
	// NativeAsset is the native asset on ScalarNet
	NativeAsset = "sclr"

	Scalarnet = common.Chain{
		Name:                  "ScalarNet",
		SupportsForeignAssets: true,
		KeyType:               tss.None,
		Module:                ModuleName,
	}
)
