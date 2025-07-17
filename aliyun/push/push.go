package push

import (
	"encoding/json"
	number "github.com/alibabacloud-go/darabonba-number/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	push20160801 "github.com/alibabacloud-go/push-20160801/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/preceeder/go/base"
	"log/slog"
	"strings"
)

type AliPushConfig struct {
	Name     string `json:"name"`
	KeyId    string `json:"keyId"`
	Secret   string `json:"secret"`
	EndPoint string `json:"endPoint"`
	RegionId string `json:"regionId"`
	AppKey   string `json:"appKey"`
	Env      string `json:"env"`
}

type AliPushClient struct {
	Client *push20160801.Client
	Config AliPushConfig
}

// 注意 KeyId 和 Secret 是阿里云账号的权限accessKeyId 没有的话就会报错
func NewAliPushClient(config AliPushConfig) AliPushClient {
	client, err := CreateClient(&(config.KeyId), &(config.Secret), &(config.EndPoint), &(config.RegionId))
	if err != nil {
		slog.Error("阿里云push创建失败", "error", err.Error())
		panic("阿里云push创建失败：" + err.Error())
	}
	return AliPushClient{Client: client, Config: config}
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string, endpoint *string, regionId *string) (_result *push20160801.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
		RegionId:        regionId,
		Endpoint:        endpoint,
	}
	// Endpoint 请参考 https://api.aliyun.com/product/Push
	//config.Endpoint = endpoint // tea.String("cloudpush.aliyuncs.com")
	_result = &push20160801.Client{}
	_result, _err = push20160801.NewClient(config)
	return _result, _err
}

/**
 *  @param userIds []string 用户id 列表
 * @param title string 转为通知的标题
 * @param message map[string]any 发送给用户的消息， 自定义信息
 * @param StoreOffline bool 是否离线推送
 * @param alter bool 是否离线弹窗
 * @param content string 离线弹窗时的内容
 * @param env string PRODUCT | DEV
 */
func (p AliPushClient) GetMessageFormat(ctx base.Context, userIds []string, title string, message map[string]any, StoreOffline bool,
	alter bool, content string, env string) *push20160801.MassPushRequestPushTask {

	//message := map[string]any{
	//	"type":    "",
	//	"data":    message,
	//}
	if env == "" {
		env = p.Config.Env
	}
	extParameters, err := json.Marshal(map[string]any{"push": message})
	if err != nil {
		slog.ErrorContext(ctx, "message json marshal error", "error", err.Error())
		return nil
	}
	body, err := json.Marshal(message)
	if err != nil {
		slog.ErrorContext(ctx, "message json marshal error", "error", err.Error())
		return nil
	}

	pushTask := &push20160801.MassPushRequestPushTask{
		PushType:                       tea.String("MESSAGE"),
		DeviceType:                     tea.String("ALL"),
		StoreOffline:                   tea.Bool(StoreOffline),
		Target:                         tea.String("ACCOUNT"),
		TargetValue:                    tea.String(strings.Join(userIds, ",")),
		Title:                          tea.String(title),
		AndroidNotifyType:              tea.String("VIBRATE"),
		AndroidOpenType:                tea.String("APPLICATION"),
		AndroidActivity:                tea.String(""),
		AndroidNotificationBarType:     tea.Int32(50),
		AndroidNotificationBarPriority: tea.Int32(0),
		AndroidExtParameters:           tea.String(string(extParameters)),
		AndroidNotificationChannel:     tea.String("静默提醒"),
		IOSApnsEnv:                     tea.String(env),
		IOSSilentNotification:          tea.Bool(true),
		IOSMutableContent:              tea.Bool(true),
		IOSExtParameters:               tea.String(string(extParameters)),
		IOSBadgeAutoIncrement:          tea.Bool(false),
		Body:                           tea.String(string(body)),
	}
	if alter {
		pushTask.AndroidRemind = tea.Bool(true)
		pushTask.AndroidPopupActivity = tea.String("")
		pushTask.AndroidPopupTitle = tea.String(title)
		pushTask.AndroidPopupBody = tea.String(content)
		pushTask.IOSRemind = tea.Bool(true)
		pushTask.IOSRemindBody = tea.String(content)
	} else {
		pushTask.AndroidRemind = tea.Bool(false)
	}
	return pushTask
}

/*
 * @param pushTask *push20160801.MassPushRequestPushTask 先调用GetMessageFormat 拿到结果就是这里的参数
 * @param appKey string   由于android 和ios 可能不一样所以这里需要给个参数
 */
func (p AliPushClient) PushMessage(ctx base.Context, pushTask *push20160801.MassPushRequestPushTask) {
	request := &push20160801.MassPushRequest{
		AppKey:   number.ParseLong(&p.Config.AppKey),
		PushTask: []*push20160801.MassPushRequestPushTask{pushTask},
	}

	// request.pushTask = new Push20160801.MassPushRequest.pushTask{};
	res, _err := p.Client.MassPush(request)
	if _err != nil {
		slog.ErrorContext(ctx, "阿里云 推送消息失败", "error", _err.Error(), "task", pushTask.String())
		return
	}
	if *res.StatusCode != 200 {
		slog.ErrorContext(ctx, "阿里云 推送消息失败", "response body", res.String(), "task", pushTask.String())
	}
}

// 使用的时候 需要先调用GetMessageFormat， 然后在调用 PushMessage
