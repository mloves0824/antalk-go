package apigw

import (
	"antalk-go/internal/common"
	"github.com/smallnest/rpcx/v5/server"
)

type Apigw struct {
	online bool
	config *common.Config

	server *server.Server
}

func NewApigw(c *common.Config) (*Apigw, error) {
	s := &Apigw{}
	s.config = c

	if err := s.setup(); err != nil {
		return nil, err
	}

	if err := s.register(c); err != nil {
		return nil, err
	}

	if err := s.serve(c); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Apigw) setup() error {
	s.server = server.NewServer()
	return nil
}

func (s *Apigw) register(c *common.Config) error {
	svc := &CmdService{}
	svc.Init(c)
	s.server.Register(svc, "")

	return nil
}

func (s *Apigw) serve(c *common.Config) error {
	s.server.Serve(c.GetString("rpcServer.proto"), c.GetString("rpcServer.addr"))
	return nil
}

func (s *Apigw) Close() {

}
