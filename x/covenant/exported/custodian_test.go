package exported_test

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	"github.com/scalarorg/scalar-core/x/covenant/exported"
	"golang.org/x/crypto/sha3"
	"gotest.tools/assert"
)

func TestCalculateUID(t *testing.T) {
	uid := exported.CalculateUID("scalarv32")
	base64UID := base64.StdEncoding.EncodeToString(uid[:])
	assert.Equal(t, "PnkyapSTiW4Tr2IZTmlN/0yTAHAEB0STY1ZLDq6vB+g=", base64UID)
	assert.Equal(t, "3e79326a9493896e13af62194e694dff4c9300700407449363564b0eaeaf07e8", hex.EncodeToString(uid[:]))
	assert.Equal(t, uid, sha3.Sum256([]byte("scalarv32")))
}
