package testutils

import (
	"github.com/scalarorg/scalar-core/testutils/rand"
	"github.com/scalarorg/scalar-core/utils"
)

// RandThreshold returns a random Threshold
func RandThreshold() utils.Threshold {
	return utils.NewThreshold(rand.I64Between(1, 101), 100)
}
