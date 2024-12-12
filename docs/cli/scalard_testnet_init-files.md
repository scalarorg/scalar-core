## scalard testnet init-files

Initialize config directories & files for a multi-validator testnet running locally via separate processes (e.g. Docker Compose or similar)

### Synopsis

init-files will setup "v" number of directories and populate each with
necessary files (private validator, genesis, config, etc.) for running "v" validator nodes.

Booting up a network with these validator folders is intended to be used with Docker Compose,
or a similar setup where each node has a manually configurable IP address.

Note, strict routability for addresses is turned off in the config file.

Example:
scalard testnet init-files --v 4 --output-dir ./.testnets --node-domain scalarnode --supported-chains=./chains --env-file=.env.local

```
scalard testnet init-files [flags]
```

### Options

```
      --base-fee string             The params base_fee in the feemarket module in geneis (default "1000000000")
      --chain-id string             genesis file chain-id, if left blank will be randomly created
      --env-file string             Path to environment file to load (optional)
  -h, --help                        help for init-files
      --key-type string             Key signing algorithm to generate keys for (default "secp256k1")
      --keyring-backend string      Select keyring's backend (os|file|test) (default "os")
      --min-gas-price string        The params min_gas_price in the feemarket module in geneis (default "0.001")
      --minimum-gas-prices string   Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01photino,0.001stake) (default "0.000006ascal")
      --node-daemon-home string     Home directory of the node's daemon configuration (default "scalard")
      --node-dir-prefix string      Prefix the directory name for each node with (node results in node1, node2, ...) (default "node")
      --node-domain string          Node domain: 
                                    		*scalarnode* results in persistent peers list ID0@scalarnode1:46656, ID1@scalarnode2:46656, ...
                                    		*192.168.0.1* results in persistent peers list ID0@192.168.0.11:46656, ID1@192.168.0.12:46656, ...
                                    		
  -o, --output-dir string           Directory to store initialization data for the testnet (default "./.testnets")
      --port-offset int             Port offset for the testnet
      --supported-chains string     Supported chains directory, each chain family is config in a seperated json file under this directory: 
                                    		*./chains/evm.json* stores all evm chain configs ...
                                    		*./chains/btc.json* stores all btc chain configs ...
                                    		 (default "./chains")
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
