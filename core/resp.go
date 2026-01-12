package core

import (
	"net/http"
)

type ApiResp struct {
	StatusCode int         `json:"-"`
	Header     http.Header `json:"-"`
	RawBody    []byte      `json:"-"`
}

func (resp ApiResp) UnmarshalBody(val any, config *Config) error {
	return config.Serializer.Unmarshal(resp.RawBody, val)
}

type CodeError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
