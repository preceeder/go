package jigou

import (
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/go-resty/resty/v2"
	"slices"
	"strings"
	"time"
)

// AppId = ""
// ServerSecret = ""

type JigouClient struct {
	Config JigouConfig
	Client *resty.Client
	Rtc
}

func NewJiGouClient(config JigouConfig) *JigouClient {
	httpClient := resty.New()
	httpClient.SetTimeout(time.Duration(5 * time.Second))

	client := &JigouClient{
		Client: httpClient,
		Config: config,
		Rtc:    Rtc{Client: httpClient, Config: config},
	}

	return client
}

// CallDataCheck 回调参数校验
func (j *JigouClient) CallDataCheck(timestamp, nonce, signature string) bool {
	data := []string{j.Config.CallBackSecret, timestamp, nonce}
	slices.Sort(data)
	chd := cryptor.Sha1(strings.Join(data, ""))
	if chd == signature {
		return true
	}
	return false
}
