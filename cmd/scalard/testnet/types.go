package testnet

import (
	"encoding/json"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/scalarorg/scalar-core/vald/config"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	ValidatorKeyName   = "priv_validator"
	BroadcasterKeyName = "broadcaster"
	ScalarKeyName      = "scalar"
	GovKeyName         = "govenance"
	FaucetKeyName      = "faucet"
)

type GenesisState map[string]json.RawMessage

type ScalarRelayer struct {
	PubKey  cryptotypes.PubKey
	Balance banktypes.Balance
}
type Protocol struct {
	BitcoinPubKey  []byte
	PubKey         cryptotypes.PubKey
	Balance        banktypes.Balance
	Tag            string
	LiquidityModel string
}
type ProtocolConfig struct {
	ScalarMnemonic string `json:"scalar_mnemonic" mapstructure:"scalar_mnemonic"`
	BitcoinPrivKey string `json:"bitcoin_privkey" mapstructure:"bitcoin_privkey"`
	Tag            string `json:"tag" mapstructure:"tag"`
	LiquidityModel string `json:"liquidity_model" mapstructure:"liquidity_model"`
}
type DeployInfo struct {
	ID           string `json:"id" mapstructure:"id"`
	ChainId      uint64 `json:"chain_id" mapstructure:"chain_id"`
	Name         string `json:"name" mapstructure:"name"`
	Gateway      string `json:"gateway" mapstructure:"gateway"`
	TokenAddress string `json:"token_address" mapstructure:"token_address"`
	TxHash       string `json:"tx_hash" mapstructure:"tx_hash"`
}
type Token struct {
	ID             string       `json:"id" mapstructure:"id"`
	ChainID        int64        `json:"chain_id" mapstructure:"chain_id"`
	Asset          string       `json:"asset" mapstructure:"asset"`
	Symbol         string       `json:"symbol" mapstructure:"symbol"`
	Name           string       `json:"name" mapstructure:"name"`
	Capacity       int64        `json:"capacity" mapstructure:"capacity"`
	Decimals       uint8        `json:"decimals" mapstructure:"decimals"`
	DailyMintLimit string       `json:"daily_mint_limit" mapstructure:"daily_mint_limit"`
	LiquidityModel string       `json:"liquidity_model" mapstructure:"liquidity_model"`
	TokenAddress   string       `json:"token_address" mapstructure:"token_address"`
	Deployments    []DeployInfo `json:"deployments" mapstructure:"deployments"`
}
type InternalToken struct {
	ChainID int64  `json:"chain_id" mapstructure:"chain_id"`
	Address string `json:"address" mapstructure:"address"`
}
type TokenPool struct {
	ChainId string          `json:"chain_d" mapstructure:"chain_id"`
	Asset   string          `json:"asset" mapstructure:"asset"`
	Chains  []InternalToken `json:"chain" mapstructure:"chain"`
}
type ValidatorInfo struct {
	Host        string
	Moniker     string
	NodeID      string
	NodeDir     string
	SeedAddress string
	RPCPort     int

	ValPubKey  cryptotypes.PubKey
	ValBalance banktypes.Balance

	ValNodePubKey      cryptotypes.PubKey
	ValNodeBalance     banktypes.Balance
	Broadcaster        cryptotypes.PubKey
	BroadcasterBalance banktypes.Balance

	GovPubKey  cryptotypes.PubKey
	GovBalance banktypes.Balance

	FaucetPubKey  cryptotypes.PubKey
	FaucetBalance banktypes.Balance

	GenesisValidator tmtypes.GenesisValidator
	GenFile          string
	AdditionalKeys   config.AdditionalKeys
}
