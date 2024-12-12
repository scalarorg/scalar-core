package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/scalarorg/scalar-core/x/btc/keeper"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/axelarnetwork/axelar-core/testutils/fake"
	"github.com/scalarorg/scalar-core/x/btc/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
)

func TestAddressValidator(t *testing.T) {
	ctx := sdk.NewContext(fake.NewMultiStore(), tmproto.Header{}, false, log.TestingLogger())
	validator := keeper.NewAddressValidator()
	assert.NotNil(t, validator)

	addr := nexus.CrossChainAddress{Chain: exported.Bitcoin, Address: "16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM"}
	assert.NoError(t, validator(ctx, addr))
}
