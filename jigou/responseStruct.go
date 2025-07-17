package jigou

type PublicResponse struct {
	Code      int    `json:"Code"`
	Message   string `json:"Message"`
	RequestID string `json:"RequestId"`
}

type RtcBroadcastMessageResponse struct {
	Code int `json:"Code"`
	Data struct {
		MessageId int `json:"MessageId"`
	} `json:"Data"`
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
}

type RtcBarrageMessageMessageResponse struct {
	Code int `json:"Code"`
	Data struct {
		MessageId int `json:"MessageId"`
	} `json:"Data"`
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
}

type RtcUserCountList struct {
	RoomID         string `json:"RoomId"`
	UserCount      int    `json:"UserCount"`
	AdminUserCount int    `json:"AdminUserCount"`
}
type RtcDataUserCountList struct {
	UserCountList []RtcUserCountList `json:"UserCountList"`
}

type RtcRoomNumbersResponse struct {
	Code      int                  `json:"Code"`
	Message   string               `json:"Message"`
	RequestID string               `json:"RequestId"`
	Data      RtcDataUserCountList `json:"Data"`
}

type RtcFailUsers struct {
	UID  string `json:"Uid"`
	Code int    `json:"Code"`
}
type RtcDataFailUsers struct {
	FailUsers []RtcFailUsers `json:"FailUsers"`
}
type RtcSendCustomCommandResponse struct {
	Code      int              `json:"Code"`
	Message   string           `json:"Message"`
	RequestID string           `json:"RequestId"`
	Data      RtcDataFailUsers `json:"Data"`
}

// RtcUserStatusResponse 查询用户状态
type RtcUserStatusResponse struct {
	Code int `json:"Code"`
	Data struct {
		UserStatusList []struct {
			UserId    string `json:"UserId"`
			Status    int    `json:"Status"`
			LoginTime int64  `json:"LoginTime"`
			UserRole  int    `json:"UserRole"`
		} `json:"UserStatusList"`
	} `json:"Data"`
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
}

// RtcUserListResponse 查询房间内用户列表
type RtcUserListResponse struct {
	Code      int    `json:"Code"`
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
	Data      struct {
		Marker   string `json:"Marker"`
		UserList []struct {
			UserId   string `json:"UserId"`
			UserName string `json:"UserName"`
			UserRole int    `json:"UserRole"` // 用户角色。 1：主播。2：观众。4：管理员;  该返回参数，仅在接入 LiveRoom 服务时有实际意义，接入 Express 服务时请忽略此参数
		} `json:"UserList"`
	} `json:"Data"`
}

// RtcSimpleStreamListResponse 获取房间内简易流列表
type RtcSimpleStreamListResponse struct {
	Code      int    `json:"Code"`
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
	Data      struct {
		StreamList []struct {
			StreamId       string `json:"StreamId"`
			UserId         string `json:"UserId"`
			UserName       string `json:"UserName"`
			CreateTime     int64  `json:"CreateTime"`
			StreamNumberId int    `json:"StreamNumberId"`
		} `json:"StreamList"`
	} `json:"Data"`
}

type GenerateIdentifyToken struct {
	Code      int         `json:"Code"`
	Data      IdentiToken `json:"Data"`
	Message   string      `json:"Message"`
	RequestID string      `json:"RequestId"`
}
type IdentiToken struct {
	IdentifyToken string `json:"IdentifyToken"`
	RemainTime    int    `json:"RemainTime"`
}

// token业务扩展：权限认证属性
type RtcRoomPayLoad struct {
	RoomId       string      `json:"room_id"`        //房间 id（必填）；用于对接口的房间 id 进行强验证
	Privilege    map[int]int `json:"privilege"`      //权限位开关列表；用于对接口的操作权限进行强验证
	StreamIdList []string    `json:"stream_id_list"` //流列表；用于对接口的流 id 进行强验证；允许为空，如果为空，则不对流 id 验证
}
