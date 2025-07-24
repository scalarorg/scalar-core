package types_test

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	fmt "fmt"
	"testing"

	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/stretchr/testify/require"
)

const (
	QUORUM      = 3
	VSIZE_LIMIT = 200000
)

func TestPayloadDecode(t *testing.T) {
	payloadHex1 := "0000000000000000000000000000000000000000000000000000000000000003e800000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001c0000000000000000000000000000000000000000000000000000000000000020026acc6a0a92e0b94afd58716442f3fdf1a10ee0c1b69131637a6709db3dd33c30000000000000000000000000000000000000000000000000000000000000016001463dc22751d9a7778aa4450ceeb0b5c3ee214401c0000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004230783733373763626133333663666636653064333934373837323234366339396664383039353035336631386461623239616137633663303632336431376330353200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000003e8"
	payloadHex2 := "0000000000000000000000000000000000000000000000000000000000000003e800000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000002004443abfc175541b06e1d9f9410e5fbf89fdb00e5d7ab1a32c830357250e39abf0000000000000000000000000000000000000000000000000000000000000016001463dc22751d9a7778aa4450ceeb0b5c3ee214401c0000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004230783733373763626133333663666636653064333934373837323234366339396664383039353035336631386461623239616137633663303632336431376330353200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000003e8"
	payload1Bytes, err := hex.DecodeString(payloadHex1)
	require.NoError(t, err)
	payload2Bytes, err := hex.DecodeString(payloadHex2)
	require.NoError(t, err)
	payload1 := types.RedeemTokenPayloadWithType{}
	payload2 := types.RedeemTokenPayloadWithType{}
	err = payload1.AbiUnpack(payload1Bytes)
	require.NoError(t, err)
	err = payload2.AbiUnpack(payload2Bytes)
	require.NoError(t, err)
	for _, utxo := range payload1.Utxos {
		t.Logf("utxo txid: %s, vout: %d, amount: %d", utxo.TxID.Hex(), utxo.Vout, utxo.AmountInSats)
	}
	for _, utxo := range payload2.Utxos {
		t.Logf("utxo txid: %s, vout: %d, amount: %d", utxo.TxID.Hex(), utxo.Vout, utxo.AmountInSats)
	}
	t.Logf("payload2: %+v", payload2.Utxos)
}

