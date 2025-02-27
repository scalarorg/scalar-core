// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/auxiliary/v1beta1/tx.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "github.com/scalarorg/scalar-core/x/permission/exported"
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

type BatchRequest struct {
	Sender   github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	Messages []types.Any                                   `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages"`
}

func (m *BatchRequest) Reset()         { *m = BatchRequest{} }
func (m *BatchRequest) String() string { return proto.CompactTextString(m) }
func (*BatchRequest) ProtoMessage()    {}
func (*BatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f81c516a80cf9221, []int{0}
}
func (m *BatchRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BatchRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchRequest.Merge(m, src)
}
func (m *BatchRequest) XXX_Size() int {
	return m.Size()
}
func (m *BatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchRequest proto.InternalMessageInfo

func (m *BatchRequest) GetSender() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *BatchRequest) GetMessages() []types.Any {
	if m != nil {
		return m.Messages
	}
	return nil
}

type BatchResponse struct {
	Responses []BatchResponse_Response `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses"`
}

func (m *BatchResponse) Reset()         { *m = BatchResponse{} }
func (m *BatchResponse) String() string { return proto.CompactTextString(m) }
func (*BatchResponse) ProtoMessage()    {}
func (*BatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f81c516a80cf9221, []int{1}
}
func (m *BatchResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BatchResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchResponse.Merge(m, src)
}
func (m *BatchResponse) XXX_Size() int {
	return m.Size()
}
func (m *BatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BatchResponse proto.InternalMessageInfo

func (m *BatchResponse) GetResponses() []BatchResponse_Response {
	if m != nil {
		return m.Responses
	}
	return nil
}

type BatchResponse_Response struct {
	// Types that are valid to be assigned to Res:
	//	*BatchResponse_Response_Result
	//	*BatchResponse_Response_Err
	Res isBatchResponse_Response_Res `protobuf_oneof:"res"`
}

func (m *BatchResponse_Response) Reset()         { *m = BatchResponse_Response{} }
func (m *BatchResponse_Response) String() string { return proto.CompactTextString(m) }
func (*BatchResponse_Response) ProtoMessage()    {}
func (*BatchResponse_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f81c516a80cf9221, []int{1, 0}
}
func (m *BatchResponse_Response) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BatchResponse_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BatchResponse_Response.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BatchResponse_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchResponse_Response.Merge(m, src)
}
func (m *BatchResponse_Response) XXX_Size() int {
	return m.Size()
}
func (m *BatchResponse_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchResponse_Response.DiscardUnknown(m)
}

var xxx_messageInfo_BatchResponse_Response proto.InternalMessageInfo

type isBatchResponse_Response_Res interface {
	isBatchResponse_Response_Res()
	MarshalTo([]byte) (int, error)
	Size() int
}

type BatchResponse_Response_Result struct {
	Result *types1.Result `protobuf:"bytes,1,opt,name=result,proto3,oneof" json:"result,omitempty"`
}
type BatchResponse_Response_Err struct {
	Err string `protobuf:"bytes,2,opt,name=err,proto3,oneof" json:"err,omitempty"`
}

func (*BatchResponse_Response_Result) isBatchResponse_Response_Res() {}
func (*BatchResponse_Response_Err) isBatchResponse_Response_Res()    {}

func (m *BatchResponse_Response) GetRes() isBatchResponse_Response_Res {
	if m != nil {
		return m.Res
	}
	return nil
}

func (m *BatchResponse_Response) GetResult() *types1.Result {
	if x, ok := m.GetRes().(*BatchResponse_Response_Result); ok {
		return x.Result
	}
	return nil
}

func (m *BatchResponse_Response) GetErr() string {
	if x, ok := m.GetRes().(*BatchResponse_Response_Err); ok {
		return x.Err
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BatchResponse_Response) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BatchResponse_Response_Result)(nil),
		(*BatchResponse_Response_Err)(nil),
	}
}

func init() {
	proto.RegisterType((*BatchRequest)(nil), "scalar.auxiliary.v1beta1.BatchRequest")
	proto.RegisterType((*BatchResponse)(nil), "scalar.auxiliary.v1beta1.BatchResponse")
	proto.RegisterType((*BatchResponse_Response)(nil), "scalar.auxiliary.v1beta1.BatchResponse.Response")
}

func init() { proto.RegisterFile("scalar/auxiliary/v1beta1/tx.proto", fileDescriptor_f81c516a80cf9221) }

