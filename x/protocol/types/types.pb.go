// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/protocol/v1beta1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types2 "github.com/scalarorg/scalar-core/x/chains/btc/types"
	types1 "github.com/scalarorg/scalar-core/x/chains/evm/types"
	types "github.com/scalarorg/scalar-core/x/chains/types"
	types3 "github.com/scalarorg/scalar-core/x/covenant/types"
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
	Params  *types.Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Address string        `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are valid to be assigned to Token:
	//	*SupportedChain_Erc20
	//	*SupportedChain_Btc
	Token isSupportedChain_Token `protobuf_oneof:"token"`
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

type isSupportedChain_Token interface {
	isSupportedChain_Token()
	MarshalTo([]byte) (int, error)
	Size() int
}

type SupportedChain_Erc20 struct {
	Erc20 *types1.ERC20TokenMetadata `protobuf:"bytes,3,opt,name=erc20,proto3,oneof" json:"erc20,omitempty"`
}
type SupportedChain_Btc struct {
	Btc *types2.BtcToken `protobuf:"bytes,4,opt,name=btc,proto3,oneof" json:"btc,omitempty"`
}

func (*SupportedChain_Erc20) isSupportedChain_Token() {}
func (*SupportedChain_Btc) isSupportedChain_Token()   {}

func (m *SupportedChain) GetToken() isSupportedChain_Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *SupportedChain) GetParams() *types.Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *SupportedChain) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SupportedChain) GetErc20() *types1.ERC20TokenMetadata {
	if x, ok := m.GetToken().(*SupportedChain_Erc20); ok {
		return x.Erc20
	}
	return nil
}

func (m *SupportedChain) GetBtc() *types2.BtcToken {
	if x, ok := m.GetToken().(*SupportedChain_Btc); ok {
		return x.Btc
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SupportedChain) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SupportedChain_Erc20)(nil),
		(*SupportedChain_Btc)(nil),
	}
}

type Protocol struct {
	Pubkey         []byte                 `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Address        []byte                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Name           string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Tag            string                 `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	Attribute      *ProtocolAttribute     `protobuf:"bytes,5,opt,name=attribute,proto3" json:"attribute,omitempty"`
	Status         Status                 `protobuf:"varint,6,opt,name=status,proto3,enum=scalar.protocol.v1beta1.Status" json:"status,omitempty"`
	CustodianGroup *types3.CustodianGroup `protobuf:"bytes,7,opt,name=custodian_group,json=custodianGroup,proto3" json:"custodian_group,omitempty"`
	Chains         []*SupportedChain      `protobuf:"bytes,8,rep,name=chains,proto3" json:"chains,omitempty"`
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

func (m *Protocol) GetAddress() []byte {
	if m != nil {
		return m.Address
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

func (m *Protocol) GetCustodianGroup() *types3.CustodianGroup {
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
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x41, 0x6f, 0xda, 0x48,
	0x14, 0x80, 0xed, 0x10, 0x20, 0x4c, 0x12, 0x42, 0x46, 0xab, 0x8d, 0x65, 0x29, 0x2c, 0x4b, 0xb4,
	0x4a, 0x14, 0x29, 0x90, 0xb0, 0xad, 0x72, 0xaa, 0x2a, 0x02, 0x34, 0x41, 0x25, 0x40, 0x07, 0x53,
	0xa9, 0xbd, 0xa0, 0xf1, 0x78, 0x4a, 0xac, 0x80, 0xc7, 0xb5, 0xc7, 0xa8, 0xf9, 0x01, 0x95, 0x2a,
	0x4e, 0xfd, 0x03, 0x9c, 0xfa, 0x67, 0x7a, 0xcc, 0xb1, 0xc7, 0x2a, 0xb9, 0xf5, 0x17, 0xf4, 0x58,
	0x79, 0x6c, 0x43, 0x20, 0xe4, 0x36, 0xcf, 0xfa, 0xde, 0xe7, 0xf7, 0xde, 0x3c, 0x0d, 0xd8, 0x73,
	0x09, 0x1e, 0x60, 0xa7, 0x68, 0x3b, 0x8c, 0x33, 0xc2, 0x06, 0xc5, 0xd1, 0x89, 0x4e, 0x39, 0x3e,
	0x29, 0xf2, 0x1b, 0x9b, 0xba, 0x05, 0xf1, 0x19, 0xee, 0x04, 0x50, 0x21, 0x82, 0x0a, 0x21, 0xa4,
	0xfe, 0xd5, 0x67, 0x7d, 0x26, 0xbe, 0x16, 0xfd, 0x53, 0x00, 0xa8, 0x91, 0x93, 0xb0, 0x11, 0xb5,
	0xb0, 0xc5, 0x97, 0x39, 0xd5, 0x7f, 0x23, 0xe8, 0x0a, 0x9b, 0x96, 0xbb, 0x14, 0xc9, 0x2f, 0x47,
	0x6c, 0xec, 0xe0, 0x61, 0xc4, 0xfc, 0x37, 0xcf, 0xd0, 0xd1, 0x70, 0xa9, 0x6a, 0x01, 0xd3, 0x39,
	0x59, 0x86, 0xe5, 0x11, 0xd8, 0x6e, 0x87, 0x3d, 0x96, 0x39, 0x77, 0x4c, 0xdd, 0xe3, 0x14, 0xbe,
	0x00, 0xf1, 0x21, 0x33, 0xe8, 0x40, 0x91, 0x73, 0xf2, 0x41, 0xba, 0xb4, 0x5f, 0x78, 0x62, 0x1a,
	0x85, 0x86, 0xf9, 0xd1, 0x33, 0x0d, 0x93, 0xdf, 0x5c, 0xfa, 0x38, 0x0a, 0xb2, 0xf2, 0xbf, 0x64,
	0x90, 0xee, 0x78, 0xb6, 0xcd, 0x1c, 0x4e, 0x8d, 0x8a, 0xff, 0x7f, 0xf8, 0x1c, 0x24, 0x82, 0x26,
	0x84, 0x72, 0xbd, 0xb4, 0x1b, 0x29, 0x83, 0xf2, 0xa6, 0xc2, 0xb6, 0x80, 0x50, 0x08, 0x43, 0x05,
	0x24, 0xb1, 0x61, 0x38, 0xd4, 0x75, 0x95, 0x95, 0x9c, 0x7c, 0x90, 0x42, 0x51, 0x08, 0x6b, 0x20,
	0x4e, 0x1d, 0x52, 0x3a, 0x56, 0x62, 0xc2, 0x77, 0xb4, 0xe0, 0xa3, 0xa3, 0xe1, 0xd4, 0x59, 0x43,
	0x95, 0xd2, 0xb1, 0xc6, 0xae, 0xa9, 0x75, 0x49, 0x39, 0x36, 0x30, 0xc7, 0x17, 0x12, 0x0a, 0xb2,
	0xe1, 0x29, 0x88, 0xe9, 0x9c, 0x28, 0xab, 0x42, 0xb2, 0xb7, 0x20, 0xd1, 0x39, 0x99, 0x4a, 0xce,
	0x38, 0x11, 0x8a, 0x0b, 0x09, 0xf9, 0x19, 0x67, 0x49, 0x10, 0xe7, 0x7e, 0x9c, 0xff, 0xbd, 0x02,
	0xd6, 0xa2, 0x09, 0xc2, 0xbf, 0x41, 0xc2, 0xf6, 0xf4, 0x6b, 0x7a, 0x23, 0xda, 0xdc, 0x40, 0x61,
	0xb4, 0xd8, 0xc7, 0xc6, 0xac, 0x0f, 0x08, 0x56, 0x2d, 0x3c, 0xa4, 0xa2, 0x8d, 0x14, 0x12, 0x67,
	0x98, 0x01, 0x31, 0x8e, 0xfb, 0xa2, 0xa8, 0x14, 0xf2, 0x8f, 0xf0, 0x02, 0xa4, 0x70, 0x74, 0x3b,
	0x4a, 0x5c, 0x14, 0x7b, 0xf8, 0xe4, 0xa5, 0x3c, 0xba, 0x4f, 0x34, 0x4b, 0x86, 0xa7, 0x20, 0xe1,
	0x72, 0xcc, 0x3d, 0x57, 0x49, 0x88, 0xbb, 0xfd, 0xe7, 0x49, 0x4d, 0x47, 0x60, 0x28, 0xc4, 0x61,
	0x1b, 0x6c, 0x11, 0xcf, 0xe5, 0xcc, 0x30, 0xb1, 0xd5, 0xeb, 0x3b, 0xcc, 0xb3, 0x95, 0xa4, 0x28,
	0x64, 0xba, 0x1d, 0xd1, 0xf2, 0x4f, 0x0d, 0x95, 0x88, 0x3f, 0xf7, 0x71, 0x94, 0x26, 0x73, 0x31,
	0x7c, 0x09, 0x12, 0xc1, 0xa0, 0x95, 0xb5, 0x5c, 0xec, 0xa1, 0xe8, 0x71, 0x29, 0x73, 0xcb, 0x84,
	0xc2, 0xb4, 0xc3, 0x11, 0x48, 0xcf, 0x2f, 0x20, 0x3c, 0x00, 0x3b, 0x8d, 0xfa, 0x9b, 0x6e, 0xbd,
	0x5a, 0xd7, 0xde, 0xf5, 0x2e, 0x5b, 0xd5, 0x5a, 0xa3, 0xd7, 0x6e, 0xb5, 0x1a, 0xf5, 0xe6, 0x79,
	0x46, 0x52, 0xd7, 0xc7, 0x93, 0x5c, 0xb2, 0xcd, 0xd8, 0xc0, 0xb4, 0xfa, 0xf0, 0x19, 0xd8, 0x5d,
	0x24, 0x35, 0x54, 0x6e, 0x76, 0xca, 0x15, 0xad, 0xde, 0x6a, 0x96, 0x1b, 0x19, 0x59, 0xdd, 0x1e,
	0x4f, 0x72, 0x9b, 0x9a, 0x83, 0x2d, 0x17, 0x13, 0x6e, 0x32, 0x0b, 0x0f, 0xd4, 0xd5, 0x2f, 0xdf,
	0xb2, 0xd2, 0xe1, 0x67, 0x19, 0x24, 0x82, 0xe9, 0xc0, 0x7d, 0x00, 0x3b, 0x5a, 0x59, 0xeb, 0x76,
	0x7a, 0xdd, 0x66, 0xa7, 0x5d, 0xab, 0xd4, 0x5f, 0xd5, 0x6b, 0xd5, 0x8c, 0xa4, 0x6e, 0x8d, 0x27,
	0xb9, 0xf5, 0xae, 0xe5, 0xda, 0x94, 0x98, 0x1f, 0x4c, 0x6a, 0xc0, 0x3d, 0x90, 0x09, 0x41, 0xff,
	0x0f, 0x6f, 0xcb, 0x5a, 0xad, 0x9a, 0x91, 0xd5, 0xcd, 0xf1, 0x24, 0x97, 0x2a, 0x13, 0x6e, 0x8e,
	0x30, 0xa7, 0xc6, 0x03, 0x5b, 0xb5, 0x36, 0xc3, 0x56, 0x02, 0x5b, 0x95, 0xe2, 0x08, 0x0c, 0xea,
	0x38, 0x7b, 0xfd, 0xfd, 0x2e, 0x2b, 0xdf, 0xde, 0x65, 0xe5, 0x9f, 0x77, 0x59, 0xf9, 0xeb, 0x7d,
	0x56, 0xba, 0xbd, 0xcf, 0x4a, 0x3f, 0xee, 0xb3, 0xd2, 0xfb, 0x93, 0xbe, 0xc9, 0xaf, 0x3c, 0xbd,
	0x40, 0xd8, 0xb0, 0x18, 0x0c, 0x95, 0x39, 0xfd, 0xf0, 0x74, 0x44, 0x98, 0x43, 0x8b, 0x9f, 0x66,
	0xef, 0x9f, 0x78, 0x0e, 0xf4, 0x84, 0x88, 0xff, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xd9,
	0x47, 0x76, 0x1f, 0x05, 0x00, 0x00,
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
	if m.Token != nil {
		{
			size := m.Token.Size()
			i -= size
			if _, err := m.Token.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SupportedChain_Erc20) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupportedChain_Erc20) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Erc20 != nil {
		{
			size, err := m.Erc20.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *SupportedChain_Btc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupportedChain_Btc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Btc != nil {
		{
			size, err := m.Btc.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
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
			dAtA[i] = 0x42
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
		dAtA[i] = 0x3a
	}
	if m.Status != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
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
		dAtA[i] = 0x2a
	}
	if len(m.Tag) > 0 {
		i -= len(m.Tag)
		copy(dAtA[i:], m.Tag)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Tag)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
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
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Token != nil {
		n += m.Token.Size()
	}
	return n
}

func (m *SupportedChain_Erc20) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Erc20 != nil {
		l = m.Erc20.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *SupportedChain_Btc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Btc != nil {
		l = m.Btc.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
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
	l = len(m.Address)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if m.Params == nil {
				m.Params = &types.Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20", wireType)
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
			v := &types1.ERC20TokenMetadata{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Token = &SupportedChain_Erc20{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Btc", wireType)
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
			v := &types2.BtcToken{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Token = &SupportedChain_Btc{v}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 3:
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
		case 4:
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
		case 5:
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
		case 6:
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
		case 7:
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
				m.CustodianGroup = &types3.CustodianGroup{}
			}
			if err := m.CustodianGroup.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
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