func TestAbiPack(t *testing.T) {

	//payload1 := "0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 3 232 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 192 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 192 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 2 0 38 172 198 160 169 46 11 148 175 213 135 22 68 47 63 223 26 16 238 12 27 105 19 22 55 166 112 157 179 221 51 195 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 22 0 20 99 220 34 117 29 154 119 120 170 68 80 206 235 11 92 62 226 20 64 28 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 32 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 66 48 120 55 51 55 55 99 98 97 51 51 54 99 102 102 54 101 48 100 51 57 52 55 56 55 50 50 52 54 99 57 57 102 100 56 48 57 53 48 53 51 102 49 56 100 97 98 50 57 97 97 55 99 54 99 48 54 50 51 100 49 55 99 48 53 50 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 3 232"
	//payload2 := "0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 3 232 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 192 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 192 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 2 0 68 67 171 252 23 85 65 176 110 29 159 148 16 229 251 248 159 219 0 229 215 171 26 50 200 48 53 114 80 227 154 191 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 22 0 20 99 220 34 117 29 154 119 120 170 68 80 206 235 11 92 62 226 20 64 28 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 32 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 66 48 120 55 51 55 55 99 98 97 51 51 54 99 102 102 54 101 48 100 51 57 52 55 56 55 50 50 52 54 99 57 57 102 100 56 48 57 53 48 53 51 102 49 56 100 97 98 50 57 97 97 55 99 54 99 48 54 50 51 100 49 55 99 48 53 50 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 3 232"
	//payload1Bytes := strings.ReplaceAll(payload1, " ", ",")
	//payload2Bytes := strings.ReplaceAll(payload2, " ", ",")
	payload1Bytes := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 232, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 192, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 192, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 38, 172, 198, 160, 169, 46, 11, 148, 175, 213, 135, 22, 68, 47, 63, 223, 26, 16, 238, 12, 27, 105, 19, 22, 55, 166, 112, 157, 179, 221, 51, 195, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 22, 0, 20, 99, 220, 34, 117, 29, 154, 119, 120, 170, 68, 80, 206, 235, 11, 92, 62, 226, 20, 64, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 66, 48, 120, 55, 51, 55, 55, 99, 98, 97, 51, 51, 54, 99, 102, 102, 54, 101, 48, 100, 51, 57, 52, 55, 56, 55, 50, 50, 52, 54, 99, 57, 57, 102, 100, 56, 48, 57, 53, 48, 53, 51, 102, 49, 56, 100, 97, 98, 50, 57, 97, 97, 55, 99, 54, 99, 48, 54, 50, 51, 100, 49, 55, 99, 48, 53, 50, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 232}
	payload2Bytes := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 232, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 192, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 192, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 68, 67, 171, 252, 23, 85, 65, 176, 110, 29, 159, 148, 16, 229, 251, 248, 159, 219, 0, 229, 215, 171, 26, 50, 200, 48, 53, 114, 80, 227, 154, 191, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 22, 0, 20, 99, 220, 34, 117, 29, 154, 119, 120, 170, 68, 80, 206, 235, 11, 92, 62, 226, 20, 64, 28, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 66, 48, 120, 55, 51, 55, 55, 99, 98, 97, 51, 51, 54, 99, 102, 102, 54, 101, 48, 100, 51, 57, 52, 55, 56, 55, 50, 50, 52, 54, 99, 57, 57, 102, 100, 56, 48, 57, 53, 48, 53, 51, 102, 49, 56, 100, 97, 98, 50, 57, 97, 97, 55, 99, 54, 99, 48, 54, 50, 51, 100, 49, 55, 99, 48, 53, 50, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 232}
	params1 := types.RedeemCustodianPayload{}
	err := params1.AbiUnpack(payload1Bytes[1:])
	require.NoError(t, err)
	t.Logf("payloadBytes: %x", params1)
	payload1WithType := types.RedeemTokenPayloadWithType{
		Type:                   types.RedeemTypeFromBytes(payload1Bytes[0]),
		RedeemCustodianPayload: &params1,
	}
	bytes, err := payload1WithType.AbiPack()
	require.NoError(t, err)
	t.Logf("bytes: %x", bytes)

	var payload2WithType types.RedeemTokenPayloadWithType
	err = payload2WithType.AbiUnpack(bytes)
	require.NoError(t, err)
	fmt.Printf("payloadBytes: %v\n", payload2WithType)

	payload1WithTypeBytes, err := payload1WithType.AbiPack()
	require.NoError(t, err)
	fmt.Printf("payloadBytes: %x\n", payload1WithTypeBytes)
	fmt.Println("--------------------------------------------")
	params2 := types.RedeemCustodianPayload{}
	err = params2.AbiUnpack(payload2Bytes[1:])
	require.NoError(t, err)
	t.Logf("payloadBytes: %x", params2)

}
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
	utxos, err := snapshot.ReserveUtxos("request1", 150, QUORUM, VSIZE_LIMIT)
	require.NoError(t, err)
	require.Equal(t, 2, len(utxos))
	require.Equal(t, uint64(0), snapshot.Utxos[0].AvailableAmount())
	require.Equal(t, uint64(50), snapshot.Utxos[1].GetReservedAmount())
	utxos, err = snapshot.ReserveUtxos("request2", 50, QUORUM, VSIZE_LIMIT)
	require.NoError(t, err)
	require.Equal(t, 1, len(utxos))
	require.Equal(t, uint64(0), snapshot.Utxos[0].AvailableAmount())
	require.Equal(t, uint64(100), snapshot.Utxos[1].GetReservedAmount())
}

