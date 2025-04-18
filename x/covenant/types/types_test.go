package types_test

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/stretchr/testify/require"
)

func TestPayloadAbiPack(t *testing.T) {
	params := types.RedeemTokenParams{}
	//paramHex := "00000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000003f54e38e86ec074d00f88ed8a4ec6e50d9fd10731d6b5991b98ee09b3d4d68200000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000003e87c85f0bf8ebc27060fafc126f5880347fda6faf704e12a52f644e0cd0a4a26000000000000000000000000000000000000000000000000000000000000000066000000000000000000000000000000000000000000000000000000000000000c65766d7c31313135353131310000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002a3078304364373766316630654230463534373764633546354436453661373230306466393465344631320000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000047342746300000000000000000000000000000000000000000000000000000000"
	payloadHex := "0000000000000000000000000000000000000000000000000000000000000003E800000000000000000000000000000000000000000000000000000000000000C0000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001C0000000000000000000000000000000000000000000000000000000000000020026ACC6A0A92E0B94AFD58716442F3FDF1A10EE0C1B69131637A6709DB3DD33C30000000000000000000000000000000000000000000000000000000000000016001463DC22751D9A7778AA4450CEEB0B5C3EE214401C0000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004230783733373763626133333663666636653064333934373837323234366339396664383039353035336631386461623239616137633663303632336431376330353200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000003E8"
	payloadBytes, err := hex.DecodeString(payloadHex)
	require.NoError(t, err)
	err = params.AbiUnpack(payloadBytes)
	require.NoError(t, err)
	t.Logf("params: %+v", params)
}
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
	t.Logf("Payload pack and unpack successfully")
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
	t.Logf("packed: %x", packed)
	require.NoError(t, err)

	unpackedParams := types.RedeemTokenParams{}
	err = unpackedParams.AbiUnpack(packed)
	require.NoError(t, err)
	t.Logf("unpackedParams: %+v", unpackedParams)
	//require.Equal(t, params, unpackedParams)

	// remove prefix
	// _, prefix := unpackedParams.RawPayload[1:], unpackedParams.RawPayload[0:1]
	// require.Equal(t, encode.ContractCallWithTokenPayloadType_CustodianOnly.Bytes(), prefix)

	require.Equal(t, mockUtxos[0].TxID.Hex(), unpackedParams.Payload.Utxos[0].TxID.Hex())
	require.Equal(t, mockUtxos[1].TxID.Hex(), unpackedParams.Payload.Utxos[1].TxID.Hex())
	require.Equal(t, mockUtxos[0].Vout, unpackedParams.Payload.Utxos[0].Vout)
	require.Equal(t, mockUtxos[1].Vout, unpackedParams.Payload.Utxos[1].Vout)
	require.Equal(t, mockUtxos[0].AmountInSats, unpackedParams.Payload.Utxos[0].AmountInSats)
	require.Equal(t, mockUtxos[1].AmountInSats, unpackedParams.Payload.Utxos[1].AmountInSats)
	require.Equal(t, mockCmdID.Bytes(), unpackedParams.Payload.RequestId)
	require.Equal(t, mockSymbol, unpackedParams.Symbol)
	require.Equal(t, mokRequestAmount.Uint64(), unpackedParams.Amount)
	require.Equal(t, strings.TrimPrefix(mockCustodianGrUID.Hex(), "0x"), hex.EncodeToString(unpackedParams.CustodianGroupUID[:]))
	require.Equal(t, uint64(1), unpackedParams.SessionSequence)
}
