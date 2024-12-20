package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/spf13/cobra"

	"github.com/scalarorg/scalar-core/x/chains/types"

	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	btcTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		TraverseChildren:           true,
		RunE:                       client.ValidateCmd,
	}

	btcTxCmd.AddCommand(getCmdCreateConfirmSourceTxs())
	btcTxCmd.AddCommand(getCmdCreateConfirmDestTxs())

	return btcTxCmd
}

func getCmdCreateConfirmSourceTxs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-source-txs <chain> <txID>...",
		Short: "Confirm source transactions in a chain",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainName := utils.NormalizeString(args[0])

			var txIDs []types.Hash
			for _, arg := range args[1:] {
				txHash, err := types.HashFromHexStr(arg)
				if err != nil {
					return fmt.Errorf("failed to parse txID %s: %v", arg, err)
				}
				txIDs = append(txIDs, *txHash)
			}

			msg := types.NewConfirmSourceTxsRequest(cliCtx.GetFromAddress(), nexus.ChainName(chainName), txIDs)
			if err := msg.ValidateBasic(); err != nil {
				return fmt.Errorf("failed to validate message: %v", err)
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd

}

func getCmdCreateConfirmDestTxs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-dest-txs <chain> <txID>...",
		Short: "Confirm destination transactions in a chain",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainName := utils.NormalizeString(args[0])

			var txIDs []types.Hash
			for _, arg := range args[1:] {
				txHash, err := types.HashFromHexStr(arg)
				if err != nil {
					return fmt.Errorf("failed to parse txID %s: %v", arg, err)
				}
				txIDs = append(txIDs, *txHash)
			}

			msg := types.NewConfirmDestTxsRequest(cliCtx.GetFromAddress(), nexus.ChainName(chainName), txIDs)
			if err := msg.ValidateBasic(); err != nil {
				clog.Red("getCmdCreateConfirmDestTxs", "err", err)
				return fmt.Errorf("failed to validate message: %v", err)
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
