package apigw

import "github.com/smallnest/rpcx/v5/server"

type Apigw struct {
	online bool
	config *Config

	server *server.Server
}

func NewApigw(config *Config) (*Apigw, error) {
	apigw := &Apigw{}
	apigw.config = config

	if err := apigw.setup(); err != nil {
		return nil, err
	}

	if err := apigw.register(); err != nil {
		return nil, err
	}

	if err := apigw.serve(); err != nil {
		return nil, err
	}

	return apigw, nil
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
