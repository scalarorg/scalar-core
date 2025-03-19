package types

import (
	covExported "github.com/scalarorg/scalar-core/x/covenant/exported"
)

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
