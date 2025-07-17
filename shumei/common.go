package shumei

import (
	"github.com/go-resty/resty/v2"
	"math/rand/v2"
	"net/http"
	"time"
)

// 注意事项   SetResult() 方法在 status_code > 399的时候失效， 需要设置SetErr()
// res, err := rc.R().
//
//	SetResult(response).
//	ForceContentType("application/json").  // 接受数据强制转为 指定类型
//	SetHeader("Accept", "application/json").
//	SetHeader("Content-Type", "application/json").
//	EnableTrace().
//	SetBody(body).
//	Post(url)
func NewRequestClient() *resty.Client {
	var RequestClient = resty.New()
	RequestClient.SetHeaders(
		map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
	)
	RequestClient.SetTimeout(3 * time.Second)
	RequestClient.SetTransport(&http.Transport{
		MaxIdleConnsPerHost:   100,              // 对于每个主机，保持最大空闲连接数为 10
		IdleConnTimeout:       30 * time.Second, // 空闲连接超时时间为 30 秒
		TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时时间为 10 秒
		ResponseHeaderTimeout: 20 * time.Second, // 等待响应头的超时时间为 20 秒
	})
	return RequestClient
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.IntN(len(letters))]
	}
	return string(rand_bytes)
}
