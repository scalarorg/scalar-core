package types_test

import (
	"testing"

	"github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/stretchr/testify/assert"
)

func TestDefaultGenesisState(t *testing.T) {
	assert.NoError(t, types.NewGenesisState(types.DefaultParams(), []types.SigningSession{}, []*types.Custodian{types.DefaultCustodian()}, []*types.CustodianGroup{types.DefaultCustodianGroup()}).Validate())
}
