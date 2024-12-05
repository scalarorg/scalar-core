package types

import (
	"bytes"
	"encoding/hex"
	fmt "fmt"

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

type Hash chainhash.Hash

var ZeroHash = Hash{}

func (h Hash) Bytes() []byte {
	return h[:]
}

func (h Hash) IsZero() bool {
	return bytes.Equal(h.Bytes(), ZeroHash.Bytes())
}

func (h Hash) Size() int {
	return chainhash.HashSize
}

func (h Hash) Marshal() ([]byte, error) {
	return h[:], nil
}

func (h Hash) MarshalTo(data []byte) (int, error) {
	bytesCopied := copy(data, h[:])
	if bytesCopied != chainhash.HashSize {
		return 0, fmt.Errorf("failed to marshal TxHash: expected %d bytes, got %d", chainhash.HashSize, bytesCopied)
	}
	return bytesCopied, nil
}

func (h *Hash) Unmarshal(data []byte) error {
	hash, err := chainhash.NewHash(data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal TxHash: %w", err)
	}
	*h = Hash(*hash)
	return nil
}

func HashFromHexStr(hexStr string) (*Hash, error) {
	hash, err := chainhash.NewHashFromStr(hexStr)
	if err != nil {
		return nil, err
	}
	return (*Hash)(hash), nil
}

func (h Hash) HexStr() string {
	return (*chainhash.Hash)(&h).String()
}

func (nk NetworkKind) Validate() error {
	if nk != Mainnet && nk != Testnet {
		return fmt.Errorf("invalid network kind: %d", nk)
	}
	return nil
}
