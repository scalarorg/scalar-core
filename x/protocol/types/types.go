package types

import (
	"errors"

	"github.com/scalarorg/scalar-core/utils/evm"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
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
