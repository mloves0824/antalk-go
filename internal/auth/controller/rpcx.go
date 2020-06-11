package controller

import (
	"antalk-go/internal/auth/service"
	"antalk-go/internal/proto/pb"
	"context"
)

type RpcxController struct {
	svc *service.AuthService
}

func New(svc *service.AuthService) *RpcxController {
	c := &RpcxController{svc:svc}
	return c
}

func (c *RpcxController) Cmd(ctx context.Context, req *proto.CmdReq, resp *proto.CmdResp) error {
	if req.Meta.CmdType == proto.CmdType_CMD_LOGIN {
		return c.svc.Login(ctx, req, resp)
	}

	return nil
}


