package types

import (
	"bytes"
	"encoding/hex"
	fmt "fmt"
	"slices"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

type VaultTag [6]byte

func (VaultTag) Size() int {
	return 6
}

func (v VaultTag) MarshalTo(data []byte) (int, error) {
	return copy(data, v[:]), nil
}

func (v VaultTag) Unmarshal(data []byte) error {
	copy(v[:], data)
	return nil
}

func TagFromAscii(str string) VaultTag {
	var tag VaultTag
	copy(tag[:], []byte(str))
	return tag
}

func (v VaultTag) HexStr() string {
	return hex.EncodeToString(v[:])
}

type VaultVersion [1]byte

func VersionFromInt(version int) VaultVersion {
	var v VaultVersion
	v[0] = byte(version)
	return v
}

func (VaultVersion) Size() int {
	return 1
}

func (v VaultVersion) MarshalTo(data []byte) (int, error) {
	return copy(data, v[:]), nil
}

func (v VaultVersion) Unmarshal(data []byte) error {
	copy(v[:], data)
	return nil
}

type TxHash chainhash.Hash

var ZeroTxHash = TxHash{}

func (h TxHash) Bytes() []byte {
	return h[:]
}

func (h TxHash) IsZero() bool {
	return bytes.Equal(h.Bytes(), ZeroTxHash.Bytes())
}

func (h TxHash) Size() int {
	return chainhash.HashSize
}

func (h TxHash) Marshal() ([]byte, error) {
	return h[:], nil
}

func (h TxHash) MarshalTo(data []byte) (int, error) {
	bytesCopied := copy(data, h[:])
	if bytesCopied != chainhash.HashSize {
		return 0, fmt.Errorf("failed to marshal TxHash: expected %d bytes, got %d", chainhash.HashSize, bytesCopied)
	}
	return bytesCopied, nil
}

func (h *TxHash) Unmarshal(data []byte) error {
	hash, err := chainhash.NewHash(data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal TxHash: %w", err)
	}
	*h = TxHash(*hash)
	return nil
}

func TxHashFromHexStr(hexStr string) (*TxHash, error) {
	hexBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}

	if len(hexBytes) != chainhash.HashSize {
		return nil, fmt.Errorf("invalid hex string length: %d", len(hexBytes))
	}

	littleEndianBytes := slices.Clone(hexBytes)
	slices.Reverse(littleEndianBytes)

	hash, err := chainhash.NewHash(littleEndianBytes)
	if err != nil {
		return nil, err
	}
	return (*TxHash)(hash), nil
}

func (h TxHash) HexStr() string {
	return (*chainhash.Hash)(&h).String()
}
