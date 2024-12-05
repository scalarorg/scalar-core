package rpc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/axelarnetwork/axelar-core/utils"
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	"github.com/scalarorg/scalar-core/client/rpc"
	btcTypes "github.com/scalarorg/scalar-core/x/btc/types"
)

const (
	mockTxHash = "07b50c84f889e2f1315da875fc91734e2bac8d0153ff9a98d9da14caa4fc7d57"
)

func TestConfirmBtcTx(t *testing.T) {
	require.NotNil(t, mockNetworkClient)

	chain := nexus.ChainName(utils.NormalizeString(chainNameBtcTestnet4))
	txHash, err := btcTypes.TxHashFromHexStr(mockTxHash)
	require.NoError(t, err)
	msg := btcTypes.NewConfirmGatewayTxsRequest(mockNetworkClient.GetAddress(), chain, []btcTypes.TxHash{*txHash})

	tx, err := rpc.ConfirmBtcTx(context.Background(), mockNetworkClient, msg)
	require.NoError(t, err)
	require.NotNil(t, tx)
}
