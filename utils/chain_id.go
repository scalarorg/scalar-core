package utils

import "github.com/scalarorg/bitcoin-vault/go-utils/chain"

func ChainInfoBytesFromID(chainID string) (chain.ChainInfoBytes, error) {
	return ChainInfoBytesFromString(chainID)
}

func ChainInfoBytesFromString(chainID string) (chain.ChainInfoBytes, error) {
	chainInfoBytes := chain.ChainInfoBytes{}
	err := chainInfoBytes.FromString(chainID)
	if err != nil {
		return chain.ChainInfoBytes{}, err
	}
	return chainInfoBytes, nil
}
