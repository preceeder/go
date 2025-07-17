package rdb

import (
	"github.com/preceeder/go/base"
	"github.com/redis/go-redis/v9"
)

// HSET key field value
func (rdm *RedisClient) HSet(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HSET, args, includeArgs...)
}

// HGET key field
func (rdm *RedisClient) HGet(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HGET, args, includeArgs...)
}

// HDEL key field [field2 ...], 删除字段，可以同时删除多个
func (rdm *RedisClient) HDel(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HDEL, args, includeArgs...)
}

// HGETALL key
func (rdm *RedisClient) HGetAll(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HGETALL, args, includeArgs...)
}

// HMSET key field1 value1 field2 value2
func (rdm *RedisClient) HMSet(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HMSET, args, includeArgs...)
}

// HMGET key field1  field2
func (rdm *RedisClient) HMGet(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HMGET, args, includeArgs...)
}

// HSETNX key field value , 设置键下字段的值，存在则不操作返回0，不存在并创建成功则返回1
func (rdm *RedisClient) HSetNx(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HSETNX, args, includeArgs...)
}

// HINCRBY key field1  value   , 指定键指定字段自增指定的整数
func (rdm *RedisClient) HIncrBy(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HINCRBY, args, includeArgs...)
}

// HINCRBYFLOAT key field1  value   , 指定键指定字段自增指定的浮点数
func (rdm *RedisClient) HIncrByFloat(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HINCRBYFLOAT, args, includeArgs...)
}

// HKEYS key  , 获取键下的所有字段列表
func (rdm *RedisClient) HKeys(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HKEYS, args, includeArgs...)
}

// HLEN key  , 获取键下字段的数量
func (rdm *RedisClient) HLen(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HLEN, args, includeArgs...)
}

// HVALS key  , 返回哈希表所有的值
func (rdm *RedisClient) HVals(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HVALS, args, includeArgs...)
}

// HEXISTS key field, 键下是否存在指定的字段
func (rdm *RedisClient) HExists(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, HEXISTS, args, includeArgs...)
}
