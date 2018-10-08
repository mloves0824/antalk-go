// Code generated by protoc-gen-go. DO NOT EDIT.
// source: apigw.proto

package apigw

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

type ResultType int32

const (
	ResultType_REFUSE_REASON_NONE               ResultType = 0
	ResultType_REFUSE_REASON_NO_MSG_SERVER      ResultType = 1
	ResultType_REFUSE_REASON_MSG_SERVER_FULL    ResultType = 2
	ResultType_REFUSE_REASON_NO_DB_SERVER       ResultType = 3
	ResultType_REFUSE_REASON_NO_LOGIN_SERVER    ResultType = 4
	ResultType_REFUSE_REASON_NO_ROUTE_SERVER    ResultType = 5
	ResultType_REFUSE_REASON_DB_VALIDATE_FAILED ResultType = 6
	ResultType_REFUSE_REASON_VERSION_TOO_OLD    ResultType = 7
	ResultType_REFUSE_REASON_INVALID_PASSWD     ResultType = 8
	ResultType_ERROR_CONNECT_TO_AUTH            ResultType = 9
	ResultType_ERROR_RPC_TO_AUTH                ResultType = 10
)

var ResultType_name = map[int32]string{
	0:  "REFUSE_REASON_NONE",
	1:  "REFUSE_REASON_NO_MSG_SERVER",
	2:  "REFUSE_REASON_MSG_SERVER_FULL",
	3:  "REFUSE_REASON_NO_DB_SERVER",
	4:  "REFUSE_REASON_NO_LOGIN_SERVER",
	5:  "REFUSE_REASON_NO_ROUTE_SERVER",
	6:  "REFUSE_REASON_DB_VALIDATE_FAILED",
	7:  "REFUSE_REASON_VERSION_TOO_OLD",
	8:  "REFUSE_REASON_INVALID_PASSWD",
	9:  "ERROR_CONNECT_TO_AUTH",
	10: "ERROR_RPC_TO_AUTH",
}

var ResultType_value = map[string]int32{
	"REFUSE_REASON_NONE":               0,
	"REFUSE_REASON_NO_MSG_SERVER":      1,
	"REFUSE_REASON_MSG_SERVER_FULL":    2,
	"REFUSE_REASON_NO_DB_SERVER":       3,
	"REFUSE_REASON_NO_LOGIN_SERVER":    4,
	"REFUSE_REASON_NO_ROUTE_SERVER":    5,
	"REFUSE_REASON_DB_VALIDATE_FAILED": 6,
	"REFUSE_REASON_VERSION_TOO_OLD":    7,
	"REFUSE_REASON_INVALID_PASSWD":     8,
	"ERROR_CONNECT_TO_AUTH":            9,
	"ERROR_RPC_TO_AUTH":                10,
}

func (x ResultType) String() string {
	return proto.EnumName(ResultType_name, int32(x))
}

func (ResultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_275560cf792c6ca5, []int{0}
}

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
	return fileDescriptor_275560cf792c6ca5, []int{1}
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
	return fileDescriptor_275560cf792c6ca5, []int{2}
}

// 登录请求
type LoginReq struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_275560cf792c6ca5, []int{0}
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

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResp struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ResultCode           ResultType `protobuf:"varint,2,opt,name=result_code,json=resultCode,proto3,enum=apigw.ResultType" json:"result_code,omitempty"`
	ResultString         string     `protobuf:"bytes,3,opt,name=result_string,json=resultString,proto3" json:"result_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_275560cf792c6ca5, []int{1}
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

func (m *LoginResp) GetResultCode() ResultType {
	if m != nil {
		return m.ResultCode
	}
	return ResultType_REFUSE_REASON_NONE
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
	return fileDescriptor_275560cf792c6ca5, []int{2}
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
	return fileDescriptor_275560cf792c6ca5, []int{3}
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
	return fileDescriptor_275560cf792c6ca5, []int{4}
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
	return fileDescriptor_275560cf792c6ca5, []int{5}
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
	return fileDescriptor_275560cf792c6ca5, []int{6}
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

type MsgNotify struct {
	MsgInfo              *MsgInfo `protobuf:"bytes,1,opt,name=msg_info,json=msgInfo,proto3" json:"msg_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgNotify) Reset()         { *m = MsgNotify{} }
