package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/scalarorg/scalar-core/x/covenant/types"
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

	// protocolTxCmd.AddCommand(
	// 	GetCmdAddProtocol(),
	// )

	return protocolTxCmd
}
