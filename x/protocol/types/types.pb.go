// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/protocol/v1beta1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/scalarorg/scalar-core/x/covenant/types"
	types "github.com/scalarorg/scalar-core/x/evm/types"
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

type LiquidityModel int32

const (
	Pooling       LiquidityModel = 0
	Transactional LiquidityModel = 1
)

var LiquidityModel_name = map[int32]string{
	0: "LIQUIDITY_MODEL_POOLING",
	1: "LIQUIDITY_MODEL_TRANSACTIONAL",
}

var LiquidityModel_value = map[string]int32{
	"LIQUIDITY_MODEL_POOLING":       0,
	"LIQUIDITY_MODEL_TRANSACTIONAL": 1,
}

func (x LiquidityModel) String() string {
	return proto.EnumName(LiquidityModel_name, int32(x))
}

func (LiquidityModel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d53a37c7b7ae195, []int{0}
}

type Status int32

const (
	Unspecified Status = 0
	Activated   Status = 1
	Deactivated Status = 2
)

var Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_ACTIVATED",
	2: "STATUS_DEACTIVATED",
}

var Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"STATUS_ACTIVATED":   1,
	"STATUS_DEACTIVATED": 2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d53a37c7b7ae195, []int{1}
}

type ProtocolAttribute struct {
	Model LiquidityModel `protobuf:"varint,1,opt,name=model,proto3,enum=scalar.protocol.v1beta1.LiquidityModel" json:"model,omitempty"`
}

func (m *ProtocolAttribute) Reset()         { *m = ProtocolAttribute{} }
func (m *ProtocolAttribute) String() string { return proto.CompactTextString(m) }
func (*ProtocolAttribute) ProtoMessage()    {}
func (*ProtocolAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d53a37c7b7ae195, []int{0}
}
func (m *ProtocolAttribute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolAttribute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolAttribute.Merge(m, src)
}
func (m *ProtocolAttribute) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolAttribute proto.InternalMessageInfo

func (m *ProtocolAttribute) GetModel() LiquidityModel {
	if m != nil {
		return m.Model
	}
	return Pooling
}

// DestinationChain represents a blockchain where tokens can be sent
type SupportedChain struct {
	ChainName                 string                   `protobuf:"bytes,1,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	ChainId                   uint64                   `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ChainType                 string                   `protobuf:"bytes,3,opt,name=chain_type,json=chainType,proto3" json:"chain_type,omitempty"`
	ChainSmartContractAddress []byte                   `protobuf:"bytes,4,opt,name=chain_smart_contract_address,json=chainSmartContractAddress,proto3" json:"chain_smart_contract_address,omitempty"`
	Token                     types.ERC20TokenMetadata `protobuf:"bytes,5,opt,name=token,proto3" json:"token"`
}

func (m *SupportedChain) Reset()         { *m = SupportedChain{} }
func (m *SupportedChain) String() string { return proto.CompactTextString(m) }
func (*SupportedChain) ProtoMessage()    {}
func (*SupportedChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d53a37c7b7ae195, []int{1}
}
func (m *SupportedChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SupportedChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SupportedChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SupportedChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportedChain.Merge(m, src)
}
func (m *SupportedChain) XXX_Size() int {
	return m.Size()
}
func (m *SupportedChain) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportedChain.DiscardUnknown(m)
}

var xxx_messageInfo_SupportedChain proto.InternalMessageInfo

func (m *SupportedChain) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *SupportedChain) GetChainId() uint64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *SupportedChain) GetChainType() string {
	if m != nil {
		return m.ChainType
	}
	return ""
}

func (m *SupportedChain) GetChainSmartContractAddress() []byte {
	if m != nil {
		return m.ChainSmartContractAddress
	}
	return nil
}

func (m *SupportedChain) GetToken() types.ERC20TokenMetadata {
	if m != nil {
		return m.Token
	}
	return types.ERC20TokenMetadata{}
}

type Protocol struct {
	Pubkey         []byte                 `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tag            string                 `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	Attribute      *ProtocolAttribute     `protobuf:"bytes,4,opt,name=attribute,proto3" json:"attribute,omitempty"`
	Status         Status                 `protobuf:"varint,5,opt,name=status,proto3,enum=scalar.protocol.v1beta1.Status" json:"status,omitempty"`
	CustodianGroup *types1.CustodianGroup `protobuf:"bytes,6,opt,name=custodian_group,json=custodianGroup,proto3" json:"custodian_group,omitempty"`
	Chains         []*SupportedChain      `protobuf:"bytes,7,rep,name=chains,proto3" json:"chains,omitempty"`
}

