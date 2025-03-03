// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/covenant/exported/v1beta1/types.proto

package exported

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type PsbtState int32

const (
	NonExistent PsbtState = 0
	Pending     PsbtState = 1
	Creating    PsbtState = 2
	Signing     PsbtState = 3
	Completed   PsbtState = 4
)

var PsbtState_name = map[int32]string{
	0: "PSBT_STATE_UNSPECIFIED",
	1: "PSBT_STATE_PENDING",
	2: "PSBT_STATE_CREATING",
	3: "PSBT_STATE_SIGNING",
	4: "PSBT_STATE_COMPLETED",
}

var PsbtState_value = map[string]int32{
	"PSBT_STATE_UNSPECIFIED": 0,
	"PSBT_STATE_PENDING":     1,
	"PSBT_STATE_CREATING":    2,
	"PSBT_STATE_SIGNING":     3,
	"PSBT_STATE_COMPLETED":   4,
}

func (x PsbtState) String() string {
	return proto.EnumName(PsbtState_name, int32(x))
}

func (PsbtState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_690bea15c48caea4, []int{0}
}

type KeyState int32

const (
	Inactive KeyState = 0
	Assigned KeyState = 1
	Active   KeyState = 2
)

var KeyState_name = map[int32]string{
	0: "KEY_STATE_UNSPECIFIED",
	1: "KEY_STATE_ASSIGNED",
	2: "KEY_STATE_ACTIVE",
}

var KeyState_value = map[string]int32{
	"KEY_STATE_UNSPECIFIED": 0,
	"KEY_STATE_ASSIGNED":    1,
	"KEY_STATE_ACTIVE":      2,
}

func (x KeyState) String() string {
	return proto.EnumName(KeyState_name, int32(x))
}

func (KeyState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_690bea15c48caea4, []int{1}
}

type TapScriptSig struct {
	KeyXOnly  *KeyXOnly  `protobuf:"bytes,1,opt,name=key_x_only,json=keyXOnly,proto3,customtype=KeyXOnly" json:"key_x_only,omitempty"`
	LeafHash  *LeafHash  `protobuf:"bytes,2,opt,name=leaf_hash,json=leafHash,proto3,customtype=LeafHash" json:"leaf_hash,omitempty"`
	Signature *Signature `protobuf:"bytes,3,opt,name=signature,proto3,customtype=Signature" json:"signature,omitempty"`
}

func (m *TapScriptSig) Reset()         { *m = TapScriptSig{} }
func (m *TapScriptSig) String() string { return proto.CompactTextString(m) }
func (*TapScriptSig) ProtoMessage()    {}
func (*TapScriptSig) Descriptor() ([]byte, []int) {
	return fileDescriptor_690bea15c48caea4, []int{0}
}
func (m *TapScriptSig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TapScriptSig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TapScriptSig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TapScriptSig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapScriptSig.Merge(m, src)
}
func (m *TapScriptSig) XXX_Size() int {
	return m.Size()
}
func (m *TapScriptSig) XXX_DiscardUnknown() {
	xxx_messageInfo_TapScriptSig.DiscardUnknown(m)
}

var xxx_messageInfo_TapScriptSig proto.InternalMessageInfo

type TapScriptSigsList struct {
	List []TapScriptSig `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
}

func (m *TapScriptSigsList) Reset()         { *m = TapScriptSigsList{} }
func (m *TapScriptSigsList) String() string { return proto.CompactTextString(m) }
func (*TapScriptSigsList) ProtoMessage()    {}
func (*TapScriptSigsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_690bea15c48caea4, []int{1}
}
func (m *TapScriptSigsList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TapScriptSigsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TapScriptSigsList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TapScriptSigsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapScriptSigsList.Merge(m, src)
}
func (m *TapScriptSigsList) XXX_Size() int {
	return m.Size()
}
func (m *TapScriptSigsList) XXX_DiscardUnknown() {
	xxx_messageInfo_TapScriptSigsList.DiscardUnknown(m)
}

var xxx_messageInfo_TapScriptSigsList proto.InternalMessageInfo

type TapScriptSigsEntry struct {
	Index uint64            `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Sigs  TapScriptSigsList `protobuf:"bytes,2,opt,name=sigs,proto3" json:"sigs"`
}

