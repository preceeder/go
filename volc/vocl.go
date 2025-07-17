package volc

import (
	"crypto/sha256"
	"fmt"
	rtcbase "github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/rtc"
	"log/slog"
	"net/url"
	"sort"
	"strings"
	"time"
)

type VoclClient struct {
	Client *rtc.Rtc
	Config Config
}

func NewClient(config Config) *VoclClient {

	client := rtc.NewInstanceWithRegion(config.Region)
	client.SetCredential(rtcbase.Credentials{
		AccessKeyID:     config.AccessKey,
		SecretAccessKey: config.SecretKey,
	})
	return &VoclClient{Client: client, Config: config}
}

func (v VoclClient) GetToken(roomId, userId string) (string, error) {
	token := New(v.Config.AppId, v.Config.AppKey, roomId, userId)
	expt := time.Now().Add(time.Hour * 24)
	token.ExpireTime(expt)
	token.AddPrivilege(PrivSubscribeStream, expt)
	token.AddPrivilege(PrivPublishStream, expt)
	s, err := token.Serialize()
	if err != nil {
		slog.Error("生成token失败", "error", err.Error())
		return "", err
	}
	return s, nil
}

/**封禁房间, 用户 这里用来关闭房间
*appid  string 你的音视频应用的唯一标志
* roomId string 指定房间 ID
* userId string 希望封禁用户的 ID    只是关闭房间  传 ""
* ForbiddenInterval int 封禁时长，单位为秒  只是关闭房间  传 ""
 */
func (v VoclClient) CloseRoom(roomId, appId, userId, forbiddenInterval string) error {
	form := url.Values{}
	if appId == "" {
		form.Set("AppId", v.Config.AppId)
	} else {
		form.Set("AppId", appId)
	}
	form.Set("RoomId", roomId)
	if userId != "" {
		form.Set("UserId", userId)
	}
	if forbiddenInterval != "" {
		form.Set("ForbiddenInterval", forbiddenInterval)
	}
	res, status, err := v.Client.Post("BanRoomUser", nil, form)
	if err != nil {
		slog.Info("CloseRoom", "res", res, "status", status, "err", err)
	}
	return err
}

// 验证签名
func (v VoclClient) CheckSignature(event []string, Signature, secretKey string) bool {
	if secretKey == "" {
		secretKey = v.Config.CallBackSecretKey
	}
	data := append(event, secretKey)
	sort.Strings(data)
	payload := strings.Join(data, "")
	hashData := sha256.Sum256([]byte(payload))
	signature := fmt.Sprintf("%x", hashData)
	return signature == Signature
}
