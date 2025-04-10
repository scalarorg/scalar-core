package covenant

import (
	"bytes"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/scalarorg/scalar-core/x/covenant/types/mock"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"github.com/stretchr/testify/assert"
)

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestAggregatePsbtFromCommandBatch$ github.com/scalarorg/scalar-core/x/covenant -v -count=1
func TestAggregatePsbtFromCommandBatch(t *testing.T) {
	// Setup test data
	ctx := sdk.Context{}
	chainName := nexus.ChainName("bitcoin")

	var (
		mockScalarnetKeeper *mock.ScalarnetKeeperMock
		mockBaseKeeper      *mock.BaseKeeperMock
	)

	// Create test group
	group := &exported.CustodianGroup{
		BitcoinPubkey: []byte(bytes.Repeat([]byte{0x2}, 32)),
		Custodians: []*exported.Custodian{
			{BitcoinPubkey: []byte(bytes.Repeat([]byte{0x2}, 32))},
			{BitcoinPubkey: []byte(bytes.Repeat([]byte{0x3}, 32))},
		},
		Quorum: 2,
	}

	amount := uint64(100000)
	txID := chains.Hash(bytes.Repeat([]byte{0x0}, 32))
	vout := uint32(0)
	amountInSat := uint64(100000)

	params := types.RedeemTokenParams{
		DestinationChain:   "test_chain",
		DestinationAddress: "test_address",
		RawPayload:         []byte("test_payload"),
		Symbol:             "test_symbol",
		Amount:             amount,
		CustodianGroupUID:  [32]byte{},
		SessionSequence:    1,
		Payload: types.RedeemTokenPayload{
			Amount:        amount,
			LockingScript: []byte("test_script"),
			Utxos:         []*types.UTXO{{TxID: txID, Vout: vout, AmountInSats: amountInSat}},
			RequestId:     chains.ZeroHash,
		},
	}

	payload, err := params.Payload.AbiPack()

	md := chainsTypes.CommandBatchMetadata{
		ExtraData: [][]byte{payload},
	}

	setter := func(m chainsTypes.CommandBatchMetadata) {}
	commandBatch := chainsTypes.NewCommandBatch(md, setter)

	// Execute test
	psbt, err := aggregatePsbtFromCommandBatch(
		ctx,
		mockScalarnetKeeper,
		mockBaseKeeper,
		chainName,
		commandBatch,
		group,
	)

	t.Logf("psbt: %v", psbt)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, psbt)
}
