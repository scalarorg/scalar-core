## scalard query scalarnet

Querying commands for the scalarnet module

```
scalard query scalarnet [flags]
```

### Options

```
  -h, --help   help for scalarnet
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
- [scalard query scalarnet chain-by-ibc-path](scalard_query_scalarnet_chain-by-ibc-path.md) - Returns the Cosmos chain for the given IBC path
- [scalard query scalarnet ibc-path](scalard_query_scalarnet_ibc-path.md) - Returns the registered IBC path for the given Cosmos chain
- [scalard query scalarnet ibc-transfer-count](scalard_query_scalarnet_ibc-transfer-count.md) - returns the number of pending IBC transfers per chain
- [scalard query scalarnet params](scalard_query_scalarnet_params.md) - Returns the params for the scalarnet module
