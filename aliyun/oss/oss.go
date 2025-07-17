package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log/slog"
	"os"
)

type AliOssConfig struct {
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
	EndPoint     string `json:"endPoint"`
	Region       string `json:"region"`
	BucketName   string `json:"bucketName"`
	Url          string `json:"url"`
	Cdn          string `json:"cdn"`
}

type AliOssClient struct {
	Config AliOssConfig
	Client *oss.Client
	Bucket *oss.Bucket
}

func NewAliOssClient(config AliOssConfig) AliOssClient {
	client, err := oss.New(config.EndPoint, config.AccessKey, config.AccessSecret, oss.Region(config.Region))
	if err != nil {
		slog.Error("阿里云push创建失败", "error", err.Error())
		os.Exit(-1)
	}
	slog.Info("创建oss 客户端")
	bucket, err := client.Bucket(config.BucketName)
	if err != nil {
		slog.Error("阿里云push创建失败", "error", err.Error())
		os.Exit(-1)
	}
	return AliOssClient{Config: config, Client: client, Bucket: bucket}
}
