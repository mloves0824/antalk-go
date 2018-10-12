package main

import (
	"fmt"
	//"io"
	"log"
	"net"
	//"strings"
	//"time"

	common_pb "github.com/mloves0824/antalk-go/proto/common"
	data_pb "github.com/mloves0824/antalk-go/proto/data"
	seq_pb "github.com/mloves0824/antalk-go/proto/seq"

	pb "github.com/mloves0824/antalk-go/proto/msg"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	data_addr = "localhost:50052"
	seq_addr  = "localhost:50055"
)

type server struct {
	name        string
	data_client data_pb.DataServiceClient
	seq_client  seq_pb.SeqServiceClient
}

func newServer() *server {
	conn, err := grpc.Dial(data_addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("conn to data error, %v", err)
	}
	new_data_client := data_pb.NewDataServiceClient(conn)

	seq_conn, err := grpc.Dial(seq_addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("conn to seq error, %v", err)
	}
	new_seq_client := seq_pb.NewSeqServiceClient(seq_conn)

	return &server{name: "MsgServer", data_client: new_data_client, seq_client: new_seq_client}
}

func (s *server) MsgSend(ctx context.Context, req *pb.MsgSendReq) (*pb.MsgSendResp, error) {
	seq_req := seq_pb.GetSeqReq{Uid: req.GetMsgInfo().GetRecvUid()}
	seq_resp, err := s.seq_client.GetSeq(context.Background(), &seq_req)
	if err != nil {
		//TODO: return InnerError
		return &pb.MsgSendResp{Result: common_pb.ResultType_ResultErrCheckAuth}, nil
	}
	log.Printf("Get seq success, seq=%lu", seq_resp.GetSeqId())
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
