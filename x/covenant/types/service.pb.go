// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scalar/covenant/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

func init() {
	proto.RegisterFile("scalar/covenant/v1beta1/service.proto", fileDescriptor_b96b94daf72c97ef)
}

var fileDescriptor_b96b94daf72c97ef = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x41, 0x6b, 0x13, 0x41,
	0x18, 0x86, 0x3b, 0x0a, 0x01, 0x47, 0x41, 0x98, 0x0a, 0x4a, 0x90, 0x55, 0x53, 0x6d, 0x6d, 0xaa,
	0xbb, 0xa6, 0x8d, 0xa0, 0xb9, 0x88, 0x16, 0xf4, 0x50, 0x04, 0x4d, 0xea, 0xc5, 0xcb, 0x32, 0xd9,
	0x1d, 0xd7, 0xc5, 0xee, 0xce, 0x76, 0x66, 0x36, 0x34, 0xd7, 0x9e, 0x3d, 0x88, 0xfd, 0x17, 0xfe,
	0x02, 0x05, 0xc1, 0x93, 0xd0, 0x63, 0xc1, 0x8b, 0x47, 0x49, 0xfc, 0x21, 0x92, 0xd9, 0x9d, 0x6c,
	0xbb, 0xcd, 0x64, 0xa7, 0x78, 0xcb, 0xe1, 0x79, 0x27, 0xcf, 0xfb, 0xcd, 0xf0, 0x2d, 0xbc, 0xc3,
	0x3d, 0xbc, 0x83, 0x99, 0xe3, 0xd1, 0x01, 0x89, 0x71, 0x2c, 0x9c, 0x41, 0xab, 0x4f, 0x04, 0x6e,
	0x39, 0x9c, 0xb0, 0x41, 0xe8, 0x11, 0x3b, 0x61, 0x54, 0x50, 0x74, 0x35, 0xc3, 0x6c, 0x85, 0xd9,
	0x39, 0x56, 0xbf, 0x1e, 0x50, 0x1a, 0xec, 0x10, 0x07, 0x27, 0xa1, 0x83, 0xe3, 0x98, 0x0a, 0x2c,
	0x42, 0x1a, 0xf3, 0x2c, 0x56, 0x5f, 0xd2, 0x9d, 0xbe, 0x9b, 0x12, 0x36, 0xcc, 0xa1, 0x9b, 0x3a,
	0x48, 0xec, 0x65, 0xc4, 0xfa, 0xcf, 0x8b, 0x10, 0xbe, 0xe4, 0x41, 0x2f, 0x53, 0x42, 0x5f, 0x00,
	0xbc, 0xbc, 0xc9, 0x08, 0x16, 0x64, 0x33, 0xe5, 0x82, 0xfa, 0x21, 0x8e, 0x91, 0x63, 0x6b, 0x0c,
	0xed, 0x12, 0xd9, 0x25, 0xbb, 0x29, 0xe1, 0xa2, 0xfe, 0xc0, 0x3c, 0xc0, 0x13, 0x1a, 0x73, 0xd2,
	0x68, 0xef, 0xff, 0xfa, 0x7b, 0x70, 0xce, 0x6e, 0xac, 0x3a, 0x3a, 0x61, 0x4f, 0x26, 0x5d, 0x4f,
	0x45, 0x3b, 0xa0, 0x29, 0x65, 0xdf, 0x24, 0xbe, 0xa1, 0x6c, 0x89, 0xac, 0x96, 0x3d, 0x15, 0x30,
	0x96, 0x4d, 0x65, 0xf2, 0xa4, 0xec, 0x0f, 0x00, 0xaf, 0x94, 0xea, 0xbf, 0x60, 0x34, 0x4d, 0x50,
	0xdb, 0x74, 0x5a, 0x12, 0x57, 0xda, 0x0f, 0xcf, 0x98, 0xca, 0xdd, 0x3b, 0xd2, 0xbd, 0xdd, 0x70,
	0x8c, 0x07, 0xed, 0x06, 0x93, 0x03, 0x54, 0x83, 0xd2, 0x4c, 0xaa, 0x1a, 0xcc, 0xc2, 0xab, 0x1b,
	0xcc, 0x4e, 0x19, 0x37, 0x28, 0x4f, 0xbf, 0x68, 0xf0, 0x0d, 0xc0, 0xc5, 0xa7, 0xbe, 0x3f, 0x3d,
	0x79, 0x9b, 0x66, 0x05, 0x36, 0xb4, 0x2a, 0x33, 0x68, 0xe5, 0xdf, 0xd2, 0xdf, 0xc0, 0xa9, 0x84,
	0xb1, 0x3b, 0xf6, 0xfd, 0x63, 0xe2, 0x82, 0x16, 0xee, 0x87, 0x00, 0x5e, 0xeb, 0x92, 0x88, 0x0e,
	0x8a, 0xc1, 0x3c, 0x67, 0x34, 0xca, 0x0a, 0x3c, 0xd2, 0xba, 0xe8, 0x22, 0xff, 0xd1, 0xe2, 0x89,
	0x6c, 0xf1, 0xb8, 0xd1, 0xd6, 0xb6, 0x60, 0xf2, 0x4f, 0x8f, 0x15, 0x79, 0xc7, 0x68, 0x54, 0x54,
	0xf9, 0x0c, 0xe0, 0x85, 0xee, 0x64, 0x9d, 0x91, 0x2d, 0x32, 0x44, 0xab, 0x7a, 0x77, 0xc5, 0x28,
	0xd9, 0xa6, 0x09, 0x9a, 0x5b, 0xda, 0xd2, 0xf2, 0x6e, 0x63, 0x49, 0x6f, 0x29, 0x33, 0xee, 0x07,
	0x32, 0x9c, 0x48, 0x7d, 0x07, 0x70, 0xb1, 0x97, 0xf6, 0xa3, 0x50, 0x6c, 0xe3, 0xa4, 0xe7, 0xb1,
	0x30, 0x11, 0xbd, 0x30, 0xe0, 0x73, 0xde, 0xc6, 0x0c, 0x5a, 0x89, 0xb6, 0xcf, 0x16, 0x32, 0x7e,
	0x1e, 0x5c, 0xa6, 0x5d, 0x81, 0x13, 0x97, 0xcb, 0xbc, 0xcb, 0xc3, 0x80, 0x77, 0x40, 0x73, 0xfd,
	0xeb, 0x79, 0x78, 0xe9, 0xf5, 0x64, 0xf3, 0xab, 0x4d, 0x7e, 0x00, 0x20, 0x9c, 0x5e, 0x21, 0x47,
	0xcd, 0xea, 0x7b, 0x9e, 0xda, 0xaf, 0x19, 0xb1, 0xb9, 0xf4, 0x3d, 0x29, 0xbd, 0x8c, 0x6e, 0x17,
	0xd2, 0x71, 0x79, 0xa5, 0x14, 0x1a, 0x1f, 0x01, 0xac, 0xc9, 0xd7, 0xc4, 0xd1, 0xb2, 0xf6, 0x5f,
	0x32, 0x40, 0xd9, 0xac, 0x54, 0x72, 0xb9, 0x49, 0x4b, 0x9a, 0xac, 0xa1, 0x39, 0x1f, 0x91, 0x93,
	0x2b, 0x81, 0xa3, 0x7d, 0x00, 0x6b, 0xaf, 0x30, 0xc3, 0xd1, 0x3c, 0x9d, 0x0c, 0xa8, 0xd6, 0x51,
	0x5c, 0xae, 0xb3, 0x22, 0x75, 0x6e, 0xa1, 0x1b, 0x5a, 0x9d, 0x44, 0x06, 0x9e, 0x6d, 0x1d, 0x8e,
	0x2c, 0x70, 0x34, 0xb2, 0xc0, 0x9f, 0x91, 0x05, 0x3e, 0x8d, 0xad, 0x85, 0xa3, 0xb1, 0xb5, 0xf0,
	0x7b, 0x6c, 0x2d, 0xbc, 0x6d, 0x05, 0xa1, 0x78, 0x9f, 0xf6, 0x6d, 0x8f, 0x46, 0xf9, 0x21, 0x94,
	0x05, 0xf9, 0xaf, 0xfb, 0x1e, 0x65, 0xc4, 0xd9, 0x2b, 0x4e, 0x15, 0xc3, 0x84, 0xf0, 0x7e, 0x4d,
	0x7e, 0xd6, 0x37, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xaa, 0xe5, 0x4e, 0x7d, 0x08, 0x00,
	0x00,
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
	// Create custodian
	CreateCustodian(ctx context.Context, in *CreateCustodianRequest, opts ...grpc.CallOption) (*CreateCustodianResponse, error)
	// Update custodian
	UpdateCustodian(ctx context.Context, in *UpdateCustodianRequest, opts ...grpc.CallOption) (*UpdateCustodianResponse, error)
	// Create custodian group
	CreateCustodianGroup(ctx context.Context, in *CreateCustodianGroupRequest, opts ...grpc.CallOption) (*CreateCustodianGroupResponse, error)
	// Update Custodian group
	UpdateCustodianGroup(ctx context.Context, in *UpdateCustodianGroupRequest, opts ...grpc.CallOption) (*UpdateCustodianGroupResponse, error)
	// Add Custodian to custodian group
	// recalculate taproot pubkey when adding custodian to custodian group
	AddCustodianToGroup(ctx context.Context, in *AddCustodianToGroupRequest, opts ...grpc.CallOption) (*CustodianToGroupResponse, error)
	// Remove Custodian from custodian group
	// recalculate taproot address when deleting custodian from custodian group
	RemoveCustodianFromGroup(ctx context.Context, in *RemoveCustodianFromGroupRequest, opts ...grpc.CallOption) (*CustodianToGroupResponse, error)
	RotateKey(ctx context.Context, in *RotateKeyRequest, opts ...grpc.CallOption) (*RotateKeyResponse, error)
	SubmitTapScriptSigs(ctx context.Context, in *SubmitTapScriptSigsRequest, opts ...grpc.CallOption) (*SubmitTapScriptSigsResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) CreateCustodian(ctx context.Context, in *CreateCustodianRequest, opts ...grpc.CallOption) (*CreateCustodianResponse, error) {
	out := new(CreateCustodianResponse)
	err := c.cc.Invoke(ctx, "/scalar.covenant.v1beta1.MsgService/CreateCustodian", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) UpdateCustodian(ctx context.Context, in *UpdateCustodianRequest, opts ...grpc.CallOption) (*UpdateCustodianResponse, error) {
	out := new(UpdateCustodianResponse)
	err := c.cc.Invoke(ctx, "/scalar.covenant.v1beta1.MsgService/UpdateCustodian", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) CreateCustodianGroup(ctx context.Context, in *CreateCustodianGroupRequest, opts ...grpc.CallOption) (*CreateCustodianGroupResponse, error) {
	out := new(CreateCustodianGroupResponse)
	err := c.cc.Invoke(ctx, "/scalar.covenant.v1beta1.MsgService/CreateCustodianGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) UpdateCustodianGroup(ctx context.Context, in *UpdateCustodianGroupRequest, opts ...grpc.CallOption) (*UpdateCustodianGroupResponse, error) {
	out := new(UpdateCustodianGroupResponse)
	err := c.cc.Invoke(ctx, "/scalar.covenant.v1beta1.MsgService/UpdateCustodianGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) AddCustodianToGroup(ctx context.Context, in *AddCustodianToGroupRequest, opts ...grpc.CallOption) (*CustodianToGroupResponse, error) {
	out := new(CustodianToGroupResponse)
	err := c.cc.Invoke(ctx, "/scalar.covenant.v1beta1.MsgService/AddCustodianToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RemoveCustodianFromGroup(ctx context.Context, in *RemoveCustodianFromGroupRequest, opts ...grpc.CallOption) (*CustodianToGroupResponse, error) {
	out := new(CustodianToGroupResponse)
	err := c.cc.Invoke(ctx, "/scalar.covenant.v1beta1.MsgService/RemoveCustodianFromGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RotateKey(ctx context.Context, in *RotateKeyRequest, opts ...grpc.CallOption) (*RotateKeyResponse, error) {
	out := new(RotateKeyResponse)
	err := c.cc.Invoke(ctx, "/scalar.covenant.v1beta1.MsgService/RotateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SubmitTapScriptSigs(ctx context.Context, in *SubmitTapScriptSigsRequest, opts ...grpc.CallOption) (*SubmitTapScriptSigsResponse, error) {
	out := new(SubmitTapScriptSigsResponse)
	err := c.cc.Invoke(ctx, "/scalar.covenant.v1beta1.MsgService/SubmitTapScriptSigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	// Create custodian
	CreateCustodian(context.Context, *CreateCustodianRequest) (*CreateCustodianResponse, error)
	// Update custodian
	UpdateCustodian(context.Context, *UpdateCustodianRequest) (*UpdateCustodianResponse, error)
	// Create custodian group
	CreateCustodianGroup(context.Context, *CreateCustodianGroupRequest) (*CreateCustodianGroupResponse, error)
	// Update Custodian group
	UpdateCustodianGroup(context.Context, *UpdateCustodianGroupRequest) (*UpdateCustodianGroupResponse, error)
	// Add Custodian to custodian group
	// recalculate taproot pubkey when adding custodian to custodian group
	AddCustodianToGroup(context.Context, *AddCustodianToGroupRequest) (*CustodianToGroupResponse, error)
	// Remove Custodian from custodian group
	// recalculate taproot address when deleting custodian from custodian group
	RemoveCustodianFromGroup(context.Context, *RemoveCustodianFromGroupRequest) (*CustodianToGroupResponse, error)
	RotateKey(context.Context, *RotateKeyRequest) (*RotateKeyResponse, error)
	SubmitTapScriptSigs(context.Context, *SubmitTapScriptSigsRequest) (*SubmitTapScriptSigsResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) CreateCustodian(ctx context.Context, req *CreateCustodianRequest) (*CreateCustodianResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustodian not implemented")
}
func (*UnimplementedMsgServiceServer) UpdateCustodian(ctx context.Context, req *UpdateCustodianRequest) (*UpdateCustodianResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustodian not implemented")
}
func (*UnimplementedMsgServiceServer) CreateCustodianGroup(ctx context.Context, req *CreateCustodianGroupRequest) (*CreateCustodianGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustodianGroup not implemented")
}
func (*UnimplementedMsgServiceServer) UpdateCustodianGroup(ctx context.Context, req *UpdateCustodianGroupRequest) (*UpdateCustodianGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustodianGroup not implemented")
}
func (*UnimplementedMsgServiceServer) AddCustodianToGroup(ctx context.Context, req *AddCustodianToGroupRequest) (*CustodianToGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCustodianToGroup not implemented")
}
func (*UnimplementedMsgServiceServer) RemoveCustodianFromGroup(ctx context.Context, req *RemoveCustodianFromGroupRequest) (*CustodianToGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCustodianFromGroup not implemented")
}
func (*UnimplementedMsgServiceServer) RotateKey(ctx context.Context, req *RotateKeyRequest) (*RotateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateKey not implemented")
}
func (*UnimplementedMsgServiceServer) SubmitTapScriptSigs(ctx context.Context, req *SubmitTapScriptSigsRequest) (*SubmitTapScriptSigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTapScriptSigs not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_CreateCustodian_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustodianRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).CreateCustodian(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.covenant.v1beta1.MsgService/CreateCustodian",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).CreateCustodian(ctx, req.(*CreateCustodianRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_UpdateCustodian_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustodianRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).UpdateCustodian(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.covenant.v1beta1.MsgService/UpdateCustodian",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).UpdateCustodian(ctx, req.(*UpdateCustodianRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_CreateCustodianGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustodianGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).CreateCustodianGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.covenant.v1beta1.MsgService/CreateCustodianGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).CreateCustodianGroup(ctx, req.(*CreateCustodianGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_UpdateCustodianGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustodianGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).UpdateCustodianGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.covenant.v1beta1.MsgService/UpdateCustodianGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).UpdateCustodianGroup(ctx, req.(*UpdateCustodianGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_AddCustodianToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCustodianToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).AddCustodianToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.covenant.v1beta1.MsgService/AddCustodianToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).AddCustodianToGroup(ctx, req.(*AddCustodianToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RemoveCustodianFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCustodianFromGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RemoveCustodianFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.covenant.v1beta1.MsgService/RemoveCustodianFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RemoveCustodianFromGroup(ctx, req.(*RemoveCustodianFromGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RotateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RotateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RotateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.covenant.v1beta1.MsgService/RotateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RotateKey(ctx, req.(*RotateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SubmitTapScriptSigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTapScriptSigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SubmitTapScriptSigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.covenant.v1beta1.MsgService/SubmitTapScriptSigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SubmitTapScriptSigs(ctx, req.(*SubmitTapScriptSigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scalar.covenant.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustodian",
			Handler:    _MsgService_CreateCustodian_Handler,
		},
		{
			MethodName: "UpdateCustodian",
			Handler:    _MsgService_UpdateCustodian_Handler,
		},
		{
			MethodName: "CreateCustodianGroup",
			Handler:    _MsgService_CreateCustodianGroup_Handler,
		},
		{
			MethodName: "UpdateCustodianGroup",
			Handler:    _MsgService_UpdateCustodianGroup_Handler,
		},
		{
			MethodName: "AddCustodianToGroup",
			Handler:    _MsgService_AddCustodianToGroup_Handler,
		},
		{
			MethodName: "RemoveCustodianFromGroup",
			Handler:    _MsgService_RemoveCustodianFromGroup_Handler,
		},
		{
			MethodName: "RotateKey",
			Handler:    _MsgService_RotateKey_Handler,
		},
		{
			MethodName: "SubmitTapScriptSigs",
			Handler:    _MsgService_SubmitTapScriptSigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scalar/covenant/v1beta1/service.proto",
}

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// Get custodians
	Custodians(ctx context.Context, in *CustodiansRequest, opts ...grpc.CallOption) (*CustodiansResponse, error)
	// Get custodian groups
	Groups(ctx context.Context, in *GroupsRequest, opts ...grpc.CallOption) (*GroupsResponse, error)
	Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) Custodians(ctx context.Context, in *CustodiansRequest, opts ...grpc.CallOption) (*CustodiansResponse, error) {
	out := new(CustodiansResponse)
	err := c.cc.Invoke(ctx, "/scalar.covenant.v1beta1.QueryService/Custodians", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) Groups(ctx context.Context, in *GroupsRequest, opts ...grpc.CallOption) (*GroupsResponse, error) {
	out := new(GroupsResponse)
	err := c.cc.Invoke(ctx, "/scalar.covenant.v1beta1.QueryService/Groups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error) {
	out := new(ParamsResponse)
	err := c.cc.Invoke(ctx, "/scalar.covenant.v1beta1.QueryService/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// Get custodians
	Custodians(context.Context, *CustodiansRequest) (*CustodiansResponse, error)
	// Get custodian groups
	Groups(context.Context, *GroupsRequest) (*GroupsResponse, error)
	Params(context.Context, *ParamsRequest) (*ParamsResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) Custodians(ctx context.Context, req *CustodiansRequest) (*CustodiansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Custodians not implemented")
}
func (*UnimplementedQueryServiceServer) Groups(ctx context.Context, req *GroupsRequest) (*GroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Groups not implemented")
}
func (*UnimplementedQueryServiceServer) Params(ctx context.Context, req *ParamsRequest) (*ParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_Custodians_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustodiansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Custodians(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.covenant.v1beta1.QueryService/Custodians",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Custodians(ctx, req.(*CustodiansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_Groups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Groups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.covenant.v1beta1.QueryService/Groups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Groups(ctx, req.(*GroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scalar.covenant.v1beta1.QueryService/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Params(ctx, req.(*ParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scalar.covenant.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Custodians",
			Handler:    _QueryService_Custodians_Handler,
		},
		{
			MethodName: "Groups",
			Handler:    _QueryService_Groups_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _QueryService_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scalar/covenant/v1beta1/service.proto",
}
