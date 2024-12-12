## scalard query feegrant grant

Query details of a single grant

### Synopsis

Query details for a grant.
You can find the fee-grant of a granter and grantee.

Example:
$ <appd> query feegrant grant [granter] [grantee]

```
scalard query feegrant grant [granter] [grantee] [flags]
```

### Options

```
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for grant
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")
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

- [scalard query feegrant](scalard_query_feegrant.md) - Querying commands for the feegrant module
