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

	Retry   int
	Limiter *rate.Limiter

	HttpClient HttpClient
	ReqTimeout time.Duration

	Header     http.Header
	Serializer Serializer
}
