package grpcm

type RpcConfig struct {
	Addr string `json:"addr"`
	Name string `json:"name"`
}

// RpcWithEtcdConfig
// 最后给到grpc 客户端的target地址是: schema://env/servername; 最终etcd 会拿 [schema/servername/] 去获取真正的grpc地址
type RpcWithEtcdConfig struct {
	EtcdAddrs   string // etcd 的地址 多个地址以 ; 分隔  "127.0.0.1:2379;101.23.12.3:2379"
	Schema      string // "full"  这里就是一个字符串 可以随意
	ServerName  string // 服务名字 客户端和服务端要一致
	ServiceAddr string // 127.0.0.1:3000
	Env         string // 环境  product / test
	TTL         int64  // etcd 心跳时间
}
