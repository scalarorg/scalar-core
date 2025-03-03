// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/chains/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_scalarorg_bitcoin_vault_go_utils_types "github.com/scalarorg/bitcoin-vault/go-utils/types"
	utils "github.com/scalarorg/scalar-core/utils"
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

type Params struct {
	Chain               github_com_scalarorg_scalar_core_x_nexus_exported.ChainName   `protobuf:"bytes,1,opt,name=chain,proto3,casttype=github.com/scalarorg/scalar-core/x/nexus/exported.ChainName" json:"chain,omitempty"`
	ConfirmationHeight  uint64                                                        `protobuf:"varint,2,opt,name=confirmation_height,json=confirmationHeight,proto3" json:"confirmation_height,omitempty"`
	NetworkKind         github_com_scalarorg_bitcoin_vault_go_utils_types.NetworkKind `protobuf:"varint,3,opt,name=network_kind,json=networkKind,proto3,casttype=github.com/scalarorg/bitcoin-vault/go-utils/types.NetworkKind" json:"network_kind,omitempty"`
	TokenCode           []byte                                                        `protobuf:"bytes,4,opt,name=token_code,json=tokenCode,proto3" json:"token_code,omitempty"`
	Burnable            []byte                                                        `protobuf:"bytes,5,opt,name=burnable,proto3" json:"burnable,omitempty"`
	RevoteLockingPeriod int64                                                         `protobuf:"varint,6,opt,name=revote_locking_period,json=revoteLockingPeriod,proto3" json:"revote_locking_period,omitempty"`
	ChainID             github_com_cosmos_cosmos_sdk_types.Int                        `protobuf:"bytes,7,opt,name=chain_id,json=chainId,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"chain_id"`
	VotingThreshold     utils.Threshold                                               `protobuf:"bytes,8,opt,name=voting_threshold,json=votingThreshold,proto3" json:"voting_threshold"`
	MinVoterCount       int64                                                         `protobuf:"varint,9,opt,name=min_voter_count,json=minVoterCount,proto3" json:"min_voter_count,omitempty"`
	CommandsGasLimit    uint32                                                        `protobuf:"varint,10,opt,name=commands_gas_limit,json=commandsGasLimit,proto3" json:"commands_gas_limit,omitempty"`
	VotingGracePeriod   int64                                                         `protobuf:"varint,11,opt,name=voting_grace_period,json=votingGracePeriod,proto3" json:"voting_grace_period,omitempty"`
	EndBlockerLimit     int64                                                         `protobuf:"varint,12,opt,name=end_blocker_limit,json=endBlockerLimit,proto3" json:"end_blocker_limit,omitempty"`
	TransferLimit       uint64                                                        `protobuf:"varint,13,opt,name=transfer_limit,json=transferLimit,proto3" json:"transfer_limit,omitempty"`
	Metadata            map[string]string                                             `protobuf:"bytes,14,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_145e1080a7ed4505, []int{0}
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

func (m *Params) GetChain() github_com_scalarorg_scalar_core_x_nexus_exported.ChainName {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *Params) GetConfirmationHeight() uint64 {
	if m != nil {
		return m.ConfirmationHeight
	}
	return 0
}

func (m *Params) GetNetworkKind() github_com_scalarorg_bitcoin_vault_go_utils_types.NetworkKind {
	if m != nil {
		return m.NetworkKind
	}
	return 0
}

func (m *Params) GetTokenCode() []byte {
	if m != nil {
		return m.TokenCode
	}
	return nil
}

func (m *Params) GetBurnable() []byte {
	if m != nil {
		return m.Burnable
	}
	return nil
}

func (m *Params) GetRevoteLockingPeriod() int64 {
	if m != nil {
		return m.RevoteLockingPeriod
	}
	return 0
}

func (m *Params) GetVotingThreshold() utils.Threshold {
	if m != nil {
		return m.VotingThreshold
	}
	return utils.Threshold{}
}

func (m *Params) GetMinVoterCount() int64 {
	if m != nil {
		return m.MinVoterCount
	}
	return 0
}

func (m *Params) GetCommandsGasLimit() uint32 {
	if m != nil {
		return m.CommandsGasLimit
	}
	return 0
}

func (m *Params) GetVotingGracePeriod() int64 {
	if m != nil {
		return m.VotingGracePeriod
	}
	return 0
}

func (m *Params) GetEndBlockerLimit() int64 {
	if m != nil {
		return m.EndBlockerLimit
	}
	return 0
}

func (m *Params) GetTransferLimit() uint64 {
	if m != nil {
		return m.TransferLimit
	}
	return 0
}

func (m *Params) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "scalar.chains.v1beta1.Params")
	proto.RegisterMapType((map[string]string)(nil), "scalar.chains.v1beta1.Params.MetadataEntry")
}

func init() {
	proto.RegisterFile("scalar/chains/v1beta1/params.proto", fileDescriptor_145e1080a7ed4505)
}

var fileDescriptor_145e1080a7ed4505 = []byte{
	// 661 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0xd6, 0xad, 0x6b, 0xdd, 0x75, 0x3f, 0xbc, 0x4d, 0x8a, 0x2a, 0xd1, 0x56, 0x13, 0x4c,
	0x11, 0xd0, 0x44, 0x1b, 0x17, 0x60, 0x42, 0x88, 0x0e, 0x34, 0x2a, 0xc6, 0x34, 0x45, 0x83, 0x03,
	0x97, 0xc8, 0x89, 0xbd, 0xd4, 0x6a, 0x62, 0x57, 0x8e, 0x53, 0xb6, 0xff, 0x82, 0x3f, 0x6b, 0xc7,
	0x1d, 0x11, 0x87, 0x0a, 0x75, 0x57, 0xfe, 0x82, 0x9d, 0x90, 0xed, 0xb4, 0x1b, 0xd2, 0x0e, 0x9c,
	0x62, 0x7f, 0xdf, 0xe7, 0xcf, 0xef, 0x3d, 0xbf, 0x17, 0xb0, 0x93, 0x45, 0x28, 0x41, 0xc2, 0x8b,
	0x06, 0x88, 0xb2, 0xcc, 0x1b, 0xef, 0x85, 0x44, 0xa2, 0x3d, 0x6f, 0x84, 0x04, 0x4a, 0x33, 0x77,
	0x24, 0xb8, 0xe4, 0x70, 0xdb, 0x68, 0x5c, 0xa3, 0x71, 0x0b, 0x4d, 0x73, 0x2b, 0xe6, 0x31, 0xd7,
	0x0a, 0x4f, 0xad, 0x8c, 0xb8, 0xf9, 0xb8, 0x30, 0xcc, 0x25, 0x4d, 0xee, 0xfc, 0xe4, 0x40, 0x90,
	0x6c, 0xc0, 0x13, 0x6c, 0x54, 0x3b, 0x7f, 0x2a, 0xa0, 0x72, 0xaa, 0xef, 0x80, 0x5f, 0xc0, 0x92,
	0x36, 0xb6, 0xad, 0x8e, 0xe5, 0xd4, 0x7a, 0x6f, 0x6f, 0x27, 0xed, 0x83, 0x98, 0xca, 0x41, 0x1e,
	0xba, 0x11, 0x4f, 0x3d, 0x63, 0xc7, 0x45, 0x5c, 0xac, 0xba, 0x11, 0x17, 0xc4, 0xbb, 0xf0, 0x18,
	0xb9, 0xc8, 0x33, 0x8f, 0x5c, 0x8c, 0xb8, 0x90, 0x04, 0xbb, 0x87, 0xca, 0xe2, 0x04, 0xa5, 0xc4,
	0x37, 0x6e, 0xd0, 0x03, 0x9b, 0x11, 0x67, 0xe7, 0x54, 0xa4, 0x48, 0x52, 0xce, 0x82, 0x01, 0xa1,
	0xf1, 0x40, 0xda, 0x0b, 0x1d, 0xcb, 0x59, 0xf4, 0xe1, 0x7d, 0xea, 0xa3, 0x66, 0x20, 0x06, 0x2b,
	0x8c, 0xc8, 0xef, 0x5c, 0x0c, 0x83, 0x21, 0x65, 0xd8, 0x2e, 0x77, 0x2c, 0xa7, 0xd1, 0x7b, 0x77,
	0x3b, 0x69, 0xbf, 0x79, 0x30, 0x9c, 0x90, 0xca, 0x88, 0x53, 0xd6, 0x1d, 0xa3, 0x3c, 0x91, 0x5e,
	0xcc, 0xbb, 0x26, 0x63, 0x79, 0x39, 0x22, 0x99, 0x7b, 0x62, 0x9c, 0x3e, 0x51, 0x86, 0xfd, 0x3a,
	0xbb, 0xdb, 0xc0, 0x47, 0x00, 0x48, 0x3e, 0x24, 0x2c, 0x88, 0x38, 0x26, 0xf6, 0x62, 0xc7, 0x72,
	0x56, 0xfc, 0x9a, 0x46, 0x0e, 0x39, 0x26, 0xb0, 0x09, 0xaa, 0x61, 0x2e, 0x18, 0x0a, 0x13, 0x62,
	0x2f, 0x69, 0x72, 0xbe, 0x87, 0xfb, 0x60, 0x5b, 0x90, 0x31, 0x97, 0x24, 0x48, 0x78, 0x34, 0xa4,
	0x2c, 0x0e, 0x46, 0x44, 0x50, 0x8e, 0xed, 0x4a, 0xc7, 0x72, 0xca, 0xfe, 0xa6, 0x21, 0x8f, 0x0d,
	0x77, 0xaa, 0x29, 0x78, 0x06, 0xaa, 0xba, 0x1c, 0x01, 0xc5, 0xf6, 0xb2, 0xf2, 0xeb, 0xbd, 0xba,
	0x9a, 0xb4, 0x4b, 0xbf, 0x26, 0xed, 0xdd, 0x7b, 0x49, 0x45, 0x3c, 0x4b, 0x79, 0x56, 0x7c, 0xba,
	0x19, 0x1e, 0x16, 0x59, 0xf4, 0x99, 0x9c, 0x4e, 0xda, 0xcb, 0xba, 0xb6, 0xfd, 0xf7, 0xfe, 0xb2,
	0xb6, 0xea, 0x63, 0x78, 0x0a, 0xd6, 0xc7, 0x5c, 0xaa, 0x08, 0xe6, 0xef, 0x6a, 0x57, 0x3b, 0x96,
	0x53, 0xdf, 0x6f, 0xbb, 0x45, 0xaf, 0xe8, 0x62, 0xcc, 0x5a, 0xc5, 0x3d, 0x9b, 0xc9, 0x7a, 0x8b,
	0xea, 0x7a, 0x7f, 0xcd, 0x1c, 0x9f, 0xc3, 0x70, 0x17, 0xac, 0xa5, 0x94, 0x05, 0x2a, 0x01, 0x11,
	0x44, 0x3c, 0x67, 0xd2, 0xae, 0xe9, 0xac, 0x1a, 0x29, 0x65, 0x5f, 0x15, 0x7a, 0xa8, 0x40, 0xf8,
	0x1c, 0xc0, 0x88, 0xa7, 0x29, 0x62, 0x38, 0x0b, 0x62, 0x94, 0x05, 0x09, 0x4d, 0xa9, 0xb4, 0x81,
	0x7a, 0x2a, 0x7f, 0x7d, 0xc6, 0x1c, 0xa1, 0xec, 0x58, 0xe1, 0xd0, 0x05, 0x9b, 0x45, 0x9c, 0xb1,
	0x40, 0x11, 0x99, 0xd5, 0xab, 0xae, 0x9d, 0x37, 0x0c, 0x75, 0xa4, 0x98, 0xa2, 0x5a, 0x4f, 0xc1,
	0x06, 0x61, 0x38, 0x08, 0x55, 0x7d, 0x89, 0x28, 0xcc, 0x57, 0xb4, 0x7a, 0x8d, 0x30, 0xdc, 0x33,
	0xb8, 0xf1, 0x7e, 0x02, 0x56, 0xa5, 0x40, 0x2c, 0x3b, 0x9f, 0x0b, 0x1b, 0xba, 0xb5, 0x1a, 0x33,
	0xd4, 0xc8, 0x8e, 0x40, 0x35, 0x25, 0x12, 0x61, 0x24, 0x91, 0xbd, 0xda, 0x29, 0x3b, 0xf5, 0xfd,
	0x67, 0xee, 0x83, 0xe3, 0xe4, 0x9a, 0x71, 0x70, 0x3f, 0x17, 0xea, 0x0f, 0x4c, 0x8a, 0x4b, 0x7f,
	0x7e, 0xb8, 0x79, 0x00, 0x1a, 0xff, 0x50, 0x70, 0x1d, 0x94, 0x87, 0xe4, 0xd2, 0x4c, 0x8d, 0xaf,
	0x96, 0x70, 0x0b, 0x2c, 0x8d, 0x51, 0x92, 0x13, 0xdd, 0xe4, 0x35, 0xdf, 0x6c, 0x5e, 0x2f, 0xbc,
	0xb4, 0x7a, 0xfd, 0xab, 0x69, 0xcb, 0xba, 0x9e, 0xb6, 0xac, 0xdf, 0xd3, 0x96, 0xf5, 0xe3, 0xa6,
	0x55, 0xba, 0xbe, 0x69, 0x95, 0x7e, 0xde, 0xb4, 0x4a, 0xdf, 0xbc, 0xff, 0x18, 0xb5, 0xe2, 0xdf,
	0xa0, 0x7b, 0x22, 0xac, 0xe8, 0x01, 0x7e, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xed, 0x89, 0x0d,
	0xe8, 0x39, 0x04, 0x00, 0x00,
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
	if len(m.Metadata) > 0 {
		for k := range m.Metadata {
			v := m.Metadata[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintParams(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintParams(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintParams(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x72
		}
	}
	if m.TransferLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TransferLimit))
		i--
		dAtA[i] = 0x68
	}
	if m.EndBlockerLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EndBlockerLimit))
		i--
		dAtA[i] = 0x60
	}
	if m.VotingGracePeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.VotingGracePeriod))
		i--
		dAtA[i] = 0x58
	}
	if m.CommandsGasLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CommandsGasLimit))
		i--
		dAtA[i] = 0x50
	}
	if m.MinVoterCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinVoterCount))
		i--
		dAtA[i] = 0x48
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
	dAtA[i] = 0x42
	{
		size := m.ChainID.Size()
		i -= size
		if _, err := m.ChainID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.RevoteLockingPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RevoteLockingPeriod))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Burnable) > 0 {
		i -= len(m.Burnable)
		copy(dAtA[i:], m.Burnable)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Burnable)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TokenCode) > 0 {
		i -= len(m.TokenCode)
		copy(dAtA[i:], m.TokenCode)
		i = encodeVarintParams(dAtA, i, uint64(len(m.TokenCode)))
		i--
		dAtA[i] = 0x22
	}
	if m.NetworkKind != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.NetworkKind))
		i--
		dAtA[i] = 0x18
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
	if m.NetworkKind != 0 {
		n += 1 + sovParams(uint64(m.NetworkKind))
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
	l = m.ChainID.Size()
	n += 1 + l + sovParams(uint64(l))
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
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovParams(uint64(len(k))) + 1 + len(v) + sovParams(uint64(len(v)))
			n += mapEntrySize + 1 + sovParams(uint64(mapEntrySize))
		}
	}
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkKind", wireType)
			}
			m.NetworkKind = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NetworkKind |= github_com_scalarorg_bitcoin_vault_go_utils_types.NetworkKind(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
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
		case 5:
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
		case 6:
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
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
			if err := m.ChainID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
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
		case 9:
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
		case 10:
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
		case 11:
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
		case 12:
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
		case 13:
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
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
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
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthParams
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthParams
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthParams
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthParams
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipParams(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthParams
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
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
