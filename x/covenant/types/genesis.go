package types

import (
	cov "github.com/scalarorg/scalar-core/x/covenant/exported"
)

func NewGenesisState(params *Params, signingSessions []SigningSession, custodians []*cov.Custodian, groups []*cov.CustodianGroup) GenesisState {
	return GenesisState{
		Params:          *params,
		Custodians:      custodians,
		Groups:          groups,
		SigningSessions: signingSessions,
	}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	custodians := []*cov.Custodian{cov.DefaultCustodian()}
	group := cov.DefaultCustodianGroup()
	params := DefaultParams()
	session := []SigningSession{}
	return NewGenesisState(params, session, custodians, []*cov.CustodianGroup{group})
}
func (m GenesisState) Validate() error {
	return nil
}
