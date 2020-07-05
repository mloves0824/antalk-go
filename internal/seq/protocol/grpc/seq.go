package grpc

import (
	pb "antalk-go/internal/proto/grpc"
	"antalk-go/internal/seq/controller"
	"context"
	"google.golang.org/grpc"
)

func New() {
	s := grpc.NewServer()
	pb.RegisterSeqServiceServer(s, &server{})
}

type server struct {
	controller *controller.SeqController
	pb.UnimplementedSeqServiceServer
}

func (s *server) GetSeq(c context.Context, r *pb.GetSeqReq) (*pb.GetSeqResp, error) {
	req := &controller.GetSeqReq{}
	bindReq(r, req)
	resp, err := s.controller.GetSeq(c, req)
	if err != nil {
		return nil, err
	}
	pbResp := &pb.GetSeqResp{}
	bindResp(resp, pbResp)
	return pbResp, nil
}

func bindReq(in *pb.GetSeqReq, out *controller.GetSeqReq)  {
	out.Uid = in.Uid
}

func bindResp(in *controller.GetSeqResp, out *pb.GetSeqResp)  {
	out.SeqId = in.SeqID
}