package service

import (
	"antalk-go/internal/common"
	proto "antalk-go/internal/proto/pb"
	"context"
)

type Push struct {

}

func New(c *common.Config) *Push {
	s := &Push{}
	return s
}


func (p *Push) msgNotify() error {
	return nil
}


func (p *Push) MsgNotify(ctx context.Context, req *proto.MsgNotifyReq, resp *proto.MsgNotifyResp) proto.ErrorCode {
	go p.msgNotify()

	resp.MsgId = req.MsgId
	return proto.ErrorCode_ERROR_NONE
}
