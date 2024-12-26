package config

import (
	"encoding/json"
	"os"

	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
	"github.com/scalarorg/scalar-core/client/rpc/codec"
)

const (
	DEFAULT_GAS_ADJUSTMENT = 1.2
)

var GlobalConfig *CosmosNetworkConfig

var GlobalTxConfig = tx.NewTxConfig(codec.GetProtoCodec(), []signing.SignMode{signing.SignMode_SIGN_MODE_DIRECT})

type CosmosNetworkConfig struct {
	ChainID       uint64   `mapstructure:"chain_id" json:"chain_id"`
	ID            string   `mapstructure:"id" json:"id"`
	Name          string   `mapstructure:"name" json:"name"`
	RPCUrl        string   `mapstructure:"rpc_url" json:"rpc_url"`
	LCDUrl        string   `mapstructure:"lcd_url" json:"lcd_url"`
	WSUrl         string   `mapstructure:"ws_url" json:"ws_url"`
	Denom         string   `mapstructure:"denom" json:"denom"`
	Mnemonic      string   `mapstructure:"mnemonic" json:"mnemonic"`
	GasPrice      float64  `mapstructure:"gas_price" json:"gas_price"`
	BroadcastMode string   `mapstructure:"broadcast_mode" json:"broadcast_mode"`
	MaxRetries    int      `mapstructure:"max_retries" json:"max_retries"`
	RetryInterval int64    `mapstructure:"retry_interval" json:"retry_interval"` //milliseconds
	PrivateKeys   []string `mapstructure:"private_keys" json:"private_keys"`
	PublicKeys    []string `mapstructure:"public_keys" json:"public_keys"`
	SignerNetwork string   `mapstructure:"signer_network" json:"signer_network"`
	GasAdjustment float64  `mapstructure:"gas_adjustment" json:"gas_adjustment"`
}

func (c *CosmosNetworkConfig) GetFamily() string {
	return chain.ChainTypeCosmos.String()
}

func (c *CosmosNetworkConfig) GetChainId() uint64 {
	return c.ChainID
}

func (c *CosmosNetworkConfig) GetId() string {
	return c.ID
}

func (c *CosmosNetworkConfig) GetName() string {
	return c.Name
}

func ReadConfig(jsonPath string) error {
	file, err := os.Open(jsonPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := &CosmosNetworkConfig{}
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}
	cfg.GasAdjustment = DEFAULT_GAS_ADJUSTMENT
	GlobalConfig = cfg
	return nil
}
