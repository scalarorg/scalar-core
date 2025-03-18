package types

import "github.com/scalarorg/scalar-core/x/chains/types"

func DefaultRedeemSession() RedeemSession {
	return RedeemSession{
		Symbol:                "",
		Sequence:              0,
		CurrentPhase:          Unspecified,
		LastRedeemTx:          nil,
		Utxos:                 make([]*UTXO, 0),
		RequestedAmountInSats: make(map[string]uint64),
	}
}

func NewRedeemSession(symbol string, sequence uint64, currentPhase Phase, lastRedeemTx *types.Hash, utxos []*UTXO, requestedAmountInSats map[string]uint64) RedeemSession {
	return RedeemSession{
		Symbol:                symbol,
		Sequence:              sequence,
		CurrentPhase:          currentPhase,
		LastRedeemTx:          lastRedeemTx,
		Utxos:                 utxos,
		RequestedAmountInSats: requestedAmountInSats,
	}
}

var EmptyRedeemSession = DefaultRedeemSession()

func (rs RedeemSession) IsEmpty() bool {
	if rs.Symbol == "" || rs.Sequence == 0 {
		return true
	}
	return false
}
