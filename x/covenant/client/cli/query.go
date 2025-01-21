package cli

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/spf13/cobra"
)

const (
	FlagPubKey = "pubkey"
	FlagName   = "name"
	FlagStatus = "status"
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
		GetCmdFindGroup(),
	)

	return protocolQueryCmd
}
func GetCmdFindCustodian() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "custodians",
		Short: "Find custodians",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryServiceClient(clientCtx)
			request := types.CustodiansRequest{}
			readCustodianFlags(cmd, &request)
			res, err := queryClient.Custodians(cmd.Context(), &request)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	addCustodianFlagsToCmd(cmd)
	return cmd
}

func GetCmdFindGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "groups",
		Short: "Find custodian groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryServiceClient(clientCtx)
			request := types.GroupsRequest{}
			readGroupFlags(cmd, &request)
			res, err := queryClient.Groups(cmd.Context(), &request)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	addGroupFlagsToCmd(cmd)
	return cmd
}
func addCustodianFlagsToCmd(cmd *cobra.Command) {
	cmd.Flags().String(FlagPubKey, "", "Custodian pubkey")
	cmd.Flags().String(FlagName, "", "Name of the custodian")
	cmd.Flags().String(FlagStatus, "", "Status of the custodian")
}
func readCustodianFlags(cmd *cobra.Command, request *types.CustodiansRequest) {
	if request == nil {
		return
	}
	var err error
	pubkey, _ := cmd.Flags().GetString(FlagPubKey)
	request.Pubkey, err = hex.DecodeString(pubkey)
	if err != nil {
		log.Fatal("Failed to decode pubkey", err)
	}
	request.Name, err = cmd.Flags().GetString(FlagName)
	if err != nil {
		log.Fatal("Failed to get name", err)
	}
}
func addGroupFlagsToCmd(cmd *cobra.Command) {
	cmd.Flags().String(FlagPubKey, "", "Custodian pubkey")
	cmd.Flags().String(FlagStatus, "", "Status of the custodian")
	cmd.Flags().String(FlagName, "", "Name of the custodian")
}

func readGroupFlags(cmd *cobra.Command, request *types.GroupsRequest) {
	if request == nil {
		return
	}
	var err error
	pubkey, _ := cmd.Flags().GetString(FlagPubKey)
	request.GroupPubkey, err = hex.DecodeString(pubkey)
	if err != nil {
		log.Fatal("Failed to decode pubkey", err)
	}
	name, _ := cmd.Flags().GetString(FlagName)
	request.Name = name
}
