package types

import (
	"time"

	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/vald/evm/rpc"
)

// EVMConfig contains all EVM module configuration values

type EVMConfig struct {
	ChainID    uint64        `json:"chainId"`
	ID         string        `json:"id"`
	Name       string        `json:"name" mapstructure:"name"`
	Gateway    string        `json:"gateway"`
	Finality   int           `json:"finality"`
	LastBlock  uint64        `json:"lastBlock"`
	GasLimit   uint64        `json:"gasLimit"`
	BlockTime  time.Duration `json:"blockTime"`
	MaxRetry   int           `json:"maxRetry"`
	RetryDelay time.Duration `json:"retryDelay"`
	TxTimeout  time.Duration `json:"txTimeout"`

	RPCAddr          string               `json:"rpcAddr" mapstructure:"rpc_addr"`
	WithBridge       bool                 `mapstructure:"start-with-bridge"`
	L1ChainName      *string              `mapstructure:"l1_chain_name"` // Deprecated: Do not use.
	FinalityOverride rpc.FinalityOverride `mapstructure:"finality_override"`
}

// DefaultConfig returns a configuration populated with default values
func DefaultConfig() []EVMConfig {
	return []EVMConfig{{
		Name:       "Ethereum",
		ID:         "evm|11155111",
		RPCAddr:    "http://127.0.0.1:7545",
		WithBridge: true,
	}}
}

func (c *EVMConfig) ValidateBasic() error {
	_, err := utils.ChainInfoBytesFromID(c.ID)
	if err != nil {
		return err
	}
	return nil
}
