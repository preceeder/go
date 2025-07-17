package ding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/fanjindong/go-cache"
	"github.com/go-resty/resty/v2"
	"github.com/preceeder/base-aliyun/common"
	"log/slog"
	"net/url"
	"strconv"
	"time"
)

type DingConfig struct {
	AppKey      string `json:"appKey"`      // dingding 应用的 appkey
	AppSecret   string `json:"appSecret"`   // dingding 应用的 appSecret
	Secret      string `json:"secret"`      // 钉钉自定义机器人的 secret
	AccessToken string `json:"accessToken"` // 钉钉自定义机器人的 accessToken
}

type DingDing struct {
	HttpClient *resty.Client
	Config     DingConfig
	Cache      cache.ICache
}

func NewDingDing(config DingConfig) *DingDing {
	return &DingDing{
		HttpClient: resty.New(),
		Config:     config,
		Cache:      cache.NewMemCache(),
	}
}

func (d *DingDing) GetAccessToken() string {
	key := "accessToken"
	accessToken, ok := d.Cache.Get(key)
	if ok {
		return accessToken.(string)
	}
	uri := "https://oapi.dingtalk.com/gettoken"
	query := map[string]string{
		"appkey":    d.Config.AppKey,
		"appsecret": d.Config.AppSecret,
	}
	res := GetToken{}
	response := d.SendRequest(func(request *resty.Request) (*resty.Response, error) {
		return request.SetQueryParams(query).SetResult(&res).Get(uri)
	})
	if response.StatusCode() != 200 {
		return ""
	}

	d.Cache.Set(key, res.AccessToken, cache.WithEx(time.Second*time.Duration(res.ExpiresIn)))
	return res.AccessToken
}

// 获取表格属性
func (d *DingDing) ExcelAttribute(workbookId string, sheetId string) *ExcelAttribute {
	// 获取管理员信息
	unionid := d.GetOneAdminUserUnionId()
	if unionid != nil {
		headers := map[string]string{
			"Content-Type":                "application/json",
			"x-acs-dingtalk-access-token": d.GetAccessToken(),
		}
		res := ExcelAttribute{}
		uri := "https://api.dingtalk.com/v1.0/doc/workbooks/%s/sheets/%s?operatorId=%s"
		uri = fmt.Sprintf(uri, workbookId, sheetId, unionid)

		//uri := "https://api.dingtalk.com/v1.0/doc/workbooks/{{workbookId}}/sheets/{{sheetId}}?operatorId={{operatorId}}"
		//uri, _ = utils.StrBindName(uri, map[string]any{"workbookId": workbookId, "sheetId": sheetId, "operatorId": unionid}, []byte(""))
		httpRes := d.SendRequest(func(request *resty.Request) (*resty.Response, error) {
			return request.SetHeaders(headers).SetResult(&res).Get(uri)
		})
		if httpRes.StatusCode() != 200 {
			slog.Error("", "error", "", "code", httpRes.StatusCode())
		}

		return &res
	}

	return nil
}

// 获取表格数据
func (d *DingDing) GetExcelData(workbookId string, sheetId string, rangeStartAlpha string, rangeStartNum string, rangeEndAlpha string, rangeEndNum string) *ExcelData {
	// 获取表格的行数, 列数
	unionid := d.GetOneAdminUserUnionId()
	if unionid != nil {
		headers := map[string]string{
			"Content-Type":                "application/json",
			"x-acs-dingtalk-access-token": d.GetAccessToken(),
		}
		res := ExcelData{}
		ranges := d.ExcelRangeHandler(workbookId, sheetId, rangeStartAlpha, rangeStartNum, rangeEndAlpha, rangeEndNum)
		//uri := "https://api.dingtalk.com/v1.0/doc/workbooks/{{workbookId}}/sheets/{{sheetId}}/ranges/{{ranges}}"
		//uri, _ = utils.StrBindName(uri, map[string]any{"workbookId": workbookId, "sheetId": sheetId, "ranges": ranges}, []byte(""))

		uri := "https://api.dingtalk.com/v1.0/doc/workbooks/%s/sheets/%s/ranges/%s"
		uri = fmt.Sprintf(uri, workbookId, sheetId, ranges)

		query := map[string]string{
			"select":     "displayValues", // "values,formulas,displayValues,backgroundColors"
			"operatorId": *unionid,
		}
		d.SendRequest(func(request *resty.Request) (*resty.Response, error) {
			return request.SetHeaders(headers).SetQueryParams(query).SetResult(&res).Get(uri)
		})
		return &res
	}
	return nil
}

