package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/scalarorg/scalar-core/x/chains/types"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"

	"github.com/cosmos/cosmos-sdk/store"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

func TestGetTokenAddress(t *testing.T) {
	details := nexus.TokenDetails{
		TokenName: "Scalar Pool",
		Symbol:    "sBtc",
		Decimals:  8,
		Capacity:  sdk.NewUint(1000000000000000000),
	}
	gatewayAddr := types.Address(common.HexToAddress("0x0000000000000000000000000000000000000000"))
	bytecode := []byte("token")
	tokenAddr, err := createTokenAddress(details, gatewayAddr, bytecode)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Token address: %s", tokenAddr)
}

func createTokenAddress(details nexus.TokenDetails, gatewayAddr types.Address, bytecode []byte) (types.Address, error) {
	var saltToken [32]byte
	copy(saltToken[:], crypto.Keccak256Hash([]byte(details.Symbol)).Bytes())

	uint8Type, err := abi.NewType("uint8", "uint8", nil)
	if err != nil {
		return types.Address{}, err
	}

	uint256Type, err := abi.NewType("uint256", "uint256", nil)
	if err != nil {
		return types.Address{}, err
	}

	stringType, err := abi.NewType("string", "string", nil)
	if err != nil {
		return types.Address{}, err
	}

	arguments := abi.Arguments{{Type: stringType}, {Type: stringType}, {Type: uint8Type}, {Type: uint256Type}}
	packed, err := arguments.Pack(details.TokenName, details.Symbol, details.Decimals, details.Capacity.BigInt())
	if err != nil {
		return types.Address{}, err
	}

	tokenInitCode := append(bytecode, packed...)
	tokenInitCodeHash := crypto.Keccak256Hash(tokenInitCode)

	tokenAddr := types.Address(crypto.CreateAddress2(common.Address(gatewayAddr), saltToken, tokenInitCodeHash.Bytes()))
	return tokenAddr, nil
}

func NewTestContextWithRocksDB(t *testing.T) sdk.Context {
	// Use a temp dir for test DB
	dir := t.TempDir()
	db, err := dbm.NewDB("testdb", dbm.RocksDBBackend, dir)
	if err != nil {
		t.Fatalf("failed to create rocksdb: %v", err)
	}

	storeKey := sdk.NewKVStoreKey("testStoreKey")
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(storeKey, sdk.StoreTypeDB, db)
	if err := ms.LoadLatestVersion(); err != nil {
		t.Fatalf("failed to load latest version: %v", err)
	}

	ctx := sdk.NewContext(ms, tmproto.Header{}, false, nil)
	return ctx
}
