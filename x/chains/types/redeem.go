package types

import fmt "fmt"

// Cmp compares two redeem sessions
// Returns -1 if m is less than other, 0 if they are equal, 1 if m is greater than other
func (m *RedeemSession) ToPhaseNumber() int64 {
	return int64(m.Sequence)*2 + int64(m.CurrentPhase)
}

func (m *RedeemSession) Cmp(other *RedeemSession) int64 {
	val := m.ToPhaseNumber()
	if other == nil {
		return val
	}
	return val - other.ToPhaseNumber()
}

func (m *RedeemSession) ToString() string {
	return fmt.Sprintf("custodian group uid: %s, sequence: %d, phase: %v", m.CustodianGroupUID.Hex(), m.Sequence, m.CurrentPhase)
}
