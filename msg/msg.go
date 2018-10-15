package main

import (
	"fmt"
	//"io"
	"log"
	"net"

	//"strings"
	//"time"

	apigw_pb "github.com/mloves0824/antalk-go/proto/apigw"
	common_pb "github.com/mloves0824/antalk-go/proto/common"
	data_pb "github.com/mloves0824/antalk-go/proto/data"
	seq_pb "github.com/mloves0824/antalk-go/proto/seq"

	pb "github.com/mloves0824/antalk-go/proto/msg"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	data_addr  = "localhost:50052"
	seq_addr   = "localhost:50055"
	apigw_addr = "localhost:50051"
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

/*
1. check param
2. get seqid from seqsvr
3. set msg to datasvr
4. resp to msg sender
5. notify msg to receiver
*/
func (s *server) MsgSend(ctx context.Context, req *pb.MsgSendReq) (*pb.MsgSendResp, error) {
	log.Printf("Recv MsgSend Request, %v", req.GetMsgInfo())

	seq_req := seq_pb.GetSeqReq{Uid: req.GetMsgInfo().GetRecvUid()}
	seq_resp, err := s.seq_client.GetSeq(context.Background(), &seq_req)
	if err != nil {
		log.Printf("GetSeqReq Rpc Error, %v", err)
		return &pb.MsgSendResp{Result: common_pb.ResultType_ResultErrInner}, nil
	}
	if seq_resp.GetSeqId() <= 0 {
		log.Printf("GetSeqReq Logic Error, seqid=%d", seq_resp.GetSeqId())
		return &pb.MsgSendResp{Result: common_pb.ResultType_ResultErrInner}, nil
	}
	log.Printf("GetSeqReq Logic Success, seqid=%lu", seq_resp.GetSeqId())

	msg_info := req.GetMsgInfo()
	msg_info.MsgId = seq_resp.GetSeqId()
	data_req := data_pb.SaveMsgReq{MsgInfo: msg_info}
	data_resp, err := s.data_client.SaveMsg(context.Background(), &data_req)
	if err != nil {
		log.Printf("SaveMsgReq Rpc Error, %v", err)
		return &pb.MsgSendResp{Result: common_pb.ResultType_ResultErrInner}, nil
	}
	if data_resp.GetResult() != common_pb.ResultType_ResultOK {
		log.Printf("SaveMsgReq Logic Error, result=%d", data_resp.GetResult())
		return &pb.MsgSendResp{Result: data_resp.GetResult()}, nil
	}

	go func() {
		//get session info from datasvr
		recv_uid := req.GetMsgInfo().GetRecvUid()
		getsession_req := data_pb.GetSessionReq{Uid: recv_uid}
		getsession_resp, err := s.data_client.GetSession(context.Background(), &getsession_req)
		if err != nil {
			log.Printf("Get session Rpc error, uid=%s", recv_uid)
			return
		}
		recv_serverinfo := getsession_resp.GetServerInfo()
		log.Println("GetSessionReq Rpc Success, uid=%s, serverinfo=%s", recv_uid, recv_serverinfo)
		if recv_serverinfo == "" { //TODO: push to apns or jpush?
			log.Printf("Recv ServerInfo Is Empty, uid=%s", recv_uid)
			return
		}

		//TODO: The Apigw Is Not Exist.
		//push to apigw
		apigw_conn, err := grpc.Dial(recv_serverinfo, grpc.WithInsecure())
		if err != nil {
			log.Printf("Dial Apigw Error, uid=%s, recv_serverinfo=%s", recv_uid, recv_serverinfo)
			return
		}
		msgpush_req := apigw_pb.MsgPushReq{MsgInfo: msg_info}
		msgpush_resp, err := apigw_pb.NewApigwServiceClient(apigw_conn).MsgPush(context.Background(), &msgpush_req)
		if err != nil {
			log.Printf("MsgPushReq RPC Error, uid=%s, recv_serverinfo=%s", recv_uid, recv_serverinfo)
			return
		}
		log.Printf("MsgPushReq RPC Success, uid=%s, recv_serverinfo=%s, result=%d", recv_uid, recv_serverinfo, msgpush_resp.GetResult())
		//TODO: set msg acked and readed.
	}()

	return &pb.MsgSendResp{Result: common_pb.ResultType_ResultOK}, nil
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
