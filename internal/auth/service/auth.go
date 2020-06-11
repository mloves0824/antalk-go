package service

import (
	"antalk-go/internal/auth/config"
	"antalk-go/internal/auth/dao"
	proto "antalk-go/internal/proto/pb"
	"context"
	_ "github.com/go-sql-driver/mysql"
)

type AuthService struct {
	c   *config.Config
	dao *dao.Dao
}

func New(c *config.Config) *AuthService {
	s := &AuthService{
		c:   c,
		dao: dao.New(c),
	}
	return s
}

func (t *AuthService) Login(ctx context.Context, req *proto.CmdReq, resp *proto.CmdResp) error {
	loginReq := &proto.LoginReq{}
	respMeta := &proto.CmdMeta{}
	if err := loginReq.Unmarshal(req.Data); err != nil {
		respMeta.ErrorCode = proto.ErrorCode_ERROR_INTERNAL
		resp.Meta = respMeta
		return nil
	}

	//根据userid获取用户信息
	user, err := t.dao.Get(loginReq.UserId)
	if err != nil {
		respMeta.ErrorCode = proto.ErrorCode_ERROR_INTERNAL
		resp.Meta = respMeta
		return nil
	}

	if user.Password != loginReq.Password {
		respMeta.ErrorCode = proto.ErrorCode_ERROR_PASSWORD
		resp.Meta = respMeta
		return nil
	}

	respMeta.ErrorCode = proto.ErrorCode_ERROR_NONE
	resp.Meta = respMeta
	return nil
}
