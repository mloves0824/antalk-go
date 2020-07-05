package controller

import (
	"antalk-go/internal/seq/service"
	"context"
	"errors"
)

type SeqController struct {
	service *service.SeqService
}

func (s *SeqController) GetSeq(c context.Context, req *GetSeqReq) (*GetSeqResp, error) {
	if req.Uid == "" {
		return nil, errors.New("uid is empty")
	}

	reqId, err := s.service.Get(req.Uid)
	if err != nil {
		return nil, err
	}

	return &GetSeqResp{SeqID:reqId}, nil
}