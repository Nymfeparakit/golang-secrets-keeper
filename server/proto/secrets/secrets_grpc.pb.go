// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: secrets.proto

package secrets

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

// SecretsManagementClient is the client API for SecretsManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretsManagementClient interface {
	AddCredentials(ctx context.Context, in *Password, opts ...grpc.CallOption) (*AddResponse, error)
	AddCardInfo(ctx context.Context, in *CardInfo, opts ...grpc.CallOption) (*AddResponse, error)
	AddTextInfo(ctx context.Context, in *TextInfo, opts ...grpc.CallOption) (*AddResponse, error)
	AddBinaryInfo(ctx context.Context, in *BinaryInfo, opts ...grpc.CallOption) (*AddResponse, error)
	ListSecrets(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListSecretResponse, error)
	GetCredentialsByID(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetCredentialsResponse, error)
	GetCardByID(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetCardResponse, error)
	GetTextByID(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetTextResponse, error)
	GetBinaryByID(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetBinaryResponse, error)
	UpdateCredentials(ctx context.Context, in *Password, opts ...grpc.CallOption) (*EmptyResponse, error)
	UpdateCardInfo(ctx context.Context, in *CardInfo, opts ...grpc.CallOption) (*EmptyResponse, error)
	UpdateTextInfo(ctx context.Context, in *TextInfo, opts ...grpc.CallOption) (*EmptyResponse, error)
	UpdateBinaryInfo(ctx context.Context, in *BinaryInfo, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type secretsManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretsManagementClient(cc grpc.ClientConnInterface) SecretsManagementClient {
	return &secretsManagementClient{cc}
}

func (c *secretsManagementClient) AddCredentials(ctx context.Context, in *Password, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/AddCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsManagementClient) AddCardInfo(ctx context.Context, in *CardInfo, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/AddCardInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsManagementClient) AddTextInfo(ctx context.Context, in *TextInfo, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/AddTextInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsManagementClient) AddBinaryInfo(ctx context.Context, in *BinaryInfo, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/AddBinaryInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsManagementClient) ListSecrets(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListSecretResponse, error) {
	out := new(ListSecretResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/ListSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsManagementClient) GetCredentialsByID(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetCredentialsResponse, error) {
	out := new(GetCredentialsResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/GetCredentialsByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsManagementClient) GetCardByID(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetCardResponse, error) {
	out := new(GetCardResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/GetCardByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsManagementClient) GetTextByID(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetTextResponse, error) {
	out := new(GetTextResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/GetTextByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsManagementClient) GetBinaryByID(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetBinaryResponse, error) {
	out := new(GetBinaryResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/GetBinaryByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsManagementClient) UpdateCredentials(ctx context.Context, in *Password, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/UpdateCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsManagementClient) UpdateCardInfo(ctx context.Context, in *CardInfo, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/UpdateCardInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsManagementClient) UpdateTextInfo(ctx context.Context, in *TextInfo, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/UpdateTextInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretsManagementClient) UpdateBinaryInfo(ctx context.Context, in *BinaryInfo, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/proto.SecretsManagement/UpdateBinaryInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretsManagementServer is the server API for SecretsManagement service.
// All implementations must embed UnimplementedSecretsManagementServer
// for forward compatibility
type SecretsManagementServer interface {
	AddCredentials(context.Context, *Password) (*AddResponse, error)
	AddCardInfo(context.Context, *CardInfo) (*AddResponse, error)
	AddTextInfo(context.Context, *TextInfo) (*AddResponse, error)
	AddBinaryInfo(context.Context, *BinaryInfo) (*AddResponse, error)
	ListSecrets(context.Context, *EmptyRequest) (*ListSecretResponse, error)
	GetCredentialsByID(context.Context, *GetSecretRequest) (*GetCredentialsResponse, error)
	GetCardByID(context.Context, *GetSecretRequest) (*GetCardResponse, error)
	GetTextByID(context.Context, *GetSecretRequest) (*GetTextResponse, error)
	GetBinaryByID(context.Context, *GetSecretRequest) (*GetBinaryResponse, error)
	UpdateCredentials(context.Context, *Password) (*EmptyResponse, error)
	UpdateCardInfo(context.Context, *CardInfo) (*EmptyResponse, error)
	UpdateTextInfo(context.Context, *TextInfo) (*EmptyResponse, error)
	UpdateBinaryInfo(context.Context, *BinaryInfo) (*EmptyResponse, error)
	mustEmbedUnimplementedSecretsManagementServer()
}

// UnimplementedSecretsManagementServer must be embedded to have forward compatible implementations.
type UnimplementedSecretsManagementServer struct {
}

func (UnimplementedSecretsManagementServer) AddCredentials(context.Context, *Password) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCredentials not implemented")
}
func (UnimplementedSecretsManagementServer) AddCardInfo(context.Context, *CardInfo) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCardInfo not implemented")
}
func (UnimplementedSecretsManagementServer) AddTextInfo(context.Context, *TextInfo) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTextInfo not implemented")
}
func (UnimplementedSecretsManagementServer) AddBinaryInfo(context.Context, *BinaryInfo) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBinaryInfo not implemented")
}
func (UnimplementedSecretsManagementServer) ListSecrets(context.Context, *EmptyRequest) (*ListSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecrets not implemented")
}
func (UnimplementedSecretsManagementServer) GetCredentialsByID(context.Context, *GetSecretRequest) (*GetCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCredentialsByID not implemented")
}
func (UnimplementedSecretsManagementServer) GetCardByID(context.Context, *GetSecretRequest) (*GetCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardByID not implemented")
}
func (UnimplementedSecretsManagementServer) GetTextByID(context.Context, *GetSecretRequest) (*GetTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTextByID not implemented")
}
func (UnimplementedSecretsManagementServer) GetBinaryByID(context.Context, *GetSecretRequest) (*GetBinaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBinaryByID not implemented")
}
func (UnimplementedSecretsManagementServer) UpdateCredentials(context.Context, *Password) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCredentials not implemented")
}
func (UnimplementedSecretsManagementServer) UpdateCardInfo(context.Context, *CardInfo) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCardInfo not implemented")
}
func (UnimplementedSecretsManagementServer) UpdateTextInfo(context.Context, *TextInfo) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTextInfo not implemented")
}
func (UnimplementedSecretsManagementServer) UpdateBinaryInfo(context.Context, *BinaryInfo) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBinaryInfo not implemented")
}
func (UnimplementedSecretsManagementServer) mustEmbedUnimplementedSecretsManagementServer() {}

// UnsafeSecretsManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretsManagementServer will
// result in compilation errors.
type UnsafeSecretsManagementServer interface {
	mustEmbedUnimplementedSecretsManagementServer()
}

func RegisterSecretsManagementServer(s grpc.ServiceRegistrar, srv SecretsManagementServer) {
	s.RegisterService(&SecretsManagement_ServiceDesc, srv)
}

func _SecretsManagement_AddCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Password)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).AddCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/AddCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).AddCredentials(ctx, req.(*Password))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsManagement_AddCardInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).AddCardInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/AddCardInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).AddCardInfo(ctx, req.(*CardInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsManagement_AddTextInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).AddTextInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/AddTextInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).AddTextInfo(ctx, req.(*TextInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsManagement_AddBinaryInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).AddBinaryInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/AddBinaryInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).AddBinaryInfo(ctx, req.(*BinaryInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsManagement_ListSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).ListSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/ListSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).ListSecrets(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsManagement_GetCredentialsByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).GetCredentialsByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/GetCredentialsByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).GetCredentialsByID(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsManagement_GetCardByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).GetCardByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/GetCardByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).GetCardByID(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsManagement_GetTextByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).GetTextByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/GetTextByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).GetTextByID(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsManagement_GetBinaryByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).GetBinaryByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/GetBinaryByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).GetBinaryByID(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsManagement_UpdateCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Password)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).UpdateCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/UpdateCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).UpdateCredentials(ctx, req.(*Password))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsManagement_UpdateCardInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).UpdateCardInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/UpdateCardInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).UpdateCardInfo(ctx, req.(*CardInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsManagement_UpdateTextInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).UpdateTextInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/UpdateTextInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).UpdateTextInfo(ctx, req.(*TextInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretsManagement_UpdateBinaryInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretsManagementServer).UpdateBinaryInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretsManagement/UpdateBinaryInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretsManagementServer).UpdateBinaryInfo(ctx, req.(*BinaryInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretsManagement_ServiceDesc is the grpc.ServiceDesc for SecretsManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretsManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SecretsManagement",
	HandlerType: (*SecretsManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCredentials",
			Handler:    _SecretsManagement_AddCredentials_Handler,
		},
		{
			MethodName: "AddCardInfo",
			Handler:    _SecretsManagement_AddCardInfo_Handler,
		},
		{
			MethodName: "AddTextInfo",
			Handler:    _SecretsManagement_AddTextInfo_Handler,
		},
		{
			MethodName: "AddBinaryInfo",
			Handler:    _SecretsManagement_AddBinaryInfo_Handler,
		},
		{
			MethodName: "ListSecrets",
			Handler:    _SecretsManagement_ListSecrets_Handler,
		},
		{
			MethodName: "GetCredentialsByID",
			Handler:    _SecretsManagement_GetCredentialsByID_Handler,
		},
		{
			MethodName: "GetCardByID",
			Handler:    _SecretsManagement_GetCardByID_Handler,
		},
		{
			MethodName: "GetTextByID",
			Handler:    _SecretsManagement_GetTextByID_Handler,
		},
		{
			MethodName: "GetBinaryByID",
			Handler:    _SecretsManagement_GetBinaryByID_Handler,
		},
		{
			MethodName: "UpdateCredentials",
			Handler:    _SecretsManagement_UpdateCredentials_Handler,
		},
		{
			MethodName: "UpdateCardInfo",
			Handler:    _SecretsManagement_UpdateCardInfo_Handler,
		},
		{
			MethodName: "UpdateTextInfo",
			Handler:    _SecretsManagement_UpdateTextInfo_Handler,
		},
		{
			MethodName: "UpdateBinaryInfo",
			Handler:    _SecretsManagement_UpdateBinaryInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secrets.proto",
}
