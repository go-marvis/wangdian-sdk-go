package core

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"time"
)

const (
	defaultContentType = "application/json; charset=utf-8"
)

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

	body, err := config.Serializer.Marshal([]any{apiReq.Body})
	if err != nil {
		return nil, err
	}

	var values = make(url.Values, 0)
	for key := range apiReq.QueryParams {
		values.Set(key, apiReq.QueryParams.Get(key))
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
	if config.Debug {
		payload := body
		buf := bytes.NewBuffer(make([]byte, 0, len(body)+1024))
		if err := json.Indent(buf, body, "", "  "); err == nil {
			payload = buf.Bytes()
		}
		log.Printf("[DEBUG] [API] %s %s payload:\n%s\n", apiReq.HttpMethod, url, payload)
	}

	req, err := http.NewRequestWithContext(ctx, apiReq.HttpMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", defaultContentType)

	var apiResp *ApiResp
	for range config.Retry {
		if config.Limiter != nil {
			if err = config.Limiter.Wait(ctx); err != nil {
				return nil, err
			}
		}

		apiResp, err = doSend(config.HttpClient, req)
		if apiResp.StatusCode != http.StatusOK {
			slog.Error("http error", "status", apiResp.StatusCode)
			return apiResp, fmt.Errorf("http error, status=%d", apiResp.StatusCode)
		}

		codeError := &CodeError{}
		if err = config.Serializer.Unmarshal(apiResp.RawBody, codeError); err != nil {
			slog.Error("unmarshal error", "err", err)
			return apiResp, err
		}

		if codeError.Status != 0 {
			slog.Error("api error", "status", codeError.Status, "message", codeError.Message)
			err = fmt.Errorf("api error, code=%v, err=%v", codeError.Status, codeError.Message)

			// retry
			if codeError.Status == 100 {
				if codeError.Message == "超过每分钟最大调用频率限制，请稍后重试" {
					time.Sleep(3 * time.Second)
					continue
				}
			}
		}

		err = nil
		break
	}
	return apiResp, nil
}

func sign(params url.Values, secret string) string {
	var content string
	for _, key := range slices.Sorted(maps.Keys(params)) {
		content = content + key + params.Get(key)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(secret+content+secret)))
}
