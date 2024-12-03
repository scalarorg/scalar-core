package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func DefaultCovenant() Covenant {
	return Covenant{
		Name: "test",
	}
}

func DefaultCovenantGroup() CovenantGroup {
	return CovenantGroup{
		Name: "test",
	}
}

func TestDefaultGenesisState(t *testing.T) {
	assert.NoError(t, NewGenesisState([]Covenant{DefaultCovenant()}, DefaultCovenantGroup()).Validate())
}
