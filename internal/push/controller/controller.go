package controller

import (
	proto "antalk-go/internal/proto/pb"
	"antalk-go/internal/push/service"
	"context"
	"log"
)

type Controller struct {
	msg *service.Push
}

func New() *Controller {
	c := &Controller{}
	return c
}

func (c *Controller) Cmd(ctx context.Context, req *proto.CmdReq, resp *proto.CmdResp) error {
	if req.Meta.CmdType == proto.CmdType_CMD_MSG {
		reply := &proto.CmdResp{}
		err := t.authClient.client.Call(context.Background(), "Cmd", req, reply)
		if err != nil {
			log.Fatal("failed to call: ", err)
		}

	}
}
