package btc

import (
	"crypto/sha256"
)

func DoubleSha256(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

func reverseBytes(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return b
}

func CalculateMerkleRoot(txHashes [][]byte) []byte {
	if len(txHashes) == 0 {
		return nil
	}
	if len(txHashes) == 1 {
		return txHashes[0]
	}
	// Double SHA256 each transaction hash
	hashes := make([][]byte, len(txHashes))
	for i, txHash := range txHashes {
		hashes[i] = DoubleSha256(txHash)
	}
	for len(hashes) > 1 {
		nextLevel := make([][]byte, 0, (len(hashes)+1)/2)
		for i := 0; i < len(hashes); i += 2 {
			if i+1 < len(hashes) {
				combined := append(hashes[i], hashes[i+1]...)
				nextLevel = append(nextLevel, DoubleSha256(combined))
			} else {
				combined := append(hashes[i], hashes[i]...)
				nextLevel = append(nextLevel, DoubleSha256(combined))
			}
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
		return reverseBytes(txid)
	}
	return txid
}
