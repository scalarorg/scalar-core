package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/scalarorg/scalar-core/x/protocol/types"
	"github.com/spf13/cobra"
)

const (
	FlagPubKey  = "pubkey"
	FlagAddress = "address"
	FlagName    = "name"
	FlagStatus  = "status"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	protocolQueryCmd := &cobra.Command{
		Use:                        "protocol",
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		//RunE:                       client.ValidateCmd,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			request := types.ProtocolsRequest{}
			readProtocolFlags(cmd, &request)
			res, err := queryClient.Protocols(cmd.Context(), &request)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(protocolQueryCmd)
	addProtocolFlagsToCmd(protocolQueryCmd)
	return protocolQueryCmd
}

func addProtocolFlagsToCmd(cmd *cobra.Command) {
	cmd.Flags().String(FlagPubKey, "", "Protocol pubkey")
	cmd.Flags().String(FlagAddress, "", "Protocol address to find")
	cmd.Flags().String(FlagName, "", "Name of the protocol")
	cmd.Flags().StringP(FlagStatus, "", "Activated", "Status of the protocol (Unspecified|Activated|Deactived)")
}
func readProtocolFlags(cmd *cobra.Command, request *types.ProtocolsRequest) {
	if request == nil {
		return
	}
	pubkey, _ := cmd.Flags().GetString(FlagPubKey)
	request.Pubkey = pubkey
	address, _ := cmd.Flags().GetString(FlagAddress)
	request.Address = address
	name, _ := cmd.Flags().GetString(FlagName)
	request.Name = name
}
