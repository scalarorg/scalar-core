package types

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/scalarorg/scalar-core/testutils/rand"
	"github.com/scalarorg/scalar-core/utils/funcs"
	"github.com/scalarorg/scalar-core/utils/slices"
	multisigTestutils "github.com/scalarorg/scalar-core/x/multisig/exported/testutils"
)

func TestNewCommandBatchMetadata(t *testing.T) {
	chainID := sdk.NewInt(1)
	commands := []Command{
		{
			ID:     CommandID(common.HexToHash("0xc5baf525fe191e3e9e35c2012ff5f86954c04677a1e4df56079714fc4949409f")),
			Type:   COMMAND_TYPE_DEPLOY_TOKEN,
			Params: common.Hex2Bytes("00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000271000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010416e20417765736f6d6520546f6b656e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000034141540000000000000000000000000000000000000000000000000000000000"),
		},
	}

	expectedData := "0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000001c5baf525fe191e3e9e35c2012ff5f86954c04677a1e4df56079714fc4949409f00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000b6465706c6f79546f6b656e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000271000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010416e20417765736f6d6520546f6b656e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000034141540000000000000000000000000000000000000000000000000000000000"
	actual, err := NewCommandBatchMetadata(
		rand.PosI64(),
		chainID,
		multisigTestutils.KeyID(),
		commands,
	)

	assert.NoError(t, err)
	assert.Equal(t, expectedData, common.Bytes2Hex(actual.Data))
}

func TestDeployToken(t *testing.T) {
	chainID := sdk.NewInt(1)
	keyID := multisigTestutils.KeyID()

	details := TokenDetails{
		TokenName: rand.Str(10),
		Symbol:    rand.Str(3),
		Decimals:  uint8(rand.I64Between(3, 10)),
		Capacity:  sdk.NewIntFromBigInt(big.NewInt(rand.I64Between(100, 100000))),
	}
	address := Address(common.BytesToAddress(rand.Bytes(common.AddressLength)))
	asset := rand.Str(5)

	capBz := make([]byte, 8)
	binary.BigEndian.PutUint64(capBz, details.Capacity.Uint64())
	capHex := hex.EncodeToString(capBz)

	dailyMintLimit := sdk.NewUint(uint64(rand.PosI64()))
	dailyMintLimitHex := hex.EncodeToString(dailyMintLimit.BigInt().Bytes())

	expectedParams := fmt.Sprintf("00000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000%s%s000000000000000000000000%s%s000000000000000000000000000000000000000000000000000000000000000a%s000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003%s0000000000000000000000000000000000000000000000000000000000",
		hex.EncodeToString([]byte{byte(details.Decimals)}),
		strings.Repeat("0", 64-len(capHex))+capHex,
		hex.EncodeToString(address.Bytes()),
		strings.Repeat("0", 64-len(dailyMintLimitHex))+dailyMintLimitHex,
		hex.EncodeToString([]byte(details.TokenName)),
		hex.EncodeToString([]byte(details.Symbol)),
	)
	expectedCommandID := NewCommandID([]byte(asset+"_"+details.Symbol), chainID)
	actual := NewDeployTokenCommand(chainID, keyID, asset, details, address, dailyMintLimit)

	assert.Equal(t, expectedParams, hex.EncodeToString(actual.Params))
	assert.Equal(t, expectedCommandID, actual.ID)

	decodedName, decodedSymbol, decodedDecs, decodedCap, tokenAddress, decodedDailyMintLimit := DecodeDeployTokenParams(actual.Params)
	assert.Equal(t, details.TokenName, decodedName)
	assert.Equal(t, details.Symbol, decodedSymbol)
	assert.Equal(t, details.Decimals, decodedDecs)
	assert.Equal(t, details.Capacity.BigInt(), decodedCap)
	assert.Equal(t, address, Address(tokenAddress))
	assert.Equal(t, decodedDailyMintLimit, dailyMintLimit)
}

func TestGetSignHash(t *testing.T) {
	data := common.FromHex("0000000000000000000000000000000000000000000000000000000000000001ec78d9c22c08bb9f0ecd5d95571ae83e3f22219c5a9278c3270691d50abfd91b000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000096d696e74546f6b656e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000014141540000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000063fc2ad3d021a4d7e64323529a55a9442c444da00000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000270f")

	expected := "0xe7bce8f57491e71212d930096bacf9288c711e5f27200946edd570e3a93546bf"
	actual := GetSignHash(data)

	assert.Equal(t, expected, actual.Hex())
}

