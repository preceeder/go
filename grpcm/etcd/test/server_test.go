package proto

import (
	"context"
	"fmt"
	"github.com/preceeder/grpcm/etcd"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func Test_server(t *testing.T) {
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", 3000))
	if err != nil {
		fmt.Println("监听网络失败：", err)
		return
	}
	defer listener.Close()

	//创建grpc句柄
	srv := grpc.NewServer()
	defer srv.GracefulStop()

	// 注册服务到grpc
	RegisterGreetServer(srv, &greetServer{})

	etcdAddr := "127.0.0.1:2379"
	schema := "full"
	serviceName := "match/coin"
	serverAddr := "127.0.0.1:3000"
	var ttl int64 = 5
	etcdObj, err := etcd.NewEtcdServer(etcdAddr, schema, serviceName, serverAddr, ttl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	etcdObj.Register()

	//关闭信号处理
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		etcdObj.UnRegister()
		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()

	//监听服务
	err = srv.Serve(listener)
	if err != nil {
		fmt.Println("监听异常：", err)
		return
	}

}

// rpc服务接口
type greetServer struct{}

func (gs *greetServer) Morning(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	fmt.Printf("Morning 调用: %s\n", req.Name)
	return &GreetResponse{
		Message: "Good morning, " + req.Name,
		From:    fmt.Sprintf("127.0.0.1:%d", 3000),
	}, nil
}

func (gs *greetServer) Night(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	fmt.Printf("Night 调用: %s\n", req.Name)
	return &GreetResponse{
		Message: "Good night, " + req.Name,
		From:    fmt.Sprintf("127.0.0.1:%d", 3000),
	}, nil
}
