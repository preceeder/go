package proto

import (
	"context"
	"fmt"
	"github.com/preceeder/grpcm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"testing"
	"time"
)

func Test_client(t *testing.T) {
	c := &grpcm.RpcWithEtcdConfig{
		EtcdAddrs:   "127.0.0.1:2379",
		Schema:      "full",
		ServerName:  "test",
		ServiceAddr: "127.0.0.1:3000",
		Env:         "test",
	}
	conn := grpcm.ClientWithEtcd(c, OrderUnaryClientInterceptor)
	//获得grpc句柄
	gtr := NewGreetClient(conn)
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Println("Morning 调用...")
		resp1, err := gtr.Morning(
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
		resp2, err := gtr.Night(
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

func OrderUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "secret", "123", "timestamp", fmt.Sprintf("%d", time.Now().Unix()))
	// Invoking the remote method
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}
