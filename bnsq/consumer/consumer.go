package consumer

import (
	"fmt"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/preceeder/go-nsq"
	"github.com/preceeder/go/bnsq/common"
	"log/slog"
	"strings"
	"time"
)

type NsqConfig struct {
	NsqLookupd   []string `json:"lookupds"`     // nsqLookupds的http 地址
	PollInterval int64    `json:"pollInterval"` // ms 轮询间隔时间   500ms就可以
	MaxInFlight  int      `json:"maxInFlight"`  // 一次最多处理的消息数量   // 最少设置 5个
	MaxAttempts  uint16   `json:"maxAttempts"`  // 最大重试次数， 超过这个次数消息就作废
	LogLevel     string   `json:"logLevel"`     //"debug", "info", "warring", "error"
	MsgTimeout   int      `json:"msg_timeout"`  //  s  单条消息处理的超时时间   nsqd 默认的最大值是15min
	SlogLogger   *slog.Logger
	//MaxRdyCount  int      `json:"max_rdy_count"` // 每个nsqd 可以接受的最大消费者数量  默认2500
	TopicPrefix  string        `json:"topicPrefix"` // topic的前缀， 一版用来区分正式服测试服   默认 default
	pollInterval time.Duration // PollInterval处理后的时间
	logLevel     nsq.LogLevel  //LogLevel 对于 ["debug", "info", "warring", "error"] 的下标
	msgTimeout   time.Duration // MsgTimeout 处理后的时间

}

var logLevel = []string{"debug", "info", "warring", "error"}

// SubRouter 子路由
type SubRouter struct {
	Topic   string
	Channel string
	Handler []MqMiddlewareFunc
}

// Routers 路由列表
var Routers []SubRouter = make([]SubRouter, 0)

// RegisterRouter 添加路由到 队列
func RegisterRouter(topic, channel string, handler ...MqMiddlewareFunc) {
	Routers = append(Routers, SubRouter{topic, channel, handler})
}

// 中间件类型定义
type MqMiddlewareFunc func(MqHandlerFunc) MqHandlerFunc

// NsqConsumer 定义 新的 NsqConsumer 对象
type NsqConsumer struct {
	config *NsqConfig
	// 注册的 consumers 对象
	consumers   []*nsq.Consumer
	logger      *common.Logger
	middlewares []MqMiddlewareFunc // 存储中间件链
}

// MqHandlerFunc 消息 handler 函数定义
type MqHandlerFunc func(msg *nsq.Message) error

// NewNsqConsumer 初始化 Nsq consumer 对象
func NewNsqConsumer(cf *NsqConfig) *NsqConsumer {
	cf.logLevel = 2
	if slice.Contain(logLevel, cf.LogLevel) {
		cf.logLevel = nsq.LogLevel(slice.IndexOf(logLevel, cf.LogLevel))
	}
	cf.msgTimeout = time.Second * 60
	if cf.MsgTimeout > 0 {
		cf.msgTimeout = time.Second * time.Duration(cf.MsgTimeout)
	}

	if cf.MaxAttempts == 0 {
		cf.MaxAttempts = 5
	}

	if cf.PollInterval == 0 {
		cf.PollInterval = 500
	}
	cf.pollInterval = time.Millisecond * time.Duration(cf.PollInterval)

	if cf.TopicPrefix == "" {
		cf.TopicPrefix = "default"
	}

	if cf.SlogLogger == nil {
		fmt.Println("日志没有配置")
	}

	return &NsqConsumer{
		config: cf,
		logger: &common.Logger{
			Logger: cf.SlogLogger,
		},
	}
}

func (n *NsqConsumer) AddMiddlewares(middlewares ...MqMiddlewareFunc) {
	n.middlewares = append(n.middlewares, middlewares...)
}

// RegisterHandler 注册 topic 和 handler Func
func (n *NsqConsumer) RegisterHandler(topic, channel string, handler ...MqMiddlewareFunc) error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = n.config.pollInterval
	cfg.MsgTimeout = n.config.msgTimeout
	cfg.MaxAttempts = n.config.MaxAttempts
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		return err
	}
	c.SetLogger(n.logger, n.config.logLevel)
	c.ChangeMaxInFlight(n.config.MaxInFlight)

	// 将多个中间件按顺序链式连接起来
	finalHandler := func(msg *nsq.Message) error {
		return nil // 最后一个 handler 如果不做处理，可以返回 nil
	}

	//添加指定接口中间
	for i := len(handler) - 1; i >= 0; i-- {
		finalHandler = handler[i](finalHandler)
	}

	// 添加全局中间件
	for i := len(n.middlewares) - 1; i >= 0; i-- {
		finalHandler = n.middlewares[i](finalHandler)
	}

	c.AddHandler(n.toNsqHandler(finalHandler))
	n.consumers = append(n.consumers, c)
	return nil
}

// 将 MqHandlerFunc 转换成 nsq.Consumer 内接受的 nsq.HandlerFunc
func (n *NsqConsumer) toNsqHandler(handlerFunc MqHandlerFunc) nsq.HandlerFunc {
	return func(msg *nsq.Message) error {
		return handlerFunc(msg)
	}
}

func (n *NsqConsumer) Start() error {
	for _, router := range Routers {
		if err := n.RegisterHandler(strings.Join([]string{n.config.TopicPrefix, router.Topic}, "."), router.Channel, router.Handler...); err != nil {
			return err
		}
	}

	for _, h := range n.consumers {
		if err := h.ConnectToNSQLookupds(n.config.NsqLookupd); err != nil {
			return err
		}
	}
	n.logger.Logger.Info("启动nsq 消费者成功")
	return nil
}

// 开始并且阻塞
func (n *NsqConsumer) StartWithBlock(f func()) error {
	for _, router := range Routers {
		n.logger.Logger.Info("nsq消费时创建", "topic", router.Topic, "channel", router.Channel)
		if err := n.RegisterHandler(router.Topic, router.Channel, router.Handler...); err != nil {
			return err
		}
	}

	for _, h := range n.consumers {
		if err := h.ConnectToNSQLookupds(n.config.NsqLookupd); err != nil {
			return err
		}
	}
	n.logger.Logger.Info("启动nsq 消费者成功")
	//开启信号监听
	common.StartSignalLister(func() {
		//平滑关闭
		n.Close()
		f()
	})
	return nil
}

// Close consumer
func (n *NsqConsumer) Close() {
	for _, h := range n.consumers {
		h.Stop()
	}
}
