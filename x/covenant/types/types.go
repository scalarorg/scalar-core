package types

import (
	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	covExported "github.com/scalarorg/scalar-core/x/covenant/exported"
)

func DefaultRedeemSession() *RedeemSession {
	return &RedeemSession{
		Symbol:       "",
		Sequence:     0,
		CurrentPhase: Unspecified,
		LastRedeemTx: nil,
		Utxos:        make([]*UTXO, 0),
	}
}

func NewRedeemSession(symbol string, sequence uint64, currentPhase Phase, lastRedeemTx *chains.Hash, utxos []*UTXO, requestedAmountInSats map[string]uint64) *RedeemSession {
	return &RedeemSession{
		Symbol:       symbol,
		Sequence:     sequence,
		CurrentPhase: currentPhase,
		LastRedeemTx: lastRedeemTx,
		Utxos:        utxos,
	}
}

var EmptyRedeemSession = DefaultRedeemSession()

func (rs RedeemSession) IsEmpty() bool {
	if rs.Symbol == "" || rs.Sequence == 0 {
		return true
	}
	return false
}

type ReservedUtxo struct {
	TxID         string
	Vout         uint32
	ScriptPubKey []byte
	Amount       uint64
	//Detail reserved amount for each request. This structure is use for release reservation when redeemtx is not broadcasted to the network
	Reserved map[string]uint64
}

type ReservedTx struct {
	RequestID string              `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Amount    uint64              `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Utxos     []*covExported.UTXO `protobuf:"bytes,3,rep,name=utxos,proto3" json:"utxos,omitempty"`
}