func (m *Protocol) Reset()         { *m = Protocol{} }
func (m *Protocol) String() string { return proto.CompactTextString(m) }
func (*Protocol) ProtoMessage()    {}
func (*Protocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d53a37c7b7ae195, []int{2}
}
func (m *Protocol) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Protocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Protocol.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Protocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Protocol.Merge(m, src)
}
func (m *Protocol) XXX_Size() int {
	return m.Size()
}
func (m *Protocol) XXX_DiscardUnknown() {
	xxx_messageInfo_Protocol.DiscardUnknown(m)
}

var xxx_messageInfo_Protocol proto.InternalMessageInfo

func (m *Protocol) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *Protocol) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Protocol) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Protocol) GetAttribute() *ProtocolAttribute {
	if m != nil {
		return m.Attribute
	}
	return nil
}

func (m *Protocol) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Unspecified
}

func (m *Protocol) GetCustodianGroup() *types1.CustodianGroup {
	if m != nil {
		return m.CustodianGroup
	}
	return nil
}

func (m *Protocol) GetChains() []*SupportedChain {
	if m != nil {
		return m.Chains
	}
	return nil
}

func init() {
	proto.RegisterEnum("scalar.protocol.v1beta1.LiquidityModel", LiquidityModel_name, LiquidityModel_value)
	proto.RegisterEnum("scalar.protocol.v1beta1.Status", Status_name, Status_value)
	proto.RegisterType((*ProtocolAttribute)(nil), "scalar.protocol.v1beta1.ProtocolAttribute")
	proto.RegisterType((*SupportedChain)(nil), "scalar.protocol.v1beta1.SupportedChain")
	proto.RegisterType((*Protocol)(nil), "scalar.protocol.v1beta1.Protocol")
}

func init() {
	proto.RegisterFile("scalar/protocol/v1beta1/types.proto", fileDescriptor_1d53a37c7b7ae195)
}

