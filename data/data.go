package main

import (
	"fmt"
	//"io"
	"log"
	"net"
	//"strings"
	"time"

	redis "github.com/mloves0824/antalk-go/pkg/utils/redis"
	pb "github.com/mloves0824/antalk-go/proto/data"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type redisConf struct {
	addr    string
	auth    string
	timeout time.Duration
}

type server struct {
	name         string
	redis_client *redis.Client
}

const ()

var redis_conf redisConf = redisConf{"localhost:6379", "antalk-auth", 1}

func newServer() *server {

	client, err := redis.NewClient(redis_conf.addr, redis_conf.auth, redis_conf.timeout)
	if err != nil {
		log.Fatalf("failed to NewClient: %v", err)
	}

	return &server{name: "DataServer", redis_client: client}
}

func (s *server) CheckAuth(context.Context, *pb.CheckAuthReq) (*pb.CheckAuthResp, error) {
	return &pb.CheckAuthResp{}, nil
}
func (s *server) SetSession(context.Context, *pb.SetSessionReq) (*pb.SetSessionResp, error) {
	return &pb.SetSessionResp{}, nil
}
func (s *server) GetSession(context.Context, *pb.GetSessionReq) (*pb.GetSessionResp, error) {
	return &pb.GetSessionResp{}, nil
}
func (s *server) SaveMsg(context.Context, *pb.SaveMsgReq) (*pb.SaveMsgResp, error) {
	return &pb.SaveMsgResp{}, nil
}

func main() {
	fmt.Printf("DataServer Start!")
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	myServer := newServer()
	pb.RegisterDataServiceServer(s, myServer)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
