// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/btc/v1beta1/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_scalarorg_scalar_core_x_evm_types "github.com/scalarorg/scalar-core/x/evm/types"
	github_com_scalarorg_scalar_core_x_nexus_exported "github.com/scalarorg/scalar-core/x/nexus/exported"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Event_Status int32

const (
	EventNonExistent Event_Status = 0
	EventConfirmed   Event_Status = 1
	EventCompleted   Event_Status = 2
	EventFailed      Event_Status = 3
)

var Event_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_CONFIRMED",
	2: "STATUS_COMPLETED",
	3: "STATUS_FAILED",
}

var Event_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"STATUS_CONFIRMED":   1,
	"STATUS_COMPLETED":   2,
	"STATUS_FAILED":      3,
}

func (x Event_Status) String() string {
	return proto.EnumName(Event_Status_name, int32(x))
}

func (Event_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_352080992c371856, []int{1, 0}
}

type VoteEvents struct {
	Chain  github_com_scalarorg_scalar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/scalarorg/scalar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
	Events []Event                                                     `protobuf:"bytes,2,rep,name=events,proto3" json:"events"`
}

func (m *VoteEvents) Reset()         { *m = VoteEvents{} }
func (m *VoteEvents) String() string { return proto.CompactTextString(m) }
func (*VoteEvents) ProtoMessage()    {}
func (*VoteEvents) Descriptor() ([]byte, []int) {
	return fileDescriptor_352080992c371856, []int{0}
}
func (m *VoteEvents) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoteEvents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoteEvents.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoteEvents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoteEvents.Merge(m, src)
}
func (m *VoteEvents) XXX_Size() int {
	return m.Size()
}
func (m *VoteEvents) XXX_DiscardUnknown() {
	xxx_messageInfo_VoteEvents.DiscardUnknown(m)
}

var xxx_messageInfo_VoteEvents proto.InternalMessageInfo

type Event struct {
	Chain  github_com_scalarorg_scalar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/scalarorg/scalar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
	TxID   Hash                                                        `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3,customtype=Hash" json:"tx_id"`
	Status Event_Status                                                `protobuf:"varint,3,opt,name=status,proto3,enum=scalar.btc.v1beta1.Event_Status" json:"status,omitempty"`
	Index  uint64                                                      `protobuf:"varint,4,opt,name=index,proto3" json:"index,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*Event_StakingTx
	Event isEvent_Event `protobuf_oneof:"event"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_352080992c371856, []int{1}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type isEvent_Event interface {
	isEvent_Event()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Event_StakingTx struct {
	StakingTx *EventStakingTx `protobuf:"bytes,5,opt,name=staking_tx,json=stakingTx,proto3,oneof" json:"staking_tx,omitempty"`
}

func (*Event_StakingTx) isEvent_Event() {}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Event) GetStakingTx() *EventStakingTx {
	if x, ok := m.GetEvent().(*Event_StakingTx); ok {
		return x.StakingTx
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_StakingTx)(nil),
	}
}

type EventStakingTx struct {
	PrevOutPoint string                                            `protobuf:"bytes,1,opt,name=prev_out_point,json=prevOutPoint,proto3" json:"prev_out_point,omitempty"`
	Amount       uint64                                            `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Asset        string                                            `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset,omitempty"`
	PayloadHash  github_com_scalarorg_scalar_core_x_evm_types.Hash `protobuf:"bytes,4,opt,name=payload_hash,json=payloadHash,proto3,customtype=github.com/scalarorg/scalar-core/x/evm/types.Hash" json:"payload_hash"`
	Metadata     StakingTxMetadata                                 `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata"`
}

func (m *EventStakingTx) Reset()         { *m = EventStakingTx{} }
func (m *EventStakingTx) String() string { return proto.CompactTextString(m) }
func (*EventStakingTx) ProtoMessage()    {}
func (*EventStakingTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_352080992c371856, []int{2}
}
func (m *EventStakingTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventStakingTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventStakingTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventStakingTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStakingTx.Merge(m, src)
}
func (m *EventStakingTx) XXX_Size() int {
	return m.Size()
}
func (m *EventStakingTx) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStakingTx.DiscardUnknown(m)
}

var xxx_messageInfo_EventStakingTx proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("scalar.btc.v1beta1.Event_Status", Event_Status_name, Event_Status_value)
	proto.RegisterType((*VoteEvents)(nil), "scalar.btc.v1beta1.VoteEvents")
	proto.RegisterType((*Event)(nil), "scalar.btc.v1beta1.Event")
	proto.RegisterType((*EventStakingTx)(nil), "scalar.btc.v1beta1.EventStakingTx")
}

func init() { proto.RegisterFile("scalar/btc/v1beta1/events.proto", fileDescriptor_352080992c371856) }

