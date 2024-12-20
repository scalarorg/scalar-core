package types

type ChainID = uint64

const (
	BTCMainnetChainID ChainID = iota
	BTCTestnetChainID
	BTCLitecoinChainID
	BTCDogecoinChainID
)

func ValidateChainID(c ChainID) error {
	return nil
}
