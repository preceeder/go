package proto

import (
	"context"
	"fmt"
	"github.com/preceeder/grpcm/etcd"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func Test_client(t *testing.T) {
	etcdAddr := "127.0.0.1:2379"
	schema := "full"
	serverName := "match/coin"
	env := "test"
	r, err := etcd.NewResolver(etcdAddr, schema, serverName, env)
	if err != nil {
		fmt.Println("etc 链接失败")
	}
	conn, err := grpc.Dial(r.GetGrpcDialTarget(), grpc.WithInsecure())
	if err != nil {
		fmt.Println("服务器 链接失败")
	}
	defer conn.Close()

	//获得grpc句柄
	c := NewGreetClient(conn)
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Println("Morning 调用...")
		resp1, err := c.Morning(
			context.Background(),
			&GreetRequest{Name: "JetWu"},
		)
		if err != nil {
			fmt.Println("Morning调用失败：", err)
			time.Sleep(time.Second * 5)
			continue
		}
		fmt.Printf("Morning 响应：%s，来自：%s\n", resp1.Message, resp1.From)

		fmt.Println("Night 调用...")
		resp2, err := c.Night(
			context.Background(),
			&GreetRequest{Name: "JetWu"},
		)
		if err != nil {
			fmt.Println("Night调用失败：", err)
			time.Sleep(time.Second * 5)
			continue
		}
		fmt.Printf("Night 响应：%s，来自：%s\n", resp2.Message, resp2.From)
	}
}
