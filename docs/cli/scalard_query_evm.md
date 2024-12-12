## scalard query evm

Querying commands for the evm module

```
scalard query evm [flags]
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

- [scalard query](scalard_query.md) - Querying subcommands
- [scalard query evm address](scalard_query_evm_address.md) - Returns the EVM address
- [scalard query evm batched-commands](scalard_query_evm_batched-commands.md) - Get the signed batched commands that can be wrapped in an EVM transaction to be executed in Axelar Gateway
- [scalard query evm burner-info](scalard_query_evm_burner-info.md) - Get information about a burner address
- [scalard query evm bytecode](scalard_query_evm_bytecode.md) - Fetch the bytecode of an EVM contract [contract] for chain [chain]
- [scalard query evm chains](scalard_query_evm_chains.md) - Return the supported EVM chains by status
- [scalard query evm command](scalard_query_evm_command.md) - Get information about an EVM gateway command given a chain and the command ID
- [scalard query evm confirmation-height](scalard_query_evm_confirmation-height.md) - Returns the minimum confirmation height for the given chain
- [scalard query evm erc20-tokens](scalard_query_evm_erc20-tokens.md) - Returns the ERC20 tokens for the given chain
- [scalard query evm event](scalard_query_evm_event.md) - Returns an event for the given chain
- [scalard query evm gateway-address](scalard_query_evm_gateway-address.md) - Query the Axelar Gateway contract address
- [scalard query evm latest-batched-commands](scalard_query_evm_latest-batched-commands.md) - Get the latest batched commands that can be wrapped in an EVM transaction to be executed in Axelar Gateway
- [scalard query evm params](scalard_query_evm_params.md) - Returns the params for the evm module
- [scalard query evm pending-commands](scalard_query_evm_pending-commands.md) - Get the list of commands not yet added to a batch
- [scalard query evm token-address](scalard_query_evm_token-address.md) - Query a token address by by either symbol or asset
- [scalard query evm token-info](scalard_query_evm_token-info.md) - Returns the info of token by either symbol, asset, or address
