package shumei

// ---------------   请求参数 --------------
type ShumeiImage struct {
	ImageUrl       string
	UserId         string
	ReceiveTokenId string
	MType          string // 检查类型
	BusinessType   string // 业务检查类型
	Lang           string
	Ip             string
	EventId        string //  默认 IMAGE
	ThroughParams  map[string]any
	NeedCallBack   bool
	CallBaskUrl    string
}

// 同步
type ShumeiMultiImage struct {
	ImageUrl       []string
	UserId         string
	ReceiveTokenId string
	MType          string
	EventId        string //  默认 IMAGE
	Lang           string
	Ip             string
	ThroughParams  map[string]any
	NeedCallBack   bool
	CallBaskUrl    string
}

type ShumeiText struct {
	Text           string
	UserId         string
	ReceiveTokenId string
	MType          string
	Lang           string
	EventId        string
	Ip             string
	DeviceId       string
}

type ShumeiVoiceFile struct {
	VoiceUrl       string
	UserId         string
	ReceiveTokenId string
	MType          string
	EventId        string
	CallbackUrl    string         // 异步回调需要
	Lang           string         // 异步回调需要
	CallbackParams map[string]any // 异步回调需要
}

// 只有异步
type ShumeiAsyncVideoFile struct {
	VideoUrl       string
	UserId         string
	ReceiveTokenId string
	VideoType      string
	VoiceType      string
	EventId        string
	Lang           string
	CallBackUrl    string
	ThroughParams  map[string]any
}

// 只有异步
type ShumeiAsyncAudioStream struct {
	RtcParams        map[string]any
	StreamType       string // 目前默认值 ZEGO
	UserId           string
	ReceiveTokenId   string
	VoiceType        string // 检测的风险类型
	BusinessType     string // 业务标签
	EventId          string
	Callback         string
	Lang             string
	AudioDetectStep  int
	RoomId           string
	ReturnAllText    int // 0：返回风险等级为非pass的音频片段  1：返回所有风险等级的音频片段   默认0
	ReturnFinishInfo int
	ThroughParams    map[string]any
}

// 只有异步
type ShumeiAsyncVideoStream struct {
	UserId            string
	ReceiveTokenId    string
	VideoType         string `json:"imgType"` // 风险类型
	VoiceType         string // 风险类型
	ImgBusinessType   string // 业务类型
	AudioBusinessType string // 业务类型
	EventId           string
	ImgCallback       string // 视频流只检查 画面
	AudioCallback     string // 音频画面
	ReturnAllImg      int
	ReturnAllText     int

	//Callback       string
	ReturnFinishInfo int
	Lang             string
	RtcParams        map[string]any
	StreamType       string // 目前默认值 ZEGO
	RoomId           string
	DetectFrequency  int // 检测频次
	DetectStep       int // 视频流截帧图片检测步长。已截帧图片每个步长只会检测一次，取值大于等于1。不使用该功能时所有截帧全部过审
	// 检测频次 和 视频流截帧图片检测步长 的关系如下
	//10秒一个截帧片段    截帧步长设置为2
	//检测截帧片段为：0 10  20  30  40  50  60
	//截帧步长检测为：0   30  60
	//截帧步长会跳过两个10秒检测一次,  回调也会跳过两个10秒 才会有回调
	ImgBusinessDetectStep int // 图片业务标签检测步长。每个步长只会检测一次imgBusinessType，取值大于等于1。默认值=1，代表所有片段都审核业务标签。
	ThroughParams         map[string]any
}

// ---------------- 响应数据 --------------

// 共长返回数据

