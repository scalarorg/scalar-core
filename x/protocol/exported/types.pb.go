// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/protocol/exported/v1beta1/types.proto

package exported

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/scalarorg/scalar-core/x/covenant/types"
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

type LiquidityModel int32

const (
	LIQUIDITY_MODEL_UNSPECIFIED LiquidityModel = 0
	LIQUIDITY_MODEL_POOL        LiquidityModel = 1
	LIQUIDITY_MODEL_UPC         LiquidityModel = 2
)

var LiquidityModel_name = map[int32]string{
	0: "LIQUIDITY_MODEL_UNSPECIFIED",
	1: "LIQUIDITY_MODEL_POOL",
	2: "LIQUIDITY_MODEL_UPC",
}

var LiquidityModel_value = map[string]int32{
	"LIQUIDITY_MODEL_UNSPECIFIED": 0,
	"LIQUIDITY_MODEL_POOL":        1,
	"LIQUIDITY_MODEL_UPC":         2,
}

func (x LiquidityModel) String() string {
	return proto.EnumName(LiquidityModel_name, int32(x))
}

func (LiquidityModel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_417dfcc53daddf21, []int{0}
}

type Status int32

const (
	Unspecified Status = 0
	Activated   Status = 1
	Deactivated Status = 2
	Pending     Status = 3
)

var Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "STATUS_ACTIVATED",
	2: "STATUS_DEACTIVATED",
	3: "STATUS_PENDING",
}

var Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"STATUS_ACTIVATED":   1,
	"STATUS_DEACTIVATED": 2,
	"STATUS_PENDING":     3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_417dfcc53daddf21, []int{1}
}

type ProtocolAttributes struct {
	Model LiquidityModel `protobuf:"varint,1,opt,name=model,proto3,enum=scalar.protocol.exported.v1beta1.LiquidityModel" json:"model,omitempty"`
}

func (m *ProtocolAttributes) Reset()         { *m = ProtocolAttributes{} }
func (m *ProtocolAttributes) String() string { return proto.CompactTextString(m) }
func (*ProtocolAttributes) ProtoMessage()    {}
func (*ProtocolAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_417dfcc53daddf21, []int{0}
}
func (m *ProtocolAttributes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolAttributes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolAttributes.Merge(m, src)
}
func (m *ProtocolAttributes) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolAttributes proto.InternalMessageInfo

type MinorAddress struct {
	ChainName github_com_scalarorg_scalar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain_name,json=chainName,proto3,casttype=github.com/scalarorg/scalar-core/x/nexus/exported.ChainName" json:"chain_name,omitempty"`
	Address   string                                                      `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MinorAddress) Reset()         { *m = MinorAddress{} }
func (m *MinorAddress) String() string { return proto.CompactTextString(m) }
func (*MinorAddress) ProtoMessage()    {}
func (*MinorAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_417dfcc53daddf21, []int{1}
}
func (m *MinorAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MinorAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MinorAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MinorAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinorAddress.Merge(m, src)
}
func (m *MinorAddress) XXX_Size() int {
	return m.Size()
}
func (m *MinorAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_MinorAddress.DiscardUnknown(m)
}

var xxx_messageInfo_MinorAddress proto.InternalMessageInfo

type SupportedChain struct {
	Chain   github_com_scalarorg_scalar_core_x_nexus_exported.ChainName `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/scalarorg/scalar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
	Name    string                                                      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string                                                      `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *SupportedChain) Reset()         { *m = SupportedChain{} }
func (m *SupportedChain) String() string { return proto.CompactTextString(m) }
func (*SupportedChain) ProtoMessage()    {}
func (*SupportedChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_417dfcc53daddf21, []int{2}
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

type ProtocolInfo struct {
	// string key_id = 1 [
	//   (gogoproto.customname) = "KeyID",
	//   (gogoproto.casttype) =
	//       "github.com/scalarorg/scalar-core/x/multisig/exported.KeyID"
	// ];
	// bytes custodians_pubkey = 2 [ (gogoproto.customname) = "CustodiansPubkey"
	// ];
	CustodiansGroupUID string                                                      `protobuf:"bytes,1,opt,name=custodians_group_uid,json=custodiansGroupUid,proto3" json:"custodians_group_uid,omitempty"`
	LiquidityModel     LiquidityModel                                              `protobuf:"varint,2,opt,name=liquidity_model,json=liquidityModel,proto3,enum=scalar.protocol.exported.v1beta1.LiquidityModel" json:"liquidity_model,omitempty"`
	Symbol             string                                                      `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	OriginChain        github_com_scalarorg_scalar_core_x_nexus_exported.ChainName `protobuf:"bytes,4,opt,name=origin_chain,json=originChain,proto3,casttype=github.com/scalarorg/scalar-core/x/nexus/exported.ChainName" json:"origin_chain,omitempty"`
	MinorAddresses     []*MinorAddress                                             `protobuf:"bytes,5,rep,name=minor_addresses,json=minorAddresses,proto3" json:"minor_addresses,omitempty"`
}

