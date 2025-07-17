package grpcm

import (
	"fmt"
	"github.com/preceeder/grpcm/etcd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

// Server 没有etcd grpc 的服务
func Server(config *RpcConfig, server *grpc.Server) {
	// 创建 Tcp 连接

	if config == nil {
		slog.Error("rpc 还未初始化 请先调用 gobase.Init(rpc:true)")
	}
	rpcLi, err := net.Listen("tcp", config.Addr)
	if err != nil {
		slog.Error("监听失败: %v", "error", err.Error())
	}
	slog.Info("开启监听： ", "addr", config.Addr)

	////开启信号监听
	//c := utils.StartSignalLister()
	//
	////开启信号处理
	//go utils.SignalHandler(c, func() {
	//	//平滑关闭
	//	server.GracefulStop()
	//})

	//初始化 注册路由
	InitRpcRouter(server)
	// 在 gRPC 服务上注册反射服务
	reflection.Register(server)

	err = server.Serve(rpcLi)
	if err != nil {
		slog.Error("failed to serve: %v", err)
	}

}

// ServerWithEtcd 带有etcd 的 grpc服务
func ServerWithEtcd(c *RpcWithEtcdConfig, server *grpc.Server) {
	if c == nil {
		slog.Error("rpc 还未初始化 请先调用 gobase.Init(rpc:true)")
	}
	rpcLi, err := net.Listen("tcp", c.ServiceAddr)
	if err != nil {
		slog.Error("监听失败: %v", "error", err.Error())
	}
	slog.Info("开启监听： ", "addr", c.ServiceAddr)

	////开启信号监听
	//c := utils.StartSignalLister()
	//
	////开启信号处理
	//go utils.SignalHandler(c, func() {
	//	//平滑关闭
	//	server.GracefulStop()
	//})

	//初始化 注册路由
	InitRpcRouter(server)
	// 在 gRPC 服务上注册反射服务
	reflection.Register(server)

	// 将服务注册到  etcd 中
	etcdObj, err := etcd.NewEtcdServer(c.EtcdAddrs, c.Schema, c.ServerName, c.ServiceAddr, c.TTL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	etcdObj.Register()

	err = server.Serve(rpcLi)
	if err != nil {
		slog.Error("failed to serve: %v", err)
	}

}
