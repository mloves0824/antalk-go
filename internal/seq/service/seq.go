package service

import (
	"antalk-go/internal/seq/dao"
	"errors"
	"sync"
)

type SeqService struct {
	d *dao.SeqDao
	m *sync.Map
}

type seqInfo struct {
	curSeq uint64
	maxSeq uint64
}

var (
	ErrEtcdLoad = errors.New("etcd load error")
)

const step = 100 //TODO chenbang: etcd

func New() (*SeqService, error) {
	s := &SeqService{}
	return s, nil
}

func (s *SeqService) Get(uid string) (uint64, error) {
	var curSeq, maxSeq uint64 = 0, 0
	v, ok := s.m.Load(uid)
	if !ok {
		curSeq = 0
	} else {
		info, ok := v.(seqInfo)
		if !ok {
			curSeq = 0
		}
		curSeq = info.curSeq
	}

	go func() {
		newCurSeq := curSeq + 1
		//当前序号大于最大序号，则提升一个步长，并持久化max_seq
		if newCurSeq > maxSeq {
			maxSeq += step
		}

		s.d.Set(uid, string(maxSeq))
	}()

	return curSeq, nil
}
