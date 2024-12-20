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
		Use:                        "covenant",
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	protocolQueryCmd.AddCommand(
		GetCmdFindCustodian(),
		GetCmdFindCustodianGroup(),
	)

	return protocolQueryCmd
}
func GetCmdFindCustodian() *cobra.Command {
	return &cobra.Command{
		Use:   "custodians",
		Short: "Find custodians",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func GetCmdFindCustodianGroup() *cobra.Command {
	return &cobra.Command{
		Use:   "custodianGroups",
		Short: "Find custodian groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
