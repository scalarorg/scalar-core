package testnet

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/cometbft/cometbft/privval"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/rs/zerolog/log"
	cmtjson "github.com/tendermint/tendermint/libs/json"
)

const (
	mnemonic  = "scan move load math desert bleak correct bread jacket tide salt depend dance program good country service ostrich urban punch digital provide file surface"
	keyName   = "test"
	bip44Path = "m/44'/118'/0'/0/0"
)

var (
	kb   keyring.Keyring
	algo keyring.SignatureAlgo
)

func TestMain(m *testing.M) {
	initArgs := initArgs{
		keyringBackend: "test",
		algo:           "secp256k1",
	}
	var err error
	kb, algo, err = createKeyring(bufio.NewReader(os.Stdin), initArgs, ".scalar")
	if err != nil {
		log.Error().Err(err).Msg("Create keyring")
	}

	os.Exit(m.Run())
}
func Test_createKeyringAccountFromMnemonic(t *testing.T) {
	pubkey, address, _ := createKeyringAccountFromMnemonic(kb,
		keyName,
		mnemonic,
		bip44Path,
		algo,
	)
	pkType := pubkey.Type()
	fmt.Println("PK Type:", pkType)
	fmt.Println("Address:", address)
	//Load secp256k1 key into ed25519
	keyJSONBytes, err := os.ReadFile(".scalar/keyring-test/validator.info")
	if err != nil {
		fmt.Println("Error reading key:", err)
		return
	}
	pvKey := privval.FilePVKey{}
	err = cmtjson.Unmarshal(keyJSONBytes, &pvKey)
	if err != nil {
		fmt.Println("Error Unmarshal key:", err)
		return
	}
	newType := pvKey.PubKey.Type()
	fmt.Println("PK type:", newType)
}
