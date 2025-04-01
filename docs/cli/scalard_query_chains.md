## scalard query chains

Querying commands for the chains module

```
scalard query chains [flags]
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

* [scalard query](scalard_query.md)	 - Querying subcommands
* [scalard query chains address](scalard_query_chains_address.md)	 - Returns the EVM address
* [scalard query chains batched-commands](scalard_query_chains_batched-commands.md)	 - Get the signed batched commands that can be wrapped in an transaction to be executed in Scalar Gateway
* [scalard query chains burner-info](scalard_query_chains_burner-info.md)	 - Get information about a burner address
* [scalard query chains bytecode](scalard_query_chains_bytecode.md)	 - Fetch the bytecode of an EVM contract [contract] for chain [chain]
* [scalard query chains chains](scalard_query_chains_chains.md)	 - Return the supported EVM chains by status
* [scalard query chains command](scalard_query_chains_command.md)	 - Get information about an EVM gateway command given a chain and the command ID
* [scalard query chains confirmation-height](scalard_query_chains_confirmation-height.md)	 - Returns the minimum confirmation height for the given chain
* [scalard query chains erc20-tokens](scalard_query_chains_erc20-tokens.md)	 - Returns the ERC20 tokens for the given chain
* [scalard query chains event](scalard_query_chains_event.md)	 - Returns an event for the given chain
* [scalard query chains gateway-address](scalard_query_chains_gateway-address.md)	 - Query the Scalar Gateway contract address
* [scalard query chains latest-batched-commands](scalard_query_chains_latest-batched-commands.md)	 - Get the latest batched commands that can be wrapped in an EVM transaction to be executed in Scalar Gateway
* [scalard query chains params](scalard_query_chains_params.md)	 - Returns the params for the evm module
* [scalard query chains pending-commands](scalard_query_chains_pending-commands.md)	 - Get the list of commands not yet added to a batch
* [scalard query chains token-info](scalard_query_chains_token-info.md)	 - Returns the info of token by either symbol, asset, or address

