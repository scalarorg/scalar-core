## scalard query chains chains

Return the supported EVM chains by status

```
scalard query chains chains [flags]
```

### Options

```
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for chains
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")
      --status string   the chain status [activated|deactivated]
```

### Options inherited from parent commands

```
      --chain-id string     The network chain ID (default "scalar")
      --home string         directory for config and data (default "$HOME/.scalar")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

### SEE ALSO

* [scalard query chains](scalard_query_chains.md)	 - Querying commands for the chains module

