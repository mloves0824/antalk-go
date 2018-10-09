// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ResultType int32

const (
	ResultType_ResultOK           ResultType = 0
	ResultType_ResultErrCheckAuth ResultType = 1
	ResultType_ResultErrRedis     ResultType = 2
)

var ResultType_name = map[int32]string{
	0: "ResultOK",
	1: "ResultErrCheckAuth",
	2: "ResultErrRedis",
}

var ResultType_value = map[string]int32{
	"ResultOK":           0,
	"ResultErrCheckAuth": 1,
	"ResultErrRedis":     2,
}

func (x ResultType) String() string {
	return proto.EnumName(ResultType_name, int32(x))
}

func (ResultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type MsgInfo struct {
	SendUid              string   `protobuf:"bytes,1,opt,name=send_uid,json=sendUid,proto3" json:"send_uid,omitempty"`
	RecvUid              string   `protobuf:"bytes,2,opt,name=recv_uid,json=recvUid,proto3" json:"recv_uid,omitempty"`
	MsgId                int64    `protobuf:"varint,3,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	MsgType              string   `protobuf:"bytes,4,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
	MsgBody              string   `protobuf:"bytes,5,opt,name=msg_body,json=msgBody,proto3" json:"msg_body,omitempty"`
	Attachment           string   `protobuf:"bytes,6,opt,name=attachment,proto3" json:"attachment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgInfo) Reset()         { *m = MsgInfo{} }
func (m *MsgInfo) String() string { return proto.CompactTextString(m) }
func (*MsgInfo) ProtoMessage()    {}
func (*MsgInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *MsgInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgInfo.Unmarshal(m, b)
}
func (m *MsgInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgInfo.Marshal(b, m, deterministic)
}
func (m *MsgInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInfo.Merge(m, src)
}
func (m *MsgInfo) XXX_Size() int {
	return xxx_messageInfo_MsgInfo.Size(m)
}
func (m *MsgInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInfo proto.InternalMessageInfo

func (m *MsgInfo) GetSendUid() string {
	if m != nil {
		return m.SendUid
	}
	return ""
}

func (m *MsgInfo) GetRecvUid() string {
	if m != nil {
		return m.RecvUid
	}
	return ""
}

func (m *MsgInfo) GetMsgId() int64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *MsgInfo) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *MsgInfo) GetMsgBody() string {
	if m != nil {
		return m.MsgBody
	}
	return ""
}

func (m *MsgInfo) GetAttachment() string {
	if m != nil {
		return m.Attachment
	}
	return ""
}

func init() {
	proto.RegisterEnum("common.ResultType", ResultType_name, ResultType_value)
	proto.RegisterType((*MsgInfo)(nil), "common.MsgInfo")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0x45, 0x4d, 0x6b, 0xa7, 0xf5, 0x29, 0x52, 0x22, 0xea, 0xb8, 0x91, 0xe2, 0x6a, 0x70, 0x31,
	0x2e, 0xfc, 0x02, 0xa7, 0x28, 0x14, 0x11, 0xcb, 0xa0, 0xeb, 0x32, 0x4d, 0x9e, 0x99, 0x60, 0x33,
	0x09, 0x49, 0x46, 0x9c, 0x7f, 0xf2, 0x23, 0x25, 0x09, 0x16, 0x57, 0x8f, 0x73, 0x0f, 0xf7, 0x2d,
	0x2e, 0x9c, 0x31, 0xad, 0x94, 0xee, 0xee, 0xd2, 0x29, 0x8d, 0xd5, 0x5e, 0xd3, 0x2c, 0xd1, 0xcd,
	0x0f, 0x81, 0xe9, 0x8b, 0x13, 0xab, 0xee, 0x43, 0xd3, 0x2b, 0x98, 0x39, 0xec, 0xf8, 0xa6, 0x97,
	0x3c, 0x27, 0x0b, 0x52, 0x1c, 0xd5, 0xd3, 0xc0, 0xef, 0x92, 0x07, 0x65, 0x91, 0x7d, 0x45, 0x35,
	0x4a, 0x2a, 0x70, 0x50, 0xe7, 0x90, 0x29, 0x27, 0x36, 0x92, 0xe7, 0xe3, 0x05, 0x29, 0xc6, 0xf5,
	0x44, 0x39, 0xb1, 0x8a, 0x8d, 0x10, 0xfb, 0xc1, 0x60, 0x7e, 0x98, 0x1a, 0xca, 0x89, 0xb7, 0xc1,
	0xe0, 0x9f, 0xda, 0x6a, 0x3e, 0xe4, 0x93, 0xbd, 0xaa, 0x34, 0x1f, 0xe8, 0x35, 0x40, 0xe3, 0x7d,
	0xc3, 0x5a, 0x85, 0x9d, 0xcf, 0xb3, 0x28, 0xff, 0x25, 0xb7, 0x4f, 0x00, 0x35, 0xba, 0x7e, 0xe7,
	0xe3, 0xa3, 0x13, 0x98, 0x25, 0x7a, 0x7d, 0x9e, 0x1f, 0xd0, 0x0b, 0xa0, 0x89, 0x1e, 0xad, 0x5d,
	0xb6, 0xc8, 0x3e, 0x1f, 0x7a, 0xdf, 0xce, 0x09, 0xa5, 0x70, 0xba, 0xcf, 0x6b, 0xe4, 0xd2, 0xcd,
	0x47, 0x55, 0x01, 0x97, 0x52, 0x97, 0xc2, 0x1a, 0x56, 0xe2, 0x77, 0xa3, 0xcc, 0x0e, 0x5d, 0x99,
	0x16, 0xa9, 0x8e, 0x97, 0xf1, 0xae, 0xc3, 0x4c, 0x6b, 0xb2, 0xcd, 0xe2, 0x5e, 0xf7, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x4c, 0xd0, 0x47, 0x52, 0x46, 0x01, 0x00, 0x00,
}