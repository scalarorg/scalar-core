package types

func NewGenesisState(covenants []Covenant, group CovenantGroup) *GenesisState {
	return &GenesisState{
		Covenants: covenants,
		Groups:    []CovenantGroup{group},
	}
}

func (m GenesisState) Validate() error {
	return nil
}
