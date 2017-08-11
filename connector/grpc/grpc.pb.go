// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	grpc.proto

It has these top-level messages:
	Scopes
	Identity
	LoginRequest
	LoginResponse
	RefreshRequest
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type Scopes struct {
	OfflineAccess bool `protobuf:"varint,1,opt,name=offline_access,json=offlineAccess" json:"offline_access,omitempty"`
	Groups        bool `protobuf:"varint,2,opt,name=groups" json:"groups,omitempty"`
}

func (m *Scopes) Reset()                    { *m = Scopes{} }
func (m *Scopes) String() string            { return proto.CompactTextString(m) }
func (*Scopes) ProtoMessage()               {}
func (*Scopes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Scopes) GetOfflineAccess() bool {
	if m != nil {
		return m.OfflineAccess
	}
	return false
}

func (m *Scopes) GetGroups() bool {
	if m != nil {
		return m.Groups
	}
	return false
}

type Identity struct {
	UserId        string   `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Username      string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email         string   `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	EmailVerified bool     `protobuf:"varint,4,opt,name=email_verified,json=emailVerified" json:"email_verified,omitempty"`
	Groups        []string `protobuf:"bytes,5,rep,name=groups" json:"groups,omitempty"`
}

func (m *Identity) Reset()                    { *m = Identity{} }
func (m *Identity) String() string            { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()               {}
func (*Identity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Identity) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Identity) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Identity) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Identity) GetEmailVerified() bool {
	if m != nil {
		return m.EmailVerified
	}
	return false
}

func (m *Identity) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

