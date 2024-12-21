package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/scalarorg/scalar-core/x/protocol/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(route string) *cobra.Command {
	protocolQueryCmd := &cobra.Command{
		Use:                        "protocol",
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	protocolQueryCmd.AddCommand(
		GetCmdFindProtocol(),
	)

	return protocolQueryCmd
}
func GetCmdFindProtocol() *cobra.Command {
	return &cobra.Command{
		Use:   "find",
		Short: "Find a protocol",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
