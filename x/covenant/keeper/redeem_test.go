package keeper_test

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/scalarorg/bitcoin-vault/go-utils/encode"
	"github.com/scalarorg/scalar-core/x/covenant/keeper"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/stretchr/testify/require"
)

func TestCreateAbiRedeemTokenParams(t *testing.T) {
	mockCmdID := types.NewCommandID(bytes.Repeat([]byte("a"), 32))
	mockAddress := "0x0000000000000000000000000000000000000000"
	mockChainName := nexus.ChainName("bitcoin|4")
	mokRequestAmount := big.NewInt(100_000)
	mockSymbol := "BTC"
	mockLockingScript := []byte("locking-script")
	mockUtxos := struct {
		amountsInSats []uint64
		txIds         []string
		vouts         []uint32
	}{
		amountsInSats: []uint64{50_000, 70_000},
		txIds:         []string{"txId1", "txId2"},
		vouts:         []uint32{1, 2},
	}

	params, err := keeper.CreateAbiRedeemTokenParams(
		&types.ReserveRedeemUtxoRequest{
			Amount:        mokRequestAmount.Uint64(),
			DestChain:     mockChainName,
			Address:       mockAddress,
			Symbol:        mockSymbol,
			LockingScript: mockLockingScript,
		},
		mockCmdID.Bytes(),
		mockUtxos.txIds,
		mockUtxos.vouts,
		mockUtxos.amountsInSats,
	)
	require.NoError(t, err)

	unpacked, err := keeper.CallContractWithTokenArguments.Unpack(params)

	require.NoError(t, err)
	require.NotNil(t, unpacked)
	t.Log(unpacked)
	t.Logf("hex: %x", unpacked)

	destChain := unpacked[0].(string)
	require.Equal(t, mockChainName.String(), destChain)
	address := unpacked[1].(string)
	require.Equal(t, mockAddress, address)
	payload := unpacked[2].([]byte)
	symbol := unpacked[3].(string)
	require.Equal(t, mockSymbol, symbol)
	amount := unpacked[4].(*big.Int)
	require.Equal(t, mokRequestAmount, amount)

	t.Logf("payload: %x", payload)

	// remove prefix
	payload, prefix := payload[1:], payload[0:1]
	require.Equal(t, encode.ContractCallWithTokenPayloadType_CustodianOnly.Bytes(), prefix)

	// decode payload
	unpacked, err = keeper.RedeemTokenPayloadArguments.Unpack(payload)
	require.NoError(t, err)
	require.NotNil(t, unpacked)

	_amount := unpacked[0].(uint64)
	require.Equal(t, mokRequestAmount.Uint64(), _amount)
	lockingScript := unpacked[1].([]byte)
	require.Equal(t, mockLockingScript, lockingScript)
	_txIds := unpacked[2].([]string)
	require.Equal(t, mockUtxos.txIds, _txIds)
	_vouts := unpacked[3].([]uint32)
	require.Equal(t, mockUtxos.vouts, _vouts)
	_amounts := unpacked[4].([]uint64)
	require.Equal(t, mockUtxos.amountsInSats, _amounts)
	_reqId := unpacked[5].([32]byte)
	t.Logf("_reqId: %x", _reqId)

	if bytes.Compare(_reqId[:], mockCmdID[:]) != 0 {
		t.Errorf("expected %x, got %x", mockCmdID.Bytes(), _reqId[:])
	}
}

func TestAbiTypes(t *testing.T) {
	uint64Type, err := abi.NewType("uint64", "", nil)
	if err != nil {
		panic(err)
	}
	arg := abi.Argument{Type: uint64Type}
	args := abi.Arguments{arg}
	data, err := args.Pack(uint64(12345))
	require.NoError(t, err)
	if err != nil {
		panic(err)
	}
	t.Log(data)
}
