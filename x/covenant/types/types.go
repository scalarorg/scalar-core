package types

import (
	fmt "fmt"

	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
)

func DefaultRedeemSession() *RedeemSession {
	return &RedeemSession{
		CustodianGroupUid: "",
		Sequence:          0,
		CurrentPhase:      Unspecified,
		LastRedeemTx:      nil,
	}
}

func NewRedeemSession(custodianGroupUid string, sequence uint64, currentPhase Phase, lastRedeemTx *chains.Hash) *RedeemSession {
	return &RedeemSession{
		CustodianGroupUid: custodianGroupUid,
		Sequence:          sequence,
		CurrentPhase:      currentPhase,
		LastRedeemTx:      lastRedeemTx,
	}
}

var EmptyRedeemSession = DefaultRedeemSession()

func (rs RedeemSession) IsEmpty() bool {
	if rs.CustodianGroupUid == "" || rs.Sequence == 0 {
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
	RequestID string  `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Amount    uint64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Utxos     []*UTXO `protobuf:"bytes,3,rep,name=utxos,proto3" json:"utxos,omitempty"`
}

func (utxo *UTXO) AppendReserved(requestID string, amount uint64) {
	utxo.Reserved[requestID] += amount
}
func (utxo *UTXO) GetReservedAmount() uint64 {
	amount := uint64(0)
	for _, reserved := range utxo.Reserved {
		amount += reserved
	}
	return amount
}

func (utxo *UTXO) AvailableAmount() uint64 {
	return utxo.AmountInSats - utxo.GetReservedAmount()
}

func (utxo *UTXO) Release(requestID string) uint64 {
	amount, ok := utxo.Reserved[requestID]
	if !ok {
		return 0
	}
	delete(utxo.Reserved, requestID)
	return amount
}

func (rs UTXOSnapshot) ReleaseUtxos(requestID string) uint64 {
	releasedAmount := uint64(0)
	for _, utxo := range rs.Utxos {
		releasedAmount += utxo.Release(requestID)
	}
	return releasedAmount
}

// Utxos list is sorted by amount in sats and txid for deterministic results
// Each utxo is reserved if it is part of the optimal combination
func (rs UTXOSnapshot) ReserveUtxos(requestID string, amount uint64) ([]*UTXO, error) {
	remainingAmount := amount
	reserveUtxos := make([]*UTXO, 0)
	for _, utxo := range rs.Utxos {
		availableAmount := utxo.AvailableAmount()
		if availableAmount > 0 {
			//Reserve amount is min(availableAmount, remainingAmount)
			reserveAmount := availableAmount
			if reserveAmount > remainingAmount {
				reserveAmount = remainingAmount
			}
			reserveUtxos = append(reserveUtxos, utxo)
			utxo.AppendReserved(requestID, reserveAmount)
			remainingAmount -= reserveAmount
		}
		if remainingAmount == 0 {
			break
		}
	}

	return reserveUtxos, nil
}

func (m RedeemTxConfirmed) ValidateBasic() error {
	if !chainsTypes.IsBitcoinChain(m.Chain) {
		return fmt.Errorf("invalid chain")
	}
	return nil
}
