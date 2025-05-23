// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--dev
// source: api/authn/service/v1/authn.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Authn_Login_FullMethodName    = "/api.authn.service.v1.Authn/Login"
	Authn_Register_FullMethodName = "/api.authn.service.v1.Authn/Register"
	Authn_Verify_FullMethodName   = "/api.authn.service.v1.Authn/Verify"
)

// AuthnClient is the client API for Authn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthnClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthReply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*AuthReply, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyReply, error)
}

type authnClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthnClient(cc grpc.ClientConnInterface) AuthnClient {
	return &authnClient{cc}
}

func (c *authnClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, Authn_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, Authn_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyReply)
	err := c.cc.Invoke(ctx, Authn_Verify_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthnServer is the server API for Authn service.
// All implementations must embed UnimplementedAuthnServer
// for forward compatibility.
type AuthnServer interface {
	Login(context.Context, *LoginRequest) (*AuthReply, error)
	Register(context.Context, *RegisterRequest) (*AuthReply, error)
	Verify(context.Context, *VerifyRequest) (*VerifyReply, error)
	mustEmbedUnimplementedAuthnServer()
}

// UnimplementedAuthnServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthnServer struct{}

func (UnimplementedAuthnServer) Login(context.Context, *LoginRequest) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthnServer) Register(context.Context, *RegisterRequest) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthnServer) Verify(context.Context, *VerifyRequest) (*VerifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedAuthnServer) mustEmbedUnimplementedAuthnServer() {}
func (UnimplementedAuthnServer) testEmbeddedByValue()               {}

// UnsafeAuthnServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthnServer will
// result in compilation errors.
type UnsafeAuthnServer interface {
	mustEmbedUnimplementedAuthnServer()
}

func RegisterAuthnServer(s grpc.ServiceRegistrar, srv AuthnServer) {
	// If the following call pancis, it indicates UnimplementedAuthnServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Authn_ServiceDesc, srv)
}

func _Authn_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authn_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authn_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authn_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authn_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authn_Verify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authn_ServiceDesc is the grpc.ServiceDesc for Authn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.authn.service.v1.Authn",
	HandlerType: (*AuthnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Authn_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Authn_Register_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Authn_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/authn/service/v1/authn.proto",
}
