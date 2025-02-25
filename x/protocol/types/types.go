package types

import (
	"errors"

	"github.com/scalarorg/scalar-core/utils/evm"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
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

// Get Unique keyId, which later can tell us how to sign btc psbt
func (p *Protocol) ToProtocolInfo() *exported.ProtocolInfo {
	minorAddreses := make([]*exported.MinorAddress, len(p.Chains))
	for i, chain := range p.Chains {
		minorAddreses[i] = &exported.MinorAddress{
			ChainName: chain.Chain,
			Address:   chain.Address,
		}
	}

	return &exported.ProtocolInfo{
		// KeyID:            multisig.KeyID(keyID),
		CustodiansGroupUID: p.CustodianGroupUID,
		LiquidityModel:     p.Attributes.Model,
		OriginChain:        p.Asset.Chain,
		Symbol:             p.Asset.Name,
		MinorAddresses:     minorAddreses,
	}
}

func (p *Protocol) AddSupportedChain(chain nexus.ChainName, address string, name string) {
	p.Chains = append(p.Chains, &exported.SupportedChain{
		Chain:   chain,
		Address: address,
		Name:    name,
	})
}
