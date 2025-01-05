package types_test

import (
	"encoding/hex"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scalarorg/scalar-core/utils"
	"github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/stretchr/testify/assert"
)

var (
	details = types.TokenDetails{
		TokenName: "SepoliaBtc",
		Symbol:    "pBtc",
		Decimals:  8,
		Capacity:  sdk.NewInt(100000000),
	}
	gatewayAddr    = common.HexToAddress("0xBdce70581713ceaecb63b9dE1Fd6188FD80654E8")
	tokenDeployer  = common.HexToAddress("0x230C52C6a30c3CC63fcBa2CBE01C9B07a44cF812")
	expectedSalt   = "2377374275e65139f8b13b342d1fb862fc39ca23e76356fe24f7ca72c5ba5757"
	salt           = crypto.Keccak256Hash([]byte(details.Symbol)).Bytes()
	uint8Type, _   = abi.NewType("uint8", "uint8", nil)
	uint256Type, _ = abi.NewType("uint256", "uint256", nil)
	stringType, _  = abi.NewType("string", "string", nil)
	bytesType, _   = abi.NewType("bytes", "bytes", nil)
)

func TestContractAddress(t *testing.T) {
	t.Logf("Detail %v", details)
	callerAddr := gatewayAddr
	expextedAddress := "0x01bA44229f9909f86D347584bad11c9a2136dE6d"
	var saltToken [32]byte
	copy(saltToken[:], salt)
	assert.Equal(t, expectedSalt, hex.EncodeToString(saltToken[:]))

	arguments := abi.Arguments{{Type: stringType}, {Type: stringType}, {Type: uint8Type}, {Type: uint256Type}}
	packed, err := arguments.Pack(details.TokenName, details.Symbol, details.Decimals, details.Capacity.BigInt())
	assert.NoError(t, err)
	bytecode, err := utils.HexDecode(types.Token)

	assert.NoError(t, err)
	tokenInitCode := append(bytecode, packed...)
	tokenInitCodeHash := crypto.Keccak256Hash(tokenInitCode)

	tokenAddr := types.Address(crypto.CreateAddress2(callerAddr, saltToken, tokenInitCodeHash.Bytes()))
	t.Logf("Generated token address %s", tokenAddr.String())
	assert.Equal(t, expextedAddress, tokenAddr.String())
}
