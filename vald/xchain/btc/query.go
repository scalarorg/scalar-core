package btc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/rs/zerolog/log"
	chainsExported "github.com/scalarorg/scalar-core/x/chains/exported"
	cov "github.com/scalarorg/scalar-core/x/covenant/types"
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

func (c *BtcClient) getUtxoList(taprootAddress string) ([]*cov.UTXO, []uint64, error) {
	if c.mempoolUrl == "" {
		return nil, nil, fmt.Errorf("mempool URL is not set")
	}
	url := fmt.Sprintf("%s/address/%s/utxo", c.mempoolUrl, taprootAddress)

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get UTXOs: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	//log.Debug().Msgf("[BtcClient] [GetAddressTxsUtxo] body: %v", string(body))

	var utxos []MempoolUtxo
	if err := json.Unmarshal(body, &utxos); err != nil {
		return nil, nil, fmt.Errorf("failed to decode UTXOs: %w", err)
	}
	log.Info().Msgf("[GetUtxoList] utxos length: %d", len(utxos))
	utxos = SortUTXOsByBlockHeight(utxos)
	utxosList := []*cov.UTXO{}
	blockHeights := make([]uint64, len(utxos))
	for _, utxo := range utxos {
		//log.Info().Msgf("[GetUtxoList] block height: %d, txid: %s, vout: %d, amount: %d", utxo.Status.BlockHeight, utxo.Txid, utxo.Vout, utxo.Value)
		txID, err := chainsExported.HashFromHex(utxo.Txid)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert txid to hash: %w", err)
		}
		utxosList = append(utxosList, &cov.UTXO{
			TxID:         txID,
			Vout:         utxo.Vout,
			AmountInSats: utxo.Value,
			Reservations: []*cov.Reservation{},
		})
		blockHeights = append(blockHeights, utxo.Status.BlockHeight)
	}

	return utxosList, blockHeights, nil
}

func SortUTXOsByBlockHeight(utxos []MempoolUtxo) []MempoolUtxo {
	sort.Slice(utxos, func(i, j int) bool {
		return (utxos[i].Status.BlockHeight < utxos[j].Status.BlockHeight) ||
			(utxos[i].Status.BlockHeight == utxos[j].Status.BlockHeight && utxos[i].Txid < utxos[j].Txid) ||
			(utxos[i].Status.BlockHeight == utxos[j].Status.BlockHeight && utxos[i].Txid == utxos[j].Txid && utxos[i].Vout < utxos[j].Vout)
	})
	return utxos
}
