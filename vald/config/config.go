package config

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils"
	tss "github.com/scalarorg/scalar-core/x/tss/types"
)

type BTCConfig struct {
	ID         string `json:"id" mapstructure:"id"`
	Chain      string `json:"chain" mapstructure:"chain"`
	Tag        string `json:"tag" mapstructure:"tag"`
	Version    uint64 `json:"version" mapstructure:"version"`
	WithBridge bool   `json:"with_bridge" mapstructure:"with_bridge"`
	RPCHost    string `json:"rpc_host" mapstructure:"rpc_host"`
	RPCPort    uint64 `json:"rpc_port" mapstructure:"rpc_port"`
	RPCUser    string `json:"rpc_user" mapstructure:"rpc_user"`
	RPCPass    string `json:"rpc_pass" mapstructure:"rpc_pass"`

	DisableTLS           bool `json:"disable_tls" mapstructure:"disable_tls"`
	DisableConnectOnNew  bool `json:"disable_connect_on_new" mapstructure:"disable_connect_on_new"`
	DisableAutoReconnect bool `json:"disable_auto_reconnect" mapstructure:"disable_auto_reconnect"`
	HttpPostMode         bool `json:"http_post_mode" mapstructure:"http_post_mode"`
}

func (c *BTCConfig) ValidateBasic() error {
	_, err := utils.ChainInfoBytesFromID(c.ID)
	if err != nil {
		return fmt.Errorf("invalid chain ID %s", c.ID)
	}
	return nil
}

type EVMConfig struct {
	ID               string `json:"id" mapstructure:"id"`
	RPCAddr          string `json:"rpc_addr" mapstructure:"rpc_addr"`
	WithBridge       bool   `json:"with_bridge" mapstructure:"with_bridge"`
	FinalityOverride string `json:"finality_override" mapstructure:"finality_override"`
}

func (c *EVMConfig) ValidateBasic() error {
	_, err := utils.ChainInfoBytesFromID(c.ID)
	if err != nil {
		return fmt.Errorf("invalid chain ID %s", c.ID)
	}
	return nil
}

type AdditionalKeys struct {
	BtcPrivKey string `json:"btc_priv_key" mapstructure:"btc_priv_key"`
}

// ValdConfig contains all necessary vald configurations
type ValdConfig struct {
	tss.TssConfig                `mapstructure:"tss"`
	BroadcastConfig              `mapstructure:"broadcast"`
	BatchSizeLimit               int           `mapstructure:"max_batch_size"`
	BatchThreshold               int           `mapstructure:"batch_threshold"`
	MaxBlocksBehindLatest        int64         `mapstructure:"max_blocks_behind_latest"` // The max amount of blocks behind the latest until which the cached height is considered valid.
	EventNotificationsMaxRetries int           `mapstructure:"event_notifications_max_retries"`
	EventNotificationsBackOff    time.Duration `mapstructure:"event_notifications_back_off"`
	MaxLatestBlockAge            time.Duration `mapstructure:"max_latest_block_age"`  // If a block is older than this, vald does not consider it to be the latest block. This is supposed to be sufficiently larger than the block production time.
	NoNewBlockPanicTimeout       time.Duration `mapstructure:"no_new_blocks_timeout"` // At times vald stalls completely. Until the bug is found it is better to panic and allow users to restart the process instead of doing nothing. Once at least one block has been seen vald will panic if it does not see another before the timout expires.

	BTCMgrConfig   []BTCConfig `mapstructure:"scalar_bridge_btc"`
	EVMMgrConfig   []EVMConfig `mapstructure:"scalar_bridge_evm"`
	AdditionalKeys `mapstructure:"additional_keys"`
}

// DefaultValdConfig returns a configurations populated with default values
func DefaultValdConfig() ValdConfig {
	return ValdConfig{
		TssConfig:                    tss.DefaultConfig(),
		BroadcastConfig:              DefaultBroadcastConfig(),
		BatchSizeLimit:               250,
		BatchThreshold:               3,
		MaxBlocksBehindLatest:        10, // Max voting/sign/heartbeats periods are under 10 blocks
		MaxLatestBlockAge:            15 * time.Second,
		EventNotificationsMaxRetries: 3,
		EventNotificationsBackOff:    1 * time.Second,
		NoNewBlockPanicTimeout:       2 * time.Minute,
	}
}

// BroadcastConfig is the configuration for transaction broadcasting
type BroadcastConfig struct {
	MaxRetries          int            `mapstructure:"max_retries"`
	MinSleepBeforeRetry time.Duration  `mapstructure:"min_sleep_before_retry"`
	MaxTimeout          time.Duration  `mapstructure:"max_timeout"`
	FeeGranter          sdk.AccAddress `mapstructure:"fee_granter"`
}

// DefaultBroadcastConfig returns a configurations populated with default values
func DefaultBroadcastConfig() BroadcastConfig {
	return BroadcastConfig{
		MaxRetries:          3,
		MinSleepBeforeRetry: 5 * time.Second,
		MaxTimeout:          15 * time.Second,
	}
}
