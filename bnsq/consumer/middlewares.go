package consumer

import (
	"github.com/preceeder/go-nsq"
	"log/slog"
	"runtime/debug"
	"time"
)

// json 字符串可以不加转议符"\" 输出
type LogStr string

func (d LogStr) MarshalJSON() ([]byte, error) {
	return []byte(d), nil
}

// 日志中间件
func LogMiddleware(next MqHandlerFunc) MqHandlerFunc {
	return func(msg *nsq.Message) error {
		slog.Info("消费消息",
			"topic", msg.Topic,
			"channel", msg.Channel, "messageId", string(msg.ID[:]),
			"Attempts", msg.Attempts,
			"body", LogStr(msg.Body))
		err := next(msg)
		if err != nil {
			slog.Error("日志中间件发生错误", "error", err.Error())
		}
		return err
	}
}

func CatchError(next MqHandlerFunc) MqHandlerFunc {
	return func(msg *nsq.Message) error {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("处理发生错误", "topic", msg.Topic, "channel", msg.Channel, "data", LogStr(msg.Body), "messageId", string(msg.ID[:]), "Attempts", msg.Attempts, "error", debug.Stack())
				msg.Requeue(time.Second * 10)
			}
		}()

		return next(msg)
	}
}
