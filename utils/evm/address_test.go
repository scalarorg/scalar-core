package evm_test

import (
	"fmt"
	"testing"

	"github.com/scalarorg/scalar-core/utils/evm"
	"gotest.tools/assert"
)

func TestNormalizeAddress(t *testing.T) {
	address := "def3456789412345678941234567894123456abc"
	normalized1, err := evm.NormalizeAddress(address)
	assert.NilError(t, err)
	fmt.Println(normalized1)
	address = "0xDef3456789412345678941234567894123456ABC"
	normalized2, err := evm.NormalizeAddress(address)
	assert.NilError(t, err)
	fmt.Println(normalized2)
	assert.Equal(t, normalized1, normalized2)
}
