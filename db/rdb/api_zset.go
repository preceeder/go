package rdb

import (
	"github.com/preceeder/go/base"
	"github.com/redis/go-redis/v9"
)

// ZADD key score1 member1 [score2 member2] , 向有序集合添加一个或多个成员，或者更新已存在成员的分数。
// return 被成功添加的新成员的数量，不包括那些被更新的、已经存在的成员。
func (rdm *RedisClient) ZAdd(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZADD, args, includeArgs...)
}

// ZCARD key , 获取有序集合的成员数
// return 当 key 存在且是有序集类型时，返回有序集的基数。 当 key 不存在时，返回 0 。
func (rdm *RedisClient) ZCard(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZCARD, args, includeArgs...)
}

// ZCOUNT key min max ,计算在有序集合中指定分数区间的成员数   [1,3]
// return  分数值在 min 和 max 之间的成员的数量。
func (rdm *RedisClient) ZCount(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZCOUNT, args, includeArgs...)
}

// ZINCRBY key increment member,  有序集合中对指定成员的分数加上增量 increment,
// 可以通过传递一个负数值 increment ，让分数减去相应的值，比如 ZINCRBY key -5 member ，就是让 member 的 score 值减去 5 。
// 当 key 不存在，或分数不是 key 的成员时， ZINCRBY key increment member 等同于 ZADD key increment member 。
// return member 成员的新分数值，以字符串形式表示。
func (rdm *RedisClient) ZIncrBy(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZINCRBY, args, includeArgs...)
}

// ZLEXCOUNT key min max,  计算有序集合中指定字典区间内成员数量， 分数相同的元素的顺序按照元素的字典序排列
// 排序规则总结：
//  1. 优先按 score 升序排列
//  2. score 相同时，按成员（member）的字典序排列
//     •	ASCII 字符小的排在前面（如 “a” < “b”）
//     •	数字字符排在字母前（如 “123” < “abc”）
//     •	大写 < 小写（“A” < “a”）
//
// 如： ZADD myzset 0 a 0 b 0 c 0 d 0 e 0 f 0 g
// ZLEXCOUNT myzset - + 获取所有的，   - 负无穷， + 正无穷， 结果 7
// ZLEXCOUNT myzset [b (f    获取包含b不包含f的之间的所有成员数 结果： 4
// return 指定区间内的成员数量。
func (rdm *RedisClient) ZLexCount(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZLEXCOUNT, args, includeArgs...)
}

// ZRANGE  key start stop [WITHSCORES] , 通过索引区间返回有序集合指定区间内的成员， 和下面的类似， 只是排序不一样
// 其中成员的位置按分数值递增(从小到大)来排序。
// 具有相同分数值的成员按字典序(lexicographical order )来排列。
// 下标参数 start 和 stop 都以 0 为底，也就是说，以 0 表示有序集第一个成员，以 1 表示有序集第二个成员，以此类推。
// 你也可以使用负数下标，以 -1 表示最后一个成员， -2 表示倒数第二个成员，以此类推。
// return 指定区间内，带有分数值(可选)的有序集成员的列表。
// [key1, score1, key2, score2, ...]
func (rdm *RedisClient) ZRange(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZRANGE, args, includeArgs...)
}

// ZREVRANGE key start stop [WITHSCORES], 通过索引区间返回有序集合指定区间内的成员, 和上面的类似， 只是排序不一样
// 其中成员的位置按分数值递减(从大到小)来排列。
// 具有相同分数值的成员按字典序的逆序(reverse lexicographical order)排列。
// return 指定区间内，带有分数值(可选)的有序集成员的列表。
// [keyn, scoren, keyn1, scoren1, ...]
func (rdm *RedisClient) ZRevRange(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZREVRANGE, args, includeArgs...)
}

// ZRANGEBYLEX key min max [LIMIT offset count],  在有序集合（sorted set）中按照字典序（lexicographical order）获取指定范围内的成员。这个命令主要用于那些成员是字符串的有序集合
// return	 指定区间内的元素列表。
// 这个的具体解释说明 看菜鸟教程  https://www.runoob.com/redis/sorted-sets-zrangebylex.html
func (rdm *RedisClient) ZRangeByLex(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZRANGEBYLEX, args, includeArgs...)
}

// ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count] , 通过分数返回有序集合指定区间内的成员, 有序集成员按分数值递增(从小到大)次序排列。
// 具有相同分数值的成员按字典序来排列(该属性是有序集提供的，不需要额外的计算)。
// 默认情况下，区间的取值使用闭区间 (小于等于或大于等于)，你也可以通过给参数前增加 ( 符号来使用可选的开区间 (小于或大于)。
// ZRANGEBYSCORE zset (1 5 , 返回所有符合条件 1 < score <= 5 的成员
// return 指定区间内，带有分数值(可选)的有序集成员的列表。
// [key1, score1, key2, score2, ...]
func (rdm *RedisClient) ZRangeByScore(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZRANGEBYSCORE, args, includeArgs...)
}

