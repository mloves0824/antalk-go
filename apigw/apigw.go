package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	pb "github.com/mloves0824/antalk-go/proto/apigw"
	auth_pb "github.com/mloves0824/antalk-go/proto/auth"
	common_pb "github.com/mloves0824/antalk-go/proto/common"
	data_pb "github.com/mloves0824/antalk-go/proto/data"
	msg_pb "github.com/mloves0824/antalk-go/proto/msg"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	auth_addr = "localhost:50053"
	data_addr = "localhost:50052"
	msg_addr  = "localhost:50054"
)

type server struct {
	name        string
	subStrm     map[string]pb.ApigwService_SubscribeServer
	auth_client auth_pb.AuthServiceClient
	data_client data_pb.DataServiceClient
	msg_client  msg_pb.MsgServiceClient
}

func newServer() *server {
	// Setup a connection with the authserver
	auth_conn, err := grpc.Dial(auth_addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect auth: %v", err)
	}
	new_auth_client := auth_pb.NewAuthServiceClient(auth_conn)

	// Setup a connection with the dataserver
	data_conn, err := grpc.Dial(data_addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect data: %v", err)
	}
	new_data_client := data_pb.NewDataServiceClient(data_conn)

	// Setup a connection with the msgserver
	msg_conn, err := grpc.Dial(msg_addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect msg: %v", err)
	}
	new_msg_client := msg_pb.NewMsgServiceClient(msg_conn)

	return &server{name: "ApigwServer",
		subStrm:     make(map[string]pb.ApigwService_SubscribeServer),
		auth_client: new_auth_client,
		data_client: new_data_client,
		msg_client:  new_msg_client}
}

/*
	1.check param
	2.check auth from authsvr
	3.set session to datasvr
	4.resp to client
*/
func (s *server) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	log.Printf("Received a Login Request uid=(%s),token=(%s)", req.GetUid(), req.GetToken())

	if req.GetUid() == "" || req.GetToken() == "" {
		log.Printf("Param Error!")
		return &pb.LoginResp{Uid: req.GetUid(), ResultCode: common_pb.ResultType_ResultErrInner}, nil
	}

	auth_req := &auth_pb.CheckAuthReq{Uid: req.GetUid(), Token: req.GetToken()}
	auth_resp, err := s.auth_client.CheckAuth(context.Background(), auth_req)
	if err != nil {
		log.Printf("CheckAuth RPC Error, err=%v", err)
		return &pb.LoginResp{Uid: req.GetUid(), ResultCode: common_pb.ResultType_ResultErrInner}, nil
	}
	if auth_resp.GetResult() != common_pb.ResultType_ResultOK {
		log.Printf("CheckAuth Logic Error, result=%v", auth_resp.GetResult())
		return &pb.LoginResp{Uid: req.GetUid(), ResultCode: auth_resp.GetResult()}, nil
	}

	setsession_req := data_pb.SetSessionReq{Uid: req.GetUid(), ServerInfo: "localhost:50051"}
	setsession_resp, err := s.data_client.SetSession(context.Background(), &setsession_req)
	if err != nil {
		log.Printf("SetSessionReq RPC Error, err=%v", err)
		return &pb.LoginResp{Uid: req.GetUid(), ResultCode: common_pb.ResultType_ResultErrInner}, nil
	}
	if setsession_resp.GetResult() != common_pb.ResultType_ResultOK {
		log.Printf("SetSessionReq Logic Error, result=%v", setsession_resp.GetResult())
		return &pb.LoginResp{Uid: req.GetUid(), ResultCode: setsession_resp.GetResult()}, nil
	}

	return &pb.LoginResp{Uid: req.GetUid(), ResultCode: common_pb.ResultType_ResultOK}, nil
}

func (s *server) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutResp, error) {
	log.Printf("Received a Logout Request (%s)", req.GetUid())
	return &pb.LogoutResp{}, nil
}

