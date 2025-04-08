package keeper_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestQueryCustodians(t *testing.T) {
	suite.Run(t, new(suite.Suite))
}

func TestCreateExecuteDataAndSigs(t *testing.T) {
	commandId := []byte{234, 203, 18, 30, 233, 172, 50, 70, 96, 231, 253, 11, 18, 61, 68, 104, 118, 249, 64, 175, 101, 134, 129, 212, 253, 200, 19, 189, 216, 204, 63, 83}
	commandHex := hex.EncodeToString(commandId)
	require.Equal(t, commandHex, "eacb121ee9ac324660e7fd0b123d446876f940af658681d4fdc813bdd8cc3f53")
	fmt.Println(commandHex)

	// key := []byte("key")
	// cmdData := []byte("cmdData")
	// signature := []byte("test")

	// executeData, _, err := keeper.CreateExecuteDataAndSigs(key, cmdData, signature)
	// require.NoError(t, err)
	// data, err := evm.AbiUnpack(executeData, "bytes32", "bytes")
	// require.NoError(t, err)
	// fmt.Println(data)

}
