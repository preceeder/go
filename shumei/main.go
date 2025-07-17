package shumei

import (
	"encoding/json"
	mapset "github.com/deckarep/golang-set"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/go-resty/resty/v2"
	"github.com/preceeder/go/base"
	"log/slog"
	"net/url"
	"strings"
	"time"
)

var ShumeiClient *ShuMei

var DefauiltHttpClient = NewRequestClient()
var imageLangSet = mapset.NewSet("zh", "en", "ar")
var textLangSet = mapset.NewSet("zh", "en", "ar", "hi", "es", "fr", "ru", "pt", "id", "de", "ja", "tr", "vi", "it", "th", "tl", "ko", "ms", "auto")
var voiceLangSet = mapset.NewSet("zh", "en", "ar", "hi", "es", "fr", "ru", "pt", "id", "de", "ja", "tr", "vi", "it", "th", "tl", "ko", "ms")

type ShumeiUrl struct {
	VideoStreamCloseUrl string `json:"videoStreamCloseUrl"` // 视频流检查关闭的数美url
	VoiceStreamCloseUrl string `json:"voiceStreamCloseUrl"` // 语音流检查关闭的数美url
	ImageUrl            string `json:"imageUrl"`            // 单张图片的检查的数美url
	MultiImageUrl       string `json:"multiImageUrl"`       // 同步多张图片的检查的数美url
	TextUrl             string `json:"textUrl"`             // 文本检查的数美url
	VoiceUrl            string `json:"voiceUrl"`            // 语音文件检查的数美url  同步返回结果,   只支持国内, 国外不支持
	AsyncVoiceUrl       string `json:"asyncVoiceUrl"`       // 异步语音文件检查的数美url
	AsyncVideoUrl       string `json:"asyncVideoUrl"`       // 视频文件检查的数美url
	VoiceStreamUrl      string `json:"voiceStreamUrl"`      // 音频流检查url
	VideoStreamUrl      string `json:"videoStreamUrl"`      // 视频流检查url
}
type CallBackUrl struct {
	ImageCallBackUrl       string `json:"imageCallBackUrl"`       // 图片回调 url
	MultiImageCallBackUrl  string `json:"multiImageCallBackUrl"`  // 图片回调 url
	VoiceCallBackUrl       string `json:"voiceCallBackUrl"`       // 语音文件检查回调url
	VideoCallBackUrl       string `json:"videoCallBackUrl"`       // 视频文件检查回调url
	VoiceStreamCallBackUrl string `json:"voiceStreamCallBackUrl"` // 音频流回调url
	VideoStreamCallBackUrl string `json:"videoStreamCallBackUrl"` // 视频流回调url
}

type ShumeiConfig struct {
	AppId          string      `json:"appid"`
	AccessKey      string      `json:"accessKey"`
	CdnUrl         string      `json:"cdnUrl"`         // cdn的url
	TokenPrefix    string      `json:"tokenPrefix"`    // 用户token的前缀
	CallBackDomain string      `json:"callBackDomain"` // 回调的 url
	ShumeiUrl      ShumeiUrl   `json:"shumeiUrl"`
	CallBackUrl    CallBackUrl `json:"callBackUrl"`
}

func NewShumeiClient(config ShumeiConfig) *ShuMei {
	return initShumei(config)
}

func initShumei(config ShumeiConfig) *ShuMei {
	client, err := NewShuMei(config.AppId, config.AccessKey,
		OptionWithTokenPrefix(config.TokenPrefix),
		OptionWithCdnUrl(config.CdnUrl),
		OptionWithCallBackDomain(config.CallBackDomain))
	if err != nil {
		slog.Error("数美初始化失败", "error", err.Error())
		panic("数美初始化失败")
	}
	ShumeiClient = client
	ShumeiClient.ShumeiUrl = config.ShumeiUrl
	ShumeiClient.CallBackUrl = config.CallBackUrl
	return ShumeiClient
}

//// 使用 viper读取的配置初始化
//func InitShumeiWithViperConfig(config viper.Viper) {
//	shumeisConfig := ShumeiConfig{}
//	utils.ReadViperConfig(config, "shumei", &shumeisConfig)
//	initShumei(shumeisConfig)
//}

