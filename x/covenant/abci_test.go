package covenant_test

// import (
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/scalarorg/scalar-core/x/chains/types"
// 	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
// 	"github.com/scalarorg/scalar-core/x/covenant/exported"
// 	"github.com/scalarorg/scalar-core/x/covenant/types/mock"
// 	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
// 	"github.com/stretchr/testify/assert"
// )

// func TestAggregatePsbtFromCommandBatch(t *testing.T) {
// 	// Setup test data
// 	ctx := sdk.Context{}
// 	chainName := nexus.ChainName("bitcoin")

// 	var (
// 		keeper          *mock.KeeperMock
// 		scalarnetKeeper *mock.ScalarnetKeeperMock
// 		chainKeeper     *mock.ChainKeeperMock
// 	)

// 	// Create test group
// 	group := &exported.CustodianGroup{
// 		BitcoinPubkey: []byte("test_pubkey"),
// 		Custodians: []*exported.Custodian{
// 			{BitcoinPubkey: []byte("custodian1_pubkey")},
// 			{BitcoinPubkey: []byte("custodian2_pubkey")},
// 		},
// 		Quorum: 2,
// 	}

// 	// Create test command batch
// 	uint256Type, _ := abi.NewType("uint256", "uint256", nil)
// 	bytesType, _ := abi.NewType("bytes", "bytes", nil)
// 	stringArrayType, _ := abi.NewType("string[]", "string[]", nil)
// 	uint256ArrayType, _ := abi.NewType("uint256[]", "uint256[]", nil)

// 	args := abi.Arguments{
// 		{Type: uint256Type},
// 		{Type: bytesType},
// 		{Type: stringArrayType},
// 		{Type: uint256ArrayType},
// 		{Type: uint256ArrayType},
// 	}

// 	amount := uint64(100000)
// 	txID := "abc123"
// 	vout := uint64(0)
// 	amountInSat := uint64(100000)

// 	// Encode test payload
// 	payload, _ := args.Pack(
// 		&amount,
// 		[]byte("test_script"),
// 		[]string{txID},
// 		[]*uint64{&vout},
// 		[]*uint64{&amountInSat},
// 	)

// 	md := chainsTypes.CommandBatchMetadata{
// 		ExtraData: [][]byte{payload},
// 	}

// 	setter := func(m types.CommandBatchMetadata) {}
// 	batch := types.NewCommandBatch(md, setter)

// 	// Execute test
// 	psbt, err := aggregatePsbtFromCommandBatch(
// 		ctx,
// 		mockScalarnet,
// 		mockBase,
// 		chainName,
// 		commandBatch,
// 		group,
// 	)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.NotNil(t, psbt)

// 	// Verify mock expectations
// 	mockScalarnet.AssertExpectations(t)
// 	mockBase.AssertExpectations(t)
// 	mockChain.AssertExpectations(t)
// }

// func TestAggregatePsbtFromCommandBatch_InvalidPayload(t *testing.T) {
// 	ctx := sdk.Context{}
// 	chainName := nexus.ChainName("bitcoin")

// 	mockScalarnet := new(mockScalarnetKeeper)
// 	mockBase := new(mockBaseKeeper)

// 	group := &exported.CustodianGroup{
// 		BitcoinPubkey: []byte("test_pubkey"),
// 		Custodians: []*exported.Custodian{
// 			{BitcoinPubkey: []byte("custodian1_pubkey")},
// 		},
// 		Quorum: 1,
// 	}

// 	// Create command batch with invalid payload
// 	commandBatch := chainsTypes.CommandBatch{
// 		ExtraData: [][]byte{[]byte("invalid_payload")},
// 	}

// 	// Execute test
// 	psbt, err := aggregatePsbtFromCommandBatch(
// 		ctx,
// 		mockScalarnet,
// 		mockBase,
// 		chainName,
// 		commandBatch,
// 		group,
// 	)

// 	// Assertions
// 	assert.Error(t, err)
// 	assert.Nil(t, psbt)
// }
