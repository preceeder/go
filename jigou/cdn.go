package jigou

import (
	"github.com/duke-git/lancet/v2/cryptor"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// GenerateCdnTxtSecret 腾讯云cdn
// return txSecret, txTime
func (j Rtc) generateCdnTxSecret(streamId string, expire int) string {
	if expire == 0 {
		// 默认给24小时的过期时间
		expire = 3600 * 24
	}
	// 将时间转换成16进制
	txTime := strconv.FormatInt(time.Now().Add(time.Duration(expire)*time.Second).Unix(), 16)
	param := url.Values{}
	param.Set("txSecret", cryptor.Md5String(strings.Join([]string{j.Config.CdnKey, streamId, txTime}, "")))
	param.Set("txTime", txTime)
	return param.Encode()
}

// GenerateCdnWsSecret, 即构自己的cdn就用这个
// return WsSecret, WsTime
func (j Rtc) generateCdnWsSecret(streamId string, expire int) string {
	if expire == 0 {
		// 默认给24小时的过期时间
		expire = 3600 * 24
	}
	// 将时间转换成16进制
	wsTime := strconv.FormatInt(time.Now().Add(time.Duration(expire)*time.Second).Unix(), 16)
	streamName := "/" + j.Config.CdnLive + "/" + streamId
	param := url.Values{}
	param.Set("wsSecret", cryptor.Md5String(strings.Join([]string{wsTime, streamName, j.Config.CdnKey}, "")))
	param.Set("wsABStime", wsTime)
	return param.Encode()
}

// GenerateCdnHvSecret, 华为云cdn
// return HvSecret, HvTime
func (j Rtc) generateCdnHvSecret(streamId string, expire int) string {
	if expire == 0 {
		// 默认给24小时的过期时间
		expire = 3600 * 24
	}
	// 将时间转换成16进制
	hwTime := strconv.FormatInt(time.Now().Add(time.Duration(expire)*time.Second).Unix(), 16)
	param := url.Values{}
	param.Set("hwSecret", cryptor.HmacSha256(j.Config.CdnKey, strings.Join([]string{hwTime, streamId}, "")))
	param.Set("hwTime", hwTime)
	return param.Encode()
}

// 获取完整的cdn推流url
func (j Rtc) GetPublishUrl(streamId string, expire int) string {
	// 将时间转换成16进制
	authParams := ""
	switch j.Config.CdnType {
	case "ws":
		authParams = j.generateCdnWsSecret(streamId, expire)
	case "hv":
		authParams = j.generateCdnHvSecret(streamId, expire)
	case "tx":
		authParams = j.generateCdnTxSecret(streamId, expire)
	}

	uri, _ := url.JoinPath(j.Config.CdnPublishUrl, j.Config.CdnLive, streamId)
	if authParams != "" {
		uri += "?" + authParams
	}
	return uri
}

// 获取cdn rtmp协议的 拉流地址
func (j Rtc) GetRtmpPullUrl(streamId string) string {
	uri, _ := url.JoinPath("rtmp://"+j.Config.CdnLaUrl, j.Config.CdnLive, streamId)
	return uri
}

// 获取cdn rtmp协议的 拉流地址
func (j Rtc) GetHdlPullUrl(streamId string) string {
	uri, _ := url.JoinPath("http://"+j.Config.CdnLaUrl, j.Config.CdnLive, streamId)
	return uri + ".flv"
}

// 获取cdn rtmp协议的 拉流地址
func (j Rtc) GetHlsPullUrl(streamId string) string {
	uri, _ := url.JoinPath("http://"+j.Config.CdnLaUrl, j.Config.CdnLive, streamId, "playlist.m3u8")
	return uri
}
