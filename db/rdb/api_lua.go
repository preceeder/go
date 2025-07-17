package rdb

import (
	"fmt"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/preceeder/go/base"
	"github.com/redis/go-redis/v9"
	"time"
)

type LuaScript struct {
	Script  string
	Keys    []string
	Args    []string
	Default map[string]any
}

// 缓存Lua脚本到redis
// return 给定脚本的 SHA1 校验和
func (rdm *RedisClient) ScriptLoad(ctx base.BaseContext, lua string) string {
	cmd := rdm.Client.ScriptLoad(ctx, lua)
	//hesHasScript := cryptor.Sha1(lua)
	return cmd.Val()
}

// 这里还需要实验一下
func (rdm *RedisClient) EvalSha(ctx base.BaseContext, lua string, keys []string, values []any) *redis.Cmd {
	hesHasScript := cryptor.Sha1(lua)
	cmd := rdm.Client.EvalSha(ctx, hesHasScript, keys, values)
	if cmd.Err() != nil {
		if redis.HasErrorPrefix(cmd.Err(), "NOSCRIPT") {
			// 如果是没有 sha的报错需要重新load
			rdm.ScriptLoad(ctx, lua)
			cmd = rdm.Client.EvalSha(ctx, hesHasScript, keys, values)
			return cmd
		}
	}
	return cmd
}

func (rdm *RedisClient) ExecScript(ctx base.BaseContext, lua LuaScript, keyInfo map[string]string, valueInfo map[string]any) *redis.Cmd {
	var defaultData map[string]any = make(map[string]any)
	if len(lua.Default) > 0 {
		defaultData = handlerDefaultValue(lua.Default)
	}
	var err error
	keys := []string{}
	if len(lua.Keys) > 0 {
		keys, err = getValues(lua.Keys, keyInfo, defaultData)
	}
	values := []any{}
	if len(lua.Args) > 0 {
		values, err = getValues(lua.Args, valueInfo, defaultData)
	}
	if err != nil {
		cmd := redis.Cmd{}
		cmd.SetErr(err)
		return &cmd
	}

	return rdm.EvalSha(ctx, lua.Script, keys, values)
}

func handlerDefaultValue(data map[string]any) map[string]any {
	for k, v := range data {
		if f, ok := v.(func() time.Duration); ok {
			data[k] = int64(f() / time.Second) // 一半都是过期时间， 计算到秒
		}
	}
	return data
}

func getValues[T string | any](keyNames []string, keyInfo map[string]T, defaultData map[string]any) ([]T, error) {
	var keys []T = make([]T, 0, len(keyNames))
	for _, key := range keyNames {
		if v, ok := keyInfo[key]; ok {
			keys = append(keys, v)
		} else {
			if dv, exit := defaultData[key]; exit {
				keys = append(keys, dv.(T))
			} else {
				return nil, fmt.Errorf("key %s not found in default data", key)
			}
		}

	}
	return keys, nil
}

//func getValues(keyNames []string, keyInfo map[string]any, defaultData map[string]any) []any {
//	var values []any = make([]any, 0, len(keyNames))
//	for i, key := range keyNames {
//		if v, ok := defaultData[key]; ok {
//			values[i] = v
//		}
//
//		if v, ok := keyInfo[key]; ok {
//			values[i] = v
//		} else {
//			// TODO error 参数缺失
//		}
//	}
//	return values
//}
