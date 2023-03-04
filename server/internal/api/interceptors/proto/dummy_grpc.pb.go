// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: dummy.proto

package dummy

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

// DummyServiceClient is the client API for DummyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DummyServiceClient interface {
	DoNothing(ctx context.Context, in *DummyRequest, opts ...grpc.CallOption) (*DummyResponse, error)
}

type dummyServiceClient struct {
	cc grpc.ClientConnInterface
}

// NewDummyServiceClient.
func NewDummyServiceClient(cc grpc.ClientConnInterface) DummyServiceClient {
	return &dummyServiceClient{cc}
}

// DoNothing.
func (c *dummyServiceClient) DoNothing(ctx context.Context, in *DummyRequest, opts ...grpc.CallOption) (*DummyResponse, error) {
	out := new(DummyResponse)
	err := c.cc.Invoke(ctx, "/proto.DummyService/DoNothing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DummyServiceServer is the server API for DummyService service.
// All implementations must embed UnimplementedDummyServiceServer
// for forward compatibility
type DummyServiceServer interface {
	DoNothing(context.Context, *DummyRequest) (*DummyResponse, error)
	mustEmbedUnimplementedDummyServiceServer()
}

// UnimplementedDummyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDummyServiceServer struct {
}

// DoNothing.
func (UnimplementedDummyServiceServer) DoNothing(context.Context, *DummyRequest) (*DummyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoNothing not implemented")
}
func (UnimplementedDummyServiceServer) mustEmbedUnimplementedDummyServiceServer() {}

// UnsafeDummyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DummyServiceServer will
// result in compilation errors.
type UnsafeDummyServiceServer interface {
	mustEmbedUnimplementedDummyServiceServer()
}

// RegisterDummyServiceServer.
func RegisterDummyServiceServer(s grpc.ServiceRegistrar, srv DummyServiceServer) {
	s.RegisterService(&DummyService_ServiceDesc, srv)
}

func _DummyService_DoNothing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DummyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyServiceServer).DoNothing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DummyService/DoNothing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyServiceServer).DoNothing(ctx, req.(*DummyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DummyService_ServiceDesc is the grpc.ServiceDesc for DummyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DummyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DummyService",
	HandlerType: (*DummyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoNothing",
			Handler:    _DummyService_DoNothing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dummy.proto",
}
