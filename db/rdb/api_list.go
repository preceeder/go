package rdb

import (
	"github.com/preceeder/go/base"
	"github.com/redis/go-redis/v9"
)

// LINDEX key index, 用于获取列表中指定索引位置上的元素
func (rdm *RedisClient) LIndex(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, LINDEX, args, includeArgs...)
}

// LINSERT key BEFORE|AFTER pivot value , 将值 value 插入到列表 key 当中，位于值 pivot 之前或之后,
// 在列表的元素前或者后插入元素,当指定元素不存在于列表中时，不执行任何操作
// LINSERT mylist BEFORE "World" "There"
func (rdm *RedisClient) LInsert(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, LINSERT, args, includeArgs...)
}

// LLEN mylist , 获取列表中元素数量
func (rdm *RedisClient) LLen(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, LLEN, args, includeArgs...)
}

// LPUSH mylist value [value2 ...] , 将一个或多个值插入到列表头部, 如果 key 不存在，一个空列表会被创建并执行 LPUSH 操作
func (rdm *RedisClient) LPush(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, LPUSH, args, includeArgs...)
}

// LPUSHX mylist value [value2 ...] , 将一个或多个值插入到列表头部, 列表不存在时操作无效
func (rdm *RedisClient) LPushx(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, LPUSHX, args, includeArgs...)
}

// LPOP mylist , 移出并获取列表的第一个元素
func (rdm *RedisClient) LPop(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, LPOP, args, includeArgs...)
}

// LRANGE mylist start stop, 获取列表指定范围内的元素
// 其中 0 表示列表的第一个元素， 1 表示列表的第二个元素，以此类推。 你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。
func (rdm *RedisClient) LRange(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, LRANGE, args, includeArgs...)
}

// LREM key count value, 根据参数 COUNT 的值，移除列表中与参数 VALUE 相等的元素。
// count表示要删除的元素数量
// count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
// count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
// count = 0 : 移除表中所有与 VALUE 相等的值。
// return 被移除元素的数量。 列表不存在时返回 0 。
func (rdm *RedisClient) LRem(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, LREM, args, includeArgs...)
}

// LSET key index value,  当索引参数超出范围，或对一个空列表进行 LSET 时，返回一个错误。
func (rdm *RedisClient) LSet(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, LSET, args, includeArgs...)
}

// LTRIM key start stop, 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素([start, stop])，不在指定区间之内的元素都将被删除。
// 下标 0 表示列表的第一个元素，以 1 表示列表的第二个元素，以此类推。 你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。
func (rdm *RedisClient) LTrim(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, LTRIM, args, includeArgs...)
}

// RPOP key, 移除列表的最后一个元素，返回值为移除的元素。
func (rdm *RedisClient) RPop(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, RPOP, args, includeArgs...)
}

// RPOPLPUSH source target, 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
// return 返回这个元素
func (rdm *RedisClient) RPopLPush(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, RPOPLPUSH, args, includeArgs...)
}

// RPUSH key value [value2 ...], 在列表中添加一个或多个值到列表尾部
// return 执行 RPUSH 操作后，列表的长度。
func (rdm *RedisClient) RPush(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, RPUSH, args, includeArgs...)
}

// RPUSHX key value [value2 ...], 将值插入到已存在的列表尾部(最右边)。如果列表不存在，操作无效。
// return 执行 Rpushx 操作后，列表的长度。
func (rdm *RedisClient) RPushx(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, RPUSHX, args, includeArgs...)
}
