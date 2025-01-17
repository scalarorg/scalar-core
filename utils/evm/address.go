package evm

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func NormalizeAddress(address string) (string, error) {
	address = strings.TrimPrefix(address, "0x")
	if len(address) != 40 {
		return "", errors.New("invalid address length")
	}
	if !IsHexAddress(address) {
		return "", errors.New("not a valid hex address")
	}
	return common.HexToAddress(address).Hex(), nil
}

func IsHexAddress(address string) bool {
	for _, c := range address {
		if !isHexCharacter(byte(c)) {
			return false
		}
	}
	return true
}

func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}
