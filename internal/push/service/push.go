package service

import (
	"antalk-go/internal/common"
	proto "antalk-go/internal/proto/pb"
	"antalk-go/internal/push/dao"
	"context"
)

type Push struct {
	db *dao.Dao
}

var msgId int64 = 0

func New(c *common.Config) *Push {
	s := &Push{}
	return s
}

func (s *Push) MsgNotify(ctx context.Context, req *proto.CmdReq, resp *proto.CmdResp) (proto.ErrorCode, error) {
	return proto.ErrorCode_ERROR_NONE, nil
}
