// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/auth/v0alpha/auth.proto

package authv0alphapb

import (
	context "context"
	fmt "fmt"
	rpc "github.com/cernbox/go-cs3apis/cs3/rpc"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GenerateAccessTokenRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateAccessTokenRequest) Reset()         { *m = GenerateAccessTokenRequest{} }
func (m *GenerateAccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateAccessTokenRequest) ProtoMessage()    {}
func (*GenerateAccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ae21de9dc3d6621, []int{0}
}

func (m *GenerateAccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateAccessTokenRequest.Unmarshal(m, b)
}
func (m *GenerateAccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateAccessTokenRequest.Marshal(b, m, deterministic)
}
func (m *GenerateAccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateAccessTokenRequest.Merge(m, src)
}
func (m *GenerateAccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateAccessTokenRequest.Size(m)
}
func (m *GenerateAccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateAccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateAccessTokenRequest proto.InternalMessageInfo

func (m *GenerateAccessTokenRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GenerateAccessTokenRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GenerateAccessTokenResponse struct {
	Status               *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	AccessToken          string      `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GenerateAccessTokenResponse) Reset()         { *m = GenerateAccessTokenResponse{} }
func (m *GenerateAccessTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateAccessTokenResponse) ProtoMessage()    {}
func (*GenerateAccessTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ae21de9dc3d6621, []int{1}
}

func (m *GenerateAccessTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateAccessTokenResponse.Unmarshal(m, b)
}
func (m *GenerateAccessTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateAccessTokenResponse.Marshal(b, m, deterministic)
}
func (m *GenerateAccessTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateAccessTokenResponse.Merge(m, src)
}
func (m *GenerateAccessTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateAccessTokenResponse.Size(m)
}
func (m *GenerateAccessTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateAccessTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateAccessTokenResponse proto.InternalMessageInfo

func (m *GenerateAccessTokenResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GenerateAccessTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type WhoAmIRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WhoAmIRequest) Reset()         { *m = WhoAmIRequest{} }
func (m *WhoAmIRequest) String() string { return proto.CompactTextString(m) }
func (*WhoAmIRequest) ProtoMessage()    {}
func (*WhoAmIRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ae21de9dc3d6621, []int{2}
}

func (m *WhoAmIRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoAmIRequest.Unmarshal(m, b)
}
func (m *WhoAmIRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoAmIRequest.Marshal(b, m, deterministic)
}
func (m *WhoAmIRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoAmIRequest.Merge(m, src)
}
func (m *WhoAmIRequest) XXX_Size() int {
	return xxx_messageInfo_WhoAmIRequest.Size(m)
}
func (m *WhoAmIRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoAmIRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WhoAmIRequest proto.InternalMessageInfo

func (m *WhoAmIRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type WhoAmIResponse struct {
	Status               *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	User                 *User       `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WhoAmIResponse) Reset()         { *m = WhoAmIResponse{} }
func (m *WhoAmIResponse) String() string { return proto.CompactTextString(m) }
func (*WhoAmIResponse) ProtoMessage()    {}
func (*WhoAmIResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ae21de9dc3d6621, []int{3}
}

func (m *WhoAmIResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoAmIResponse.Unmarshal(m, b)
}
func (m *WhoAmIResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoAmIResponse.Marshal(b, m, deterministic)
}
func (m *WhoAmIResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoAmIResponse.Merge(m, src)
}
func (m *WhoAmIResponse) XXX_Size() int {
	return xxx_messageInfo_WhoAmIResponse.Size(m)
}
func (m *WhoAmIResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoAmIResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WhoAmIResponse proto.InternalMessageInfo

func (m *WhoAmIResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *WhoAmIResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*GenerateAccessTokenRequest)(nil), "cs3.authv0alpha.GenerateAccessTokenRequest")
	proto.RegisterType((*GenerateAccessTokenResponse)(nil), "cs3.authv0alpha.GenerateAccessTokenResponse")
	proto.RegisterType((*WhoAmIRequest)(nil), "cs3.authv0alpha.WhoAmIRequest")
	proto.RegisterType((*WhoAmIResponse)(nil), "cs3.authv0alpha.WhoAmIResponse")
}

func init() { proto.RegisterFile("cs3/auth/v0alpha/auth.proto", fileDescriptor_1ae21de9dc3d6621) }

var fileDescriptor_1ae21de9dc3d6621 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xdd, 0x4e, 0xea, 0x40,
	0x10, 0x4e, 0xc9, 0x49, 0x73, 0xd8, 0x1e, 0x0e, 0x49, 0xd1, 0x84, 0x94, 0x44, 0x91, 0x1b, 0x35,
	0x9a, 0x85, 0xb4, 0x4f, 0xd0, 0x72, 0x61, 0xb8, 0x92, 0x14, 0xfc, 0x89, 0x31, 0x9a, 0x65, 0x99,
	0xa4, 0x44, 0xe9, 0xae, 0xbb, 0x5b, 0x7c, 0x1f, 0x2f, 0x7d, 0x14, 0x9e, 0xca, 0xec, 0xb6, 0x28,
	0x52, 0x48, 0xf4, 0x72, 0xe6, 0xfb, 0xd9, 0xf9, 0x66, 0x07, 0xb5, 0xa8, 0x0c, 0xba, 0x24, 0x53,
	0x49, 0x77, 0xd1, 0x23, 0xcf, 0x3c, 0x21, 0xa6, 0xc0, 0x5c, 0x30, 0xc5, 0xdc, 0x3a, 0x95, 0x01,
	0xd6, 0x75, 0x81, 0x79, 0xed, 0x12, 0x5b, 0x80, 0x64, 0x99, 0xa0, 0x20, 0x73, 0x89, 0xb7, 0xa7,
	0x19, 0x82, 0xd3, 0xae, 0x54, 0x44, 0x65, 0x45, 0xb7, 0x33, 0x46, 0xde, 0x05, 0xa4, 0x20, 0x88,
	0x82, 0x90, 0x52, 0x90, 0x72, 0xcc, 0x9e, 0x20, 0x8d, 0xe1, 0x25, 0x03, 0xa9, 0x5c, 0x0f, 0xfd,
	0xcd, 0x24, 0x88, 0x94, 0xcc, 0xa1, 0x69, 0xb5, 0xad, 0x93, 0x6a, 0xfc, 0x59, 0x6b, 0x8c, 0x13,
	0x29, 0x5f, 0x99, 0x98, 0x36, 0x2b, 0x39, 0xb6, 0xaa, 0x3b, 0x33, 0xd4, 0xda, 0xea, 0x2a, 0x39,
	0x4b, 0x25, 0xb8, 0xc7, 0xc8, 0xce, 0x87, 0x30, 0xa6, 0x8e, 0x5f, 0xc7, 0x3a, 0x8e, 0xe0, 0x14,
	0x8f, 0x4c, 0x3b, 0x2e, 0x60, 0xf7, 0x08, 0xfd, 0x23, 0x46, 0xff, 0xa8, 0xb4, 0x41, 0xf1, 0x8e,
	0x43, 0xbe, 0x3c, 0x3b, 0x3e, 0xaa, 0xdd, 0x24, 0x2c, 0x9c, 0x0f, 0x56, 0x33, 0x6f, 0x6a, 0xac,
	0xb2, 0x66, 0x8a, 0xfe, 0xaf, 0x34, 0xbf, 0x9d, 0xe8, 0x14, 0xfd, 0xd1, 0x1b, 0x30, 0x93, 0x38,
	0xfe, 0x3e, 0xde, 0xf8, 0x07, 0x7c, 0x25, 0x41, 0xc4, 0x86, 0xe2, 0x2f, 0x2d, 0xe4, 0x84, 0x99,
	0x4a, 0x46, 0x20, 0x16, 0x33, 0x0a, 0x2e, 0x47, 0x8d, 0x2d, 0x4b, 0x71, 0xcf, 0x4a, 0x1e, 0xbb,
	0x3f, 0xc4, 0x3b, 0xff, 0x19, 0xb9, 0x48, 0x35, 0x40, 0x76, 0x9e, 0xd3, 0x3d, 0x28, 0xe9, 0xbe,
	0x2d, 0xcd, 0x3b, 0xdc, 0x89, 0xe7, 0x56, 0xd1, 0x03, 0x6a, 0x50, 0x36, 0xdf, 0x64, 0x45, 0x55,
	0x1d, 0x70, 0xa8, 0x2f, 0x69, 0x68, 0xdd, 0xd5, 0xd6, 0x10, 0x3e, 0x79, 0xab, 0xd8, 0xfd, 0xe8,
	0xf2, 0x36, 0x8c, 0xde, 0x2b, 0xf5, 0xfe, 0x28, 0xc0, 0x9a, 0x79, 0xdd, 0x0b, 0x35, 0xb8, 0x34,
	0x9d, 0xfb, 0xb5, 0xce, 0xc4, 0x36, 0xe7, 0x18, 0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x16,
	0xf7, 0x58, 0xf6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error)
	WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) GenerateAccessToken(ctx context.Context, in *GenerateAccessTokenRequest, opts ...grpc.CallOption) (*GenerateAccessTokenResponse, error) {
	out := new(GenerateAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/cs3.authv0alpha.AuthService/GenerateAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) WhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error) {
	out := new(WhoAmIResponse)
	err := c.cc.Invoke(ctx, "/cs3.authv0alpha.AuthService/WhoAmI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	GenerateAccessToken(context.Context, *GenerateAccessTokenRequest) (*GenerateAccessTokenResponse, error)
	WhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResponse, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_GenerateAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GenerateAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.authv0alpha.AuthService/GenerateAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GenerateAccessToken(ctx, req.(*GenerateAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_WhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).WhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.authv0alpha.AuthService/WhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).WhoAmI(ctx, req.(*WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.authv0alpha.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateAccessToken",
			Handler:    _AuthService_GenerateAccessToken_Handler,
		},
		{
			MethodName: "WhoAmI",
			Handler:    _AuthService_WhoAmI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/auth/v0alpha/auth.proto",
}