type ShuMei struct {
	AppId            string
	AccessKey        string
	DefaultImageType string // 默认值 POLITICS_PORN_AD
	DefaultTextType  string // 默认值 AD
	DefaultVoiceType string // 默认值 PORN_MOAN_AD
	DefaultVideoType string // 默认值 POLITY_EROTIC_ADVERT
	TokenPrefix      string // 用户id的统一前缀
	HttpClient       *resty.Client
	CdnUrl           string      // 资源的 url 最后不加 /
	CallBackDomain   string      // 回调域名  url 最后不加 /
	ShumeiUrl        ShumeiUrl   // 数美接口的urls
	CallBackUrl      CallBackUrl // 回调url
}

func NewShuMei(appId string, accessKey string, optionals ...func(*ShuMei) error) (*ShuMei, error) {
	tp := DefauiltHttpClient
	sh := &ShuMei{
		AppId:      appId,
		AccessKey:  accessKey,
		HttpClient: tp,
	}
	for _, op := range optionals {
		err := op(sh)
		if err != nil {
			return nil, err
		}
	}
	if sh.DefaultImageType == "" {
		sh.DefaultImageType = "POLITICS_PORN_AD"
	}
	if sh.DefaultVoiceType == "" {
		sh.DefaultVoiceType = "PORN_MOAN_AD"
	}

	if sh.DefaultTextType == "" {
		sh.DefaultTextType = "AD"
	}
	if sh.DefaultVideoType == "" {
		sh.DefaultVideoType = "POLITY_EROTIC_ADVERT"
	}

	return sh, nil
}

func OptionWithTokenPrefix(t string) func(*ShuMei) error {
	return func(s *ShuMei) error {
		s.TokenPrefix = t
		return nil
	}
}

func OptionWithCdnUrl(t string) func(mei *ShuMei) error {
	return func(s *ShuMei) error {
		s.CdnUrl = t
		return nil
	}
}

func OptionWithCallBackDomain(t string) func(mei *ShuMei) error {
	return func(s *ShuMei) error {
		s.CallBackDomain = t
		return nil
	}
}
func OptionWithUrl(t ShumeiUrl) func(mei *ShuMei) error {
	return func(s *ShuMei) error {
		s.ShumeiUrl = t
		return nil
	}
}

func OptionWithImageType(t string) func(*ShuMei) error {
	return func(s *ShuMei) error {
		s.DefaultImageType = t
		return nil
	}
}

func OptionWithTextType(t string) func(*ShuMei) error {
	return func(s *ShuMei) error {
		s.DefaultTextType = t
		return nil
	}
}

