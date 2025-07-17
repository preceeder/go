package jigou

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/preceeder/go/base"
	"github.com/zegoim/zego_server_assistant/token/go/src/token04"
	"log/slog"
	"strconv"
)

type Rtc struct {
	Config JigouConfig
	Client *resty.Client
}

func (j Rtc) GetAction(key string) string {
	return UrlMap[key]
}

func (j Rtc) Get(ctx base.BaseContext, action string, reqData any, resBody any) error {
	publicParams := GetPublicParams(j.Config)
	url := ServerUrl["rtc"]
	values, err := QueryParamsFromValues(reqData)
	if err != nil {
		return err
	}
	res, err := j.Client.R().ForceContentType("application/json").EnableTrace().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetResult(resBody).SetQueryParam("Action", action).SetQueryParams(publicParams).SetQueryParamsFromValues(values).Get(url)
	if err != nil {
		slog.ErrorContext(ctx, "jigou request", "error", err.Error(), "params", res.Request.QueryParam.Encode())
		return err
	}
	slog.InfoContext(ctx, "jigou request", "res", res.String(), "headers", res.Header())
	return nil
}

// GenerateIdentifyToken 获取 音视频流审核鉴权 Token
func (j Rtc) GenerateIdentifyToken(ctx base.BaseContext) (GenerateIdentifyToken, error) {
	resBody := &GenerateIdentifyToken{}
	err := j.Get(ctx, j.GetAction("IdentifyToken"), nil, resBody)
	return *resBody, err
}

// GetToken 获取jigou权限认证token   客户端使用
func (j Rtc) GetToken(ctx base.BaseContext, userId, roomId string) (string, error) {
	var effectiveTimeInSeconds int64 = 3600 // token 的有效时长，单位：秒
	//业务权限认证配置，可以配置多个权限位
	privilege := make(map[int]int)
	privilege[token04.PrivilegeKeyLogin] = token04.PrivilegeEnable   // 有房间登录权限
	privilege[token04.PrivilegeKeyPublish] = token04.PrivilegeEnable // 无推流权限
	//token业务扩展配置
	payloadData := &RtcRoomPayLoad{
		RoomId:       roomId,
		Privilege:    privilege,
		StreamIdList: nil,
	}
	payload, err := json.Marshal(payloadData)
	if err != nil {
		slog.ErrorContext(ctx, "GetToken error", "error", err.Error())
		return "", err
	}
	//生成token
	appid, _ := strconv.Atoi(j.Config.AppId)
	token, err := token04.GenerateToken04(uint32(appid), userId, j.Config.ServerSecret, effectiveTimeInSeconds, string(payload))
	if err != nil {
		fmt.Println(err)
	}
	return token, nil
}
