package types

import (
	types "github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/scalarorg/scalar-core/x/nexus/exported"
)

func NewGenesisState(protocols []*Protocol) *GenesisState {
	return &GenesisState{
		Protocols: protocols,
	}
}

func (m GenesisState) Validate() error {
	return nil
}

// DefaultGenesisState returns a genesis state with default parameters
func DefaultGenesisState() *GenesisState {
	return NewGenesisState([]*Protocol{})
}

func DefaultProtocol() *Protocol {
	// sepoliaErc20token := evmtypes.ERC20TokenMetadata{
	// 	Asset:        "sBtc",
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
	sepoliaToken := SupportedChain{
		Chain:   exported.ChainName("evm|111551111"),
		Address: "0xaBbeEcbBfE4732b9DA50CE6b298EDf47E351Fc05",
	}
	bnbToken := SupportedChain{
		Chain:   exported.ChainName("evm|97"),
		Address: "0xaa36A8a917D1804376A7b6Cd54AE1C74Cf83654d",
	}
	protocol := &Protocol{
		Name:              DefaultProtocolName,
		CustodianGroupUID: types.DefaultCustodianGroup().UID,
		Chains: []*SupportedChain{
			&sepoliaToken, &bnbToken,
		},
	}
	return protocol
}