func (m *TapScriptSigsEntry) Reset()         { *m = TapScriptSigsEntry{} }
func (m *TapScriptSigsEntry) String() string { return proto.CompactTextString(m) }
func (*TapScriptSigsEntry) ProtoMessage()    {}
func (*TapScriptSigsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_690bea15c48caea4, []int{2}
}
func (m *TapScriptSigsEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TapScriptSigsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TapScriptSigsEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TapScriptSigsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapScriptSigsEntry.Merge(m, src)
}
func (m *TapScriptSigsEntry) XXX_Size() int {
	return m.Size()
}
func (m *TapScriptSigsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TapScriptSigsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TapScriptSigsEntry proto.InternalMessageInfo

// The reason we use a list instead of a map is because the map is not ensured
// the deterministic order of the entries
type TapScriptSigsMap struct {
	Inner []TapScriptSigsEntry `protobuf:"bytes,1,rep,name=inner,proto3" json:"inner"`
}

func (m *TapScriptSigsMap) Reset()         { *m = TapScriptSigsMap{} }
func (m *TapScriptSigsMap) String() string { return proto.CompactTextString(m) }
func (*TapScriptSigsMap) ProtoMessage()    {}
func (*TapScriptSigsMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_690bea15c48caea4, []int{3}
}
func (m *TapScriptSigsMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TapScriptSigsMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TapScriptSigsMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TapScriptSigsMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapScriptSigsMap.Merge(m, src)
}
func (m *TapScriptSigsMap) XXX_Size() int {
	return m.Size()
}
func (m *TapScriptSigsMap) XXX_DiscardUnknown() {
	xxx_messageInfo_TapScriptSigsMap.DiscardUnknown(m)
}

var xxx_messageInfo_TapScriptSigsMap proto.InternalMessageInfo

type ListOfTapScriptSigsMap struct {
	Inner []*TapScriptSigsMap `protobuf:"bytes,1,rep,name=inner,proto3" json:"inner,omitempty"`
}

func (m *ListOfTapScriptSigsMap) Reset()         { *m = ListOfTapScriptSigsMap{} }
func (m *ListOfTapScriptSigsMap) String() string { return proto.CompactTextString(m) }
func (*ListOfTapScriptSigsMap) ProtoMessage()    {}
func (*ListOfTapScriptSigsMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_690bea15c48caea4, []int{4}
}
func (m *ListOfTapScriptSigsMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListOfTapScriptSigsMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListOfTapScriptSigsMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListOfTapScriptSigsMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOfTapScriptSigsMap.Merge(m, src)
}
func (m *ListOfTapScriptSigsMap) XXX_Size() int {
	return m.Size()
}
func (m *ListOfTapScriptSigsMap) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOfTapScriptSigsMap.DiscardUnknown(m)
}

var xxx_messageInfo_ListOfTapScriptSigsMap proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("scalar.covenant.exported.v1beta1.PsbtState", PsbtState_name, PsbtState_value)
	proto.RegisterEnum("scalar.covenant.exported.v1beta1.KeyState", KeyState_name, KeyState_value)
	proto.RegisterType((*TapScriptSig)(nil), "scalar.covenant.exported.v1beta1.TapScriptSig")
	proto.RegisterType((*TapScriptSigsList)(nil), "scalar.covenant.exported.v1beta1.TapScriptSigsList")
	proto.RegisterType((*TapScriptSigsEntry)(nil), "scalar.covenant.exported.v1beta1.TapScriptSigsEntry")
	proto.RegisterType((*TapScriptSigsMap)(nil), "scalar.covenant.exported.v1beta1.TapScriptSigsMap")
	proto.RegisterType((*ListOfTapScriptSigsMap)(nil), "scalar.covenant.exported.v1beta1.ListOfTapScriptSigsMap")
}

