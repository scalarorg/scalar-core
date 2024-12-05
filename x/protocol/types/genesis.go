package types

func NewGenesisState(protocol Protocol) *GenesisState {
	return &GenesisState{
		Protocols: []Protocol{protocol},
	}
}

func (m GenesisState) Validate() error {
	return nil
}
