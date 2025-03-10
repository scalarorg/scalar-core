package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/spf13/cobra"

	"github.com/scalarorg/scalar-core/x/chains/types"

	covenant "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

const (
	flagAddress = "address"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	chainsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		TraverseChildren:           true,
		RunE:                       client.ValidateCmd,
	}

	chainsTxCmd.AddCommand(
		GetCmdCreateConfirmSourceTxs(),
		GetCmdSetGateway(),
		GetCmdLink(),
		GetCmdConfirmERC20TokenDeployment(),
		GetCmdConfirmERC20Deposit(),
		GetCmdConfirmTransferOperatorship(),
		GetCmdCreatePendingTransfers(),
		GetCmdCreateDeployToken(),
		GetCmdCreateBurnTokens(),
		GetCmdCreateTransferOperatorship(),
		GetCmdSignCommands(),
		GetCmdSignBtcCommands(),
		GetCmdSignPsbtCommand(),
		GetCmdAddChain(),
	)
	return chainsTxCmd
}

func GetCmdCreateConfirmSourceTxs() *cobra.Command {
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
				txHash, err := types.HashFromHex(arg)
				if err != nil {
					return fmt.Errorf("failed to parse txID %s: %v", arg, err)
				}
				txIDs = append(txIDs, txHash)
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

// GetCmdSetGateway sets the gateway address for the given evm chain
func GetCmdSetGateway() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-gateway [chain] [address]",
		Short: "Set the gateway address for the given evm chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chain := args[0]
			addressHex := args[1]
			if !common.IsHexAddress(addressHex) {
				return fmt.Errorf("invalid address %s", addressHex)
			}

			msg := types.NewSetGatewayRequest(cliCtx.GetFromAddress(), chain, types.Address(common.HexToAddress(addressHex)))

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdLink links a cross chain address to an EVM chain address created by Scalar
func GetCmdLink() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link [chain] [recipient chain] [recipient address] [asset name]",
		Short: "Link a cross chain address to an EVM chain address created by Scalar",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewLinkRequest(cliCtx.GetFromAddress(), args[0], args[1], args[2], args[3])

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdConfirmERC20TokenDeployment returns the cli command to confirm a ERC20 token deployment
func GetCmdConfirmERC20TokenDeployment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-erc20-token [chain] [origin chain] [origin asset] [txID]",
		Short: "Confirm an ERC20 token deployment in an EVM chain transaction for a given asset of some origin chain and gateway address",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chain := args[0]
			originChain := args[1]
			originAsset := args[2]
			asset := types.NewAsset(originChain, originAsset)
			txID := common.HexToHash(args[3])
			msg := types.NewConfirmTokenRequest(cliCtx.GetFromAddress(), chain, asset, txID)

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdConfirmERC20Deposit returns the cli command to confirm an ERC20 deposit
func GetCmdConfirmERC20Deposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-erc20-deposit [chain] [txID] [burnerAddr]",
		Short: "Confirm ERC20 deposits in an EVM chain transaction to a burner address",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chain := args[0]
			txID := common.HexToHash(args[1])
			burnerAddr := common.HexToAddress(args[2])

			msg := types.NewConfirmDepositRequest(cliCtx.GetFromAddress(), chain, txID, burnerAddr)

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdConfirmTransferOperatorship returns the cli command to confirm a transfer operatorship for the gateway contract
func GetCmdConfirmTransferOperatorship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-transfer-operatorship [chain] [txID]",
		Short: "Confirm a transfer operatorship in an EVM chain transaction",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chain := args[0]
			txID := common.HexToHash(args[1])
			msg := types.NewConfirmTransferKeyRequest(cliCtx.GetFromAddress(), chain, txID)

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCreatePendingTransfers returns the cli command to create commands for handling all pending token transfers to an EVM chain
func GetCmdCreatePendingTransfers() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pending-transfers [chain]",
		Short: "Create commands for handling all pending transfers to an EVM chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewCreatePendingTransfersRequest(cliCtx.GetFromAddress(), args[0])

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCreateDeployToken returns the cli command to create deploy-token command for an EVM chain
func GetCmdCreateDeployToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-deploy-token [evm chain] [origin chain] [origin asset] [token name] [symbol] [decimals] [capacity] [mintLimit]",
		Short: "Create a deploy token command with the ScalarGateway contract",
		Args:  cobra.ExactArgs(8),
	}
	address := cmd.Flags().String(flagAddress, types.ZeroAddress.Hex(), "existing ERC20 token's address")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		cliCtx, err := client.GetClientTxContext(cmd)
		if err != nil {
			return err
		}

		chain := args[0]
		tokenName := args[1]
		symbol := args[2]
		if !common.IsHexAddress(*address) {
			return fmt.Errorf("could not parse address")
		}

		msg := types.NewCreateDeployTokenRequest(cliCtx.GetFromAddress(), chain, symbol, tokenName, types.Address(common.HexToAddress(*address)))

		return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCreateBurnTokens returns the cli command to create burn commands for all confirmed token deposits in an EVM chain
func GetCmdCreateBurnTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-burn-tokens [chain]",
		Short: "Create burn commands for all confirmed token deposits in an EVM chain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewCreateBurnTokensRequest(cliCtx.GetFromAddress(), args[0])

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCreateTransferOperatorship returns the cli command to create transfer-operatorship command for an EVM chain contract
func GetCmdCreateTransferOperatorship() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-operatorship [chain] [keyID]",
		Short: "Create transfer operatorship command for an EVM chain contract",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewCreateTransferOperatorshipRequest(cliCtx.GetFromAddress(), args[0], args[1])

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdSignCommands returns the cli command to sign pending commands for an EVM chain contract
func GetCmdSignCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-commands [chain]",
		Short: "Sign pending commands for an EVM chain contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewSignCommandsRequest(cliCtx.GetFromAddress(), args[0])

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdAddChain returns the cli command to add a new evm chain command
func GetCmdAddChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-chain [name] [chain config]",
		Short: "Add a new EVM chain",
		Long:  "Add a new EVM chain. The chain config parameter should be the path to a json file containing the key requirements and the evm module parameters",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			name := args[0]
			jsonFile := args[1]

			byteValue, err := ioutil.ReadFile(jsonFile)
			if err != nil {
				return err
			}
			var chainConf struct {
				Params types.Params `json:"params"`
			}
			err = json.Unmarshal([]byte(byteValue), &chainConf)
			if err != nil {
				return err
			}

			msg := types.NewAddChainRequest(cliCtx.GetFromAddress(), name, chainConf.Params)

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdSignBtcCommands returns the cli command to sign pending commands for a BTC transaction
func GetCmdSignBtcCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-btc-commands [chain]",
		Short: "Sign pending commands for a BTC chain contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewSignBtcCommandsRequest(cliCtx.GetFromAddress(), args[0])

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdSignPsbtCommand returns the cli command to sign psbt commands for a BTC transaction
func GetCmdSignPsbtCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-psbt [chain] [psbt]",
		Short: "Sign psbt commands for a BTC transaction",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			psbt, err := covenant.PsbtFromHex(args[1])
			if err != nil {
				return err
			}

			msg := types.NewSignPsbtCommandRequest(cliCtx.GetFromAddress(), args[0], psbt)

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
