package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func DefaultProtocol() Protocol {
	// token := chainsTypes.ERC20TokenMetadata{
	// 	Asset:        "pBtc",
	// 	ChainID:      sdk.NewInt(1115511),
	// 	TokenAddress: chainsTypes.Address(common.HexToAddress("0x5f214989a5f49ab3c56fd5003c2858e24959c018")),
	// 	Status:       chainsTypes.Confirmed,
	// 	Details: chainsTypes.TokenDetails{
	// 		TokenName: "pBtc",
	// 		Symbol:    "pBtc",
	// 		Decimals:  8,
	// 		Capacity:  sdk.NewInt(100000000),
	// 	},
	// }
	protocol := Protocol{
		// Name:          DefaultProtocolName,
		// CovenantGroup: covenanttypes.DefaultCovenantGroupName,
		// Tokens: []chainsTypes.ERC20TokenMetadata{
		// 	token,
		// },
	}
	return protocol
}
func TestDefaultGenesisState(t *testing.T) {
	assert.NoError(t, NewGenesisState(DefaultProtocol()).Validate())
}