var fileDescriptor_1d53a37c7b7ae195 = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xc1, 0x4e, 0xdb, 0x4c,
	0x10, 0xc7, 0xe3, 0x10, 0x02, 0xd9, 0x40, 0x08, 0xab, 0x4f, 0x1f, 0xc1, 0x2a, 0xc6, 0x02, 0xa9,
	0x44, 0x48, 0x4d, 0x4a, 0x5a, 0xa9, 0xa7, 0x0a, 0x99, 0x24, 0xa5, 0x56, 0x43, 0x92, 0xda, 0x4e,
	0xa5, 0xf6, 0x62, 0x6d, 0xec, 0xad, 0xb1, 0x88, 0xbd, 0xae, 0xbd, 0x8e, 0x9a, 0x07, 0xa8, 0x54,
	0xe5, 0xd4, 0x17, 0xc8, 0xa9, 0x2f, 0xc3, 0x91, 0x63, 0x4f, 0xa8, 0x82, 0x7b, 0x9f, 0xa1, 0xf2,
	0xda, 0x26, 0x50, 0xe0, 0x36, 0xb3, 0xfb, 0x9b, 0xbf, 0xe7, 0x3f, 0x3b, 0x09, 0xd8, 0x0d, 0x0c,
	0x34, 0x42, 0x7e, 0xdd, 0xf3, 0x09, 0x25, 0x06, 0x19, 0xd5, 0xc7, 0x07, 0x43, 0x4c, 0xd1, 0x41,
	0x9d, 0x4e, 0x3c, 0x1c, 0xd4, 0xd8, 0x31, 0xdc, 0x88, 0xa1, 0x5a, 0x0a, 0xd5, 0x12, 0x88, 0xff,
	0xcf, 0x22, 0x16, 0x61, 0xa7, 0xf5, 0x28, 0x8a, 0x01, 0x3e, 0xd5, 0x34, 0xc8, 0x18, 0xbb, 0xc8,
	0xa5, 0x0f, 0x69, 0xf2, 0x42, 0x02, 0xe1, 0xb1, 0xf3, 0xd0, 0xfd, 0x8e, 0x02, 0xd6, 0xfb, 0xc9,
	0xe7, 0x24, 0x4a, 0x7d, 0x7b, 0x18, 0x52, 0x0c, 0x5f, 0x83, 0x45, 0x87, 0x98, 0x78, 0x54, 0xe1,
	0x44, 0xae, 0x5a, 0x6a, 0xec, 0xd5, 0x1e, 0x69, 0xac, 0xd6, 0xb1, 0xbf, 0x84, 0xb6, 0x69, 0xd3,
	0xc9, 0x49, 0x84, 0x2b, 0x71, 0xd5, 0xce, 0x1f, 0x0e, 0x94, 0xd4, 0xd0, 0xf3, 0x88, 0x4f, 0xb1,
	0xd9, 0x3c, 0x45, 0xb6, 0x0b, 0xb7, 0x00, 0x30, 0xa2, 0x40, 0x77, 0x91, 0x83, 0x99, 0x6c, 0x41,
	0x29, 0xb0, 0x93, 0x2e, 0x72, 0x30, 0xdc, 0x04, 0xcb, 0xf1, 0xb5, 0x6d, 0x56, 0xb2, 0x22, 0x57,
	0xcd, 0x29, 0x4b, 0x2c, 0x97, 0xcd, 0x79, 0x65, 0xd4, 0x75, 0x65, 0xe1, 0x56, 0xa5, 0x36, 0xf1,
	0x30, 0x3c, 0x04, 0x4f, 0xe2, 0xeb, 0xc0, 0x41, 0x3e, 0xd5, 0x0d, 0xe2, 0x52, 0x1f, 0x19, 0x54,
	0x47, 0xa6, 0xe9, 0xe3, 0x20, 0xa8, 0xe4, 0x44, 0xae, 0xba, 0xa2, 0x6c, 0x32, 0x46, 0x8d, 0x90,
	0x66, 0x42, 0x48, 0x31, 0x00, 0x8f, 0xc0, 0x22, 0x25, 0x67, 0xd8, 0xad, 0x2c, 0x8a, 0x5c, 0xb5,
	0xd8, 0x78, 0x9a, 0x7a, 0xc5, 0x63, 0xe7, 0xc6, 0x66, 0x5b, 0x69, 0x36, 0x9e, 0x6b, 0x11, 0x75,
	0x82, 0x29, 0x32, 0x11, 0x45, 0x47, 0xb9, 0xf3, 0xcb, 0xed, 0x8c, 0x12, 0x97, 0xee, 0x5c, 0x66,
	0xc1, 0x72, 0x3a, 0x45, 0xf8, 0x3f, 0xc8, 0x7b, 0xe1, 0xf0, 0x0c, 0x4f, 0x98, 0xcd, 0x15, 0x25,
	0xc9, 0x20, 0x04, 0x39, 0x66, 0x3e, 0xcb, 0x2c, 0xb0, 0x18, 0x96, 0xc1, 0x02, 0x45, 0x56, 0xe2,
	0x2a, 0x0a, 0xe1, 0x5b, 0x50, 0x40, 0xe9, 0x3b, 0xb0, 0xe6, 0x8b, 0x8d, 0xfd, 0x47, 0xc7, 0x7f,
	0xef, 0xe5, 0x94, 0x79, 0x31, 0x7c, 0x05, 0xf2, 0x01, 0x45, 0x34, 0x0c, 0x98, 0xb3, 0x52, 0x63,
	0xfb, 0x51, 0x19, 0x95, 0x61, 0x4a, 0x82, 0xc3, 0x3e, 0x58, 0x33, 0xc2, 0x80, 0x12, 0xd3, 0x46,
	0xae, 0x6e, 0xf9, 0x24, 0xf4, 0x2a, 0x79, 0xd6, 0xc8, 0xcd, 0x1e, 0xa4, 0x1b, 0x77, 0xa3, 0xd0,
	0x4c, 0xf9, 0xe3, 0x08, 0x57, 0x4a, 0xc6, 0x9d, 0x1c, 0x1e, 0x82, 0x3c, 0x7b, 0x80, 0xa0, 0xb2,
	0x24, 0x2e, 0xdc, 0x16, 0xba, 0xdf, 0xca, 0x9d, 0xb5, 0x51, 0x92, 0xb2, 0xfd, 0x31, 0x28, 0xdd,
	0x5d, 0x35, 0x58, 0x05, 0x1b, 0x1d, 0xf9, 0xfd, 0x40, 0x6e, 0xc9, 0xda, 0x47, 0xfd, 0xa4, 0xd7,
	0x6a, 0x77, 0xf4, 0x7e, 0xaf, 0xd7, 0x91, 0xbb, 0xc7, 0xe5, 0x0c, 0x5f, 0x9c, 0xce, 0xc4, 0xa5,
	0x3e, 0x21, 0x23, 0xdb, 0xb5, 0xe0, 0x4b, 0xb0, 0xf5, 0x2f, 0xa9, 0x29, 0x52, 0x57, 0x95, 0x9a,
	0x9a, 0xdc, 0xeb, 0x4a, 0x9d, 0x32, 0xc7, 0xaf, 0x4f, 0x67, 0xe2, 0xaa, 0xe6, 0x23, 0x37, 0x40,
	0x06, 0xb5, 0x89, 0x8b, 0x46, 0x7c, 0xee, 0xfb, 0x4f, 0x21, 0xb3, 0xff, 0x8d, 0x03, 0xf9, 0x78,
	0x3a, 0x70, 0x0f, 0x40, 0x55, 0x93, 0xb4, 0x81, 0xaa, 0x0f, 0xba, 0x6a, 0xbf, 0xdd, 0x94, 0xdf,
	0xc8, 0xed, 0x56, 0x39, 0xc3, 0xaf, 0x4d, 0x67, 0x62, 0x71, 0xe0, 0x06, 0x1e, 0x36, 0xec, 0xcf,
	0x36, 0x36, 0xe1, 0x2e, 0x28, 0x27, 0x60, 0xf4, 0x85, 0x0f, 0x92, 0xd6, 0x6e, 0x95, 0x39, 0x7e,
	0x75, 0x3a, 0x13, 0x0b, 0x92, 0x41, 0xed, 0x31, 0xa2, 0xd8, 0xbc, 0xa5, 0xd6, 0x6a, 0xcf, 0xb1,
	0x6c, 0xac, 0xd6, 0xc2, 0x28, 0x05, 0xe3, 0x3e, 0x8e, 0xde, 0x9d, 0x5f, 0x09, 0xdc, 0xc5, 0x95,
	0xc0, 0xfd, 0xbe, 0x12, 0xb8, 0x1f, 0xd7, 0x42, 0xe6, 0xe2, 0x5a, 0xc8, 0xfc, 0xba, 0x16, 0x32,
	0x9f, 0x0e, 0x2c, 0x9b, 0x9e, 0x86, 0xc3, 0x9a, 0x41, 0x9c, 0x7a, 0x3c, 0x54, 0xe2, 0x5b, 0x49,
	0xf4, 0xcc, 0x20, 0x3e, 0xae, 0x7f, 0x9d, 0xff, 0xe9, 0xb0, 0x1f, 0xfe, 0x30, 0xcf, 0xf2, 0x17,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x81, 0x0e, 0x09, 0x74, 0x94, 0x04, 0x00, 0x00,
}

