package rongyun

import (
	"encoding/json"
	"github.com/rongcloud/server-sdk-go/v3/sdk"
)

// rcMsg rcMsg接口
type RcMsg interface {
	ToString() (string, error)
}

type CustomMsg struct {
	Type    string          `json:"type,omitempty"`
	User    sdk.MsgUserInfo `json:"user,omitempty"`
	Content any             `json:"content"`
	Extra   any             `json:"extra,omitempty"`
}

func (c CustomMsg) ToString() (string, error) {
	msg, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(msg), nil
}
