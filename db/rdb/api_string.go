package rdb

import (
	"github.com/preceeder/go/base"
	"github.com/redis/go-redis/v9"
)

func (rdm *RedisClient) Set(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SET, args, includeArgs...)
}

func (rdm *RedisClient) MSet(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, MSET, args, includeArgs...)
}

// SETRANGE key offset value   , 用 value 参数覆写给定 key 所储存的字符串值，从偏移量 offset 开始。
func (rdm *RedisClient) SetRange(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SETRANGE, args, includeArgs...)
}

// 将值 value 关联到 key ，并将 key 的过期时间设为 seconds (以秒为单位)。
func (rdm *RedisClient) SetEx(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SETEX, args, includeArgs...)
}

// 只有在 key 不存在时设置 key 的值。
func (rdm *RedisClient) SetNx(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SETNX, args, includeArgs...)
}

func (rdm *RedisClient) Del(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, DEL, args, includeArgs...)
}

func (rdm *RedisClient) Get(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, GET, args, includeArgs...)
}

func (rdm *RedisClient) GetSet(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, GETSET, args, includeArgs...)
}

func (rdm *RedisClient) GetRange(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, GETRANGE, args, includeArgs...)
}

func (rdm *RedisClient) MGet(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, MGET, args, includeArgs...)
}

func (rdm *RedisClient) Incr(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, INCR, args, includeArgs...)
}

func (rdm *RedisClient) IncrBy(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, INCRBY, args, includeArgs...)
}
func (rdm *RedisClient) IncrByFloat(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, INCRBYFLOAT, args, includeArgs...)
}

func (rdm *RedisClient) DecrBy(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, DECRBY, args, includeArgs...)
}

func (rdm *RedisClient) Decr(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, DECR, args, includeArgs...)
}

func (rdm *RedisClient) StringAppend(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, APPEND, args, includeArgs...)
}
