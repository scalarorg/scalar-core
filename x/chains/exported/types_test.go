package exported_test

import (
	"bytes"
	"testing"

	"github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	bytes := bytes.Repeat([]byte{0x01}, 32)
	hash := exported.Hash(bytes)
	hex := hash.Hex()
	require.Equal(t, "0x0101010101010101010101010101010101010101010101010101010101010101", hex)
	hash2, err := exported.HashFromHex(hex)
	require.NoError(t, err)
	require.Equal(t, hash, hash2)
}
