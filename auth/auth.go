package main

import (
	"fmt"
	//"io"
	"log"
	"net"
	//"strings"
	//"time"

	pb "github.com/mloves0824/antalk-go/proto/auth"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	name string
}

func newServer() *server {
	return &server{name: "AuthServer"}
}

func (s *server) CheckAuth(context.Context, *pb.CheckAuthReq) (*pb.CheckAuthResp, error) {
	return &pb.CheckAuthResp{}, nil
}

func main() {
	fmt.Printf("AuthServer Start!")
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	myServer := newServer()
	pb.RegisterAuthServiceServer(s, myServer)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
