package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/scalarorg/scalar-core/x/protocol/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	protocolTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	protocolTxCmd.AddCommand(
		GetCmdAddProtocol(),
	)

	return protocolTxCmd
}

func GetCmdAddProtocol() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add a protocol",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
