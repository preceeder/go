package grpcm

import (
	"fmt"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/preceeder/grpcm/etcd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"log/slog"
	"time"
)

// Client 没有服务发现的grpc客户端
func Client(config RpcConfig, interceptor ...grpc.UnaryClientInterceptor) *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.Config{
				BaseDelay:  time.Second * 1,
				Multiplier: 1.6,
				MaxDelay:   time.Second * 15,
			},
			MinConnectTimeout: time.Second * 15,
		}),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithMax(10),
			grpc_retry.WithBackoff(grpc_retry.BackoffExponential(1*time.Second)),
		)),
		grpc.WithChainUnaryInterceptor(interceptor...),
	}
	clt, err := grpc.Dial(config.Addr, opts...)
	if err != nil {
		slog.Error("连接 gPRC 服务失败", "error", err)
	}
	return clt
}

// ClientWithEtcd 带有etcd 服务发现的 grpc 客户端
func ClientWithEtcd(c *RpcWithEtcdConfig, interceptor ...grpc.UnaryClientInterceptor) *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.Config{
				BaseDelay:  time.Second * 1,
				Multiplier: 1.6,
				MaxDelay:   time.Second * 15,
			},
			MinConnectTimeout: time.Second * 15,
		}),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithMax(10),
			grpc_retry.WithBackoff(grpc_retry.BackoffExponential(1*time.Second)),
		)),
		grpc.WithChainUnaryInterceptor(interceptor...),
	}
	r, err := etcd.NewResolver(c.EtcdAddrs, c.Schema, c.ServerName, c.Env)
	if err != nil {
		fmt.Println("etc 链接失败")
	}
	conn, err := grpc.Dial(r.GetGrpcDialTarget(), opts...)
	if err != nil {
		slog.Error("连接 gPRC 服务失败", "error", err)
	}
	return conn
}
