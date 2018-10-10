package main

import (
	"fmt"
	//"io"
	"log"
	"net"
	//	"strings"
	"time"

	redis "github.com/mloves0824/antalk-go/pkg/utils/redis"
	commonpb "github.com/mloves0824/antalk-go/proto/common"
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

const (
	key_prefix_token = "TOEKN:"
)

var redis_conf redisConf = redisConf{"127.0.0.1:19000", "antalk-auth", 10}

func newServer() *server {

	client, err := redis.NewClientNoAuth(redis_conf.addr, redis_conf.timeout)
	if err != nil {
		log.Fatalf("failed to NewClient: %v", err)
	}

	return &server{name: "DataServer", redis_client: client}
}

func (s *server) CheckAuth(ctx context.Context, req *pb.CheckAuthReq) (*pb.CheckAuthResp, error) {
	cmd := fmt.Sprintf("GET %s", key_prefix_token+req.GetToken())
	value, err := s.redis_client.Do(cmd)
	if err != nil {
		return &pb.CheckAuthResp{Uid: req.GetUid(), Result: commonpb.ResultType_ResultErrRedis}, nil
	}
	if value != nil {
		return &pb.CheckAuthResp{Uid: req.GetUid(), Result: commonpb.ResultType_ResultErrCheckAuth}, nil
	}
	return &pb.CheckAuthResp{Uid: req.GetUid(), Result: commonpb.ResultType_ResultOK}, nil
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
