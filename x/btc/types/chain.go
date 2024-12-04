package types

type ChainId = uint64

const (
	MainnetChainId ChainId = iota
	TestnetChainId
	LitecoinChainId
	DogecoinChainId
)