// 处理获取表格的范围
func (d *DingDing) ExcelRangeHandler(workbookId string, sheetId string, rangeStartAlpha string, rangeStartNum string, rangeEndAlpha string, rangeEndNum string) string {
	ranges := ""
	tableAttr := d.ExcelAttribute(workbookId, sheetId)
	lastRow := tableAttr.LastNonEmptyRow + 1 // 从0开始的
	lastClom := tableAttr.LastNonEmptyColumn // 从0开始的; 26 进制的 A = 0   需要转发为对应的字符串 AAA
	if rangeStartAlpha != "" {
		ranges += rangeStartAlpha
	} else {
		ranges += "A"
	}
	if rangeStartNum != "" {
		ranges += rangeStartNum
	} else {
		ranges += "1"
	}
	ranges += ":"

	if rangeEndAlpha != "" {
		ranges += rangeEndAlpha
	} else {
		ranges += common.BaseTo26(lastClom)
	}

	if rangeEndNum != "" {
		ranges += rangeEndNum
	} else {
		ranges += strconv.FormatInt(int64(lastRow), 10)
	}
	return ranges
}

// 获取一个管理用户的 unionid
func (d *DingDing) GetOneAdminUserUnionId() *string {
	adminList := d.GetAdminUsers()
	if adminList != nil {
		userId := adminList.Result[0].Userid
		userInfo := d.GetuserInfo(userId)
		if userInfo != nil {
			return &userInfo.Result.Unionid
		}
	}
	return nil
}

// 获取管理员信息
func (d *DingDing) GetAdminUsers() *AmdinUsersList {
	key := "adminUserIds"
	userdis, ok := d.Cache.Get(key)
	if ok {
		return userdis.(*AmdinUsersList)
	}
	uri := "https://oapi.dingtalk.com/topapi/user/listadmin?access_token=" + d.GetAccessToken()

	res := AmdinUsersList{}
	d.SendRequest(func(request *resty.Request) (*resty.Response, error) {
		return request.SetResult(&res).Post(uri)
	})
	if res.Errcode == 0 {
		d.Cache.Set(key, &res, cache.WithEx(time.Second*7200))
		return &res
	}
	return nil
}

// 获取用户的unionid
func (d *DingDing) GetuserInfo(userId string) *UserInfo {
	key := "userInfo:" + userId
	userInfo, ok := d.Cache.Get(key)
	if ok {
		return userInfo.(*UserInfo)
	}
	uri := "https://oapi.dingtalk.com/topapi/v2/user/get?access_token=" + d.GetAccessToken()
	res := UserInfo{}
	body := map[string]any{
		"userid":   userId,
		"language": "zh_CN",
	}
	d.SendRequest(func(request *resty.Request) (*resty.Response, error) {
		return request.SetBody(body).SetResult(&res).Post(uri)
	})
	if res.Errcode == 0 {
		d.Cache.Set(key, &res, cache.WithEx(time.Second*7200))
		return &res
	}
	return nil
}

// 发送机器人信息
func (d *DingDing) JiqirenSendMessage(data any) {
	// 机器人
	uri := "https://oapi.dingtalk.com/robot/send"
	query := map[string]string{
		"access_token": d.Config.AccessToken,
	}
	if len(d.Config.Secret) > 0 {
		timestamp := time.Now().UnixMilli()
		signStr := fmt.Sprintf("%d\n%s", timestamp, d.Config.Secret)
		decodeString := base64.StdEncoding.EncodeToString(HmacSha256(d.Config.Secret, signStr))
		sign := url.QueryEscape(decodeString)
		query["timestamp"] = strconv.FormatInt(timestamp, 10)
		query["sign"] = sign
	}

	res := d.SendRequest(func(request *resty.Request) (*resty.Response, error) {
		return request.SetHeaders(map[string]string{"Content-Type": "application/json"}).
			SetQueryParams(query).
			SetBody(data).Post(uri)
	})
	slog.Info("钉钉机器人发送消息", "res", res.String())
}

func (d *DingDing) SendRequest(f func(request *resty.Request) (*resty.Response, error)) *resty.Response {
	resuest := d.HttpClient.R()
	res, err := f(resuest)
	if err != nil {
		fmt.Println("钉钉请求错误 ", err.Error())
		fmt.Println("url:", res.Request.URL)
		return nil
	}
	if res.StatusCode() != 200 {
		slog.Error(res.Request.URL, "code", res.StatusCode(), "response", res.String())
	}
	return res
}

func HmacSha256(key string, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))
	return mac.Sum(nil)
}
