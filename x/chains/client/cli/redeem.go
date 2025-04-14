package cli

import (
	"log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/spf13/cobra"
)

const (
	flagChain = "chain"
)

func getCmdRedeemSession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redeem-session [chain]",
		Short: "Get redeem session by custodian uid",
		Args:  cobra.ExactArgs(1),
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
	cmd.Flags().String(flagChain, "", "Evm Chain name")
}
func readRedeemSessionFlags(cmd *cobra.Command, request *types.RedeemSessionRequest) {
	if request == nil {
		return
	}
	var err error
	chain, err := cmd.Flags().GetString(flagChain)
	request.Chain = chain
	if err != nil {
		log.Fatal("Failed to decode pubkey", err)
	}
}
