package feishu

import (
	"context"
	"encoding/json"
	"github.com/fanjindong/go-cache"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkauth "github.com/larksuite/oapi-sdk-go/v3/service/auth/v3"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"log/slog"
	"time"
)

type FConfig struct {
	Appid     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type FeiShu struct {
	Client    *lark.Client
	AppId     string
	AppSecret string
	Cache     cache.ICache
}

func NewFeiShuClient(appId, appSecret string) *FeiShu {
	client := lark.NewClient("cli_a55848a6d038100c", "f8XZICshhQtsGBZMysxLXg3MKQSiOGvF")
	return &FeiShu{
		Client:    client,
		AppId:     appId,
		AppSecret: appSecret,
		Cache:     cache.NewMemCache(),
	}
}

type TenantAccessTokenResponse struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}

// GetTenantAccessToken
// 获取 TenantAccessToken   两小时有效
func (client *FeiShu) GetTenantAccessToken(ctx context.Context) string {
	resData, ok := client.Cache.Get("TenantAccessToken")
	if ok && resData != nil {
		return resData.(string)
	}
	// 创建请求对象
	req := larkauth.NewInternalTenantAccessTokenReqBuilder().
		Body(larkauth.NewInternalTenantAccessTokenReqBodyBuilder().
			AppId(client.AppId).
			AppSecret(client.AppSecret).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Client.Auth.TenantAccessToken.Internal(context.Background(), req)

	// 处理错误
	if err != nil {
		slog.InfoContext(ctx, "飞书 GetTenantAccessToken", "error", err.Error())
		return ""
	}

	// 服务端错误处理
	if !resp.Success() {
		slog.InfoContext(ctx, "飞书 GetTenantAccessToken", "code", resp.Code, "message", resp.Msg)
		return ""
	}
	// 业务处理
	res := TenantAccessTokenResponse{}
	_ = json.Unmarshal(resp.RawBody, &res)
	client.Cache.Set("TenantAccessToken", res.TenantAccessToken, cache.WithEx(time.Duration(res.Expire)*time.Second))
	return res.TenantAccessToken
}

type AppAccessTokenResponse struct {
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	AppAccessToken string `json:"app_access_token"`
	Expire         int    `json:"expire"`
}

// GetAppAccessToken
// 获取 AppAccessToken   两小时有效
func (client *FeiShu) GetAppAccessToken(ctx context.Context) string {
	resData, ok := client.Cache.Get("AppAccessToken")
	if ok && resData != nil {
		return resData.(string)
	}
	// 创建请求对象
	req := larkauth.NewInternalAppAccessTokenReqBuilder().
		Body(larkauth.NewInternalAppAccessTokenReqBodyBuilder().
			AppId(client.AppId).
			AppSecret(client.AppSecret).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Client.Auth.AppAccessToken.Internal(context.Background(), req)

	// 处理错误
	if err != nil {
		slog.InfoContext(ctx, "飞书 GetAppAccessToken", "error", err.Error())
		return ""
	}

	// 服务端错误处理
	if !resp.Success() {
		slog.InfoContext(ctx, "飞书 GetAppAccessToken", "code", resp.Code, "message", resp.Msg)
		return ""
	}
	// 业务处理
	res := AppAccessTokenResponse{}
	_ = json.Unmarshal(resp.RawBody, &res)
	client.Cache.Set("AppAccessToken", res.AppAccessToken, cache.WithEx(time.Duration(res.Expire)*time.Second))
	return res.AppAccessToken
}

// GetChatsInfo
// 获取用户或机器人所在的所有群信息
func (client *FeiShu) GetChatsInfo(ctx context.Context) (AllChatItems map[string]larkim.ListChat) {
	// 创建请求对象
	AllChatItems = map[string]larkim.ListChat{}
	haseMore := true
	for haseMore {
		haseMore = false
		pageToken := ""
		req := larkim.NewListChatReqBuilder().
			SortType(`ByCreateTimeAsc`).
			PageToken(pageToken).
			PageSize(100).
			Build()

		// 发起请求
		resp, err := client.Client.Im.Chat.List(ctx, req)

		// 处理错误
		if err != nil {
			slog.ErrorContext(ctx, "飞书 GetChatsInfo", "error", err.Error())
			return
		}

		// 服务端错误处理
		if !resp.Success() {
			slog.InfoContext(ctx, "飞书 GetChatsInfo", "code", resp.Code, "message", resp.Msg)
			return
		}
		haseMore = *resp.Data.HasMore
		for _, d := range resp.Data.Items {
			AllChatItems[*d.Name] = *d
		}
		pageToken = *resp.Data.PageToken
	}
	return AllChatItems
}

// QueryChatId
// 更具群名称 查询群的 id
func (client *FeiShu) QueryChatId(ctx context.Context, name string) (result larkim.ListChat) {
	resData, ok := client.Cache.Get("QueryChatId")
	if ok && resData != nil {
		data := resData.(map[string]larkim.ListChat)
		if result, ok = data[name]; ok {
			return result
		}
	}
	chatInfos := client.GetChatsInfo(ctx)
	client.Cache.Set("QueryChatId", chatInfos, cache.WithEx(time.Duration(3600)*time.Second))
	if result, ok = chatInfos[name]; ok {
		return result
	}
	return
}

// SendMessage
// 发送消息
// receiveIdType chat_id 往群里发消息
// receiveId 接收者id
// msgType 发送消息类型  "text"
// content 发送的消息  `{"text":"test content"}`
func (client *FeiShu) SendMessage(ctx context.Context, receiveIdType, receiveId, msgType, content string) {
	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(receiveIdType).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(receiveId).
			MsgType(msgType).
			Content(content).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Client.Im.Message.Create(context.Background(), req)

	// 处理错误
	if err != nil {
		slog.ErrorContext(ctx, "飞书 SendMessage", "error", err.Error())
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		slog.InfoContext(ctx, "飞书 SendMessage", "code", resp.Code, "message", resp.Msg)
		return
	}
	//res, _ :=json.Marshal(resp.Data)
	slog.InfoContext(ctx, "飞书 SendMessage", "code", resp.Code, "msg", resp.Msg)
}

// SendMessageToChat
// 给指定群发送消息   注意群名称不要重复了
// name 群名
// msgType 消息类型
// content  消息数据
func (client *FeiShu) SendMessageToChat(ctx context.Context, name, msgType, content string) {
	dd := client.QueryChatId(ctx, name)
	if dd.Name != nil {
		client.SendMessage(ctx, "chat_id", *dd.ChatId, msgType, content)
	}
}

// QueryBiTableInfos
// 查询多维表数据
// fieldNames `["字段1","字段2"]`
func (client *FeiShu) QueryBiTableInfos(ctx context.Context, appToken, tableId, fieldNames string) []*larkbitable.AppTableRecord {
	// 创建请求对象
	hasMore := true
	pageToken := ""
	result := []*larkbitable.AppTableRecord{}
	for hasMore {
		hasMore = false
		req := larkbitable.NewListAppTableRecordReqBuilder().
			AppToken(appToken).
			TableId(tableId).
			PageSize(200).
			PageToken(pageToken).
			FieldNames(fieldNames).
			Build()

		// 发起请求
		resp, err := client.Client.Bitable.AppTableRecord.List(ctx, req)

		// 处理错误
		if err != nil {
			slog.ErrorContext(ctx, "error", err.Error())
			return result
		}

		// 服务端错误处理
		if !resp.Success() {
			slog.InfoContext(ctx, "code", resp.Code, "message", resp.Msg)
			return result
		}
		//fmt.Println(*resp.Data.HasMore,resp.Data.PageToken)
		hasMore = *resp.Data.HasMore
		if hasMore {
			pageToken = *resp.Data.PageToken
		}
		result = append(result, resp.Data.Items...)
	}
	return result

}
