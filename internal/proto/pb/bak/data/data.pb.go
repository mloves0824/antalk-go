// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data/data.proto

package data

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
	return fileDescriptor_48e8d492558f417f, []int{0}
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
	Uid                  string            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Result               common.ResultType `protobuf:"varint,2,opt,name=result,proto3,enum=common.ResultType" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckAuthResp) Reset()         { *m = CheckAuthResp{} }
func (m *CheckAuthResp) String() string { return proto.CompactTextString(m) }
func (*CheckAuthResp) ProtoMessage()    {}
func (*CheckAuthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_48e8d492558f417f, []int{1}
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

func (m *CheckAuthResp) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *CheckAuthResp) GetResult() common.ResultType {
	if m != nil {
		return m.Result
	}
	return common.ResultType_ResultOK
}

type SetSessionReq struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ServerInfo           string   `protobuf:"bytes,2,opt,name=server_info,json=serverInfo,proto3" json:"server_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetSessionReq) Reset()         { *m = SetSessionReq{} }
func (m *SetSessionReq) String() string { return proto.CompactTextString(m) }
func (*SetSessionReq) ProtoMessage()    {}
func (*SetSessionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48e8d492558f417f, []int{2}
}

func (m *SetSessionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetSessionReq.Unmarshal(m, b)
}
func (m *SetSessionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetSessionReq.Marshal(b, m, deterministic)
}
func (m *SetSessionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSessionReq.Merge(m, src)
}
func (m *SetSessionReq) XXX_Size() int {
	return xxx_messageInfo_SetSessionReq.Size(m)
}
func (m *SetSessionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSessionReq.DiscardUnknown(m)
}

var xxx_messageInfo_SetSessionReq proto.InternalMessageInfo

func (m *SetSessionReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SetSessionReq) GetServerInfo() string {
	if m != nil {
		return m.ServerInfo
	}
	return ""
}

