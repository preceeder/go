package rongyun

import (
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/rongcloud/server-sdk-go/v3/sdk"
	"strings"
)

type RongyunConfig struct {
	Prefix           string `json:"prefix"`         // 用户前缀
	ChatRoomPrefix   string `json:"chatRoomPrefix"` // 聊天室前缀
	AppKey           string `json:"appKey"`
	AppSecret        string `json:"appSecret"`
	WithRongCloudURI string `json:"withRongCloudURI"`
}
type RongYunClient struct {
	Client *sdk.RongCloud
	Config RongyunConfig
}

func NewRongYunClient(config RongyunConfig) *RongYunClient {
	Rc := sdk.NewRongCloud(config.AppKey,
		config.AppSecret,
		sdk.Region{}, // TODO 这里的数据还没给
		// 每个域名最大活跃连接数
		sdk.WithMaxIdleConnsPerHost(100),
		sdk.WithTimeout(10),
		sdk.WithRongCloudURI(config.WithRongCloudURI),
	)

	return &RongYunClient{Client: Rc, Config: config}
}

// 回调参数校验
func (rc RongYunClient) CallDataCheck(timestamp, nonce, signature string) bool {
	data := []string{rc.Config.AppSecret, nonce, timestamp}
	chd := cryptor.Sha1(strings.Join(data, ""))
	if chd == signature {
		return true
	}
	return false
}
