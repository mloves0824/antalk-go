// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/auth.proto

package auth

import (
	common "github.com/mloves0824/antalk-go/proto/common"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CheckAuthReq struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAuthReq) Reset()         { *m = CheckAuthReq{} }
func (m *CheckAuthReq) String() string { return proto.CompactTextString(m) }
func (*CheckAuthReq) ProtoMessage()    {}
func (*CheckAuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_712ec48c1eaf43a2, []int{0}
}

func (m *CheckAuthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAuthReq.Unmarshal(m, b)
}
func (m *CheckAuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAuthReq.Marshal(b, m, deterministic)
}
func (m *CheckAuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAuthReq.Merge(m, src)
}
func (m *CheckAuthReq) XXX_Size() int {
	return xxx_messageInfo_CheckAuthReq.Size(m)
}
func (m *CheckAuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAuthReq proto.InternalMessageInfo

func (m *CheckAuthReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CheckAuthReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CheckAuthResp struct {
	Result               common.ResultType `protobuf:"varint,1,opt,name=result,proto3,enum=common.ResultType" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckAuthResp) Reset()         { *m = CheckAuthResp{} }
func (m *CheckAuthResp) String() string { return proto.CompactTextString(m) }
func (*CheckAuthResp) ProtoMessage()    {}
func (*CheckAuthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_712ec48c1eaf43a2, []int{1}
}

func (m *CheckAuthResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAuthResp.Unmarshal(m, b)
}
func (m *CheckAuthResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAuthResp.Marshal(b, m, deterministic)
}
func (m *CheckAuthResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAuthResp.Merge(m, src)
}
func (m *CheckAuthResp) XXX_Size() int {
	return xxx_messageInfo_CheckAuthResp.Size(m)
}
func (m *CheckAuthResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAuthResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAuthResp proto.InternalMessageInfo

func (m *CheckAuthResp) GetResult() common.ResultType {
	if m != nil {
		return m.Result
	}
	return common.ResultType_ResultOK
}

func init() {
	proto.RegisterType((*CheckAuthReq)(nil), "auth.CheckAuthReq")
	proto.RegisterType((*CheckAuthResp)(nil), "auth.CheckAuthResp")
}

func init() { proto.RegisterFile("auth/auth.proto", fileDescriptor_712ec48c1eaf43a2) }

var fileDescriptor_712ec48c1eaf43a2 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2c, 0x2d, 0xc9,
	0xd0, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x94, 0x70, 0x72,
	0x7e, 0x6e, 0x6e, 0x7e, 0x9e, 0x3e, 0x84, 0x82, 0x48, 0x29, 0x99, 0x71, 0xf1, 0x38, 0x67, 0xa4,
	0x26, 0x67, 0x3b, 0x96, 0x96, 0x64, 0x04, 0xa5, 0x16, 0x0a, 0x09, 0x70, 0x31, 0x97, 0x66, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x25, 0xf9, 0xd9,
	0xa9, 0x79, 0x12, 0x4c, 0x60, 0x31, 0x08, 0x47, 0xc9, 0x9a, 0x8b, 0x17, 0x49, 0x5f, 0x71, 0x81,
	0x90, 0x16, 0x17, 0x5b, 0x51, 0x6a, 0x71, 0x69, 0x4e, 0x09, 0x58, 0x2f, 0x9f, 0x91, 0x90, 0x1e,
	0xd4, 0x9e, 0x20, 0xb0, 0x68, 0x48, 0x65, 0x41, 0x6a, 0x10, 0x54, 0x85, 0x91, 0x33, 0x17, 0x37,
	0x48, 0x5f, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x09, 0x17, 0x27, 0xdc, 0x2c, 0x21,
	0x21, 0x3d, 0xb0, 0xc3, 0x91, 0x1d, 0x25, 0x25, 0x8c, 0x21, 0x56, 0x5c, 0xe0, 0xa4, 0xc2, 0x25,
	0x9a, 0x99, 0xaf, 0x97, 0x5e, 0x54, 0x90, 0xac, 0x97, 0x5a, 0x91, 0x98, 0x5b, 0x90, 0x93, 0x5a,
	0x0c, 0x56, 0xe6, 0xc4, 0x09, 0x52, 0x12, 0x00, 0xf2, 0x5d, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0x9b,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0xd2, 0x35, 0xe7, 0x14, 0x01, 0x00, 0x00,
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
	CheckAuth(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*CheckAuthResp, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) CheckAuth(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*CheckAuthResp, error) {
	out := new(CheckAuthResp)
	err := c.cc.Invoke(ctx, "/auth.AuthService/CheckAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	CheckAuth(context.Context, *CheckAuthReq) (*CheckAuthResp, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/CheckAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).CheckAuth(ctx, req.(*CheckAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAuth",
			Handler:    _AuthService_CheckAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
