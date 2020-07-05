package service

import (
	"antalk-go/internal/common"
	proto "antalk-go/internal/proto/pb"
	"antalk-go/internal/push/dao"
	"context"
)

type Push struct {
	dao *dao.Dao
	apigwClient *ApigwClient
}

var msgId int64 = 0

func New(c *common.Config) *Push {
	s := &Push{}
	return s
}

func (s *Push) MsgNotify(ctx context.Context, req *proto.MsgNotifyReq, resp *proto.MsgNotifyResp) proto.ErrorCode {
	go func() {
		//通过ReceiveUid查找对应的apigw服务
		route, err := s.dao.GetUserRouteInfo(req.ReceiveUid, "")
		if err != nil {
			return
		}

		s.apigwClient.SendMsgNotifyToApigw(route.ServerInfo, req)
	}()

	resp.MsgId = req.MsgId
	return proto.ErrorCode_ERROR_NONE
}
