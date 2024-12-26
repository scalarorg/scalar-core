package rpc

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/scalar-core/client/rpc/cosmos"
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
	log.Debug().Msgf("Created account with address: %s from mnemonic: %s", addr.String(), mnemonic)
	return privKey, addr, nil
}

func ConfirmSourceTx(ctx context.Context, client *cosmos.NetworkClient, msg *chainTypes.ConfirmSourceTxsRequest) (*types.TxResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}
	return client.SignAndBroadcastMsgs(ctx, msg)
}
