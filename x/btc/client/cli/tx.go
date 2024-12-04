package cli

import (
	"fmt"

	"github.com/axelarnetwork/axelar-core/utils"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/scalarorg/scalar-core/x/btc/types"

	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
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

	btcTxCmd.AddCommand(getCmdCreateConfirmGatewayTxs())

	return btcTxCmd
}

func getCmdCreateConfirmGatewayTxs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-gateway-txs <chain> <txID>...",
		Short: "Confirm gateway transactions in an EVM chain",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chain := nexus.ChainName(utils.NormalizeString(args[0]))
			var txIDs []types.TxHash
			for _, arg := range args[1:] {
				txHash, err := types.TxHashFromHexStr(arg)
				if err != nil {
					return fmt.Errorf("failed to parse txID %s: %v", arg, err)
				}
				txIDs = append(txIDs, *txHash)
			}

			msg := types.NewConfirmGatewayTxsRequest(cliCtx.GetFromAddress(), chain, txIDs)
			if err := msg.ValidateBasic(); err != nil {
				return fmt.Errorf("failed to validate message: %v", err)
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd

}
