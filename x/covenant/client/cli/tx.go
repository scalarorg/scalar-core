package cli

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/scalarorg/scalar-core/utils"
	chainexported "github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdUpdateUtxoForRedeemSession(),
	)

	return cmd
}

type UTXO struct {
	TxID         string `json:"txid"`
	Vout         uint32 `json:"vout"`
	ScriptPubKey string `json:"script_pubkey"`
	Value        uint64 `json:"value"`
}

func GetCmdUpdateUtxoForRedeemSession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-utxo-for-redeem-session <symbol> <list_of_utxos> <block_height>...",
		Short: "Update UTXO for redeem session",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainName := utils.NormalizeString(args[0])

			var rawUtxos []UTXO
			err = json.Unmarshal([]byte(args[1]), &rawUtxos)
			if err != nil {
				return err
			}

			var utxos []*types.UTXO
			for _, utxo := range rawUtxos {
				if utxo.TxID == "" {
					return fmt.Errorf("txid is required")
				}

				if utxo.ScriptPubKey == "" {
					return fmt.Errorf("script_pubkey is required")
				}

				txId, err := chainexported.HashFromHex(utxo.TxID)
				if err != nil {
					return err
				}

				scriptPubKey, err := hex.DecodeString(utxo.ScriptPubKey)
				if err != nil {
					return err
				}

				utxos = append(utxos, &types.UTXO{
					TxID:         txId,
					Vout:         utxo.Vout,
					ScriptPubkey: scriptPubKey,
					AmountInSats: utxo.Value,
				})
			}

			blockHeight, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse block height: %v", err)
			}

			msg := &types.UpdateUtxoForRedeemSessionRequest{
				Symbol:      chainName,
				ListOfUtxos: utxos,
				BlockHeight: blockHeight,
				Sender:      cliCtx.GetFromAddress(),
			}
			if err := msg.ValidateBasic(); err != nil {
				return fmt.Errorf("failed to validate message: %v", err)
			}
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