var fileDescriptor_352080992c371856 = []byte{
	// 635 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xbd, 0x89, 0x93, 0xaf, 0xdd, 0xe4, 0x0b, 0xd1, 0x2a, 0x42, 0x21, 0x07, 0xc7, 0x8a,
	0x40, 0x32, 0x52, 0xb1, 0xd5, 0x72, 0x00, 0xc4, 0x01, 0x35, 0x89, 0x43, 0x83, 0xda, 0xb4, 0x72,
	0x52, 0x0e, 0x08, 0x29, 0xda, 0xc4, 0x4b, 0x62, 0x11, 0x7b, 0x2d, 0xef, 0x26, 0x72, 0xdf, 0x00,
	0x95, 0x0b, 0x0f, 0x40, 0x4f, 0x70, 0xe0, 0x11, 0x78, 0x84, 0x1e, 0x7b, 0x44, 0x1c, 0x22, 0x48,
	0xdf, 0x82, 0x13, 0xf2, 0x7a, 0xa9, 0x54, 0xd1, 0x4a, 0xbd, 0x70, 0xdb, 0x19, 0xff, 0xfc, 0xdf,
	0x99, 0xfd, 0xcf, 0xc0, 0x3a, 0x1b, 0xe3, 0x19, 0x8e, 0xac, 0x11, 0x1f, 0x5b, 0x8b, 0xcd, 0x11,
	0xe1, 0x78, 0xd3, 0x22, 0x0b, 0x12, 0x70, 0x66, 0x86, 0x11, 0xe5, 0x14, 0xa1, 0x14, 0x30, 0x47,
	0x7c, 0x6c, 0x4a, 0xa0, 0x56, 0x99, 0xd0, 0x09, 0x15, 0x9f, 0xad, 0xe4, 0x94, 0x92, 0x35, 0xed,
	0x0a, 0x29, 0x7e, 0x14, 0x12, 0xa9, 0xd4, 0xf8, 0x08, 0x20, 0x7c, 0x49, 0x39, 0xb1, 0x85, 0x3c,
	0x3a, 0x84, 0xb9, 0xf1, 0x14, 0x7b, 0x41, 0x15, 0xe8, 0xc0, 0x58, 0x6f, 0x3e, 0xfb, 0xb5, 0xac,
	0x3f, 0x9d, 0x78, 0x7c, 0x3a, 0x1f, 0x99, 0x63, 0xea, 0x5b, 0xa9, 0x18, 0x8d, 0x26, 0xf2, 0xf4,
	0x60, 0x4c, 0x23, 0x62, 0xc5, 0x56, 0x40, 0xe2, 0x39, 0xb3, 0x48, 0x1c, 0xd2, 0x88, 0x13, 0xd7,
	0x6c, 0x25, 0x12, 0x3d, 0xec, 0x13, 0x27, 0x55, 0x43, 0x8f, 0x60, 0x3e, 0xad, 0xbf, 0x9a, 0xd1,
	0xb3, 0x46, 0x61, 0xeb, 0x8e, 0xf9, 0x77, 0x03, 0xa6, 0x28, 0xa1, 0xa9, 0x9e, 0x2e, 0xeb, 0x8a,
	0x23, 0xf1, 0xc6, 0x2a, 0x0b, 0x73, 0x22, 0xff, 0xaf, 0x2a, 0xbb, 0x0f, 0x73, 0x3c, 0x1e, 0x7a,
	0x6e, 0x35, 0xa3, 0x03, 0xa3, 0xd8, 0xac, 0x24, 0xb7, 0x7f, 0x5f, 0xd6, 0xd5, 0x1d, 0xcc, 0xa6,
	0xab, 0x65, 0x5d, 0x1d, 0xc4, 0xdd, 0xb6, 0xa3, 0xf2, 0xb8, 0xeb, 0xa2, 0xc7, 0x30, 0xcf, 0x38,
	0xe6, 0x73, 0x56, 0xcd, 0xea, 0xc0, 0x28, 0x6d, 0xe9, 0xd7, 0x36, 0x61, 0xf6, 0x05, 0xe7, 0x48,
	0x1e, 0x55, 0x60, 0xce, 0x0b, 0x5c, 0x12, 0x57, 0x55, 0x1d, 0x18, 0xaa, 0x93, 0x06, 0xa8, 0x05,
	0x21, 0xe3, 0xf8, 0xad, 0x17, 0x4c, 0x86, 0x3c, 0xae, 0xe6, 0x74, 0x60, 0x14, 0xb6, 0x1a, 0xd7,
	0x6a, 0xf6, 0x53, 0x74, 0x10, 0xef, 0x28, 0xce, 0x3a, 0xfb, 0x13, 0x34, 0xbe, 0x02, 0x98, 0x4f,
	0x6f, 0x43, 0x1b, 0x10, 0xf5, 0x07, 0xdb, 0x83, 0xc3, 0xfe, 0xf0, 0xb0, 0xd7, 0x3f, 0xb0, 0x5b,
	0xdd, 0x4e, 0xd7, 0x6e, 0x97, 0x95, 0x5a, 0xe5, 0xf8, 0x44, 0x2f, 0x0b, 0x8d, 0x1e, 0x0d, 0xec,
	0xd8, 0x63, 0x3c, 0x79, 0x4f, 0x03, 0x96, 0x25, 0xdd, 0xda, 0xef, 0x75, 0xba, 0xce, 0x9e, 0xdd,
	0x2e, 0x83, 0x1a, 0x3a, 0x3e, 0xd1, 0x4b, 0x82, 0x6d, 0xd1, 0xe0, 0x8d, 0x17, 0xf9, 0xc4, 0xbd,
	0x44, 0xee, 0x1d, 0xec, 0xda, 0x03, 0xbb, 0x5d, 0xce, 0x5c, 0x22, 0xfd, 0x70, 0x46, 0x38, 0x71,
	0x51, 0x03, 0xfe, 0x2f, 0xc9, 0xce, 0x76, 0x77, 0xd7, 0x6e, 0x97, 0xb3, 0xb5, 0x5b, 0xc7, 0x27,
	0x7a, 0x41, 0x60, 0x1d, 0xec, 0xcd, 0x88, 0x5b, 0x5b, 0x7b, 0xf7, 0x49, 0x53, 0xbe, 0x7c, 0xd6,
	0x40, 0xf3, 0x3f, 0x98, 0x13, 0x2e, 0x37, 0xde, 0x67, 0x60, 0xe9, 0x72, 0x8f, 0xe8, 0x2e, 0x2c,
	0x85, 0x11, 0x59, 0x0c, 0xe9, 0x9c, 0x0f, 0x43, 0xea, 0x05, 0x3c, 0xb5, 0xdd, 0x29, 0x26, 0xd9,
	0xfd, 0x39, 0x3f, 0x48, 0x72, 0xe8, 0x36, 0xcc, 0x63, 0x9f, 0xce, 0x03, 0x2e, 0xdc, 0x53, 0x1d,
	0x19, 0x25, 0xef, 0x8d, 0x19, 0x23, 0x5c, 0x18, 0xb5, 0xee, 0xa4, 0x01, 0x7a, 0x0d, 0x8b, 0x21,
	0x3e, 0x9a, 0x51, 0xec, 0x0e, 0xa7, 0x98, 0x4d, 0x85, 0x19, 0xc5, 0xe6, 0x13, 0xe9, 0xf8, 0xe6,
	0x0d, 0x86, 0x89, 0x2c, 0x7c, 0xb9, 0x3c, 0xc9, 0x78, 0x38, 0x05, 0x29, 0x97, 0x04, 0xe8, 0x39,
	0x5c, 0xf3, 0x09, 0xc7, 0x2e, 0xe6, 0x58, 0x7a, 0x79, 0xef, 0x2a, 0x2f, 0x2f, 0x5a, 0xdc, 0x93,
	0xb0, 0x1c, 0xf8, 0x8b, 0x9f, 0x9b, 0x2f, 0x4e, 0x7f, 0x6a, 0xca, 0xe9, 0x4a, 0x03, 0x67, 0x2b,
	0x0d, 0xfc, 0x58, 0x69, 0xe0, 0xc3, 0xb9, 0xa6, 0x9c, 0x9d, 0x6b, 0xca, 0xb7, 0x73, 0x4d, 0x79,
	0xb5, 0x71, 0x83, 0x32, 0x93, 0x5d, 0x17, 0x65, 0x8e, 0xf2, 0x62, 0xc9, 0x1f, 0xfe, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0xd0, 0x38, 0x3c, 0x13, 0x51, 0x04, 0x00, 0x00,
}

