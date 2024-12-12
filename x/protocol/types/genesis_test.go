package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	covenanttypes "github.com/scalarorg/scalar-core/x/covenant/types"
	evmtypes "github.com/scalarorg/scalar-core/x/evm/types"
	"github.com/stretchr/testify/assert"
)

func DefaultProtocol() Protocol {
	token := evmtypes.ERC20TokenMetadata{
		Asset:        "pBtc",
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
	protocol := Protocol{
		Name:          DefaultProtocolName,
		CovenantGroup: covenanttypes.DefaultCovenantGroupName,
		Tokens: []evmtypes.ERC20TokenMetadata{
			token,
		},
	}
	return protocol
}
func TestDefaultGenesisState(t *testing.T) {
	assert.NoError(t, NewGenesisState(DefaultProtocol()).Validate())
}
