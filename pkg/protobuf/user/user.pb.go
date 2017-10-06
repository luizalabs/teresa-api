// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/protobuf/user/user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	pkg/protobuf/user/user.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	SetPasswordRequest
	DeleteRequest
	CreateRequest
	Empty
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Email     string  `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password  string  `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	ExpiresIn float64 `protobuf:"fixed64,3,opt,name=expires_in,json=expiresIn" json:"expires_in,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetExpiresIn() float64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SetPasswordRequest struct {
	Password string `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
}

func (m *SetPasswordRequest) Reset()                    { *m = SetPasswordRequest{} }
func (m *SetPasswordRequest) String() string            { return proto.CompactTextString(m) }
func (*SetPasswordRequest) ProtoMessage()               {}
func (*SetPasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SetPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type DeleteRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Admin    bool   `protobuf:"varint,4,opt,name=admin" json:"admin,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateRequest) GetAdmin() bool {
	if m != nil {
		return m.Admin
	}
	return false
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "user.LoginResponse")
	proto.RegisterType((*SetPasswordRequest)(nil), "user.SetPasswordRequest")
	proto.RegisterType((*DeleteRequest)(nil), "user.DeleteRequest")
	proto.RegisterType((*CreateRequest)(nil), "user.CreateRequest")
	proto.RegisterType((*Empty)(nil), "user.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for User service

type UserClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/user.User/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetPassword(ctx context.Context, in *SetPasswordRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.User/SetPassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.User/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/user.User/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	SetPassword(context.Context, *SetPasswordRequest) (*Empty, error)
	Delete(context.Context, *DeleteRequest) (*Empty, error)
	Create(context.Context, *CreateRequest) (*Empty, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetPassword(ctx, req.(*SetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "SetPassword",
			Handler:    _User_SetPassword_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _User_Delete_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _User_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protobuf/user/user.proto",
}

func init() { proto.RegisterFile("pkg/protobuf/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xd9, 0x36, 0xe9, 0xbf, 0x9d, 0xfe, 0x7b, 0x19, 0x7b, 0x08, 0x45, 0xa1, 0x04, 0x84,
	0xe0, 0xa1, 0x2d, 0xea, 0x1b, 0xa8, 0x07, 0xc1, 0x83, 0x44, 0x3c, 0x97, 0x94, 0x8e, 0x25, 0xa4,
	0xd9, 0x5d, 0x77, 0x37, 0xa8, 0x2f, 0xe8, 0x73, 0x49, 0x76, 0x37, 0x36, 0xb1, 0xe8, 0x25, 0xec,
	0x37, 0xfb, 0x65, 0xbe, 0x99, 0x1f, 0x0b, 0xa7, 0xb2, 0xd8, 0x2d, 0xa5, 0x12, 0x46, 0x6c, 0xaa,
	0x97, 0x65, 0xa5, 0x49, 0xd9, 0xcf, 0xc2, 0x96, 0x30, 0xa8, 0xcf, 0xf1, 0x1a, 0xfe, 0x3f, 0x88,
	0x5d, 0xce, 0x53, 0x7a, 0xad, 0x48, 0x1b, 0x9c, 0x42, 0x48, 0x65, 0x96, 0xef, 0x23, 0x36, 0x67,
	0xc9, 0x28, 0x75, 0x02, 0x67, 0x30, 0x94, 0x99, 0xd6, 0x6f, 0x42, 0x6d, 0xa3, 0x9e, 0xbd, 0xf8,
	0xd6, 0x78, 0x06, 0x40, 0xef, 0x32, 0x57, 0xa4, 0xd7, 0x39, 0x8f, 0xfa, 0x73, 0x96, 0xb0, 0x74,
	0xe4, 0x2b, 0xf7, 0x3c, 0x3e, 0x87, 0x89, 0x0f, 0xd0, 0x52, 0x70, 0x4d, 0x75, 0x82, 0x11, 0x05,
	0xf1, 0x26, 0xc1, 0x8a, 0x78, 0x05, 0xf8, 0x44, 0xe6, 0xd1, 0x37, 0x6d, 0xa6, 0x69, 0xe7, 0xb2,
	0x6e, 0x6e, 0xdd, 0xf8, 0x96, 0xf6, 0x64, 0xe8, 0xcf, 0xd1, 0xe3, 0x02, 0x26, 0x37, 0x8a, 0xb2,
	0x83, 0x0d, 0x21, 0xe0, 0x59, 0x49, 0xde, 0x65, 0xcf, 0x87, 0x5f, 0x7b, 0xbf, 0x6d, 0xdd, 0xff,
	0xb1, 0xf5, 0x14, 0xc2, 0x6c, 0x5b, 0xe6, 0x3c, 0x0a, 0xe6, 0x2c, 0x19, 0xa6, 0x4e, 0xc4, 0xff,
	0x20, 0xbc, 0x2b, 0xa5, 0xf9, 0xb8, 0xfc, 0x64, 0x10, 0x3c, 0x6b, 0x52, 0xb8, 0x82, 0xd0, 0xae,
	0x8f, 0xb8, 0xb0, 0xec, 0xdb, 0xb0, 0x67, 0x27, 0x9d, 0x9a, 0xe7, 0x73, 0x0d, 0xe3, 0x16, 0x09,
	0x8c, 0x9c, 0xe7, 0x18, 0xce, 0x6c, 0xec, 0x6e, 0x6c, 0x20, 0x5e, 0xc0, 0xc0, 0xd1, 0x40, 0xdf,
	0xb4, 0xc3, 0xe6, 0xc8, 0xeb, 0x90, 0x34, 0xde, 0x0e, 0xa0, 0x8e, 0x77, 0x33, 0xb0, 0x8f, 0xe5,
	0xea, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x06, 0xd6, 0x8b, 0x61, 0x4c, 0x02, 0x00, 0x00,
}
