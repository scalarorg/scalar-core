// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/evm/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	utils "github.com/scalarorg/scalar-core/utils"
	exported "github.com/scalarorg/scalar-core/x/nexus/exported"
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

// Params is the parameter set for this module
type Params struct {
	Chain               github_com_scalarorg_scalar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/scalarorg/scalar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
	ConfirmationHeight  uint64                                                      `protobuf:"varint,2,opt,name=confirmation_height,json=confirmationHeight,proto3" json:"confirmation_height,omitempty"`
	Network             string                                                      `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	TokenCode           []byte                                                      `protobuf:"bytes,5,opt,name=token_code,json=tokenCode,proto3" json:"token_code,omitempty"`
	Burnable            []byte                                                      `protobuf:"bytes,6,opt,name=burnable,proto3" json:"burnable,omitempty"`
	RevoteLockingPeriod int64                                                       `protobuf:"varint,7,opt,name=revote_locking_period,json=revoteLockingPeriod,proto3" json:"revote_locking_period,omitempty"`
	Networks            []NetworkInfo                                               `protobuf:"bytes,8,rep,name=networks,proto3" json:"networks"`
	VotingThreshold     utils.Threshold                                             `protobuf:"bytes,9,opt,name=voting_threshold,json=votingThreshold,proto3" json:"voting_threshold"`
	MinVoterCount       int64                                                       `protobuf:"varint,10,opt,name=min_voter_count,json=minVoterCount,proto3" json:"min_voter_count,omitempty"`
	CommandsGasLimit    uint32                                                      `protobuf:"varint,11,opt,name=commands_gas_limit,json=commandsGasLimit,proto3" json:"commands_gas_limit,omitempty"`
	VotingGracePeriod   int64                                                       `protobuf:"varint,13,opt,name=voting_grace_period,json=votingGracePeriod,proto3" json:"voting_grace_period,omitempty"`
	EndBlockerLimit     int64                                                       `protobuf:"varint,14,opt,name=end_blocker_limit,json=endBlockerLimit,proto3" json:"end_blocker_limit,omitempty"`
	TransferLimit       uint64                                                      `protobuf:"varint,15,opt,name=transfer_limit,json=transferLimit,proto3" json:"transfer_limit,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e217ba5e728e664, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

type PendingChain struct {
	Params Params         `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Chain  exported.Chain `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain"`
}

func (m *PendingChain) Reset()         { *m = PendingChain{} }
func (m *PendingChain) String() string { return proto.CompactTextString(m) }
func (*PendingChain) ProtoMessage()    {}
func (*PendingChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e217ba5e728e664, []int{1}
}
func (m *PendingChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingChain.Merge(m, src)
}
func (m *PendingChain) XXX_Size() int {
	return m.Size()
}
func (m *PendingChain) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingChain.DiscardUnknown(m)
}

var xxx_messageInfo_PendingChain proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "scalar.evm.v1beta1.Params")
	proto.RegisterType((*PendingChain)(nil), "scalar.evm.v1beta1.PendingChain")
}

func init() { proto.RegisterFile("scalar/evm/v1beta1/params.proto", fileDescriptor_6e217ba5e728e664) }

