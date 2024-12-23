package rpc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/scalarorg/scalar-core/client/rpc"
	"github.com/scalarorg/scalar-core/utils"
	chainTypes "github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

const (
	mockTxHash = "07b50c84f889e2f1315da875fc91734e2bac8d0153ff9a98d9da14caa4fc7d57"
)

func TestConfirmSourceTx(t *testing.T) {
	require.NotNil(t, mockNetworkClient)

	chain := nexus.ChainName(utils.NormalizeString(chainNameBtcTestnet4))
	txHash, err := chainTypes.HashFromHexStr(mockTxHash)
	require.NoError(t, err)
	msg := chainTypes.NewConfirmSourceTxsRequest(mockNetworkClient.GetAddress(), chain, []chainTypes.Hash{*txHash})

	tx, err := rpc.ConfirmSourceTx(context.Background(), mockNetworkClient, msg)
	require.NoError(t, err)
	require.NotNil(t, tx)
}
