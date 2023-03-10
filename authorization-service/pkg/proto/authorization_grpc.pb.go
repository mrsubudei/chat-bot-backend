// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authorization

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

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationClient interface {
	SignUp(ctx context.Context, in *UserSingle, opts ...grpc.CallOption) (*IdRequest, error)
	SignIn(ctx context.Context, in *UserSingle, opts ...grpc.CallOption) (*Empty, error)
	VerifyRegistration(ctx context.Context, in *UserSingle, opts ...grpc.CallOption) (*Empty, error)
	GetByPhone(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*UserSingle, error)
	GetByEmail(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*UserSingle, error)
	GetById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserSingle, error)
	GetBySessionToken(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*UserSingle, error)
}

type authorizationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationClient(cc grpc.ClientConnInterface) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) SignUp(ctx context.Context, in *UserSingle, opts ...grpc.CallOption) (*IdRequest, error) {
	out := new(IdRequest)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) SignIn(ctx context.Context, in *UserSingle, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) VerifyRegistration(ctx context.Context, in *UserSingle, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/VerifyRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) GetByPhone(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*UserSingle, error) {
	out := new(UserSingle)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/GetByPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) GetByEmail(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*UserSingle, error) {
	out := new(UserSingle)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/GetByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) GetById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserSingle, error) {
	out := new(UserSingle)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) GetBySessionToken(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*UserSingle, error) {
	out := new(UserSingle)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/GetBySessionToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
// All implementations must embed UnimplementedAuthorizationServer
// for forward compatibility
type AuthorizationServer interface {
	SignUp(context.Context, *UserSingle) (*IdRequest, error)
	SignIn(context.Context, *UserSingle) (*Empty, error)
	VerifyRegistration(context.Context, *UserSingle) (*Empty, error)
	GetByPhone(context.Context, *StringRequest) (*UserSingle, error)
	GetByEmail(context.Context, *StringRequest) (*UserSingle, error)
	GetById(context.Context, *IdRequest) (*UserSingle, error)
	GetBySessionToken(context.Context, *StringRequest) (*UserSingle, error)
	mustEmbedUnimplementedAuthorizationServer()
}

// UnimplementedAuthorizationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (UnimplementedAuthorizationServer) SignUp(context.Context, *UserSingle) (*IdRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAuthorizationServer) SignIn(context.Context, *UserSingle) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthorizationServer) VerifyRegistration(context.Context, *UserSingle) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyRegistration not implemented")
}
func (UnimplementedAuthorizationServer) GetByPhone(context.Context, *StringRequest) (*UserSingle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByPhone not implemented")
}
func (UnimplementedAuthorizationServer) GetByEmail(context.Context, *StringRequest) (*UserSingle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByEmail not implemented")
}
func (UnimplementedAuthorizationServer) GetById(context.Context, *IdRequest) (*UserSingle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedAuthorizationServer) GetBySessionToken(context.Context, *StringRequest) (*UserSingle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBySessionToken not implemented")
}
func (UnimplementedAuthorizationServer) mustEmbedUnimplementedAuthorizationServer() {}

// UnsafeAuthorizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationServer will
// result in compilation errors.
type UnsafeAuthorizationServer interface {
	mustEmbedUnimplementedAuthorizationServer()
}

func RegisterAuthorizationServer(s grpc.ServiceRegistrar, srv AuthorizationServer) {
	s.RegisterService(&Authorization_ServiceDesc, srv)
}

func _Authorization_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSingle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).SignUp(ctx, req.(*UserSingle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSingle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).SignIn(ctx, req.(*UserSingle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_VerifyRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSingle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).VerifyRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/VerifyRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).VerifyRegistration(ctx, req.(*UserSingle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_GetByPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).GetByPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/GetByPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).GetByPhone(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_GetByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).GetByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/GetByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).GetByEmail(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).GetById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_GetBySessionToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).GetBySessionToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/GetBySessionToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).GetBySessionToken(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authorization_ServiceDesc is the grpc.ServiceDesc for Authorization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authorization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authorization.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Authorization_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Authorization_SignIn_Handler,
		},
		{
			MethodName: "VerifyRegistration",
			Handler:    _Authorization_VerifyRegistration_Handler,
		},
		{
			MethodName: "GetByPhone",
			Handler:    _Authorization_GetByPhone_Handler,
		},
		{
			MethodName: "GetByEmail",
			Handler:    _Authorization_GetByEmail_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _Authorization_GetById_Handler,
		},
		{
			MethodName: "GetBySessionToken",
			Handler:    _Authorization_GetBySessionToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/authorization.proto",
}
