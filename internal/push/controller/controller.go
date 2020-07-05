package controller

import (
	"antalk-go/internal/common"
	proto "antalk-go/internal/proto/pb"
	"antalk-go/internal/push/service"
	"context"
)

type Controller struct {
	push *service.Push
}

func New(c *common.Config) (*Controller, error) {
	s := &Controller{}
	return s, nil
}

func (c *Controller) Cmd(ctx context.Context, req *proto.CmdReq, resp *proto.CmdResp) error {
	if req.Meta.CmdType == proto.CmdType_CMD_MSG_NOTIFY {
		meta := &proto.CmdMeta{}
		notifyReq := &proto.MsgNotifyReq{}
		if err := notifyReq.Unmarshal(req.Data); err != nil {
			meta.ErrorCode = proto.ErrorCode_ERROR_INTERNAL
			resp.Meta = meta
			//TODO messi: log
			return nil
		}

		//TODO messi: 参数判断
		notifyResp := &proto.MsgNotifyResp{}
		errCode := c.push.MsgNotify(ctx, notifyReq, notifyResp)
		if errCode != proto.ErrorCode_ERROR_NONE {
			meta.ErrorCode = errCode
			resp.Meta = meta
			return nil
		}

		data, err := notifyResp.Marshal()
		if err != nil {
			meta.ErrorCode = proto.ErrorCode_ERROR_INTERNAL
			resp.Meta = meta
			return nil
		}

		//success
		meta.ErrorCode = proto.ErrorCode_ERROR_NONE
		resp.Meta = meta
		resp.Data = data
		return nil
	}

	return nil
}
