## scalard testnet start

Launch an in-process multi-validator testnet

### Synopsis

testnet will launch an in-process multi-validator testnet,
and generate "v" directories, populated with necessary validator configuration files
(private validator, genesis, config, etc.).

Example:
scalard testnet start --v 4 --base-dir ./.testnets

```
scalard testnet start [flags]
```

### Options

```
      --base-dir string             the base directory to store the testnet (default "./.testnets")
      --base-fee string             The params base_fee in the feemarket module in geneis (default "1000000000")
      --block-height int            The block height to stop the testnet (default 100)
      --chain-id string             genesis file chain-id, if left blank will be randomly created
  -h, --help                        help for start
      --key-type string             Key signing algorithm to generate keys for (default "secp256k1")
      --min-gas-price string        The params min_gas_price in the feemarket module in geneis (default "0.001")
      --minimum-gas-prices string   Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01photino,0.001stake) (default "0.000006ascal")
  -o, --output-dir string           Directory to store initialization data for the testnet (default "./.testnets")
      --timeout int                 The testnet run time. Default is 1800 seconds (default 1800)
      --v int                       Number of validators to initialize the testnet with (default 4)
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

- [scalard testnet](scalard_testnet.md) - subcommands for starting or configuring local testnets
