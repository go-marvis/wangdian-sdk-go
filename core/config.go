package core

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Config struct {
	Sid     string
	Key     string
	Secret  string
	Salt    string
	BaseUrl string
	Version string

	Retry   int           // 重试次数
	Limiter *rate.Limiter // 请求限速

	HttpClient HttpClient
	ReqTimeout time.Duration

	Header     http.Header
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
