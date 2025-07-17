package rdb

import (
	"github.com/preceeder/go/base"
	"github.com/redis/go-redis/v9"
)

//	SADD key member [member ...], 向集合添加一个或多个成员
//
// return 被添加到集合中的新元素的数量，不包括被忽略的元素。
func (rdm *RedisClient) SAdd(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SADD, args, includeArgs...)
}

// SCARD key, 获取集合的成员数
// return 集合的数量。 当集合 key 不存在时，返回 0 。
func (rdm *RedisClient) SCard(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SCARD, args, includeArgs...)
}

// SDIFF key, 计算第一个集合与其他集合的差集
// return 返回第一个集合与其他集合之间的差异，也可以认为说第一个集合中独有的元素
// key1 = {a,b,c,d}
// key2 = {c}
// key3 = {a,c,e}
// SDIFF key1 key2 key3 = {b,d}
func (rdm *RedisClient) SDiff(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SDIFF, args, includeArgs...)
}

// SDIFFSTORE destination key [key …] ,给定所有集合的差集并存储在 destination 中, 如果指定的集合 destination 已存在，则会被覆盖。
// return 结果集中的元素数量。
func (rdm *RedisClient) SDiffStore(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SDIFFSTORE, args, includeArgs...)
}

// SINTER key key1  ...keyn  , 返回给定所有给定集合的交集。 不存在的集合 key 被视为空集。 当给定集合当中有一个空集时，结果也为空集(根据集合运算定律)。
// return 交集的集合
func (rdm *RedisClient) SInter(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SINTER, args, includeArgs...)
}

// SINTERSTORE destination key key1 ...,  将给定集合之间的交集存储在指定的集合中。如果指定的集合已经存在，则将其覆盖。
// return 返回存储交集的集合的元素数量。
func (rdm *RedisClient) SInterStore(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SINTERSTORE, args, includeArgs...)
}

// SISMEMBER key member ,  判断member是否存在于key对应的集合中
// return 如果成员元素是集合的成员，返回 1 。 如果成员元素不是集合的成员，或 key 不存在，返回 0 。
func (rdm *RedisClient) SIsMember(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SISMEMBER, args, includeArgs...)
}

// SMEMBERS key, 返回集合中的所有的成员。 不存在的集合 key 被视为空集合。
// return 集合中的所有成员。
func (rdm *RedisClient) SMembers(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SMEMBERS, args, includeArgs...)
}

// SMOVE  source destination member, 将指定成员 member 元素从 source 集合移动到 destination 集合。
// 如果 source 集合不存在或不包含指定的 member 元素，则 SMOVE 命令不执行任何操作，仅返回 0 。否则， member 元素从 source 集合中被移除，并添加到 destination 集合中去。
// 当 destination 集合已经包含 member 元素时， SMOVE 命令只是简单地将 source 集合中的 member 元素删除。
// 当 source 或 destination 不是集合类型时，返回一个错误。
// return 如果成员元素被成功移除，返回 1 。 如果成员元素不是 source 集合的成员，并且没有任何操作对 destination 集合执行，那么返回 0 。
func (rdm *RedisClient) SMove(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SMOVE, args, includeArgs...)
}

// SREM key member1 member2 ... , 移除集合中的一个或多个成员元素，不存在的成员元素会被忽略。
// return 被成功移除的元素的数量，不包括被忽略的元素。
func (rdm *RedisClient) SRem(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SREM, args, includeArgs...)
}

// SUNION key key1 key2 ..., 计算给定集合的并集。不存在的集合 key 被视为空集。
// return 并集成员的列表。
func (rdm *RedisClient) SUnion(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SUNION, args, includeArgs...)
}

// SUNIONSTORE destination key [key …], 将给定集合的并集存储在指定的集合 destination 中。如果 destination 已经存在，则将其覆盖。
// return 结果集中的元素数量。
func (rdm *RedisClient) SUnionStore(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, SUNIONSTORE, args, includeArgs...)
}
