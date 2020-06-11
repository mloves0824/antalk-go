package service

import (
	"antalk-go/internal/common"
	"antalk-go/internal/msg/dao"
	"antalk-go/internal/msg/model"
	proto "antalk-go/internal/proto/pb"
	"context"
	"sync/atomic"
	"time"
)

type Msg struct {
	db *dao.Dao
}

var msgId int64 = 0

func New(c *common.Config) *Msg  {
	s := &Msg{}
	return s
}

func getMsgId() int64 {
	atomic.AddInt64(&msgId, 1)
	return msgId
}

func (s *Msg) SendMsg(ctx context.Context, req *proto.CmdReq, resp *proto.CmdResp) (proto.ErrorCode, error) {
	sendMsgReq := &proto.SendMsgReq{}
	sendMsgResp := &proto.SendMsgResp{}

	if err := sendMsgReq.Unmarshal(req.Data); err != nil {
		return proto.ErrorCode_ERROR_INTERNAL, nil
	}

	//save msg
	msg := &model.Msg{}
	msg.SenderId = sendMsgReq.SenderId
	msg.RecvId = sendMsgReq.RecvId
	msg.Content = sendMsgReq.Content
	msg.MsgId = getMsgId()	//TODO messi: get id from seq server
	if err := s.db.SaveMsg(msg); err != nil {
		return proto.ErrorCode_ERROR_INTERNAL, nil
	}

	//reply to sender
	sendMsgResp.MsgId = msg.MsgId
	sendMsgResp.Timestamp = time.Now().Unix()

	//notify to recv

	return proto.ErrorCode_ERROR_NONE, nil
}