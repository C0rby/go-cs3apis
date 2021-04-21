// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/admin/user/v1beta1/user_api.proto

package userv1beta1

import (
	context "context"
	fmt "fmt"
	v1beta11 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	v1beta12 "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	v1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateUserRequest struct {
	// OPTIONAL.
	// Opaque information. Allow to send any arbitrary data a service might use that is outside the API boundaries
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The information of user to be created.
	User                 *v1beta11.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_127a7770d33b8726, []int{0}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *CreateUserRequest) GetUser() *v1beta11.User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The user information.
	User                 *v1beta11.User `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_127a7770d33b8726, []int{1}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateUserResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *CreateUserResponse) GetUser() *v1beta11.User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteUserRequest struct {
	// OPTIONAL.
	// Opaque information. Allow to send any arbitrary data a service might use that is outside the API boundaries
	Opaque *v1beta1.Opaque `protobuf:"bytes,1,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// REQUIRED.
	// The user to be deleted, given their ID.
	UserId               *v1beta11.UserId `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_127a7770d33b8726, []int{2}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func (m *DeleteUserRequest) GetUserId() *v1beta11.UserId {
	if m != nil {
		return m.UserId
	}
	return nil
}

type DeleteUserResponse struct {
	// REQUIRED.
	// The response status.
	Status *v1beta12.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// OPTIONAL.
	// Opaque information.
	Opaque               *v1beta1.Opaque `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_127a7770d33b8726, []int{3}
}

