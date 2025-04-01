package btc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/utils/clog"
	covTypes "github.com/scalarorg/scalar-core/x/covenant/types"
	voteTypes "github.com/scalarorg/scalar-core/x/vote/types"
)

func (client *BtcClient) ProcessInitializeUtxo(event *covTypes.IntializeUtxoSnapshotStarted, proxy sdk.AccAddress) ([]sdk.Msg, error) {
	// TODO: Ensure the utxos are confirmed with the correct number of confirmations
	utxos, err := client.getUtxoList(event.Address)
	if err != nil {
		return nil, err
	}

	voteEvent := covTypes.NewVoteEvents(event.Chain, covTypes.Event{
		Chain: event.Chain,
		Hash:  nil,
		Event: &covTypes.Event_IntializeUtxoSnapshotCompleted{
			IntializeUtxoSnapshotCompleted: &covTypes.IntializeUtxoSnapshotCompleted{
				UtxoSnapshot: &covTypes.UTXOSnapshot{
					CustodianGroupUID: event.CustodianGroupUID,
					BlockHeight:       1, // TODO: fill me
					Utxos:             utxos,
				},
			},
		},
		Index: 0,
	})

	events := covTypes.Event{}
	votes := []sdk.Msg{
		voteTypes.NewVoteRequest(proxy, event.PollID, voteEvent),
	}
	clog.Redf("broadcasting vote %v for poll %s", events, event.PollID.String())

	return votes, nil
}
