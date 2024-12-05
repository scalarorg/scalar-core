package types

import fmt "fmt"

type ChainId = uint64

const (
	MainnetChainId ChainId = iota
	TestnetChainId
	LitecoinChainId
	DogecoinChainId
)

func ValidateChainId(c ChainId) error {
	if c < 1 {
		return fmt.Errorf("chain id must be greater than 0")
	}
	return nil
}
