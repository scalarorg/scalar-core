## scalard tx evm

evm transactions subcommands

```
scalard tx evm [flags]
```

### Options

```
  -h, --help   help for evm
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

- [scalard tx](scalard_tx.md) - Transactions subcommands
- [scalard tx evm add-chain](scalard_tx_evm_add-chain.md) - Add a new EVM chain
- [scalard tx evm confirm-erc20-deposit](scalard_tx_evm_confirm-erc20-deposit.md) - Confirm ERC20 deposits in an EVM chain transaction to a burner address
- [scalard tx evm confirm-erc20-token](scalard_tx_evm_confirm-erc20-token.md) - Confirm an ERC20 token deployment in an EVM chain transaction for a given asset of some origin chain and gateway address
- [scalard tx evm confirm-gateway-txs](scalard_tx_evm_confirm-gateway-txs.md) - Confirm gateway transactions in an EVM chain
- [scalard tx evm confirm-transfer-operatorship](scalard_tx_evm_confirm-transfer-operatorship.md) - Confirm a transfer operatorship in an EVM chain transaction
- [scalard tx evm create-burn-tokens](scalard_tx_evm_create-burn-tokens.md) - Create burn commands for all confirmed token deposits in an EVM chain
- [scalard tx evm create-deploy-token](scalard_tx_evm_create-deploy-token.md) - Create a deploy token command with the ScalarGateway contract
- [scalard tx evm create-pending-transfers](scalard_tx_evm_create-pending-transfers.md) - Create commands for handling all pending transfers to an EVM chain
- [scalard tx evm link](scalard_tx_evm_link.md) - Link a cross chain address to an EVM chain address created by Scalar
- [scalard tx evm retry-event](scalard_tx_evm_retry-event.md) - Retry a failed event
- [scalard tx evm set-gateway](scalard_tx_evm_set-gateway.md) - Set the gateway address for the given evm chain
- [scalard tx evm sign-commands](scalard_tx_evm_sign-commands.md) - Sign pending commands for an EVM chain contract
- [scalard tx evm transfer-operatorship](scalard_tx_evm_transfer-operatorship.md) - Create transfer operatorship command for an EVM chain contract
