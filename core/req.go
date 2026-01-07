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

type Pager struct {
	PageNo   int
	PageSize int
}

type ReqOption struct {
	Header    http.Header
	Pager     *Pager
	CalcTotal bool
}

type ReqOptionFunc func(option *ReqOption)

func WithHeader(header http.Header) ReqOptionFunc {
	return func(option *ReqOption) {
		option.Header = header
	}
}

func WithPager(pageNo, pageSize int) ReqOptionFunc {
	return func(option *ReqOption) {
		option.Pager = &Pager{
			PageNo:   pageNo,
			PageSize: pageSize,
		}
	}
}

func WithCalcTotal(calcTotal bool) ReqOptionFunc {
	return func(option *ReqOption) {
		option.CalcTotal = calcTotal
	}
}
