package producer

import (
	"github.com/preceeder/go-nsq"
	"log/slog"
	"strings"
	"time"
)

var NsqProducer = map[string]*Producer{}

type NsqProducerConfig struct {
	Name        string `json:"name" default:"default"`
	Addr        string `json:"nsqdAddr" default:"127.0.0.1:8009"`
	TopicPrefix string `json:"topicPrefix"` // 默认default
}

func InitNsqProducer(nsqConfig []NsqProducerConfig) {
	//nsqConfig := []NsqProducerConfig{}
	//utils.ReadViperConfig(config, "nsq-producer", &nsqConfig)
	for _, nsqc := range nsqConfig {
		_, err := NewProduct(nsqc.Addr, nsqc.Name, nsqc.TopicPrefix)
		if err != nil {
			slog.Error("InitNsqProducer error ", "error", err.Error())
			panic("InitNsqProducer error: " + err.Error())
		}
	}
	////开启信号监听
	//signl := utils.StartSignalLister()
	//
	////开启信号处理
	//go utils.SignalHandler(signl, func() {
	//	//平滑关闭
	//	for _, v := range NsqProducer {
	//		slog.Info("stop nsq producer", "addr", v.String())
	//		v.Stop()
	//	}
	//	os.Exit(1)
	//})
}

// NewProduct
// nsqAddr nsqd的tcp地址
// name 内部用于标记的名字, 保持唯一就可以
func NewProduct(nsqdAddr, name string, topicPrefix string) (*Producer, error) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(nsqdAddr, config)
	if err != nil {
		slog.Error("create producer failed", "error", err.Error())
		return nil, err
	}
	err = producer.Ping()
	if err != nil {
		slog.Error("nsq producer ping", "error", err.Error())
	}

	if topicPrefix == "" {
		topicPrefix = "default"
	}

	NsqProducer[name] = &Producer{
		cf: NsqProducerConfig{
			Name:        name,
			Addr:        nsqdAddr,
			TopicPrefix: topicPrefix,
		},
		pr: producer,
	}
	slog.Info("开启 nsq producer", "addr", producer.String())
	return NsqProducer[name], nil
}

type Producer struct {
	cf NsqProducerConfig
	pr *nsq.Producer
}

func (p *Producer) DeferredPublish(topic string, delay time.Duration, body []byte) error {
	err := p.pr.DeferredPublish(strings.Join([]string{p.cf.TopicPrefix, topic}, "."), delay, body)
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) Publish(topic string, body []byte) error {
	err := p.pr.Publish(strings.Join([]string{p.cf.TopicPrefix, topic}, "."), body)
	if err != nil {
		return err
	}
	return nil
}

// 向同一个topic同时发布多个消息
func (p *Producer) MultiPublish(topic string, body [][]byte) error {
	err := p.pr.MultiPublish(strings.Join([]string{p.cf.TopicPrefix, topic}, "."), body)
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) Stop() {
	p.pr.Stop()
}
