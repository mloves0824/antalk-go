package rpcx

import (
	"antalk-go/internal/common"
	"antalk-go/internal/msg/service"
	"github.com/smallnest/rpcx/v5/server"
)

type Server struct {
	config *common.Config
	engine *server.Server
	svc *service.Msg
}

func New(c *common.Config) (*Server, error) {
	s := &Server{}
	s.config = c

	s.engine = server.NewServer()
	s.svc = service.New(c)
	s.engine.Register(s.svc, "")
	s.engine.Serve(c.GetString("rpcServer.proto"), c.GetString("rpcServer.addr"))

	return s, nil
}
