## scalard query protocol

Querying commands for the protocol module

```
scalard query protocol [flags]
```

### Options

```
      --address string   Protocol address to find
      --height int       Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help             help for protocol
      --name string      Name of the protocol
      --node string      <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string    Output format (text|json) (default "text")
      --pubkey string    Protocol pubkey
      --status string    Status of the protocol (Unspecified|Activated|Deactived) (default "Activated")
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

* [scalard query](scalard_query.md)	 - Querying subcommands

