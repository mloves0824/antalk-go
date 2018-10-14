package main

import (
	"fmt"
	_ "io"
	"log"
	"net"

	"database/sql"
	_ "strings"
	"time"

	redigo "github.com/garyburd/redigo/redis"
	redis "github.com/mloves0824/antalk-go/pkg/utils/redis"
	commonpb "github.com/mloves0824/antalk-go/proto/common"
	pb "github.com/mloves0824/antalk-go/proto/data"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-xorm/xorm"
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
	db_client    *sql.DB
}

const (
	key_prefix_token = "TOKEN:"
)

var redis_conf redisConf = redisConf{"127.0.0.1:19000", "antalk-auth", 3 * time.Second}

func newServer() *server {

	client, err := redis.NewClientNoAuth(redis_conf.addr, redis_conf.timeout)
	if err != nil {
		//log.Fatalf("failed to NewClient: %v", err)
	}

	//init mysql
	new_db_client, err := sql.Open("mysql", "root:123456@tcp(139.199.161.198:3006)/antalk?charset=utf8")
	if err != nil {
		log.Fatalf("Failed to Conn to mysql: %v", err)
	}
	//	rows, err := db.Query("select * from tx")
	//	if err != nil {
	//		log.Fatalf("Failed to Query: %v", err)
	//	}
	//	for rows.Next() {
	//		var id int
	//		var num int
	//		err := rows.Scan(&id, &num)
	//		if err != nil {
	//			log.Fatalf("Failed to Scan: %v", err)
	//		}
	//	}
	return &server{name: "DataServer", redis_client: client, db_client: new_db_client}
}

func String(rsp interface{}) string {
	if s, err := redigo.String(rsp, nil); err != nil {
		panic(err)
	} else {
		return s
	}
}

func (s *server) CheckAuth(ctx context.Context, req *pb.CheckAuthReq) (*pb.CheckAuthResp, error) {
	cmd := fmt.Sprintf("get %s", key_prefix_token+req.GetUid())
	value, err := s.redis_client.Do("GET", key_prefix_token+req.GetUid())
	if err != nil {
		log.Printf("redis_client.Do Err, cmd=%s, err=%v", cmd, err)
		return &pb.CheckAuthResp{Uid: req.GetUid(), Result: commonpb.ResultType_ResultErrRedis}, nil
	}
	log.Printf("redis_client.Do success, cmd=%s, value=%s, GetToken=%s,compare=%b",
		cmd, String(value), req.GetToken(), (String(value) == req.GetToken()))

	if String(value) == req.GetToken() {
		return &pb.CheckAuthResp{Uid: req.GetUid(), Result: commonpb.ResultType_ResultOK}, nil
	}
	return &pb.CheckAuthResp{Uid: req.GetUid(), Result: commonpb.ResultType_ResultErrCheckAuth}, nil
}

func (s *server) SetSession(context.Context, *pb.SetSessionReq) (*pb.SetSessionResp, error) {
	return &pb.SetSessionResp{}, nil
}

func (s *server) GetSession(context.Context, *pb.GetSessionReq) (*pb.GetSessionResp, error) {
	return &pb.GetSessionResp{}, nil
}
func (s *server) SaveMsg(ctx context.Context, req *pb.SaveMsgReq) (*pb.SaveMsgResp, error) {
	stmt, err := s.db_client.Prepare("INSERT INTO Msg SET msgId=?,srcUid=?,dstUid=?,msgBody=?")
	if err != nil {
		log.Printf("Prepare error!")
		return &pb.SaveMsgResp{Result: commonpb.ResultType_ResultErrInner}, nil
	}
	result, err := stmt.Exec(req.GetMsgInfo().GetMsgId(), req.GetMsgInfo().GetSendUid(),
		req.GetMsgInfo().GetRecvUid(), req.GetMsgInfo().GetMsgBody())
	if err != nil {
		log.Printf("Exec error!")
		return &pb.SaveMsgResp{Result: commonpb.ResultType_ResultErrInner}, nil
	}
	affect, err := result.RowsAffected()
	log.Println("Exec Success,affect=%d,err=%v!", affect, err)
	if affect != 1 {
		log.Printf("RowsAffected error!")
		return &pb.SaveMsgResp{Result: commonpb.ResultType_ResultErrInner}, nil
	}

	return &pb.SaveMsgResp{Result: commonpb.ResultType_ResultOK}, nil
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
