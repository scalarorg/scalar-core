package types

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Chains: []GenesisState_Chain{},
	}
}

func (data GenesisState) Validate() error {
	// Add validation logic
	return nil
}

func NewGenesisState(params Params, tag VaultTag, version VaultVersion) *GenesisState {
	return &GenesisState{
		Chains: []GenesisState_Chain{
			{
				Chain: params.Chain,
			},
		},
		VaultTag:     &tag,
		VaultVersion: &version,
	}
}
