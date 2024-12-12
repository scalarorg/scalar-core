// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/multisig/exported/v1beta1/types.proto

package exported

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type MultisigState int32

const (
	NonExistent MultisigState = 0
	Pending     MultisigState = 1
	Completed   MultisigState = 2
)

var MultisigState_name = map[int32]string{
	0: "MULTISIG_STATE_UNSPECIFIED",
	1: "MULTISIG_STATE_PENDING",
	2: "MULTISIG_STATE_COMPLETED",
}

var MultisigState_value = map[string]int32{
	"MULTISIG_STATE_UNSPECIFIED": 0,
	"MULTISIG_STATE_PENDING":     1,
	"MULTISIG_STATE_COMPLETED":   2,
}

func (x MultisigState) String() string {
	return proto.EnumName(MultisigState_name, int32(x))
}

func (MultisigState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a8510f4fe43ebe84, []int{0}
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
	return fileDescriptor_a8510f4fe43ebe84, []int{1}
}

func init() {
	proto.RegisterEnum("scalar.multisig.exported.v1beta1.MultisigState", MultisigState_name, MultisigState_value)
	proto.RegisterEnum("scalar.multisig.exported.v1beta1.KeyState", KeyState_name, KeyState_value)
}

func init() {
	proto.RegisterFile("scalar/multisig/exported/v1beta1/types.proto", fileDescriptor_a8510f4fe43ebe84)
}

var fileDescriptor_a8510f4fe43ebe84 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd1, 0xc1, 0x8a, 0xda, 0x40,
	0x1c, 0x06, 0xf0, 0x8c, 0x07, 0x6b, 0xc7, 0x4a, 0x43, 0x68, 0x4b, 0xc9, 0x61, 0xc8, 0xa1, 0x20,
	0xd8, 0x36, 0x83, 0xb4, 0x2f, 0x90, 0x26, 0x53, 0x19, 0xd4, 0x34, 0x34, 0xb1, 0xd0, 0x5e, 0x24,
	0x26, 0x43, 0x1a, 0xd0, 0x4c, 0x48, 0x46, 0xd1, 0x07, 0x28, 0x94, 0x9c, 0xfa, 0x02, 0x81, 0xc2,
	0xee, 0x61, 0x1f, 0xc5, 0xa3, 0xc7, 0x3d, 0xee, 0xea, 0x8b, 0x2c, 0x6b, 0x14, 0x17, 0x77, 0x6f,
	0xff, 0xc3, 0xef, 0xfb, 0xf3, 0xc1, 0x07, 0x3f, 0xe4, 0x81, 0x3f, 0xf5, 0x33, 0x3c, 0x9b, 0x4f,
	0x45, 0x9c, 0xc7, 0x11, 0x66, 0xcb, 0x94, 0x67, 0x82, 0x85, 0x78, 0xd1, 0x9d, 0x30, 0xe1, 0x77,
	0xb1, 0x58, 0xa5, 0x2c, 0xd7, 0xd3, 0x8c, 0x0b, 0xae, 0x68, 0x95, 0xd6, 0x8f, 0x5a, 0x3f, 0x6a,
	0xfd, 0xa0, 0xd5, 0x57, 0x11, 0x8f, 0xf8, 0x1e, 0xe3, 0xfb, 0xab, 0xca, 0x75, 0xfe, 0x03, 0xd8,
	0x1a, 0x1e, 0x32, 0xae, 0xf0, 0x05, 0x53, 0x30, 0x54, 0x87, 0xa3, 0x81, 0x47, 0x5d, 0xda, 0x1b,
	0xbb, 0x9e, 0xe1, 0x91, 0xf1, 0xc8, 0x76, 0x1d, 0x62, 0xd2, 0xaf, 0x94, 0x58, 0xb2, 0xa4, 0xbe,
	0x2c, 0x4a, 0xad, 0x69, 0xf3, 0x84, 0x2c, 0xe3, 0x5c, 0xb0, 0x44, 0x28, 0x6d, 0xf8, 0xe6, 0x2c,
	0xe0, 0x10, 0xdb, 0xa2, 0x76, 0x4f, 0x06, 0x6a, 0xb3, 0x28, 0xb5, 0x67, 0x0e, 0x4b, 0xc2, 0x38,
	0x89, 0x94, 0xf7, 0xf0, 0xed, 0x19, 0x34, 0xbf, 0x0d, 0x9d, 0x01, 0xf1, 0x88, 0x25, 0xd7, 0xd4,
	0x56, 0x51, 0x6a, 0xcf, 0x4d, 0x3e, 0x4b, 0xa7, 0x4c, 0xb0, 0x50, 0x6d, 0xfc, 0xbd, 0x40, 0xd2,
	0xd5, 0x25, 0x02, 0x9d, 0x3f, 0x00, 0x36, 0xfa, 0x6c, 0x55, 0xb5, 0x6b, 0xc3, 0xd7, 0x7d, 0xf2,
	0xf3, 0xc9, 0x62, 0x2f, 0x8a, 0x52, 0x6b, 0xd0, 0xc4, 0x0f, 0x44, 0xbc, 0x60, 0xca, 0x3b, 0xa8,
	0x9c, 0xa0, 0xe1, 0xba, 0xb4, 0x67, 0x13, 0x4b, 0x06, 0x95, 0x32, 0xf2, 0x3c, 0x8e, 0x12, 0x16,
	0x2a, 0x1a, 0x94, 0x1f, 0x28, 0xd3, 0xa3, 0x3f, 0x88, 0x5c, 0x53, 0x61, 0x51, 0x6a, 0x75, 0x63,
	0xff, 0xe7, 0xd4, 0xe3, 0xcb, 0xf7, 0xf5, 0x2d, 0x92, 0xd6, 0x5b, 0x04, 0x36, 0x5b, 0x04, 0x6e,
	0xb6, 0x08, 0xfc, 0xdb, 0x21, 0x69, 0xb3, 0x43, 0xd2, 0xf5, 0x0e, 0x49, 0xbf, 0x3e, 0x47, 0xb1,
	0xf8, 0x3d, 0x9f, 0xe8, 0x01, 0x9f, 0xe1, 0x6a, 0x0b, 0x9e, 0x45, 0x87, 0xeb, 0x63, 0xc0, 0x33,
	0x86, 0x97, 0x8f, 0xa7, 0x9c, 0xd4, 0xf7, 0x2b, 0x7c, 0xba, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xbb,
	0x54, 0x11, 0xa2, 0xed, 0x01, 0x00, 0x00,
}
