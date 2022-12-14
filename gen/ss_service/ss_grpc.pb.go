// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: proto/ss.proto

package ss_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SsServiceClient is the client API for SsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SsServiceClient interface {
	ActivateSsConnection(ctx context.Context, in *SsConnectionReq, opts ...grpc.CallOption) (*SsConnectionRes, error)
	DeactivateSsConnection(ctx context.Context, in *SsConnectionReq, opts ...grpc.CallOption) (*SsConnectionRes, error)
	SsConnectionStatus(ctx context.Context, in *SsConnectionReq, opts ...grpc.CallOption) (*SsConnectionRes, error)
	CheckSsPortAvailable(ctx context.Context, in *CheckSsPortAvailableReq, opts ...grpc.CallOption) (*CheckSsPortAvailableRes, error)
}

type ssServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSsServiceClient(cc grpc.ClientConnInterface) SsServiceClient {
	return &ssServiceClient{cc}
}

func (c *ssServiceClient) ActivateSsConnection(ctx context.Context, in *SsConnectionReq, opts ...grpc.CallOption) (*SsConnectionRes, error) {
	out := new(SsConnectionRes)
	err := c.cc.Invoke(ctx, "/SsService/ActivateSsConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssServiceClient) DeactivateSsConnection(ctx context.Context, in *SsConnectionReq, opts ...grpc.CallOption) (*SsConnectionRes, error) {
	out := new(SsConnectionRes)
	err := c.cc.Invoke(ctx, "/SsService/DeactivateSsConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssServiceClient) SsConnectionStatus(ctx context.Context, in *SsConnectionReq, opts ...grpc.CallOption) (*SsConnectionRes, error) {
	out := new(SsConnectionRes)
	err := c.cc.Invoke(ctx, "/SsService/SsConnectionStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssServiceClient) CheckSsPortAvailable(ctx context.Context, in *CheckSsPortAvailableReq, opts ...grpc.CallOption) (*CheckSsPortAvailableRes, error) {
	out := new(CheckSsPortAvailableRes)
	err := c.cc.Invoke(ctx, "/SsService/CheckSsPortAvailable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SsServiceServer is the server API for SsService service.
// All implementations must embed UnimplementedSsServiceServer
// for forward compatibility
type SsServiceServer interface {
	ActivateSsConnection(context.Context, *SsConnectionReq) (*SsConnectionRes, error)
	DeactivateSsConnection(context.Context, *SsConnectionReq) (*SsConnectionRes, error)
	SsConnectionStatus(context.Context, *SsConnectionReq) (*SsConnectionRes, error)
	CheckSsPortAvailable(context.Context, *CheckSsPortAvailableReq) (*CheckSsPortAvailableRes, error)
	mustEmbedUnimplementedSsServiceServer()
}

// UnimplementedSsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSsServiceServer struct {
}

func (UnimplementedSsServiceServer) ActivateSsConnection(context.Context, *SsConnectionReq) (*SsConnectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateSsConnection not implemented")
}
func (UnimplementedSsServiceServer) DeactivateSsConnection(context.Context, *SsConnectionReq) (*SsConnectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateSsConnection not implemented")
}
func (UnimplementedSsServiceServer) SsConnectionStatus(context.Context, *SsConnectionReq) (*SsConnectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SsConnectionStatus not implemented")
}
func (UnimplementedSsServiceServer) CheckSsPortAvailable(context.Context, *CheckSsPortAvailableReq) (*CheckSsPortAvailableRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSsPortAvailable not implemented")
}
func (UnimplementedSsServiceServer) mustEmbedUnimplementedSsServiceServer() {}

// UnsafeSsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SsServiceServer will
// result in compilation errors.
type UnsafeSsServiceServer interface {
	mustEmbedUnimplementedSsServiceServer()
}

func RegisterSsServiceServer(s grpc.ServiceRegistrar, srv SsServiceServer) {
	s.RegisterService(&SsService_ServiceDesc, srv)
}

func _SsService_ActivateSsConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsConnectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsServiceServer).ActivateSsConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SsService/ActivateSsConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsServiceServer).ActivateSsConnection(ctx, req.(*SsConnectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SsService_DeactivateSsConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsConnectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsServiceServer).DeactivateSsConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SsService/DeactivateSsConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsServiceServer).DeactivateSsConnection(ctx, req.(*SsConnectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SsService_SsConnectionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SsConnectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsServiceServer).SsConnectionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SsService/SsConnectionStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsServiceServer).SsConnectionStatus(ctx, req.(*SsConnectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SsService_CheckSsPortAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSsPortAvailableReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsServiceServer).CheckSsPortAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SsService/CheckSsPortAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsServiceServer).CheckSsPortAvailable(ctx, req.(*CheckSsPortAvailableReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SsService_ServiceDesc is the grpc.ServiceDesc for SsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SsService",
	HandlerType: (*SsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ActivateSsConnection",
			Handler:    _SsService_ActivateSsConnection_Handler,
		},
		{
			MethodName: "DeactivateSsConnection",
			Handler:    _SsService_DeactivateSsConnection_Handler,
		},
		{
			MethodName: "SsConnectionStatus",
			Handler:    _SsService_SsConnectionStatus_Handler,
		},
		{
			MethodName: "CheckSsPortAvailable",
			Handler:    _SsService_CheckSsPortAvailable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ss.proto",
}