func (m *ProtocolAttribute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolAttribute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolAttribute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Model != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Model))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SupportedChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SupportedChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupportedChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.ChainSmartContractAddress) > 0 {
		i -= len(m.ChainSmartContractAddress)
		copy(dAtA[i:], m.ChainSmartContractAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChainSmartContractAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ChainType) > 0 {
		i -= len(m.ChainType)
		copy(dAtA[i:], m.ChainType)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChainType)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ChainId != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Protocol) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Protocol) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Protocol) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Chains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.CustodianGroup != nil {
		{
			size, err := m.CustodianGroup.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.Status != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if m.Attribute != nil {
		{
			size, err := m.Attribute.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Tag) > 0 {
		i -= len(m.Tag)
		copy(dAtA[i:], m.Tag)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Tag)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProtocolAttribute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Model != 0 {
		n += 1 + sovTypes(uint64(m.Model))
	}
	return n
}

func (m *SupportedChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ChainId != 0 {
		n += 1 + sovTypes(uint64(m.ChainId))
	}
	l = len(m.ChainType)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ChainSmartContractAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Token.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *Protocol) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Tag)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Attribute != nil {
		l = m.Attribute.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovTypes(uint64(m.Status))
	}
	if m.CustodianGroup != nil {
		l = m.CustodianGroup.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Chains) > 0 {
		for _, e := range m.Chains {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProtocolAttribute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ProtocolAttribute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolAttribute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Model", wireType)
			}
			m.Model = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Model |= LiquidityModel(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *SupportedChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SupportedChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SupportedChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainSmartContractAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainSmartContractAddress = append(m.ChainSmartContractAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ChainSmartContractAddress == nil {
				m.ChainSmartContractAddress = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Protocol) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Protocol: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Protocol: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = append(m.Pubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.Pubkey == nil {
				m.Pubkey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attribute", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Attribute == nil {
				m.Attribute = &ProtocolAttribute{}
			}
			if err := m.Attribute.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustodianGroup", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CustodianGroup == nil {
				m.CustodianGroup = &types1.CustodianGroup{}
			}
			if err := m.CustodianGroup.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, &SupportedChain{})
			if err := m.Chains[len(m.Chains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
