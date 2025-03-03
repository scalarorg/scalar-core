// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/protocol/v1beta1/query.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_scalarorg_scalar_core_x_nexus_exported "github.com/scalarorg/scalar-core/x/nexus/exported"
	exported "github.com/scalarorg/scalar-core/x/protocol/exported"
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

type ProtocolsRequest struct {
	Pubkey  string          `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Address string          `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Name    string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Status  exported.Status `protobuf:"varint,4,opt,name=status,proto3,enum=scalar.protocol.exported.v1beta1.Status" json:"status,omitempty"`
}

func (m *ProtocolsRequest) Reset()         { *m = ProtocolsRequest{} }
func (m *ProtocolsRequest) String() string { return proto.CompactTextString(m) }
func (*ProtocolsRequest) ProtoMessage()    {}
func (*ProtocolsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17e7473aa0548727, []int{0}
}
func (m *ProtocolsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolsRequest.Merge(m, src)
}
func (m *ProtocolsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolsRequest proto.InternalMessageInfo

func (m *ProtocolsRequest) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *ProtocolsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ProtocolsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtocolsRequest) GetStatus() exported.Status {
	if m != nil {
		return m.Status
	}
	return exported.Unspecified
}

type ProtocolsResponse struct {
	Protocols []*ProtocolDetails `protobuf:"bytes,1,rep,name=protocols,proto3" json:"protocols,omitempty"`
	Total     uint64             `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (m *ProtocolsResponse) Reset()         { *m = ProtocolsResponse{} }
func (m *ProtocolsResponse) String() string { return proto.CompactTextString(m) }
func (*ProtocolsResponse) ProtoMessage()    {}
func (*ProtocolsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17e7473aa0548727, []int{1}
}
func (m *ProtocolsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolsResponse.Merge(m, src)
}
func (m *ProtocolsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolsResponse proto.InternalMessageInfo

func (m *ProtocolsResponse) GetProtocols() []*ProtocolDetails {
	if m != nil {
		return m.Protocols
	}
	return nil
}

func (m *ProtocolsResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type ProtocolRequest struct {
	OriginChain github_com_scalarorg_scalar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=origin_chain,json=originChain,proto3,casttype=github.com/scalarorg/scalar-core/x/nexus/exported.ChainName" json:"origin_chain,omitempty"`
	MinorChain  github_com_scalarorg_scalar_core_x_nexus_exported.ChainName `protobuf:"bytes,2,opt,name=minor_chain,json=minorChain,proto3,casttype=github.com/scalarorg/scalar-core/x/nexus/exported.ChainName" json:"minor_chain,omitempty"`
	Symbol      string                                                      `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Address     string                                                      `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *ProtocolRequest) Reset()         { *m = ProtocolRequest{} }
