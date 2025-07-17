package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log/slog"
	"strings"
	"time"
)

type EtcdServer struct {
	etcdAddr    string // etcd 地址
	serviceName string // 客户端要和 服务端一致
	serverAddr  string // grpc 服务地址    127.0.0.1:3000
	ttl         int64  // etcd 保持心跳时间
	schema      string
	etcdClient  *clientv3.Client
}

func NewEtcdServer(etcdAddr, schema, serviceName, serverAddr string, ttl int64) (*EtcdServer, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(etcdAddr, ";"),
		DialTimeout: 15 * time.Second,
	})
	if err != nil {
		slog.Error("连接etcd失败", err)
		return nil, err
	}
	return &EtcdServer{
		etcdClient:  cli,
		etcdAddr:    etcdAddr,
		schema:      schema,
		serviceName: serviceName,
		serverAddr:  serverAddr,
		ttl:         ttl,
	}, nil
}

func (e *EtcdServer) Register() {
	//与etcd建立长连接，并保证连接不断(心跳检测)
	ticker := time.NewTicker(time.Second * time.Duration(e.ttl))
	go func() {
		key := e.GetKey()
		for {
			resp, err := e.etcdClient.Get(context.Background(), key)
			if err != nil {
				slog.Error("获取服务地址失败", "error", err.Error())
			} else if resp.Count == 0 { //尚未注册
				err = e.keepAlive()
				if err != nil {
					slog.Error("租约自动续期操作失败", "error", err.Error())
				}
			}
			<-ticker.C
		}
	}()
}

func (e *EtcdServer) GetKey() string {
	key := e.schema + "/" + e.serviceName + "/" + e.serverAddr
	return key
}

// 保持服务器与etcd的长连接
func (e *EtcdServer) keepAlive() error {
	//创建租约
	leaseResp, err := e.etcdClient.Grant(context.Background(), e.ttl)
	if err != nil {
		slog.Error("创建租约失败", "error", err.Error())
		return err
	}

	//将服务地址注册到etcd中
	_, err = e.etcdClient.Put(context.Background(), e.GetKey(), e.serverAddr, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		slog.Error("注册服务失败", "error", err.Error())
		return err
	}

	//租约自动续约
	ch, err := e.etcdClient.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		slog.Error("创建租约自动续约失败", "error", err.Error())
		return err
	}
	//清空keepAlive返回的channel
	go func() {
		for {
			<-ch
		}
	}()
	return nil
}

// 取消注册
func (e *EtcdServer) UnRegister() {
	if e.etcdClient != nil {
		e.etcdClient.Delete(context.Background(), e.GetKey())
	}
}
