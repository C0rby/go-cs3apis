// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/appprovider/v0alpha/appprovider.proto

package appproviderv0alphapb

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

type GetIFrameRequest struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Miemtype             string   `protobuf:"bytes,2,opt,name=miemtype,proto3" json:"miemtype,omitempty"`
	AccessToken          string   `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIFrameRequest) Reset()         { *m = GetIFrameRequest{} }
func (m *GetIFrameRequest) String() string { return proto.CompactTextString(m) }
func (*GetIFrameRequest) ProtoMessage()    {}
func (*GetIFrameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f74a8f301591c1e7, []int{0}
}

func (m *GetIFrameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIFrameRequest.Unmarshal(m, b)
}
func (m *GetIFrameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIFrameRequest.Marshal(b, m, deterministic)
}
func (m *GetIFrameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIFrameRequest.Merge(m, src)
}
func (m *GetIFrameRequest) XXX_Size() int {
	return xxx_messageInfo_GetIFrameRequest.Size(m)
}
func (m *GetIFrameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIFrameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIFrameRequest proto.InternalMessageInfo

func (m *GetIFrameRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *GetIFrameRequest) GetMiemtype() string {
	if m != nil {
		return m.Miemtype
	}
	return ""
}

func (m *GetIFrameRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type GetIFrameResponse struct {
	Status               *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IframeLocation       string      `protobuf:"bytes,2,opt,name=iframe_location,json=iframeLocation,proto3" json:"iframe_location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetIFrameResponse) Reset()         { *m = GetIFrameResponse{} }
func (m *GetIFrameResponse) String() string { return proto.CompactTextString(m) }
func (*GetIFrameResponse) ProtoMessage()    {}
func (*GetIFrameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f74a8f301591c1e7, []int{1}
}