func (m *ProtocolInfo) Reset()         { *m = ProtocolInfo{} }
func (m *ProtocolInfo) String() string { return proto.CompactTextString(m) }
func (*ProtocolInfo) ProtoMessage()    {}
func (*ProtocolInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_417dfcc53daddf21, []int{3}
}
func (m *ProtocolInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolInfo.Merge(m, src)
}
func (m *ProtocolInfo) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolInfo proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("scalar.protocol.exported.v1beta1.LiquidityModel", LiquidityModel_name, LiquidityModel_value)
	proto.RegisterEnum("scalar.protocol.exported.v1beta1.Status", Status_name, Status_value)
	proto.RegisterType((*ProtocolAttributes)(nil), "scalar.protocol.exported.v1beta1.ProtocolAttributes")
	proto.RegisterType((*MinorAddress)(nil), "scalar.protocol.exported.v1beta1.MinorAddress")
	proto.RegisterType((*SupportedChain)(nil), "scalar.protocol.exported.v1beta1.SupportedChain")
	proto.RegisterType((*ProtocolInfo)(nil), "scalar.protocol.exported.v1beta1.ProtocolInfo")
}

func init() {
	proto.RegisterFile("scalar/protocol/exported/v1beta1/types.proto", fileDescriptor_417dfcc53daddf21)
}