func (m *ProtocolRequest) String() string { return proto.CompactTextString(m) }
func (*ProtocolRequest) ProtoMessage()    {}
func (*ProtocolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17e7473aa0548727, []int{2}
}
func (m *ProtocolRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolRequest.Merge(m, src)
}
func (m *ProtocolRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolRequest proto.InternalMessageInfo

func (m *ProtocolRequest) GetOriginChain() github_com_scalarorg_scalar_core_x_nexus_exported.ChainName {
	if m != nil {
		return m.OriginChain
	}
	return ""
}

func (m *ProtocolRequest) GetMinorChain() github_com_scalarorg_scalar_core_x_nexus_exported.ChainName {
	if m != nil {
		return m.MinorChain
	}
	return ""
}

func (m *ProtocolRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ProtocolRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type ProtocolResponse struct {
	Protocol *ProtocolDetails `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
}

func (m *ProtocolResponse) Reset()         { *m = ProtocolResponse{} }
func (m *ProtocolResponse) String() string { return proto.CompactTextString(m) }
func (*ProtocolResponse) ProtoMessage()    {}
func (*ProtocolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_17e7473aa0548727, []int{3}
}
func (m *ProtocolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolResponse.Merge(m, src)
}
func (m *ProtocolResponse) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolResponse proto.InternalMessageInfo

func (m *ProtocolResponse) GetProtocol() *ProtocolDetails {
	if m != nil {
		return m.Protocol
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtocolsRequest)(nil), "scalar.protocol.v1beta1.ProtocolsRequest")
	proto.RegisterType((*ProtocolsResponse)(nil), "scalar.protocol.v1beta1.ProtocolsResponse")
	proto.RegisterType((*ProtocolRequest)(nil), "scalar.protocol.v1beta1.ProtocolRequest")
	proto.RegisterType((*ProtocolResponse)(nil), "scalar.protocol.v1beta1.ProtocolResponse")
}

func init() {
	proto.RegisterFile("scalar/protocol/v1beta1/query.proto", fileDescriptor_17e7473aa0548727)
}

var fileDescriptor_17e7473aa0548727 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3d, 0xaf, 0xd3, 0x30,
	0x14, 0xad, 0xdb, 0x10, 0xa8, 0x8b, 0xf8, 0xb0, 0xaa, 0x12, 0x75, 0x08, 0x55, 0x58, 0x32, 0x40,
	0xa2, 0xc2, 0xc8, 0x00, 0x2a, 0x15, 0x0b, 0x12, 0x42, 0x61, 0x41, 0x2c, 0xe0, 0xa4, 0x56, 0x1a,
	0x91, 0xc4, 0xa9, 0xed, 0xa0, 0xe6, 0x27, 0xb0, 0x31, 0xf3, 0x8b, 0x18, 0x3b, 0x32, 0x21, 0xd4,
	0xfe, 0x0b, 0x26, 0x54, 0x7f, 0xa4, 0xa5, 0xea, 0x93, 0xde, 0x93, 0xde, 0x76, 0xef, 0xf5, 0x39,
	0xc7, 0xf7, 0x9e, 0x7b, 0xe1, 0x23, 0x9e, 0xe0, 0x1c, 0xb3, 0xb0, 0x62, 0x54, 0xd0, 0x84, 0xe6,
	0xe1, 0xd7, 0x69, 0x4c, 0x04, 0x9e, 0x86, 0xab, 0x9a, 0xb0, 0x26, 0x90, 0x65, 0xf4, 0x40, 0x81,
	0x02, 0x03, 0x0a, 0x34, 0x68, 0x7c, 0x21, 0x5b, 0x34, 0x15, 0xe1, 0x0a, 0x3f, 0x7e, 0x7c, 0x0a,
	0x22, 0xeb, 0x8a, 0x32, 0x41, 0x16, 0x67, 0xd1, 0xc3, 0x94, 0xa6, 0x54, 0x86, 0xe1, 0x3e, 0x52,
	0x55, 0xef, 0x07, 0x80, 0xf7, 0xde, 0x69, 0x3e, 0x8f, 0xc8, 0xaa, 0x26, 0x5c, 0xa0, 0x11, 0xb4,
	0xab, 0x3a, 0xfe, 0x42, 0x1a, 0x07, 0x4c, 0x80, 0xdf, 0x8f, 0x74, 0x86, 0x1c, 0x78, 0x13, 0x2f,
	0x16, 0x8c, 0x70, 0xee, 0x74, 0xe5, 0x83, 0x49, 0x11, 0x82, 0x56, 0x89, 0x0b, 0xe2, 0xf4, 0x64,
	0x59, 0xc6, 0xe8, 0x25, 0xb4, 0xb9, 0xc0, 0xa2, 0xe6, 0x8e, 0x35, 0x01, 0xfe, 0x9d, 0xa7, 0x7e,
	0x70, 0x3a, 0xad, 0xe9, 0xd7, 0x8c, 0x1d, 0xbc, 0x97, 0xf8, 0x48, 0xf3, 0xbc, 0x15, 0xbc, 0x7f,
	0xd4, 0x1b, 0xaf, 0x68, 0xc9, 0x09, 0x7a, 0x0d, 0xfb, 0x46, 0x80, 0x3b, 0x60, 0xd2, 0xf3, 0x07,
	0x67, 0x94, 0x8d, 0xa0, 0xa1, 0xcf, 0x89, 0xc0, 0x59, 0xce, 0xa3, 0x03, 0x15, 0x0d, 0xe1, 0x0d,
	0x41, 0x05, 0xce, 0xe5, 0x28, 0x56, 0xa4, 0x12, 0xef, 0x5b, 0x17, 0xde, 0x35, 0x24, 0x63, 0x47,
	0x0c, 0x6f, 0x53, 0x96, 0xa5, 0x59, 0xf9, 0x29, 0x59, 0xe2, 0xac, 0x54, 0xa6, 0xcc, 0x5e, 0xfc,
	0xfd, 0xfd, 0xf0, 0x79, 0x9a, 0x89, 0x65, 0x1d, 0x07, 0x09, 0x2d, 0x42, 0xd5, 0x02, 0x65, 0xa9,
	0x8e, 0x9e, 0x24, 0x94, 0x91, 0x70, 0x1d, 0x96, 0x64, 0x5d, 0xf3, 0x76, 0x35, 0xc1, 0xab, 0xbd,
	0xc4, 0x5b, 0x5c, 0x90, 0x68, 0xa0, 0x44, 0x65, 0x01, 0x7d, 0x86, 0x83, 0x22, 0x2b, 0x29, 0xd3,
	0x5f, 0x74, 0xaf, 0xe7, 0x0b, 0x28, 0x35, 0xd5, 0x0f, 0x23, 0x68, 0xf3, 0xa6, 0x88, 0x69, 0xae,
	0x97, 0xa4, 0xb3, 0xe3, 0xa5, 0x5a, 0xff, 0x2d, 0xd5, 0xfb, 0x70, 0x38, 0x8d, 0xd6, 0xfd, 0x39,
	0xbc, 0x65, 0x2c, 0x94, 0x3e, 0x5c, 0xc5, 0xfc, 0x96, 0x39, 0x7b, 0xf3, 0x73, 0xeb, 0x82, 0xcd,
	0xd6, 0x05, 0x7f, 0xb6, 0x2e, 0xf8, 0xbe, 0x73, 0x3b, 0x9b, 0x9d, 0xdb, 0xf9, 0xb5, 0x73, 0x3b,
	0x1f, 0xa7, 0x97, 0x18, 0xb7, 0xbd, 0x77, 0x79, 0xde, 0xb1, 0x2d, 0xf3, 0x67, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x15, 0xff, 0x4a, 0x72, 0x03, 0x00, 0x00,
}

func (m *ProtocolsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProtocolsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Total != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Protocols) > 0 {
		for iNdEx := len(m.Protocols) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Protocols[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ProtocolRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MinorChain) > 0 {
		i -= len(m.MinorChain)
		copy(dAtA[i:], m.MinorChain)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.MinorChain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OriginChain) > 0 {
		i -= len(m.OriginChain)
		copy(dAtA[i:], m.OriginChain)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.OriginChain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProtocolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Protocol != nil {
		{
			size, err := m.Protocol.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProtocolsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovQuery(uint64(m.Status))
	}
	return n
}

func (m *ProtocolsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Protocols) > 0 {
		for _, e := range m.Protocols {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Total != 0 {
		n += 1 + sovQuery(uint64(m.Total))
	}
	return n
}

func (m *ProtocolRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OriginChain)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.MinorChain)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *ProtocolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Protocol != nil {
		l = m.Protocol.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProtocolsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ProtocolsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= exported.Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *ProtocolsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ProtocolsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocols", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocols = append(m.Protocols, &ProtocolDetails{})
			if err := m.Protocols[len(m.Protocols)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *ProtocolRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ProtocolRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginChain = github_com_scalarorg_scalar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinorChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinorChain = github_com_scalarorg_scalar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *ProtocolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ProtocolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Protocol == nil {
				m.Protocol = &ProtocolDetails{}
			}
			if err := m.Protocol.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