func (s *server) Subscribe(stream pb.ApigwService_SubscribeServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("Received a Subscription Request (%s, %s)", in.GetUid(), in.GetType().String())
		_, ok := s.subStrm[in.GetUid()+":"+in.GetType().String()]
		if !ok {
			s.subStrm[in.GetUid()+":"+in.GetType().String()] = stream
		} else {
			log.Printf("Client %s already subscribed to %s", in.GetUid(), in.GetType().String())
		}
	}
	return nil
}

func (s *server) MsgSend(ctx context.Context, req *pb.MsgSendReq) (*pb.MsgSendResp, error) {
	log.Printf("Received a MsgSend Request (%v)", req.GetMsgInfo())

	//TODO: check auth locally before sending msg

	msg_req := msg_pb.MsgSendReq{MsgInfo: req.GetMsgInfo()}
	msg_resp, err := s.msg_client.MsgSend(context.Background(), &msg_req)
	if err != nil {
		log.Printf("MsgSend Rpc Error, %v", err)
		return &pb.MsgSendResp{Result: common_pb.ResultType_ResultErrInner}, nil
	}
	log.Printf("MsgSend Rpc Success, result=%s", msg_resp.GetResult())
	return &pb.MsgSendResp{Result: msg_resp.GetResult()}, nil
}

func (s *server) MsgPush(ctx context.Context, req *pb.MsgPushReq) (*pb.MsgPushResp, error) {
	log.Printf("Received a MsgPush Request (%s)", req.GetMsgInfo().GetMsgId())

	recv_uid := req.GetMsgInfo().GetRecvUid()
	value, ok := s.subStrm[recv_uid+":MSG"]
	if ok {
		if err := value.Send(&pb.Notification{Type: pb.TopicType_MSG, Msg: &pb.MsgNotify{MsgInfo: req.GetMsgInfo()}}); err != nil {
			log.Fatalf("Send failed %v", err)
		}
	}
	return &pb.MsgPushResp{}, nil
}

func (s *server) KickoutPush(ctx context.Context, req *pb.KickoutPushReq) (*pb.KickoutPushResp, error) {
	return &pb.KickoutPushResp{}, nil
}

func (s *server) HealthCheck(ctx context.Context, req *pb.HealthCheckReq) (*pb.HealthCheckResp, error) {

	setsession_req := data_pb.SetSessionReq{Uid: req.GetUid(), ServerInfo: "localhost:50051"}
	setsession_resp, err := s.data_client.SetSession(context.Background(), &setsession_req)
	if err != nil {
		log.Printf("SetSessionReq RPC Error, err=%v", err)
		return &pb.HealthCheckResp{Status: pb.HealthCheckResp_NOT_SERVING}, nil
	}
	if setsession_resp.GetResult() != common_pb.ResultType_ResultOK {
		log.Printf("SetSessionReq Logic Error, result=%v", setsession_resp.GetResult())
		return &pb.HealthCheckResp{Status: pb.HealthCheckResp_NOT_SERVING}, nil
	}

	return &pb.HealthCheckResp{Status: pb.HealthCheckResp_SERVING}, nil
}

func (s *server) pushUpdates() {
	var str string
	//var i uint32

	time.Sleep(1 * time.Second)

	for {
		//fmt.Println("walk thru the subscribe streams")
		for k, v := range s.subStrm {
			if strings.Contains(k, "KICKOUT") {
				fmt.Printf("Enter new KICKOUT for Client(%v): ", k)
				fmt.Scan(&str)
				if err := v.Send(&pb.Notification{Type: pb.TopicType_KICKOUT, Kick: &pb.KickNotify{Uid: k, Reason: 0}}); err != nil {
					log.Fatalf("Send failed %v", err)
				}
			} else {
				fmt.Printf("Not Support!", k)
			}
		}
	}
}

func main() {
	fmt.Println("Server Start!")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	myServer := newServer()
	pb.RegisterApigwServiceServer(s, myServer)

	//go myServer.pushUpdates()

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
