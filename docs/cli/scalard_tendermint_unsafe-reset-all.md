## scalard tendermint unsafe-reset-all

(unsafe) Remove all the data and WAL, reset this node's validator to genesis state

```
scalard tendermint unsafe-reset-all [flags]
```

### Options

```
  -h, --help             help for unsafe-reset-all
      --keep-addr-book   keep the address book intact
```

### Options inherited from parent commands

```
      --home string         directory for config and data (default "$HOME/.scalar")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --output string       Output format (text|json) (default "text")
      --trace               print out full stack trace on errors
```

### SEE ALSO

- [scalard tendermint](scalard_tendermint.md) - Tendermint subcommands
