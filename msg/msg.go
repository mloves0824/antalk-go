package main

import (
	"fmt"
	//"io"
	"log"
	"net"
	//"strings"
	//"time"

	pb "github.com/mloves0824/antalk-go/proto/msg"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	name string
}

func newServer() *server {
	return &server{name: "MsgServer"}
}

func (s *server) MsgSend(context.Context, *pb.MsgSendReq) (*pb.MsgSendResp, error) {
	return &pb.MsgSendResp{}, nil
}

func main() {
	fmt.Printf("MsgServer Start!")
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	myServer := newServer()
	pb.RegisterMsgServiceServer(s, myServer)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