func (m *MsgNotify) String() string { return proto.CompactTextString(m) }
func (*MsgNotify) ProtoMessage()    {}
func (*MsgNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_275560cf792c6ca5, []int{7}
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

func (m *MsgNotify) GetMsgInfo() *MsgInfo {
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
	return fileDescriptor_275560cf792c6ca5, []int{8}
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

func init() {
	proto.RegisterEnum("apigw.ResultType", ResultType_name, ResultType_value)
	proto.RegisterEnum("apigw.TopicType", TopicType_name, TopicType_value)
	proto.RegisterEnum("apigw.KickoutReason", KickoutReason_name, KickoutReason_value)
	proto.RegisterType((*LoginReq)(nil), "apigw.LoginReq")
	proto.RegisterType((*LoginResp)(nil), "apigw.LoginResp")
	proto.RegisterType((*LogoutReq)(nil), "apigw.LogoutReq")
	proto.RegisterType((*LogoutResp)(nil), "apigw.LogoutResp")
	proto.RegisterType((*Topic)(nil), "apigw.Topic")
	proto.RegisterType((*Notification)(nil), "apigw.Notification")
	proto.RegisterType((*MsgInfo)(nil), "apigw.MsgInfo")
	proto.RegisterType((*MsgNotify)(nil), "apigw.MsgNotify")
	proto.RegisterType((*KickNotify)(nil), "apigw.KickNotify")
}

func init() { proto.RegisterFile("apigw.proto", fileDescriptor_275560cf792c6ca5) }

var fileDescriptor_275560cf792c6ca5 = []byte{
	// 734 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdb, 0x6e, 0xf3, 0x44,
	0x10, 0xae, 0x73, 0xce, 0x24, 0x7f, 0x71, 0x17, 0xfa, 0x2b, 0x0d, 0xf4, 0x80, 0x29, 0xa8, 0x44,
	0x10, 0x50, 0x90, 0x10, 0x77, 0x28, 0x07, 0x37, 0x58, 0x75, 0xec, 0x68, 0xed, 0x94, 0xcb, 0x55,
	0x62, 0xbb, 0xee, 0xaa, 0x8d, 0xd7, 0x78, 0x9d, 0x96, 0x3c, 0x06, 0x0f, 0xc0, 0x1b, 0xf0, 0x66,
	0xbc, 0x04, 0xf2, 0xda, 0xce, 0xa1, 0xcd, 0x15, 0x77, 0x9e, 0xf9, 0xbe, 0xf9, 0xf2, 0xcd, 0x4c,
	0x76, 0xa0, 0x31, 0x0f, 0xa9, 0xff, 0xda, 0x0d, 0x23, 0x16, 0x33, 0x54, 0x16, 0x81, 0xf2, 0x0b,
	0xd4, 0x74, 0xe6, 0xd3, 0x00, 0x7b, 0x7f, 0x20, 0x19, 0x8a, 0x2b, 0xea, 0xb6, 0xa4, 0x2b, 0xe9,
	0xa6, 0x8e, 0x93, 0x4f, 0xd4, 0x86, 0x5a, 0x38, 0xe7, 0xfc, 0x95, 0x45, 0x6e, 0xab, 0x20, 0xd2,
	0x9b, 0x58, 0x79, 0x81, 0x7a, 0x56, 0xc9, 0xc3, 0x03, 0xa5, 0x3d, 0x68, 0x44, 0x1e, 0x5f, 0x3d,
	0xc7, 0xc4, 0x61, 0xae, 0x27, 0xaa, 0x8f, 0x7b, 0x27, 0xdd, 0xd4, 0x02, 0x16, 0x88, 0xbd, 0x0e,
	0x3d, 0x0c, 0x29, 0x6b, 0xc8, 0x5c, 0x0f, 0x7d, 0x05, 0x1f, 0xb2, 0x1a, 0x1e, 0x47, 0x34, 0xf0,
	0x5b, 0x45, 0xa1, 0xd7, 0x4c, 0x93, 0x96, 0xc8, 0x29, 0xe7, 0xe2, 0x77, 0xd9, 0x2a, 0x3e, 0x68,
	0x59, 0x69, 0x02, 0xe4, 0x30, 0x0f, 0x95, 0x5f, 0xa1, 0x6c, 0xb3, 0x90, 0x3a, 0x07, 0x0c, 0x5e,
	0x43, 0x29, 0x5e, 0x87, 0xb9, 0x33, 0x39, 0x73, 0x26, 0xd8, 0xc2, 0x98, 0x40, 0x95, 0xbf, 0x24,
	0x68, 0x1a, 0x2c, 0xa6, 0x0f, 0xd4, 0x99, 0xc7, 0x94, 0x05, 0xff, 0x57, 0x08, 0x7d, 0x0d, 0xa5,
	0x27, 0xea, 0x3c, 0x89, 0x96, 0x1a, 0x9b, 0x41, 0xdc, 0x51, 0xe7, 0x49, 0xc8, 0xaf, 0xb1, 0x80,
	0x91, 0x02, 0xc5, 0x25, 0xf7, 0x5b, 0x25, 0xc1, 0xca, 0xb5, 0x26, 0xdc, 0xcf, 0x48, 0x09, 0xa8,
	0xfc, 0x23, 0x41, 0x75, 0xc2, 0x7d, 0x2d, 0x78, 0x60, 0xe8, 0x0c, 0x6a, 0xdc, 0x0b, 0x5c, 0xb2,
	0xf5, 0x54, 0x4d, 0xe2, 0x19, 0x75, 0x13, 0x28, 0xf2, 0x9c, 0x17, 0x01, 0xa5, 0xcb, 0xab, 0x26,
	0x71, 0x02, 0x9d, 0x42, 0x65, 0xc9, 0x7d, 0x42, 0x5d, 0x61, 0xa7, 0x88, 0xcb, 0x4b, 0xee, 0x6b,
	0xa2, 0x22, 0x49, 0x8b, 0x6e, 0x4a, 0x69, 0xc5, 0x92, 0xfb, 0x49, 0x13, 0x39, 0xb4, 0x60, 0xee,
	0xba, 0x55, 0xde, 0x40, 0x03, 0xe6, 0xae, 0xd1, 0x05, 0xc0, 0x3c, 0x8e, 0xe7, 0xce, 0xe3, 0xd2,
	0x0b, 0xe2, 0x56, 0x45, 0x80, 0x3b, 0x19, 0xe5, 0x67, 0xa8, 0x6f, 0x1a, 0x40, 0xdf, 0xa6, 0x3a,
	0x34, 0x78, 0x60, 0xc2, 0x6f, 0xa3, 0x77, 0xbc, 0x6d, 0x32, 0xe9, 0x48, 0xe8, 0x26, 0x1f, 0x8a,
	0x0e, 0xb0, 0x1d, 0xcf, 0x81, 0xb9, 0x7f, 0x07, 0x95, 0xc8, 0x9b, 0x73, 0x16, 0x64, 0x93, 0xff,
	0x6c, 0x67, 0xa6, 0x62, 0xff, 0x09, 0x86, 0x33, 0x4e, 0xe7, 0xdf, 0x02, 0xc0, 0xf6, 0x6f, 0x87,
	0x3e, 0x02, 0xc2, 0xea, 0xed, 0xcc, 0x52, 0x09, 0x56, 0xfb, 0x96, 0x69, 0x10, 0xc3, 0x34, 0x54,
	0xf9, 0x08, 0x5d, 0xc2, 0xe7, 0x6f, 0xf3, 0x64, 0x62, 0x8d, 0x89, 0xa5, 0xe2, 0x7b, 0x15, 0xcb,
	0x12, 0xfa, 0x12, 0xce, 0xf7, 0x09, 0x5b, 0x94, 0xdc, 0xce, 0x74, 0x5d, 0x2e, 0xa0, 0x0b, 0x68,
	0xbf, 0xd3, 0x18, 0x0d, 0x72, 0x89, 0xe2, 0x7b, 0x09, 0xc3, 0x24, 0xba, 0x39, 0xd6, 0x8c, 0x9c,
	0x52, 0x3a, 0x48, 0xc1, 0xe6, 0xcc, 0x56, 0x73, 0x4a, 0x19, 0x5d, 0xc3, 0xd5, 0x3e, 0x65, 0x34,
	0x20, 0xf7, 0x7d, 0x5d, 0x1b, 0xf5, 0x6d, 0x95, 0xdc, 0xf6, 0x35, 0x5d, 0x1d, 0xc9, 0x95, 0xf7,
	0x42, 0xf7, 0x2a, 0xb6, 0x34, 0xd3, 0x20, 0xb6, 0x69, 0x12, 0x53, 0x1f, 0xc9, 0x55, 0x74, 0x05,
	0x5f, 0xec, 0x53, 0x34, 0x43, 0xe8, 0x90, 0x69, 0xdf, 0xb2, 0x7e, 0x1f, 0xc9, 0x35, 0x74, 0x06,
	0xa7, 0x2a, 0xc6, 0x26, 0x26, 0x43, 0xd3, 0x30, 0xd4, 0xa1, 0x4d, 0x6c, 0x93, 0xf4, 0x67, 0xf6,
	0x6f, 0x72, 0x1d, 0x9d, 0xc2, 0x49, 0x0a, 0xe1, 0xe9, 0x70, 0x93, 0x86, 0xce, 0x0f, 0x50, 0xdf,
	0x3c, 0x00, 0xd4, 0x80, 0xea, 0x9d, 0x36, 0xbc, 0x33, 0x67, 0xb6, 0x7c, 0x84, 0xaa, 0x50, 0x9c,
	0x58, 0x63, 0x59, 0x42, 0x4d, 0xa8, 0x8d, 0xcd, 0x19, 0x9e, 0x26, 0x51, 0xa1, 0x73, 0x09, 0x1f,
	0xf6, 0xf6, 0x86, 0x8e, 0x01, 0xcc, 0xf8, 0xd1, 0x8b, 0xc4, 0x8d, 0x91, 0x8f, 0x7a, 0x7f, 0x4b,
	0xd0, 0xec, 0x27, 0xfb, 0xb5, 0xbc, 0xe8, 0x85, 0x3a, 0x1e, 0xea, 0x40, 0x59, 0x60, 0xe8, 0x93,
	0x6c, 0xef, 0xf9, 0x1d, 0x6b, 0xcb, 0xfb, 0x09, 0x1e, 0xa2, 0xef, 0xa1, 0x92, 0x1e, 0x05, 0xb4,
	0x83, 0xa5, 0x27, 0xa4, 0x7d, 0xf2, 0x26, 0xc3, 0x43, 0xd4, 0x83, 0xba, 0xb5, 0x5a, 0x70, 0x27,
	0xa2, 0x0b, 0x0f, 0x35, 0x77, 0x1f, 0x74, 0xfb, 0xd3, 0x2c, 0xda, 0xbd, 0x09, 0x37, 0xd2, 0x8f,
	0xd2, 0xe0, 0x1b, 0xf8, 0x48, 0x59, 0xd7, 0x8f, 0x42, 0xa7, 0xeb, 0xfd, 0x39, 0x5f, 0x86, 0xcf,
	0x1e, 0x4f, 0xa9, 0x03, 0x10, 0xb6, 0xa7, 0xc9, 0xd5, 0x9d, 0x4a, 0x8b, 0x8a, 0x38, 0xbf, 0x3f,
	0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x20, 0xe8, 0x03, 0x3d, 0x8d, 0x05, 0x00, 0x00,
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

// ApigwServiceServer is the server API for ApigwService service.
type ApigwServiceServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Logout(context.Context, *LogoutReq) (*LogoutResp, error)
	Subscribe(ApigwService_SubscribeServer) error
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
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _ApigwService_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "apigw.proto",
}
