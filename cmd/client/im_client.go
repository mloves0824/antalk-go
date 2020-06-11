package main

import (
	"antalk-go/internal/proto/pb"
	"context"
	"github.com/smallnest/rpcx/v5/client"
	"github.com/smallnest/rpcx/v5/protocol"
	"log"
)

func main() {
	option := client.DefaultOption
	option.SerializeType = protocol.ProtoBuffer
	//获取服务发现句柄
	d := client.NewPeer2PeerDiscovery("tcp@localhost:18000", "")
	//获取xclient
	xclient := client.NewXClient("CmdService", client.Failtry, client.RandomSelect, d, option)
	if xclient == nil {
		log.Fatal("NewXClient error")
	}
	defer xclient.Close()

	req := &proto.CmdReq{
	}
	meta := &proto.CmdMeta{}
	meta.CmdType = proto.CmdType_CMD_LOGIN
	meta.ReqId = 1
	meta.DataType = proto.DataType_DT_PB
	req.Meta = meta
	loginReq := &proto.LoginReq{}
	loginReq.UserId = 1
	loginReq.Password = "123456"
	data, err := loginReq.Marshal()
	if err != nil {
		log.Fatal("loginReq.Marshal error, ", err)
	}
	req.Data = data

	resp :=&proto.CmdResp{}
	err = xclient.Call(context.Background(), "Cmd", req, resp)
	if err != nil {
		log.Fatal("call error, ", err)
	}

	log.Printf("call succ, %v", resp)


}
