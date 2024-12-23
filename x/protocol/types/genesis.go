package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/scalarorg/scalar-core/x/chains/evm/types"
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
	sepoliaErc20token := evmtypes.ERC20TokenMetadata{
		Asset:        "sBtc",
		ChainID:      sdk.NewInt(1115511),
		TokenAddress: evmtypes.Address(common.HexToAddress("0x5f214989a5f49ab3c56fd5003c2858e24959c018")),
		Status:       evmtypes.Confirmed,
		Details: evmtypes.TokenDetails{
			TokenName: "pBtc",
			Symbol:    "pBtc",
			Decimals:  8,
			Capacity:  sdk.NewInt(100000000),
		},
	}
	sepoliaChain := SupportedChain{
		Token: &SupportedChain_Erc20{Erc20: &sepoliaErc20token},
	}
	protocol := &Protocol{
		Name:           DefaultProtocolName,
		CustodianGroup: types.DefaultCustodianGroup(),
		Chains: []*SupportedChain{
			&sepoliaChain,
		},
	}
	return protocol
}
