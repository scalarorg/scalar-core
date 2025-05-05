package types

import (
	"bytes"
	fmt "fmt"
	"math/big"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/scalarorg/bitcoin-vault/go-utils/encode"
	"github.com/scalarorg/scalar-core/utils"
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

type RedeemTokenPayload struct {
	Amount        uint64
	LockingScript []byte
	Utxos         []*UTXO
	RequestId     [32]byte
}

type RedeemTokenPayloadWithType struct {
	RedeemTokenPayload
	PayloadType encode.ContractCallWithTokenPayloadType
}

func (p *RedeemTokenPayloadWithType) AbiPack() ([]byte, error) {
	payload, err := p.RedeemTokenPayload.AbiPack()
	if err != nil {
		return nil, err
	}
	return encode.AppendPayload(encode.ContractCallWithTokenPayloadType(p.PayloadType), payload), nil
}

func (p *RedeemTokenPayloadWithType) AbiUnpack(data []byte) error {
	p.PayloadType = encode.ContractCallWithTokenPayloadType(data[0])
	var payload RedeemTokenPayload
	err := payload.AbiUnpack(data[1:])
	if err != nil {
		return err
	}
	p.RedeemTokenPayload = payload
	return nil
}

func (p *RedeemTokenPayload) AbiPack() ([]byte, error) {
	txIds := make([]string, len(p.Utxos))
	vouts := make([]uint32, len(p.Utxos))
	amounts := make([]uint64, len(p.Utxos))
	for i, utxo := range p.Utxos {
		txIds[i] = utxo.TxID.Hex()
		vouts[i] = utxo.Vout
		amounts[i] = utxo.AmountInSats
	}
	return RedeemTokenPayloadArguments.Pack(
		p.Amount,
		p.LockingScript,
		txIds,
		vouts,
		amounts,
		p.RequestId,
	)
}

func (p *RedeemTokenPayload) AbiUnpack(data []byte) error {
	unpacked, err := RedeemTokenPayloadArguments.Unpack(data)
	if err != nil {
		log.Error().Err(err).Msg("redeem token payload abi unpack error")
		return err
	}
	p.Amount = unpacked[0].(uint64)
	p.LockingScript = unpacked[1].([]byte)
	txIds := unpacked[2].([]string)
	vouts := unpacked[3].([]uint32)
	amounts := unpacked[4].([]uint64)
	p.Utxos = make([]*UTXO, len(txIds))
	for i, txId := range txIds {
		hash, err := chains.HashFromHex(strings.TrimPrefix(txId, "0x"))
		if err != nil {
			log.Error().Err(err).Msg("txId hash error")
			return err
		}
		p.Utxos[i] = &UTXO{TxID: hash, Vout: vouts[i], AmountInSats: amounts[i]}
	}
	p.RequestId = unpacked[5].([32]byte)
	return nil
}

type RedeemTokenParams struct {
	DestinationChain   string //Bitcoin chain
	DestinationAddress string //Bitcoin user address
	Payload            RedeemTokenPayloadWithType
	RawPayload         []byte //Use in unpacking
	Symbol             string
	Amount             uint64
	CustodianGroupUID  [32]byte
	SessionSequence    uint64
}

func (p *RedeemTokenParams) AbiPack() ([]byte, error) {
	payload, err := p.Payload.AbiPack()
	if err != nil {
		return nil, err
	}
	packed, err := RedeemTokenArguments.Pack(
		p.DestinationChain,
		p.DestinationAddress,
		payload,
		p.Symbol,
		big.NewInt(int64(p.Amount)),
		p.CustodianGroupUID,
		p.SessionSequence,
	)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

func (p *RedeemTokenParams) AbiUnpack(data []byte) error {
	unpacked, err := RedeemTokenArguments.Unpack(data)
	if err != nil {
		log.Error().Err(err).Msg("redeem token params abi unpack error")
		return err
	}
	p.DestinationChain = unpacked[0].(string)
	p.DestinationAddress = unpacked[1].(string)
	p.RawPayload = unpacked[2].([]byte)
	//Remove prefix
	err = p.Payload.AbiUnpack(p.RawPayload)
	if err != nil {
		log.Error().Err(err).Msg("payload abi unpack error")
		return err
	}
	p.Symbol = unpacked[3].(string)
	p.Amount = unpacked[4].(*big.Int).Uint64()
	p.CustodianGroupUID = unpacked[5].([32]byte)
	p.SessionSequence = unpacked[6].(uint64)
	return nil
}

func (utxo *UTXO) AppendReserved(requestID string, amount uint64) error {
	if utxo.Reservations == nil {
		utxo.Reservations = []*Reservation{}
	}
	totalReserved := uint64(0)
	for _, reserved := range utxo.Reservations {
		if reserved.Request == requestID {
			return fmt.Errorf("requestID already reserved in this utxo %s", utxo.TxID.Hex())
		}
		totalReserved += reserved.Amount
	}
	if totalReserved+amount > utxo.AmountInSats {
		return fmt.Errorf("amount exceeds utxo amount, totalReserved %d, amount %d, utxo.AmountInSats %d", totalReserved, amount, utxo.AmountInSats)
	}
	utxo.Reservations = append(utxo.Reservations, &Reservation{
		Request: requestID,
		Amount:  amount,
	})
	return nil
}
func (utxo *UTXO) GetReservedAmount() uint64 {
	amount := uint64(0)
	for _, reserved := range utxo.Reservations {
		amount += reserved.Amount
	}
	return amount
}

func (utxo *UTXO) AvailableAmount() uint64 {
	reservedAmount := utxo.GetReservedAmount()
	return utxo.AmountInSats - reservedAmount
}

func (utxo *UTXO) IsReserved(requestID string) bool {
	for _, reserved := range utxo.Reservations {
		if reserved.Request == requestID {
			return true
		}
	}
	return false
}

func (utxo *UTXO) Release(requestID string) uint64 {
	if utxo.Reservations == nil {
		return 0
	}
	i := 0
	releaseAmount := uint64(0)
	for _, reserved := range utxo.Reservations {
		if reserved.Request != requestID {
			utxo.Reservations[i] = reserved
			i++
		} else {
			releaseAmount = reserved.Amount
		}
	}
	utxo.Reservations = utxo.Reservations[:i]
	return releaseAmount
}

func (s *UTXOSnapshot) ReleaseUtxos(requestID string) uint64 {
	releasedAmount := uint64(0)
	for _, utxo := range s.Utxos {
		releasedAmount += utxo.Release(requestID)
	}
	return releasedAmount
}
func (s *UTXOSnapshot) CountInputOutput() (int, int) {
	inputs := 0
	mapRequests := map[string]bool{}
	for _, utxo := range s.Utxos {
		if len(utxo.Reservations) > 0 {
			inputs += 1
			for _, reservation := range utxo.Reservations {
				mapRequests[reservation.Request] = true
			}
		}
	}
	return inputs, len(mapRequests)
}

// Return true if utxo is appended
func (s *UTXOSnapshot) AppendUtxo(utxo *UTXO) bool {
	exists := false
	for _, item := range s.Utxos {
		if bytes.Equal(item.TxID[:], utxo.TxID[:]) && item.Vout == utxo.Vout {
			exists = true
			break
		}
	}
	if !exists {
		s.Utxos = append(s.Utxos, utxo)
		return true
	}
	return false
}

// Utxos list is sorted by amount in sats and txid for deterministic results
// Each utxo is reserved if it is part of the optimal combination
func (s *UTXOSnapshot) ReserveUtxos(requestID string, amount uint64, quorum uint64, vSizeLimit uint64) ([]*UTXO, error) {
	currentInputs, currentOutputs := s.CountInputOutput()
	newInput := 0
	newOutput := 1
	remainingAmount := amount
	reserveUtxos := make([]*UTXO, 0)
	mapNewResevations := map[int]uint64{}
	for ind, utxo := range s.Utxos {
		if utxo.IsReserved(requestID) {
			return nil, fmt.Errorf("requestID %s is already reserved in utxo %s", requestID, utxo.TxID.Hex())
		}
		availableAmount := utxo.AvailableAmount()
		if availableAmount > 0 {
			//Reserve amount is min(availableAmount, remainingAmount)
			reserveAmount := availableAmount
			if reserveAmount > remainingAmount {
				reserveAmount = remainingAmount
			}
			remainingAmount -= reserveAmount
			mapNewResevations[ind] = reserveAmount
			reserveUtxos = append(reserveUtxos, &UTXO{
				TxID:         utxo.TxID,
				Vout:         utxo.Vout,
				AmountInSats: reserveAmount,
				ScriptPubkey: utxo.ScriptPubkey,
			})
			//First reservation
			if len(utxo.Reservations) == 0 {
				newInput += 1
			}
		}
		if remainingAmount == 0 {
			break
		}
	}
	if remainingAmount > 0 {
		return nil, fmt.Errorf("not enough utxos to reserve, remainingAmount %d", remainingAmount)
	}
	//Add extra input and output for collect change amount
	newVsize := utils.CalculateVsize(currentInputs+newInput+1, currentOutputs+newOutput+1, quorum)
	if newVsize > vSizeLimit {
		return nil, fmt.Errorf("new virtual size exceeds the limit %d > %d", newVsize, vSizeLimit)
	}

	for ind, reserveAmount := range mapNewResevations {
		utxo := s.Utxos[ind]
		utxo.AppendReserved(requestID, reserveAmount)
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

func (s *UTXOSnapshot) FindUtxo(txID chains.Hash, vout uint32) *UTXO {
	for _, utxo := range s.Utxos {
		if utxo.TxID == txID && utxo.Vout == vout {
			return utxo
		}
	}
	return nil
}

func (s *UTXOSnapshot) Compare(other *UTXOSnapshot) bool {
	if other == nil {
		return false
	}
	if len(s.Utxos) != len(other.Utxos) {
		return false
	}
	for i, utxo := range s.Utxos {
		otherUtxo := other.Utxos[i]
		if !bytes.Equal(utxo.TxID[:], otherUtxo.TxID[:]) || utxo.Vout != otherUtxo.Vout {
			return false
		}
	}
	return true
}

func (s *UTXOSnapshot) ToString() string {
	output := fmt.Sprintf("GroupUid:%s;", s.CustodianGroupUID.Hex())
	for _, utxo := range s.Utxos {
		output += fmt.Sprintf("%s:%d;", utxo.TxID.Hex(), utxo.Vout)
	}
	return output
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
