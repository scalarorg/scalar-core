// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/multisig/v1beta1/params.proto

package types

import (
	fmt "fmt"
	utils "github.com/axelarnetwork/axelar-core/utils"
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
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4e, 0xea, 0x40,
	0x14, 0x87, 0xdb, 0x0b, 0x17, 0x93, 0x31, 0xf2, 0xa7, 0x92, 0xd8, 0xb0, 0x28, 0xc4, 0x60, 0x64,
	0xa1, 0x53, 0xd1, 0x37, 0xc0, 0x18, 0x97, 0x92, 0xc6, 0x95, 0x9b, 0x66, 0x28, 0x93, 0xe9, 0xc4,
	0xb6, 0xd3, 0x4c, 0xa7, 0x04, 0x5e, 0xc2, 0xf8, 0x58, 0x2c, 0x59, 0xba, 0x32, 0x0a, 0x2f, 0x62,
	0x3a, 0x33, 0x2d, 0xb8, 0x74, 0x37, 0x39, 0xbf, 0xef, 0x7c, 0xe7, 0x24, 0x73, 0xc0, 0x30, 0x0b,
	0x50, 0x84, 0xb8, 0x1b, 0xe7, 0x91, 0xa0, 0x19, 0x25, 0xee, 0x62, 0x3c, 0xc3, 0x02, 0x8d, 0xdd,
	0x14, 0x71, 0x14, 0x67, 0x30, 0xe5, 0x4c, 0x30, 0xeb, 0x4c, 0x51, 0xb0, 0xa4, 0xa0, 0xa6, 0x7a,
	0x5d, 0xc2, 0x08, 0x93, 0x8c, 0x5b, 0xbc, 0x14, 0xde, 0x1b, 0xa2, 0x25, 0x2e, 0xa4, 0xb9, 0xa0,
	0x51, 0x56, 0x19, 0x45, 0xc8, 0x71, 0x16, 0xb2, 0x68, 0xae, 0xa8, 0xf3, 0xb7, 0x1a, 0x68, 0x4c,
	0xe5, 0x14, 0x6b, 0x0a, 0xda, 0xaf, 0x78, 0x45, 0x70, 0xe2, 0x57, 0x90, 0x6d, 0x0e, 0xcc, 0xd1,
	0xf1, 0x6d, 0x1f, 0x2a, 0x17, 0x94, 0xae, 0x72, 0x2e, 0x7c, 0x2e, 0xb1, 0x49, 0x7d, 0xfd, 0xd9,
	0x37, 0xbc, 0x96, 0x6a, 0xaf, 0xca, 0x96, 0x07, 0x3a, 0x19, 0x25, 0x09, 0x4d, 0xc8, 0x81, 0xf2,
	0xdf, 0x5f, 0x94, 0x6d, 0xdd, 0xbf, 0x77, 0x5e, 0x80, 0x66, 0xb9, 0x25, 0x8d, 0x31, 0xcb, 0x85,
	0x5d, 0x1b, 0x98, 0xa3, 0x9a, 0x77, 0xa2, 0x87, 0xab, 0xa2, 0x05, 0xc1, 0xa9, 0xc6, 0x08, 0x47,
	0x01, 0xf6, 0x53, 0xcc, 0x29, 0x9b, 0xdb, 0x75, 0xc9, 0x76, 0x54, 0xf4, 0x58, 0x24, 0x53, 0x19,
	0x58, 0x97, 0xa0, 0x55, 0xad, 0xaa, 0xbd, 0xff, 0x25, 0xdb, 0x2c, 0x37, 0xd0, 0xe2, 0x1b, 0xd0,
	0x2d, 0xc1, 0x5f, 0xe6, 0x86, 0xa4, 0x2d, 0x9d, 0x1d, 0xaa, 0xaf, 0x80, 0x85, 0x02, 0x41, 0x17,
	0xd8, 0xc7, 0x29, 0x0b, 0x42, 0x3f, 0x60, 0x79, 0x22, 0xec, 0xa3, 0x81, 0x39, 0xaa, 0x7b, 0x6d,
	0x95, 0x3c, 0x14, 0xc1, 0x7d, 0x51, 0x9f, 0x3c, 0xad, 0xbf, 0x1d, 0x63, 0xbd, 0x75, 0xcc, 0xcd,
	0xd6, 0x31, 0xbf, 0xb6, 0x8e, 0xf9, 0xbe, 0x73, 0x8c, 0xcd, 0xce, 0x31, 0x3e, 0x76, 0x8e, 0xf1,
	0x32, 0x26, 0x54, 0x84, 0xf9, 0x0c, 0x06, 0x2c, 0x76, 0xd5, 0x39, 0x30, 0x4e, 0xf4, 0xeb, 0x3a,
	0x60, 0x1c, 0xbb, 0xcb, 0xfd, 0x15, 0x89, 0x55, 0x8a, 0xb3, 0x59, 0x43, 0x7e, 0xf4, 0xdd, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xba, 0x17, 0xd5, 0x65, 0x02, 0x00, 0x00,
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