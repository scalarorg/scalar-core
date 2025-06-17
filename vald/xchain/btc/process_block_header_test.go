package btc_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/chains/types"
	votes "github.com/scalarorg/scalar-core/x/vote/exported"
	"github.com/stretchr/testify/require"
)

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestProcessNewBlockConfirmation$ github.com/scalarorg/scalar-core/vald/xchain/btc -v -count=1
func TestProcessNewBlockConfirmation(t *testing.T) {

	prevBlockHash := exported.Hash(common.HexToHash("00000000000000030642a17ebbc7ebb4215b50618983087427c0d2632fd1eeaa").Bytes())

	vote, err := mockBtcClient.ProcessNewBlockConfirmation(&types.ConfirmNewBlockStarted{
		Chain: "btc",
		PollParticipants: votes.PollParticipants{
			PollID: "1",
			Participants: []sdk.ValAddress{
				sdk.ValAddress(common.HexToAddress("0x1").Bytes()),
			},
		},
		BlockHash:          exported.Hash(common.HexToHash("0000000009cc1e47605fdd3ecd7bfedbed161ae82afd555e4cd760c8d9eecb3d").Bytes()),
		PreviousBlockHash:  &prevBlockHash,
		ConfirmationHeight: 2,
	}, []byte{})

	require.NoError(t, err)
	require.NotNil(t, vote)
}
