package proto

import (
	"context"
	"fmt"
	"github.com/preceeder/grpcm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log/slog"
	"strconv"
	"testing"
	"time"
)

func init() {
	grpcm.RpcRegister(&Greet_ServiceDesc, &greetServer{})
}

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (m interface{}, err error) {
	requestId := strconv.FormatInt(time.Now().Unix(), 10)

	defer func() {
		if r := recover(); r != nil {
			slog.Error("", "error", err, "fun", "WealthLevelHandler", "requestId", requestId)

		}
	}()

	md, ok := metadata.FromIncomingContext(ctx)
	slog.Info("", "ip", md.Get(":authority")[0], "user-agent", md.Get("user-agent")[0], "params", req, "requestId", requestId)
	if ok {
		if sec := md.Get("secret"); len(sec) > 0 {
			if sec[0] != "123" {
				return nil, status.Errorf(codes.PermissionDenied, "method MyTest not found")
			}
		}
	}
	ctx = context.WithValue(ctx, "requestId", requestId)

	m, err = handler(ctx, req)
	slog.Info("rpc 返回数据", "res", m, "requestId", requestId)
	return m, err
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

func Test_server(t *testing.T) {
	c := &grpcm.RpcWithEtcdConfig{
		EtcdAddrs:   "127.0.0.1:2379",
		Schema:      "full",
		ServerName:  "test",
		ServiceAddr: "127.0.0.1:3000",
		Env:         "test",
	}
	gserver := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptor))
	////开启信号监听
	//fsingal := utils.StartSignalLister()
	//
	////开启信号处理
	//go utils.SignalHandler(fsingal, func() {
	//	//平滑关闭
	//	gserver.GracefulStop()
	//})

	grpcm.ServerWithEtcd(c, gserver)

}
