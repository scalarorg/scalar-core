package cli

import (
	"fmt"
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/stretchr/testify/assert"
)

var scalarHome string

func TestMain(m *testing.M) {
	scalarHome = os.Getenv("SCALAR_HOME_DIR")
	os.Exit(m.Run())
}
func TestConfirmSourceTxs(t *testing.T) {
	cmd := getCmdCreateConfirmSourceTxs()
	args := []string{
		"bitcoin|4", "18fa2be86b54d9ff7e35aba97d57483f05500cd9301547607f67ea5b47fa1c87",
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
