// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

/*
Package users is a generated protocol buffer package.

It is generated from these files:
	users.proto

It has these top-level messages:
	GetRequest
	User
*/
package users

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

type GetRequest struct {
	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids" json:"ids,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetRequest) GetIds() []uint64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type User struct {
	FirstName string `protobuf:"bytes,1,opt,name=firstName" json:"firstName,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "users.GetRequest")
	proto.RegisterType((*User)(nil), "users.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Users service

type UsersClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error)
	Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/users.Users/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/users.Users/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersServer interface {
	Get(context.Context, *GetRequest) (*User, error)
	Update(context.Context, *User) (*User, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.Users/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Update(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Users_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Users_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}

func init() { proto.RegisterFile("users.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2d, 0x4e, 0x2d,
	0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0xe4, 0xb8, 0xb8, 0xdc,
	0x53, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0x33, 0x53, 0x8a,
	0x25, 0x18, 0x15, 0x98, 0x35, 0x58, 0x82, 0x40, 0x4c, 0x25, 0x15, 0x2e, 0x96, 0xd0, 0xe2, 0xd4,
	0x22, 0x21, 0x19, 0x2e, 0xce, 0xb4, 0xcc, 0xa2, 0xe2, 0x12, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x84, 0x80, 0x51, 0x18, 0x17, 0x2b, 0x48, 0x55, 0xb1, 0x90, 0x3a,
	0x17, 0xb3, 0x7b, 0x6a, 0x89, 0x90, 0xa0, 0x1e, 0xc4, 0x2a, 0x84, 0xd1, 0x52, 0xdc, 0x50, 0x21,
	0x90, 0x3a, 0x25, 0x06, 0x21, 0x15, 0x2e, 0xb6, 0xd0, 0x82, 0x94, 0xc4, 0x92, 0x54, 0x21, 0x64,
	0x09, 0x34, 0x55, 0x4e, 0xda, 0x5c, 0x0a, 0xc9, 0xf9, 0xb9, 0x7a, 0xe9, 0x99, 0x25, 0x19, 0xa5,
	0x49, 0x7a, 0xe5, 0x19, 0x99, 0x39, 0xa9, 0x39, 0xf9, 0xf9, 0x05, 0x7a, 0x39, 0xa9, 0x65, 0x99,
	0x25, 0x10, 0x8f, 0x38, 0x41, 0x6c, 0x0e, 0x60, 0x4c, 0x62, 0x03, 0xf3, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xfc, 0x0f, 0xc7, 0x2c, 0xe7, 0x00, 0x00, 0x00,
}