package types

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/scalarorg/scalar-core/testutils/rand"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	covTypes "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/scalarorg/scalar-core/x/vote/exported"
)


// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestVoteRequest_ValidateBasic$ github.com/scalarorg/scalar-core/x/vote/types -v -count=1
func TestVoteRequest_ValidateBasic(t *testing.T) {
	t.Run("correct vote events", func(t *testing.T) {
		chain := nexus.ChainName("bitcoin|4")
		voteEvent := covTypes.NewVoteEvents(chain, covTypes.Event{
			Chain: chain,
			Hash:  nil,
			Event: &covTypes.Event_RedeemTxsConfirmed{
				RedeemTxsConfirmed: &covTypes.RedeemTxsConfirmed{
					EventIDs: []chainsTypes.EventID{},
					UtxoSnapshot: &covTypes.UTXOSnapshot{
						CustodianGroupUID: chains.Hash(rand.Bytes(32)),
						BlockHeight:       1, // TODO: fill me
						Utxos:             []*covTypes.UTXO{},
					},
				},
			},
			Index: 0,
		})

		vote := NewVoteRequest(rand.AccAddr(), exported.PollID(rand.PosI64()), voteEvent)
		assert.NoError(t, vote.ValidateBasic())
	})
}