type LoginRequest struct {
	Username string  `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string  `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Scopes   *Scopes `protobuf:"bytes,3,opt,name=scopes" json:"scopes,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetScopes() *Scopes {
	if m != nil {
		return m.Scopes
	}
	return nil
}

type LoginResponse struct {
	Identity      *Identity `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	ValidPassword bool      `protobuf:"varint,2,opt,name=valid_password,json=validPassword" json:"valid_password,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LoginResponse) GetIdentity() *Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *LoginResponse) GetValidPassword() bool {
	if m != nil {
		return m.ValidPassword
	}
	return false
}

type RefreshRequest struct {
	Scopes   *Scopes   `protobuf:"bytes,1,opt,name=scopes" json:"scopes,omitempty"`
	Identity *Identity `protobuf:"bytes,2,opt,name=identity" json:"identity,omitempty"`
}

func (m *RefreshRequest) Reset()                    { *m = RefreshRequest{} }
func (m *RefreshRequest) String() string            { return proto.CompactTextString(m) }
func (*RefreshRequest) ProtoMessage()               {}
func (*RefreshRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RefreshRequest) GetScopes() *Scopes {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *RefreshRequest) GetIdentity() *Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func init() {
	proto.RegisterType((*Scopes)(nil), "grpc.Scopes")
	proto.RegisterType((*Identity)(nil), "grpc.Identity")
	proto.RegisterType((*LoginRequest)(nil), "grpc.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "grpc.LoginResponse")
	proto.RegisterType((*RefreshRequest)(nil), "grpc.RefreshRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for Connector service

type ConnectorClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc1.CallOption) (*LoginResponse, error)
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc1.CallOption) (*Identity, error)
}

type connectorClient struct {
	cc *grpc1.ClientConn
}

func NewConnectorClient(cc *grpc1.ClientConn) ConnectorClient {
	return &connectorClient{cc}
}

func (c *connectorClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc1.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc1.Invoke(ctx, "/grpc.Connector/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc1.CallOption) (*Identity, error) {
	out := new(Identity)
	err := grpc1.Invoke(ctx, "/grpc.Connector/Refresh", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Connector service

type ConnectorServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Refresh(context.Context, *RefreshRequest) (*Identity, error)
}

func RegisterConnectorServer(s *grpc1.Server, srv ConnectorServer) {
	s.RegisterService(&_Connector_serviceDesc, srv)
}

func _Connector_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Login(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Connector/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Connector_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServer).Refresh(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Connector/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Connector_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.Connector",
	HandlerType: (*ConnectorServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Connector_Login_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Connector_Refresh_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "grpc.proto",
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6a, 0xe3, 0x30,
	0x10, 0xc6, 0xe3, 0xfc, 0x71, 0x9c, 0xc9, 0x9f, 0x83, 0x36, 0xec, 0x1a, 0x9f, 0x82, 0xd9, 0x85,
	0xa5, 0x87, 0x40, 0xdd, 0x27, 0x28, 0x3d, 0x94, 0x40, 0x0f, 0x45, 0x85, 0x5e, 0x8d, 0x63, 0x8f,
	0x53, 0x81, 0x63, 0xb9, 0x92, 0x93, 0xd2, 0xf7, 0xe8, 0x03, 0x17, 0x8f, 0xe4, 0x34, 0x0e, 0xb4,
	0xb7, 0xf9, 0xbe, 0xd1, 0x78, 0x7e, 0xf3, 0x61, 0x80, 0x9d, 0xaa, 0xd2, 0x75, 0xa5, 0x64, 0x2d,
	0xd9, 0xb0, 0xa9, 0xc3, 0x7b, 0x70, 0x9f, 0x52, 0x59, 0xa1, 0x66, 0xff, 0x60, 0x21, 0xf3, 0xbc,
	0x10, 0x25, 0xc6, 0x49, 0x9a, 0xa2, 0xd6, 0xbe, 0xb3, 0x72, 0xfe, 0x7b, 0x7c, 0x6e, 0xdd, 0x5b,
	0x32, 0xd9, 0x6f, 0x70, 0x77, 0x4a, 0x1e, 0x2a, 0xed, 0xf7, 0xa9, 0x6d, 0x55, 0xf8, 0xe1, 0x80,
	0xb7, 0xc9, 0xb0, 0xac, 0x45, 0xfd, 0xce, 0xfe, 0xc0, 0xf8, 0xa0, 0x51, 0xc5, 0x22, 0xa3, 0x8f,
	0x4c, 0xb8, 0xdb, 0xc8, 0x4d, 0xc6, 0x02, 0xf0, 0x9a, 0xaa, 0x4c, 0xf6, 0x48, 0xf3, 0x13, 0x7e,
	0xd2, 0x6c, 0x09, 0x23, 0xdc, 0x27, 0xa2, 0xf0, 0x07, 0xd4, 0x30, 0xa2, 0xc1, 0xa2, 0x22, 0x3e,
	0xa2, 0x12, 0xb9, 0xc0, 0xcc, 0x1f, 0x1a, 0x2c, 0x72, 0x9f, 0xad, 0x79, 0x86, 0x35, 0x5a, 0x0d,
	0x9a, 0x85, 0x16, 0xab, 0x80, 0xd9, 0x83, 0xdc, 0x89, 0x92, 0xe3, 0xeb, 0x01, 0x75, 0xdd, 0x01,
	0x70, 0x2e, 0x00, 0x02, 0xf0, 0xaa, 0x44, 0xeb, 0x37, 0xa9, 0xb2, 0x16, 0xae, 0xd5, 0xec, 0x2f,
	0xb8, 0x9a, 0x72, 0x22, 0xba, 0x69, 0x34, 0x5b, 0x53, 0x94, 0x26, 0x3b, 0x6e, 0x7b, 0xe1, 0x16,
	0xe6, 0x76, 0x9b, 0xae, 0x64, 0xa9, 0x91, 0x5d, 0x81, 0x27, 0x6c, 0x28, 0xb4, 0x6e, 0x1a, 0x2d,
	0xcc, 0x60, 0x1b, 0x15, 0x3f, 0xf5, 0x9b, 0x4b, 0x8f, 0x49, 0x21, 0xb2, 0xb8, 0x03, 0xe1, 0xf1,
	0x39, 0xb9, 0x8f, 0xd6, 0x0c, 0xb7, 0xb0, 0xe0, 0x98, 0x2b, 0xd4, 0x2f, 0xed, 0x4d, 0x5f, 0x6c,
	0xce, 0xf7, 0x6c, 0x1d, 0x94, 0xfe, 0xcf, 0x28, 0x91, 0x82, 0xc9, 0x9d, 0x2c, 0x4b, 0x4c, 0x6b,
	0xa9, 0x58, 0x04, 0x23, 0x3a, 0x8a, 0x31, 0xf3, 0xfe, 0x3c, 0xcf, 0xe0, 0x57, 0xc7, 0x33, 0x57,
	0x87, 0x3d, 0x76, 0x0d, 0x63, 0x0b, 0xc9, 0x96, 0xe6, 0x45, 0x97, 0x39, 0xb8, 0xd8, 0x1d, 0xf6,
	0xb6, 0x2e, 0xfd, 0x96, 0x37, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x5f, 0x6c, 0xdb, 0xa4,
	0x02, 0x00, 0x00,
}