func init() {
	proto.RegisterFile("scalar/covenant/exported/v1beta1/types.proto", fileDescriptor_690bea15c48caea4)
}

var fileDescriptor_690bea15c48caea4 = []byte{
	// 628 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5d, 0x6b, 0xd3, 0x5e,
	0x1c, 0xc7, 0x93, 0xad, 0xff, 0xfd, 0xdb, 0xd3, 0x0e, 0xe3, 0x71, 0x8e, 0x92, 0x8b, 0x2c, 0x54,
	0x65, 0x73, 0xd3, 0x84, 0x3d, 0xbc, 0x81, 0x3e, 0xc4, 0xad, 0x6c, 0xed, 0x42, 0x12, 0x45, 0x05,
	0x29, 0xa7, 0xed, 0x59, 0x1a, 0x96, 0x9d, 0x84, 0x9c, 0xb3, 0xd1, 0xbc, 0x00, 0x41, 0x7a, 0x25,
	0xde, 0xf7, 0x4a, 0x2f, 0x7c, 0x29, 0xbb, 0xdc, 0xa5, 0x0c, 0x19, 0xba, 0xbd, 0x11, 0xc9, 0x43,
	0x69, 0xca, 0x04, 0xe9, 0xdd, 0x2f, 0xc9, 0x27, 0xdf, 0x07, 0xf2, 0xcb, 0x01, 0x2f, 0x68, 0x0f,
	0xb9, 0x28, 0x50, 0x7b, 0xde, 0x05, 0x26, 0x88, 0x30, 0x15, 0x0f, 0x7d, 0x2f, 0x60, 0xb8, 0xaf,
	0x5e, 0x6c, 0x77, 0x31, 0x43, 0xdb, 0x2a, 0x0b, 0x7d, 0x4c, 0x15, 0x3f, 0xf0, 0x98, 0x07, 0xe5,
	0x84, 0x56, 0x26, 0xb4, 0x32, 0xa1, 0x95, 0x94, 0x16, 0x57, 0x6c, 0xcf, 0xf6, 0x62, 0x58, 0x8d,
	0xa6, 0xe4, 0xbd, 0xca, 0x17, 0x1e, 0x94, 0x2c, 0xe4, 0x9b, 0xbd, 0xc0, 0xf1, 0x99, 0xe9, 0xd8,
	0x70, 0x13, 0x80, 0x53, 0x1c, 0x76, 0x86, 0x1d, 0x8f, 0xb8, 0x61, 0x99, 0x97, 0xf9, 0x8d, 0x52,
	0xad, 0x74, 0x7d, 0xb3, 0x96, 0x3f, 0xc4, 0xe1, 0xdb, 0x63, 0xe2, 0x86, 0x46, 0xfe, 0x34, 0x9d,
	0xe0, 0x73, 0x50, 0x70, 0x31, 0x3a, 0xe9, 0x0c, 0x10, 0x1d, 0x94, 0x17, 0xa6, 0xe8, 0x11, 0x46,
	0x27, 0x07, 0x88, 0x0e, 0x8c, 0xbc, 0x9b, 0x4e, 0x70, 0x0b, 0x14, 0xa8, 0x63, 0x13, 0xc4, 0xce,
	0x03, 0x5c, 0x5e, 0x8c, 0xd1, 0xe5, 0xeb, 0x9b, 0xb5, 0x82, 0x39, 0xb9, 0x69, 0x4c, 0x9f, 0x57,
	0x3e, 0x80, 0x87, 0xd9, 0x4c, 0xf4, 0xc8, 0xa1, 0x0c, 0x1e, 0x80, 0x9c, 0xeb, 0x50, 0x56, 0xe6,
	0xe5, 0xc5, 0x8d, 0xe2, 0x8e, 0xa2, 0xfc, 0xab, 0xb0, 0x92, 0x95, 0xa8, 0xe5, 0x2e, 0x6f, 0xd6,
	0x38, 0x23, 0x56, 0xa8, 0x84, 0x00, 0xce, 0xc8, 0x6b, 0x84, 0x05, 0x21, 0x5c, 0x01, 0xff, 0x39,
	0xa4, 0x8f, 0x87, 0x71, 0xe7, 0x9c, 0x91, 0x5c, 0xc0, 0x16, 0xc8, 0x51, 0xc7, 0xa6, 0x71, 0xbb,
	0xe2, 0xce, 0xee, 0x7c, 0xae, 0x71, 0xf0, 0x89, 0x75, 0x24, 0x53, 0xe9, 0x03, 0x61, 0x06, 0x68,
	0x21, 0x1f, 0xea, 0x91, 0x31, 0xc1, 0x41, 0xda, 0x6c, 0x6f, 0x4e, 0x8f, 0x38, 0x7d, 0x6a, 0x92,
	0x08, 0x55, 0xba, 0x60, 0x35, 0x72, 0x3e, 0x3e, 0xb9, 0xe7, 0x75, 0x30, 0xeb, 0xb5, 0x33, 0xa7,
	0x57, 0x0b, 0xf9, 0xa9, 0xc7, 0xe6, 0x4f, 0x1e, 0x14, 0x74, 0xda, 0x65, 0x26, 0x43, 0x0c, 0xc3,
	0x2d, 0xb0, 0xaa, 0x9b, 0x35, 0xab, 0x63, 0x5a, 0x55, 0x4b, 0xeb, 0xbc, 0x6e, 0x9b, 0xba, 0x56,
	0x6f, 0xbe, 0x6a, 0x6a, 0x0d, 0x81, 0x13, 0x1f, 0x8c, 0xc6, 0x72, 0xb1, 0xed, 0x11, 0x6d, 0xe8,
	0x50, 0x86, 0x09, 0x83, 0x4f, 0x00, 0xcc, 0xc0, 0xba, 0xd6, 0x6e, 0x34, 0xdb, 0xfb, 0x02, 0x2f,
	0x16, 0x47, 0x63, 0xf9, 0x7f, 0x1d, 0x93, 0xbe, 0x43, 0x6c, 0xf8, 0x0c, 0x3c, 0xca, 0x40, 0x75,
	0x43, 0xab, 0x5a, 0x11, 0xb5, 0x20, 0x96, 0x46, 0x63, 0x39, 0x5f, 0x0f, 0x30, 0x62, 0x11, 0x36,
	0xab, 0x65, 0x36, 0xf7, 0xdb, 0x11, 0xb5, 0x98, 0x68, 0x45, 0xcb, 0x15, 0x41, 0xeb, 0x60, 0x25,
	0xab, 0x75, 0xdc, 0xd2, 0x8f, 0x34, 0x4b, 0x6b, 0x08, 0x39, 0x71, 0x79, 0x34, 0x96, 0x0b, 0x75,
	0xef, 0xcc, 0x77, 0x31, 0xc3, 0x7d, 0x31, 0xff, 0xe9, 0xab, 0xc4, 0x7d, 0xff, 0x26, 0xf1, 0x9b,
	0x1f, 0x79, 0x10, 0x6d, 0x7c, 0xd2, 0x6e, 0x1d, 0x3c, 0x3e, 0xd4, 0xde, 0xfd, 0xb5, 0x5c, 0x9c,
	0xa6, 0x49, 0x50, 0x8f, 0x39, 0x17, 0x18, 0x3e, 0x05, 0x70, 0x0a, 0x56, 0xcd, 0x28, 0x8e, 0xd6,
	0x10, 0xf8, 0x84, 0xaa, 0xd2, 0x68, 0xc3, 0x71, 0x1f, 0xca, 0x40, 0xc8, 0x50, 0x75, 0xab, 0xf9,
	0x46, 0x13, 0x16, 0x44, 0x30, 0x1a, 0xcb, 0x4b, 0xd5, 0x58, 0x67, 0x9a, 0xa3, 0x66, 0x5c, 0xfe,
	0x96, 0xb8, 0xcb, 0x5b, 0x89, 0xbf, 0xba, 0x95, 0xf8, 0x5f, 0xb7, 0x12, 0xff, 0xf9, 0x4e, 0xe2,
	0xae, 0xee, 0x24, 0xee, 0xc7, 0x9d, 0xc4, 0xbd, 0xdf, 0xb3, 0x1d, 0x36, 0x38, 0xef, 0x2a, 0x3d,
	0xef, 0x4c, 0x4d, 0xbe, 0xa4, 0x17, 0xd8, 0xe9, 0xf4, 0xb2, 0xe7, 0x05, 0x58, 0x1d, 0xde, 0x3f,
	0x3f, 0xba, 0x4b, 0xf1, 0xaf, 0xbf, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xcb, 0xa4, 0xd0,
	0x62, 0x04, 0x00, 0x00,
}

