package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultGenesisState(t *testing.T) {
	assert.NoError(t, NewGenesisState([]*Custodian{DefaultCustodian()}, DefaultCustodianGroup()).Validate())
}
