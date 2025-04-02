package btc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
	covTypes "github.com/scalarorg/scalar-core/x/covenant/types"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *BtcClient) ProcessInitializeUtxo(event *covTypes.IntializeUtxoSnapshotStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {

	// info, err := client.getBlockChainInfo()
	// if err != nil {
	// 	log.Error().Msgf("[ProcessInitializeUtxo] failed to get block chain info")
	// 	return nil, err
	// }

	// diff := uint64(info.Blocks) - event.BlockCheckpoint
	// if diff != event.ConfirmationHeight || diff != event.ConfirmationHeight+1 {
	// 	return nil, fmt.Errorf("[ProcessInitializeUtxo] blockInfo %+v, event.BlockCheckpoint %d, event.ConfirmationHeight %d",
	// 		info, event.BlockCheckpoint, event.ConfirmationHeight)
	// }

	// TODO: Ensure the utxos are confirmed with the correct number of confirmations
	//Ouput utxos already sorted by block height and txId
	utxos, blockHeights, err := client.getUtxoList(event.Address)
	if err != nil {
		log.Error().Msgf("[ProcessInitializeUtxo] failed to get utxo list for address %s", event.Address)
		return nil, err
	}

	filteredUtxos := []*covTypes.UTXO{}
	youngtUtxos := []*covTypes.UTXO{}
	expectedBlockHeight := event.BlockCheckpoint - event.ConfirmationHeight
	for i, utxo := range utxos {
		if blockHeights[i] <= expectedBlockHeight {
			filteredUtxos = append(filteredUtxos, utxo)
		} else {
			log.Info().Msgf("[ProcessInitializeUtxo] utxo does not have enought confirmations %+v", utxo)
			youngtUtxos = append(youngtUtxos, utxo)
		}
	}
	utxoSnapshot := covTypes.UTXOSnapshot{
		CustodianGroupUID: event.CustodianGroupUID,
		BlockHeight:       event.BlockCheckpoint,
		Utxos:             filteredUtxos,
	}
	hash := utxoSnapshot.GetHash()
	log.Info().Msgf("[ProcessInitializeUtxo] total utxo: %d, filtered utxos: %d, utxos with small confirmations: %d. Snapshot hash: %s",
		len(utxos), len(filteredUtxos), len(youngtUtxos), hash.Hex())

	voteEvent := covTypes.NewVoteEvents(event.Chain, covTypes.Event{
		Chain: event.Chain,
		Hash:  &hash,
		Event: &covTypes.Event_IntializeUtxoSnapshotCompleted{
			IntializeUtxoSnapshotCompleted: &covTypes.IntializeUtxoSnapshotCompleted{
				UtxoSnapshot: &utxoSnapshot,
			},
		},
		Index: 0,
	})

	votes := []sdk.Msg{
		voteTypes.NewVoteRequest(proxy, event.PollID, voteEvent),
	}

	return votes, nil
}
