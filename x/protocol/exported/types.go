package exported

import (
	"encoding/hex"

	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

func (p *ProtocolInfo) GetKeyID(pk []byte) multisig.KeyID {
	return multisig.KeyID(hex.EncodeToString(pk))
}

func (p *ProtocolInfo) IsSupportedChain(chain nexus.ChainName) bool {
	for _, c := range p.MinorAddresses {
		if c.ChainName == chain {
			return true
		}
	}
	return false
}
