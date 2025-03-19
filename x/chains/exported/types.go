package exported

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Hash common.Hash

var ZeroHash = Hash{}

func (h Hash) IsZero() bool {
	return bytes.Equal(h.Bytes(), ZeroHash.Bytes())
}

func (h Hash) Bytes() []byte {
	return common.Hash(h).Bytes()
}

func (h Hash) Marshal() ([]byte, error) {
	return h[:], nil
}

// MarshalTo implements codec.ProtoMarshaler
func (h Hash) MarshalTo(data []byte) (n int, err error) {
	bytesCopied := copy(data, h[:])
	if bytesCopied != common.HashLength {
		return 0, fmt.Errorf("expected data size to be %d, actual %d", common.HashLength, len(data))
	}

	return common.HashLength, nil
}

func (h *Hash) Unmarshal(data []byte) error {
	if len(data) != common.HashLength {
		return fmt.Errorf("expected data size to be %d, actual %d", common.HashLength, len(data))
	}

	*h = Hash(common.BytesToHash(data))

	return nil
}

func (h Hash) Hex() string {
	return common.Hash(h).Hex()
}

func (h Hash) Size() int {
	return common.HashLength
}

func HashFromHex(hex string) (Hash, error) {
	if len(hex) != common.HashLength*2 {
		return Hash{}, fmt.Errorf("invalid hash length")
	}
	return Hash(common.HexToHash(hex)), nil
}
