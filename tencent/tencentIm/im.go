package tencentIm

import (
	"github.com/go-resty/resty/v2"
	"github.com/preceeder/go/base"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"tencent/tencentIm/ECDSASHA256"
	"tencent/tencentIm/HMACSHA256"
	"time"
)

type TencentImClient struct {
	Config   TencentImConfig
	Client   *resty.Client
	UserSign string // 管理员的有效 usersign
}
type TencentImConfig struct {
	Prefix     string `json:"prefix"`     // im id前缀
	AppId      int    `json:"appId"`      // appid
	Identifier string `json:"identifier"` // 管理员账户
	Key        string `json:"Key"`        // 密钥   HMAC-SHA256 算法 使用
	PrivateKey string `json:"privateKey"` // 私钥   ECDSA-SHA256 算法 加密 使用
	PublicKey  string `json:"publicKey"`  // 公钥    ECDSA-SHA256 算法 验证 使用
	UseSha     string `json:"useSha"`     // 使用那种算法  HMAC-SHA256｜ ECDSA-SHA256
	ImHost     string `json:"imHost"`     // 域名  最后不要 /
	Expire     int    `json:"expire"`     // token 过期时间 s
}

func NewTencentIm(config TencentImConfig) TencentImClient {
	client := InitClient()
	return TencentImClient{Client: client, Config: config}
}

func InitClient() *resty.Client {
	imClient := resty.New()
	imClient.SetTimeout(3 * time.Second)
	imClient.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	})

	imClient.SetTransport(&http.Transport{
		MaxIdleConnsPerHost:   50,               // 对于每个主机，保持最大空闲连接数为 10
		IdleConnTimeout:       30 * time.Second, // 空闲连接超时时间为 30 秒
		TLSHandshakeTimeout:   3 * time.Second,  // TLS 握手超时时间为 10 秒
		ResponseHeaderTimeout: 3 * time.Second,  // 等待响应头的超时时间为 3 秒
	})
	return imClient
}

// respBody 必须是指针

func (t TencentImClient) SendImRequest(ctx base.BaseContext, serverName string, requestData any, respBody any) error {
	req := t.Client.R().SetBody(requestData)
	if respBody != nil {
		req.SetResult(respBody)
	}
	durl := t.setUrl(serverName)
	_, err := req.Post(durl)
	if err != nil {
		slog.ErrorContext(ctx, "SendImRequest error", "error", err.Error(), "serverName", serverName, "data", requestData)
		return err
	}
	return nil
}

func (t TencentImClient) setUrl(serverName string) string {
	if uri, ok := ApiMap[serverName]; ok {
		query := url.Values{}
		query.Set("contenttype", "json")
		query.Set("sdkappid", strconv.Itoa(t.Config.AppId))
		query.Set("identifier", t.Config.Identifier)
		sign := ""
		if t.Config.UseSha == "HMAC-SHA256" {
			sign = t.getHmacSign()
		} else {
			sign = t.getEcdsaSign()
		}
		query.Set("usersig", sign)
		query.Set("random", RandStrInt(5))
		bp, _ := url.JoinPath(t.Config.ImHost, uri)
		return bp + "?" + query.Encode()
	} else {
		slog.Error("im server not find in ApiMap", "serverName", serverName)
	}
	return ""
}

func (t *TencentImClient) getHmacSign() string {
	var err error
	var userSignValid = false
	if len(t.UserSign) > 10 {
		err = HMACSHA256.VerifyUserSig(uint64(t.Config.AppId), t.Config.Key, t.Config.Identifier, t.UserSign, time.Now())
		if err != nil {
			slog.Error("im usersign error", "error", err.Error())
		}
		userSignValid = true
	}
	if !userSignValid {
		t.UserSign, err = HMACSHA256.GenUserSig(t.Config.AppId, t.Config.Key, t.Config.Identifier, t.Config.Expire)
		if err != nil {
			slog.Error("生成im usersign error", "error", err.Error())
		}
	}
	return t.UserSign
}

func (t *TencentImClient) getEcdsaSign() string {
	var err error
	var userSignValid = false
	if len(t.UserSign) > 10 {
		err = ECDSASHA256.VerifyUsersig(t.Config.PublicKey, t.UserSign, t.Config.AppId, t.Config.Identifier)
		if err != nil {
			slog.Error("im usersign error", "error", err.Error())
		}
		userSignValid = true
	}
	if !userSignValid {
		t.UserSign, err = ECDSASHA256.GenerateUsersigWithExpire(t.Config.PrivateKey, t.Config.AppId, t.Config.Identifier, int64(t.Config.Expire))
		if err != nil {
			slog.Error("生成im usersign error", "error", err.Error())
		}
	}
	return t.UserSign
}

// 获取用户的 token
func (t TencentImClient) GetUserSign(userId string) (userSign string, err error) {
	if t.Config.UseSha == "ECDSA-SHA256" {
		userSign, err = ECDSASHA256.GenerateUsersigWithExpire(t.Config.PrivateKey, t.Config.AppId, userId, int64(t.Config.Expire))
		if err != nil {
			slog.Error("生成im usersign error", "error", err.Error())
		}
	} else if t.Config.UseSha == "HMAC-SHA256" {
		userSign, err = HMACSHA256.GenUserSig(t.Config.AppId, t.Config.Key, userId, t.Config.Expire)
		if err != nil {
			slog.Error("生成im usersign error", "error", err.Error())
		}
	}
	return
}
