package db

import (
	"fmt"
	"iceforg/pkg/config"
	"sync"

	rPkg "github.com/go-redis/redis"
)

var redisProvider = new(redis)

type redis struct {
	config *config.Redis
	client *rPkg.Client
	locker sync.Mutex
}

func GetRedisProvider() *rPkg.Client {
	if redisProvider.client == nil {
		redisProvider.initer()
	}
	return redisProvider.client
}

func (r *redis) initer() {
	if r.client == nil {
		r.locker.Lock()
		defer r.locker.Unlock()

		redisProvider.client = rPkg.NewClient(r.redisOpt())
		pong, err := redisProvider.client.Ping().Result()
		if err != nil || pong != r.config.Pong {
			panic("redis init failed," +
				fmt.Sprintf("err:%v,pong:%v", err, pong))
		}
	}
}

func (r *redis) redisOpt() *rPkg.Options {
	return &rPkg.Options{
		Addr:     r.config.Addr,
		Password: r.config.Password,
		DB:       0,
	}
}
