package types_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/scalarorg/scalar-core/utils/funcs"
	chainsExported "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/stretchr/testify/require"
)

func TestCreateSwitchPhasePayload(t *testing.T) {
	bytes32Type := funcs.Must(abi.NewType("bytes32", "bytes32", nil))
	uint8Type := funcs.Must(abi.NewType("uint8", "uint8", nil))

	switchPhaseArguments := abi.Arguments{{Type: uint8Type}, {Type: bytes32Type}}

	gr := bytes.Repeat([]byte{0x01}, 32)

	var gr32 [32]byte
	copy(gr32[:], gr)

	fmt.Println(hex.EncodeToString(gr))

	payload := funcs.Must(switchPhaseArguments.Pack(uint8(exported.Preparing), gr32))
	// return payload

	fmt.Println(hex.EncodeToString(payload))
}

func TestCreateSwitchPhasePayload2(t *testing.T) {
	gr := bytes.Repeat([]byte{0x01}, 32)
	hash := chainsExported.Hash(gr)
	payload := types.CreateSwitchPhasePayload(hash, exported.Preparing)
	require.Equal(t, "00000000000000000000000000000000000000000000000000000000000000000101010101010101010101010101010101010101010101010101010101010101", hex.EncodeToString(payload))
}

func TestDecodeSwitchPhasePayload(t *testing.T) {
	gr := bytes.Repeat([]byte{0x01}, 32)
	hash := chainsExported.Hash(gr)
	payload := types.CreateSwitchPhasePayload(hash, exported.Preparing)
	phase, group := chainsTypes.DecodeSwitchPhaseParams(payload)
	require.Equal(t, uint8(exported.Preparing), phase)
	fmt.Println(group)
	require.Equal(t, hex.EncodeToString(gr), hex.EncodeToString(group[:]))
}
