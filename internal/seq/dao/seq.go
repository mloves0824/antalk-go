package dao

import (
	"context"
	"go.etcd.io/etcd/client"
)

type SeqDao struct {
	c client.Client
}

func (d *SeqDao) Get(key string) (string, error) {
	kapi := client.NewKeysAPI(d.c)
	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		return "", err
	}

	return resp.Node.Value, nil
}

func (d *SeqDao) Set(key string, value string) error  {
	kapi := client.NewKeysAPI(d.c)
	_, err := kapi.Set(context.Background(), key, value, nil)
	if err != nil {
		return err
	}

	return nil
}