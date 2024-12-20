package types

import (
	fmt "fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

func ScriptPubKeyToAddress(scriptPubKey []byte, params *chaincfg.Params) (btcutil.Address, error) {
	// Extract the type of script
	_, addresses, _, err := txscript.ExtractPkScriptAddrs(scriptPubKey, params)
	if err != nil {
		return nil, err
	}

	// Usually we take the first address, but some scripts might have multiple
	if len(addresses) == 0 {
		return nil, fmt.Errorf("no addresses found")
	}

	// TODO: Just support the simple case for now
	return addresses[0], nil
}