func TestCreateExecuteDataMultisig(t *testing.T) {
	commandData := common.FromHex("0000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000157e938a17a25798cd144da54195e9ef765a44ffcf2784c7990b0442c1ca02a230000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000096d696e74546f6b656e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000060000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb922660000000000000000000000000000000000000000000000000000000005f5e10000000000000000000000000000000000000000000000000000000000000000034141540000000000000000000000000000000000000000000000000000000000")
	addresses := []common.Address{
		common.HexToAddress("0x14dC79964da2C08b23698B3D3cc7Ca32193d9955"),
		common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"),
		common.HexToAddress("0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f"),
		common.HexToAddress("0x90F79bf6EB2c4f870365E785982E1f101E93b906"),
		common.HexToAddress("0x976EA74026E726554dB657fA54763abd0C3a0aa9"),
		common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc"),
	}
	weights := slices.Expand(func(idx int) sdk.Uint { return sdk.OneUint() }, len(addresses))
	signatures := [][]byte{
		common.FromHex("0009f7136165f0fc9044f9de3a88920aad0c5844797bd67924ca7bd59865901a38bd628cc2935ca2e83f8d5a97279a1a743cd3ab529e8aeb245bcf676b6491bf1c"),
		common.FromHex("a9704efbd02b99f7c46dfc02437b7143bc61dd6766011d87f7084b43c27d66843105529481d9851713b0a6091305ef84a57017e5772d63a2c97a09d53a91909b1b"),
		common.FromHex("e13769e1716162c3000d09764840ecf4ba5bde9a5ad0fdc7ba0fa6028553a4e1703dbabae51d6360ad7b8ba2ad16215df4291cf3dc6c55e8c85ba234bca6a55a1b"),
	}
	threshold := sdk.NewUint(3)

	expected := "09c5eabe00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000002a000000000000000000000000000000000000000000000000000000000000002400000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000157e938a17a25798cd144da54195e9ef765a44ffcf2784c7990b0442c1ca02a230000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000096d696e74546f6b656e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000060000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb922660000000000000000000000000000000000000000000000000000000005f5e1000000000000000000000000000000000000000000000000000000000000000003414154000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000240000000000000000000000000000000000000000000000000000000000000000600000000000000000000000014dc79964da2c08b23698b3d3cc7ca32193d995500000000000000000000000015d34aaf54267db7d7c367839aaf71a00a2c6a6500000000000000000000000023618e81e3f5cdf7f54c3d65f7fbc0abf5b21e8f00000000000000000000000090f79bf6eb2c4f870365e785982e1f101e93b906000000000000000000000000976ea74026e726554db657fa54763abd0c3a0aa90000000000000000000000009965507d1a55bcc2695c58ba16fb37d819b0a4dc00000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000410009f7136165f0fc9044f9de3a88920aad0c5844797bd67924ca7bd59865901a38bd628cc2935ca2e83f8d5a97279a1a743cd3ab529e8aeb245bcf676b6491bf1c000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041a9704efbd02b99f7c46dfc02437b7143bc61dd6766011d87f7084b43c27d66843105529481d9851713b0a6091305ef84a57017e5772d63a2c97a09d53a91909b1b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041e13769e1716162c3000d09764840ecf4ba5bde9a5ad0fdc7ba0fa6028553a4e1703dbabae51d6360ad7b8ba2ad16215df4291cf3dc6c55e8c85ba234bca6a55a1b00000000000000000000000000000000000000000000000000000000000000"
	actual, err := CreateExecuteDataMultisig(commandData, addresses, weights, threshold, signatures)

	assert.NoError(t, err)
	assert.Equal(t, expected, common.Bytes2Hex(actual))
}

func TestERC20TokenMetadata_ValidateBasic(t *testing.T) {
	t.Run("burner code for internal token is validated", func(t *testing.T) {
		internal := ERC20TokenMetadata{
			Asset:   "asset",
			ChainID: sdk.NewInt(rand.PosI64()),
			Details: TokenDetails{
				TokenName: "token",
				Symbol:    "axl",
				Decimals:  18,
				Capacity:  sdk.NewInt(rand.PosI64()),
			},
			TokenAddress: Address{},
			TxHash:       Hash(common.BytesToHash(rand.Bytes(common.HashLength))),
			Status:       Initialized,
			IsExternal:   false,
			BurnerCode:   funcs.Must(hex.DecodeString(Burnable)),
		}
		assert.NoError(t, internal.ValidateBasic())

		internal.BurnerCode = rand.Bytes(123)
		assert.Error(t, internal.ValidateBasic())
	})

	t.Run("burner code for external token must be nil", func(t *testing.T) {
		external := ERC20TokenMetadata{
			Asset:   "asset",
			ChainID: sdk.NewInt(rand.PosI64()),
			Details: TokenDetails{
				TokenName: "token",
				Symbol:    "axl",
				Decimals:  18,
				Capacity:  sdk.NewInt(rand.PosI64()),
			},
			TokenAddress: Address{},
			TxHash:       Hash(common.BytesToHash(rand.Bytes(common.HashLength))),
			Status:       Initialized,
			IsExternal:   true,
			BurnerCode:   nil,
		}

		assert.NoError(t, external.ValidateBasic())

		external.BurnerCode = funcs.Must(hex.DecodeString(Burnable))
		assert.Error(t, external.ValidateBasic())
	})

}

func TestCommandID_ValidateBasic(t *testing.T) {
	randID := NewCommandID(rand.Bytes(100), sdk.NewIntFromUint64(uint64(rand.PosI64())))
	assert.NoError(t, randID.ValidateBasic())

	var data [100]byte
	idForZeroData := NewCommandID(data[:], sdk.NewIntFromUint64(0))
	assert.NoError(t, idForZeroData.ValidateBasic())

	var emptyID CommandID
	assert.NoError(t, emptyID.ValidateBasic())
}
