package jigou

import (
	"log/slog"
	"net/url"
)

// QueryParamsFromValues 将结构体转化为 queryUrlValues 类型
func QueryParamsFromValues(r any) (url.Values, error) {
	requestData, err := StructToUrlValues(r)
	if err != nil {
		slog.Error("QueryParamsFromValues", "error", err.Error(), "data", r)
		return nil, err
	}
	return requestData, nil
}

// RtcBroadcastMessageReq 推送广播消息
type RtcBroadcastMessageReq struct {
	RoomId          string `url:"RoomId"`             // 房间 ID。
	UserId          string `url:"UserId"`             // 发送方用户 ID。
	UserName        string `url:"UserName,omitempty"` // 发送方用户名（与 UserId 一一对应）。  为空时可以不要
	MessageCategory uint32 `url:"MessageCategory"`    // 消息分类, 1 系统消息; 2 聊天消息
	MessageContent  string `url:"MessageContent"`     // 消息内容，长度不能超过 1024 个字节。
}

// RtcBarrageMessageReq 推送弹幕消息
type RtcBarrageMessageReq struct {
	RoomId          string `url:"RoomId"`             // 房间 ID。
	UserId          string `url:"UserId"`             // 发送方用户 ID。
	UserName        string `url:"UserName,omitempty"` // 发送方用户名（与 UserId 一一对应）。  为空时可以不要
	MessageCategory uint32 `url:"MessageCategory"`    // 消息分类, 1 系统消息; 2 聊天消息
	MessageContent  string `url:"MessageContent"`     // 消息内容，长度不能超过 1024 个字节。
}

// RtcCustomCommandMessageReq 发送自定义消息
type RtcCustomCommandMessageReq struct {
	RoomId         string   `url:"RoomId"`               // 房间 ID。
	FromUserId     string   `url:"FromUserId"`           // 发送方用户 ID。
	ToUserId       []string `url:"ToUserId[],omitempty"` // 不填写就是发送给房间内的所有用户, 填写了就只发送给指定用户
	MessageContent string   `url:"MessageContent"`       // 自定义消息内容，长度不能超过 1024 个字节。
}

// RtcRoomUserNumReq 获取房间人数
type RtcRoomUserNumReq struct {
	RoomId []string `url:"RoomId[]"` // 最多10个房间号
}

// RtcUserStatusReq  查询用户状态
type RtcUserStatusReq struct {
	RoomId string   `url:"RoomId"`
	UserId []string `url:"UserId[]"` // 需要查询状态的用户 ID 列表，最大支持 10 个用户 ID。
}

// RtcRoomUserListReq 查询房间用户列表
type RtcRoomUserListReq struct {
	RoomId string `url:"RoomId"`
	Mode   int32  `url:"Mode"`             //用户登录房间的时间排序，默认值为 0。 0：按时间正序; 1：按时间倒序
	Limit  int32  `url:"Limit"`            //单次请求返回的用户个数，取值范围 0-200
	Marker string `url:"Marker,omitempty"` //查询用户起始位标识，每次请求的响应有返回，为空时从头开始返回用户信息。
}

// RtcAddStreamReq 增加房间流
type RtcAddStreamReq struct {
	RoomId      string `url:"RoomId"`
	UserId      string `url:"UserId"`
	UserName    string `url:"UserName,omitempty"`
	StreamId    string `url:"StreamId"`              // 流 ID，不超过 256 字节
	StreamTitle string `url:"StreamTitle,omitempty"` // 流标题，不超过 127 字节。
	ExtraInfo   string `url:"ExtraInfo,omitempty"`   // 流附加信息，不超过 1024 字节。

}

// RtcDeleteStreamReq 删除流
type RtcDeleteStreamReq struct {
	RoomId   string `url:"RoomId"`
	UserId   string `url:"UserId"`
	UserName string `url:"UserName,omitempty"`
	StreamId string `url:"StreamId"`
}

// RtcSimpleStreamListReq 获取房间人数
type RtcSimpleStreamListReq struct {
	RoomId []string `url:"RoomId[]"` // 最多10个房间号
}

// RtcCloseRoomReq 关闭房间
type RtcCloseRoomReq struct {
	RoomId            string `url:"RoomId"`
	CustomReason      string `url:"CustomReason,omitempty"`      // 关闭原因，最大长度为 256 字节
	RoomCloseCallback bool   `url:"RoomCloseCallback,omitempty"` // 是否产生 房间关闭回调，默认为 false
}

// RtcKickOutRoomUserReq 踢出房间用户
type RtcKickOutRoomUserReq struct {
	RoomId       string   `url:"RoomId"`
	UserId       []string `url:"UserId[]"`     // 踢出房间的用户 ID 列表，最大支持 5 个用户 ID。
	CustomReason string   `url:"CustomReason"` // 踢人原因，最大长度为 256 字节。
}
