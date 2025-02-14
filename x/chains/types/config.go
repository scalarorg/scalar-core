package types

import (
	"github.com/scalarorg/bitcoin-vault/go-utils/types"
	"github.com/scalarorg/scalar-core/utils"
)

type ChainConfig struct {
	ID           string            `json:"id"`
	ChainID      uint64            `json:"chain_id"`
	Name         string            `json:"name"`
	NetworkKind  types.NetworkKind `json:"network_kind"`
	Gateway      string            `json:"gateway"`
	AuthWeighted string            `json:"authWeighted"`
	Metadata     map[string]string `json:"metadata"`
}

// DefaultConfig returns a configuration populated with default values
func DefaultConfig() []ChainConfig {
	return []ChainConfig{{
		ChainID:     4,
		NetworkKind: types.NetworkKindTestnet,
		Name:        "bitcoin-testnet4",
		ID:          "bitcoin|4",
	}}
}

func (c *ChainConfig) ValidateBasic() error {
	_, err := utils.ChainInfoBytesFromID(c.ID)
	if err != nil {
		return err
	}

	// TODO: Check if evm chain -> validate gateway
	return nil
}
