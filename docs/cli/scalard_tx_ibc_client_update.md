## scalard tx ibc client update

update existing client with a header

### Synopsis

update existing client with a header

```
scalard tx ibc client update [client-id] [path/to/header.json] [flags]
```

### Examples

```
<appd> tx ibc client update [client-id] [path/to/header.json] --from node0 --home ../node0/<app>cli --chain-id $CID
```

### Options

```
  -h, --help   help for update
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

- [scalard tx ibc client](scalard_tx_ibc_client.md) - IBC client transaction subcommands