var fileDescriptor_f81c516a80cf9221 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0x4d, 0x89, 0xda, 0x6b, 0x59, 0xac, 0x4a, 0xb8, 0x19, 0x9c, 0x50, 0x96, 0x2c,
	0xb9, 0x23, 0x61, 0xeb, 0x16, 0x4f, 0x45, 0x82, 0x01, 0x8b, 0x89, 0x05, 0x9d, 0xed, 0xc7, 0xd5,
	0x22, 0xf1, 0x99, 0x7b, 0x67, 0x94, 0x6c, 0x7c, 0x04, 0x3e, 0x0b, 0x82, 0xef, 0x10, 0x31, 0x75,
	0x64, 0xaa, 0x50, 0xf2, 0x2d, 0x98, 0x90, 0xef, 0x2e, 0x6e, 0x18, 0x98, 0xfc, 0x9e, 0xfd, 0x7b,
	0xef, 0xff, 0xbf, 0xff, 0x99, 0x3e, 0xc5, 0x5c, 0x2c, 0x84, 0xe6, 0xa2, 0x59, 0x95, 0x8b, 0x52,
	0xe8, 0x35, 0xff, 0x3c, 0xcd, 0xc0, 0x88, 0x29, 0x37, 0x2b, 0x56, 0x6b, 0x65, 0x54, 0x18, 0x39,
	0x84, 0x75, 0x08, 0xf3, 0xc8, 0xe0, 0x52, 0x2a, 0x25, 0x17, 0xc0, 0x2d, 0x97, 0x35, 0x1f, 0xb8,
	0xa8, 0xd6, 0x6e, 0x68, 0x70, 0x21, 0x95, 0x54, 0xb6, 0xe4, 0x6d, 0xe5, 0xdf, 0x32, 0xaf, 0x56,
	0x83, 0x5e, 0x96, 0x88, 0xa5, 0xaa, 0x38, 0xac, 0x6a, 0xa5, 0x0d, 0x14, 0x0f, 0xba, 0xeb, 0x1a,
	0xd0, 0xf3, 0xcf, 0x72, 0x85, 0x4b, 0x85, 0x3c, 0x13, 0x08, 0x5c, 0x64, 0x79, 0xd9, 0x51, 0x6d,
	0xe3, 0xa1, 0x4b, 0x07, 0xbd, 0x77, 0x6a, 0xae, 0x71, 0x9f, 0xae, 0xbe, 0x11, 0x7a, 0x9e, 0x08,
	0x93, 0xdf, 0xa6, 0xf0, 0xa9, 0x01, 0x34, 0xe1, 0x4b, 0xda, 0x47, 0xa8, 0x0a, 0xd0, 0x11, 0x19,
	0x91, 0xf1, 0x79, 0x32, 0xfd, 0x73, 0x3f, 0x9c, 0xc8, 0xd2, 0xdc, 0x36, 0x19, 0xcb, 0xd5, 0xd2,
	0x4f, 0xfb, 0xc7, 0x04, 0x8b, 0x8f, 0xde, 0xce, 0x3c, 0xcf, 0xe7, 0x45, 0xa1, 0x01, 0x31, 0xf5,
	0x0b, 0xc2, 0x37, 0xf4, 0x64, 0x09, 0x88, 0x42, 0x02, 0x46, 0x47, 0xa3, 0xde, 0xf8, 0x6c, 0x76,
	0xc1, 0x5c, 0x1e, 0x6c, 0x9f, 0x07, 0x9b, 0x57, 0xeb, 0x64, 0xb8, 0xb9, 0x1f, 0x06, 0x3f, 0xbf,
	0x4f, 0x9e, 0x78, 0x67, 0xed, 0x59, 0xf6, 0x09, 0xb2, 0xd7, 0x28, 0xd3, 0x6e, 0xcd, 0xf5, 0xf1,
	0x97, 0x1f, 0x11, 0xb9, 0xda, 0x10, 0xfa, 0xd8, 0x9b, 0xc6, 0x5a, 0x55, 0x08, 0xe1, 0x5b, 0x7a,
	0xaa, 0x7d, 0x8d, 0x11, 0xb1, 0x5a, 0xcf, 0xd9, 0xff, 0x6e, 0x85, 0xfd, 0x33, 0xcb, 0xf6, 0x45,
	0x72, 0xdc, 0xfa, 0x48, 0x1f, 0x16, 0x0d, 0x04, 0x3d, 0xe9, 0x14, 0xae, 0x69, 0x5f, 0x03, 0x36,
	0x0b, 0x63, 0x73, 0x39, 0x9b, 0x8d, 0xd8, 0xa1, 0x5b, 0x1b, 0xf6, 0x7e, 0x7d, 0x6a, 0xb9, 0x9b,
	0x20, 0xf5, 0x13, 0x61, 0x48, 0x7b, 0xa0, 0x75, 0x74, 0x34, 0x22, 0xe3, 0xd3, 0x9b, 0x20, 0x6d,
	0x9b, 0xe4, 0x11, 0xed, 0x69, 0xc0, 0xe4, 0xd5, 0x66, 0x1b, 0x93, 0xbb, 0x6d, 0x4c, 0x7e, 0x6f,
	0x63, 0xf2, 0x75, 0x17, 0x07, 0x77, 0xbb, 0x38, 0xf8, 0xb5, 0x8b, 0x83, 0x77, 0xb3, 0x83, 0xd0,
	0xdd, 0x49, 0x94, 0x96, 0xbe, 0x9a, 0xe4, 0x4a, 0x03, 0x5f, 0x1d, 0xfc, 0x93, 0xf6, 0x12, 0xb2,
	0xbe, 0xcd, 0xf5, 0xc5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0xa8, 0xc9, 0x65, 0xb4, 0x02,
	0x00, 0x00,
}

func (m *BatchRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BatchRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BatchResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BatchResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Responses) > 0 {
		for iNdEx := len(m.Responses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Responses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BatchResponse_Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchResponse_Response) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BatchResponse_Response) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Res != nil {
		{
			size := m.Res.Size()
			i -= size
			if _, err := m.Res.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *BatchResponse_Response_Result) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BatchResponse_Response_Result) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Result != nil {
		{
			size, err := m.Result.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *BatchResponse_Response_Err) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BatchResponse_Response_Err) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Err)
	copy(dAtA[i:], m.Err)
	i = encodeVarintTx(dAtA, i, uint64(len(m.Err)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BatchRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *BatchResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Responses) > 0 {
		for _, e := range m.Responses {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *BatchResponse_Response) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Res != nil {
		n += m.Res.Size()
	}
	return n
}

func (m *BatchResponse_Response_Result) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}
func (m *BatchResponse_Response_Err) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Err)
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BatchRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: BatchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, types.Any{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *BatchResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: BatchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Responses", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Responses = append(m.Responses, BatchResponse_Response{})
			if err := m.Responses[len(m.Responses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *BatchResponse_Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types1.Result{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Res = &BatchResponse_Response_Result{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Err", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Res = &BatchResponse_Response_Err{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
