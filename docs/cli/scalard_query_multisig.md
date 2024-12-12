## scalard query multisig

Querying commands for the multisig module

```
scalard query multisig [flags]
```

### Options

```
  -h, --help   help for multisig
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
- [scalard query multisig key](scalard_query_multisig_key.md) - Returns the key of the given ID
- [scalard query multisig key-id](scalard_query_multisig_key-id.md) - Returns the key ID assigned to a given chain
- [scalard query multisig keygen-session](scalard_query_multisig_keygen-session.md) - Returns the keygen session info for the given key ID
- [scalard query multisig next-key-id](scalard_query_multisig_next-key-id.md) - Returns the key ID assigned for the next rotation on a given chain and for the given key role
- [scalard query multisig params](scalard_query_multisig_params.md) - Returns the params for the multisig module
