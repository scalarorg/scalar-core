package btc_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestProcessNewBlockConfirmation$ github.com/scalarorg/scalar-core/vald/xchain/btc -v -count=1
func TestProcessNewBlockConfirmation(t *testing.T) {
	vote, err := mockBtcClient.ProcessNewBlockConfirmation(&types.ConfirmBtcNewBlockStarted{
		PollID: "1",
		Chain:  "btc",
		Participants: []sdk.ValAddress{
			sdk.ValAddress(common.HexToAddress("0x1").Bytes()),
		},
		BlockHash:          exported.Hash(common.HexToHash("0000000009cc1e47605fdd3ecd7bfedbed161ae82afd555e4cd760c8d9eecb3d").Bytes()),
		PreviousBlockHash:  exported.Hash(common.HexToHash("00000000000000030642a17ebbc7ebb4215b50618983087427c0d2632fd1eeaa").Bytes()),
		ConfirmationHeight: 2,
	}, []byte{})

	require.NoError(t, err)
	require.NotNil(t, vote)
}
