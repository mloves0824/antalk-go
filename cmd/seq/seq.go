package main

import (
	"log"
	"net"
	"time"

	pb "github.com/mloves0824/antalk-go/proto/seq"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	listen_addr = "localhost:50055"
)

type server struct {
	name string
}

func newServer() *server {
	return &server{"SeqServer"}
}

func (s *server) GetSeq(ctx context.Context, req *pb.GetSeqReq) (resp *pb.GetSeqResp, err error) {
	return &pb.GetSeqResp{SeqId: uint64(time.Now().Unix())}, nil
}

func main() {
	listener, err := net.Listen("tcp", listen_addr)
	if err != nil {
		log.Fatalf("Listen Error, err=%v", err)
	}
	log.Println("Seq Server Start, Listen addr=%s", listen_addr)

	s := grpc.NewServer()
	svr := newServer()
	pb.RegisterSeqServiceServer(s, svr)
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Serve Error, err=%v")
	}
}