func (m *VoteEvents) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoteEvents) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoteEvents) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Events) > 0 {
		for iNdEx := len(m.Events) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Events[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Event != nil {
		{
			size := m.Event.Size()
			i -= size
			if _, err := m.Event.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Index != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x20
	}
	if m.Status != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.TxID.Size()
		i -= size
		if _, err := m.TxID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Event_StakingTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event_StakingTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StakingTx != nil {
		{
			size, err := m.StakingTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *EventStakingTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventStakingTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventStakingTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.PayloadHash.Size()
		i -= size
		if _, err := m.PayloadHash.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Amount != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PrevOutPoint) > 0 {
		i -= len(m.PrevOutPoint)
		copy(dAtA[i:], m.PrevOutPoint)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PrevOutPoint)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VoteEvents) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Events) > 0 {
		for _, e := range m.Events {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.TxID.Size()
	n += 1 + l + sovEvents(uint64(l))
	if m.Status != 0 {
		n += 1 + sovEvents(uint64(m.Status))
	}
	if m.Index != 0 {
		n += 1 + sovEvents(uint64(m.Index))
	}
	if m.Event != nil {
		n += m.Event.Size()
	}
	return n
}

func (m *Event_StakingTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StakingTx != nil {
		l = m.StakingTx.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}
func (m *EventStakingTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PrevOutPoint)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovEvents(uint64(m.Amount))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.PayloadHash.Size()
	n += 1 + l + sovEvents(uint64(l))
	l = m.Metadata.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VoteEvents) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VoteEvents: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoteEvents: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = github_com_scalarorg_scalar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Events", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Events = append(m.Events, Event{})
			if err := m.Events[len(m.Events)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = github_com_scalarorg_scalar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Event_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakingTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &EventStakingTx{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &Event_StakingTx{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EventStakingTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventStakingTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventStakingTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevOutPoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevOutPoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayloadHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PayloadHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
