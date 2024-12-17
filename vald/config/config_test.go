package config_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/scalarorg/scalar-core/vald/config"
	"github.com/spf13/viper"
)

func TestReadConfig(t *testing.T) {

	content, err := os.ReadFile("testdata/config.toml")
	if err != nil {
		t.Fatal(err)
	}

	viper.SetConfigType("toml")
	if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
		t.Fatal(err)
	}

	viper.AddConfigPath("testdata/config.toml")

	valdConf := config.DefaultValdConfig()
	viper.RegisterAlias("broadcast.max_timeout", "rpc.timeout_broadcast_tx_commit")
	if err := viper.Unmarshal(&valdConf, config.AddDecodeHooks); err != nil {
		panic(err)
	}

	t.Logf("%+v", valdConf)
}
