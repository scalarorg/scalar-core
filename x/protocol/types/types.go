package types

import (
	"encoding/hex"
	"errors"

	"github.com/scalarorg/scalar-core/utils/evm"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	multisig "github.com/scalarorg/scalar-core/x/multisig/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/scalarorg/scalar-core/x/protocol/exported"
)

func (p *Protocol) IsAssetSupported(destinationChain nexus.ChainName, tokenAddress string) error {
	for _, chain := range p.Chains {
		if chain.Chain == destinationChain {
			if chainsTypes.IsEvmChain(chain.Chain) {
				tokenAddress, err := evm.NormalizeAddress(tokenAddress)
				if err != nil {
					return err
				}

				assetAddress, err := evm.NormalizeAddress(chain.Address)
				if err != nil {
					return err
				}
				if tokenAddress == assetAddress {
					return nil
				}
			}
		}
	}
	return errors.New("asset not supported")
}
func (p *Protocol) GetKeyID() string {
	return hex.EncodeToString(p.CustodianGroup.BtcPubkey)
}

// Get Unique keyId, which later can tell us how to sign btc psbt
func (p *Protocol) ToProtocolInfo() *exported.ProtocolInfo {
	return &exported.ProtocolInfo{
		KeyID:            multisig.KeyID(p.GetKeyID()),
		CustodiansPubkey: p.CustodianGroup.BtcPubkey,
		LiquidityModel:   p.Attribute.Model,
	}
}
