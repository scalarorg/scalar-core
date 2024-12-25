package testnet_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeGateway(t *testing.T) {
	gateway := "842C080EE1399addb76830CFe21D41e47aaaf57e"
	gwBytes, err := hex.DecodeString(gateway)
	assert.NoError(t, err)
	fmt.Printf("%v", gwBytes)
}
