package keeper_test

import (
	"bytes"
	"encoding/hex"
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
	_amount, lockingScript, _txIds, _vouts, _amounts, _reqId := decodeAbiRedeemTokenParams(t, payload)

	require.Equal(t, mokRequestAmount.Uint64(), _amount)
	require.Equal(t, mockLockingScript, lockingScript)
	require.Equal(t, mockUtxos.txIds, _txIds)
	require.Equal(t, mockUtxos.vouts, _vouts)
	require.Equal(t, mockUtxos.amountsInSats, _amounts)
	if bytes.Compare(_reqId[:], mockCmdID[:]) != 0 {
		t.Errorf("expected %x, got %x", mockCmdID.Bytes(), _reqId[:])
	}
}

func TestDecodeAbiRedeemTokenParams(t *testing.T) {
	h, _ := hex.DecodeString("00000000000000000000000000000000000000000000000000000000000000128500000000000000000000000000000000000000000000000000000000000000C000000000000000000000000000000000000000000000000000000000000000E0000000000000000000000000000000000000000000000000000000000000042000000000000000000000000000000000000000000000000000000000000004E023C06825C7784642C3DFBCCEE38E715187F54E5362906AE4D78DAF5F0406B6030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000A0000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000001A0000000000000000000000000000000000000000000000000000000000000022000000000000000000000000000000000000000000000000000000000000002A000000000000000000000000000000000000000000000000000000000000000423078373337376362613333366366663665306433393437383732323436633939666438303935303533663138646162323961613763366330363233643137633035320000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000042307833323465346437323638623062363562613666643731333539623634363061383261396661653137616637343334623438616436333265323033313536623735000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004230783764326233326537666438326436633831393662303161323166616266353765353431666637636234623966353062323738353964386337363639623566616100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000423078653439393237666532313066626237333961626266333539343538356534376361346362623830323165656363363732616535303961336434366562396639340000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000042307831326530626662393962306138643530626138313461383464623234653935333530393331316331343139636166646464383631326566633839393462346336000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000003E800000000000000000000000000000000000000000000000000000000000003E80000000000000000000000000000000000000000000000000000000000000458000000000000000000000000000000000000000000000000000000000000053300000000000000000000000000000000000000000000000000000000000003E8")

	amount, lockingScript, txIds, vouts, amounts, reqId := decodeAbiRedeemTokenParams(t, h[1:])
	t.Logf("amount: %d", amount)
	t.Logf("lockingScript: %x", lockingScript)
	t.Logf("txIds: %v", txIds)
	t.Logf("vouts: %v", vouts)
	t.Logf("amounts: %v", amounts)
	t.Logf("reqId: %x", reqId)
}

func decodeAbiRedeemTokenParams(t *testing.T, params []byte) (uint64, []byte, []string, []uint32, []uint64, [32]byte) {

	// decode payload
	unpacked, err := keeper.RedeemTokenPayloadArguments.Unpack(params)
	require.NoError(t, err)
	require.NotNil(t, unpacked)

	_amount := unpacked[0].(uint64)
	lockingScript := unpacked[1].([]byte)
	_txIds := unpacked[2].([]string)
	_vouts := unpacked[3].([]uint32)
	_amounts := unpacked[4].([]uint64)
	_reqId := unpacked[5].([32]byte)
	t.Logf("_reqId: %x", _reqId)
	return _amount, lockingScript, _txIds, _vouts, _amounts, _reqId
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
