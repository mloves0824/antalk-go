package controller

import (
	proto "antalk-go/internal/proto/pb"
	"context"
)

type Controller struct {

}

func (c *Controller) Cmd(ctx context.Context, req *proto.CmdReq, resp *proto.CmdResp) error {
	if req.Meta.CmdType == proto.CmdType_CMD_MSG_NOTIFY {
		return c.MsgNotify(ctx, req, resp)
	}

	return nil
}

func (c *Controller) MsgNotify(ctx context.Context, req *proto.CmdReq, resp *proto.CmdResp) error {
	//协议解析
	msgReq := &proto.MsgNotifyReq{}
	if err := msgReq.Unmarshal(req.Data); err != nil {
		return err
	}
}