package types

import (
	fmt "fmt"

	chains "github.com/scalarorg/scalar-core/x/chains/exported"
	"github.com/scalarorg/scalar-core/x/covenant/exported"
	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
	"golang.org/x/crypto/sha3"
)

const (
	EventTypeSwitchPhaseSign   = "switchPhaseSign"
	EventTypeReserveRedeemUtxo = "ReserveRedeemUtxo"
)

func DefaultRedeemSession() *RedeemSession {
	return &RedeemSession{
		CustodianGroupUID: chains.ZeroHash,
		Sequence:          0,
		CurrentPhase:      exported.Executing,
		IsSwitching:       true,
		LastRedeemTx:      nil,
	}
}

func NewRedeemSession(CustodianGroupUID chains.Hash, sequence uint64, currentPhase exported.Phase, isSwitching bool, lastRedeemTx *chains.Hash) *RedeemSession {
	return &RedeemSession{
		CustodianGroupUID: CustodianGroupUID,
		Sequence:          sequence,
		CurrentPhase:      currentPhase,
		IsSwitching:       isSwitching,
		LastRedeemTx:      lastRedeemTx,
	}
}

var EmptyRedeemSession = DefaultRedeemSession()

func (rs RedeemSession) IsEmpty() bool {
	if rs.CustodianGroupUID.IsZero() || rs.Sequence == 0 {
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

func (utxo *UTXO) AppendReserved(requestID string, amount uint64) error {
	if utxo.Reserved == nil {
		utxo.Reserved = make(map[string]uint64)
	}
	totalReserved := uint64(0)
	for id, reserved := range utxo.Reserved {
		if id == requestID {
			return fmt.Errorf("requestID already reserved in this utxo %s", utxo.TxID.Hex())
		}
		totalReserved += reserved
	}
	if totalReserved+amount > utxo.AmountInSats {
		return fmt.Errorf("amount exceeds utxo amount, totalReserved %d, amount %d, utxo.AmountInSats %d", totalReserved, amount, utxo.AmountInSats)
	}
	utxo.Reserved[requestID] = amount
	return nil
}
func (utxo *UTXO) GetReservedAmount() uint64 {
	amount := uint64(0)
	for _, reserved := range utxo.Reserved {
		amount += reserved
	}
	return amount
}

func (utxo *UTXO) AvailableAmount() uint64 {
	reservedAmount := utxo.GetReservedAmount()
	return utxo.AmountInSats - reservedAmount
}

func (utxo *UTXO) Release(requestID string) uint64 {
	if utxo.Reserved == nil {
		return 0
	}
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
	if remainingAmount > 0 {
		return nil, fmt.Errorf("not enough utxos to reserve, remainingAmount %d", remainingAmount)
	}
	return reserveUtxos, nil
}

func (rs *UTXOSnapshot) GetHash() chains.Hash {
	bytes, err := rs.Marshal()
	if err != nil {
		return chains.ZeroHash
	}
	hash := sha3.Sum256(bytes)
	return chains.Hash(hash[:])
}

func NewVoteEvents(chain nexus.ChainName, events ...Event) *VoteEvents {
	return &VoteEvents{
		Chain:  chain,
		Events: events,
	}
}

func NewSigMetadata(sigType SigType, chain nexus.ChainName, commandID []byte) *SigMetadata {
	return &SigMetadata{
		Type:      sigType,
		Chain:     chain,
		CommandID: commandID,
	}
}