var fileDescriptor_417dfcc53daddf21 = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x8d, 0x93, 0xfe, 0x51, 0x36, 0xfd, 0xa5, 0xd1, 0xfe, 0xaa, 0x12, 0x05, 0xc9, 0x8e, 0xda,
	0x03, 0x55, 0x05, 0x36, 0x2d, 0xdc, 0x38, 0xa0, 0x34, 0x4e, 0x8b, 0xa5, 0x34, 0x09, 0xf9, 0x83,
	0x04, 0x42, 0xb2, 0x1c, 0x7b, 0xeb, 0x2e, 0x8d, 0xbd, 0xa9, 0x77, 0x5d, 0xb5, 0xdf, 0x00, 0xf5,
	0xc4, 0x85, 0x63, 0x4f, 0x48, 0x7c, 0x0d, 0xae, 0x3d, 0xf6, 0xc8, 0x29, 0x02, 0xf7, 0x5b, 0x70,
	0x42, 0xf6, 0xda, 0x69, 0x92, 0x22, 0x15, 0xd1, 0xdb, 0xcc, 0xec, 0x7b, 0x6f, 0xde, 0xcc, 0x58,
	0x06, 0x8f, 0xa9, 0x69, 0x0c, 0x0c, 0x4f, 0x19, 0x7a, 0x84, 0x11, 0x93, 0x0c, 0x14, 0x74, 0x3a,
	0x24, 0x1e, 0x43, 0x96, 0x72, 0xb2, 0xd5, 0x47, 0xcc, 0xd8, 0x52, 0xd8, 0xd9, 0x10, 0x51, 0x39,
	0x7a, 0x87, 0x65, 0x8e, 0x96, 0x13, 0xb4, 0x9c, 0xa0, 0xe5, 0x18, 0x5d, 0x5a, 0x8f, 0xf5, 0x4c,
	0x72, 0x82, 0x5c, 0xc3, 0x65, 0x7f, 0x92, 0x29, 0xad, 0xd8, 0xc4, 0x26, 0x51, 0xa8, 0x84, 0x11,
	0xaf, 0xae, 0xbd, 0x07, 0xb0, 0x15, 0xeb, 0x56, 0x18, 0xf3, 0x70, 0xdf, 0x67, 0x88, 0xc2, 0x5d,
	0x30, 0xef, 0x10, 0x0b, 0x0d, 0x8a, 0x42, 0x59, 0xd8, 0xc8, 0x6f, 0x3f, 0x95, 0xef, 0xb2, 0x20,
	0xd7, 0xf1, 0xb1, 0x8f, 0x2d, 0xcc, 0xce, 0xf6, 0x43, 0x5e, 0x9b, 0xd3, 0xd7, 0x3e, 0x0b, 0x60,
	0x69, 0x1f, 0xbb, 0xc4, 0xab, 0x58, 0x96, 0x87, 0x28, 0x85, 0x47, 0x00, 0x98, 0x87, 0x06, 0x76,
	0x75, 0xd7, 0x70, 0x50, 0xa4, 0x9e, 0xdd, 0xa9, 0x07, 0x23, 0x29, 0x5b, 0x0d, 0xab, 0x0d, 0xc3,
	0x41, 0xbf, 0x46, 0xd2, 0x0b, 0x1b, 0xb3, 0x43, 0xbf, 0x2f, 0x9b, 0xc4, 0x51, 0x78, 0x63, 0xe2,
	0xd9, 0x71, 0xf4, 0xc4, 0x24, 0x1e, 0x52, 0x4e, 0x15, 0x17, 0x9d, 0xfa, 0x74, 0xbc, 0x37, 0x79,
	0x4c, 0x6f, 0x67, 0xcd, 0x24, 0x84, 0x45, 0xb0, 0x68, 0xf0, 0xbe, 0xc5, 0x74, 0xd8, 0xa9, 0x9d,
	0xa4, 0xa1, 0xaf, 0x7c, 0xc7, 0x1f, 0x72, 0x72, 0xc4, 0x85, 0x3d, 0x30, 0x1f, 0x31, 0x63, 0x53,
	0x2f, 0xef, 0xeb, 0x83, 0xab, 0x41, 0x08, 0xe6, 0xa2, 0x51, 0xb9, 0x81, 0x28, 0x9e, 0xf4, 0x95,
	0x99, 0xf6, 0xf5, 0x2d, 0x03, 0x96, 0x92, 0x73, 0x68, 0xee, 0x01, 0x81, 0xaf, 0xc0, 0x8a, 0xe9,
	0x53, 0x46, 0x2c, 0x6c, 0xb8, 0x54, 0xb7, 0x3d, 0xe2, 0x0f, 0x75, 0x1f, 0x5b, 0xb1, 0xc9, 0xd5,
	0x60, 0x24, 0xc1, 0xea, 0xf8, 0x7d, 0x2f, 0x7c, 0xee, 0x69, 0x6a, 0x1b, 0x9a, 0x33, 0x35, 0x6c,
	0x41, 0x07, 0x2c, 0x0f, 0x92, 0x1b, 0xe9, 0xfc, 0xb8, 0xe9, 0x7f, 0x3b, 0xee, 0x0e, 0x0c, 0x46,
	0x52, 0x7e, 0xe6, 0xe0, 0xf9, 0xc1, 0x54, 0x0e, 0x57, 0xc1, 0x02, 0x3d, 0x73, 0xfa, 0x64, 0x10,
	0x8f, 0x18, 0x67, 0xf0, 0x18, 0x2c, 0x11, 0x0f, 0xdb, 0xd8, 0xd5, 0xf9, 0xb6, 0xe7, 0xa2, 0x41,
	0x1a, 0xc1, 0x48, 0xca, 0x35, 0xa3, 0x7a, 0xb4, 0xc1, 0xfb, 0x2e, 0x3f, 0x47, 0x6e, 0xb4, 0xe0,
	0x11, 0x58, 0x76, 0xc2, 0x6f, 0x50, 0x8f, 0xb7, 0x8c, 0x68, 0x71, 0xbe, 0x9c, 0xd9, 0xc8, 0x6d,
	0xcb, 0x77, 0x4f, 0x3e, 0xf9, 0xf1, 0xf2, 0xb9, 0x27, 0x2b, 0x88, 0xb6, 0xf3, 0xce, 0x54, 0xbe,
	0xf9, 0x01, 0xcc, 0x6c, 0x06, 0x4a, 0xe0, 0x61, 0x5d, 0x7b, 0xdd, 0xd3, 0x54, 0xad, 0xfb, 0x56,
	0xdf, 0x6f, 0xaa, 0xb5, 0xba, 0xde, 0x6b, 0x74, 0x5a, 0xb5, 0xaa, 0xb6, 0xab, 0xd5, 0xd4, 0x42,
	0x0a, 0x16, 0xc1, 0xca, 0x2c, 0xa0, 0xd5, 0x6c, 0xd6, 0x0b, 0x02, 0x7c, 0x00, 0xfe, 0xbf, 0x45,
	0x6d, 0x55, 0x0b, 0xe9, 0xd2, 0xdc, 0xc7, 0x2f, 0x62, 0x6a, 0xf3, 0xab, 0x00, 0x16, 0x3a, 0xcc,
	0x60, 0x3e, 0x85, 0x8f, 0x00, 0xec, 0x74, 0x2b, 0xdd, 0x5e, 0x67, 0x5a, 0xbb, 0xb4, 0x7c, 0x7e,
	0x51, 0xce, 0xf5, 0x5c, 0x3a, 0x44, 0x26, 0x3e, 0xc0, 0xc8, 0x82, 0xeb, 0xa0, 0x10, 0x03, 0x2b,
	0xd5, 0xae, 0xf6, 0xa6, 0xd2, 0xad, 0xa9, 0x05, 0xa1, 0xf4, 0xdf, 0xf9, 0x45, 0x39, 0x5b, 0x31,
	0x19, 0x3e, 0x31, 0x18, 0xb2, 0x26, 0xd4, 0xd4, 0xda, 0x0d, 0x2c, 0xcd, 0xd5, 0x54, 0x64, 0x8c,
	0x81, 0x12, 0xc8, 0xc7, 0xc0, 0x56, 0xad, 0xa1, 0x6a, 0x8d, 0xbd, 0x42, 0xa6, 0x94, 0x3b, 0xbf,
	0x28, 0x2f, 0xb6, 0x90, 0x6b, 0x61, 0xd7, 0xe6, 0x46, 0x77, 0xda, 0x97, 0x3f, 0xc5, 0xd4, 0x65,
	0x20, 0x0a, 0x57, 0x81, 0x28, 0xfc, 0x08, 0x44, 0xe1, 0xd3, 0xb5, 0x98, 0xba, 0xba, 0x16, 0x53,
	0xdf, 0xaf, 0xc5, 0xd4, 0xbb, 0xe7, 0x7f, 0x71, 0xe9, 0x5b, 0x7f, 0xca, 0xfe, 0x42, 0x54, 0x7a,
	0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x43, 0x7d, 0x51, 0x4c, 0x05, 0x00, 0x00,
}

