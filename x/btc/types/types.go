package types

import "github.com/btcsuite/btcd/chaincfg/chainhash"

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

type VaultVersion [1]byte

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

func (h TxHash) Size() int {
	return chainhash.HashSize
}

func (h TxHash) MarshalTo(data []byte) (int, error) {
	return copy(data, h[:]), nil
}

func (h TxHash) Unmarshal(data []byte) error {
	copy(h[:], data)
	return nil
}
