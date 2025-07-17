package jigou

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// ServerUrl 服务的 url
var ServerUrl = map[string]string{
	"rtc":         "https://rtc-api.zego.im",         // rtc 接口
	"cloudrecord": "https://cloudrecord-api.zego.im", // 云端录制
}
var UrlMap = map[string]string{
	"BroadcastMessage": "SendBroadcastMessage",     // 推送广播消息
	"BarrageMessage":   "SendBarrageMessage",       // 推送弹幕消息
	"CustomMessage":    "SendCustomCommand",        // 推送自定义消息
	"UserNumber":       "DescribeUserNum",          // 查询房间人数
	"UserStatus":       "DescribeUsers",            // 查询用户状态
	"UserList":         "DescribeUserList",         // 获取房间用户列表
	"AddStream":        "AddStream",                // 增加房间流
	"DeleteStream":     "DeleteStream",             // 删除流
	"SimpleStreamList": "DescribeSimpleStreamList", // 获取房间内简易流列表
	"CloseRoom":        "CloseRoom",                // 关闭房间
	"KickoutUser":      "KickoutUser",              // 踢出房间用户
	"IdentifyToken":    "GenerateIdentifyToken",    // 获取音视频流审核鉴权token
}

type JigouConfig struct {
	Prefix         string `json:"prefix"`
	AppId          string `json:"appId"`
	ServerSecret   string `json:"serverSecret"`
	CallBackSecret string `json:"callbackSecret"`
	CdnLive        string `json:"cdnLive"`       // cdn推流的接入点（接入名称）
	CdnType        string `json:"cdnType"`       // cdn的运营商   tx(腾讯)｜ hv(华为)｜ ws(网宿)
	CdnPublishUrl  string `json:"cdnPublishUrl"` // cdn 推流的url
	CdnLaUrl       string `json:"cdnLaUrl"`      // cdn拉流的url, 不要给协议头， rtmp://, http://   这种都不要给
	CdnKey         string `json:"cdnKey"`        // cdn推流的鉴权key
}

// 生成签名
// Signature=md5(AppId + SignatureNonce + ServerSecret + Timestamp)
func generateSignature(appId string, serverSecret, signatureNonce string, timeStamp int64) string {
	data := fmt.Sprintf("%s%s%s%d", appId, signatureNonce, serverSecret, timeStamp)
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func GetPublicParams(config JigouConfig) map[string]string {
	publicParams := map[string]string{}
	timestamp := time.Now().Unix()
	// 生成16进制随机字符串(16位)
	nonce := make([]byte, 8)
	_, _ = rand.Read(nonce)
	hexNonce := hex.EncodeToString(nonce)
	// 生成签名
	signature := generateSignature(config.AppId, config.ServerSecret, hexNonce, timestamp)
	publicParams["AppId"] = config.AppId
	//公共参数中的随机数和生成签名的随机数要一致
	publicParams["SignatureNonce"] = hexNonce
	publicParams["SignatureVersion"] = "2.0"
	//公共参数中的时间戳和生成签名的时间戳要一致
	publicParams["Timestamp"] = fmt.Sprintf("%d", timestamp)
	publicParams["Signature"] = signature
	return publicParams
}