func (m *GetIFrameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIFrameResponse.Unmarshal(m, b)
}
func (m *GetIFrameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIFrameResponse.Marshal(b, m, deterministic)
}
func (m *GetIFrameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIFrameResponse.Merge(m, src)
}
func (m *GetIFrameResponse) XXX_Size() int {
	return xxx_messageInfo_GetIFrameResponse.Size(m)
}
func (m *GetIFrameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIFrameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetIFrameResponse proto.InternalMessageInfo

func (m *GetIFrameResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetIFrameResponse) GetIframeLocation() string {
	if m != nil {
		return m.IframeLocation
	}
	return ""
}

func init() {
	proto.RegisterType((*GetIFrameRequest)(nil), "cs3.appproviderv0alpha.GetIFrameRequest")
	proto.RegisterType((*GetIFrameResponse)(nil), "cs3.appproviderv0alpha.GetIFrameResponse")
}

func init() {
	proto.RegisterFile("cs3/appprovider/v0alpha/appprovider.proto", fileDescriptor_f74a8f301591c1e7)
}

var fileDescriptor_f74a8f301591c1e7 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xd1, 0x4a, 0x3a, 0x41,
	0x14, 0xc6, 0x59, 0xff, 0x20, 0x7f, 0xc7, 0x48, 0x1b, 0x44, 0x64, 0xaf, 0xca, 0x1b, 0xf5, 0x66,
	0x14, 0xf7, 0x09, 0x5c, 0xa1, 0x08, 0x82, 0x44, 0x23, 0x22, 0x02, 0x1b, 0xa7, 0x23, 0x2d, 0x39,
	0x3b, 0xa7, 0x99, 0x71, 0xa1, 0xcb, 0x5e, 0xa5, 0xcb, 0x1e, 0xa5, 0xa7, 0x8a, 0xd9, 0xd9, 0x64,
	0xc9, 0x82, 0x2e, 0xcf, 0xef, 0xfb, 0xce, 0x9c, 0x73, 0xbe, 0x21, 0x03, 0x61, 0xa2, 0x21, 0x47,
	0x44, 0xad, 0xb2, 0xe4, 0x01, 0xf4, 0x30, 0x1b, 0xf1, 0x0d, 0x3e, 0xf2, 0x32, 0x63, 0xa8, 0x95,
	0x55, 0xb4, 0x2d, 0x4c, 0xc4, 0x4a, 0xb8, 0x70, 0x86, 0x2d, 0xf7, 0x84, 0x46, 0x31, 0x34, 0x96,
	0xdb, 0xad, 0xf1, 0xee, 0xae, 0x24, 0xcd, 0x33, 0xb0, 0xe7, 0xa7, 0x9a, 0x4b, 0x98, 0xc3, 0xf3,
	0x16, 0x8c, 0xa5, 0x21, 0xf9, 0xbf, 0x4e, 0x36, 0x90, 0x72, 0x09, 0x9d, 0xe0, 0x38, 0xe8, 0xd7,
	0xe6, 0xbb, 0xda, 0x69, 0x32, 0x01, 0x69, 0x5f, 0x10, 0x3a, 0x15, 0xaf, 0x7d, 0xd5, 0xf4, 0x84,
	0x1c, 0x70, 0x21, 0xc0, 0x98, 0xa5, 0x55, 0x4f, 0x90, 0x76, 0xfe, 0xe5, 0x7a, 0xdd, 0xb3, 0x2b,
	0x87, 0xba, 0x40, 0x8e, 0x4a, 0xe3, 0x0c, 0xaa, 0xd4, 0x00, 0xed, 0x91, 0xaa, 0xdf, 0x29, 0x9f,
	0x56, 0x1f, 0x37, 0x98, 0x3b, 0x41, 0xa3, 0x60, 0x8b, 0x1c, 0xcf, 0x0b, 0x99, 0xf6, 0x48, 0x23,
	0x59, 0xbb, 0xd6, 0xe5, 0x46, 0x09, 0x6e, 0x13, 0x95, 0x16, 0x3b, 0x1c, 0x7a, 0x7c, 0x51, 0xd0,
	0x71, 0x46, 0xe8, 0x04, 0x71, 0x56, 0x24, 0xb0, 0x00, 0x9d, 0x25, 0x02, 0xe8, 0x3d, 0xa9, 0xed,
	0x86, 0xd3, 0x3e, 0xfb, 0x39, 0x27, 0xf6, 0x3d, 0x8e, 0x70, 0xf0, 0x07, 0xa7, 0xbf, 0x24, 0x7e,
	0x0d, 0x48, 0x28, 0x94, 0xfc, 0xa5, 0x21, 0x6e, 0x96, 0x96, 0x9a, 0xb9, 0xf8, 0x67, 0xc1, 0x6d,
	0x6b, 0xdf, 0x87, 0xab, 0xb7, 0x4a, 0x75, 0x1a, 0x5f, 0xde, 0x4c, 0xe2, 0xf7, 0x4a, 0x7b, 0xba,
	0x88, 0x58, 0xa9, 0xef, 0x7a, 0x34, 0x71, 0x9e, 0x8f, 0x5c, 0xb8, 0xdb, 0x17, 0x56, 0xd5, 0xfc,
	0x63, 0xa3, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x53, 0xfd, 0xa1, 0x33, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppProviderServiceClient is the client API for AppProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppProviderServiceClient interface {
	GetIFrame(ctx context.Context, in *GetIFrameRequest, opts ...grpc.CallOption) (*GetIFrameResponse, error)
}

type appProviderServiceClient struct {
	cc *grpc.ClientConn
}

func NewAppProviderServiceClient(cc *grpc.ClientConn) AppProviderServiceClient {
	return &appProviderServiceClient{cc}
}

func (c *appProviderServiceClient) GetIFrame(ctx context.Context, in *GetIFrameRequest, opts ...grpc.CallOption) (*GetIFrameResponse, error) {
	out := new(GetIFrameResponse)
	err := c.cc.Invoke(ctx, "/cs3.appproviderv0alpha.AppProviderService/GetIFrame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppProviderServiceServer is the server API for AppProviderService service.
type AppProviderServiceServer interface {
	GetIFrame(context.Context, *GetIFrameRequest) (*GetIFrameResponse, error)
}

func RegisterAppProviderServiceServer(s *grpc.Server, srv AppProviderServiceServer) {
	s.RegisterService(&_AppProviderService_serviceDesc, srv)
}

func _AppProviderService_GetIFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIFrameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppProviderServiceServer).GetIFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cs3.appproviderv0alpha.AppProviderService/GetIFrame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppProviderServiceServer).GetIFrame(ctx, req.(*GetIFrameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppProviderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cs3.appproviderv0alpha.AppProviderService",
	HandlerType: (*AppProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIFrame",
			Handler:    _AppProviderService_GetIFrame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cs3/appprovider/v0alpha/appprovider.proto",
}
