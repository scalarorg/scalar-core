package types_test

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/scalarorg/bitcoin-vault/go-utils/encode"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
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

func TestRedeemTokenParamsAbiPack(t *testing.T) {
	mockCmdID := types.NewCommandID(bytes.Repeat([]byte("a"), 32))
	mockCustodianGrUID := chains.Hash(bytes.Repeat([]byte("b"), 32))
	mockAddress := "0x0000000000000000000000000000000000000000"
	mockChainName := nexus.ChainName("bitcoin|4")
	mokRequestAmount := big.NewInt(100_000)
	mockSymbol := "BTC"
	mockLockingScript := []byte("locking-script")
	mockUtxos := []*types.UTXO{
		{
			TxID:         chains.Hash(bytes.Repeat([]byte{0x01}, 32)),
			Vout:         0,
			AmountInSats: 50_000,
		},
		{
			TxID:         chains.Hash(bytes.Repeat([]byte{0x02}, 32)),
			Vout:         0,
			AmountInSats: 70_000,
		},
	}
	payload := types.RedeemTokenPayload{
		Amount:        mokRequestAmount.Uint64(),
		LockingScript: mockLockingScript,
		Utxos:         mockUtxos,
		RequestId:     mockCmdID,
	}
	payloadBytes, err := payload.AbiPack()
	require.NoError(t, err)
	payload2 := types.RedeemTokenPayload{}
	err = payload2.AbiUnpack(payloadBytes)
	require.NoError(t, err)
	require.Equal(t, payload, payload2)
	params := types.RedeemTokenParams{
		DestinationChain:   mockChainName.String(),
		DestinationAddress: mockAddress,
		Payload:            payload,
		Symbol:             mockSymbol,
		Amount:             mokRequestAmount.Uint64(),
		CustodianGroupUID:  mockCustodianGrUID,
		SessionSequence:    1,
	}

	packed, err := params.AbiPack()
	require.NoError(t, err)

	unpackedParams := types.RedeemTokenParams{}
	err = unpackedParams.AbiUnpack(packed)
	require.NoError(t, err)
	t.Logf("unpackedParams: %+v", unpackedParams)
	require.Equal(t, params, unpackedParams)

	// remove prefix
	_, prefix := unpackedParams.RawPayload[1:], unpackedParams.RawPayload[0:1]
	require.Equal(t, encode.ContractCallWithTokenPayloadType_CustodianOnly.Bytes(), prefix)

	require.Equal(t, mockUtxos[0].TxID.Hex(), unpackedParams.Payload.Utxos[0].TxID.Hex())
	require.Equal(t, mockUtxos[1].TxID.Hex(), unpackedParams.Payload.Utxos[1].TxID.Hex())
	require.Equal(t, mockUtxos[0].Vout, unpackedParams.Payload.Utxos[0].Vout)
	require.Equal(t, mockUtxos[1].Vout, unpackedParams.Payload.Utxos[1].Vout)
	require.Equal(t, mockUtxos[0].AmountInSats, unpackedParams.Payload.Utxos[0].AmountInSats)
	require.Equal(t, mockUtxos[1].AmountInSats, unpackedParams.Payload.Utxos[1].AmountInSats)
	require.Equal(t, mockCmdID, unpackedParams.Payload.RequestId)
	require.Equal(t, mockSymbol, unpackedParams.Symbol)
	require.Equal(t, mokRequestAmount.Uint64(), unpackedParams.Amount)
	require.Equal(t, mockCustodianGrUID, unpackedParams.CustodianGroupUID)
	require.Equal(t, uint64(1), unpackedParams.SessionSequence)
}
