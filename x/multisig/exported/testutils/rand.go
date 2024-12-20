package testutils

import (
	"github.com/scalarorg/scalar-core/testutils/rand"
	"github.com/scalarorg/scalar-core/x/multisig/exported"
)

// KeyID returns a random key ID
func KeyID() exported.KeyID {
	return exported.KeyID(rand.NormalizedStrBetween(exported.KeyIDLengthMin, exported.KeyIDLengthMax+1))
}
