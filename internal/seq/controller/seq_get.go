package controller

type GetSeqReq struct {
	Uid string
}

type GetSeqResp struct {
	SeqID uint64
}