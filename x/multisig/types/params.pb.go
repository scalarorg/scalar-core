// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/multisig/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	utils "github.com/scalarorg/scalar-core/utils"
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

// Params represent the genesis parameters for the module
type Params struct {
	KeygenThreshold    utils.Threshold `protobuf:"bytes,1,opt,name=keygen_threshold,json=keygenThreshold,proto3" json:"keygen_threshold"`
	SigningThreshold   utils.Threshold `protobuf:"bytes,2,opt,name=signing_threshold,json=signingThreshold,proto3" json:"signing_threshold"`
	KeygenTimeout      int64           `protobuf:"varint,3,opt,name=keygen_timeout,json=keygenTimeout,proto3" json:"keygen_timeout,omitempty"`
	KeygenGracePeriod  int64           `protobuf:"varint,4,opt,name=keygen_grace_period,json=keygenGracePeriod,proto3" json:"keygen_grace_period,omitempty"`
	SigningTimeout     int64           `protobuf:"varint,5,opt,name=signing_timeout,json=signingTimeout,proto3" json:"signing_timeout,omitempty"`
	SigningGracePeriod int64           `protobuf:"varint,6,opt,name=signing_grace_period,json=signingGracePeriod,proto3" json:"signing_grace_period,omitempty"`
	ActiveEpochCount   uint64          `protobuf:"varint,7,opt,name=active_epoch_count,json=activeEpochCount,proto3" json:"active_epoch_count,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_173ab6c85922430b, []int{0}
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

func init() {
	proto.RegisterType((*Params)(nil), "scalar.multisig.v1beta1.Params")
}

func init() {
	proto.RegisterFile("scalar/multisig/v1beta1/params.proto", fileDescriptor_173ab6c85922430b)
}

var fileDescriptor_173ab6c85922430b = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x4a, 0xeb, 0x40,
	0x18, 0x86, 0x93, 0xd3, 0x9e, 0x1e, 0x98, 0xc3, 0xe9, 0x4f, 0x4e, 0xc1, 0xd0, 0x45, 0x5a, 0x44,
	0xb1, 0x0b, 0x4d, 0xac, 0xde, 0x41, 0x45, 0x5c, 0x5a, 0x82, 0x2b, 0x37, 0x21, 0x9d, 0x0e, 0x93,
	0xc1, 0x24, 0x13, 0x66, 0x26, 0xc5, 0xde, 0x84, 0x78, 0x59, 0x5d, 0x76, 0xe9, 0x4a, 0xb4, 0xbd,
	0x11, 0xc9, 0xfc, 0xa4, 0x75, 0xe9, 0x6e, 0xf8, 0xde, 0x67, 0x9e, 0xf7, 0x23, 0x19, 0x70, 0xc2,
	0x61, 0x9c, 0xc6, 0x2c, 0xc8, 0xca, 0x54, 0x10, 0x4e, 0x70, 0xb0, 0x9c, 0xcc, 0x91, 0x88, 0x27,
	0x41, 0x11, 0xb3, 0x38, 0xe3, 0x7e, 0xc1, 0xa8, 0xa0, 0xce, 0x91, 0xa2, 0x7c, 0x43, 0xf9, 0x9a,
	0x1a, 0xf4, 0x31, 0xc5, 0x54, 0x32, 0x41, 0x75, 0x52, 0xf8, 0xc0, 0x48, 0x4b, 0x41, 0x52, 0x5e,
	0x1b, 0x45, 0xc2, 0x10, 0x4f, 0x68, 0xba, 0x50, 0xd4, 0xf1, 0x4b, 0x03, 0xb4, 0x66, 0xb2, 0xc5,
	0x99, 0x81, 0xee, 0x13, 0x5a, 0x61, 0x94, 0x47, 0x35, 0xe4, 0xda, 0x23, 0x7b, 0xfc, 0xf7, 0x6a,
	0xe8, 0xeb, 0x6a, 0xe9, 0x32, 0xbd, 0xfe, 0x83, 0xc1, 0xa6, 0xcd, 0xf5, 0xfb, 0xd0, 0x0a, 0x3b,
	0xea, 0x7a, 0x3d, 0x76, 0x42, 0xd0, 0xe3, 0x04, 0xe7, 0x24, 0xc7, 0x07, 0xca, 0x5f, 0x3f, 0x51,
	0x76, 0xf5, 0xfd, 0xbd, 0xf3, 0x14, 0xb4, 0xcd, 0x96, 0x24, 0x43, 0xb4, 0x14, 0x6e, 0x63, 0x64,
	0x8f, 0x1b, 0xe1, 0x3f, 0x5d, 0xae, 0x86, 0x8e, 0x0f, 0xfe, 0x6b, 0x0c, 0xb3, 0x18, 0xa2, 0xa8,
	0x40, 0x8c, 0xd0, 0x85, 0xdb, 0x94, 0x6c, 0x4f, 0x45, 0x77, 0x55, 0x32, 0x93, 0x81, 0x73, 0x06,
	0x3a, 0xf5, 0xaa, 0xda, 0xfb, 0x5b, 0xb2, 0x6d, 0xb3, 0x81, 0x16, 0x5f, 0x82, 0xbe, 0x01, 0xbf,
	0x99, 0x5b, 0x92, 0x76, 0x74, 0x76, 0xa8, 0x3e, 0x07, 0x4e, 0x0c, 0x05, 0x59, 0xa2, 0x08, 0x15,
	0x14, 0x26, 0x11, 0xa4, 0x65, 0x2e, 0xdc, 0x3f, 0x23, 0x7b, 0xdc, 0x0c, 0xbb, 0x2a, 0xb9, 0xad,
	0x82, 0x9b, 0x6a, 0x3e, 0xbd, 0x5f, 0x7f, 0x7a, 0xd6, 0x7a, 0xeb, 0xd9, 0x9b, 0xad, 0x67, 0x7f,
	0x6c, 0x3d, 0xfb, 0x75, 0xe7, 0x59, 0x9b, 0x9d, 0x67, 0xbd, 0xed, 0x3c, 0xeb, 0x71, 0x82, 0x89,
	0x48, 0xca, 0xb9, 0x0f, 0x69, 0x16, 0xa8, 0x0f, 0x48, 0x19, 0xd6, 0xa7, 0x0b, 0x48, 0x19, 0x0a,
	0x9e, 0xf7, 0xaf, 0x48, 0xac, 0x0a, 0xc4, 0xe7, 0x2d, 0xf9, 0xa3, 0xaf, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x7e, 0x49, 0x5a, 0x74, 0x65, 0x02, 0x00, 0x00,
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
	if m.ActiveEpochCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ActiveEpochCount))
		i--
		dAtA[i] = 0x38
	}
	if m.SigningGracePeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SigningGracePeriod))
		i--
		dAtA[i] = 0x30
	}
	if m.SigningTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SigningTimeout))
		i--
		dAtA[i] = 0x28
	}
	if m.KeygenGracePeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.KeygenGracePeriod))
		i--
		dAtA[i] = 0x20
	}
	if m.KeygenTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.KeygenTimeout))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.SigningThreshold.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.KeygenThreshold.MarshalToSizedBuffer(dAtA[:i])
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
	l = m.KeygenThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SigningThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.KeygenTimeout != 0 {
		n += 1 + sovParams(uint64(m.KeygenTimeout))
	}
	if m.KeygenGracePeriod != 0 {
		n += 1 + sovParams(uint64(m.KeygenGracePeriod))
	}
	if m.SigningTimeout != 0 {
		n += 1 + sovParams(uint64(m.SigningTimeout))
	}
	if m.SigningGracePeriod != 0 {
		n += 1 + sovParams(uint64(m.SigningGracePeriod))
	}
	if m.ActiveEpochCount != 0 {
		n += 1 + sovParams(uint64(m.ActiveEpochCount))
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
				return fmt.Errorf("proto: wrong wireType = %d for field KeygenThreshold", wireType)
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
			if err := m.KeygenThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningThreshold", wireType)
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
			if err := m.SigningThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeygenTimeout", wireType)
			}
			m.KeygenTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeygenTimeout |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeygenGracePeriod", wireType)
			}
			m.KeygenGracePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeygenGracePeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningTimeout", wireType)
			}
			m.SigningTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigningTimeout |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningGracePeriod", wireType)
			}
			m.SigningGracePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigningGracePeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveEpochCount", wireType)
			}
			m.ActiveEpochCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActiveEpochCount |= uint64(b&0x7F) << shift
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
