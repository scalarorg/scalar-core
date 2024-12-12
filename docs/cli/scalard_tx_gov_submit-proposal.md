## scalard tx gov submit-proposal

Submit a proposal along with an initial deposit

### Synopsis

Submit a proposal along with an initial deposit.
Proposal title, description, type and deposit can be given directly or through a proposal JSON file.

Example:
$ <appd> tx gov submit-proposal --proposal="path/to/proposal.json" --from mykey

Where proposal.json contains:

{
"title": "Test Proposal",
"description": "My awesome proposal",
"type": "Text",
"deposit": "10test"
}

Which is equivalent to:

$ <appd> tx gov submit-proposal --title="Test Proposal" --description="My awesome proposal" --type="Text" --deposit="10test" --from mykey

```
scalard tx gov submit-proposal [flags]
```

### Options

```
  -a, --account-number uint      The account number of the signing account (offline mode only)
  -b, --broadcast-mode string    Transaction broadcasting mode (sync|async|block) (default "block")
      --deposit string           The proposal deposit
      --description string       The proposal description
      --dry-run                  ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it (when enabled, the local Keybase is not accessible)
      --fee-account string       Fee account pays fees for the transaction instead of deducting from the signer
      --fees string              Fees to pay along with transaction; eg: 10uatom
      --from string              Name or address of private key with which to sign
      --gas string               gas limit to set per-transaction; set to "auto" to calculate sufficient gas automatically (default 200000)
      --gas-adjustment float     adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set manually this flag is ignored  (default 1)
      --gas-prices string        Gas prices in decimal format to determine the transaction fee (e.g. 0.1uatom) (default "0.007scal")
      --generate-only            Build an unsigned transaction and write it to STDOUT (when enabled, the local Keybase is not accessible)
  -h, --help                     help for submit-proposal
      --keyring-backend string   Select keyring's backend (os|file|kwallet|pass|test|memory) (default "file")
      --keyring-dir string       The client Keyring directory; if omitted, the default 'home' directory will be used
      --ledger                   Use a connected Ledger device
      --node string              <host>:<port> to tendermint rpc interface for this chain (default "tcp://localhost:26657")
      --note string              Note to add a description to the transaction (previously --memo)
      --offline                  Offline mode (does not allow any online functionality
  -o, --output string            Output format (text|json) (default "json")
      --proposal string          Proposal file path (if this path is given, other proposal flags are ignored)
  -s, --sequence uint            The sequence number of the signing account (offline mode only)
      --sign-mode string         Choose sign mode (direct|amino-json), this is an advanced feature
      --timeout-height uint      Set a block timeout height to prevent the tx from being committed past a certain height
      --title string             The proposal title
      --type string              The proposal Type
  -y, --yes                      Skip tx broadcasting prompt confirmation (default true)
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

- [scalard tx gov](scalard_tx_gov.md) - Governance transactions subcommands
- [scalard tx gov submit-proposal call-contracts](scalard_tx_gov_submit-proposal_call-contracts.md) - Submit a call contracts proposal
- [scalard tx gov submit-proposal cancel-software-upgrade](scalard_tx_gov_submit-proposal_cancel-software-upgrade.md) - Cancel the current software upgrade proposal
- [scalard tx gov submit-proposal community-pool-spend](scalard_tx_gov_submit-proposal_community-pool-spend.md) - Submit a community pool spend proposal
- [scalard tx gov submit-proposal ibc-upgrade](scalard_tx_gov_submit-proposal_ibc-upgrade.md) - Submit an IBC upgrade proposal
- [scalard tx gov submit-proposal param-change](scalard_tx_gov_submit-proposal_param-change.md) - Submit a parameter change proposal
- [scalard tx gov submit-proposal software-upgrade](scalard_tx_gov_submit-proposal_software-upgrade.md) - Submit a software upgrade proposal
- [scalard tx gov submit-proposal update-client](scalard_tx_gov_submit-proposal_update-client.md) - Submit an update IBC client proposal
