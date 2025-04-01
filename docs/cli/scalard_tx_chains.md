## scalard tx chains

chains transactions subcommands

```
scalard tx chains [flags]
```

### Options

```
  -h, --help   help for chains
```

### Options inherited from parent commands

```
      --chain-id string     The network chain ID (default "scalar")
      --home string         directory for config and data (default "$HOME/.scalar")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --output string       Output format (text|json) (default "text")
      --trace               print out full stack trace on errors
```

### SEE ALSO

* [scalard tx](scalard_tx.md)	 - Transactions subcommands
* [scalard tx chains add-chain](scalard_tx_chains_add-chain.md)	 - Add a new EVM chain
* [scalard tx chains confirm-erc20-deposit](scalard_tx_chains_confirm-erc20-deposit.md)	 - Confirm ERC20 deposits in an EVM chain transaction to a burner address
* [scalard tx chains confirm-erc20-token](scalard_tx_chains_confirm-erc20-token.md)	 - Confirm an ERC20 token deployment in an EVM chain transaction for a given asset of some origin chain and gateway address
* [scalard tx chains confirm-source-txs](scalard_tx_chains_confirm-source-txs.md)	 - Confirm source transactions in a chain
* [scalard tx chains confirm-transfer-operatorship](scalard_tx_chains_confirm-transfer-operatorship.md)	 - Confirm a transfer operatorship in an EVM chain transaction
* [scalard tx chains create-burn-tokens](scalard_tx_chains_create-burn-tokens.md)	 - Create burn commands for all confirmed token deposits in an EVM chain
* [scalard tx chains create-deploy-token](scalard_tx_chains_create-deploy-token.md)	 - Create a deploy token command with the ScalarGateway contract
* [scalard tx chains create-pending-transfers](scalard_tx_chains_create-pending-transfers.md)	 - Create commands for handling all pending transfers to an EVM chain
* [scalard tx chains link](scalard_tx_chains_link.md)	 - Link a cross chain address to an EVM chain address created by Scalar
* [scalard tx chains set-gateway](scalard_tx_chains_set-gateway.md)	 - Set the gateway address for the given evm chain
* [scalard tx chains sign-btc-commands](scalard_tx_chains_sign-btc-commands.md)	 - Sign pending commands for a BTC chain contract
* [scalard tx chains sign-commands](scalard_tx_chains_sign-commands.md)	 - Sign pending commands for an EVM chain contract
* [scalard tx chains transfer-operatorship](scalard_tx_chains_transfer-operatorship.md)	 - Create transfer operatorship command for an EVM chain contract

