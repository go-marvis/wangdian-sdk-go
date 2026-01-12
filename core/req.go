package core

import (
	"net/http"
	"net/url"
)

type ApiReq struct {
	HttpMethod  string
	Method      string
	Body        any
	QueryParams QueryParams
}

type QueryParams url.Values

func (u QueryParams) Set(key, value string) {
	u[key] = []string{value}
}

func (u QueryParams) Get(key string) string {
	vs := u[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

type ReqOption struct {
	Header http.Header
}

type ReqOptionFunc func(option *ReqOption)

func WithHeader(header http.Header) ReqOptionFunc {
	return func(option *ReqOption) {
		option.Header = header
	}
}
