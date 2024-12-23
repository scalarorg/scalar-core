package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/scalarorg/scalar-core/x/tss/types"
)

func TestDefaultParams(t *testing.T) {
	params := types.DefaultParams()

	assert.NoError(t, params.Validate())
}
