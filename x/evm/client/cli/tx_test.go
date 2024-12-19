package cli

import (
	"fmt"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
)

var scalarHome string

func TestMain(m *testing.M) {
	scalarHome = os.Getenv("SCALAR_HOME_DIR")
	os.Exit(m.Run())
}
func TestEthToWei_IsInteger(t *testing.T) {
	amount, _ := sdk.NewDecFromStr("3.2")
	eth := sdk.DecCoin{
		Denom:  "eth",
		Amount: amount,
	}
	wei := eth
	wei.Amount = eth.Amount.MulInt64(params.Ether)

	assert.True(t, wei.Amount.IsInteger())
}

func TestGweiToWei_IsNotInteger(t *testing.T) {
	amount, _ := sdk.NewDecFromStr("3.0000000000002")
	gwei := sdk.DecCoin{
		Denom:  "gwei",
		Amount: amount,
	}
	wei := gwei
	wei.Amount = gwei.Amount.MulInt64(params.GWei)

	assert.False(t, wei.Amount.IsInteger())
}

func TestConfirmGatewayTxs(t *testing.T) {
	cmd := GetCmdCreateConfirmGatewayTxs()
	args := []string{
		"evm|11155111", "0x983bff649adbdc2766948e280f5c7c11d07081f02234dd254d06b3c02f21d5fd",
	}
	cmd.SetArgs(args)
	home := fmt.Sprintf("%s/scalar/node1/scalard", scalarHome)
	cmd.Flags().Set(flags.FlagFrom, "broadcaster")
	cmd.Flags().Set(flags.FlagKeyringBackend, "test")
	cmd.Flags().Set(flags.FlagHome, home)
	cmd.Flags().Set(flags.FlagGas, "300000")
	cmd.Flags().Set(flags.FlagChainID, "scalar-testnet-1")
	err := cmd.Execute()
	assert.NoError(t, err)
}
