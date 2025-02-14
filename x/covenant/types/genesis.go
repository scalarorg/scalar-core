package types

import (
	"github.com/google/uuid"
)

func NewGenesisState(params *Params, signingSessions []SigningSession, custodians []*Custodian, groups []*CustodianGroup) GenesisState {
	return GenesisState{
		Params:          *params,
		Custodians:      custodians,
		Groups:          groups,
		SigningSessions: signingSessions,
	}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	custodians := []*Custodian{DefaultCustodian()}
	group := DefaultCustodianGroup()
	params := DefaultParams()
	session := []SigningSession{}
	return NewGenesisState(params, session, custodians, []*CustodianGroup{group})
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
