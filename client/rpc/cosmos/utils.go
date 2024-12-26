package cosmos

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	chainTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

const (
	AccountAddressPrefix   = "scalar"
	ValidatorAddressPrefix = AccountAddressPrefix + types.PrefixValidator + types.PrefixOperator
)

func CreateAccountFromMnemonic(mnemonic string, bip44Path string) (*secp256k1.PrivKey, types.AccAddress, error) {
	// Derive the seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "")
	path := "m/44'/118'/0'/0/0"
	if bip44Path != "" {
		path = bip44Path
	}
	// Create master key and derive the private key
	// Using "m/44'/118'/0'/0/0" for Cosmos
	master, ch := hd.ComputeMastersFromSeed(seed)
	privKeyBytes, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		return nil, nil, err
	}

	// Create private key and get address
	privKey := &secp256k1.PrivKey{Key: privKeyBytes}
	addr := types.AccAddress(privKey.PubKey().Address())
	return privKey, addr, nil
}

func ConfirmSourceTx(ctx context.Context, client *NetworkClient, msg *chainTypes.ConfirmSourceTxsRequest) (*types.TxResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}
	return client.SignAndBroadcastMsgs(ctx, msg)
}

func StringArrayToBytes(s string) ([]byte, error) {
	// Remove brackets and split by comma
	s = strings.Trim(s, "[]")
	numberStrings := strings.Split(s, ",")

	bytes := make([]byte, len(numberStrings))
	for i, numStr := range numberStrings {
		// Convert string to uint64
		val, err := strconv.ParseUint(strings.TrimSpace(numStr), 10, 8)
		if err != nil {
			return nil, fmt.Errorf("failed to parse number at index %d: %w", i, err)
		}
		bytes[i] = byte(val)
	}

	return bytes, nil
}
