package apigw

import "github.com/smallnest/rpcx/v5/server"

type Apigw struct {
	online bool
	config *Config

	server *server.Server
}

func NewApigw(config *Config) (*Apigw, error) {
	s := &Apigw{}
	s.config = config

	if err := s.setup(); err != nil {
		return nil, err
	}

	if err := s.register(); err != nil {
		return nil, err
	}

	if err := s.serve(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Apigw) setup() error {
	s.server = server.NewServer()
	return nil
}

func (s *Apigw) register() error {
	s.server.Register(new(AuthService), "")

	return nil
}

func (s *Apigw) serve() error {
	s.server.Serve("tcp", ":11111")
	return nil
}

func (s *Apigw) Close() {

}
