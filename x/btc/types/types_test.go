package types_test

import (
	"testing"

	"github.com/scalarorg/scalar-core/x/btc/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultVaultTag(t *testing.T) {
	tag := types.TagFromAscii("SCALAR")
	require.Equal(t, tag.HexStr(), "5343414c4152")
	t.Log(tag)
}

func TestHashFromHexStr(t *testing.T) {
	hash, err := types.HashFromHexStr("07b50c84f889e2f1315da875fc91734e2bac8d0153ff9a98d9da14caa4fc7d57")
	require.NoError(t, err)
	require.Equal(t, hash.HexStr(), "07b50c84f889e2f1315da875fc91734e2bac8d0153ff9a98d9da14caa4fc7d57")
	t.Log(hash)
}
