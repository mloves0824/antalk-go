package apigw

import (
	"antalk-go/internal/common"
	"github.com/smallnest/rpcx/v5/client"
	"github.com/smallnest/rpcx/v5/protocol"
	"strings"
)

type AuthClient struct {
	client client.XClient
}

func (ac *AuthClient) Init(c *common.Config)  {
	servers := strings.Split(c.GetString("auth.servers"), ";")
	var pairs []*client.KVPair
	for _,v := range servers {
		pairs = append(pairs, &client.KVPair{Key: v})
	}

	d := client.NewMultipleServersDiscovery(pairs)
	opt := client.DefaultOption
	opt.SerializeType = protocol.ProtoBuffer
	ac.client = client.NewXClient("RpcxController", client.Failtry, client.RandomSelect, d, opt)
}