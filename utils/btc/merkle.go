package btc

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/slices"
	"github.com/scalarorg/scalar-core/x/chains/exported"
)

func DoubleSha256(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

func ReverseBytes(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return b
}

func CalculateMerkleRoot(hashes [][]byte) []byte {
	if len(hashes) == 0 {
		return nil
	}
	if len(hashes) == 1 {
		return hashes[0]
	}
	for len(hashes) > 1 {
		var nextLevel [][]byte
		for i := 0; i < len(hashes); i += 2 {
			left := hashes[i]
			var right []byte
			if i+1 < len(hashes) {
				right = hashes[i+1]
			} else {
				right = hashes[i]
			}
			nextLevel = append(nextLevel, DoubleSha256(append(left, right...)))
		}
		hashes = nextLevel
	}
	return hashes[0]
}

func GetMerkleRootFromPath(txid []byte, txIndex uint64, path [][]byte, reverse bool) []byte {
	for _, p := range path {
		if txIndex%2 == 0 {
			txid = DoubleSha256(append(txid, p...))
		} else {
			txid = DoubleSha256(append(p, txid...))
		}
		txIndex /= 2
	}
	if reverse {
		return ReverseBytes(txid)
	}
	return txid
}

func ValidateTxProof(txId []byte, txIndex uint64, merklePath []exported.Hash, blockMerkleRoot exported.Hash) error {
	merkleRoot := GetMerkleRootFromPath(txId, txIndex, slices.Map(merklePath, func(p exported.Hash) []byte {
		return p.Bytes()
	}), true)

	if !bytes.Equal(merkleRoot, blockMerkleRoot.Bytes()) {
		clog.Redf("txId: %x, txIndex: %d, merkleRoot: %x, blockMerkleRoot: %x",
			txId, txIndex, merkleRoot, blockMerkleRoot.Bytes())

		for i, p := range merklePath {
			clog.Redf("merklePath[%d]: %x", i, p.Bytes())
		}

		return fmt.Errorf("merkle root mismatch: %x != %x", merkleRoot, blockMerkleRoot.Bytes())
	}

	return nil
}