var (
	txIds = []string{
		"DtE38OXXPtSeDJX1McDOvqlpA7QY0owukq/+JHp/yGo=",
		"A4MifD/2ZQ8+XoDE4bQmDalvLHafHFLL0eN3ClSLWWs=",
		"IMnmI7MO8y8O4pGDW6ast/nOIOTCy7302oQE1wv06dI=",
		"bXDcUWrCenIHuwVfsuFn//cbAfJKXJRbvsf9qgpGajU=",
		"caVZhW0jFrEeTt1WnC1lXyhFuZJNjsMQQ4q1RELcVAM=",
		"c/ce/peWHROEs/db8C0j7py5JlSKLQj8uBbs30ZK+0Q=",
		"dmSDyPUN35mfAwvy8lMm92R9VDJmvpXsQzPNHXMbZTw=",
		"d0vF6FeHj0HFwE9m8lDiyFx+FHuiP7Yvt70WRBNDNYI=",
		"fUPKsCOnEuR9QC/RLBnM2yPAWSDZNBMFPxvsF5yI2pU=",
		"f9YhleGkBdsSkEoLADWJe1l6CqZqdG4RCNhMnhbDsUA=",
		"gHVuE72yFcYoSeModt2kh3IWW3SbN2sLKZrT6zlfh9A=",
		"h3foaQSifWhcE0CutgfanfqffB3al9LR3LCKwqkNB8U=",
		"iZT1OkAe+qr02sUgYgfaluSAmuTBm5vUMxjRU1rUZ6U=",
		"ipQf4KfGa2blrVaz5OCKcT3p5GLRgdM0CX/PcUKQNqw=",
		"ip0yZcVLq7WWyf2CM3ce9Qr2nIhbeAXxvU0+nSRCkiA=",
		"isTIh4+UWo8vbRAD6GhBzJLUIu0dUCEpGyKWMbLvGaQ=",
		"kPpncUSzoSW7F38IIxghxOOi+ztXW/g+DaGSWaC2/Ck=",
		"kufqqZneDQVxkTk9sgHmjofqPx2ioj91NM3x1RzhgAk=",
		"ly44A+v0XzVGCRnb4N9Dlva0eeB5VnGkh9nCQVLamcY=",
		"l/EYur4aD9/zRKFaq0DMTs/MVHGNBwmG7LNQfZ/cYbo=",
		"mB3YienkRi/BrYyg7aOE+mKQSzd0/jnoplsRzK6m204=",
		"oWkL7mdwn9Heh9wJJM/G/WCAlTTqnm1bQiAieKHu/wI=",
		"pS247KHP+3Q4VUvR5w+J6zSh47+Wq6zxWv5A9zefqlY=",
		"qCGPplDmlnmiHEIWkuYVEOpiGz+j/Oj3QzFA9bKQIYs=",
		"qTn9kgMW11T5E8Hhj7j4xpuFQss6JbunZIvL6TaRJhc=",
		"qidWE+BanBpCBFjaRfdOBc6YAaPea6g3CE2C5ToGjsU=",
		"q/lJivYxhAksYfMOv3vYd8U0H81ia+76tKvnL13O4JE=",
		"rkI4yOLAYleyqOiheW+d0clT0yEToXN1i5ADr2JZ96A=",
		"sA0Zm45jDTtbADka2/CX3e5q0fZ8QGVIqXAlkT0dCSc=",
		"tMAOjB+IOZF0zgqJLD2zi/oLSb/aw4oVe/T0b7p0qek=",
		"uISNFqPVBbTbBD5KYrAaZGl0n1W8JpiNCGQxe+IcfTM=",
		"wRFrgUzWmZVglFgXcppmWw9WhedxuQYTy7YqpXXIVLQ=",
		"wxvFIsAnqO/vF4lscuvJ1lQpRLzpbKWwpiHy729T/wk=",
		"w53EgwaTFuezQrtF3hHMX8lK/lmC4md8y+WlHowBoK8=",
		"xGOH1mRdhLMALmBY9GmnMi/XQDMQnoU3J6bTVCa/2ho=",
		"x5BSBa3S7JJPcmM4PDJeFlBXVQbMjYEyvPNPzgHNVUg=",
		"yDGUc0tUaPfc2tiE5Kc1eT7Qgw+d26rfjPWcmp+LX7Q=",
		"zAM0lel3qIhjX4CMAvSV3sAj+fvIEKJpgA+2i0cpoHo=",
		"zGntlm7WtT3fHEqTXZjFc8odX5wvMBm//M0ckjneVPQ=",
		"zUgQslTPKCBJ6WNJt5QSQZIU6GvjY+5NZxUIMvoHLTc=",
		"zc1KAug3L26J06m2mvTiXwe3RkQm3RNP15T2CYw8dWc=",
		"zhebeNR7PVIzKKO5gKNDpWC9NPLwFn0bvZh18FSMqs8=",
		"0H7gvdCvK1ReTIXb9CjL5hChGXQt1lRUxF4E80Niuwk=",
		"02aP8qWsopMUyaP/8rLvji1SVawikqC22bPahWgg1jc=",
		"1jaY8qHFlY5zFghVaxdJa0D+pZuNQKc9yJriDiYGZtI=",
		"1r0HGcB11kiOPZiAZHO0CWkGRpfNttKuMHNbnjACdFY=",
		"2QrY3QVQ+9Fb+2t7ccAbVB56J9N6uJvWvQFXFkwXmMk=",
		"2bnHwvW6pwT2o2RXZebQ637uafnlL9MBJSBaJqZcnOc=",
		"2rM1eUc4nJOxcubVafhlgAqVnkrx5j0yvObPB2yfxZ0=",
		"3Ul1d5rLBFDikq3ojCEEcMiKtKONglTRx53G9ioMzkM=",
		"3r/Jj3/e4i+6fTfJVzv4XLtaxeKgWZPjFAVaQSOm2/o=",
		"3xF5AI03TgkljLMFjo/33PmGN6xqakPnXZi+GcM1rX8=",
		"332Av0pagSuG4QoSDX6QIAWYg8sDgVJwmXJvLzKoBls=",
		"4HSdmRppNXQz9hLU8J1/MLRR0sXMJRzTfnSwvjPzQXM=",
		"5izDdl7ERcMFmlQqGosSE/M7G4ki7UTt5Vle1d3OAsU=",
		"6DTfQIUVmsPlhVtC1FAkDZVbkd1zti+Wn0ROGUU3h/w=",
		"64nnFw7EB4vAKttVXBX0A2DPLVYN60o9E52C4COxkBQ=",
		"7EEBvse987eTmZfrStxWxocOpFzNB1tbJVsCtFFe9ng=",
		"80/qEBy07cBsjZob7qcJKm/QBxFaZ55nxHDQdah1Lj4=",
		"9Yd0/lmpcPPqqb/0gppNhLjQHD66hPyiyjeQHtQ2Lng=",
		"9ltbitShzcY3Xs349rpTRwdtrOBNvMDo9Bqot7zqiGc=",
		"9x40Olv/pOQQfQBXkR5GqLJO2JHLTfIPToeswHKMKaw=",
		"/lSvsziPodPD5fEX2ugMSOjmUg2xWGDOPVTUh4dEDu0=",
		"/0qs3S3kVeDLt9DSUYuhiEcYuxIwpyU8wnJnY/us8RQ=",
		"uWADWJFFomq5IM6+62TDnOqWEtYbqAMjN/BoEcbiXeI=",
	}
)