// ZREVRANGEBYSCORE key max min [WITHSCORES],  返回有序集中指定分数区间内的成员，分数从高到低排序,具有相同分数值的成员按字典序的逆序(reverse lexicographical order )排列。
// return 指定区间内，带有分数值(可选)的有序集成员的列表。
// [keyn, scoren, keyn1, scoren1, ...]
func (rdm *RedisClient) ZRevRangeByScore(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZREVRANGEBYSCORE, args, includeArgs...)
}

// ZRANK key member , 返回有序集中指定成员的排名。其中有序集成员按分数值递增(从小到大)顺序排列。
// return 如果成员是有序集 key 的成员，返回 member 的排名。 如果成员不是有序集 key 的成员，返回 nil 。
// 排名是从0开始的
func (rdm *RedisClient) ZRank(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZRANK, args, includeArgs...)
}

// ZREVRANK key member , 返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序,  排名以 0 为底，也就是说， 分数值最大的成员排名为 0 。
// return 如果成员是有序集 key 的成员，返回成员的排名。 如果成员不是有序集 key 的成员，返回 nil 。
func (rdm *RedisClient) ZRevRank(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZREVRANK, args, includeArgs...)
}

// ZREM key member [member2 ...], 移除有序集合中的一个或多个成员,不存在的成员将被忽略。
// return 被成功移除的成员的数量，不包括被忽略的成员。
func (rdm *RedisClient) ZRem(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZREM, args, includeArgs...)
}

// ZREMRANGEBYLEX key min max,  移除有序集合中给定的字典区间的所有成员。
// return 被成功移除的成员的数量，不包括被忽略的成员。
// ZREMRANGEBYLEX myzset [alpha [omega
func (rdm *RedisClient) ZRemRangeByLex(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZREMRANGEBYLEX, args, includeArgs...)
}

// ZREMRANGEBYRANK key start stop, 移除有序集中，指定排名(rank)区间内的所有成员。
// return 被移除成员的数量。
// ZREMRANGEBYRANK salary 0 1     # 移除下标 0 至 1 区间内的成员
func (rdm *RedisClient) ZRemRangeByRank(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZREMRANGEBYRANK, args, includeArgs...)
}

// ZREMRANGEBYSCORE key min max,  移除有序集中，指定分数（score）区间内的所有成员。
// return 被移除成员的数量。
// ZREMRANGEBYSCORE salary 1500 3500      # 移除所有薪水在 1500 到 3500 内的员工
func (rdm *RedisClient) ZRemRangeByScore(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZREMRANGEBYSCORE, args, includeArgs...)
}

// ZSCORE key member, 返回有序集中，成员的分数值
// return  成员的分数值，以字符串形式表示。
func (rdm *RedisClient) ZScore(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZSCORE, args, includeArgs...)
}

// ZINTERSTORE  destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE sum|min|max], 计算给定的一个或多个有序集的交集并将结果集存储在新的有序集合 destination 中
// destination：结果有序集合的名称。
// numkeys：要计算交集的有序集合的数量。
// key：参与计算的有序集合的名称。
// WEIGHTS weight [weight ...]：指定每个有序集合的权重。默认权重为 1。  作用对对应集合的元素分数加权， 然后在进行 AGGREGATE 操作
// AGGREGATE sum|min|max：指定交集结果的聚合方式。默认是 sum（求和）。  得到的交集元素的分数最后要经过AGGREGATE 操作
// 如：
// ZADD zset1 1 one 2 two
// ZADD zset2 1 one 3 two
// ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3 AGGREGATE SUM
// WEIGHTS 2 3：表示 zset1 的权重是 2，zset2 的权重是 3。
// AGGREGATE SUM：合并分数使用加法。
// 计算：
//
//	one: zset1 中的分数是 1 × 2 = 2；zset2 是 1 × 3 = 3；总分数是 2 + 3 = 5
//	two: zset1 中的分数是 2 × 2 = 4；zset2 是 3 × 3 = 9；总分数是 4 + 9 = 13
//
// 最终 out 集合的内容是： one: 5, two: 13
// return  保存到目标结果集的的成员数量。
func (rdm *RedisClient) ZInterStore(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZINTERSTORE, args, includeArgs...)
}

// ZINTER numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
// 从redis6.2开始支持， 要注意版本
// return  [key1, score1, key2, score2, ...]
func (rdm *RedisClient) ZInter(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZINTER, args, includeArgs...)
}

// ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX],
// 计算给定的一个或多个有序集的并集，并存储在新的 key 中, 默认情况下，结果集中某个成员的分数值是所有给定集下该成员分数值之和 。
// return 保存到 destination 的结果集的成员数量。
func (rdm *RedisClient) ZUnionStore(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZUNIONSTORE, args, includeArgs...)
}

// ZUNION numkeys key [key ...] [WEIGHTS weight [weight ...]] [AGGREGATE SUM|MIN|MAX] [WITHSCORES]
// 从redis6.2开始支持， 要注意版本
// return  [key1, score1, key2, score2, ...]
func (rdm *RedisClient) ZUnion(ctx base.BaseContext, cmd RdCmd, args map[string]any, includeArgs ...any) *redis.Cmd {
	return rdm.Handler(ctx, cmd, ZUNION, args, includeArgs...)
}
