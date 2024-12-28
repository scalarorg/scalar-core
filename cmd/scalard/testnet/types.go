package testnet

import (
	"encoding/json"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
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

type ScalarProtocol struct {
	ScalarPubKey  cryptotypes.PubKey
	ScalarBalance banktypes.Balance
}
type Token struct {
	ID              string `json:"id" mapstructure:"id"`
	ChainID         int64  `json:"chain_id" mapstructure:"chain_id"`
	Asset           string `json:"asset" mapstructure:"asset"`
	Symbol          string `json:"symbol" mapstructure:"symbol"`
	Name            string `json:"name" mapstructure:"name"`
	Capacity        int64  `json:"capacity" mapstructure:"capacity"`
	Decimals        uint8  `json:"decimals" mapstructure:"decimals"`
	TokenAddress    string `json:"token_address" mapstructure:"token_address"`
	ProtocolAddress string `json:"protocol_address" mapstructure:"protocol_address"`
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

	ValNodePubKey  cryptotypes.PubKey
	ValNodeBalance banktypes.Balance
	//Balance of broadcaster
	Broadcaster        cryptotypes.PubKey
	BroadcasterBalance banktypes.Balance

	GovPubKey  cryptotypes.PubKey
	GovBalance banktypes.Balance

	FaucetPubKey  cryptotypes.PubKey
	FaucetBalance banktypes.Balance

	GenesisValidator tmtypes.GenesisValidator
	BtcPubkey        string
	GenFile          string
}
