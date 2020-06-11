package auth

import (
	"antalk-go/internal/auth/config"
	"antalk-go/internal/auth/controller"
	"antalk-go/internal/auth/service"
	"github.com/smallnest/rpcx/v5/server"
)

type AuthServer struct {
	online bool
	config *config.Config

	auth *service.AuthService
	server *server.Server
}

func New(c *config.Config) (*AuthServer, error) {
	s := &AuthServer{}
	s.config = c

	if err := s.setup(c); err != nil {
		return nil, err
	}

	if err := s.register(); err != nil {
		return nil, err
	}

	if err := s.serve(c); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *AuthServer) setup(c *config.Config) error {
	s.server = server.NewServer()

	s.auth = service.New(c)
	return nil
}

func (s *AuthServer) register() error {
	s.server.Register(controller.New(s.auth), "")

	return nil
}

func (s *AuthServer) serve(c *config.Config) error {
	s.server.Serve(c.GetString("rpcServer.proto"), c.GetString("rpcServer.addr"))
	return nil
}

func (s *AuthServer) Close() {

}
