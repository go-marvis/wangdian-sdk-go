package core

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Config struct {
	Sid     string // 卖家账号，旺店通分配
	Key     string // 接口账号，旺店通分配
	Secret  string
	Salt    string // 密钥, 加密请求
	BaseUrl string // 服务地址
	Version string // 版本号
	Debug   bool   // 调试模式, 输出请求和响应

	Retry   int           // 重试次数
	Limiter *rate.Limiter // 请求限速

	HttpClient HttpClient
	ReqTimeout time.Duration

	Serializer Serializer
}

func Init(config *Config) {
	if config.HttpClient == nil {
		if config.ReqTimeout == 0 {
			config.HttpClient = http.DefaultClient
		} else {
			config.HttpClient = &http.Client{Timeout: config.ReqTimeout}
		}
	}

	if config.Serializer == nil {
		config.Serializer = &defaultSerializer{}
	}
}
