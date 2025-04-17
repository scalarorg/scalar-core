package types

// Cmp compares two redeem sessions
// Returns -1 if m is less than other, 0 if they are equal, 1 if m is greater than other
func (m *RedeemSession) Cmp(other *RedeemSession) int64 {
	val := int64(m.Sequence)*2 + int64(m.CurrentPhase)
	otherVal := int64(other.Sequence)*2 + int64(other.CurrentPhase)
	return val - otherVal
}
