package rdb

import (
	"fmt"
	"github.com/preceeder/go/base"
	"testing"
	"time"
)

func TestRedisClient_ZRange(t *testing.T) {
	client := InitRedis()
	cmd := client.ZAdd(base.Context{}, ade, map[string]any{
		"score1": 3, "key1": 2, "score2": 5, "key2": 4, "score3": 5, "key3": 6,
	})
	if cmd.Err() != nil {
		fmt.Println(cmd.Err())
		return
	}
	fmt.Println(cmd.Val())

	cmd = client.ZRange(base.Context{}, ade, map[string]any{"start": 0, "stop": -1})
	fmt.Println(cmd.Val())
}

var ade = RdCmd{
	Key: "haha",
	CMD: map[Command]RdSubCmd{
		ZADD: {
			Params: "{{score1}} {{key1}} {{score2}} {{key2}} {{score3}} {{key3}}",
			Exp: func() time.Duration {
				return time.Second * 30
			},
		},
		ZRANGE: {
			Params: "{{start}} {{stop}} WITHSCORES",
		},
	},
}
