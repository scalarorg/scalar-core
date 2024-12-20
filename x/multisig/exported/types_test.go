package exported_test

import (
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/stretchr/testify/assert"

	"github.com/scalarorg/scalar-core/testutils/rand"
	"github.com/scalarorg/scalar-core/utils/funcs"
	. "github.com/scalarorg/scalar-core/utils/test"
	"github.com/scalarorg/scalar-core/x/multisig/exported"
	typestestutils "github.com/scalarorg/scalar-core/x/multisig/types/testutils"
)

func TestPublicKey(t *testing.T) {
	var (
		pubKey exported.PublicKey
	)

	t.Run("ValidateBasic", func(t *testing.T) {
		Given("valid public key", func() {
			pubKey = typestestutils.PublicKey()
		}).
			When("", func() {}).
			Then("should return nil", func(t *testing.T) {
				assert.NoError(t, pubKey.ValidateBasic())
			}).
			Run(t, 5)

		Given("invalid public key", func() {
			pubKey = rand.Bytes(int(rand.I64Between(1, 101)))
		}).
			When("", func() {}).
			Then("should return error", func(t *testing.T) {
				assert.Error(t, pubKey.ValidateBasic())
			}).
			Run(t, 5)

		Given("uncompressed public key", func() {
			pubKey = funcs.Must(btcec.NewPrivateKey()).PubKey().SerializeUncompressed()
		}).
			When("", func() {}).
			Then("should return error", func(t *testing.T) {
				assert.Error(t, pubKey.ValidateBasic())
			}).
			Run(t, 5)
	})
}
