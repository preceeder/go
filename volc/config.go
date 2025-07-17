package volc

import (
	rtcbase "github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/service/rtc"
	"net/http"
	"net/url"
)

type Config struct {
	AppKey            string `json:"appKey"` // 生成token 的时候需要
	AccessKey         string `json:"accessKey"`
	SecretKey         string `json:"secretKey"`
	AppId             string `json:"appId"`
	Region            string `json:"region"`            // 华北 cn-north-1 | 新加坡 ap-singapore-1 | 美国 us-east-1
	CallBackSecretKey string `json:"callBackSecretKey"` // 回调签名密钥, 验证密钥
}

func init() {
	rtc.ServiceInfoMap = map[string]rtcbase.ServiceInfo{
		// 华北
		"cn-north-1": {
			Timeout: rtc.DefaultTimeout,
			Scheme:  "https",
			Host:    "rtc.volcengineapi.com",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
			Credentials: rtcbase.Credentials{
				Region:  "cn-north-1",
				Service: rtc.ServiceName,
			},
		},
		// 新加坡
		"ap-singapore-1": {
			Timeout: rtc.DefaultTimeout,
			Scheme:  "https",
			Host:    "rtc.volcengineapi.com",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
			Credentials: rtcbase.Credentials{
				Region:  "ap-singapore-1",
				Service: rtc.ServiceName,
			},
		},
		"us-east-1": {
			Timeout: rtc.DefaultTimeout,
			Scheme:  "https",
			Host:    "rtc.volcengineapi.com",
			Header: http.Header{
				"Accept":       []string{"application/json"},
				"Content-Type": []string{"application/json"},
			},
			Credentials: rtcbase.Credentials{
				Region:  "us-east-1",
				Service: rtc.ServiceName,
			}},
	}
	rtc.ApiListInfo = map[string]*rtcbase.ApiInfo{
		// 封禁房间
		"BanRoomUser": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BanRoomUser"},
				"Version": []string{"2020-12-01"},
			},
			Form: map[string][]string{
				"AppId":             []string{},
				"RoomId":            []string{},
				"UserId":            []string{}, // 非必填
				"ForbiddenInterval": []string{}, // 非必填
			},
		},
		// 移出用户
		"KickUser": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"KickUser"},
				"Version": []string{"2020-12-01"},
			},
			Form: map[string][]string{
				"AppId":  []string{},
				"RoomId": []string{},
				"UserId": []string{},
			},
		},
		// 解散房间
		"DismissRoom": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"DismissRoom"},
				"Version": []string{"2020-12-01"},
			},
			Form: map[string][]string{
				"AppId":  []string{},
				"RoomId": []string{},
			},
		},
		// 获取实时用户列表 获取指定房间的实时用户列表
		"GetRoomOnlineUsers": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetRoomOnlineUsers"},
				"Version": []string{"2023-08-01"},
				"AppId":   []string{},
				"RoomId":  []string{},
			},
		},
		// 房间外点对点消息  向指定的一个应用客户端发送房间外点对点消息
		"SendUnicast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendUnicast"},
				"Version": []string{"2023-07-20"},
			},
			Form: map[string][]string{
				"AppId":   []string{},
				"From":    []string{}, // 服务商的 user_id
				"To":      []string{}, // 接收消息的user_id
				"Binary":  []string{}, //字段为 true，发送二进制消息； false，发送文本消息。
				"Message": []string{}, //点对点消息内容。如果是二进制消息，需进行 base64 编码
			},
		},
		// 发送房间内点对点消息  向指定 RTC 房间内指定的一个应用客户端发送消息
		"SendRoomUnicast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendRoomUnicast"},
				"Version": []string{"2023-07-20"},
			},
			Form: map[string][]string{
				"AppId":   []string{},
				"RoomId":  []string{},
				"From":    []string{}, // 服务商的 user_id
				"To":      []string{}, // 接收消息的user_id
				"Binary":  []string{}, //字段为 true，发送二进制消息； false，发送文本消息。
				"Message": []string{}, //点对点消息内容。如果是二进制消息，需进行 base64 编码
			},
		},
		//发送房间内广播消息 向指定一个 RTC 房间内的所有用户广播消息
		"SendBroadcast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"SendBroadcast"},
				"Version": []string{"2023-07-20"},
			},
			Form: map[string][]string{
				"AppId":   []string{},
				"RoomId":  []string{},
				"From":    []string{}, // 服务商的 user_id
				"Binary":  []string{}, //字段为 true，发送二进制消息； false，发送文本消息。
				"Message": []string{}, //点对点消息内容。如果是二进制消息，需进行 base64 编码
			},
		},
		//批量发送房间内点对点消息  向指定 RTC 房间内的批量用户发送点对点消息
		"BatchSendRoomUnicast": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"BatchSendRoomUnicast"},
				"Version": []string{"2023-07-20"},
			},
			Form: map[string][]string{
				"AppId":   []string{},
				"RoomId":  []string{},
				"From":    []string{}, // 服务商的 user_id
				"To":      []string{}, // []string{}
				"Binary":  []string{}, //字段为 true，发送二进制消息； false，发送文本消息。
				"Message": []string{}, //点对点消息内容。如果是二进制消息，需进行 base64 编码
			},
		},
	}
}
