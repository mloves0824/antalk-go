// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apigw/apigw.proto

package apigw

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

type TopicType int32

const (
	TopicType_KICKOUT  TopicType = 0
	TopicType_MSG      TopicType = 1
	TopicType_GOURPMSG TopicType = 2
)

var TopicType_name = map[int32]string{
	0: "KICKOUT",
	1: "MSG",
	2: "GOURPMSG",
}

var TopicType_value = map[string]int32{
	"KICKOUT":  0,
	"MSG":      1,
	"GOURPMSG": 2,
}

func (x TopicType) String() string {
	return proto.EnumName(TopicType_name, int32(x))
}

func (TopicType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{0}
}

type KickoutReason int32

const (
	KickoutReason_OtherLogin KickoutReason = 0
)

var KickoutReason_name = map[int32]string{
	0: "OtherLogin",
}

var KickoutReason_value = map[string]int32{
	"OtherLogin": 0,
}

func (x KickoutReason) String() string {
	return proto.EnumName(KickoutReason_name, int32(x))
}

func (KickoutReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{1}
}

// 登录请求
type LoginReq struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *LoginReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginReq) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type LoginResp struct {
	Uid                  string            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ResultCode           common.ResultType `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3,enum=common.ResultType" json:"result_code,omitempty"`
	ResultString         string            `protobuf:"bytes,3,opt,name=result_string,json=resultString,proto3" json:"result_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{1}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *LoginResp) GetResultCode() common.ResultType {
	if m != nil {
		return m.ResultCode
	}
	return common.ResultType_ResultOK
}

func (m *LoginResp) GetResultString() string {
	if m != nil {
		return m.ResultString
	}
	return ""
}

// 登出请求
type LogoutReq struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutReq) Reset()         { *m = LogoutReq{} }
func (m *LogoutReq) String() string { return proto.CompactTextString(m) }
func (*LogoutReq) ProtoMessage()    {}
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{2}
}

func (m *LogoutReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutReq.Unmarshal(m, b)
}
func (m *LogoutReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutReq.Marshal(b, m, deterministic)
}
func (m *LogoutReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutReq.Merge(m, src)
}
func (m *LogoutReq) XXX_Size() int {
	return xxx_messageInfo_LogoutReq.Size(m)
}
func (m *LogoutReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutReq.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutReq proto.InternalMessageInfo

func (m *LogoutReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type LogoutResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResp) Reset()         { *m = LogoutResp{} }
func (m *LogoutResp) String() string { return proto.CompactTextString(m) }
func (*LogoutResp) ProtoMessage()    {}
func (*LogoutResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{3}
}

func (m *LogoutResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResp.Unmarshal(m, b)
}
func (m *LogoutResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResp.Marshal(b, m, deterministic)
}
func (m *LogoutResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResp.Merge(m, src)
}
func (m *LogoutResp) XXX_Size() int {
	return xxx_messageInfo_LogoutResp.Size(m)
}
func (m *LogoutResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResp.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResp proto.InternalMessageInfo

type Topic struct {
	Uid                  string    `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type                 TopicType `protobuf:"varint,2,opt,name=type,proto3,enum=apigw.TopicType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{4}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Topic) GetType() TopicType {
	if m != nil {
		return m.Type
	}
	return TopicType_KICKOUT
}

type Notification struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Type                 TopicType   `protobuf:"varint,2,opt,name=type,proto3,enum=apigw.TopicType" json:"type,omitempty"`
	Kick                 *KickNotify `protobuf:"bytes,3,opt,name=kick,proto3" json:"kick,omitempty"`
	Msg                  *MsgNotify  `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{5}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Notification) GetType() TopicType {
	if m != nil {
		return m.Type
	}
	return TopicType_KICKOUT
}

func (m *Notification) GetKick() *KickNotify {
	if m != nil {
		return m.Kick
	}
	return nil
}

func (m *Notification) GetMsg() *MsgNotify {
	if m != nil {
		return m.Msg
	}
	return nil
}