func (m *ProtocolAttributes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolAttributes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolAttributes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *MinorAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MinorAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MinorAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
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
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
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
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProtocolInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinorAddresses) > 0 {
		for iNdEx := len(m.MinorAddresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinorAddresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.OriginChain) > 0 {
		i -= len(m.OriginChain)
		copy(dAtA[i:], m.OriginChain)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.OriginChain)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if m.LiquidityModel != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.LiquidityModel))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CustodiansGroupUID) > 0 {
		i -= len(m.CustodiansGroupUID)
		copy(dAtA[i:], m.CustodiansGroupUID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CustodiansGroupUID)))
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
func (m *ProtocolAttributes) Size() (n int) {
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

func (m *MinorAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *SupportedChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ProtocolInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CustodiansGroupUID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.LiquidityModel != 0 {
		n += 1 + sovTypes(uint64(m.LiquidityModel))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.OriginChain)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.MinorAddresses) > 0 {
		for _, e := range m.MinorAddresses {
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
func (m *ProtocolAttributes) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ProtocolAttributes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolAttributes: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MinorAddress) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MinorAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MinorAddress: illegal tag %d (wire type %d)", fieldNum, wire)
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
			m.ChainName = github_com_scalarorg_scalar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
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
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
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
			m.Chain = github_com_scalarorg_scalar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
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
func (m *ProtocolInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ProtocolInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustodiansGroupUID", wireType)
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
			m.CustodiansGroupUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityModel", wireType)
			}
			m.LiquidityModel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LiquidityModel |= LiquidityModel(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
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
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginChain", wireType)
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
			m.OriginChain = github_com_scalarorg_scalar_core_x_nexus_exported.ChainName(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinorAddresses", wireType)
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
			m.MinorAddresses = append(m.MinorAddresses, &MinorAddress{})
			if err := m.MinorAddresses[len(m.MinorAddresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
