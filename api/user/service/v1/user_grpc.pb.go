// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--dev
// source: api/user/service/v1/user.proto

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
	User_CreateUser_FullMethodName              = "/api.user.service.v1.User/CreateUser"
	User_UpdateUser_FullMethodName              = "/api.user.service.v1.User/UpdateUser"
	User_DeleteUser_FullMethodName              = "/api.user.service.v1.User/DeleteUser"
	User_GetUser_FullMethodName                 = "/api.user.service.v1.User/GetUser"
	User_ListUser_FullMethodName                = "/api.user.service.v1.User/ListUser"
	User_GetUserByUsername_FullMethodName       = "/api.user.service.v1.User/GetUserByUsername"
	User_CreateUserDomain_FullMethodName        = "/api.user.service.v1.User/CreateUserDomain"
	User_ListUserByParentId_FullMethodName      = "/api.user.service.v1.User/ListUserByParentId"
	User_UpdateUserLastLoginTime_FullMethodName = "/api.user.service.v1.User/UpdateUserLastLoginTime"
	User_GetUserDomain_FullMethodName           = "/api.user.service.v1.User/GetUserDomain"
	User_ListUserDomains_FullMethodName         = "/api.user.service.v1.User/ListUserDomains"
	User_ListUserDomainsByUserId_FullMethodName = "/api.user.service.v1.User/ListUserDomainsByUserId"
	User_DeleteDomain_FullMethodName            = "/api.user.service.v1.User/DeleteDomain"
	User_GetUserByDomain_FullMethodName         = "/api.user.service.v1.User/GetUserByDomain"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserReply, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error)
	GetUserByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	CreateUserDomain(ctx context.Context, in *CreateUserDomainRequest, opts ...grpc.CallOption) (*CreateUserDomainReply, error)
	ListUserByParentId(ctx context.Context, in *ListUserByParentIdReq, opts ...grpc.CallOption) (*ListUserByParentIdReply, error)
	UpdateUserLastLoginTime(ctx context.Context, in *UpdateUserLastLoginTimeReq, opts ...grpc.CallOption) (*UpdateUserLastLoginTimeReply, error)
	// 获取域名
	GetUserDomain(ctx context.Context, in *GetUserDomainRequest, opts ...grpc.CallOption) (*GetUserDomainReply, error)
	ListUserDomains(ctx context.Context, in *ListUserDomainsRequest, opts ...grpc.CallOption) (*ListUserDomainsReply, error)
	ListUserDomainsByUserId(ctx context.Context, in *ListUserDomainsByUserIdRequest, opts ...grpc.CallOption) (*ListUserDomainsByUserIdReply, error)
	DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*DeleteDomainReply, error)
	GetUserByDomain(ctx context.Context, in *GetUserByDomainRequest, opts ...grpc.CallOption) (*GetUserByDomainReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserReply)
	err := c.cc.Invoke(ctx, User_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserReply)
	err := c.cc.Invoke(ctx, User_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserReply)
	err := c.cc.Invoke(ctx, User_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, User_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListUser(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserReply)
	err := c.cc.Invoke(ctx, User_ListUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, User_GetUserByUsername_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateUserDomain(ctx context.Context, in *CreateUserDomainRequest, opts ...grpc.CallOption) (*CreateUserDomainReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserDomainReply)
	err := c.cc.Invoke(ctx, User_CreateUserDomain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListUserByParentId(ctx context.Context, in *ListUserByParentIdReq, opts ...grpc.CallOption) (*ListUserByParentIdReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserByParentIdReply)
	err := c.cc.Invoke(ctx, User_ListUserByParentId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserLastLoginTime(ctx context.Context, in *UpdateUserLastLoginTimeReq, opts ...grpc.CallOption) (*UpdateUserLastLoginTimeReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserLastLoginTimeReply)
	err := c.cc.Invoke(ctx, User_UpdateUserLastLoginTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserDomain(ctx context.Context, in *GetUserDomainRequest, opts ...grpc.CallOption) (*GetUserDomainReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserDomainReply)
	err := c.cc.Invoke(ctx, User_GetUserDomain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListUserDomains(ctx context.Context, in *ListUserDomainsRequest, opts ...grpc.CallOption) (*ListUserDomainsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserDomainsReply)
	err := c.cc.Invoke(ctx, User_ListUserDomains_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListUserDomainsByUserId(ctx context.Context, in *ListUserDomainsByUserIdRequest, opts ...grpc.CallOption) (*ListUserDomainsByUserIdReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserDomainsByUserIdReply)
	err := c.cc.Invoke(ctx, User_ListUserDomainsByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*DeleteDomainReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDomainReply)
	err := c.cc.Invoke(ctx, User_DeleteDomain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserByDomain(ctx context.Context, in *GetUserByDomainRequest, opts ...grpc.CallOption) (*GetUserByDomainReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserByDomainReply)
	err := c.cc.Invoke(ctx, User_GetUserByDomain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility.
type UserServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	ListUser(context.Context, *ListUserRequest) (*ListUserReply, error)
	GetUserByUsername(context.Context, *GetUserByUsernameRequest) (*GetUserReply, error)
	CreateUserDomain(context.Context, *CreateUserDomainRequest) (*CreateUserDomainReply, error)
	ListUserByParentId(context.Context, *ListUserByParentIdReq) (*ListUserByParentIdReply, error)
	UpdateUserLastLoginTime(context.Context, *UpdateUserLastLoginTimeReq) (*UpdateUserLastLoginTimeReply, error)
	// 获取域名
	GetUserDomain(context.Context, *GetUserDomainRequest) (*GetUserDomainReply, error)
	ListUserDomains(context.Context, *ListUserDomainsRequest) (*ListUserDomainsReply, error)
	ListUserDomainsByUserId(context.Context, *ListUserDomainsByUserIdRequest) (*ListUserDomainsByUserIdReply, error)
	DeleteDomain(context.Context, *DeleteDomainRequest) (*DeleteDomainReply, error)
	GetUserByDomain(context.Context, *GetUserByDomainRequest) (*GetUserByDomainReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServer struct{}

func (UnimplementedUserServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServer) GetUser(context.Context, *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) ListUser(context.Context, *ListUserRequest) (*ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUser not implemented")
}
func (UnimplementedUserServer) GetUserByUsername(context.Context, *GetUserByUsernameRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUsername not implemented")
}
func (UnimplementedUserServer) CreateUserDomain(context.Context, *CreateUserDomainRequest) (*CreateUserDomainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserDomain not implemented")
}
func (UnimplementedUserServer) ListUserByParentId(context.Context, *ListUserByParentIdReq) (*ListUserByParentIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserByParentId not implemented")
}
func (UnimplementedUserServer) UpdateUserLastLoginTime(context.Context, *UpdateUserLastLoginTimeReq) (*UpdateUserLastLoginTimeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserLastLoginTime not implemented")
}
func (UnimplementedUserServer) GetUserDomain(context.Context, *GetUserDomainRequest) (*GetUserDomainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDomain not implemented")
}
func (UnimplementedUserServer) ListUserDomains(context.Context, *ListUserDomainsRequest) (*ListUserDomainsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserDomains not implemented")
}
func (UnimplementedUserServer) ListUserDomainsByUserId(context.Context, *ListUserDomainsByUserIdRequest) (*ListUserDomainsByUserIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserDomainsByUserId not implemented")
}
func (UnimplementedUserServer) DeleteDomain(context.Context, *DeleteDomainRequest) (*DeleteDomainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDomain not implemented")
}
func (UnimplementedUserServer) GetUserByDomain(context.Context, *GetUserByDomainRequest) (*GetUserByDomainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByDomain not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}
func (UnimplementedUserServer) testEmbeddedByValue()              {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	// If the following call pancis, it indicates UnimplementedUserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ListUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListUser(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByUsername(ctx, req.(*GetUserByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateUserDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUserDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CreateUserDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUserDomain(ctx, req.(*CreateUserDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListUserByParentId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserByParentIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListUserByParentId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ListUserByParentId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListUserByParentId(ctx, req.(*ListUserByParentIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserLastLoginTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserLastLoginTimeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserLastLoginTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateUserLastLoginTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserLastLoginTime(ctx, req.(*UpdateUserLastLoginTimeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserDomain(ctx, req.(*GetUserDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListUserDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListUserDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ListUserDomains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListUserDomains(ctx, req.(*ListUserDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListUserDomainsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserDomainsByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListUserDomainsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ListUserDomainsByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListUserDomainsByUserId(ctx, req.(*ListUserDomainsByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeleteDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteDomain(ctx, req.(*DeleteDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserByDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserByDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserByDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserByDomain(ctx, req.(*GetUserByDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.service.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _User_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _User_DeleteUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "ListUser",
			Handler:    _User_ListUser_Handler,
		},
		{
			MethodName: "GetUserByUsername",
			Handler:    _User_GetUserByUsername_Handler,
		},
		{
			MethodName: "CreateUserDomain",
			Handler:    _User_CreateUserDomain_Handler,
		},
		{
			MethodName: "ListUserByParentId",
			Handler:    _User_ListUserByParentId_Handler,
		},
		{
			MethodName: "UpdateUserLastLoginTime",
			Handler:    _User_UpdateUserLastLoginTime_Handler,
		},
		{
			MethodName: "GetUserDomain",
			Handler:    _User_GetUserDomain_Handler,
		},
		{
			MethodName: "ListUserDomains",
			Handler:    _User_ListUserDomains_Handler,
		},
		{
			MethodName: "ListUserDomainsByUserId",
			Handler:    _User_ListUserDomainsByUserId_Handler,
		},
		{
			MethodName: "DeleteDomain",
			Handler:    _User_DeleteDomain_Handler,
		},
		{
			MethodName: "GetUserByDomain",
			Handler:    _User_GetUserByDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/service/v1/user.proto",
}
