package covenant

import (
	"encoding/base64"
	"encoding/binary"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scalarorg/scalar-core/testutils/fake"
	"github.com/scalarorg/scalar-core/testutils/rand"
	"github.com/scalarorg/scalar-core/utils/slices"
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
	mockChains "github.com/scalarorg/scalar-core/x/chains/types/mock"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/scalarorg/scalar-core/x/covenant/types/mock"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	scalarnetTypes "github.com/scalarorg/scalar-core/x/scalarnet/types"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// CGO_LDFLAGS="-L./lib -lbitcoin_vault_ffi" CGO_CFLAGS="-I./lib" go test -timeout 10m -run ^TestAggregatePsbtFromCommandBatch$ github.com/scalarorg/scalar-core/x/covenant -v -count=1
func TestAggregatePsbtFromCommandBatch(t *testing.T) {
	// Setup
	ctx, chainName := setupTestContext()
	mockKeepers := setupMockKeepers()
	group := createTestCustodianGroup()

	// Create test command batch with multiple redeem payloads
	commandBatch := createTestCommandBatch([]*redeemParams{
		{amount: 7000, utxoAmounts: []uint64{3000, 9000}},
		{amount: 8000, utxoAmounts: []uint64{9000, 10000}},
	})

	// Execute test
	psbt, err := aggregatePsbtFromCommandBatch(
		ctx,
		mockKeepers.keeper,
		mockKeepers.scalarnet,
		mockKeepers.base,
		chainName,
		commandBatch,
		group,
		0,
	)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, psbt)
	t.Log("Psbtbase64: ", base64.StdEncoding.EncodeToString(psbt))

}

type mockKeepers struct {
	scalarnet *mock.ScalarnetKeeperMock
	base      *mock.BaseKeeperMock
	keeper    *mock.KeeperMock
}

type redeemParams struct {
	amount      uint64
	utxoAmounts []uint64
}

func setupTestContext() (sdk.Context, nexus.ChainName) {
	ctx := sdk.NewContext(
		fake.NewMultiStore(),
		tmproto.Header{},
		false,
		log.TestingLogger(),
	).WithBlockHeight(rand.I64Between(10, 100))

	return ctx, nexus.ChainName("bitcoin")
}

func setupMockKeepers() mockKeepers {
	scalarnetKeeper := &mock.ScalarnetKeeperMock{
		GetParamsFunc: func(ctx sdk.Context) scalarnetTypes.Params {
			return scalarnetTypes.Params{
				Version: 1,
				Tag:     []byte("SCALAR"),
			}
		},
	}

	keeper := &mock.KeeperMock{}

	baseKeeper := &mock.BaseKeeperMock{
		ForChainFunc: func(ctx sdk.Context, chain nexus.ChainName) (chainsTypes.ChainKeeper, error) {
			return &mockChains.ChainKeeperMock{
				GetParamsFunc: func(ctx sdk.Context) chainsTypes.Params {
					return chainsTypes.Params{}
				},
			}, nil
		},
	}

	return mockKeepers{
		scalarnet: scalarnetKeeper,
		base:      baseKeeper,
		keeper:    keeper,
	}
}

func createTestCustodianGroup() *exported.CustodianGroup {
	return &exported.CustodianGroup{
		BitcoinPubkey: rand.PublicKey(),
		Custodians: slices.Map([]uint8{1, 2, 3, 4}, func(t uint8) *exported.Custodian {
			return &exported.Custodian{
				BitcoinPubkey: rand.PublicKey(),
			}
		}),
		Quorum: 3,
	}
}

func createTestCommandBatch(params []*redeemParams) chainsTypes.CommandBatch {
	var payloads [][]byte

	for _, param := range params {
		payload := createRedeemPayload(param.amount, param.utxoAmounts)
		encodedPayload, _ := payload.AbiPack()
		payloads = append(payloads, encodedPayload)
	}

	md := chainsTypes.CommandBatchMetadata{
		ExtraData: payloads,
	}

	return chainsTypes.NewCommandBatch(md, func(m chainsTypes.CommandBatchMetadata) {})
}

func createRedeemPayload(amount uint64, utxoAmounts []uint64) types.RedeemTokenPayload {
	return types.RedeemTokenPayload{
		Amount:        amount,
		LockingScript: []byte("test_script"),
		Utxos:         createTestUTXOs(utxoAmounts),
		RequestId:     chains.ZeroHash,
	}
}

func createTestUTXOs(amounts []uint64) []*types.UTXO {
	return slices.Map(amounts, func(amount uint64) *types.UTXO {
		var buf = make([]byte, 32)
		binary.LittleEndian.PutUint64(buf[:], amount)
		return &types.UTXO{
			TxID:         chains.Hash(crypto.Keccak256Hash(buf)),
			Vout:         0,
			AmountInSats: amount,
		}
	})
}
