// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/auxiliary/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("scalar/auxiliary/v1beta1/service.proto", fileDescriptor_cb00b5d875b7013f)
}
func init() {
	golang_proto.RegisterFile("scalar/auxiliary/v1beta1/service.proto", fileDescriptor_cb00b5d875b7013f)
}

var fileDescriptor_cb00b5d875b7013f = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x4e, 0x4e, 0xcc,
	0x49, 0x2c, 0xd2, 0x4f, 0x2c, 0xad, 0xc8, 0xcc, 0xc9, 0x4c, 0x2c, 0xaa, 0xd4, 0x2f, 0x33, 0x4c,
	0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x92, 0x80, 0xa8, 0xd3, 0x83, 0xab, 0xd3, 0x83, 0xaa, 0x93, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0x2b, 0xd2, 0x07, 0xb1, 0x20, 0xea, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73,
	0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3,
	0x8a, 0xa1, 0xb2, 0x8a, 0x38, 0x6d, 0x2d, 0xa9, 0x80, 0x28, 0x31, 0xea, 0x62, 0xe4, 0xe2, 0xf2,
	0x2d, 0x4e, 0x0f, 0x86, 0xb8, 0x42, 0xa8, 0x86, 0x8b, 0xd5, 0x29, 0xb1, 0x24, 0x39, 0x43, 0x48,
	0x4d, 0x0f, 0x97, 0x4b, 0xf4, 0xc0, 0x0a, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0xd4,
	0x09, 0xaa, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x55, 0x52, 0x6a, 0xba, 0xfc, 0x64, 0x32, 0x93,
	0x8c, 0x92, 0xb8, 0x3e, 0x86, 0xa3, 0x92, 0x40, 0x0a, 0xad, 0x18, 0xb5, 0x9c, 0x02, 0x4e, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x0e,
	0x3c, 0x96, 0x63, 0xbc, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xa3, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xa8, 0x19, 0xf9, 0x45, 0xe9, 0x50, 0x96, 0x6e,
	0x72, 0x7e, 0x51, 0xaa, 0x7e, 0x05, 0x92, 0xa1, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60,
	0x5f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x26, 0x7b, 0x22, 0x80, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	Batch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) Batch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponse, error) {
	out := new(BatchResponse)
	err := c.cc.Invoke(ctx, "/scalar.auxiliary.v1beta1.MsgService/Batch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	Batch(context.Context, *BatchRequest) (*BatchResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) Batch(ctx context.Context, req *BatchRequest) (*BatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Batch not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_Batch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Batch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.auxiliary.v1beta1.MsgService/Batch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Batch(ctx, req.(*BatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scalar.auxiliary.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Batch",
			Handler:    _MsgService_Batch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scalar/auxiliary/v1beta1/service.proto",
}
