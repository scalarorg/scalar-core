package cli

import (
	"errors"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	legacyTx "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/spf13/cobra"

	"github.com/scalarorg/scalar-core/x/snapshot/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	snapshotTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	snapshotTxCmd.AddCommand(
		GetCmdRegisterProxy(),
		GetCmdDeregisterProxy(),
		GetCmdSendTokens(),
	)

	return snapshotTxCmd
}

// GetCmdRegisterProxy returns the command to register a proxy
func GetCmdRegisterProxy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-proxy [proxy address]",
		Short: "Register a proxy account for a specific validator principal to broadcast transactions in its stead",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return sdkerrors.Wrap(types.ErrSnapshot, "proxy invalid")
			}

			msg := types.NewRegisterProxyRequest(sdk.ValAddress(clientCtx.FromAddress), addr)
			return legacyTx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdDeregisterProxy returns the command to register a proxy
func GetCmdDeregisterProxy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deactivate-proxy",
		Short: "Deactivate the proxy account of the sender",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewDeactivateProxyRequest(sdk.ValAddress(clientCtx.FromAddress))
			return legacyTx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdSendTokens returns the command to send stake to a number of addresses
func GetCmdSendTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-tokens [amount] [address 1] ... [address n]",
		Short: "Sends the specified amount of tokens to the designated addresses",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			decCoins, err := sdk.ParseDecCoins(args[0])
			if err != nil {
				return err
			}

			if decCoins.Len() != 1 {
				return errors.New("only a single amount is permitted")
			}

			coins, decimals := decCoins.TruncateDecimal()
			if !decimals.IsZero() {
				return errors.New("amount must be an integer value")
			}

			inputs := make([]banktypes.Input, 0)
			outputs := make([]banktypes.Output, 0)

			for _, addr := range args[1:] {

				to, err := sdk.AccAddressFromBech32(addr)
				if err != nil {
					return err
				}

				inputs = append(inputs, banktypes.NewInput(clientCtx.FromAddress, coins))
				outputs = append(outputs, banktypes.NewOutput(to, coins))

			}

			msg := banktypes.NewMsgMultiSend(inputs, outputs)
			return legacyTx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
