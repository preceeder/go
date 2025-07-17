package tencentIm

import (
	"encoding/json"
	"errors"
	"log/slog"
)

type MsgContent interface {
	GetMsgType() string
	GetData() any
}

// 文本类型
type TextContent struct {
	Text string `json:"Text"`
}

func (c TextContent) GetMsgType() string {
	return "TIMTextElem"
}
func (c TextContent) GetData() any {
	return c
}

// 自定义类型
type CustomContent struct {
	Data any `json:"Data"`
	//自定义消息描述信息。当接收方为 iOS 或 Android 后台在线时，做离线推送文本展示。
	//若发送自定义消息的同时设置了OfflinePushInfo.Desc字段，此字段会被覆盖，请优先填 OfflinePushInfo.Desc 字段。
	//说明：当消息中只有一个 TIMCustomElem 自定义消息元素时，如果 Desc 字段和 OfflinePushInfo.Desc 字段都不填写，将收不到该条消息的离线推送，需要填写 OfflinePushInfo.Desc 字段才能收到该消息的离线推送。
	Desc string `json:"Desc,omitempty"`
	//扩展字段。当接收方为 iOS 系统且应用处在后台时，此字段作为 APNs 请求包 Payloads 中的 Ext 键值下发，Ext 的协议格式由业务方确定，APNs 只做透传。
	Ext any `json:"Ext,omitempty"`
	//自定义 APNs 推送铃音。
	Sound string `json:"Sound,omitempty"`
}

func (c CustomContent) GetMsgType() string {
	return "TIMCustomElem"
}
func (c *CustomContent) GetData() any {
	var err error
	if _, ok := c.Data.(string); !ok {
		dd, err1 := json.Marshal(c.Data)
		if err1 != nil {
			err = errors.Join(err, err1)
		}
		c.Data = string(dd)
	}
	if _, ok := c.Ext.(string); !ok {
		dd, err2 := json.Marshal(c.Ext)
		if err2 != nil {
			err = errors.Join(err, err2)
		}
		c.Ext = string(dd)
	}
	if err != nil {
		slog.Error("CustomContent Get Data error", "error", err.Error(), "data", c.Data, "ext", c.Ext)
	}
	return c
}

// 声音类型
type SoundContent struct {
	//语音下载地址，可通过该 URL 地址直接下载相应语音。
	URL string `json:"Url"`
	//语音的唯一标识，客户端用于索引语音的键值。
	UUID string `json:"UUID"`
	//语音数据大小，单位：字节。
	Size int `json:"Size,omitempty"`
	//语音时长，单位：秒。
	Second int `json:"Second,omitempty"`
	//语音下载方式标记。目前 Download_Flag 取值只能为2，表示可通过Url字段值的 URL 地址直接下载语音
	DownloadFlag int `json:"Download_Flag"`
}

func (c SoundContent) GetMsgType() string {
	return "TIMSoundElem"
}
func (c SoundContent) GetData() any {
	return c
}

// 图片类型
type ImageContent struct {
	//图片的唯一标识，客户端用于索引图片的键值。
	UUID string `json:"UUID"`
	//图片格式。
	//JPG = 1
	//GIF = 2
	//PNG = 3
	//BMP = 4
	//其他 = 255
	ImageFormat int `json:"ImageFormat,omitempty"`
	//原图、缩略图或者大图下载信息。
	ImageInfoArray []ImageInfoArray `json:"ImageInfoArray"`
}
type ImageInfoArray struct {
	//图片类型：
	//1-原图
	//2-大图
	//3-缩略图
	Type int `json:"Type"`
	//图片数据大小，单位：字节。
	Size int `json:"Size,omitempty"`
	//图片宽度，单位为像素。
	Width int `json:"Width"`
	//图片高度，单位为像素。
	Height int `json:"Height"`
	//图片下载地址。
	URL string `json:"URL"`
}

func (c ImageContent) GetMsgType() string {
	return "TIMImageElem"
}
func (c ImageContent) GetData() any {
	return c
}

// 视频类型
type VideoContent struct {
	//视频下载地址。可通过该 URL 地址直接下载相应视频。
	VideoURL string `json:"VideoUrl"`
	//视频的唯一标识，客户端用于索引视频的键值。
	VideoUUID string `json:"VideoUUID"`
	//视频数据大小，单位：字节。
	VideoSize int `json:"VideoSize,omitempty"`
	//视频时长，单位：秒。Web 端不支持获取视频时长，值为0。
	VideoSecond int `json:"VideoSecond,omitempty"`
	//视频格式，例如 mp4。
	VideoFormat string `json:"VideoFormat,omitempty"`
	//视频下载方式标记。目前 VideoDownloadFlag 取值只能为2，表示可通过 VideoUrl 字段值的 URL 地址直接下载视频。
	VideoDownloadFlag int `json:"VideoDownloadFlag"`
	//视频缩略图下载地址。可通过该 URL 地址直接下载相应视频缩略图。
	ThumbURL string `json:"ThumbUrl"`
	//视频缩略图的唯一标识，客户端用于索引视频缩略图的键值。
	ThumbUUID string `json:"ThumbUUID"`
	//缩略图大小，单位：字节。
	ThumbSize int `json:"ThumbSize,omitempty"`
	//缩略图宽度，单位为像素。
	ThumbWidth int `json:"ThumbWidth"`
	//缩略图高度，单位为像素。
	ThumbHeight int `json:"ThumbHeight"`
	//缩略图格式，例如 JPG、BMP 等。
	ThumbFormat string `json:"ThumbFormat,omitempty"`
	//视频缩略图下载方式标记。目前 ThumbDownloadFlag 取值只能为2，表示可通过 ThumbUrl 字段值的 URL 地址直接下载视频缩略图。
	ThumbDownloadFlag int `json:"ThumbDownloadFlag"`
}

func (c VideoContent) GetMsgType() string {
	return "TIMVideoFileElem"
}

func (c VideoContent) GetData() any {
	return c
}

type HistoryMessageList struct {
	FromAccount  string `json:"From_Account"`
	ToAccount    string `json:"To_Account"`
	MsgSeq       int    `json:"MsgSeq"`
	MsgRandom    int    `json:"MsgRandom"`
	MsgTimeStamp int    `json:"MsgTimeStamp"`
	MsgFlagBits  int    `json:"MsgFlagBits"`
	IsPeerRead   int    `json:"IsPeerRead"`
	MsgKey       string `json:"MsgKey"`
	MsgBody      []struct {
		MsgType    string `json:"MsgType"` // TIMTextElem 文本格式, TIMLocationElem 地理位置 格式, TIMFaceElem 表情格式, TIMCustomElem 自定义信息, TIMSoundElem 语音消息, TIMImageElem 图片内容, TIMVideoFileElem 视频文件
		MsgContent struct {
			TextContent
			CustomContent
			SoundContent
			ImageContent
			VideoContent
		} `json:"MsgContent"`
	} `json:"MsgBody"`
	CloudCustomData string `json:"CloudCustomData"`
}

// 查询历史信息
type HistoryMessage struct {
	ActionStatus string               `json:"ActionStatus"`
	ErrorInfo    string               `json:"ErrorInfo"`
	ErrorCode    int                  `json:"ErrorCode"`
	Complete     int                  `json:"Complete"`
	MsgCnt       int                  `json:"MsgCnt"`
	LastMsgTime  int                  `json:"LastMsgTime"`
	LastMsgKey   string               `json:"LastMsgKey"`
	MsgList      []HistoryMessageList `json:"MsgList"`
}
