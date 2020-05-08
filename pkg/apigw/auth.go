package apigw

import (
	"antalk-go/pkg/proto"
	"context"
)

type AuthService int

func (t *AuthService) Auth(ctx context.Context, req *proto.AuthReq, resp *proto.AuthResp) error {
	return nil
}
