package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/spf13/cobra"

	"github.com/scalarorg/scalar-core/x/btc/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	btcQueryCmd := &cobra.Command{
		Use:                        "btc",
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	btcQueryCmd.AddCommand(
		getCmdQueryBatchedCommands(),
	)

	return btcQueryCmd

}

// getCmdQueryBatchedCommands returns the query to get the batched commands
func getCmdQueryBatchedCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batched-commands [chain] [batchedCommandsID]",
		Short: "Get the signed batched commands that can be wrapped in an EVM transaction to be executed in Scalar Gateway",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			chain := args[0]
			idHex := args[1]

			queryClient := types.NewQueryServiceClient(clientCtx)

			res, err := queryClient.BatchedCommands(cmd.Context(),
				&types.BatchedCommandsRequest{
					Chain: utils.NormalizeString(chain),
					Id:    utils.NormalizeString(idHex),
				})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
