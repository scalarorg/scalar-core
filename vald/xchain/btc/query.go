package btc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"

	chainsExported "github.com/scalarorg/scalar-core/x/chains/exported"
	covExported "github.com/scalarorg/scalar-core/x/covenant/exported"
)

type MempoolUtxo struct {
	Txid   string `json:"txid"`
	Vout   uint32 `json:"vout"`
	Status struct {
		Confirmed   bool   `json:"confirmed"`
		BlockHeight uint64 `json:"block_height"`
		BlockHash   string `json:"block_hash"`
		BlockTime   uint64 `json:"block_time"`
	} `json:"status"`
	Value uint64 `json:"value"`
}

func (c *BtcClient) getUtxoList(taprootAddress string) ([]*covExported.UTXO, error) {
	if c.mempoolUrl == "" {
		return nil, fmt.Errorf("mempool URL is not set")
	}
	url := fmt.Sprintf("%s/address/%s/utxo", c.mempoolUrl, taprootAddress)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get UTXOs: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	//log.Debug().Msgf("[BtcClient] [GetAddressTxsUtxo] body: %v", string(body))

	var utxos []MempoolUtxo
	if err := json.Unmarshal(body, &utxos); err != nil {
		return nil, fmt.Errorf("failed to decode UTXOs: %w", err)
	}
	utxos = sortUTXOsByValue(utxos)

	utxosList := []*covExported.UTXO{}
	for _, utxo := range utxos {
		txID, err := chainsExported.HashFromHex(utxo.Txid)
		if err != nil {
			return nil, fmt.Errorf("failed to convert txid to hash: %w", err)
		}
		utxosList = append(utxosList, &covExported.UTXO{
			TxID:         txID,
			Vout:         utxo.Vout,
			AmountInSats: utxo.Value,
			Reserved:     make(map[string]uint64),
		})
	}

	return utxosList, nil

}

func sortUTXOsByValue(utxos []MempoolUtxo) []MempoolUtxo {
	sort.Slice(utxos, func(i, j int) bool {
		return utxos[i].Value > utxos[j].Value
	})
	return utxos
}