func (m *DeleteUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserResponse.Unmarshal(m, b)
}
func (m *DeleteUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserResponse.Marshal(b, m, deterministic)
}
func (m *DeleteUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserResponse.Merge(m, src)
}
func (m *DeleteUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteUserResponse.Size(m)
}
func (m *DeleteUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserResponse proto.InternalMessageInfo

func (m *DeleteUserResponse) GetStatus() *v1beta12.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DeleteUserResponse) GetOpaque() *v1beta1.Opaque {
	if m != nil {
		return m.Opaque
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "cs3.admin.user.v1beta1.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "cs3.admin.user.v1beta1.CreateUserResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "cs3.admin.user.v1beta1.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "cs3.admin.user.v1beta1.DeleteUserResponse")
}

func init() {
	proto.RegisterFile("cs3/admin/user/v1beta1/user_api.proto", fileDescriptor_127a7770d33b8726)
}

var fileDescriptor_127a7770d33b8726 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xbf, 0x6a, 0xe3, 0x40,
	0x10, 0xc6, 0x91, 0x7c, 0xc8, 0xb0, 0xbe, 0xc6, 0x5b, 0xf8, 0x7c, 0xe2, 0x8e, 0xbb, 0x33, 0x1c,
	0xc4, 0x29, 0x56, 0xc8, 0xea, 0xd2, 0xd9, 0x4e, 0xe3, 0x2a, 0x46, 0xc1, 0x29, 0x82, 0x21, 0xc8,
	0xab, 0x29, 0x04, 0xb1, 0xb5, 0xde, 0x5d, 0x85, 0x98, 0x74, 0x79, 0x94, 0x94, 0x81, 0xbc, 0x48,
	0xda, 0xbc, 0x50, 0xd8, 0xd1, 0xfa, 0x1f, 0x8e, 0x49, 0x4c, 0x20, 0x95, 0x10, 0xbf, 0x6f, 0xe6,
	0xfb, 0x66, 0x77, 0x96, 0xfc, 0xe7, 0x2a, 0x0a, 0x92, 0x74, 0x9a, 0xcd, 0x82, 0x42, 0x81, 0x0c,
	0x6e, 0xc2, 0x09, 0xe8, 0x24, 0xc4, 0x9f, 0xab, 0x44, 0x64, 0x4c, 0xc8, 0x5c, 0xe7, 0xb4, 0xc1,
	0x55, 0xc4, 0x50, 0xc6, 0x0c, 0x61, 0x56, 0xe6, 0xb7, 0x4d, 0x79, 0x96, 0xc2, 0x4c, 0x67, 0x7a,
	0xb1, 0xdd, 0x41, 0x82, 0xca, 0x0b, 0xc9, 0x41, 0x95, 0x2d, 0xfc, 0x5f, 0x46, 0x2a, 0x05, 0x5f,
	0x09, 0x94, 0x4e, 0x74, 0xb1, 0xa4, 0xbf, 0x0d, 0xd5, 0x0b, 0x01, 0x6a, 0xc5, 0xf1, 0xaf, 0xc4,
	0xad, 0x3b, 0x52, 0xef, 0x4b, 0x48, 0x34, 0x8c, 0x14, 0xc8, 0x18, 0xe6, 0x05, 0x28, 0x4d, 0x43,
	0xe2, 0xe5, 0x22, 0x99, 0x17, 0xd0, 0x74, 0xfe, 0x3a, 0x47, 0xb5, 0xce, 0x4f, 0x66, 0x52, 0x96,
	0x65, 0xb6, 0x09, 0x3b, 0x43, 0x41, 0x6c, 0x85, 0x34, 0x22, 0xdf, 0x4c, 0xc8, 0xa6, 0x8b, 0x05,
	0x7f, 0xb0, 0x60, 0x19, 0x7f, 0x6b, 0x32, 0x86, 0x46, 0x28, 0x6e, 0x3d, 0x39, 0x84, 0x6e, 0xba,
	0x2b, 0x91, 0xcf, 0x14, 0xd0, 0x80, 0x78, 0xe5, 0x08, 0xd6, 0xfe, 0x07, 0x76, 0x93, 0x82, 0xaf,
	0x7a, 0x9c, 0x23, 0x8e, 0xad, 0x6c, 0x23, 0xaf, 0x7b, 0x68, 0xde, 0xca, 0x21, 0x79, 0xef, 0x1d,
	0x52, 0x3f, 0x85, 0x6b, 0xf8, 0xf4, 0x69, 0x9d, 0x90, 0x2a, 0xee, 0x41, 0x96, 0xda, 0xc4, 0xff,
	0xde, 0x09, 0x30, 0x48, 0x63, 0xaf, 0xc0, 0x6f, 0xeb, 0x96, 0xd0, 0xcd, 0x0c, 0x5f, 0x77, 0x66,
	0x9d, 0x17, 0x87, 0x54, 0x8d, 0x69, 0x77, 0x38, 0xa0, 0x9c, 0x90, 0xf5, 0xcd, 0xd1, 0x36, 0x7b,
	0x7b, 0x8d, 0xd9, 0xce, 0x6e, 0xf9, 0xc7, 0x1f, 0x91, 0xda, 0xa1, 0x38, 0x21, 0xeb, 0x51, 0xf7,
	0x9b, 0xec, 0x5c, 0xc9, 0x7e, 0x93, 0xdd, 0x93, 0xeb, 0xcd, 0x89, 0xcf, 0xf3, 0xe9, 0x9e, 0x82,
	0xde, 0x77, 0x1c, 0x58, 0x64, 0x43, 0xf3, 0x5a, 0x86, 0xce, 0x65, 0xcd, 0x50, 0x0b, 0x1f, 0xdc,
	0x4a, 0xbf, 0x3b, 0x7a, 0x74, 0x1b, 0x7d, 0x15, 0xb1, 0x2e, 0xd6, 0x1a, 0x35, 0xbb, 0x08, 0x7b,
	0x06, 0x3f, 0x23, 0x18, 0x23, 0x18, 0x1b, 0x30, 0xb6, 0x60, 0xe2, 0xe1, 0xdb, 0x8b, 0x5e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xc9, 0xbf, 0x7d, 0x37, 0x24, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserAPIClient is the client API for UserAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserAPIClient interface {
	// Create a user account.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// Delete a user account.
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type userAPIClient struct {
	cc *grpc.ClientConn
}

func NewUserAPIClient(cc *grpc.ClientConn) UserAPIClient {
	return &userAPIClient{cc}
}

func (c *userAPIClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/cs3.admin.user.v1beta1.UserAPI/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAPIClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/cs3.admin.user.v1beta1.UserAPI/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAPIServer is the server API for UserAPI service.
type UserAPIServer interface {
	// Create a user account.
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// Delete a user account.
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
}

// UnimplementedUserAPIServer can be embedded to have forward compatible implementations.
type UnimplementedUserAPIServer struct {
}

func (*UnimplementedUserAPIServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserAPIServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterUserAPIServer(s *grpc.Server, srv UserAPIServer) {
	s.RegisterService(&_UserAPI_serviceDesc, srv)
}

func _UserAPI_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.admin.user.v1beta1.UserAPI/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAPI_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAPIServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.admin.user.v1beta1.UserAPI/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAPIServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.admin.user.v1beta1.UserAPI",
	HandlerType: (*UserAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserAPI_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserAPI_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/admin/user/v1beta1/user_api.proto",
}
