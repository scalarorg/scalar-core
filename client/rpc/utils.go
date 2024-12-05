package rpc

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
	"github.com/scalarorg/relayers/pkg/clients/cosmos"
	btcTypes "github.com/scalarorg/scalar-core/x/btc/types"
)

const (
	AccountAddressPrefix   = "scalar"
	ValidatorAddressPrefix = AccountAddressPrefix + types.PrefixValidator + types.PrefixOperator
)

func CreateAccountFromKey(key string) (*secp256k1.PrivKey, types.AccAddress, error) {
	privKeyBytes, err := hex.DecodeString(key)
	if err != nil {
		return nil, nil, err
	}
	privKey := &secp256k1.PrivKey{Key: privKeyBytes}
	config := types.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, ValidatorAddressPrefix)
	addr := types.AccAddress(privKey.PubKey().Address())
	log.Debug().Msgf("Created account with address: %s from key: %s", addr.String(), key)
	return privKey, addr, nil
}

func ConfirmBtcTx(ctx context.Context, client *cosmos.NetworkClient, msg *btcTypes.ConfirmStakingTxsRequest) (*types.TxResponse, error) {
	if client == nil {
		return nil, fmt.Errorf("client is not initialized")
	}
	return client.SignAndBroadcastMsgs(ctx, msg)
}
