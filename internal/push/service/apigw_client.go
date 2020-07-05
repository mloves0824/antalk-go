package service

import (
	"antalk-go/internal/common"
	proto "antalk-go/internal/proto/pb"
	"context"
	"github.com/smallnest/rpcx/v5/client"
	"github.com/smallnest/rpcx/v5/protocol"
	"log"
	"strings"
)

type ApigwClient struct {
	client client.XClient
}

func (ac *ApigwClient) Init(c *common.Config)  {
	servers := strings.Split(c.GetString("apigw.servers"), ";")
	var pairs []*client.KVPair
	for _,v := range servers {
		pairs = append(pairs, &client.KVPair{Key: v})
	}

	d := client.NewMultipleServersDiscovery(pairs)
	opt := client.DefaultOption
	opt.SerializeType = protocol.ProtoBuffer
	ac.client = client.NewXClient("RpcxController", client.Failtry, client.RandomSelect, d, opt)
}

func (ac *ApigwClient) SendMsgNotifyToApigw(addr string, req *proto.MsgNotifyReq) error {
	d := client.NewPeer2PeerDiscovery(addr, "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.ProtoBuffer
	xclient := client.NewXClient("CmdService", client.Failtry, client.RandomSelect, d, opt)

	cmdReq := &proto.CmdReq{
	}
	meta := &proto.CmdMeta{}
	meta.CmdType = proto.CmdType_CMD_MSG_NOTIFY
	meta.ReqId = 1
	meta.DataType = proto.DataType_DT_PB
	cmdReq.Meta = meta

	data, err := req.Marshal()
	if err != nil {
		log.Fatal("loginReq.Marshal error, ", err)
	}
	cmdReq.Data = data
	resp :=&proto.CmdResp{}
	err = xclient.Call(context.Background(), "Cmd", req, resp)
	if err != nil {
		log.Fatal("call error, ", err)
	}
}