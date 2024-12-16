package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/scalarorg/bitcoin-vault/go-utils/chain"
)

type BtcChain int32

const (
	MainnetBtcChain  BtcChain = 0
	Testnet3BtcChain BtcChain = 1
	SignetBtcChain   BtcChain = 2
	RegtestBtcChain  BtcChain = 3
	Testnet4BtcChain BtcChain = 4
)

var BtcChainName = map[BtcChain]string{
	MainnetBtcChain:  chaincfg.MainNetParams.Name,
	Testnet3BtcChain: chaincfg.TestNet3Params.Name,
	SignetBtcChain:   chaincfg.SigNetParams.Name,
	RegtestBtcChain:  chaincfg.RegressionNetParams.Name,
	Testnet4BtcChain: chaincfg.TestNet3Params.Name, // TODO: Add TestNet4Params
}

var BtcChainValue = map[string]BtcChain{
	chaincfg.MainNetParams.Name:       MainnetBtcChain,
	chaincfg.TestNet3Params.Name:      Testnet3BtcChain,
	chaincfg.SigNetParams.Name:        SignetBtcChain,
	chaincfg.RegressionNetParams.Name: RegtestBtcChain,
	"testnet4":                        Testnet4BtcChain,
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

type BTCConfig struct {
	ChainID     uint64          `json:"chainID" mapstructure:"chainID"`
	ChainInfo   chain.ChainInfo `json:"chainInfo"`
	Chain       BtcChain        `json:"chain" mapstructure:"chain"`
	NetworkKind NetworkKind     `json:"networkKind" mapstructure:"networkKind"`
	Name        string          `json:"name" mapstructure:"name"`
	ID          string          `json:"id"`
	Gateway     string          `json:"gateway" mapstructure:"gateway"` //Taproot address
	Finality    int             `json:"finality"`
	LastBlock   uint64          `json:"lastBlock"`
	GasLimit    uint64          `json:"gasLimit"`
	BlockTime   time.Duration   `json:"blockTime"` //Timeout im ms for pending txs
	MaxRetry    int             `json:"maxRetry"`
	RetryDelay  time.Duration   `json:"retryDelay"`
	TxTimeout   time.Duration   `json:"txTimeout"` //Timeout for send txs (~3s)
	Tag         string          `json:"tag" mapstructure:"tag"`
	Version     byte            `json:"version" mapstructure:"version"`
	WithBridge  bool            `json:"withBridge" mapstructure:"withBridge"`
	RpcHost     string          `json:"rpcHost" mapstructure:"rpcHost"`
	RpcPort     int             `json:"rpcPort" mapstructure:"rpcPort"`
	RpcUser     string          `json:"rpcUser" mapstructure:"rpcUser"`
	RpcPass     string          `json:"rpcPass" mapstructure:"rpcPass"`
}

// DefaultConfig returns a configuration populated with default values
func DefaultConfig() []BTCConfig {
	return []BTCConfig{{
		ChainID:     4,
		Chain:       Testnet4BtcChain,
		NetworkKind: Testnet,
		Name:        "bitcoin-testnet4",
		ID:          "bitcoin-testnet4",
		Gateway:     "",
		Finality:    10,
		LastBlock:   0,
		GasLimit:    1000000,
		BlockTime:   1000 * time.Millisecond,
		MaxRetry:    3,
		RetryDelay:  100 * time.Millisecond,
		TxTimeout:   3 * time.Second,
		RpcHost:     "http://127.0.0.1:48332",
		RpcUser:     "user",
		RpcPass:     "password",
		Tag:         "SCALAR",
		Version:     1,
		WithBridge:  false,
	}}
}