func (s ShuMei) Send(url string, body any, response any) (map[string]any, error) {
	res, err := s.HttpClient.R().
		SetResult(response).
		ForceContentType("application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		EnableTrace().
		SetBody(body).
		Post(url)
	if response == nil {
		var data map[string]interface{}
		_ = json.Unmarshal(res.Body(), &data)
		return data, err
	} else {
		_ = json.Unmarshal(res.Body(), response)
		return nil, err
	}
}

func (s ShuMei) urlHandler(imageUrl string) string {
	if !(strings.HasPrefix(imageUrl, "http://") || strings.HasPrefix(imageUrl, "https://")) {
		imageUrl, _ = url.JoinPath(s.CdnUrl, imageUrl)
	}
	return imageUrl
}

func (s ShuMei) tokenHandler(userId string) string {
	if userId == "" {
		return ""
	}
	if !strings.HasPrefix(userId, s.TokenPrefix) {
		userId = s.TokenPrefix + userId
	}
	return userId
}

func (s ShuMei) voiceLangeHandler(lang string) string {
	if !voiceLangSet.Contains(lang) {
		lang = "zh"
	}
	return lang
}

func (s ShuMei) imageLangHandler(lang string) string {
	if !imageLangSet.Contains(lang) {
		lang = "zh"
	}
	return lang
}

func (s ShuMei) textLangHandler(lang string) string {
	if !textLangSet.Contains(lang) {
		lang = "auto"
	}
	return lang
}

// 回调路径处理
func (s ShuMei) HandlerCallBackUrl(urlStr string) string {
	if len(urlStr) == 0 {
		return ""
	}
	if !(strings.HasPrefix(urlStr, "http://") || strings.HasPrefix(urlStr, "https://")) {
		urlStr, _ = url.JoinPath(s.CallBackDomain, urlStr)
	}
	return urlStr
}

// 传了callBackUrl  就是走回调
func (s ShuMei) Image(ctx base.BaseContext, p ShumeiImage) (bool, *PublicLongResponse) {
	//turl := "http://api-img-xjp.fengkongcloud.com/image/v4"
	turl := s.ShumeiUrl.ImageUrl //  "http://api-img-sh.fengkongcloud.com/image/v4"

	data := map[string]interface{}{
		"img":            s.urlHandler(p.ImageUrl),
		"tokenId":        s.tokenHandler(p.UserId),
		"receiveTokenId": s.tokenHandler(p.ReceiveTokenId),
		"lang":           s.imageLangHandler(p.Lang),
	}
	if p.Ip != "" {
		data["ip"] = p.Ip
	}

	if p.MType == "" {
		p.MType = s.DefaultImageType
	}

	if p.EventId == "" {
		p.EventId = "IMAGE"
	}

	if p.BusinessType == "" {
		p.BusinessType = "FACE"
	}

	data["extra"] = map[string]any{
		"passThrough": p.ThroughParams,
	}

	payload := map[string]interface{}{
		"accessKey":    s.AccessKey,
		"type":         p.MType,
		"eventId":      p.EventId,
		"businessType": p.BusinessType,
		"appId":        s.AppId,
		"data":         data,
	}
	if p.NeedCallBack {
		payload["callback"] = s.CallBackUrl.ImageCallBackUrl
		if len(p.CallBaskUrl) > 0 {
			payload["callback"] = s.HandlerCallBackUrl(p.CallBaskUrl)
		}
	}

	res := &PublicLongResponse{}
	_, err := s.Send(turl, payload, res)
	if err != nil {
		slog.ErrorContext(ctx, "shumei image request", "error", err.Error())
		return true, nil
	}

	if res.Code == 1100 {
		if res.RiskLevel == "REJECT" {
			return false, res
		}
	}
	return true, res
}

// 多张图片同时检查
//
//	传了callBackUrl  就是走回调
func (s ShuMei) MultiImage(ctx base.BaseContext, p ShumeiMultiImage) (bool, *PublicLongResponse) {
	turl := s.ShumeiUrl.MultiImageUrl //  "http://api-img-sh.fengkongcloud.com/image/v4"

	data := map[string]interface{}{
		"tokenId":        s.tokenHandler(p.UserId),
		"receiveTokenId": s.tokenHandler(p.ReceiveTokenId),
		"lang":           s.imageLangHandler(p.Lang),
	}

	imgs := []map[string]string{}
	btiImagemap := map[string]string{}
	for _, img := range p.ImageUrl {
		btid := cryptor.Md5String(img)[8:]
		imgs = append(imgs, map[string]string{
			"img":  s.urlHandler(img),
			"btId": btid,
		})
		btiImagemap[btid] = img
	}
	data["imgs"] = imgs
	if p.Ip != "" {
		data["ip"] = p.Ip
	}

	if p.MType == "" {
		p.MType = s.DefaultImageType
	}
	if p.EventId == "" {
		p.EventId = "IMAGE"
	}

	if p.ThroughParams == nil {
		p.ThroughParams = map[string]any{
			"btIdMap": btiImagemap,
		}
	} else {
		p.ThroughParams["btIdMap"] = btiImagemap
	}

	data["extra"] = map[string]any{
		"passThrough": p.ThroughParams,
	}
	payload := map[string]interface{}{
		"accessKey":    s.AccessKey,
		"appId":        s.AppId,
		"eventId":      p.EventId,
		"type":         p.MType,
		"businessType": "FACE",
		"data":         data,
	}

	if p.NeedCallBack {
		payload["callback"] = s.CallBackUrl.MultiImageCallBackUrl
		if len(p.CallBaskUrl) > 0 {
			payload["callback"] = s.HandlerCallBackUrl(p.CallBaskUrl)
		}
	}

	res := &PublicLongResponse{}
	_, err := s.Send(turl, payload, res)
	if err != nil {
		slog.ErrorContext(ctx, "shumei image request", "error", err.Error(), "error", err.Error())
		return false, nil
	}

	if res.Code != 1100 {
		slog.ErrorContext(ctx, "AsyncVoiceFile", "error", res.Message, "requestId", res.RequestID, "code", res.Code)
		return false, res
	}
	return true, res
}

func (s ShuMei) Text(ctx base.BaseContext, p ShumeiText) (bool, *PublicLongResponse) {
	turl := s.ShumeiUrl.TextUrl //"http://api-text-sh.fengkongcloud.com/text/v4"
	data := map[string]interface{}{
		"text":     p.Text,
		"tokenId":  s.tokenHandler(p.UserId),
		"lang":     s.textLangHandler(p.Lang),
		"ip":       p.Ip,
		"deviceId": p.DeviceId,
	}
	if p.EventId == "" {
		p.EventId = "text"
	}
	if p.EventId == "message" {
		data["extra"] = map[string]any{"receiveTokenId": s.tokenHandler(p.ReceiveTokenId)}
	}

	if p.MType == "" {
		p.MType = s.DefaultTextType
	}

	payload := map[string]interface{}{
		"accessKey": s.AccessKey,
		"appId":     s.AppId,
		"eventId":   p.EventId,
		"type":      p.MType,
		"data":      data,
	}
	res := &PublicLongResponse{}
	_, err := s.Send(turl, payload, res)
	if err != nil {
		slog.ErrorContext(ctx, "shumei image request", "error", err.Error())
		return true, nil
	}
	if res.Code == 1100 {
		if res.RiskLevel == "REJECT" {
			return false, res
		}
	}
	return true, res
}

// 只支持 url的 同步 国内支持, 国外不支持
func (s ShuMei) VoiceFile(ctx base.BaseContext, p ShumeiVoiceFile) (bool, *VoiceFileResponse) {
	turl := s.ShumeiUrl.VoiceUrl // "http://api-audio-sh.fengkongcloud.com/audiomessage/v4"
	data := map[string]interface{}{
		"tokenId": s.tokenHandler(p.UserId),
	}
	if p.EventId == "" {
		p.EventId = "default"
	}
	if p.MType == "" {
		p.MType = s.DefaultVoiceType
	}

	payload := map[string]interface{}{
		"accessKey":   s.AccessKey,
		"appId":       s.AppId,
		"eventId":     p.EventId,
		"type":        p.MType,
		"contentType": "URL",
		"content":     s.urlHandler(p.VoiceUrl),
		"data":        data,
		"btId":        RandStr(16),
	}

	res := &VoiceFileResponse{}
	_, err := s.Send(turl, payload, res)
	if err != nil {
		slog.ErrorContext(ctx, "shumei image request", "error", err.Error())
		return true, nil
	}
	if res.Code == 1100 {
		if res.Detail.RiskLevel == "REJECT" {
			return false, res
		}
	}
	return true, res
}

// AsyncVoiceFile 异步语音文件检查 国外只支持异步的,   回调的时候透传参数是 requestParams,  这里main装的是 此处data参数里的所有数据
func (s ShuMei) AsyncVoiceFile(ctx base.BaseContext, p ShumeiVoiceFile) bool {
	turl := s.ShumeiUrl.AsyncVoiceUrl //"http://api-audio-sh.fengkongcloud.com/audio/v4"
	data := map[string]any{
		"tokenId":     s.tokenHandler(p.UserId),
		"lang":        s.voiceLangeHandler(p.Lang),
		"passThrough": p.CallbackParams,
	}

	if p.EventId == "" {
		p.EventId = "default"
	}
	if p.MType == "" {
		p.MType = s.DefaultVoiceType
	}

	payload := map[string]interface{}{
		"accessKey":   s.AccessKey,
		"appId":       s.AppId,
		"eventId":     p.EventId,
		"type":        p.MType,
		"contentType": "URL",
		"content":     s.urlHandler(p.VoiceUrl),
		"data":        data,
		"btId":        RandStr(16),
		"callback":    s.CallBackUrl.VoiceCallBackUrl,
	}

	if p.CallbackUrl != "" {
		payload["callback"] = s.HandlerCallBackUrl(p.CallbackUrl)
	}

	res := &PublicShortResponse{}
	_, err := s.Send(turl, payload, res)
	if err != nil {
		slog.ErrorContext(ctx, "shumei image request", "error", err.Error())
		return false
	}
	if res.Code != 1100 {
		slog.ErrorContext(ctx, "AsyncVoiceFile", "error", res.Message, "requestId", res.RequestID, "code", res.Code)
		return false
	}
	return true
}

// AsyncVideoFile 异步视频文件检查  国外只支持异步的
func (s ShuMei) AsyncVideoFile(ctx base.BaseContext, p ShumeiAsyncVideoFile) (bool, *VideoFileResponse) {
	//上海节点
	turl := s.ShumeiUrl.AsyncVideoUrl // "http://api-video-sh.fengkongcloud.com/video/v4"
	data := map[string]interface{}{
		"tokenId": s.tokenHandler(p.UserId),
		"lang":    s.imageLangHandler(p.Lang),
		"btId":    RandStr(16),
		"url":     s.urlHandler(p.VideoUrl),
		"extra":   map[string]any{"passThrough": p.ThroughParams},
	}

	if p.EventId == "" {
		p.EventId = "default"
	}

	if p.VideoType == "" {
		p.VideoType = s.DefaultVideoType
	}

	if p.VoiceType == "" {
		p.VoiceType = s.DefaultVoiceType
	}

	payload := map[string]interface{}{
		"accessKey": s.AccessKey,
		"appId":     s.AppId,
		"eventId":   p.EventId,
		"imgType":   p.VideoType,
		"audioType": p.VoiceType,
		"callback":  s.CallBackUrl.VideoCallBackUrl,
		"data":      data,
	}

	if p.CallBackUrl != "" {
		payload["callback"] = s.HandlerCallBackUrl(p.CallBackUrl)
	}

	res := &VideoFileResponse{}
	_, err := s.Send(turl, payload, res)
	if err != nil {
		slog.ErrorContext(ctx, "shumei image request", "error", err.Error())
		return false, nil
	}
	if res.Code != 1100 {
		slog.ErrorContext(ctx, "AsyncVoiceFile", "error", res.Message, "requestId", res.RequestID, "code", res.Code)
		return false, res
	}
	return true, res
}

// 音频流检查
func (s ShuMei) AudioStream(ctx base.BaseContext, p ShumeiAsyncAudioStream) (bool, *AudioStreamResponse) {
	turl := s.ShumeiUrl.VoiceStreamUrl //"http://api-audiostream-sh.fengkongcloud.com/audiostream/v4"
	data := map[string]interface{}{
		"tokenId":          s.tokenHandler(p.UserId),
		"lang":             s.voiceLangeHandler(p.Lang),
		"btId":             RandStr(16),
		"streamType":       p.StreamType,
		"returnAllText":    p.ReturnAllText,
		"room":             p.RoomId,
		"returnFinishInfo": p.ReturnFinishInfo,
		"audioDetectStep":  p.AudioDetectStep,
		"extra":            map[string]any{"passThrough": p.ThroughParams},
	}

	for k, v := range p.RtcParams {
		data[k] = v
	}

	if p.EventId == "" {
		p.EventId = "default"
	}

	if p.VoiceType == "" {
		p.VoiceType = s.DefaultVoiceType
	}

	payload := voiceStreamRequest{
		AccessKey:    s.AccessKey,
		AppId:        s.AppId,
		EventId:      p.EventId,
		Type:         p.VoiceType,
		BusinessType: p.BusinessType,
		Callback:     s.HandlerCallBackUrl(p.Callback),
		Data:         data,
	}

	res := &AudioStreamResponse{}
	_, err := s.Send(turl, payload, res)
	if err != nil {
		slog.ErrorContext(ctx, "shumei AsyncVoiceFile request", "error", err.Error())
		return false, nil
	}
	if res.Code != 1100 {
		slog.ErrorContext(ctx, "AudioStream", "error", res.Message, "requestData", payload, "requestId", res.RequestID, "code", res.Code)
		return false, nil
	} else if res.Code == 1100 && res.Detail.Errorcode != 0 {
		slog.ErrorContext(ctx, "AsyncVoiceFile", "error", res.Message, "requestId", res.RequestID, "errorCode", res.Detail.Errorcode)
		return false, nil
	}

	return true, res
}

// 视频流检查
func (s ShuMei) VideoStream(ctx base.BaseContext, p ShumeiAsyncVideoStream) (bool, *AudioStreamResponse) {
	turl := s.ShumeiUrl.VideoStreamUrl //"http://api-videostream-sh.fengkongcloud.com/videostream/v4"
	data := map[string]interface{}{
		"tokenId":          s.tokenHandler(p.UserId),
		"lang":             s.imageLangHandler(p.Lang),
		"btId":             RandStr(16),
		"streamType":       p.StreamType,
		"room":             p.RoomId,
		"returnFinishInfo": p.ReturnFinishInfo,
		"detectFrequency":  p.DetectFrequency, // 通知的频次   秒/次
		"returnAllImg":     p.ReturnAllImg,
		"returnAllText":    p.ReturnAllText,
		//"audioDetectStep":  20,
		"extra": map[string]any{"passThrough": p.ThroughParams},
	}

	if p.DetectStep > 0 {
		data["detectStep"] = p.DetectStep
	}
	// rtc 参数
	for k, v := range p.RtcParams {
		data[k] = v
	}
	if p.EventId == "" {
		p.EventId = "default"
	}

	if p.VideoType == "" {
		p.VideoType = s.DefaultVideoType
	}
	payload := videoStreamRequest{
		AccessKey:         s.AccessKey,
		AppId:             s.AppId,
		EventId:           p.EventId,
		ImgType:           p.VideoType,
		AudioType:         p.VoiceType,
		ImgCallback:       s.HandlerCallBackUrl(p.ImgCallback),
		AudioCallback:     s.HandlerCallBackUrl(p.AudioCallback),
		Data:              data,
		ImgBusinessType:   p.ImgBusinessType,
		AudioBusinessType: p.AudioBusinessType,
	}

	res := &AudioStreamResponse{}

	_, err := s.Send(turl, payload, res)
	if err != nil {
		slog.ErrorContext(ctx, "shumei AsyncVoiceFile request", "error", err.Error())
		return false, nil
	}
	if res.Code != 1100 {
		slog.ErrorContext(ctx, "AudioStream", "error", res.Message, "requestId", res.RequestID, "requestData", payload, "code", res.Code)
		return false, nil
	} else if res.Code == 1100 && res.Detail.Errorcode != 0 {
		slog.ErrorContext(ctx, "AsyncVoiceFile", "error", res.Message, "requestId", res.RequestID, "errorCode", res.Detail.Errorcode)
		return false, nil
	}

	return true, res
}

/** 流关闭接口
 * @param requestId string 请求id
 * @ltype string 类型 voice｜video
 */
func (s ShuMei) CloseStreamCheck(ctx base.BaseContext, requestId string, ltype string) (bool, *CloseStreamResponse) {
	turl := ""
	if ltype == "video" {
		turl = s.ShumeiUrl.VideoStreamCloseUrl
	} else if ltype == "voice" {
		turl = s.ShumeiUrl.VoiceStreamCloseUrl
	}

	payload := map[string]interface{}{
		"accessKey": s.AccessKey,
		"requestId": requestId,
	}

	res := &CloseStreamResponse{}
	_, err := s.Send(turl, payload, res)
	if err != nil {
		slog.ErrorContext(ctx, "shumei close stream faild", "error", err.Error())
		return false, nil
	}
	return true, res

}

func (s ShuMei) TianWang(ctx base.BaseContext, p TianWangParams) map[string]any {
	data := map[string]any{
		"accessKey": s.AccessKey,
		"appId":     s.AppId,
		"eventId":   p.EventId,
		"data": map[string]any{
			"tokenId":    s.tokenHandler(p.TokenId),
			"ip":         p.Ip,
			"timestamp":  time.Now().UnixMilli(),
			"deviceId":   p.SmDeviceId,
			"phone":      p.Phone,
			"os":         p.Channel,
			"appVersion": p.Version,
			"type":       p.RegisterMethod,
			"userAgent":  p.UserAgent,
		},
	}
	response, err := s.Send("http://api-skynet-bj.fengkongcloud.com/v4/event", data, nil)
	if err != nil {
		slog.ErrorContext(ctx, "tianwang 事件接口访问失败", "error", err.Error())
	}
	return response
}
