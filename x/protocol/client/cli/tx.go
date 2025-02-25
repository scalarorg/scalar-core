package cli

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	chainTypes "github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/scalarorg/scalar-core/x/protocol/exported"
	"github.com/scalarorg/scalar-core/x/protocol/types"
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

	protocolTxCmd.AddCommand(
		GetCmdAddProtocol(),
	)

	return protocolTxCmd
}

type CreateProtocolArgs struct {
	Attribute         *exported.ProtocolAttributes `json:"attribute"`
	Avatar            string                       `json:"avatar"`
	BitcoinPubkey     string                       `json:"bitcoin_pubkey"`
	CustodianGroupUid string                       `json:"custodian_group_uid"`
	Name              string                       `json:"name"`
	Tag               string                       `json:"tag"`
	Asset             struct {
		ChainName string `json:"chain_name"`
		AssetName string `json:"asset_name"`
	} `json:"asset"`
}

func GetCmdAddProtocol() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [protocol-json]",
		Short: "Add a protocol",
		Long: `Add a new protocol with the specified parameters in JSON format.
Example:
$ scalard tx protocol add '{"attribute":{"model":0},"avatar":"base64_encoded_avatar","bitcoin_pubkey":"hex_encoded_pubkey","custodian_group_uid":"uuid","name":"protocol_name","tag":"tag_value","asset":{"chain_name":"bitcoin","asset_name":"btc"}}'`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Get account information even in offline mode
			address := clientCtx.GetFromAddress()
			var scalarPubkeyBytes []byte

			account, err := clientCtx.AccountRetriever.GetAccount(clientCtx, address)
			if err != nil {
				return fmt.Errorf("failed to get account: %w", err)
			}
			scalarPubkey := account.GetPubKey()
			scalarPubkeyBytes = scalarPubkey.Bytes()

			fmt.Printf("scalarPubkeyBytes: %+x\n", scalarPubkeyBytes)
			fmt.Printf("address: %+v\n", address)

			var createProtocolArgs CreateProtocolArgs
			if err := json.Unmarshal([]byte(args[0]), &createProtocolArgs); err != nil {
				return fmt.Errorf("failed to parse protocol JSON: %w", err)
			}

			fmt.Printf("createProtocolArgs: %+v\n", createProtocolArgs)

			bitcoinPubkey, err := hex.DecodeString(createProtocolArgs.BitcoinPubkey)
			if err != nil {
				return fmt.Errorf("invalid bitcoin pubkey hex: %w", err)
			}

			avatar, err := base64.StdEncoding.DecodeString(createProtocolArgs.Avatar)
			if err != nil {
				return fmt.Errorf("invalid avatar base64: %w", err)
			}

			msg := types.NewCreateProtocolRequest(address, createProtocolArgs.Name, bitcoinPubkey, scalarPubkeyBytes, createProtocolArgs.Tag, createProtocolArgs.Attribute, createProtocolArgs.CustodianGroupUid, avatar, &chainTypes.Asset{
				Chain: nexus.ChainName(createProtocolArgs.Asset.ChainName),
				Name:  createProtocolArgs.Asset.AssetName,
			})
			if err := msg.ValidateBasic(); err != nil {
				return fmt.Errorf("failed to validate message: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd

}