type MsgNotify struct {
	MsgInfo              *common.MsgInfo `protobuf:"bytes,1,opt,name=msg_info,json=msgInfo,proto3" json:"msg_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MsgNotify) Reset()         { *m = MsgNotify{} }
func (m *MsgNotify) String() string { return proto.CompactTextString(m) }
func (*MsgNotify) ProtoMessage()    {}
func (*MsgNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{6}
}

func (m *MsgNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgNotify.Unmarshal(m, b)
}
func (m *MsgNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgNotify.Marshal(b, m, deterministic)
}
func (m *MsgNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNotify.Merge(m, src)
}
func (m *MsgNotify) XXX_Size() int {
	return xxx_messageInfo_MsgNotify.Size(m)
}
func (m *MsgNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNotify.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNotify proto.InternalMessageInfo

func (m *MsgNotify) GetMsgInfo() *common.MsgInfo {
	if m != nil {
		return m.MsgInfo
	}
	return nil
}

type KickNotify struct {
	Uid                  string        `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Reason               KickoutReason `protobuf:"varint,2,opt,name=reason,proto3,enum=apigw.KickoutReason" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *KickNotify) Reset()         { *m = KickNotify{} }
func (m *KickNotify) String() string { return proto.CompactTextString(m) }
func (*KickNotify) ProtoMessage()    {}
func (*KickNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{7}
}

func (m *KickNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KickNotify.Unmarshal(m, b)
}
func (m *KickNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KickNotify.Marshal(b, m, deterministic)
}
func (m *KickNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KickNotify.Merge(m, src)
}
func (m *KickNotify) XXX_Size() int {
	return xxx_messageInfo_KickNotify.Size(m)
}
func (m *KickNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_KickNotify.DiscardUnknown(m)
}

var xxx_messageInfo_KickNotify proto.InternalMessageInfo

func (m *KickNotify) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *KickNotify) GetReason() KickoutReason {
	if m != nil {
		return m.Reason
	}
	return KickoutReason_OtherLogin
}

type MsgSendReq struct {
	MsgInfo              *common.MsgInfo `protobuf:"bytes,1,opt,name=msg_info,json=msgInfo,proto3" json:"msg_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MsgSendReq) Reset()         { *m = MsgSendReq{} }
func (m *MsgSendReq) String() string { return proto.CompactTextString(m) }
func (*MsgSendReq) ProtoMessage()    {}
func (*MsgSendReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{8}
}

func (m *MsgSendReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSendReq.Unmarshal(m, b)
}
func (m *MsgSendReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSendReq.Marshal(b, m, deterministic)
}
func (m *MsgSendReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendReq.Merge(m, src)
}
func (m *MsgSendReq) XXX_Size() int {
	return xxx_messageInfo_MsgSendReq.Size(m)
}
func (m *MsgSendReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendReq.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendReq proto.InternalMessageInfo

func (m *MsgSendReq) GetMsgInfo() *common.MsgInfo {
	if m != nil {
		return m.MsgInfo
	}
	return nil
}

type MsgSendResp struct {
	Result               common.ResultType `protobuf:"varint,1,opt,name=result,proto3,enum=common.ResultType" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MsgSendResp) Reset()         { *m = MsgSendResp{} }
func (m *MsgSendResp) String() string { return proto.CompactTextString(m) }
func (*MsgSendResp) ProtoMessage()    {}
func (*MsgSendResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{9}
}

func (m *MsgSendResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSendResp.Unmarshal(m, b)
}
func (m *MsgSendResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSendResp.Marshal(b, m, deterministic)
}
func (m *MsgSendResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendResp.Merge(m, src)
}
func (m *MsgSendResp) XXX_Size() int {
	return xxx_messageInfo_MsgSendResp.Size(m)
}
func (m *MsgSendResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendResp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendResp proto.InternalMessageInfo

func (m *MsgSendResp) GetResult() common.ResultType {
	if m != nil {
		return m.Result
	}
	return common.ResultType_ResultOK
}

type MsgPushReq struct {
	MsgInfo              *common.MsgInfo `protobuf:"bytes,1,opt,name=msg_info,json=msgInfo,proto3" json:"msg_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MsgPushReq) Reset()         { *m = MsgPushReq{} }
func (m *MsgPushReq) String() string { return proto.CompactTextString(m) }
func (*MsgPushReq) ProtoMessage()    {}
func (*MsgPushReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{10}
}

func (m *MsgPushReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgPushReq.Unmarshal(m, b)
}
func (m *MsgPushReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgPushReq.Marshal(b, m, deterministic)
}
func (m *MsgPushReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPushReq.Merge(m, src)
}
func (m *MsgPushReq) XXX_Size() int {
	return xxx_messageInfo_MsgPushReq.Size(m)
}
func (m *MsgPushReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPushReq.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPushReq proto.InternalMessageInfo

func (m *MsgPushReq) GetMsgInfo() *common.MsgInfo {
	if m != nil {
		return m.MsgInfo
	}
	return nil
}

type MsgPushResp struct {
	Result               common.ResultType `protobuf:"varint,1,opt,name=result,proto3,enum=common.ResultType" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MsgPushResp) Reset()         { *m = MsgPushResp{} }
func (m *MsgPushResp) String() string { return proto.CompactTextString(m) }
func (*MsgPushResp) ProtoMessage()    {}
func (*MsgPushResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a764574211fe91, []int{11}
}

func (m *MsgPushResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgPushResp.Unmarshal(m, b)
}
func (m *MsgPushResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgPushResp.Marshal(b, m, deterministic)
}
func (m *MsgPushResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPushResp.Merge(m, src)
}
func (m *MsgPushResp) XXX_Size() int {
	return xxx_messageInfo_MsgPushResp.Size(m)
}
func (m *MsgPushResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPushResp.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPushResp proto.InternalMessageInfo

func (m *MsgPushResp) GetResult() common.ResultType {
	if m != nil {
		return m.Result
	}
	return common.ResultType_ResultOK
}

func init() {
	proto.RegisterEnum("apigw.TopicType", TopicType_name, TopicType_value)
	proto.RegisterEnum("apigw.KickoutReason", KickoutReason_name, KickoutReason_value)
	proto.RegisterType((*LoginReq)(nil), "apigw.LoginReq")
	proto.RegisterType((*LoginResp)(nil), "apigw.LoginResp")
	proto.RegisterType((*LogoutReq)(nil), "apigw.LogoutReq")
	proto.RegisterType((*LogoutResp)(nil), "apigw.LogoutResp")
	proto.RegisterType((*Topic)(nil), "apigw.Topic")
	proto.RegisterType((*Notification)(nil), "apigw.Notification")
	proto.RegisterType((*MsgNotify)(nil), "apigw.MsgNotify")
	proto.RegisterType((*KickNotify)(nil), "apigw.KickNotify")
	proto.RegisterType((*MsgSendReq)(nil), "apigw.MsgSendReq")
	proto.RegisterType((*MsgSendResp)(nil), "apigw.MsgSendResp")
	proto.RegisterType((*MsgPushReq)(nil), "apigw.MsgPushReq")
	proto.RegisterType((*MsgPushResp)(nil), "apigw.MsgPushResp")
}

func init() { proto.RegisterFile("apigw/apigw.proto", fileDescriptor_a9a764574211fe91) }

var fileDescriptor_a9a764574211fe91 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4b, 0x6f, 0xda, 0x40,
	0x10, 0x8e, 0x43, 0x78, 0x78, 0x70, 0x12, 0x33, 0x89, 0x2a, 0x84, 0x5a, 0x35, 0x72, 0x1f, 0x42,
	0xa8, 0x85, 0x88, 0x1c, 0xda, 0x9e, 0xaa, 0x26, 0x87, 0x28, 0x22, 0x14, 0x64, 0xc8, 0x39, 0x02,
	0xb3, 0x38, 0x2b, 0x62, 0xef, 0xd6, 0x6b, 0x9a, 0x72, 0xec, 0x4f, 0xe8, 0x3f, 0xae, 0x3c, 0x5e,
	0x83, 0x13, 0x71, 0x49, 0x2e, 0xe0, 0x9d, 0xef, 0x9b, 0xd7, 0x37, 0xb3, 0x0b, 0xb5, 0x89, 0xe4,
	0xfe, 0x43, 0x87, 0x7e, 0xdb, 0x32, 0x12, 0xb1, 0xc0, 0x22, 0x1d, 0x1a, 0x47, 0x9e, 0x08, 0x02,
	0x11, 0x76, 0xd2, 0xbf, 0x14, 0x73, 0x86, 0x50, 0xb9, 0x16, 0x3e, 0x0f, 0x5d, 0xf6, 0x0b, 0x6d,
	0x28, 0x2c, 0xf9, 0xac, 0x6e, 0x9c, 0x18, 0x4d, 0xd3, 0x4d, 0x3e, 0xf1, 0x18, 0x8a, 0xb1, 0x58,
	0xb0, 0xb0, 0xbe, 0x4b, 0xb6, 0xf4, 0x80, 0xaf, 0xc1, 0x8c, 0x79, 0xc0, 0x54, 0x3c, 0x09, 0x64,
	0xbd, 0x70, 0x62, 0x34, 0x0b, 0xee, 0xc6, 0xe0, 0x3c, 0x80, 0xa9, 0x23, 0x2a, 0xb9, 0x25, 0xe4,
	0x19, 0x54, 0x23, 0xa6, 0x96, 0xf7, 0xf1, 0xad, 0x27, 0x66, 0x8c, 0x02, 0x1f, 0x74, 0xb1, 0xad,
	0x8b, 0x72, 0x09, 0x1a, 0xaf, 0x24, 0x73, 0x21, 0xa5, 0x5d, 0x88, 0x19, 0xc3, 0x77, 0xb0, 0xaf,
	0x9d, 0x54, 0x1c, 0xf1, 0xd0, 0xa7, 0xac, 0xa6, 0x6b, 0xa5, 0xc6, 0x11, 0xd9, 0x9c, 0x37, 0x94,
	0x58, 0x2c, 0xe3, 0xad, 0xbd, 0x38, 0x16, 0x40, 0x06, 0x2b, 0xe9, 0x7c, 0x87, 0xe2, 0x58, 0x48,
	0xee, 0x6d, 0xa9, 0xf0, 0x3d, 0xec, 0xc5, 0x2b, 0x99, 0x95, 0x66, 0xb7, 0x53, 0x29, 0x89, 0x4d,
	0x85, 0x11, 0xea, 0xfc, 0x33, 0xc0, 0xfa, 0x29, 0x62, 0x3e, 0xe7, 0xde, 0x24, 0xe6, 0x22, 0x7c,
	0x69, 0x20, 0xfc, 0x00, 0x7b, 0x0b, 0xee, 0x2d, 0xa8, 0xa5, 0x6a, 0xb7, 0xa6, 0x59, 0x3d, 0xee,
	0x2d, 0x28, 0xfc, 0xca, 0x25, 0x18, 0x1d, 0x28, 0x04, 0xca, 0xaf, 0xef, 0x11, 0x2b, 0x8b, 0xd5,
	0x57, 0xbe, 0x26, 0x25, 0xa0, 0xf3, 0x05, 0xcc, 0xb5, 0x05, 0x5b, 0x50, 0x09, 0x94, 0x7f, 0xcb,
	0xc3, 0xb9, 0xa0, 0xa2, 0xaa, 0xdd, 0xc3, 0x4c, 0xe5, 0xbe, 0xf2, 0xaf, 0xc2, 0xb9, 0x70, 0xcb,
	0x41, 0xfa, 0xe1, 0x5c, 0x03, 0x6c, 0x12, 0x6e, 0xe9, 0xe4, 0x13, 0x94, 0x22, 0x36, 0x51, 0x22,
	0xd4, 0xbd, 0x1c, 0xe7, 0xaa, 0x24, 0x45, 0x13, 0xcc, 0xd5, 0x1c, 0xe7, 0x2b, 0x40, 0x5f, 0xf9,
	0x23, 0x16, 0xce, 0x92, 0x49, 0x3c, 0xa7, 0x8e, 0x6f, 0x50, 0x5d, 0x7b, 0x2a, 0x89, 0xad, 0x24,
	0x6d, 0x32, 0x61, 0x72, 0xdc, 0xbe, 0x26, 0x9a, 0xa1, 0x93, 0x0e, 0x97, 0xea, 0xee, 0x65, 0x49,
	0x53, 0xcf, 0xe7, 0x25, 0x6d, 0x75, 0xc0, 0x5c, 0x8f, 0x13, 0xab, 0x50, 0xee, 0x5d, 0x5d, 0xf4,
	0x06, 0x37, 0x63, 0x7b, 0x07, 0xcb, 0x50, 0xe8, 0x8f, 0x2e, 0x6d, 0x03, 0x2d, 0xa8, 0x5c, 0x0e,
	0x6e, 0xdc, 0x61, 0x72, 0xda, 0x6d, 0xbd, 0x85, 0xfd, 0x47, 0x9a, 0xe1, 0x01, 0xc0, 0x20, 0xbe,
	0x63, 0x11, 0x5d, 0x19, 0x7b, 0xa7, 0xfb, 0x77, 0x17, 0xac, 0x1f, 0x89, 0xb6, 0x23, 0x16, 0xfd,
	0xe6, 0x1e, 0xc3, 0x16, 0x14, 0x09, 0xc3, 0x43, 0xad, 0x79, 0x76, 0x5d, 0x1b, 0xf6, 0x63, 0x83,
	0x92, 0xf8, 0x19, 0x4a, 0xe9, 0x8a, 0x63, 0x0e, 0x4b, 0x2f, 0x44, 0xa3, 0xf6, 0xc4, 0xa2, 0x24,
	0x76, 0xc1, 0x1c, 0x2d, 0xa7, 0xca, 0x8b, 0xf8, 0x94, 0xa1, 0x95, 0x5f, 0xcf, 0xc6, 0x91, 0x3e,
	0xe5, 0x37, 0xbc, 0x69, 0x9c, 0x1a, 0x78, 0x0a, 0x65, 0x3d, 0x21, 0xac, 0x6d, 0x96, 0x50, 0xcf,
	0xba, 0x81, 0x4f, 0x4d, 0x4a, 0x6a, 0x8f, 0x44, 0xde, 0xbc, 0x87, 0x1e, 0x54, 0xde, 0x23, 0x9b,
	0xc0, 0xf9, 0x47, 0x78, 0xc5, 0x45, 0xdb, 0x8f, 0xa4, 0xd7, 0x66, 0x7f, 0x26, 0x81, 0xbc, 0x67,
	0x2a, 0x65, 0x9d, 0x03, 0x49, 0x33, 0x4c, 0x5e, 0xae, 0xa1, 0x31, 0x2d, 0xd1, 0x13, 0x76, 0xf6,
	0x3f, 0x00, 0x00, 0xff, 0xff, 0xee, 0x97, 0x51, 0x97, 0xf3, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApigwServiceClient is the client API for ApigwService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApigwServiceClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error)
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (ApigwService_SubscribeClient, error)
	MsgSend(ctx context.Context, in *MsgSendReq, opts ...grpc.CallOption) (*MsgSendResp, error)
	MsgPush(ctx context.Context, in *MsgPushReq, opts ...grpc.CallOption) (*MsgPushResp, error)
}

type apigwServiceClient struct {
	cc *grpc.ClientConn
}

func NewApigwServiceClient(cc *grpc.ClientConn) ApigwServiceClient {
	return &apigwServiceClient{cc}
}

func (c *apigwServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/apigw.ApigwService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apigwServiceClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error) {
	out := new(LogoutResp)
	err := c.cc.Invoke(ctx, "/apigw.ApigwService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apigwServiceClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (ApigwService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ApigwService_serviceDesc.Streams[0], "/apigw.ApigwService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &apigwServiceSubscribeClient{stream}
	return x, nil
}

type ApigwService_SubscribeClient interface {
	Send(*Topic) error
	Recv() (*Notification, error)
	grpc.ClientStream
}

type apigwServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *apigwServiceSubscribeClient) Send(m *Topic) error {
	return x.ClientStream.SendMsg(m)
}

func (x *apigwServiceSubscribeClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *apigwServiceClient) MsgSend(ctx context.Context, in *MsgSendReq, opts ...grpc.CallOption) (*MsgSendResp, error) {
	out := new(MsgSendResp)
	err := c.cc.Invoke(ctx, "/apigw.ApigwService/MsgSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apigwServiceClient) MsgPush(ctx context.Context, in *MsgPushReq, opts ...grpc.CallOption) (*MsgPushResp, error) {
	out := new(MsgPushResp)
	err := c.cc.Invoke(ctx, "/apigw.ApigwService/MsgPush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApigwServiceServer is the server API for ApigwService service.
type ApigwServiceServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Logout(context.Context, *LogoutReq) (*LogoutResp, error)
	Subscribe(ApigwService_SubscribeServer) error
	MsgSend(context.Context, *MsgSendReq) (*MsgSendResp, error)
	MsgPush(context.Context, *MsgPushReq) (*MsgPushResp, error)
}

func RegisterApigwServiceServer(s *grpc.Server, srv ApigwServiceServer) {
	s.RegisterService(&_ApigwService_serviceDesc, srv)
}

func _ApigwService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApigwServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apigw.ApigwService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApigwServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApigwService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApigwServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apigw.ApigwService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApigwServiceServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApigwService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ApigwServiceServer).Subscribe(&apigwServiceSubscribeServer{stream})
}

type ApigwService_SubscribeServer interface {
	Send(*Notification) error
	Recv() (*Topic, error)
	grpc.ServerStream
}

type apigwServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *apigwServiceSubscribeServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func (x *apigwServiceSubscribeServer) Recv() (*Topic, error) {
	m := new(Topic)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ApigwService_MsgSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApigwServiceServer).MsgSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apigw.ApigwService/MsgSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApigwServiceServer).MsgSend(ctx, req.(*MsgSendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApigwService_MsgPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPushReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApigwServiceServer).MsgPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apigw.ApigwService/MsgPush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApigwServiceServer).MsgPush(ctx, req.(*MsgPushReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApigwService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apigw.ApigwService",
	HandlerType: (*ApigwServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ApigwService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _ApigwService_Logout_Handler,
		},
		{
			MethodName: "MsgSend",
			Handler:    _ApigwService_MsgSend_Handler,
		},
		{
			MethodName: "MsgPush",
			Handler:    _ApigwService_MsgPush_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _ApigwService_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "apigw/apigw.proto",
}
