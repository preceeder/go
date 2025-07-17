package rdb

import (
	"time"
)

var UserWealthCmd = RdCmd{
	Key: "redis_tessd009",
	CMD: map[Command]RdSubCmd{
		"HSET": {
			Params: "",
			Exp: func() time.Duration {
				return time.Second * 30
			},
		},
		"HGET": {
			Params: "",
		},
		"expire": {
			Params: "{{expireTime}}",
			DefaultParams: map[string]any{
				"expireTime": time.Hour * 49,
			},
		},
		"HMGET": {
			Params: "",
		},
		"HGETALL": {
			Params: "",
		},
		"ZRANK": {
			ReturnNilError: true,
		},
	},
}

func InitRedis() *RedisClient {
	config := Config{
		Host:        "127.0.0.1",
		Port:        "6377",
		Password:    "QDjk9UdkoD6cv",
		Db:          13,
		MaxIdle:     2,
		IdleTimeout: 240,
		PoolSize:    13,
	}
	return NewRedisClient(config)
}
