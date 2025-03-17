package types_test

import (
	"testing"

	covenant "github.com/scalarorg/scalar-core/x/covenant/exported"
	"github.com/scalarorg/scalar-core/x/covenant/types"
	"github.com/stretchr/testify/assert"
)

func TestDefaultGenesisState(t *testing.T) {
	assert.NoError(t, types.NewGenesisState(types.DefaultParams(), []types.SigningSession{}, []*covenant.Custodian{covenant.DefaultCustodian()}, []*covenant.CustodianGroup{covenant.DefaultCustodianGroup()}).Validate())
}