type SetSessionResp struct {
	Uid                  string            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Result               common.ResultType `protobuf:"varint,2,opt,name=result,proto3,enum=common.ResultType" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SetSessionResp) Reset()         { *m = SetSessionResp{} }
func (m *SetSessionResp) String() string { return proto.CompactTextString(m) }
func (*SetSessionResp) ProtoMessage()    {}
func (*SetSessionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_48e8d492558f417f, []int{3}
}

func (m *SetSessionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetSessionResp.Unmarshal(m, b)
}
func (m *SetSessionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetSessionResp.Marshal(b, m, deterministic)
}
func (m *SetSessionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetSessionResp.Merge(m, src)
}
func (m *SetSessionResp) XXX_Size() int {
	return xxx_messageInfo_SetSessionResp.Size(m)
}
func (m *SetSessionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SetSessionResp.DiscardUnknown(m)
}

var xxx_messageInfo_SetSessionResp proto.InternalMessageInfo

func (m *SetSessionResp) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SetSessionResp) GetResult() common.ResultType {
	if m != nil {
		return m.Result
	}
	return common.ResultType_ResultOK
}

type GetSessionReq struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionReq) Reset()         { *m = GetSessionReq{} }
func (m *GetSessionReq) String() string { return proto.CompactTextString(m) }
func (*GetSessionReq) ProtoMessage()    {}
func (*GetSessionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48e8d492558f417f, []int{4}
}

func (m *GetSessionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionReq.Unmarshal(m, b)
}
func (m *GetSessionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionReq.Marshal(b, m, deterministic)
}
func (m *GetSessionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionReq.Merge(m, src)
}
func (m *GetSessionReq) XXX_Size() int {
	return xxx_messageInfo_GetSessionReq.Size(m)
}
func (m *GetSessionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionReq proto.InternalMessageInfo

func (m *GetSessionReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type GetSessionResp struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ServerInfo           string   `protobuf:"bytes,2,opt,name=server_info,json=serverInfo,proto3" json:"server_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSessionResp) Reset()         { *m = GetSessionResp{} }
func (m *GetSessionResp) String() string { return proto.CompactTextString(m) }
func (*GetSessionResp) ProtoMessage()    {}
func (*GetSessionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_48e8d492558f417f, []int{5}
}

func (m *GetSessionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSessionResp.Unmarshal(m, b)
}
func (m *GetSessionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSessionResp.Marshal(b, m, deterministic)
}
func (m *GetSessionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSessionResp.Merge(m, src)
}
func (m *GetSessionResp) XXX_Size() int {
	return xxx_messageInfo_GetSessionResp.Size(m)
}
func (m *GetSessionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSessionResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetSessionResp proto.InternalMessageInfo

func (m *GetSessionResp) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *GetSessionResp) GetServerInfo() string {
	if m != nil {
		return m.ServerInfo
	}
	return ""
}

type SaveMsgReq struct {
	MsgInfo              *common.MsgInfo `protobuf:"bytes,1,opt,name=msg_info,json=msgInfo,proto3" json:"msg_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SaveMsgReq) Reset()         { *m = SaveMsgReq{} }
func (m *SaveMsgReq) String() string { return proto.CompactTextString(m) }
func (*SaveMsgReq) ProtoMessage()    {}
func (*SaveMsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_48e8d492558f417f, []int{6}
}

func (m *SaveMsgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveMsgReq.Unmarshal(m, b)
}
func (m *SaveMsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveMsgReq.Marshal(b, m, deterministic)
}
func (m *SaveMsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveMsgReq.Merge(m, src)
}
func (m *SaveMsgReq) XXX_Size() int {
	return xxx_messageInfo_SaveMsgReq.Size(m)
}
func (m *SaveMsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveMsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_SaveMsgReq proto.InternalMessageInfo

func (m *SaveMsgReq) GetMsgInfo() *common.MsgInfo {
	if m != nil {
		return m.MsgInfo
	}
	return nil
}

type SaveMsgResp struct {
	Result               common.ResultType `protobuf:"varint,1,opt,name=result,proto3,enum=common.ResultType" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SaveMsgResp) Reset()         { *m = SaveMsgResp{} }
func (m *SaveMsgResp) String() string { return proto.CompactTextString(m) }
func (*SaveMsgResp) ProtoMessage()    {}
func (*SaveMsgResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_48e8d492558f417f, []int{7}
}

func (m *SaveMsgResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveMsgResp.Unmarshal(m, b)
}
func (m *SaveMsgResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveMsgResp.Marshal(b, m, deterministic)
}
func (m *SaveMsgResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveMsgResp.Merge(m, src)
}
func (m *SaveMsgResp) XXX_Size() int {
	return xxx_messageInfo_SaveMsgResp.Size(m)
}
func (m *SaveMsgResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveMsgResp.DiscardUnknown(m)
}

var xxx_messageInfo_SaveMsgResp proto.InternalMessageInfo

func (m *SaveMsgResp) GetResult() common.ResultType {
	if m != nil {
		return m.Result
	}
	return common.ResultType_ResultOK
}

func init() {
	proto.RegisterType((*CheckAuthReq)(nil), "data.CheckAuthReq")
	proto.RegisterType((*CheckAuthResp)(nil), "data.CheckAuthResp")
	proto.RegisterType((*SetSessionReq)(nil), "data.SetSessionReq")
	proto.RegisterType((*SetSessionResp)(nil), "data.SetSessionResp")
	proto.RegisterType((*GetSessionReq)(nil), "data.GetSessionReq")
	proto.RegisterType((*GetSessionResp)(nil), "data.GetSessionResp")
	proto.RegisterType((*SaveMsgReq)(nil), "data.SaveMsgReq")
	proto.RegisterType((*SaveMsgResp)(nil), "data.SaveMsgResp")
}

func init() { proto.RegisterFile("data/data.proto", fileDescriptor_48e8d492558f417f) }

var fileDescriptor_48e8d492558f417f = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x4d, 0xfd, 0x00, 0x99, 0xca, 0x87, 0x0b, 0x26, 0xa4, 0x17, 0xb5, 0xf1, 0x60, 0x38, 0x94,
	0x04, 0x8d, 0x1f, 0x47, 0xc1, 0xa4, 0xf1, 0x80, 0x21, 0xad, 0x77, 0x53, 0xcb, 0x50, 0x1a, 0x68,
	0x77, 0xed, 0x6e, 0x89, 0xfe, 0x60, 0xff, 0x87, 0xd9, 0x6e, 0xc5, 0x0a, 0x44, 0x0e, 0x5e, 0xda,
	0xe9, 0xeb, 0xbc, 0x99, 0xf7, 0x66, 0x77, 0xa0, 0x3e, 0xf6, 0x84, 0xd7, 0x95, 0x0f, 0x8b, 0x25,
	0x54, 0x50, 0xb2, 0x27, 0x63, 0xa3, 0xe9, 0xd3, 0x28, 0xa2, 0x71, 0x57, 0xbd, 0xd4, 0x2f, 0xf3,
	0x1a, 0x0e, 0x07, 0x53, 0xf4, 0x67, 0xf7, 0xa9, 0x98, 0x3a, 0xf8, 0x46, 0x1a, 0xb0, 0x9b, 0x86,
	0xe3, 0xb6, 0x76, 0xaa, 0x5d, 0x54, 0x1c, 0x19, 0x92, 0x16, 0xec, 0x0b, 0x3a, 0xc3, 0xb8, 0xbd,
	0x93, 0x61, 0xea, 0xc3, 0x1c, 0x42, 0xb5, 0xc0, 0xe3, 0x6c, 0x03, 0xb1, 0x03, 0xa5, 0x04, 0x79,
	0x3a, 0x17, 0x19, 0xb3, 0xd6, 0x23, 0x56, 0xde, 0xd9, 0xc9, 0xd0, 0xe7, 0x0f, 0x86, 0x4e, 0x9e,
	0x61, 0xf6, 0xa1, 0xea, 0xa2, 0x70, 0x91, 0xf3, 0x90, 0xc6, 0x9b, 0x75, 0x9c, 0x80, 0xce, 0x31,
	0x59, 0x60, 0xf2, 0x12, 0xc6, 0x13, 0x9a, 0xab, 0x01, 0x05, 0x3d, 0xc6, 0x13, 0x6a, 0x3e, 0x41,
	0xad, 0x58, 0xe3, 0xdf, 0x9a, 0xce, 0xa0, 0x6a, 0xff, 0xad, 0xc9, 0x1c, 0x40, 0xcd, 0xde, 0xd6,
	0x72, 0xab, 0xee, 0x5b, 0x00, 0xd7, 0x5b, 0xe0, 0x90, 0x07, 0xb2, 0x49, 0x07, 0x0e, 0x22, 0x1e,
	0xa8, 0x5c, 0x59, 0x45, 0xef, 0xd5, 0xbf, 0x35, 0x0e, 0x79, 0x20, 0x09, 0x4e, 0x39, 0x52, 0x81,
	0x79, 0x07, 0xfa, 0x92, 0xc9, 0x59, 0xc1, 0x9c, 0xb6, 0xcd, 0x5c, 0xef, 0x53, 0x03, 0xfd, 0xc1,
	0x13, 0x9e, 0x8b, 0xc9, 0x22, 0xf4, 0x91, 0x5c, 0x41, 0x65, 0x79, 0x9e, 0x84, 0x58, 0xd9, 0xe5,
	0x29, 0x5e, 0x0c, 0xa3, 0xb9, 0x86, 0x71, 0x46, 0x6e, 0x00, 0x7e, 0x46, 0x4e, 0xf2, 0x94, 0x5f,
	0x07, 0x69, 0xb4, 0xd6, 0x41, 0x45, 0xb4, 0xd7, 0x88, 0xf6, 0x26, 0xe2, 0xca, 0x7c, 0x2d, 0x28,
	0xe7, 0x96, 0x49, 0x23, 0xaf, 0xbc, 0x9c, 0x9d, 0x71, 0xb4, 0x82, 0x70, 0xd6, 0x3f, 0x87, 0xe3,
	0x90, 0x5a, 0x41, 0xc2, 0x7c, 0x0b, 0xdf, 0xbd, 0x88, 0xcd, 0x91, 0x5b, 0x5e, 0x2a, 0xa6, 0xfd,
	0x8a, 0x34, 0x31, 0x92, 0x3b, 0x30, 0xd2, 0x5e, 0x4b, 0xd9, 0x32, 0x5c, 0x7e, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x5d, 0x1d, 0x6b, 0x9a, 0x3a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataServiceClient interface {
	CheckAuth(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*CheckAuthResp, error)
	SetSession(ctx context.Context, in *SetSessionReq, opts ...grpc.CallOption) (*SetSessionResp, error)
	GetSession(ctx context.Context, in *GetSessionReq, opts ...grpc.CallOption) (*GetSessionResp, error)
	SaveMsg(ctx context.Context, in *SaveMsgReq, opts ...grpc.CallOption) (*SaveMsgResp, error)
}

type dataServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataServiceClient(cc *grpc.ClientConn) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) CheckAuth(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*CheckAuthResp, error) {
	out := new(CheckAuthResp)
	err := c.cc.Invoke(ctx, "/data.DataService/CheckAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) SetSession(ctx context.Context, in *SetSessionReq, opts ...grpc.CallOption) (*SetSessionResp, error) {
	out := new(SetSessionResp)
	err := c.cc.Invoke(ctx, "/data.DataService/SetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) GetSession(ctx context.Context, in *GetSessionReq, opts ...grpc.CallOption) (*GetSessionResp, error) {
	out := new(GetSessionResp)
	err := c.cc.Invoke(ctx, "/data.DataService/GetSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) SaveMsg(ctx context.Context, in *SaveMsgReq, opts ...grpc.CallOption) (*SaveMsgResp, error) {
	out := new(SaveMsgResp)
	err := c.cc.Invoke(ctx, "/data.DataService/SaveMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServiceServer is the server API for DataService service.
type DataServiceServer interface {
	CheckAuth(context.Context, *CheckAuthReq) (*CheckAuthResp, error)
	SetSession(context.Context, *SetSessionReq) (*SetSessionResp, error)
	GetSession(context.Context, *GetSessionReq) (*GetSessionResp, error)
	SaveMsg(context.Context, *SaveMsgReq) (*SaveMsgResp, error)
}

func RegisterDataServiceServer(s *grpc.Server, srv DataServiceServer) {
	s.RegisterService(&_DataService_serviceDesc, srv)
}

func _DataService_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.DataService/CheckAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).CheckAuth(ctx, req.(*CheckAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_SetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).SetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.DataService/SetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).SetSession(ctx, req.(*SetSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.DataService/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).GetSession(ctx, req.(*GetSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_SaveMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).SaveMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.DataService/SaveMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).SaveMsg(ctx, req.(*SaveMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAuth",
			Handler:    _DataService_CheckAuth_Handler,
		},
		{
			MethodName: "SetSession",
			Handler:    _DataService_SetSession_Handler,
		},
		{
			MethodName: "GetSession",
			Handler:    _DataService_GetSession_Handler,
		},
		{
			MethodName: "SaveMsg",
			Handler:    _DataService_SaveMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data/data.proto",
}
