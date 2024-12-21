package types

import (
	"github.com/scalarorg/scalar-core/x/covenant/types"
)

func NewGenesisState(protocols []*Protocol) *GenesisState {
	return &GenesisState{
		Protocols: protocols,
	}
}

// DefaultGenesisState returns a genesis state with default parameters
func DefaultGenesisState() *GenesisState {
	return NewGenesisState([]*Protocol{DefaultProtocol()})
}
func (m GenesisState) Validate() error {
	return nil
}

func DefaultProtocol() *Protocol {
	// token := evmtypes.ERC20TokenMetadata{
	// 	Asset:        "pBtc",
	// 	ChainID:      sdk.NewInt(1115511),
	// 	TokenAddress: evmtypes.Address(common.HexToAddress("0x5f214989a5f49ab3c56fd5003c2858e24959c018")),
	// 	Status:       evmtypes.Confirmed,
	// 	Details: evmtypes.TokenDetails{
	// 		TokenName: "pBtc",
	// 		Symbol:    "pBtc",
	// 		Decimals:  8,
	// 		Capacity:  sdk.NewInt(100000000),
	// 	},
	// }
	protocol := &Protocol{
		Name:           DefaultProtocolName,
		CustodianGroup: types.DefaultCustodianGroup(),
		Chains:         []*SupportedChain{},
	}
	return protocol
}
