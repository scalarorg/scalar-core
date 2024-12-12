## scalard tx bank

Bank transaction subcommands

```
scalard tx bank [flags]
```

### Options

```
  -h, --help   help for bank
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
- [scalard tx bank send](scalard_tx_bank_send.md) - Send funds from one account to another.
  Note, the'--from' flag is ignored as it is implied from [from_key_or_address].
  When using '--dry-run' a key name cannot be used, only a bech32 address.
