package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/scalarorg/scalar-core/utils"
)

type BtcChain int32

const (
	MainnetBtcChain  BtcChain = 0
	Testnet3BtcChain BtcChain = 1
	SignetBtcChain   BtcChain = 2
	RegtestBtcChain  BtcChain = 3
	Testnet4BtcChain BtcChain = 4
)

const chaincfgTestnet4ParamsName = "testnet4"

var BtcChainName = map[BtcChain]string{
	MainnetBtcChain:  chaincfg.MainNetParams.Name,
	Testnet3BtcChain: chaincfg.TestNet3Params.Name,
	SignetBtcChain:   chaincfg.SigNetParams.Name,
	RegtestBtcChain:  chaincfg.RegressionNetParams.Name,
	Testnet4BtcChain: chaincfgTestnet4ParamsName,
}

var BtcChainValue = map[string]BtcChain{
	chaincfg.MainNetParams.Name:       MainnetBtcChain,
	chaincfg.TestNet3Params.Name:      Testnet3BtcChain,
	chaincfg.SigNetParams.Name:        SignetBtcChain,
	chaincfg.RegressionNetParams.Name: RegtestBtcChain,
	chaincfgTestnet4ParamsName:        Testnet4BtcChain,
}

func (c BtcChain) String() string {
	return BtcChainName[c]
}

func (c *BtcChain) FromString(s string) error {
	chain, ok := BtcChainValue[s]
	if !ok {
		return fmt.Errorf("invalid btc chain: %s", s)
	}
	*c = chain
	return nil
}

func (c BtcChain) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *BtcChain) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return c.FromString(s)
}

func (c *BtcChain) UnmarshalText(text []byte) error {
	return c.FromString(string(text))
}

type BTCConfig struct {
	ID                   string        `json:"id" mapstructure:"id"`
	ChainID              uint64        `json:"chainID" mapstructure:"chain_id"`
	Name                 string        `json:"name" mapstructure:"name"`
	Chain                BtcChain      `json:"chain" mapstructure:"chain"`
	NetworkKind          NetworkKind   `json:"networkKind" mapstructure:"network_kind"`
	Gateway              string        `json:"gateway" mapstructure:"gateway"` //Taproot address
	Finality             int           `json:"finality"`
	LastBlock            uint64        `json:"lastBlock"`
	GasLimit             uint64        `json:"gasLimit"`
	BlockTime            time.Duration `json:"blockTime"` //Timeout im ms for pending txs
	MaxRetry             int           `json:"maxRetry"`
	RetryDelay           time.Duration `json:"retryDelay"`
	TxTimeout            time.Duration `json:"txTimeout"` //Timeout for send txs (~3s)
	Tag                  string        `json:"tag" mapstructure:"tag"`
	Version              uint8         `json:"version" mapstructure:"version"`
	Rbf                  bool          `json:"rbf"`
	WithBridge           bool          `json:"withBridge" mapstructure:"with_bridge"`
	RpcHost              string        `json:"rpcHost" mapstructure:"rpc_host"`
	RpcPort              int           `json:"rpcPort" mapstructure:"rpc_port"`
	RpcUser              string        `json:"rpcUser" mapstructure:"rpc_user"`
	RpcPass              string        `json:"rpcPassword" mapstructure:"rpc_pass"`
	DisableTLS           bool          `json:"disableTLS" mapstructure:"disable_tls"`
	DisableConnectOnNew  bool          `json:"disableConnectOnNew" mapstructure:"disable_connect_on_new"`
	DisableAutoReconnect bool          `json:"disableAutoReconnect" mapstructure:"disable_auto_reconnect"`
	HttpPostMode         bool          `json:"httpPostMode" mapstructure:"http_post_mode"`
}

// DefaultConfig returns a configuration populated with default values
func DefaultConfig() []BTCConfig {
	return []BTCConfig{{
		ChainID:              4,
		Chain:                Testnet4BtcChain,
		NetworkKind:          Testnet,
		Name:                 "bitcoin-testnet4",
		ID:                   "bitcoin|4",
		Gateway:              "",
		Finality:             10,
		LastBlock:            0,
		GasLimit:             1000000,
		BlockTime:            1000 * time.Millisecond,
		MaxRetry:             3,
		RetryDelay:           100 * time.Millisecond,
		TxTimeout:            3 * time.Second,
		RpcHost:              "http://127.0.0.1:48332",
		RpcUser:              "user",
		RpcPass:              "password",
		Tag:                  "SCALAR",
		Version:              1,
		WithBridge:           false,
		DisableTLS:           true,
		DisableConnectOnNew:  true,
		DisableAutoReconnect: false,
		HttpPostMode:         true,
	}}
}

func (c *BTCConfig) ValidateBasic() error {
	_, err := utils.ChainInfoBytesFromID(c.ID)
	if err != nil {
		return err
	}
	return nil
}
