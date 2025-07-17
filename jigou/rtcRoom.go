package jigou

import "github.com/preceeder/go/base"

// BroadcastMessage 推送广播消息
func (j Rtc) BroadcastMessage(ctx base.BaseContext, reqData RtcBroadcastMessageReq) (*RtcBroadcastMessageResponse, error) {
	resBody := &RtcBroadcastMessageResponse{}
	err := j.Get(ctx, j.GetAction("BroadcastMessage"), reqData, resBody)
	return resBody, err
}

// BarrageMessage 推送弹幕消息
func (j Rtc) BarrageMessage(ctx base.BaseContext, reqData RtcBarrageMessageReq) (*RtcBarrageMessageMessageResponse, error) {
	resBody := &RtcBarrageMessageMessageResponse{}
	err := j.Get(ctx, j.GetAction("BarrageMessage"), reqData, resBody)
	return resBody, err
}

// SendCustomCommand 发送自定义消息
func (j Rtc) SendCustomCommand(ctx base.BaseContext, reqData RtcCustomCommandMessageReq) (*RtcSendCustomCommandResponse, error) {
	resBody := &RtcSendCustomCommandResponse{}
	err := j.Get(ctx, j.GetAction("CustomMessage"), reqData, resBody)
	return resBody, err
}

// GetRoomNumbers 获取房间的人数
func (j Rtc) GetRoomNumbers(ctx base.BaseContext, reqData RtcRoomUserNumReq) (*RtcRoomNumbersResponse, error) {
	resBody := &RtcRoomNumbersResponse{}
	err := j.Get(ctx, j.GetAction("UserNumber"), reqData, resBody)
	return resBody, err
}

// QueryRoomUserStatus 查询房间内用户状态
func (j Rtc) QueryRoomUserStatus(ctx base.BaseContext, reqData RtcUserStatusReq) (*RtcUserStatusResponse, error) {
	resBody := &RtcUserStatusResponse{}
	err := j.Get(ctx, j.GetAction("UserStatus"), reqData, resBody)
	return resBody, err
}

// QueryRoomUserList 查询房间内用户列表
func (j Rtc) QueryRoomUserList(ctx base.BaseContext, reqData RtcRoomUserListReq) (*RtcUserListResponse, error) {
	resBody := &RtcUserListResponse{}
	err := j.Get(ctx, j.GetAction("UserList"), reqData, resBody)
	return resBody, err
}

// AddRoomStream 增加房间流
func (j Rtc) AddRoomStream(ctx base.BaseContext, reqData RtcAddStreamReq) (*PublicResponse, error) {
	resBody := &PublicResponse{}
	err := j.Get(ctx, j.GetAction("AddStream"), reqData, resBody)
	return resBody, err
}

// DeleteRoomStream 删除房间流
func (j Rtc) DeleteRoomStream(ctx base.BaseContext, reqData RtcDeleteStreamReq) (*PublicResponse, error) {
	resBody := &PublicResponse{}
	err := j.Get(ctx, j.GetAction("DeleteStream"), reqData, resBody)
	return resBody, err
}

func (j Rtc) QuerySimpleStreamList(ctx base.BaseContext, reqData RtcSimpleStreamListReq) (*RtcSimpleStreamListResponse, error) {
	resBody := &RtcSimpleStreamListResponse{}
	err := j.Get(ctx, j.GetAction("SimpleStreamList"), reqData, resBody)
	return resBody, err
}

// CloseRoom 关闭房间
func (j Rtc) CloseRoom(ctx base.BaseContext, reqData RtcCloseRoomReq) (*PublicResponse, error) {
	resBody := &PublicResponse{}
	err := j.Get(ctx, j.GetAction("CloseRoom"), reqData, resBody)
	return resBody, err
}

// KickOutRoomUser 踢出房间用户
func (j Rtc) KickOutRoomUser(ctx base.BaseContext, reqData RtcKickOutRoomUserReq) (*PublicResponse, error) {
	resBody := &PublicResponse{}
	err := j.Get(ctx, j.GetAction("KickoutUser"), reqData, resBody)
	return resBody, err
}
