package core

import "net/http"

type ApiResp struct {
	StatusCode int
	Header     http.Header
	RawBody    []byte
}

type Response[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
