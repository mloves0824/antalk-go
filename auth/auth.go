package main

import (
	"fmt"
	//"io"
	"log"
	"net"
	//"strings"
	//"time"

	pb "github.com/mloves0824/antalk-go/proto/auth"
	common_pb "github.com/mloves0824/antalk-go/proto/common"
	data_pb "github.com/mloves0824/antalk-go/proto/data"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	name        string
	data_client data_pb.DataServiceClient
}

const (
	data_addr string = "localhost:50052"
)

func newServer() *server {
	conn, err := grpc.Dial(data_addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to Dial DataSvr: %v", err)
	}
	new_data_client := data_pb.NewDataServiceClient(conn)
	return &server{name: "AuthServer", data_client: new_data_client}
}

func (s *server) CheckAuth(ctx context.Context, req *pb.CheckAuthReq) (*pb.CheckAuthResp, error) {
	data_req := &data_pb.CheckAuthReq{Uid: req.GetUid(), Token: req.GetToken()}
	data_resp, err := s.data_client.CheckAuth(context.Background(), data_req)
	if err != nil {
		return &pb.CheckAuthResp{Result: common_pb.ResultType_ResultErrCheckAuth}, nil
	}
	log.Printf("send CheckAuthReq success: result=%s", data_resp.GetResult())
	return &pb.CheckAuthResp{Result: data_resp.GetResult()}, nil
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
