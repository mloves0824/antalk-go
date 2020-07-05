package dao

import (
	"antalk-go/internal/common"
	"antalk-go/internal/push/model"
	_ "github.com/go-sql-driver/mysql"
	redis "github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

type Dao struct {
	c           *common.Config
	redisClient *redis.Pool
}

func New(c *common.Config) *Dao {
	d := &Dao{
		c:           c,
		redisClient: newPool(c.GetString("redis.addr")),
	}
	return d
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

func (d *Dao) GetUserRoute(uid int64, sid string) (*model.UserRouteInfo, error)  {
	conn := d.redisClient.Get()
	defer conn.Close()

	routeStr, err := conn.Do("hget", strconv.FormatInt(uid, 10), sid)
	if err != nil {
		return nil, err
	}

	route := &model.UserRouteInfo{}
	route.ServerInfo = string(routeStr)
	return route, nil

}