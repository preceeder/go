package rdb

import (
	"encoding/json"
	"fmt"
	"github.com/preceeder/go/base"
	"testing"
	"time"
)

// var client *RedisClient
//
//	func init() {
//		config := Config{
//			Host:        "127.0.0.1",
//			Port:        "6377",
//			Password:    "QDjk9UdkoD6cv",
//			Db:          0,
//			MaxIdle:     2,
//			IdleTimeout: 240,
//			PoolSize:    13,
//		}
//		client = NewRedisClient(config)
//	}
func Test_handlerDefaultValue(t *testing.T) {
	data := map[string]any{
		"name": 23,
		"exp":  func() time.Duration { return time.Second * 5 },
	}
	fmt.Println(handlerDefaultValue(data))
	fmt.Println(int(time.Second * 5 / time.Second))
}

func Test_getValues(t *testing.T) {
	values, err := getValues[any]([]string{"name", "age"}, nil, map[string]any{"name": "niy", "age": 23})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(values)

	var d []string = nil
	fmt.Println(len(d))
}

func TestRedisClient_EvalSha(t *testing.T) {
	client := InitRedis()
	client.EvalSha(base.Context{}, RkSetUserList.Script, RkSetUserList.Keys, nil)
}
func TestRedisClient_ExecScript(t *testing.T) {
	client := InitRedis()
	cmd := client.ZAdd(base.Context{}, ade, nil, 3, 2, 5, 4, 5, 6, 9, 7)
	if cmd.Err() != nil {
		fmt.Println(cmd.Err())
		return
	}
	fmt.Println(cmd.Val())
	cmd = client.ExecScript(base.Context{},
		RkSetUserList,
		map[string]string{"paramsKey": "haha"},
		map[string]any{"userIds": StUserIds{1, 2, 3, 4}, "size": 30, "expireT": 23},
	)
	fmt.Println(cmd.Val())
}

type StUserIds []any

func (s StUserIds) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s StUserIds) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}

var RkSetUserList = LuaScript{
	Script: SAVE_DISCONT_USER_LIST,
	Keys:   []string{"paramsKey"},
	Args:   []string{"userIds", "size", "expireT"},
	Default: map[string]any{
		"size":    20,
		"expireT": int((time.Hour * 24).Seconds()),
	},
}

// 缓存用户列表 去重 并返回用户数据  lua版本小于5.2用 unpack， 大于等于5.2 用table.unpack
var SAVE_DISCONT_USER_LIST string = `
	local redisK =  KEYS[1]
	local user_ids_data = cjson.decode(ARGV[1])	
	local size = ARGV[2]
	local expireT = ARGV[3]
	local tempKey = redisK .. "_temp"

	-- 用户列表 保存到一个临时key
	redis.call("ZADD", tempKey, unpack(user_ids_data))
	
    -- 用户数据缓存数据 和临时数据 做并集 原本的数据分数不会改变  去掉重复的 保存到使用的key 中
    redis.call("ZUNIONSTORE", redisK, 2,  redisK, tempKey, "AGGREGATE", "MIN")
	
    -- 删除临时数据
	redis.call("DEL", tempKey)	
	
	-- 设置缓存数据的过期时间
	redis.call("expire", redisK, expireT)
	
	-- 获取当前分数的 数据
	local resData = redis.call("ZRANGEBYSCORE", redisK, user_ids_data[1], user_ids_data[1]+size)
	return resData`
