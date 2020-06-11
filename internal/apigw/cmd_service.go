package apigw

import (
	"antalk-go/internal/common"
	"antalk-go/internal/proto/pb"
	"context"
	"github.com/smallnest/rpcx/v5/server"
	"log"
	"net"
)

type CmdService struct {
	authClient AuthClient
	sessions []*Session
}

func (t *CmdService) Init(c *common.Config)  {
	t.authClient.Init(c)
}

func (t *CmdService) Cmd(ctx context.Context, req *proto.CmdReq, resp *proto.CmdResp) error {
	if req.Meta.CmdType == proto.CmdType_CMD_LOGIN {
		reply := &proto.CmdResp{}
		err := t.authClient.client.Call(context.Background(), "Cmd", req, reply)
		if err != nil {
			log.Fatal("failed to call: ", err)
		}

		if reply.Meta.ErrorCode == proto.ErrorCode_ERROR_NONE {
			clientConn := ctx.Value(server.RemoteConnContextKey).(net.Conn)
			t.sessions = append(t.sessions, NewSession(clientConn))
		}
	}


	return nil
}
