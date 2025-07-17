package tencentyun

import (
	"github.com/preceeder/go/base"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	v20180301 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/faceid/v20180301"
	"log/slog"
)

type TencentIdentityClient struct {
	Client *v20180301.Client
	Config TencentIdentityConfig
}

// api 配置
type TencentIdentityConfig struct {
	SecretId  string `json:"secretId"`
	SecretKey string `json:"secretKey"`
	Endpoint  string `json:"endpoint"` // "faceid.tencentcloudapi.com"
}

func NewTencentIdentityClient(config TencentIdentityConfig) *TencentIdentityClient {
	credential := common.NewCredential(
		config.SecretId,
		config.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = config.Endpoint
	cpf.HttpProfile.ReqMethod = "POST"

	// 强制使用ipv4
	//ipv4Transport := &http.Transport{
	//	DialContext: (&net.Dialer{
	//		Timeout:   30 * time.Second,
	//		KeepAlive: 30 * time.Second,
	//		Resolver: &net.Resolver{
	//			PreferGo: true,
	//			Dial: func(ctx context.Context, _, address string) (net.Conn, error) {
	//				return net.Dial("tcp4", address)
	//			},
	//		},
	//	}).DialContext,
	//}
	//common.DefaultHttpClient = &http.Client{Transport: ipv4Transport}

	client, err := v20180301.NewClient(credential, regions.Shanghai, cpf)
	if err != nil {
		slog.Error("初始化 NewTencentIdentityClient", "error", err.Error())
	}
	return &TencentIdentityClient{Client: client, Config: config}
}

func (t *TencentIdentityClient) Identity(ctx base.BaseContext, name string, idCard string) (v20180301.IdCardVerificationResponseParams, error) {
	defer func() {
		if err := recover(); err != nil {
			slog.ErrorContext(ctx, "identity", "error", err)
		}
	}()

	//创建common client
	request := v20180301.NewIdCardVerificationRequest()
	request.IdCard = &idCard
	request.Name = &name
	res, err := t.Client.IdCardVerification(request)
	if err != nil {
		return v20180301.IdCardVerificationResponseParams{}, err
	}
	return *res.Response, err
}