type AllLabels struct {
	Probability     float64        `json:"probability"`
	RiskDescription string         `json:"riskDescription"`
	RiskDetail      map[string]any `json:"riskDetail"`
	RiskLabel1      string         `json:"riskLabel1"`
	RiskLabel2      string         `json:"riskLabel2"`
	RiskLabel3      string         `json:"riskLabel3"`
	RiskLevel       string         `json:"riskLevel"`
}
type BusinessLabels struct {
	BusinessDescription string         `json:"businessDescription"`
	BusinessDetail      map[string]any `json:"businessDetail"`
	BusinessLabel1      string         `json:"businessLabel1"`
	BusinessLabel2      string         `json:"businessLabel2"`
	BusinessLabel3      string         `json:"businessLabel3"`
	ConfidenceLevel     int            `json:"confidenceLevel"`
	Probability         float64        `json:"probability"`
}
type PublicLongResponse struct {
	RequestID       string           `json:"requestId"`
	Code            int              `json:"code"`
	Message         string           `json:"message"`
	RiskLevel       string           `json:"riskLevel"`
	RiskLabel1      string           `json:"riskLabel1"`
	RiskLabel2      string           `json:"riskLabel2"`
	RiskLabel3      string           `json:"riskLabel3"`
	RiskDescription string           `json:"riskDescription"`
	RiskDetail      map[string]any   `json:"riskDetail"`
	AuxInfo         map[string]any   `json:"auxInfo"`
	AllLabels       []AllLabels      `json:"allLabels"`
	BusinessLabels  []BusinessLabels `json:"businessLabels"`
	TokenLabels     map[string]any   `json:"tokenLabels"`
}

// 公共短返回数据
type PublicShortResponse struct {
	RequestID string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

type VideoFileResponse struct {
	RequestID string `json:"requestId"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	BtId      string `json:"btId"`
}

type AudioStreamResponse struct {
	RequestID string                    `json:"requestId"`
	Code      int                       `json:"code"`
	Message   string                    `json:"message"`
	Detail    AudioStreamResponseDetail `json:"detail"`
}

type AudioStreamResponseDetail struct {
	Errorcode    int    `json:"errorcode`
	DupRequestId string `json:"dupRequestId"`
}

// 语音文件 检查同步返回数据
type VoiceFileResponse struct {
	Code      int             `json:"code"`
	Message   string          `json:"message"`
	RequestID string          `json:"requestId"`
	BtID      string          `json:"btId"`
	Detail    VoiceFileDetail `json:"detail"`
}
type VoiceFileDetail struct {
	AudioDetail   []map[string]any `json:"audioDetail"`
	AudioTags     map[string]any   `json:"audioTags"`
	AudioText     string           `json:"audioText"`
	AudioTime     int              `json:"audioTime"`
	Code          int              `json:"code"`
	RequestParams map[string]any   `json:"requestParams"`
	RiskLevel     string           `json:"riskLevel"`
}

// 流关闭回调
type CloseStreamResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

type TianWangResponse struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	RequestID string         `json:"requestId"`
	RiskLevel string         `json:"riskLevel"`
	Detail    map[string]any `json:"detail"`
}

type TianWangParams struct {
	EventId        string `json:"eventId"`
	TokenId        string `json:"tokenId"`
	Ip             string `json:"ip"`
	SmDeviceId     string `json:"smDeviceId"`
	Phone          string `json:"phone"`
	Channel        string `json:"channel"`
	Version        string `json:"version"`
	RegisterMethod string `json:"registerMethod"`
	UserAgent      string `json:"userAgent"`
}

// 音频流请求数据
type voiceStreamRequest struct {
	AccessKey    string `json:"accessKey"`
	AppId        string `json:"appId"`
	EventId      string `json:"eventId"`
	Type         string `json:"type"`
	BusinessType string `json:"businessType,omitempty"`
	Callback     string `json:"callback"`
	Data         any    `json:"data"`
}

// 视频流请求数据
type videoStreamRequest struct {
	AccessKey         string `json:"accessKey"`
	AppId             string `json:"appId"`
	EventId           string `json:"eventId"`
	ImgType           string `json:"imgType"`
	AudioType         string `json:"audioType"`
	ImgCallback       string `json:"imgCallback,omitempty"`
	AudioCallback     string `json:"audioCallback,omitempty"`
	Data              any    `json:"data"`
	ImgBusinessType   string `json:"imgBusinessType,omitempty"`
	AudioBusinessType string `json:"audioBusinessType,omitempty"`
}