func (m *TapScriptSig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TapScriptSig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TapScriptSig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Signature != nil {
		{
			size := m.Signature.Size()
			i -= size
			if _, err := m.Signature.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.LeafHash != nil {
		{
			size := m.LeafHash.Size()
			i -= size
			if _, err := m.LeafHash.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.KeyXOnly != nil {
		{
			size := m.KeyXOnly.Size()
			i -= size
			if _, err := m.KeyXOnly.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TapScriptSigsList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TapScriptSigsList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TapScriptSigsList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.List[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TapScriptSigsEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TapScriptSigsEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TapScriptSigsEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Sigs.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Index != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TapScriptSigsMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TapScriptSigsMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TapScriptSigsMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Inner) > 0 {
		for iNdEx := len(m.Inner) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Inner[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ListOfTapScriptSigsMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListOfTapScriptSigsMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListOfTapScriptSigsMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Inner) > 0 {
		for iNdEx := len(m.Inner) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Inner[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *TapScriptSig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.KeyXOnly != nil {
		l = m.KeyXOnly.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.LeafHash != nil {
		l = m.LeafHash.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Signature != nil {
		l = m.Signature.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *TapScriptSigsList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *TapScriptSigsEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovTypes(uint64(m.Index))
	}
	l = m.Sigs.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *TapScriptSigsMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Inner) > 0 {
		for _, e := range m.Inner {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *ListOfTapScriptSigsMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Inner) > 0 {
		for _, e := range m.Inner {
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
func (m *TapScriptSig) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TapScriptSig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TapScriptSig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyXOnly", wireType)
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
			var v KeyXOnly
			m.KeyXOnly = &v
			if err := m.KeyXOnly.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeafHash", wireType)
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
			var v LeafHash
			m.LeafHash = &v
			if err := m.LeafHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
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
			var v Signature
			m.Signature = &v
			if err := m.Signature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *TapScriptSigsList) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TapScriptSigsList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TapScriptSigsList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
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
			m.List = append(m.List, TapScriptSig{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *TapScriptSigsEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TapScriptSigsEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TapScriptSigsEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sigs", wireType)
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
			if err := m.Sigs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *TapScriptSigsMap) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TapScriptSigsMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TapScriptSigsMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inner", wireType)
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
			m.Inner = append(m.Inner, TapScriptSigsEntry{})
			if err := m.Inner[len(m.Inner)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ListOfTapScriptSigsMap) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListOfTapScriptSigsMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListOfTapScriptSigsMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inner", wireType)
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
			m.Inner = append(m.Inner, &TapScriptSigsMap{})
			if err := m.Inner[len(m.Inner)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
