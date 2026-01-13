package core

import (
	"net/http"
	"net/url"
)

type ApiReq struct {
	HttpMethod  string
	Method      string
	Body        any
	QueryParams url.Values
}

func NewApiReq() *ApiReq {
	return &ApiReq{
		QueryParams: url.Values{},
	}
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
