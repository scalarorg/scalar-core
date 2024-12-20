package types

import (
	"github.com/google/uuid"
)

func NewGenesisState(custodians []*Custodian, group *CustodianGroup) GenesisState {
	return GenesisState{
		Custodians: custodians,
		Groups:     []*CustodianGroup{group},
	}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	custodians := []*Custodian{DefaultCustodian()}
	group := DefaultCustodianGroup()
	return NewGenesisState(custodians, group)
}
func (m GenesisState) Validate() error {
	return nil
}

func DefaultCustodian() *Custodian {
	custodian := &Custodian{
		Name:   DefaultCustodianName,
		Status: Activated,
	}
	return custodian
}

func DefaultCustodianGroup() *CustodianGroup {
	return &CustodianGroup{
		Uid:  uuid.NewString(),
		Name: DefaultCustodianName,
	}
}
