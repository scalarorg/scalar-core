## scalard query ibc client status

Query client status

### Synopsis

Query client activity status. Any client without an 'Active' status is considered inactive

```
scalard query ibc client status [client-id] [flags]
```

### Examples

```
<appd> query ibc client status [client-id]
```

### Options

```
  -h, --help   help for status
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

- [scalard query ibc client](scalard_query_ibc_client.md) - IBC client query subcommands
