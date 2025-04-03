package cli

import (
	"log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	chainsExported "github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/spf13/cobra"
)

func GetCmdRedeemSession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redeem-session",
		Short: "Get redeem session by custodian uid",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryServiceClient(clientCtx)
			request := types.RedeemSessionRequest{}
			readRedeemSessionFlags(cmd, &request)
			res, err := queryClient.RedeemSession(cmd.Context(), &request)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	addRedeemSessionFlagsToCmd(cmd)
	return cmd
}
func addRedeemSessionFlagsToCmd(cmd *cobra.Command) {
	cmd.Flags().String(FlagUID, "", "Custodian UID")
	//cmd.Flags().String("taproot-pubkey", "", "Taproot pubkey in hex")
}
func readRedeemSessionFlags(cmd *cobra.Command, request *types.RedeemSessionRequest) {
	if request == nil {
		return
	}
	var err error
	uid, _ := cmd.Flags().GetString(FlagUID)
	request.UID, err = chainsExported.HashFromHex(uid)
	if err != nil {
		log.Fatal("Failed to decode pubkey", err)
	}
}
func GetCmdUTXOSnapshot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "utxo-snapshot",
		Short: "Get the UTXO snapshot",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryServiceClient(clientCtx)
			request := types.UTXOSnapshotRequest{}
			readUTXOSnapshotFlags(cmd, &request)
			res, err := queryClient.UTXOSnapshot(cmd.Context(), &request)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	addUTXOSnapshotFlagsToCmd(cmd)
	return cmd
}
func addUTXOSnapshotFlagsToCmd(cmd *cobra.Command) {
	cmd.Flags().String(FlagUID, "", "Custodian UID")
}
func readUTXOSnapshotFlags(cmd *cobra.Command, request *types.UTXOSnapshotRequest) {
	if request == nil {
		return
	}
	var err error
	uid, _ := cmd.Flags().GetString(FlagUID)
	request.UID, err = chainsExported.HashFromHex(uid)
	if err != nil {
		log.Fatal("Failed to decode pubkey", err)
	}
}
