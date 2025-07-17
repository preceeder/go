package rdb

import (
	"context"
	"errors"
	"github.com/preceeder/go/base"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"time"
)

type Config struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Password    string `json:"password"`
	Db          int    `json:"db"`
	MaxIdle     int    `json:"maxIdle"`
	MinIdle     int    `json:"minIdle"`
	IdleTimeout int    `json:"idleTimeout"`
	PoolSize    int    `json:"PoolSize"`
}

type RedisClient struct {
	Config Config
	Client *redis.Client
}

func NewRedisClient(config Config) *RedisClient {
	return &RedisClient{Client: initRedis(config), Config: config}
}

func initRedis(c Config) *redis.Client {
	//Rd := new(redis.Client)
	slog.Info("redisDb connect", "info", c)
	addr := c.Host + ":" + c.Port
	redisOpt := &redis.Options{
		Addr:         addr,
		Password:     c.Password,
		DB:           c.Db,
		PoolSize:     c.PoolSize,
		MaxIdleConns: c.MaxIdle,
		MinIdleConns: c.MinIdle,
	}
	rdb := redis.NewClient(redisOpt)
	rdb.AddHook(RKParesHook{})
	cmd := rdb.Ping(context.Background())
	if cmd.Err() != nil {
		panic("redis connect fail, " + cmd.Err().Error())
	}

	return rdb
}

func (rdm *RedisClient) RedisClose() {
	err := rdm.Client.Close()
	if err != nil {
		slog.Error("close redisDb", "index", rdm.Config.Db, "error", err.Error())
	} else {
		slog.Info("close redisDb", "index", rdm.Config.Db)
	}

}

func (rdm *RedisClient) Handler(ctx base.BaseContext, cmd RdCmd, cmdName Command, args map[string]any, includeArgs ...any) *redis.Cmd {
	cmdList, key, subCmd := Build(ctx, cmd, cmdName, args, includeArgs...)
	resultCmd := rdm.Client.Do(ctx, cmdList...)
	// 如果是 nil 且不返回的话就处理一下
	if !subCmd.ReturnNilError {
		if errors.Is(resultCmd.Err(), redis.Nil) {
			resultCmd.SetErr(nil)
		}
	}
	rdm.setExp(ctx, key, subCmd)
	return resultCmd
}
func (rdm *RedisClient) setExp(ctx base.BaseContext, key string, subCmd RdSubCmd) {
	if subCmd.Exp != nil {
		exp := subCmd.Exp()
		expireCmd := rdm.Client.Expire(ctx, key, exp)
		if expireCmd.Err() != nil {
			slog.ErrorContext(ctx, "set expire", "key", key, "error", expireCmd.Err())
		}
	}
}

// Exec 执行 Redis 命令
func Exec(ctx context.Context, rdb *redis.Client, cmd redis.Cmder, key string, exp time.Duration) error {
	// 使用 go-redis 执行命令
	err := rdb.Process(ctx, cmd) // 这一步可以单独执行
	if err == nil && exp > 0 {
		if er := rdb.Expire(ctx, key, exp).Err(); er != nil {
			slog.ErrorContext(ctx, "set exp", "key", key, "exp", exp, "error", er.Error())
		}
	}
	return err
}
