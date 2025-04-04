package keeper_test

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/scalarorg/scalar-core/x/covenant/keeper"
	types "github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/stretchr/testify/require"
)

func TestCreateAbiRedeemTokenParams(t *testing.T) {
	cmdId := types.NewCommandID(bytes.Repeat([]byte("a"), 32))

	params, err := keeper.CreateAbiRedeemTokenParams(
		&types.ReserveRedeemUtxoRequest{
			Amount: 100000000,
		},
		cmdId.Bytes(),
		[]string{"txId1", "txId2"},
		[]uint32{1, 2},
		[]uint64{100000000, 200000000},
	)
	require.NoError(t, err)
	require.NotNil(t, params)
	t.Log(params)
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
