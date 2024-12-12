## scalard tx distribution

Distribution transactions subcommands

```
scalard tx distribution [flags]
```

### Options

```
  -h, --help   help for distribution
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
- [scalard tx distribution fund-community-pool](scalard_tx_distribution_fund-community-pool.md) - Funds the community pool with the specified amount
- [scalard tx distribution set-withdraw-addr](scalard_tx_distribution_set-withdraw-addr.md) - change the default withdraw address for rewards associated with an address
- [scalard tx distribution withdraw-all-rewards](scalard_tx_distribution_withdraw-all-rewards.md) - withdraw all delegations rewards for a delegator
- [scalard tx distribution withdraw-rewards](scalard_tx_distribution_withdraw-rewards.md) - Withdraw rewards from a given delegation address, and optionally withdraw validator commission if the delegation address given is a validator operator
