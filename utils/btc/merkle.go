package btc

import "crypto/sha256"

func doubleSha256(b []byte) []byte {
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

func GetMerkleRootFromPath(txid []byte, txIndex uint64, path [][]byte, reverse bool) []byte {
	for _, p := range path {
		if txIndex%2 == 0 {
			txid = doubleSha256(append(txid, p...))
		} else {
			txid = doubleSha256(append(p, txid...))
		}
		txIndex /= 2
	}
	if reverse {
		return reverseBytes(txid)
	}
	return txid
}
