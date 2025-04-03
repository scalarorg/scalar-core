package types_test

import (
	"bytes"
	"testing"

	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/stretchr/testify/require"
)

func TestUTXOAppendReserved(t *testing.T) {
	utxo := types.UTXO{
		TxID:         chains.Hash{},
		Vout:         0,
		AmountInSats: 100,
	}
	err := utxo.AppendReserved("request1", 50)
	require.NoError(t, err)
	require.Equal(t, uint64(50), utxo.AvailableAmount())
	require.Equal(t, uint64(50), utxo.GetReservedAmount())
	err = utxo.AppendReserved("request2", 50)
	require.NoError(t, err)
	require.Equal(t, uint64(0), utxo.AvailableAmount())
	require.Equal(t, uint64(100), utxo.GetReservedAmount())
	err = utxo.AppendReserved("request3", 50)
	require.Error(t, err)
	utxo.Release("request1")
	require.Equal(t, uint64(50), utxo.AvailableAmount())
	require.Equal(t, uint64(50), utxo.GetReservedAmount())
}

func TestUTXOSnapshotReserveUtxos(t *testing.T) {
	snapshot := types.UTXOSnapshot{
		Utxos: []*types.UTXO{
			{
				TxID:         chains.Hash(bytes.Repeat([]byte{0x01}, 32)),
				Vout:         0,
				AmountInSats: 100,
			},
			{
				TxID:         chains.Hash(bytes.Repeat([]byte{0x02}, 32)),
				Vout:         0,
				AmountInSats: 200,
			},
		},
	}
	utxos, err := snapshot.ReserveUtxos("request1", 150)
	require.NoError(t, err)
	require.Equal(t, 2, len(utxos))
	require.Equal(t, uint64(0), snapshot.Utxos[0].AvailableAmount())
	require.Equal(t, uint64(50), snapshot.Utxos[1].GetReservedAmount())
	utxos, err = snapshot.ReserveUtxos("request2", 50)
	require.NoError(t, err)
	require.Equal(t, 1, len(utxos))
	require.Equal(t, uint64(0), snapshot.Utxos[0].AvailableAmount())
	require.Equal(t, uint64(100), snapshot.Utxos[1].GetReservedAmount())
}