var fileDescriptor_6e217ba5e728e664 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xd3, 0x3c,
	0x1c, 0x6f, 0xb6, 0xb6, 0xeb, 0xbc, 0x75, 0xed, 0xbc, 0xe7, 0x91, 0xa2, 0x4a, 0xa4, 0xd1, 0x34,
	0x50, 0x41, 0x23, 0xd1, 0xc6, 0x05, 0x89, 0x03, 0xd0, 0x1d, 0x06, 0xd3, 0x34, 0x55, 0x11, 0x70,
	0xe0, 0x12, 0xb9, 0xc9, 0x7f, 0x89, 0xb5, 0xc4, 0xae, 0x1c, 0xb7, 0x8c, 0xaf, 0xc0, 0x09, 0x89,
	0x2f, 0xb5, 0xe3, 0x8e, 0x9c, 0x26, 0xd8, 0xbe, 0x05, 0x27, 0x14, 0xdb, 0x89, 0x26, 0xe8, 0x81,
	0x9b, 0xf3, 0x7b, 0xb3, 0xf3, 0x7f, 0x41, 0xc3, 0x22, 0x22, 0x19, 0x11, 0x3e, 0x2c, 0x72, 0x7f,
	0x71, 0x30, 0x05, 0x49, 0x0e, 0xfc, 0x19, 0x11, 0x24, 0x2f, 0xbc, 0x99, 0xe0, 0x92, 0x63, 0xac,
	0x05, 0x1e, 0x2c, 0x72, 0xcf, 0x08, 0x06, 0x7b, 0xc6, 0x34, 0x97, 0x34, 0x2b, 0x6a, 0x9b, 0x4c,
	0x05, 0x14, 0x29, 0xcf, 0x62, 0xed, 0x1c, 0x38, 0x4b, 0xa2, 0xe5, 0xe7, 0x19, 0x98, 0xe4, 0xc1,
	0x7f, 0x09, 0x4f, 0xb8, 0x3a, 0xfa, 0xe5, 0xc9, 0xa0, 0x8f, 0x8d, 0x8b, 0xc1, 0xe5, 0xbc, 0xf0,
	0xe1, 0x72, 0xc6, 0x85, 0x84, 0x78, 0x59, 0xc0, 0xee, 0xb7, 0x16, 0x6a, 0x4f, 0xd4, 0x5b, 0xf1,
	0x7b, 0xd4, 0x8a, 0x52, 0x42, 0x99, 0x6d, 0xb9, 0xd6, 0x68, 0x7d, 0xfc, 0xf2, 0xd7, 0xcd, 0xf0,
	0x45, 0x42, 0x65, 0x3a, 0x9f, 0x7a, 0x11, 0xcf, 0x7d, 0x9d, 0xc9, 0x45, 0x62, 0x4e, 0x4f, 0x23,
	0x2e, 0xc0, 0xbf, 0xfc, 0xe3, 0x12, 0xef, 0xa8, 0x8c, 0x38, 0x23, 0x39, 0x04, 0x3a, 0x0d, 0xfb,
	0x68, 0x27, 0xe2, 0xec, 0x9c, 0x8a, 0x9c, 0x48, 0xca, 0x59, 0x98, 0x02, 0x4d, 0x52, 0x69, 0xaf,
	0xb8, 0xd6, 0xa8, 0x19, 0xe0, 0xfb, 0xd4, 0x1b, 0xc5, 0x60, 0x1b, 0xad, 0x31, 0x90, 0x9f, 0xb8,
	0xb8, 0xb0, 0x57, 0xcb, 0x97, 0x04, 0xd5, 0x27, 0x7e, 0x80, 0x90, 0xe4, 0x17, 0xc0, 0xc2, 0x88,
	0xc7, 0x60, 0xb7, 0x5c, 0x6b, 0xb4, 0x19, 0xac, 0x2b, 0xe4, 0x88, 0xc7, 0x80, 0x07, 0xa8, 0x33,
	0x9d, 0x0b, 0x46, 0xa6, 0x19, 0xd8, 0x6d, 0x45, 0xd6, 0xdf, 0xf8, 0x10, 0xfd, 0x2f, 0x60, 0xc1,
	0x25, 0x84, 0x19, 0x8f, 0x2e, 0x28, 0x4b, 0xc2, 0x19, 0x08, 0xca, 0x63, 0x7b, 0xcd, 0xb5, 0x46,
	0xab, 0xc1, 0x8e, 0x26, 0x4f, 0x35, 0x37, 0x51, 0x14, 0x7e, 0x8d, 0x3a, 0xe6, 0xe6, 0xc2, 0xee,
	0xb8, 0xab, 0xa3, 0x8d, 0xc3, 0xa1, 0xf7, 0x77, 0x27, 0xbd, 0x33, 0xad, 0x79, 0xcb, 0xce, 0xf9,
	0xb8, 0x79, 0x75, 0x33, 0x6c, 0x04, 0xb5, 0x0d, 0x4f, 0x50, 0x7f, 0xc1, 0x65, 0x79, 0x5d, 0xdd,
	0x59, 0x7b, 0xdd, 0xb5, 0xee, 0x47, 0xa9, 0x01, 0xa8, 0xc3, 0xde, 0x55, 0x32, 0x13, 0xd5, 0xd3,
	0xf6, 0x1a, 0xc6, 0x8f, 0x50, 0x2f, 0xa7, 0x2c, 0x2c, 0x5f, 0x2b, 0xc2, 0x88, 0xcf, 0x99, 0xb4,
	0x91, 0xfa, 0x85, 0x6e, 0x4e, 0xd9, 0x87, 0x12, 0x3d, 0x2a, 0x41, 0xbc, 0x8f, 0x70, 0xc4, 0xf3,
	0x9c, 0xb0, 0xb8, 0x08, 0x13, 0x52, 0x84, 0x19, 0xcd, 0xa9, 0xb4, 0x37, 0x5c, 0x6b, 0xd4, 0x0d,
	0xfa, 0x15, 0x73, 0x4c, 0x8a, 0xd3, 0x12, 0xc7, 0x1e, 0xda, 0x31, 0xef, 0x4c, 0x04, 0x89, 0xa0,
	0x2a, 0x4e, 0x57, 0x25, 0x6f, 0x6b, 0xea, 0xb8, 0x64, 0x4c, 0x69, 0x9e, 0xa0, 0x6d, 0x60, 0x71,
	0x38, 0x2d, 0x8b, 0x09, 0xc2, 0x84, 0x6f, 0x29, 0x75, 0x0f, 0x58, 0x3c, 0xd6, 0xb8, 0xce, 0x7e,
	0x88, 0xb6, 0xa4, 0x20, 0xac, 0x38, 0xaf, 0x85, 0x3d, 0xd5, 0xfb, 0x6e, 0x85, 0x2a, 0xd9, 0x49,
	0xb3, 0xd3, 0xec, 0xb7, 0x4e, 0x9a, 0x9d, 0xcd, 0x7e, 0x77, 0xf7, 0x8b, 0x85, 0x36, 0x27, 0xc0,
	0x62, 0xca, 0x12, 0x35, 0x4f, 0xf8, 0x39, 0x6a, 0xeb, 0x8d, 0x52, 0xc3, 0xb9, 0x71, 0x38, 0x58,
	0xd6, 0x08, 0x3d, 0xc7, 0xa6, 0x70, 0x46, 0x8f, 0x5f, 0x55, 0x53, 0xbd, 0xa2, 0x8c, 0x7b, 0x95,
	0x51, 0x8d, 0xad, 0x57, 0x8f, 0x6d, 0x95, 0xa1, 0xae, 0x33, 0x11, 0xda, 0x38, 0x3e, 0xb9, 0xfa,
	0xe9, 0x34, 0xae, 0x6e, 0x1d, 0xeb, 0xfa, 0xd6, 0xb1, 0x7e, 0xdc, 0x3a, 0xd6, 0xd7, 0x3b, 0xa7,
	0x71, 0x7d, 0xe7, 0x34, 0xbe, 0xdf, 0x39, 0x8d, 0x8f, 0xfb, 0xff, 0xb0, 0x22, 0xe5, 0xf6, 0xaa,
	0xa5, 0x9b, 0xb6, 0xd5, 0xd6, 0x3d, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x3c, 0xd4, 0xdb,
	0x33, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TransferLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TransferLimit))
		i--
		dAtA[i] = 0x78
	}
	if m.EndBlockerLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EndBlockerLimit))
		i--
		dAtA[i] = 0x70
	}
	if m.VotingGracePeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.VotingGracePeriod))
		i--
		dAtA[i] = 0x68
	}
	if m.CommandsGasLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CommandsGasLimit))
		i--
		dAtA[i] = 0x58
	}
	if m.MinVoterCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinVoterCount))
		i--
		dAtA[i] = 0x50
	}
	{
		size, err := m.VotingThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.Networks) > 0 {
		for iNdEx := len(m.Networks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Networks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.RevoteLockingPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RevoteLockingPeriod))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Burnable) > 0 {
		i -= len(m.Burnable)
		copy(dAtA[i:], m.Burnable)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Burnable)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TokenCode) > 0 {
		i -= len(m.TokenCode)
		copy(dAtA[i:], m.TokenCode)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TokenCode)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ConfirmationHeight != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ConfirmationHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PendingChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Chain.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.ConfirmationHeight != 0 {
		n += 1 + sovParams(uint64(m.ConfirmationHeight))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.TokenCode)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.Burnable)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.RevoteLockingPeriod != 0 {
		n += 1 + sovParams(uint64(m.RevoteLockingPeriod))
	}
	if len(m.Networks) > 0 {
		for _, e := range m.Networks {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.VotingThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MinVoterCount != 0 {
		n += 1 + sovParams(uint64(m.MinVoterCount))
	}
	if m.CommandsGasLimit != 0 {
		n += 1 + sovParams(uint64(m.CommandsGasLimit))
	}
	if m.VotingGracePeriod != 0 {
		n += 1 + sovParams(uint64(m.VotingGracePeriod))
	}
	if m.EndBlockerLimit != 0 {
		n += 1 + sovParams(uint64(m.EndBlockerLimit))
	}
	if m.TransferLimit != 0 {
		n += 1 + sovParams(uint64(m.TransferLimit))
	}
	return n
}

func (m *PendingChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Chain.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = github_com_scalarorg_scalar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfirmationHeight", wireType)
			}
			m.ConfirmationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConfirmationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenCode", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenCode = append(m.TokenCode[:0], dAtA[iNdEx:postIndex]...)
			if m.TokenCode == nil {
				m.TokenCode = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Burnable", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Burnable = append(m.Burnable[:0], dAtA[iNdEx:postIndex]...)
			if m.Burnable == nil {
				m.Burnable = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevoteLockingPeriod", wireType)
			}
			m.RevoteLockingPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RevoteLockingPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Networks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Networks = append(m.Networks, NetworkInfo{})
			if err := m.Networks[len(m.Networks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VotingThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinVoterCount", wireType)
			}
			m.MinVoterCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinVoterCount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommandsGasLimit", wireType)
			}
			m.CommandsGasLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommandsGasLimit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingGracePeriod", wireType)
			}
			m.VotingGracePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingGracePeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndBlockerLimit", wireType)
			}
			m.EndBlockerLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndBlockerLimit |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferLimit", wireType)
			}
			m.TransferLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransferLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *PendingChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: PendingChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Chain.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
