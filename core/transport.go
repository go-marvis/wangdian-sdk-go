package core

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"log/slog"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"time"
)

const (
	defaultContentType = "application/json; charset=utf-8"
)

func NewHttpClient(config *Config) {
	if config.HttpClient == nil {
		if config.ReqTimeout == 0 {
			config.HttpClient = http.DefaultClient
		} else {
			config.HttpClient = &http.Client{Timeout: config.ReqTimeout}
		}
	}
}

func NewSerializer(config *Config) {
	if config.Serializer == nil {
		config.Serializer = &defaultSerializer{}
	}
}

func doSend(client HttpClient, req *http.Request) (*ApiResp, error) {

	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := readResponse(resp)
	if err != nil {
		return nil, err
	}

	return &ApiResp{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		RawBody:    body,
	}, nil
}

func readResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func Request(ctx context.Context, apiReq *ApiReq, config *Config, options ...ReqOptionFunc) (*ApiResp, error) {
	option := &ReqOption{}
	for _, optf := range options {
		optf(option)
	}

	body, err := config.Serializer.Marshal(apiReq.Body)
	if err != nil {
		return nil, err
	}

	var values = make(url.Values, 0)
	for key := range apiReq.QueryParams {
		values.Set(key, apiReq.QueryParams.Get(key))
	}

	if option.Pager != nil {
		values.Set("page_size", strconv.Itoa(option.Pager.PageSize))
		values.Set("page_no", strconv.Itoa(option.Pager.PageNo))
	}
	if option.CalcTotal {
		values.Set("calc_total", "1")
	} else {
		values.Set("calc_total", "0")
	}

	values.Set("method", apiReq.Method)
	values.Set("sid", config.Sid)
	values.Set("key", config.Key)
	values.Set("salt", config.Salt)
	values.Set("timestamp", fmt.Sprintf("%v", time.Now().Unix()-1325347200)) // 2012-01-01 00:00:00
	values.Set("v", config.Version)
	values.Set("body", string(body))
	values.Set("sign", sign(values, config.Secret))
	values.Del("body")

	url := config.BaseUrl + "?" + values.Encode()
	slog.Debug("request", "url", url)
	req, err := http.NewRequestWithContext(ctx, apiReq.HttpMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", defaultContentType)

	for k, vs := range option.Header {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}

	for k, vs := range config.Header {
		for _, v := range vs {
			req.Header.Add(k, v)
		}
	}

	return doSend(config.HttpClient, req)
}

func sign(params url.Values, secret string) string {
	var content string
	for _, key := range slices.Sorted(maps.Keys(params)) {
		content = content + key + params.Get(key)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(secret+content+secret)))
}
