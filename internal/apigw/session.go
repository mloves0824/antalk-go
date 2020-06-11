package apigw

import "net"

type Session struct {
	conn net.Conn
}

func NewSession(sock net.Conn) *Session {
	s := &Session{}
	s.conn = sock
	return s
}