package etcd

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"log/slog"
	"net/url"
	"strings"
	"time"
)

type EtcdResolver struct {
	EtcdAddr   string              // 127.0.0.1:8080
	clientConn resolver.ClientConn // grpc 解析链接
	schema     string              // 这个可以自定定义
	etcdClient *clientv3.Client
	serverName string // 服务端etcd注册的服务名
	Env        string // test, product ...
}

func NewResolver(etcdAddr, schema, serverName, env string) (*EtcdResolver, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(etcdAddr, ";"),
		DialTimeout: 15 * time.Second,
	})
	if err != nil {
		slog.Error("连接etcd失败", err)
		return nil, err
	}
	r := &EtcdResolver{EtcdAddr: etcdAddr, etcdClient: cli, schema: schema, serverName: serverName, Env: env}
	// 注册这个 解析器
	resolver.Register(r)
	return r, nil
}

// GetGrpcDialTarget 返回 grpc 的链接 url;   和正常的 url是不一样的; grpc 是要通过这个解析器 解析这个url,
// 然后在通过Build函数拿到这个url对应的 grpc服务地址, grpc 会等待 e.clientConn.UpdateState 更新后才会真正去链接grpc 服务
func (e *EtcdResolver) GetGrpcDialTarget() string {
	return e.Scheme() + "://" + e.Env + "/" + e.serverName
}

func (e *EtcdResolver) ResolveNow(options resolver.ResolveNowOptions) {
	//TODO implement me
	slog.Info("ResolveNow", "options", options)
}

func (e *EtcdResolver) Close() {
	//TODO implement me
	slog.Info("close Resolver", "etcdresolver info", e.EtcdAddr)
}

func (e *EtcdResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	//TODO implement me
	e.clientConn = cc
	prfixKey, err := url.JoinPath(target.URL.Scheme, target.URL.Path)
	if err != nil {
		slog.Error("EtcdResolver error", "error", err.Error(), "target", target.String())
		return nil, err
	}
	go e.watch(prfixKey + "/")

	return e, nil
}

func (e *EtcdResolver) Scheme() string {
	//TODO implement me
	return e.schema
}

// 监听etcd中某个key前缀的服务地址列表的变化
func (e *EtcdResolver) watch(keyPrefix string) {
	//初始化服务地址列表
	var addrList []resolver.Address

	resp, err := e.etcdClient.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		slog.Error("获取服务地址列表失败", "error", err.Error())
	} else {
		for i := range resp.Kvs {
			addrList = append(addrList, resolver.Address{Addr: strings.TrimPrefix(string(resp.Kvs[i].Key), keyPrefix)})
		}
	}
	e.clientConn.UpdateState(resolver.State{Addresses: addrList})

	//监听服务地址列表的变化
	rch := e.etcdClient.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	for n := range rch {
		for _, ev := range n.Events {
			addr := strings.TrimPrefix(string(ev.Kv.Key), keyPrefix)
			switch ev.Type {
			case mvccpb.PUT:
				if !exists(addrList, addr) {
					addrList = append(addrList, resolver.Address{Addr: addr})
					e.clientConn.UpdateState(resolver.State{Addresses: addrList})
				}
			case mvccpb.DELETE:
				if s, ok := remove(addrList, addr); ok {
					addrList = s
					e.clientConn.UpdateState(resolver.State{Addresses: addrList})
				}
			}
		}
	}
}

func exists(l []resolver.Address, addr string) bool {
	for i := range l {
		if l[i].Addr == addr {
			return true
		}
	}
	return false
}

func remove(s []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}