type PrintUTXO struct {
	// Reserved amount for each request id
	Reservations []*types.Reservation `json:"reservations,omitempty"`
	TxID         string               `json:"txid"`
	Vout         uint32               `json:"vout"`
	AmountInSats uint64               `protobuf:"varint,4,opt,name=amount_in_sats,json=amountInSats,proto3" json:"amount_in_sats,omitempty"`
}

func UTXOToPrint(utxo *types.UTXO) *PrintUTXO {
	return &PrintUTXO{
		Reservations: utxo.Reservations,
		TxID:         HashToBase64(utxo.TxID),
		Vout:         utxo.Vout,
		AmountInSats: utxo.AmountInSats,
	}
}
func TestRealUtxoReserve(t *testing.T) {
	snapshot := types.UTXOSnapshot{}
	for ind, txId := range txIds {
		vout := uint32(1)
		amount := uint64(9741)
		if ind == 0 {
			vout = 8
			amount = 4355
		}
		snapshot.Utxos = append(snapshot.Utxos, &types.UTXO{
			TxID:         Base64ToHash(txId),
			Vout:         vout,
			AmountInSats: amount,
		})
	}
	reserveAmount := uint64(4741)
	reserveUtxos(t, &snapshot, "86e0db317fd058c58aa30c3dc09671346987b45db0c141396264201ccbdfe1d6", reserveAmount)
	reserveUtxos(t, &snapshot, "5828158eaa7e8a6f63f0dade346e264967cddd82c19dcece6242406d849ff7e3", reserveAmount)
	reserveUtxos(t, &snapshot, "e9adbdd25fc0ffa03035335e593f24f92472122da31db29190042013a4bfc3cb", reserveAmount)
	reserveUtxos(t, &snapshot, "8ce3828a18bf953c797d3849d94e4c8a2088c42007792cfff32a060d63d822d1", reserveAmount)
	reserveUtxos(t, &snapshot, "241fece8a2c9fdbbd0cf378d528ca51d9e3a1ffecc0b00fa70afdad2487c6f9e", reserveAmount)
	reserveUtxos(t, &snapshot, "73a1fcb1a953232993d500c750cbbf1ec71e819c601a62cc2aebe130d191f74e", reserveAmount)
	reserveUtxos(t, &snapshot, "f28c898e621ab2c3168851332ac6b743159008902b2a599a602ddb80f8ea0517", reserveAmount)
	reserveUtxos(t, &snapshot, "7bd903a369d0e4374e3c0b8464369d1450ac0cc43fe90a71d2ec70223fa517cd", reserveAmount)
	reserveUtxos(t, &snapshot, "8b091183327d8fe1961022c81ad13043512ed1808171ca5d9615eed77e991a1c", reserveAmount)
	reserveUtxos(t, &snapshot, "01e0490fbcf3aa1c2b8dd7d928492f2d7b656b451b4c8cf2c07e69fde2155574", reserveAmount)
	reserveUtxos(t, &snapshot, "5775baa235a4c3af20b3fbe9f7d2e22630d0b2b03a1c3cac426583df86b8d98c", reserveAmount)
	reserveUtxos(t, &snapshot, "6d333e7c216fae72a67ae4ad6ff07be9fdbd551b920e777ed2ef5c667e0940e3", reserveAmount)
	reserveUtxos(t, &snapshot, "5775baa235a4c3af20b3fbe9f7d2e22630d0b2b03a1c3cac426583df86b8d98c", reserveAmount)
	printUtxos := make([]*PrintUTXO, 0)
	for _, utxo := range snapshot.Utxos {
		printUtxos = append(printUtxos, UTXOToPrint(utxo))
	}
	bytes, err := json.MarshalIndent(printUtxos, "", "\t")
	require.NoError(t, err)
	t.Logf("Snapshot: %s", string(bytes))
}

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestUPCUnpack$ github.com/scalarorg/scalar-core/x/covenant/types -v -count=1
func TestUPCUnpack(t *testing.T) {
	data := "01370736274ff0100960200000001de476807425a9fda442f63c53130bf62773dd93dadba25840a5dfdced4c918160100000000fdffffff030000000000000000106a0e5343414c41520301817472616e737a8501000000000016001474e29d4022d44de324a6f5b95ec2fa46a3ba27c8e069f902000000002251206317353427ad7d2e0cee59e825fdd1598e70c08e8a0146f48ea9db206ad2fec3000000000001012b80f0fa02000000002251206317353427ad7d2e0cee59e825fdd1598e70c08e8a0146f48ea9db206ad2fec30103040000000041146f63a8d030a1857a7d054fbbfd06dddc30285154d785a977b0c33445ee3262bfac7a87f9f99577615b5ae427531c5f2290e7bd27bd120eabae4e162501d994d04076d5a7a02e851d0667e9bc01ed512241f67d36126d63ebdf298e6c34a4484dd59a016f98afd97ff26594f66ffeeb99a0e57bc33b30a5e23e7bd1129ec364423c6215c150929b74c1a04954b78b4b6035e97a5e078a5a0f28ec96d547bfee9ace803ac063d85693fdcf7e1d56c8240e96cbd468d4c580a29a6540a5305aa14fd1a3785904887e344db1e070ba2141b0240b484249216fb09b56f02bfcfd98e2582e4cbdad206f63a8d030a1857a7d054fbbfd06dddc30285154d785a977b0c33445ee3262bfad2015da913b3e87b4932b1e1b87d9667c28e7250aa0ed60b3a31095f541e1641488ac20594e78c0a2968210d9c1550d4ad31b03d5e4b9659cf2f67842483bb3c2bb7811ba20e2d226cfdaec93903c3f3b81a01a81b19137627cb26e621a0afb7bcd6efbcfffba20f0f3d9beaf7a3945bcaa147e041ae1d5ca029bde7e40d8251f0783d6ecbe8fb5ba53a2c0211615da913b3e87b4932b1e1b87d9667c28e7250aa0ed60b3a31095f541e16414882501ac7a87f9f99577615b5ae427531c5f2290e7bd27bd120eabae4e162501d994d0000000002116594e78c0a2968210d9c1550d4ad31b03d5e4b9659cf2f67842483bb3c2bb78112501ac7a87f9f99577615b5ae427531c5f2290e7bd27bd120eabae4e162501d994d00000000021166f63a8d030a1857a7d054fbbfd06dddc30285154d785a977b0c33445ee3262bf2501ac7a87f9f99577615b5ae427531c5f2290e7bd27bd120eabae4e162501d994d0000000002116e2d226cfdaec93903c3f3b81a01a81b19137627cb26e621a0afb7bcd6efbcfff2501ac7a87f9f99577615b5ae427531c5f2290e7bd27bd120eabae4e162501d994d0000000002116f0f3d9beaf7a3945bcaa147e041ae1d5ca029bde7e40d8251f0783d6ecbe8fb52501ac7a87f9f99577615b5ae427531c5f2290e7bd27bd120eabae4e162501d994d000000000011820801c615d9bb268053f1c0bde9e56f1506853321f90a012f511084d13fea58c6800000000"

	hex, _ := hex.DecodeString(data)

	var p types.RedeemTokenPayloadWithType
	err := p.AbiUnpack(hex)
	if err != nil {
		t.Fatalf("AbiUnpack failed, error: %s", err)
	}

	fmt.Printf("p: %x\n", p.RedeemUPCPayload.Psbt)
}

func reserveUtxos(t *testing.T, snapshot *types.UTXOSnapshot, requestId string, amount uint64) []*types.UTXO {
	utxos, err := snapshot.ReserveUtxos(requestId, amount, QUORUM, VSIZE_LIMIT)
	require.NoError(t, err)
	return utxos
}
func Base64ToHash(input string) chains.Hash {
	hashBytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		panic(err)
	}
	var hash [32]byte
	copy(hash[:], hashBytes)
	return chains.Hash(hash)
}

func HashToBase64(hash chains.Hash) string {
	return base64.StdEncoding.EncodeToString(hash[:])
}
