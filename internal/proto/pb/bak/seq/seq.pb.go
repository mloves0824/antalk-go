// Code generated by protoc-gen-go. DO NOT EDIT.
// source: seq/seq.proto

package seq

import (
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

// 获取请求
type GetSeqReq struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSeqReq) Reset()         { *m = GetSeqReq{} }
func (m *GetSeqReq) String() string { return proto.CompactTextString(m) }
func (*GetSeqReq) ProtoMessage()    {}
func (*GetSeqReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2ae200e313a505b, []int{0}
}

func (m *GetSeqReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSeqReq.Unmarshal(m, b)
}
func (m *GetSeqReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSeqReq.Marshal(b, m, deterministic)
}
func (m *GetSeqReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSeqReq.Merge(m, src)
}
func (m *GetSeqReq) XXX_Size() int {
	return xxx_messageInfo_GetSeqReq.Size(m)
}
func (m *GetSeqReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSeqReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetSeqReq proto.InternalMessageInfo

func (m *GetSeqReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type GetSeqResp struct {
	SeqId                uint64   `protobuf:"varint,1,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSeqResp) Reset()         { *m = GetSeqResp{} }
func (m *GetSeqResp) String() string { return proto.CompactTextString(m) }
func (*GetSeqResp) ProtoMessage()    {}
func (*GetSeqResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2ae200e313a505b, []int{1}
}

func (m *GetSeqResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSeqResp.Unmarshal(m, b)
}
func (m *GetSeqResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSeqResp.Marshal(b, m, deterministic)
}
func (m *GetSeqResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSeqResp.Merge(m, src)
}
func (m *GetSeqResp) XXX_Size() int {
	return xxx_messageInfo_GetSeqResp.Size(m)
}
func (m *GetSeqResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSeqResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetSeqResp proto.InternalMessageInfo

func (m *GetSeqResp) GetSeqId() uint64 {
	if m != nil {
		return m.SeqId
	}
	return 0
}

func init() {
	proto.RegisterType((*GetSeqReq)(nil), "seq.GetSeqReq")
	proto.RegisterType((*GetSeqResp)(nil), "seq.GetSeqResp")
}

func init() { proto.RegisterFile("seq/seq.proto", fileDescriptor_f2ae200e313a505b) }

var fileDescriptor_f2ae200e313a505b = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0xd4,
	0x2f, 0x4e, 0x2d, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x4e, 0x2d, 0x54, 0x92,
	0xe5, 0xe2, 0x74, 0x4f, 0x2d, 0x09, 0x4e, 0x2d, 0x0c, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62, 0x2e,
	0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0x94, 0xb9, 0xb8, 0x60,
	0xd2, 0xc5, 0x05, 0x42, 0xa2, 0x5c, 0x6c, 0xc5, 0xa9, 0x85, 0xf1, 0x50, 0x25, 0x2c, 0x41, 0xac,
	0xc5, 0xa9, 0x85, 0x9e, 0x29, 0x46, 0xe6, 0x5c, 0x5c, 0xc1, 0xa9, 0x85, 0xc1, 0xa9, 0x45, 0x65,
	0x99, 0xc9, 0xa9, 0x42, 0x9a, 0x5c, 0x6c, 0x10, 0x2d, 0x42, 0x7c, 0x7a, 0x20, 0xcb, 0xe0, 0xc6,
	0x4b, 0xf1, 0xa3, 0xf0, 0x8b, 0x0b, 0x9c, 0xd4, 0xb8, 0xc4, 0x32, 0xf3, 0xf5, 0xd2, 0x8b, 0x0a,
	0x92, 0xf5, 0x52, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52, 0x8b, 0xf5, 0x12, 0x0b, 0x32, 0xd3, 0xcb,
	0x9d, 0xb8, 0x1c, 0x41, 0x54, 0x00, 0xc8, 0x9d, 0x01, 0x8c, 0x49, 0x6c, 0x60, 0x07, 0x1b, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xc2, 0x99, 0x92, 0xc1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SeqServiceClient is the client API for SeqService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SeqServiceClient interface {
	GetSeq(ctx context.Context, in *GetSeqReq, opts ...grpc.CallOption) (*GetSeqResp, error)
}

type seqServiceClient struct {
	cc *grpc.ClientConn
}

func NewSeqServiceClient(cc *grpc.ClientConn) SeqServiceClient {
	return &seqServiceClient{cc}
}

func (c *seqServiceClient) GetSeq(ctx context.Context, in *GetSeqReq, opts ...grpc.CallOption) (*GetSeqResp, error) {
	out := new(GetSeqResp)
	err := c.cc.Invoke(ctx, "/seq.SeqService/GetSeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeqServiceServer is the server API for SeqService service.
type SeqServiceServer interface {
	GetSeq(context.Context, *GetSeqReq) (*GetSeqResp, error)
}

func RegisterSeqServiceServer(s *grpc.Server, srv SeqServiceServer) {
	s.RegisterService(&_SeqService_serviceDesc, srv)
}

func _SeqService_GetSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeqServiceServer).GetSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seq.SeqService/GetSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeqServiceServer).GetSeq(ctx, req.(*GetSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SeqService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "seq.SeqService",
	HandlerType: (*SeqServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSeq",
			Handler:    _SeqService_GetSeq_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seq/seq.proto",
}
