package main

import (
	"antalk-go/pkg/proto"
	"context"
	"github.com/smallnest/rpcx/v5/client"
	"github.com/smallnest/rpcx/v5/protocol"
	"log"
)

func main() {
	option := client.DefaultOption
	option.SerializeType = protocol.ProtoBuffer
	//获取服务发现句柄
	d := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:11111", "")
	//获取xclient
	xclient := client.NewXClient("AuthService", client.Failtry, client.RandomSelect, d, option)
	if xclient == nil {
		log.Fatal("NewXClient error")
	}
	defer xclient.Close()

	req := &proto.AuthReq{
		Uid:    "test",
		Method: 0,
		Token:  "test",
	}
	resp :=&proto.AuthResp{}
	err := xclient.Call(context.Background(), "Auth", req, resp)
	if err != nil {
		log.Fatal("call error, ", err)
	}

	log.Printf("call succ, %v", resp)


}
