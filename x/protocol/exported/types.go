package exported

import (
	"encoding/hex"

	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
)

func (p *ProtocolInfo) GetKeyID(pk []byte) multisig.KeyID {
	return multisig.KeyID(hex.EncodeToString(pk))
}
