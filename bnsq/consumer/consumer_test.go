package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/preceeder/go-nsq"
	"github.com/preceeder/go/bnsq/common"
	"log/slog"
	"os"
	"sync"
	testing2 "testing"
	"time"
)

func init() {
	RegisterRouter("shumei", "image", HandlerTime, Image)
	//RegisterRouter("shumei", "video", Video)
}

//// MyHandler 是一个消费者类型
//type ShuMei struct {
//	Body any `json:"body"`
//}

func fos(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 20)
	wg.Done()
}

func HandlerTime(f MqHandlerFunc) MqHandlerFunc {
	return func(message *nsq.Message) error {
		fmt.Println("HandlerTime")
		return f(message)
	}
}

func Image(f MqHandlerFunc) MqHandlerFunc {
	return func(msg *nsq.Message) error {
		fmt.Println("Image")
		// 数美的 图片， 域名， 文字 都可以通过这个， 但是数美的视频必须是异步的
		dd := map[string]any{}
		err := json.Unmarshal(msg.Body, &dd)
		if err != nil {
			return err
		}
		fmt.Println("收到数据 Image", string(msg.ID[:]), dd, time.Now().Format("2006-01-02 15:04:05"))

		fmt.Println("Image 处理完成", msg.ID, dd, time.Now().Format("2006-01-02 15:04:05"))
		//msg.Requeue(time.Second * 11)
		//return nil
		//msg.Finish()
		//msg.Touch()
		// 每个消息都有一定的消息处理时间, 一般是1分钟, 最大可设置的时间是15分钟  -msg-timeout=1min
		// 重置消息的超时时间
		msg.Finish()
		return nil
	}
}

func Video(msg *nsq.Message) (err error) {
	// 数美的 图片， 域名， 文字 都可以通过这个， 但是数美的视频必须是异步的
	dd := map[string]any{}
	err = json.Unmarshal(msg.Body, &dd)
	if err != nil {
		return err
	}
	fmt.Println("收到数据 Video", dd, time.Now().UnixMilli())
	//msg.Finish()
	//msg.Touch()
	// 每个消息都有一定的消息处理时间, 一般是1分钟, 最大可设置的时间是15分钟  -msg-timeout=1min
	// 重置消息的超时时间
	msg.Finish()

	return nil
}

func TestStart(t *testing2.T) {
	//&logs.SlogConfig{
	//	InfoFileName:            "logs/nsq/out.log",
	//	MaxSize:                 100,
	//	MaxAge:                  15,
	//	MaxBackups:              20,
	//	TransparentTransmission: true,
	//	StdOut:                  "1",
	//	Compress:                true,
	//}

	config := NsqConfig{
		NsqLookupd:   []string{"127.0.0.1:4131"},
		PollInterval: 500,
		MaxInFlight:  3,
		MaxAttempts:  3,
		MsgTimeout:   10,
		LogLevel:     "warning",
		SlogLogger:   slog.Default(),
	}
	sqConsumer := NewNsqConsumer(&config)
	sqConsumer.AddMiddlewares(LogMiddleware)
	err := sqConsumer.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	//开启信号监听
	common.StartSignalLister(func() {
		//平滑关闭
		sqConsumer.Close()
	})
}
